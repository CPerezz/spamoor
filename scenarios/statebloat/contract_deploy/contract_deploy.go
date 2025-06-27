package contractdeploy

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethpandaops/spamoor/scenario"
	"github.com/ethpandaops/spamoor/scenarios/statebloat/contract_deploy/contract"
	"github.com/ethpandaops/spamoor/spamoor"
	"github.com/ethpandaops/spamoor/txbuilder"
)

type ScenarioOptions struct {
	MaxWallets       uint64 `yaml:"max_wallets"`
	BaseFee          uint64 `yaml:"base_fee"`
	TipFee           uint64 `yaml:"tip_fee"`
	ClientGroup      string `yaml:"client_group"`
	MaxTransactions  uint64 `yaml:"max_transactions"`
	RateLimitPercent uint64 `yaml:"rate_limit_percent"`
}

// ContractDeployment tracks a deployed contract with its deployer info
type ContractDeployment struct {
	ContractAddress string `json:"contract_address"`
	PrivateKey      string `json:"private_key"`
}

// PendingTransaction tracks a transaction until it's mined
type PendingTransaction struct {
	TxHash     common.Hash
	PrivateKey *ecdsa.PrivateKey
	Timestamp  time.Time
	Sent       bool // Track if transaction was successfully sent
}

// BlockDeploymentStats tracks deployment statistics per block
type BlockDeploymentStats struct {
	BlockNumber       uint64
	ContractCount     int
	TotalGasUsed      uint64
	TotalBytecodeSize int
}

type Scenario struct {
	options    ScenarioOptions
	logger     *logrus.Entry
	walletPool *spamoor.WalletPool

	// Cached chain ID to avoid repeated RPC calls
	chainID      *big.Int
	chainIDOnce  sync.Once
	chainIDError error

	// Transaction tracking
	pendingTxs      map[common.Hash]*PendingTransaction
	pendingTxsMutex sync.RWMutex

	// Results tracking
	deployedContracts []ContractDeployment
	contractsMutex    sync.Mutex

	// Block-level statistics tracking
	blockStats      map[uint64]*BlockDeploymentStats
	blockStatsMutex sync.Mutex
	lastLoggedBlock uint64

	// Block monitoring for real-time logging
	blockMonitorCancel context.CancelFunc
	blockMonitorDone   chan struct{}

	// Wallet group management for rate limiting
	currentWalletGroup uint64
	walletsPerGroup    uint64

	// Original gas prices for escalation
	originalBaseFee uint64
	originalTipFee  uint64

	// Deadlock prevention counters
	consecutiveEmptyBlocks      uint64
	consecutiveFailedBlockWaits uint64
	emptyBlockRetryCount        uint64
	maxGasPriceReached          bool

	// Async file writing system to prevent blocking main loop
	fileWriterChan   chan struct{}
	fileWriterCancel context.CancelFunc
	fileWriterDone   chan struct{}

	// Channel-based communication between sender and monitor goroutines
	sentTxChan      chan *PendingTransaction // Sender -> Monitor: new transactions
	newBlockChan    chan uint64              // Monitor -> Sender: new block signal
	gasEscalateChan chan bool                // Monitor -> Sender: escalate gas prices

	// Stuck transaction replacement channels
	requestStuckTxChan  chan bool                                // Sender -> Monitor: request stuck txs
	responseStuckTxChan chan map[common.Hash]*PendingTransaction // Monitor -> Sender: stuck tx list
}

var ScenarioName = "contract-deploy"
var ScenarioDefaultOptions = ScenarioOptions{
	MaxWallets:       0, // Use root wallet only by default
	BaseFee:          5, // Moderate base fee (5 gwei)
	TipFee:           1, // Priority fee (1 gwei)
	MaxTransactions:  0,
	RateLimitPercent: 90, // Use 90% of block capacity by default
}
var ScenarioDescriptor = scenario.Descriptor{
	Name:           ScenarioName,
	Description:    "Deploy contracts to create state bloat",
	DefaultOptions: ScenarioDefaultOptions,
	NewScenario:    newScenario,
}

func newScenario(logger logrus.FieldLogger) scenario.Scenario {
	return &Scenario{
		logger:     logger.WithField("scenario", ScenarioName),
		pendingTxs: make(map[common.Hash]*PendingTransaction),
	}
}

func (s *Scenario) Flags(flags *pflag.FlagSet) error {
	flags.Uint64Var(&s.options.MaxWallets, "max-wallets", ScenarioDefaultOptions.MaxWallets, "Maximum number of child wallets to use")
	flags.Uint64Var(&s.options.BaseFee, "basefee", ScenarioDefaultOptions.BaseFee, "Base fee per gas in gwei")
	flags.Uint64Var(&s.options.TipFee, "tipfee", ScenarioDefaultOptions.TipFee, "Tip fee per gas in gwei")
	flags.StringVar(&s.options.ClientGroup, "client-group", ScenarioDefaultOptions.ClientGroup, "Client group to use for sending transactions")
	flags.Uint64Var(&s.options.MaxTransactions, "max-transactions", ScenarioDefaultOptions.MaxTransactions, "Maximum number of transactions to send (0 = use rate limiting based on block gas limit)")
	flags.Uint64Var(&s.options.RateLimitPercent, "rate-limit-percent", ScenarioDefaultOptions.RateLimitPercent, "Percentage of block capacity to use (1-100)")
	return nil
}

