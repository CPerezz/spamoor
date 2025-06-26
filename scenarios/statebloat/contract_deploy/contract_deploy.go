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
	consecutiveEmptyBlocks    uint64
	consecutiveFailedBlockWaits uint64
	emptyBlockRetryCount      uint64
	maxGasPriceReached        bool
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

// processPendingTransactions checks for transaction confirmations and updates state
func (s *Scenario) processPendingTransactions(ctx context.Context) {
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

	for txHash, pendingTx := range s.pendingTxs {
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

		receipt, err := ethClient.TransactionReceipt(ctx, txHash)
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
		
		// Create timeout context for escalation to prevent hanging
		escalationCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		
		err := s.updateDynamicFeesWithEscalation(escalationCtx, true)
		if err != nil {
			s.logger.Warnf("Failed to escalate gas prices: %v", err)
		} else {
			// Trigger immediate transaction sending with escalated prices
			s.triggerImmediateTxSending(ctx, len(stuckTxs))
		}
	}

	// Process successful deployments
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

	// Get the actual deployed contract bytecode size
	client := s.walletPool.GetClient(spamoor.SelectClientByIndex, 0, s.options.ClientGroup)
	// TODO: This should be a constant documented on how this number is obtained.
	var bytecodeSize int = 23914

	if client != nil {
		// Get the actual deployed bytecode size using eth_getCode
		contractCode, err := client.GetEthClient().CodeAt(context.Background(), contractAddress, nil)
		if err == nil {
			bytecodeSize = len(contractCode)
		}
	}

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

	// Save the deployments.json file each time a contract is confirmed
	if err := s.saveDeploymentsMapping(); err != nil {
		s.logger.Warnf("Failed to save deployments.json: %v", err)
	}
}