func (s *Scenario) Init(options *scenario.Options) error {
	s.walletPool = options.WalletPool

	if options.Config != "" {
		err := yaml.Unmarshal([]byte(options.Config), &s.options)
		if err != nil {
			return fmt.Errorf("failed to unmarshal config: %w", err)
		}
	}

	// For rate-limiting mode (MaxTransactions == 0), we need to calculate wallet count
	// based on block capacity to implement the alternating wallet group strategy
	if s.options.MaxTransactions == 0 && s.options.MaxWallets == 0 {
		// Get a client to fetch block info
		client := s.walletPool.GetClient(spamoor.SelectClientByIndex, 0, s.options.ClientGroup)
		if client == nil {
			return fmt.Errorf("no client available for initialization")
		}

		// Get current block to calculate capacity
		ctx := context.Background()
		ethClient := client.GetEthClient()
		currentBlock, err := ethClient.BlockByNumber(ctx, nil)
		if err != nil {
			return fmt.Errorf("failed to get current block: %w", err)
		}

		blockGasLimit := currentBlock.GasLimit()
		estimatedGasPerContract := uint64(4949468)
		maxContractsPerBlock := blockGasLimit / estimatedGasPerContract

		// Apply rate limit percentage to avoid trying to fill blocks completely
		rateLimitPercent := s.options.RateLimitPercent
		if rateLimitPercent == 0 || rateLimitPercent > 100 {
			rateLimitPercent = 90 // Default to 90% if not set
		}
		contractsPerGroup := (maxContractsPerBlock * rateLimitPercent) / 100

		// Create 10x the number of wallets we need per group
		// This allows us to cycle through 10 groups, giving each wallet more time between uses
		totalWallets := contractsPerGroup * 10
		s.walletsPerGroup = contractsPerGroup

		s.logger.Infof("Rate limiting mode: creating %d wallets (%d per group across 10 groups for %d%% of %d max contracts/block)",
			totalWallets, s.walletsPerGroup, rateLimitPercent, maxContractsPerBlock)

		s.walletPool.SetWalletCount(totalWallets)
	} else if s.options.MaxWallets > 0 {
		s.walletPool.SetWalletCount(s.options.MaxWallets)
	} else {
		// Default to a reasonable number of wallets to avoid root wallet usage
		s.walletPool.SetWalletCount(10)
	}

	// Force wallet preparation to ensure they exist before nonce reset
	err := s.walletPool.PrepareWallets()
	if err != nil {
		return fmt.Errorf("failed to prepare wallets: %w", err)
	}

	// Reset nonces for all child wallets to sync with blockchain state
	// This prevents conflicts with leftover pending transactions from previous runs
	s.logger.Info("Resetting wallet nonces to sync with blockchain state...")
	walletCount := s.walletPool.GetWalletCount()
	client := s.walletPool.GetClient(spamoor.SelectClientByIndex, 0, s.options.ClientGroup)
	if client == nil {
		return fmt.Errorf("no client available for nonce reset")
	}

	ctx := context.Background()
	for i := uint64(0); i < walletCount; i++ {
		wallet := s.walletPool.GetWallet(spamoor.SelectWalletByIndex, int(i))
		if wallet != nil {
			oldNonce := wallet.GetNonce()
			wallet.ResetPendingNonce(ctx, client)
			newNonce := wallet.GetNonce()
			if oldNonce != newNonce {
				s.logger.Infof("Reset wallet %d nonce from %d to %d", i, oldNonce, newNonce)
			}
		}
	}
	s.logger.Info("Completed wallet nonce reset")

	return nil
}

// getChainID caches the chain ID to avoid repeated RPC calls
func (s *Scenario) getChainID(ctx context.Context) (*big.Int, error) {
	s.chainIDOnce.Do(func() {
		client := s.walletPool.GetClient(spamoor.SelectClientByIndex, 0, s.options.ClientGroup)
		if client == nil {
			s.chainIDError = fmt.Errorf("no client available for chain ID")
			return
		}
		s.chainID, s.chainIDError = client.GetChainId(ctx)
	})
	return s.chainID, s.chainIDError
}

// DEPRECATED: processPendingTransactions - moved to block monitor goroutine
// This function is no longer used in the refactored architecture
func (s *Scenario) processPendingTransactionsOLD(ctx context.Context) {
	// Quick timeout to prevent blocking main loop
	timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	s.pendingTxsMutex.Lock()

	client := s.walletPool.GetClient(spamoor.SelectClientByIndex, 0, s.options.ClientGroup)
	if client == nil {
		s.pendingTxsMutex.Unlock()
		return
	}

	ethClient := client.GetEthClient()
	var confirmedTxs []common.Hash
	var timedOutTxs []common.Hash
	var stuckTxs []common.Hash
	var successfulDeployments []struct {
		ContractAddress common.Address
		PrivateKey      *ecdsa.PrivateKey
		Receipt         *types.Receipt
		TxHash          common.Hash
	}

	// Limit processing to prevent blocking - only check up to 20 transactions
	checkedCount := 0
	maxChecks := 20

	for txHash, pendingTx := range s.pendingTxs {
		if checkedCount >= maxChecks {
			break // Don't check all transactions to avoid blocking
		}
		checkedCount++

		timeSinceSent := time.Since(pendingTx.Timestamp)

		// Check if transaction is stuck (>30 seconds without confirmation)
		if timeSinceSent > 30*time.Second && timeSinceSent <= 1*time.Minute {
			stuckTxs = append(stuckTxs, txHash)
		}

		// Check if transaction is too old (1 minute timeout)
		if timeSinceSent > 1*time.Minute {
			s.logger.Warnf("Transaction %s timed out after 1 minute, removing from pending", txHash.Hex())
			timedOutTxs = append(timedOutTxs, txHash)
			continue
		}

		// Quick receipt check with timeout to prevent hanging
		receipt, err := ethClient.TransactionReceipt(timeoutCtx, txHash)
		if err != nil {
			// Transaction still pending or error retrieving receipt
			continue
		}

		confirmedTxs = append(confirmedTxs, txHash)

		// Process successful deployment
		if receipt.Status == 1 && receipt.ContractAddress != (common.Address{}) {
			successfulDeployments = append(successfulDeployments, struct {
				ContractAddress common.Address
				PrivateKey      *ecdsa.PrivateKey
				Receipt         *types.Receipt
				TxHash          common.Hash
			}{
				ContractAddress: receipt.ContractAddress,
				PrivateKey:      pendingTx.PrivateKey,
				Receipt:         receipt,
				TxHash:          txHash,
			})
		}
	}

	// Trigger gas price escalation if we have stuck transactions
	escalationNeeded := len(stuckTxs) > 0

	// Remove confirmed transactions from pending map
	for _, txHash := range confirmedTxs {
		delete(s.pendingTxs, txHash)
	}

	// Remove timed out transactions from pending map
	for _, txHash := range timedOutTxs {
		delete(s.pendingTxs, txHash)
	}

	s.pendingTxsMutex.Unlock()

	// Handle gas price escalation outside of mutex lock to prevent deadlock
	if escalationNeeded {
		s.logger.Warnf("Detected %d stuck transactions (>30s), escalating gas prices", len(stuckTxs))

		// Quick escalation without additional RPC calls to prevent blocking
		err := s.updateDynamicFeesWithEscalation(timeoutCtx, true)
		if err != nil {
			s.logger.Warnf("Failed to escalate gas prices: %v", err)
		} else {
			s.logger.Infof("Gas prices escalated successfully, will use new prices in next transaction batch")
		}
	}

	// Process successful deployments (now non-blocking due to async file writing)
	for _, deployment := range successfulDeployments {
		s.recordDeployedContract(deployment.ContractAddress, deployment.PrivateKey, deployment.Receipt, deployment.TxHash)
	}
}

// recordDeployedContract records a successfully deployed contract
func (s *Scenario) recordDeployedContract(contractAddress common.Address, privateKey *ecdsa.PrivateKey, receipt *types.Receipt, txHash common.Hash) {
	s.contractsMutex.Lock()
	defer s.contractsMutex.Unlock()

	// Keep the JSON structure simple - only contract address and private key
	deployment := ContractDeployment{
		ContractAddress: contractAddress.Hex(),
		PrivateKey:      fmt.Sprintf("0x%x", crypto.FromECDSA(privateKey)),
	}

	s.deployedContracts = append(s.deployedContracts, deployment)

	// Use constant bytecode size to avoid blocking RPC calls in main loop
	// This prevents the main transaction sending loop from hanging on RPC timeouts
	const bytecodeSize int = 23914 // Average deployed contract size

	blockNumber := receipt.BlockNumber.Uint64()

	// Debug logging for block tracking
	s.logger.WithFields(logrus.Fields{
		"tx_block":        blockNumber,
		"existing_blocks": len(s.blockStats),
	}).Debug("Recording contract deployment")

	// Update block-level statistics
	s.blockStatsMutex.Lock()
	defer s.blockStatsMutex.Unlock()

	if s.blockStats == nil {
		s.blockStats = make(map[uint64]*BlockDeploymentStats)
	}

	// Create or update current block stats (removed the old logging logic)
	if s.blockStats[blockNumber] == nil {
		s.blockStats[blockNumber] = &BlockDeploymentStats{
			BlockNumber: blockNumber,
		}
		s.logger.WithField("block_number", blockNumber).Debug("Created new block stats")
	}

	blockStat := s.blockStats[blockNumber]
	blockStat.ContractCount++
	blockStat.TotalGasUsed += receipt.GasUsed
	blockStat.TotalBytecodeSize += bytecodeSize

	s.logger.WithFields(logrus.Fields{
		"block_number":       blockNumber,
		"contracts_in_block": blockStat.ContractCount,
		"gas_used":           blockStat.TotalGasUsed,
		"bytecode_size":      blockStat.TotalBytecodeSize,
	}).Debug("Updated block stats")

	// Queue an async write to deployments.json (non-blocking)
	s.queueFileWrite()
}

// Helper function for max calculation
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// writeDeploymentsFile creates/updates deployments.json with private key to contract address mapping
func (s *Scenario) writeDeploymentsFile() error {
	// Create a map from private key to array of contract addresses
	deploymentMap := make(map[string][]string)

	for _, contract := range s.deployedContracts {
		privateKey := contract.PrivateKey
		contractAddr := contract.ContractAddress
		deploymentMap[privateKey] = append(deploymentMap[privateKey], contractAddr)
	}

	// Create or overwrite the deployments.json file
	deploymentsFile, err := os.Create("deployments.json")
	if err != nil {
		return fmt.Errorf("failed to create deployments.json file: %w", err)
	}
	defer deploymentsFile.Close()

	// Write the mapping as JSON with pretty formatting
	encoder := json.NewEncoder(deploymentsFile)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(deploymentMap)
	if err != nil {
		return fmt.Errorf("failed to write deployments.json: %w", err)
	}

	return nil
}

// queueFileWrite queues a deployments.json write request (non-blocking)
func (s *Scenario) queueFileWrite() {
	// Non-blocking send to prevent main loop from hanging
	select {
	case s.fileWriterChan <- struct{}{}:
		// Successfully queued
	default:
		// Channel full, skip this write request to prevent blocking
		s.logger.Debug("File writer channel full, skipping write request")
	}
}

// DEPRECATED: startBlockMonitor - replaced by runBlockMonitor goroutine
// This function is no longer used in the refactored architecture
func (s *Scenario) startBlockMonitorOLD(ctx context.Context) {
	monitorCtx, cancel := context.WithCancel(ctx)
	s.blockMonitorCancel = cancel
	s.blockMonitorDone = make(chan struct{})

	go func() {
		defer close(s.blockMonitorDone)

		client := s.walletPool.GetClient(spamoor.SelectClientByIndex, 0, s.options.ClientGroup)
		if client == nil {
			s.logger.Warn("No client available for block monitoring")
			return
		}

		ethClient := client.GetEthClient()
		ticker := time.NewTicker(2 * time.Second) // Poll every 2 seconds
		defer ticker.Stop()

		for {
			select {
			case <-monitorCtx.Done():
				return
			case <-ticker.C:
				// Get current block number
				latestBlock, err := ethClient.BlockByNumber(monitorCtx, nil)
				if err != nil {
					s.logger.WithError(err).Debug("Failed to get latest block for monitoring")
					continue
				}

				currentBlockNumber := latestBlock.Number().Uint64()

				// Log any completed blocks that haven't been logged yet
				s.blockStatsMutex.Lock()
				for bn := s.lastLoggedBlock + 1; bn < currentBlockNumber; bn++ {
					if stats, exists := s.blockStats[bn]; exists && stats.ContractCount > 0 {
						avgGasPerByte := float64(stats.TotalGasUsed) / float64(max(stats.TotalBytecodeSize, 1))

						s.contractsMutex.Lock()
						totalContracts := len(s.deployedContracts)
						s.contractsMutex.Unlock()

						// Calculate block utilization if we know the gas limit
						var utilizationStr string
						if latestBlock != nil && latestBlock.GasLimit() > 0 {
							utilization := float64(stats.TotalGasUsed) / float64(latestBlock.GasLimit()) * 100
							utilizationStr = fmt.Sprintf("%.1f%%", utilization)
						}

						s.logger.WithFields(logrus.Fields{
							"block_number":        bn,
							"contracts_deployed":  stats.ContractCount,
							"total_gas_used":      stats.TotalGasUsed,
							"total_bytecode_size": stats.TotalBytecodeSize,
							"avg_gas_per_byte":    fmt.Sprintf("%.2f", avgGasPerByte),
							"total_contracts":     totalContracts,
							"block_utilization":   utilizationStr,
						}).Info("Block deployment summary")

						s.lastLoggedBlock = bn
					}
				}
				s.blockStatsMutex.Unlock()
			}
		}
	}()
}

// stopBlockMonitor stops the block monitoring goroutine
func (s *Scenario) stopBlockMonitor() {
	if s.blockMonitorCancel != nil {
		s.blockMonitorCancel()
	}
	if s.blockMonitorDone != nil {
		<-s.blockMonitorDone // Wait for goroutine to finish
	}
}

// startFileWriter starts a background goroutine for async file writing
// This prevents file I/O operations from blocking the main transaction sending loop
func (s *Scenario) startFileWriter(ctx context.Context) {
	fileCtx, cancel := context.WithCancel(ctx)
	s.fileWriterCancel = cancel
	s.fileWriterChan = make(chan struct{}, 100) // Buffer to prevent blocking
	s.fileWriterDone = make(chan struct{})

	go func() {
		defer close(s.fileWriterDone)

		// Batch writes every 5 seconds or when buffer is full
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		pendingWrites := 0

		for {
			select {
			case <-fileCtx.Done():
				// Final write before shutdown
				if pendingWrites > 0 {
					s.writeDeploymentsFile()
				}
				return
			case <-s.fileWriterChan:
				pendingWrites++
				// Write immediately if buffer reaches threshold
				if pendingWrites >= 50 {
					s.writeDeploymentsFile()
					pendingWrites = 0
				}
			case <-ticker.C:
				// Periodic write if there are pending changes
				if pendingWrites > 0 {
					s.writeDeploymentsFile()
					pendingWrites = 0
				}
			}
		}
	}()
}