// Helper function for max calculation
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// saveDeploymentsMapping creates/updates deployments.json with private key to contract address mapping
func (s *Scenario) saveDeploymentsMapping() error {
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

// startBlockMonitor starts a background goroutine that monitors for new blocks
// and logs block deployment summaries immediately when blocks are mined
func (s *Scenario) startBlockMonitor(ctx context.Context) {
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

func (s *Scenario) Run(ctx context.Context) error {
	s.logger.Infof("starting scenario: %s", ScenarioName)
	defer s.logger.Infof("scenario %s finished.", ScenarioName)

	// Store original gas prices for escalation
	s.originalBaseFee = s.options.BaseFee
	s.originalTipFee = s.options.TipFee
	s.logger.Infof("Stored original gas prices - Base fee: %d gwei, Tip fee: %d gwei", 
		s.originalBaseFee, s.originalTipFee)

	// Start block monitoring for real-time logging
	s.startBlockMonitor(ctx)
	defer s.stopBlockMonitor()

	// Cache chain ID at startup
	chainID, err := s.getChainID(ctx)
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %w", err)
	}

	s.logger.Infof("Chain ID: %s", chainID.String())

	// Get client first
	client := s.walletPool.GetClient(spamoor.SelectClientByIndex, 0, s.options.ClientGroup)
	if client == nil {
		return fmt.Errorf("no client available")
	}
	ethClient := client.GetEthClient()

	// Get current block for initialization
	currentBlock, err := ethClient.BlockByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to get current block: %w", err)
	}

	txIdxCounter := uint64(0)
	totalTxCount := atomic.Uint64{}
	lastBlockNumber := currentBlock.Number().Uint64()

	// Main loop for alternating wallet groups
	loopIteration := uint64(0)
	for {
		loopIteration++
		
		// Enhanced monitoring every 10 iterations
		if loopIteration%10 == 0 {
			s.logger.Infof("DEADLOCK MONITORING: Loop iteration %d, empty blocks: %d, failed waits: %d, retry count: %d, max gas reached: %v",
				loopIteration, s.consecutiveEmptyBlocks, s.consecutiveFailedBlockWaits, s.emptyBlockRetryCount, s.maxGasPriceReached)
		}
		
		// Check if we've reached max transactions (if set)
		if s.options.MaxTransactions > 0 && txIdxCounter >= s.options.MaxTransactions {
			s.logger.Infof("reached maximum number of transactions (%d)", s.options.MaxTransactions)
			break
		}

		// Only proceed if we have wallet groups configured
		if s.walletsPerGroup == 0 {
			s.logger.Error("No wallet groups configured")
			break
		}
		
		// Automatic recovery if too many consecutive failures
		if s.consecutiveFailedBlockWaits > 10 {
			s.logger.Errorf("AUTOMATIC RECOVERY: Too many consecutive failed block waits (%d), resetting counters",
				s.consecutiveFailedBlockWaits)
			s.consecutiveFailedBlockWaits = 0
			s.consecutiveEmptyBlocks = 0
			s.emptyBlockRetryCount = 0
		}

		// Send batch of transactions for current wallet group
		groupStartIdx := s.currentWalletGroup * s.walletsPerGroup
		groupEndIdx := groupStartIdx + s.walletsPerGroup

		s.logger.Infof("Sending transactions for wallet group %d of 10 (wallets %d-%d)",
			s.currentWalletGroup, groupStartIdx, groupEndIdx-1)

		// Send all transactions for this group in parallel
		err := s.sendBatchTransactions(ctx, txIdxCounter, groupStartIdx, groupEndIdx)
		if err != nil {
			s.logger.Errorf("Failed to send batch transactions: %v", err)
			break
		}

		// Update counters
		sentCount := groupEndIdx - groupStartIdx
		txIdxCounter += sentCount
		totalTxCount.Add(sentCount)

		// Process pending transactions
		s.processPendingTransactions(ctx)

		// Wait for block number to change before switching groups
		s.logger.Info("Waiting for next block...")
		waitStartTime := time.Now()
		maxWaitTime := 2 * time.Minute // Maximum wait time before forcing progression
		
		for {
			// Check if context is cancelled
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			// Check if we've exceeded maximum wait time
			waitDuration := time.Since(waitStartTime)
			if waitDuration > maxWaitTime {
				s.consecutiveFailedBlockWaits++
				s.logger.Errorf("DEADLOCK PREVENTION: Forcing progression after %v wait (failed waits: %d)", 
					waitDuration, s.consecutiveFailedBlockWaits)
				
				// Force progression to next wallet group to prevent deadlock
				s.currentWalletGroup = (s.currentWalletGroup + 1) % 10
				s.logger.Warnf("Forced switch to wallet group %d to prevent deadlock", s.currentWalletGroup)
				break
			}

			currentBlock, err := ethClient.BlockByNumber(ctx, nil)
			if err != nil {
				s.logger.Warnf("Failed to get current block: %v", err)
				time.Sleep(500 * time.Millisecond)
				continue
			}

			currentBlockNum := currentBlock.Number().Uint64()
			if currentBlockNum > lastBlockNumber {
				lastBlockNumber = currentBlockNum
				s.consecutiveFailedBlockWaits = 0 // Reset counter on successful block progression

				// Check if this block is empty (no transactions)
				blockTxCount := len(currentBlock.Transactions())
				if blockTxCount == 0 {
					s.consecutiveEmptyBlocks++
					s.logger.Warnf("Block %d is EMPTY (0 transactions), consecutive empty blocks: %d", 
						lastBlockNumber, s.consecutiveEmptyBlocks)
					
					// Circuit breaker: if too many consecutive empty blocks, stop retrying
					if s.consecutiveEmptyBlocks > 5 {
						s.logger.Errorf("CIRCUIT BREAKER: Too many consecutive empty blocks (%d), skipping retry to prevent deadlock", 
							s.consecutiveEmptyBlocks)
					} else if s.emptyBlockRetryCount < 3 && !s.maxGasPriceReached {
						// Limited retries per empty block
						s.emptyBlockRetryCount++
						
						// Escalate gas prices due to empty block
						err := s.updateDynamicFeesWithEscalation(ctx, true)
						if err != nil {
							s.logger.Warnf("Failed to escalate gas prices after empty block: %v", err)
						}
						
						// Add cooldown to prevent excessive retries
						s.logger.Infof("Adding 10-second cooldown before retry (attempt %d/3)", s.emptyBlockRetryCount)
						time.Sleep(10 * time.Second)
						
						// Immediately send transactions with new gas prices (skip waiting for next block)
						groupStartIdx := s.currentWalletGroup * s.walletsPerGroup
						groupEndIdx := groupStartIdx + s.walletsPerGroup
						
						s.logger.Infof("RETRYING transactions for wallet group %d of 10 (wallets %d-%d) with escalated prices",
							s.currentWalletGroup, groupStartIdx, groupEndIdx-1)
						
						// Send batch with escalated prices
						err = s.sendBatchTransactions(ctx, txIdxCounter, groupStartIdx, groupEndIdx)
						if err != nil {
							s.logger.Errorf("Failed to send retry batch after empty block: %v", err)
						} else {
							// Update counters for the retry batch
							sentCount := groupEndIdx - groupStartIdx
							txIdxCounter += sentCount
							totalTxCount.Add(sentCount)
						}
					} else {
						s.logger.Warnf("Skipping empty block retry (attempts: %d/3, max gas reached: %v)", 
							s.emptyBlockRetryCount, s.maxGasPriceReached)
					}
				} else {
					// Block has transactions - reset empty block counters
					s.consecutiveEmptyBlocks = 0
					s.emptyBlockRetryCount = 0
				}

				// Rotate to the next wallet group (0-9)
				s.currentWalletGroup = (s.currentWalletGroup + 1) % 10
				s.logger.Infof("New block %d: switching to wallet group %d",
					lastBlockNumber, s.currentWalletGroup)

				break
			}

			// Log if we've been waiting too long
			if waitDuration > 30*time.Second {
				s.logger.Warnf("Been waiting for new block for %v, current block: %d, waiting for block > %d",
					waitDuration, currentBlockNum, lastBlockNumber)
				waitStartTime = time.Now() // Reset to avoid log spam
			}

			time.Sleep(100 * time.Millisecond)
		}
	}

	// Wait for all pending transactions to complete with 1 second intervals
	s.logger.Info("Waiting for remaining transactions to complete...")
	for {
		s.processPendingTransactions(ctx)

		s.pendingTxsMutex.RLock()
		pendingCount := len(s.pendingTxs)
		s.pendingTxsMutex.RUnlock()

		if pendingCount == 0 {
			break
		}

		s.logger.Infof("Waiting for %d pending transactions...", pendingCount)
		time.Sleep(1 * time.Second) // Changed from 2 seconds to 1 second
	}

	// Stop block monitoring before final cleanup
	s.stopBlockMonitor()

	// Log any remaining unlogged blocks (final blocks) - keep this as final safety net
	s.blockStatsMutex.Lock()
	for bn, stats := range s.blockStats {
		if bn > s.lastLoggedBlock && stats.ContractCount > 0 {
			avgGasPerByte := float64(stats.TotalGasUsed) / float64(max(stats.TotalBytecodeSize, 1))

			s.logger.WithFields(logrus.Fields{
				"block_number":        bn,
				"contracts_deployed":  stats.ContractCount,
				"total_gas_used":      stats.TotalGasUsed,
				"total_bytecode_size": stats.TotalBytecodeSize,
				"avg_gas_per_byte":    fmt.Sprintf("%.2f", avgGasPerByte),
				"total_contracts":     len(s.deployedContracts),
			}).Info("Block deployment summary")
		}
	}
	s.blockStatsMutex.Unlock()

	// Log final summary
	s.contractsMutex.Lock()
	totalContracts := len(s.deployedContracts)
	s.contractsMutex.Unlock()

	s.logger.WithFields(logrus.Fields{
		"total_txs":       totalTxCount.Load(),
		"total_contracts": totalContracts,
	}).Info("All transactions completed")

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

	// Get the latest block to check current base fee
	latestBlock, err := ethClient.BlockByNumber(ctx, nil)
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

// triggerImmediateTxSending forces immediate transaction sending with escalated gas prices
func (s *Scenario) triggerImmediateTxSending(ctx context.Context, stuckTxCount int) {
	s.logger.Warnf("FORCING immediate transaction sending due to %d stuck transactions", stuckTxCount)
	
	// Start from wallet group 0 and send transactions for all groups
	for groupIdx := uint64(0); groupIdx < 10; groupIdx++ {
		groupStartIdx := groupIdx * s.walletsPerGroup
		groupEndIdx := groupStartIdx + s.walletsPerGroup
		
		s.logger.Infof("FORCING transactions for wallet group %d of 10 (wallets %d-%d) with escalated prices",
			groupIdx, groupStartIdx, groupEndIdx-1)
		
		// Send batch with escalated prices
		err := s.sendBatchTransactions(ctx, 0, groupStartIdx, groupEndIdx) // Use 0 for baseIdx since we're forcing
		if err != nil {
			s.logger.Errorf("Failed to send forced batch for group %d: %v", groupIdx, err)
		} else {
			s.logger.Infof("Successfully sent forced batch for wallet group %d", groupIdx)
		}
		
		// Small delay between groups to avoid overwhelming the network
		time.Sleep(1 * time.Second)
		
		// Check context cancellation between groups
		select {
		case <-ctx.Done():
			s.logger.Warnf("Context cancelled during forced transaction sending")
			return
		default:
		}
	}
	
	s.logger.Infof("Completed forced transaction sending across all 10 wallet groups")
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

	// Use BuildBoundTx which handles nonce management internally
	tx, err := wallet.BuildBoundTx(ctx, &txbuilder.TxMetadata{
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

	// Send the transaction
	err = client.SendTransaction(ctx, tx)
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