// stopFileWriter stops the async file writer and ensures final write
func (s *Scenario) stopFileWriter() {
	if s.fileWriterCancel != nil {
		s.fileWriterCancel()
	}
	if s.fileWriterDone != nil {
		<-s.fileWriterDone // Wait for final write to complete
	}
}

func (s *Scenario) Run(ctx context.Context) error {
	s.logger.Infof("starting scenario: %s", ScenarioName)
	defer s.logger.Infof("scenario %s finished.", ScenarioName)

	// Store original gas prices for escalation
	s.originalBaseFee = s.options.BaseFee
	s.originalTipFee = s.options.TipFee
	s.logger.Infof("Stored original gas prices - Base fee: %d gwei, Tip fee: %d gwei",
		s.originalBaseFee, s.originalTipFee)

	// Initialize channels for goroutine communication
	s.sentTxChan = make(chan *PendingTransaction, 1000) // Buffer for sent transactions
	s.newBlockChan = make(chan uint64, 10)              // Buffer for block numbers
	s.gasEscalateChan = make(chan bool, 10)             // Buffer for gas escalation signals

	// Initialize stuck transaction replacement channels
	s.requestStuckTxChan = make(chan bool, 1)                                 // Request channel
	s.responseStuckTxChan = make(chan map[common.Hash]*PendingTransaction, 1) // Response channel

	// Start async file writer to prevent blocking main loop
	s.startFileWriter(ctx)
	defer s.stopFileWriter()

	// Cache chain ID at startup
	chainID, err := s.getChainID(ctx)
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %w", err)
	}
	s.logger.Infof("Chain ID: %s", chainID.String())

	// Create error channel for goroutines
	errChan := make(chan error, 2)

	// Create a shared context for both goroutines
	runCtx, runCancel := context.WithCancel(ctx)
	defer runCancel()

	// Start block monitor goroutine (processes transactions and monitors blocks)
	go func() {
		if err := s.runBlockMonitor(runCtx); err != nil {
			s.logger.Errorf("Block monitor error: %v", err)
			errChan <- err
		}
	}()

	// Start transaction sender goroutine (only sends transactions)
	go func() {
		if err := s.runTransactionSender(runCtx); err != nil {
			s.logger.Errorf("Transaction sender error: %v", err)
			errChan <- err
		}
	}()

	// Wait for either goroutine to finish or context cancellation
	select {
	case <-ctx.Done():
		s.logger.Info("Context cancelled, stopping scenario")
		runCancel()
		// Give goroutines time to clean up
		time.Sleep(2 * time.Second)
		return ctx.Err()
	case err := <-errChan:
		s.logger.Errorf("Goroutine error, stopping scenario: %v", err)
		runCancel()
		// Give other goroutine time to clean up
		time.Sleep(2 * time.Second)
		return err
	}
}

// runTransactionSender runs the transaction sending goroutine
// This goroutine ONLY sends transactions and never blocks on processing
func (s *Scenario) runTransactionSender(ctx context.Context) error {
	s.logger.Info("Starting transaction sender goroutine")

	// Get initial block number
	client := s.walletPool.GetClient(spamoor.SelectClientByIndex, 0, s.options.ClientGroup)
	if client == nil {
		return fmt.Errorf("no client available")
	}

	ethClient := client.GetEthClient()
	currentBlock, err := ethClient.BlockByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to get initial block: %w", err)
	}

	lastBlockNumber := currentBlock.Number().Uint64()
	txIdxCounter := uint64(0)
	totalTxCount := atomic.Uint64{}

	// Main sending loop
	for {
		select {
		case <-ctx.Done():
			s.logger.Info("Transaction sender stopping due to context cancellation")
			return nil
		case <-s.gasEscalateChan:
			// Gas escalation signal received from monitor
			s.logger.Info("Received gas escalation signal from monitor")
			err := s.updateDynamicFeesWithEscalation(ctx, true)
			if err != nil {
				s.logger.Warnf("Failed to escalate gas prices: %v", err)
			}
		default:
			// Continue with normal flow
		}

		// Check if we've reached max transactions (if set)
		if s.options.MaxTransactions > 0 && txIdxCounter >= s.options.MaxTransactions {
			s.logger.Infof("Reached maximum number of transactions (%d)", s.options.MaxTransactions)
			close(s.sentTxChan) // Signal monitor to stop
			return nil
		}

		// Check wallet groups configured
		if s.walletsPerGroup == 0 {
			return fmt.Errorf("no wallet groups configured")
		}

		// Send batch of transactions for current wallet group
		groupStartIdx := s.currentWalletGroup * s.walletsPerGroup
		groupEndIdx := groupStartIdx + s.walletsPerGroup

		s.logger.Infof("Sending transactions for wallet group %d of 10 (wallets %d-%d)",
			s.currentWalletGroup, groupStartIdx, groupEndIdx-1)

		// Send all transactions for this group
		err := s.sendBatchTransactionsWithTracking(ctx, txIdxCounter, groupStartIdx, groupEndIdx)
		if err != nil {
			return fmt.Errorf("failed to send batch: %w", err)
		}

		// Update counters
		sentCount := groupEndIdx - groupStartIdx
		txIdxCounter += sentCount
		totalTxCount.Add(sentCount)

		// Wait for new block signal from monitor
		s.logger.Info("Waiting for next block...")

		waitTimeout := time.NewTimer(2 * time.Minute)
		select {
		case <-ctx.Done():
			waitTimeout.Stop()
			return nil
		case newBlockNum := <-s.newBlockChan:
			waitTimeout.Stop()
			lastBlockNumber = newBlockNum
			s.logger.Infof("New block %d: switching to wallet group %d",
				lastBlockNumber, (s.currentWalletGroup+1)%10)
			// Rotate to next wallet group
			previousGroup := s.currentWalletGroup
			s.currentWalletGroup = (s.currentWalletGroup + 1) % 10

			// Check if we completed a full round (went from group 9 back to 0)
			if previousGroup == 9 && s.currentWalletGroup == 0 {
				s.logger.Info("Completed full wallet group cycle, checking for stuck transactions")
				s.handleStuckTransactions(ctx)
			}

		case <-waitTimeout.C:
			// Timeout waiting for block, force progression
			s.logger.Warn("Timeout waiting for new block, forcing progression")
			previousGroup := s.currentWalletGroup
			s.currentWalletGroup = (s.currentWalletGroup + 1) % 10

			// Check if we completed a full round
			if previousGroup == 9 && s.currentWalletGroup == 0 {
				s.logger.Info("Completed full wallet group cycle, checking for stuck transactions")
				s.handleStuckTransactions(ctx)
			}
		}
	}
}

// handleStuckTransactions requests stuck transactions from monitor and sends replacements
func (s *Scenario) handleStuckTransactions(ctx context.Context) {
	// Request stuck transactions from monitor
	select {
	case s.requestStuckTxChan <- true:
		s.logger.Debug("Sent request for stuck transactions")
	case <-time.After(1 * time.Second):
		s.logger.Warn("Timeout sending stuck tx request, monitor might be busy")
		return
	}

	// Wait for response with stuck transactions
	var stuckTxs map[common.Hash]*PendingTransaction
	select {
	case stuckTxs = <-s.responseStuckTxChan:
		s.logger.Infof("Received %d stuck transactions to replace", len(stuckTxs))
	case <-time.After(5 * time.Second):
		s.logger.Warn("Timeout waiting for stuck transactions response")
		return
	}

	if len(stuckTxs) == 0 {
		s.logger.Debug("No stuck transactions found")
		return
	}

	// Send replacement transactions
	s.sendReplacementTransactions(ctx, stuckTxs)
}

// sendReplacementTransactions sends replacement transactions for stuck ones
func (s *Scenario) sendReplacementTransactions(ctx context.Context, stuckTxs map[common.Hash]*PendingTransaction) {
	s.logger.Infof("Sending replacement transactions for %d stuck transactions", len(stuckTxs))

	// Get client for querying transaction details
	client := s.walletPool.GetClient(spamoor.SelectClientByIndex, 0, s.options.ClientGroup)
	if client == nil {
		s.logger.Error("No client available for replacement transactions")
		return
	}
	ethClient := client.GetEthClient()

	replacedCount := 0
	failedCount := 0

	// Process each stuck transaction
	for txHash, pendingTx := range stuckTxs {
		// Skip transactions that were not actually sent successfully
		if !pendingTx.Sent {
			s.logger.Debugf("Skipping transaction %s - was not successfully sent", txHash.Hex())
			continue
		}

		// Get the actual transaction to find its nonce
		tx, _, err := ethClient.TransactionByHash(ctx, txHash)
		if err != nil {
			// Transaction not found - it was never actually sent or already dropped from mempool
			s.logger.Debugf("Transaction %s not found in mempool (may have been dropped): %v", txHash.Hex(), err)
			failedCount++
			continue
		}

		// Find the wallet
		walletAddr := crypto.PubkeyToAddress(pendingTx.PrivateKey.PublicKey)
		var wallet *spamoor.Wallet

		// Search for the wallet in our pool
		for i := uint64(0); i < s.walletPool.GetWalletCount(); i++ {
			w := s.walletPool.GetWallet(spamoor.SelectWalletByIndex, int(i))
			if w != nil && w.GetAddress() == walletAddr {
				wallet = w
				break
			}
		}

		if wallet == nil {
			s.logger.Errorf("Could not find wallet for address %s", walletAddr.Hex())
			failedCount++
			continue
		}

		// Use the transaction's actual nonce for replacement
		txNonce := tx.Nonce()

		// Build and send replacement transaction
		if err := s.sendSingleReplacement(ctx, wallet, txNonce); err != nil {
			s.logger.Warnf("Failed to send replacement for %s (nonce %d): %v", txHash.Hex(), txNonce, err)
			failedCount++
		} else {
			s.logger.Infof("Sent replacement for stuck tx %s with nonce %d", txHash.Hex(), txNonce)
			replacedCount++
		}
	}

	s.logger.Infof("Replacement complete: %d succeeded, %d failed", replacedCount, failedCount)
}

// sendSingleReplacement sends a single replacement transaction with higher gas
func (s *Scenario) sendSingleReplacement(ctx context.Context, wallet *spamoor.Wallet, nonce uint64) error {
	// Get client
	client := s.walletPool.GetClient(spamoor.SelectClientRoundRobin, 0, "")
	if client == nil {
		return fmt.Errorf("no client available")
	}

	// Generate random salt for unique contract
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		return fmt.Errorf("failed to generate salt: %w", err)
	}
	saltInt := new(big.Int).SetBytes(salt)

	// Calculate replacement gas prices (20% higher)
	replacementBaseFee := s.options.BaseFee + (s.options.BaseFee * 20 / 100)
	replacementTipFee := s.options.TipFee + (s.options.TipFee * 20 / 100)

	s.logger.Debugf("Replacement gas prices - Base: %d gwei (was %d), Tip: %d gwei (was %d)",
		replacementBaseFee, s.options.BaseFee, replacementTipFee, s.options.TipFee)

	// Build replacement transaction with specific nonce
	txCtx, cancel := context.WithTimeout(ctx, 10*time.Second) // Longer timeout for replacements
	defer cancel()

	// We need to use ReplaceTransaction method or build manually with specific nonce
	// Since BuildBoundTx auto-increments nonce, we'll use a manual approach
	transactor, err := bind.NewKeyedTransactorWithChainID(wallet.GetPrivateKey(), s.chainID)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %w", err)
	}

	transactor.Context = txCtx
	transactor.From = wallet.GetAddress()
	transactor.Nonce = big.NewInt(int64(nonce))
	transactor.GasTipCap = new(big.Int).Mul(big.NewInt(int64(replacementTipFee)), big.NewInt(1000000000))
	transactor.GasFeeCap = new(big.Int).Mul(big.NewInt(int64(replacementBaseFee)), big.NewInt(1000000000))
	transactor.GasLimit = 5200000
	transactor.Value = big.NewInt(0)
	transactor.NoSend = true

	// Deploy the contract
	_, tx, _, err := contract.DeployContract(transactor, client.GetEthClient(), saltInt)
	if err != nil {
		return fmt.Errorf("failed to build replacement contract deployment: %w", err)
	}

	// Send the replacement transaction
	sendCtx, cancelSend := context.WithTimeout(ctx, 10*time.Second)
	defer cancelSend()
	err = client.SendTransaction(sendCtx, tx)
	if err != nil {
		return fmt.Errorf("failed to send replacement transaction: %w", err)
	}

	// Track the replacement transaction
	select {
	case s.sentTxChan <- &PendingTransaction{
		TxHash:     tx.Hash(),
		PrivateKey: wallet.GetPrivateKey(),
		Timestamp:  time.Now(),
	}:
		// Successfully sent to monitor
	default:
		// Channel full, but not critical for replacements
	}

	return nil
}

// sendBatchTransactionsWithTracking sends transactions and tracks them via channel
func (s *Scenario) sendBatchTransactionsWithTracking(ctx context.Context, baseIdx, startIdx, endIdx uint64) error {
	var wg sync.WaitGroup
	errors := make(chan error, endIdx-startIdx)
	successfulTxs := make(chan *PendingTransaction, endIdx-startIdx)

	// Send one transaction per wallet in the group
	for walletIdx := startIdx; walletIdx < endIdx; walletIdx++ {
		wg.Add(1)
		go func(idx uint64) {
			defer wg.Done()

			// Send transaction and track it
			if tx, privKey, err := s.sendAndTrackTransaction(ctx, idx); err != nil {
				errors <- fmt.Errorf("wallet %d: %w", idx, err)
			} else if tx != nil {
				// Only track successfully sent transactions
				successfulTxs <- &PendingTransaction{
					TxHash:     tx.Hash(),
					PrivateKey: privKey,
					Timestamp:  time.Now(),
					Sent:       true,
				}
			}
		}(walletIdx)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	close(errors)
	close(successfulTxs)

	// Send successful transactions to monitor (non-blocking)
	for pendingTx := range successfulTxs {
		select {
		case s.sentTxChan <- pendingTx:
			// Successfully sent to monitor
		default:
			// Channel full, log but don't block
			s.logger.Warnf("sentTxChan full, dropping tx %s", pendingTx.TxHash.Hex())
		}
	}

	// Collect any errors
	var errs []error
	for err := range errors {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		s.logger.Warnf("Failed to send %d transactions out of %d", len(errs), endIdx-startIdx)
		// Log first few errors for debugging
		for i, err := range errs {
			if i >= 5 {
				break
			}
			s.logger.Warnf("Error %d: %v", i+1, err)
		}
	}

	s.logger.Infof("Sent %d transactions from wallet group", endIdx-startIdx-uint64(len(errs)))
	return nil
}

// sendBatchTransactions sends all transactions for a wallet group in parallel
func (s *Scenario) sendBatchTransactions(ctx context.Context, baseIdx, startIdx, endIdx uint64) error {
	var wg sync.WaitGroup
	errors := make(chan error, endIdx-startIdx)

	// Send one transaction per wallet in the group
	for walletIdx := startIdx; walletIdx < endIdx; walletIdx++ {
		wg.Add(1)
		go func(idx uint64) {
			defer wg.Done()

			// Pass the wallet index directly as the transaction index
			// This ensures each wallet is used exactly once
			if err := s.sendTransaction(ctx, idx); err != nil {
				errors <- fmt.Errorf("wallet %d: %w", idx, err)
			}
		}(walletIdx)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	close(errors)

	// Collect any errors
	var errs []error
	for err := range errors {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		s.logger.Warnf("Failed to send %d transactions out of %d", len(errs), endIdx-startIdx)
		// Log first few errors for debugging
		for i, err := range errs {
			if i >= 5 {
				break
			}
			s.logger.Warnf("Error %d: %v", i+1, err)
		}
	}

	s.logger.Infof("Sent %d transactions from wallet group", endIdx-startIdx-uint64(len(errs)))
	return nil
}

// runBlockMonitor runs the block monitoring and transaction processing goroutine
// This goroutine handles all heavy processing to keep the sender free
func (s *Scenario) runBlockMonitor(ctx context.Context) error {
	s.logger.Info("Starting block monitor goroutine")

	// Get client for monitoring
	client := s.walletPool.GetClient(spamoor.SelectClientByIndex, 0, s.options.ClientGroup)
	if client == nil {
		return fmt.Errorf("no client available for monitoring")
	}
	ethClient := client.GetEthClient()

	// Get initial block number
	currentBlock, err := ethClient.BlockByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to get initial block: %w", err)
	}

	lastBlockNumber := currentBlock.Number().Uint64()
	consecutiveEmptyBlocks := uint64(0)

	// Initialize pending transactions map
	pendingTxs := make(map[common.Hash]*PendingTransaction)

	// Block monitoring ticker
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	// Track total contracts for final summary
	totalContracts := 0

	for {
		select {
		case <-ctx.Done():
			s.logger.Info("Block monitor stopping due to context cancellation")

			// Final summary
			s.logger.WithFields(logrus.Fields{
				"total_contracts": totalContracts,
			}).Info("Block monitor final summary")

			return nil

		case pendingTx := <-s.sentTxChan:
			// New transaction sent by sender
			if pendingTx != nil {
				pendingTxs[pendingTx.TxHash] = pendingTx
				s.logger.Debugf("Tracking new transaction %s", pendingTx.TxHash.Hex())
			}

		case <-s.requestStuckTxChan:
			// Sender requested list of stuck transactions
			s.logger.Debug("Received request for stuck transactions")
			stuckTxs := s.getStuckTransactions(pendingTxs, 30*time.Second)

			// Send response
			select {
			case s.responseStuckTxChan <- stuckTxs:
				s.logger.Debugf("Sent %d stuck transactions to sender", len(stuckTxs))
			case <-time.After(1 * time.Second):
				s.logger.Warn("Timeout sending stuck transactions response")
			}

		case <-ticker.C:
			// Check for new block
			blockCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
			latestBlock, err := ethClient.BlockByNumber(blockCtx, nil)
			cancel()

			if err != nil {
				s.logger.Warnf("Failed to get latest block: %v", err)
				continue
			}

			currentBlockNum := latestBlock.Number().Uint64()

			// Process new blocks
			if currentBlockNum > lastBlockNumber {
				// Process all blocks from last to current
				for blockNum := lastBlockNumber + 1; blockNum <= currentBlockNum; blockNum++ {
					// Get block details
					blockCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
					block, err := ethClient.BlockByNumber(blockCtx, big.NewInt(int64(blockNum)))
					cancel()

					if err != nil {
						s.logger.Warnf("Failed to get block %d: %v", blockNum, err)
						continue
					}

					// Check if block is empty
					txCount := len(block.Transactions())
					if txCount == 0 {
						consecutiveEmptyBlocks++
						s.logger.Warnf("Block %d is EMPTY (0 transactions), consecutive: %d",
							blockNum, consecutiveEmptyBlocks)

						// Send gas escalation signal on first empty block
						if consecutiveEmptyBlocks == 1 {
							select {
							case s.gasEscalateChan <- true:
								s.logger.Info("Sent gas escalation signal due to empty block")
							default:
								// Channel full, skip
							}
						}
					} else {
						consecutiveEmptyBlocks = 0
					}

					// Process transactions in this block
					confirmedCount := s.processBlockTransactions(ctx, block, pendingTxs)
					if confirmedCount > 0 {
						totalContracts += confirmedCount
						s.logger.Infof("Block %d: confirmed %d contracts (total: %d)",
							blockNum, confirmedCount, totalContracts)
					}

					// Clean up old pending transactions (timeout after 2 minutes)
					s.cleanupOldTransactions(pendingTxs, 2*time.Minute)
				}

				lastBlockNumber = currentBlockNum

				// Send new block signal to sender
				select {
				case s.newBlockChan <- currentBlockNum:
					// Successfully sent block signal
				default:
					// Channel full, sender will timeout and progress anyway
					s.logger.Warnf("newBlockChan full, sender will timeout")
				}
			}
		}
	}
}

// processBlockTransactions processes all transactions in a block and returns confirmed count
func (s *Scenario) processBlockTransactions(ctx context.Context, block *types.Block, pendingTxs map[common.Hash]*PendingTransaction) int {
	confirmedCount := 0

	// Check each transaction in the block
	for _, tx := range block.Transactions() {
		// Check if this is one of our pending transactions
		if pendingTx, exists := pendingTxs[tx.Hash()]; exists {
			// Get receipt
			client := s.walletPool.GetClient(spamoor.SelectClientByIndex, 0, s.options.ClientGroup)
			if client == nil {
				continue
			}

			receiptCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
			receipt, err := client.GetEthClient().TransactionReceipt(receiptCtx, tx.Hash())
			cancel()

			if err != nil {
				s.logger.Warnf("Failed to get receipt for %s: %v", tx.Hash().Hex(), err)
				continue
			}

			// Process successful deployment
			if receipt.Status == 1 && receipt.ContractAddress != (common.Address{}) {
				s.recordDeployedContract(receipt.ContractAddress, pendingTx.PrivateKey, receipt, tx.Hash())
				confirmedCount++
			}

			// Remove from pending
			delete(pendingTxs, tx.Hash())
		}
	}

	return confirmedCount
}

// cleanupOldTransactions removes transactions that have been pending too long
func (s *Scenario) cleanupOldTransactions(pendingTxs map[common.Hash]*PendingTransaction, timeout time.Duration) {
	now := time.Now()
	var toRemove []common.Hash

	for hash, tx := range pendingTxs {
		if now.Sub(tx.Timestamp) > timeout {
			toRemove = append(toRemove, hash)
		}
	}

	for _, hash := range toRemove {
		delete(pendingTxs, hash)
		s.logger.Debugf("Removed timed out transaction %s", hash.Hex())
	}
}

// getStuckTransactions returns transactions that have been pending longer than threshold
func (s *Scenario) getStuckTransactions(pendingTxs map[common.Hash]*PendingTransaction, threshold time.Duration) map[common.Hash]*PendingTransaction {
	stuckTxs := make(map[common.Hash]*PendingTransaction)
	now := time.Now()

	for hash, tx := range pendingTxs {
		if now.Sub(tx.Timestamp) > threshold {
			stuckTxs[hash] = tx
		}
	}

	return stuckTxs
}

// sendAndTrackTransaction sends a single transaction and returns tx + private key for tracking
func (s *Scenario) sendAndTrackTransaction(ctx context.Context, walletIdx uint64) (*types.Transaction, *ecdsa.PrivateKey, error) {
	maxRetries := 3

	for attempt := 0; attempt < maxRetries; attempt++ {
		tx, privKey, err := s.attemptTransactionWithReturn(ctx, walletIdx, attempt)
		if err == nil {
			return tx, privKey, nil
		}

		// Handle different error types
		if strings.Contains(err.Error(), "replacement transaction underpriced") {
			// Skip nonce conflicts quickly
			wallet := s.walletPool.GetWallet(spamoor.SelectWalletByIndex, int(walletIdx))
			if wallet != nil {
				wallet.GetNextNonce() // Skip the conflicting nonce
			}
			continue
		}

		if strings.Contains(err.Error(), "max fee per gas less than block base fee") {
			// Update fees and retry
			s.updateDynamicFees(ctx)
			continue
		}

		// For other errors, return immediately
		return nil, nil, err
	}

	return nil, nil, fmt.Errorf("failed after %d attempts", maxRetries)
}

// sendTransaction sends a single contract deployment transaction
func (s *Scenario) sendTransaction(ctx context.Context, txIdx uint64) error {
	maxRetries := 3
	maxNonceSkips := 5

	for attempt := 0; attempt < maxRetries; attempt++ {
		err := s.attemptTransaction(ctx, txIdx, attempt)
		if err == nil {
			return nil
		}

		// Check if it's a nonce conflict error (replacement transaction underpriced)
		if strings.Contains(err.Error(), "replacement transaction underpriced") {
			s.logger.Warnf("Transaction %d nonce conflict detected, skipping to next nonce (attempt %d/%d)",
				txIdx, attempt+1, maxRetries)

			// Get the wallet and skip the conflicting nonce
			wallet := s.walletPool.GetWallet(spamoor.SelectWalletByIndex, int(txIdx))
			if wallet != nil {
				skippedNonces := 0
				for skippedNonces < maxNonceSkips {
					oldNonce := wallet.GetNonce()
					newNonce := wallet.GetNextNonce() // Skip the conflicting nonce
					s.logger.Infof("Wallet %d skipped nonce %d, now using nonce %d",
						txIdx, oldNonce, newNonce)
					skippedNonces++

					// Retry with the new nonce
					err = s.attemptTransaction(ctx, txIdx, attempt)
					if err == nil {
						return nil
					}

					// If still getting nonce conflict, continue skipping
					if !strings.Contains(err.Error(), "replacement transaction underpriced") {
						break
					}
				}

				if skippedNonces >= maxNonceSkips {
					s.logger.Errorf("Wallet %d exceeded maximum nonce skips (%d), giving up",
						txIdx, maxNonceSkips)
					return fmt.Errorf("exceeded maximum nonce skips (%d) for wallet %d", maxNonceSkips, txIdx)
				}
			}

			time.Sleep(time.Duration(attempt+1) * 200 * time.Millisecond) // Short backoff for nonce conflicts
			continue
		}

		// Check if it's a base fee error
		if strings.Contains(err.Error(), "max fee per gas less than block base fee") {
			s.logger.Warnf("Transaction %d base fee too low, adjusting fees and retrying (attempt %d/%d)",
				txIdx, attempt+1, maxRetries)

			// Update fees based on current network conditions
			if updateErr := s.updateDynamicFees(ctx); updateErr != nil {
				s.logger.Warnf("Failed to update dynamic fees: %v", updateErr)
			}

			time.Sleep(time.Duration(attempt+1) * 500 * time.Millisecond) // Exponential backoff
			continue
		}

		// For other errors, return immediately
		return err
	}

	return fmt.Errorf("failed to send transaction after %d attempts", maxRetries)
}

// updateDynamicFees queries the network and updates base fee and tip fee
func (s *Scenario) updateDynamicFees(ctx context.Context) error {
	return s.updateDynamicFeesWithEscalation(ctx, false)
}

// updateDynamicFeesWithEscalation updates fees with optional aggressive escalation
func (s *Scenario) updateDynamicFeesWithEscalation(ctx context.Context, escalate bool) error {
	client := s.walletPool.GetClient(spamoor.SelectClientByIndex, 0, s.options.ClientGroup)
	if client == nil {
		return fmt.Errorf("no client available")
	}

	if escalate {
		// Check if we've reached maximum gas price limits (prevent infinite escalation)
		maxBaseFee := uint64(1000) // 1000 gwei max base fee
		maxTipFee := uint64(100)   // 100 gwei max tip fee

		if s.options.BaseFee >= maxBaseFee || s.options.TipFee >= maxTipFee {
			s.maxGasPriceReached = true
			s.logger.Errorf("ESCALATION LIMIT: Maximum gas prices reached (base: %d/%d gwei, tip: %d/%d gwei), stopping escalation",
				s.options.BaseFee, maxBaseFee, s.options.TipFee, maxTipFee)
			return fmt.Errorf("maximum gas price limit reached")
		}

		// Aggressive escalation: add original fees to current fees
		oldBaseFee := s.options.BaseFee
		oldTipFee := s.options.TipFee

		s.options.BaseFee = s.options.BaseFee + s.originalBaseFee
		s.options.TipFee = s.options.TipFee + s.originalTipFee

		// Check if we've exceeded limits after escalation
		if s.options.BaseFee > maxBaseFee {
			s.options.BaseFee = maxBaseFee
			s.maxGasPriceReached = true
		}
		if s.options.TipFee > maxTipFee {
			s.options.TipFee = maxTipFee
			s.maxGasPriceReached = true
		}

		s.logger.Warnf("ESCALATING gas prices due to stuck transactions - Base fee: %d -> %d gwei (+%d), Tip fee: %d -> %d gwei (+%d)",
			oldBaseFee, s.options.BaseFee, s.originalBaseFee,
			oldTipFee, s.options.TipFee, s.originalTipFee)

		if s.maxGasPriceReached {
			s.logger.Warnf("Gas price escalation has reached maximum limits")
		}

		return nil
	}

	ethClient := client.GetEthClient()

	// Get the latest block to check current base fee with timeout
	blockCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	latestBlock, err := ethClient.BlockByNumber(blockCtx, nil)
	if err != nil {
		return fmt.Errorf("failed to get latest block: %w", err)
	}

	if latestBlock.BaseFee() != nil {
		// Convert base fee from wei to gwei
		currentBaseFeeGwei := new(big.Int).Div(latestBlock.BaseFee(), big.NewInt(1000000000))

		newBaseFeeGwei := new(big.Int).Add(currentBaseFeeGwei, big.NewInt(100))

		s.options.BaseFee = newBaseFeeGwei.Uint64()

		// Also increase tip fee slightly to ensure competitive priority
		if s.options.TipFee+1 > 3 {
			s.options.TipFee = s.options.TipFee + 1
		} else {
			s.options.TipFee = 2 // Minimum 3 gwei tip
		}

		s.logger.Infof("Updated dynamic fees - Base fee: %d gwei, Tip fee: %d gwei (network base fee: %s gwei)",
			s.options.BaseFee, s.options.TipFee, currentBaseFeeGwei.String())
	}

	return nil
}

// attemptTransactionWithReturn builds and sends a transaction, returning the tx and private key
func (s *Scenario) attemptTransactionWithReturn(ctx context.Context, walletIdx uint64, attempt int) (*types.Transaction, *ecdsa.PrivateKey, error) {
	// Get client
	client := s.walletPool.GetClient(spamoor.SelectClientRoundRobin, 0, "")
	if client == nil {
		return nil, nil, fmt.Errorf("no client available")
	}

	// Get wallet
	wallet := s.walletPool.GetWallet(spamoor.SelectWalletByIndex, int(walletIdx))
	if wallet == nil {
		return nil, nil, fmt.Errorf("no wallet available at index %d", walletIdx)
	}

	// Generate random salt for unique contract
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate salt: %w", err)
	}
	saltInt := new(big.Int).SetBytes(salt)

	// Use BuildBoundTx with timeout to prevent hanging
	txCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	tx, err := wallet.BuildBoundTx(txCtx, &txbuilder.TxMetadata{
		GasFeeCap: uint256.MustFromBig(new(big.Int).Mul(big.NewInt(int64(s.options.BaseFee)), big.NewInt(1000000000))),
		GasTipCap: uint256.MustFromBig(new(big.Int).Mul(big.NewInt(int64(s.options.TipFee)), big.NewInt(1000000000))),
		Gas:       5200000, // Fixed gas limit for contract deployment
		Value:     uint256.NewInt(0),
	}, func(transactOpts *bind.TransactOpts) (*types.Transaction, error) {
		// Deploy the contract using the transactOpts provided by BuildBoundTx
		_, deployTx, _, err := contract.DeployContract(transactOpts, client.GetEthClient(), saltInt)
		return deployTx, err
	})

	if err != nil {
		return nil, nil, fmt.Errorf("failed to build and deploy contract: %w", err)
	}

	// Send the transaction with timeout
	sendCtx, cancelSend := context.WithTimeout(ctx, 3*time.Second)
	defer cancelSend()
	err = client.SendTransaction(sendCtx, tx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to send transaction: %w", err)
	}

	return tx, wallet.GetPrivateKey(), nil
}

// attemptTransaction makes a single attempt to send a transaction
func (s *Scenario) attemptTransaction(ctx context.Context, txIdx uint64, attempt int) error {
	// Get client
	client := s.walletPool.GetClient(spamoor.SelectClientRoundRobin, 0, "")
	if client == nil {
		return fmt.Errorf("no client available")
	}

	// txIdx is the wallet index in batch mode
	wallet := s.walletPool.GetWallet(spamoor.SelectWalletByIndex, int(txIdx))
	if wallet == nil {
		return fmt.Errorf("no wallet available at index %d", txIdx)
	}

	// Generate random salt for unique contract
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		return fmt.Errorf("failed to generate salt: %w", err)
	}
	saltInt := new(big.Int).SetBytes(salt)

	// Use BuildBoundTx with timeout to prevent hanging
	txCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	tx, err := wallet.BuildBoundTx(txCtx, &txbuilder.TxMetadata{
		GasFeeCap: uint256.MustFromBig(new(big.Int).Mul(big.NewInt(int64(s.options.BaseFee)), big.NewInt(1000000000))),
		GasTipCap: uint256.MustFromBig(new(big.Int).Mul(big.NewInt(int64(s.options.TipFee)), big.NewInt(1000000000))),
		Gas:       5200000, // Fixed gas limit for contract deployment
		Value:     uint256.NewInt(0),
	}, func(transactOpts *bind.TransactOpts) (*types.Transaction, error) {
		// Deploy the contract using the transactOpts provided by BuildBoundTx
		_, deployTx, _, err := contract.DeployContract(transactOpts, client.GetEthClient(), saltInt)
		return deployTx, err
	})

	if err != nil {
		return fmt.Errorf("failed to build and deploy contract: %w", err)
	}

	// Send the transaction with timeout
	sendCtx, cancelSend := context.WithTimeout(ctx, 3*time.Second)
	defer cancelSend()
	err = client.SendTransaction(sendCtx, tx)
	if err != nil {
		return fmt.Errorf("failed to send transaction: %w", err)
	}

	// Track pending transaction
	pendingTx := &PendingTransaction{
		TxHash:     tx.Hash(),
		PrivateKey: wallet.GetPrivateKey(),
		Timestamp:  time.Now(),
	}

	s.pendingTxsMutex.Lock()
	s.pendingTxs[tx.Hash()] = pendingTx
	s.pendingTxsMutex.Unlock()

	return nil
}
