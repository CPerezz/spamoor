package erc20maxtransfers

import (
	"context"
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

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"

	contract "github.com/ethpandaops/spamoor/scenarios/statebloat/contract_deploy/contract"
	"github.com/ethpandaops/spamoor/scenario"
	"github.com/ethpandaops/spamoor/spamoor"
	"github.com/ethpandaops/spamoor/txbuilder"
)

// Constants for ERC20 transfer operations
const (
	// Gas costs for batch transfers (no events, no checks)
	BaseTransactionGas        = 21000  // Base transaction cost
	BatchFunctionOverhead     = 3000   // Function selector + array processing overhead
	GasPerBatchTransfer      = 52000  // Gas per transfer in batch (SSTORE + balance updates, no event) - measured from actual usage
	
	// Default fees
	DefaultBaseFeeGwei = 10
	DefaultTipFeeGwei  = 2
	
	// Token amounts
	TokenTransferAmount = 1 // 1 token in smallest unit
	
	// Gas optimization
	DefaultTargetGasRatio = 0.95  // Target 95% of block gas limit
	FallbackBlockGasLimit = 30000000
	
	// State growth tracking
	EstimatedStateGrowthPerTransfer = 100 // bytes per new recipient
	BloatingSummaryFileName         = "erc20_bloating_summary.json"
	
	// Timing constants
	BlockPollingInterval = 500 * time.Millisecond
	BlockMiningTimeout   = 30 * time.Second
)

// ScenarioOptions defines the configuration options for the scenario
type ScenarioOptions struct {
	BaseFee  uint64 `yaml:"base_fee"`
	TipFee   uint64 `yaml:"tip_fee"`
	Contract string `yaml:"contract"`
}

// DeploymentEntry represents a contract deployment from deployments.json
type DeploymentEntry map[string][]string

// ContractBloatStats tracks unique recipients per contract
type ContractBloatStats struct {
	UniqueRecipients int `json:"unique_recipients"`
}

// BloatingSummary represents the JSON file structure
type BloatingSummary struct {
	Contracts       map[string]*ContractBloatStats `json:"contracts"`
	TotalRecipients int                            `json:"total_recipients"`
	LastBlockNumber string                         `json:"last_block_number"`
	LastBlockUpdate time.Time                      `json:"last_block_update"`
}

// LogEntry represents a transfer batch log entry
type LogEntry struct {
	Contract    common.Address
	Recipients  []common.Address
	GasUsed     uint64
	BlockNumber uint64
	TxHash      common.Hash
	Timestamp   time.Time
}

// Scenario implements the ERC20 max transfers scenario
type Scenario struct {
	options    ScenarioOptions
	logger     *logrus.Entry
	walletPool *spamoor.WalletPool

	// Deployed contracts and deployer wallets
	deployerWallets    []*spamoor.Wallet
	deployedContracts  map[string][]common.Address // Private key -> contracts
	contractInstances  map[common.Address]*contract.Contract
	contractIndex      int // Round-robin index
	allContracts       []common.Address // Flattened list for round-robin

	// Contract ABI
	transferABI      abi.Method
	batchTransferABI abi.Method
	contractABI      abi.ABI

	// Async logging
	logChannel chan *LogEntry
	loggerWg   sync.WaitGroup

	// Statistics tracking
	totalRecipients   uint64
	totalGasUsed      uint64
	contractStats     map[common.Address]*ContractBloatStats
	contractStatsLock sync.Mutex

	// Gas optimization
	currentTransfersPerTx uint64
	blockGasLimit        uint64
}

var ScenarioName = "erc20-max-transfers"
var ScenarioDefaultOptions = ScenarioOptions{
	BaseFee:  DefaultBaseFeeGwei,
	TipFee:   DefaultTipFeeGwei,
	Contract: "",
}
var ScenarioDescriptor = scenario.Descriptor{
	Name:           ScenarioName,
	Description:    "Maximum ERC20 transfers per block to unique addresses",
	DefaultOptions: ScenarioDefaultOptions,
	NewScenario:    newScenario,
}

func newScenario(logger logrus.FieldLogger) scenario.Scenario {
	return &Scenario{
		logger:             logger.WithField("scenario", ScenarioName),
		contractStats:      make(map[common.Address]*ContractBloatStats),
		contractInstances:  make(map[common.Address]*contract.Contract),
		deployedContracts:  make(map[string][]common.Address),
		deployerWallets:    make([]*spamoor.Wallet, 0),
		allContracts:       make([]common.Address, 0),
		logChannel:         make(chan *LogEntry, 1000),
	}
}

func (s *Scenario) Flags(flags *pflag.FlagSet) error {
	flags.Uint64Var(&s.options.BaseFee, "basefee", ScenarioDefaultOptions.BaseFee, "Max fee per gas to use in transactions (in gwei)")
	flags.Uint64Var(&s.options.TipFee, "tipfee", ScenarioDefaultOptions.TipFee, "Max tip per gas to use in transactions (in gwei)")
	flags.StringVar(&s.options.Contract, "contract", ScenarioDefaultOptions.Contract, "Specific contract address to use (default: rotate through all)")
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

	// Load deployed contracts from deployments.json
	err := s.loadDeployedContracts()
	if err != nil {
		return fmt.Errorf("failed to load deployed contracts: %w", err)
	}

	// Load transfer function ABI
	err = s.loadTransferABI()
	if err != nil {
		return fmt.Errorf("failed to load transfer ABI: %w", err)
	}

	// Set up 100 child wallets for rotation
	s.walletPool.SetWalletCount(100)

	return nil
}

func (s *Scenario) Config() string {
	yamlBytes, _ := yaml.Marshal(&s.options)
	return string(yamlBytes)
}

// loadDeployedContracts loads contract addresses and creates deployer wallets from deployments.json
func (s *Scenario) loadDeployedContracts() error {
	data, err := os.ReadFile("deployments.json")
	if err != nil {
		return fmt.Errorf("failed to read deployments.json: %w", err)
	}

	var deployments map[string][]string
	err = json.Unmarshal(data, &deployments)
	if err != nil {
		return fmt.Errorf("failed to parse deployments.json: %w", err)
	}

	// Process all entries and create wallets
	totalContracts := 0
	for privateKey, addresses := range deployments {
		// Trim 0x prefix if present
		if strings.HasPrefix(privateKey, "0x") {
			privateKey = privateKey[2:]
		}

		// Create wallet from private key
		wallet, err := spamoor.NewWallet(privateKey)
		if err != nil {
			s.logger.Warnf("Failed to create wallet from private key: %v", err)
			continue
		}

		s.deployerWallets = append(s.deployerWallets, wallet)
		
		// Store contract addresses
		contractAddrs := make([]common.Address, len(addresses))
		for i, addr := range addresses {
			contractAddrs[i] = common.HexToAddress(addr)
			s.allContracts = append(s.allContracts, contractAddrs[i])
		}
		s.deployedContracts[privateKey] = contractAddrs
		totalContracts += len(addresses)
	}

	if len(s.deployerWallets) == 0 || totalContracts == 0 {
		return fmt.Errorf("no valid deployments found in deployments.json")
	}

	s.logger.Infof("Loaded %d deployer wallets with %d total contracts from deployments.json", 
		len(s.deployerWallets), totalContracts)

	// Initialize contract stats for all deployed contracts
	for _, contractAddr := range s.allContracts {
		s.contractStats[contractAddr] = &ContractBloatStats{
			UniqueRecipients: 0,
		}
	}

	// If specific contract requested, validate it exists
	if s.options.Contract != "" {
		contractAddr := common.HexToAddress(s.options.Contract)
		found := false
		for _, addr := range s.allContracts {
			if addr == contractAddr {
				found = true
				s.allContracts = []common.Address{contractAddr} // Use only this contract
				break
			}
		}
		if !found {
			return fmt.Errorf("specified contract %s not found in deployments", s.options.Contract)
		}
		s.logger.Infof("Using specific contract: %s", contractAddr.Hex())
	}

	return nil
}

// loadTransferABI loads the transfer and batchTransfer function ABIs from the contract
func (s *Scenario) loadTransferABI() error {
	// Parse the contract ABI to get the transfer methods
	contractABI, err := abi.JSON(strings.NewReader(contract.ContractMetaData.ABI))
	if err != nil {
		return fmt.Errorf("failed to parse contract ABI: %w", err)
	}

	transferMethod, exists := contractABI.Methods["transfer"]
	if !exists {
		return fmt.Errorf("transfer method not found in contract ABI")
	}

	batchTransferMethod, exists := contractABI.Methods["batchTransfer"]
	if !exists {
		return fmt.Errorf("batchTransfer method not found in contract ABI")
	}

	s.transferABI = transferMethod
	s.batchTransferABI = batchTransferMethod
	s.contractABI = contractABI
	return nil
}

// getNetworkBlockGasLimit retrieves the current block gas limit from the network
func (s *Scenario) getNetworkBlockGasLimit(ctx context.Context, client *spamoor.Client) uint64 {
	// Get the latest block
	block, err := client.GetEthClient().BlockByNumber(ctx, nil)
	if err != nil {
		s.logger.Warnf("failed to get latest block: %v, using fallback: %d", err, FallbackBlockGasLimit)
		return FallbackBlockGasLimit
	}

	gasLimit := block.GasLimit()
	s.logger.Debugf("network block gas limit: %d", gasLimit)
	return gasLimit
}

// generateRandomRecipient generates a random recipient address
func (s *Scenario) generateRandomRecipient() common.Address {
	var addr common.Address
	rand.Read(addr[:])
	return addr
}

// selectNextContract selects the next contract in round-robin fashion
func (s *Scenario) selectNextContract() common.Address {
	if len(s.allContracts) == 0 {
		return common.Address{}
	}
	
	contract := s.allContracts[s.contractIndex]
	s.contractIndex = (s.contractIndex + 1) % len(s.allContracts)
	return contract
}


// calculateOptimalTransfers calculates the optimal number of transfers per transaction
func (s *Scenario) calculateOptimalTransfers(blockGasLimit uint64) uint64 {
	// Calculate target gas (95% of block limit)
	targetGas := uint64(float64(blockGasLimit) * DefaultTargetGasRatio)
	
	// Calculate available gas for transfers
	availableGas := targetGas - BaseTransactionGas - BatchFunctionOverhead
	
	// Calculate number of transfers that fit
	transfers := availableGas / GasPerBatchTransfer
	
	// Apply reasonable bounds
	if transfers < 1 {
		transfers = 1
	} else if transfers > 1000 { // Cap at 1000 transfers per tx for safety
		transfers = 1000
	}
	
	return transfers
}

// initializeDeployerWallets updates all deployer wallets with chain info
func (s *Scenario) initializeDeployerWallets(ctx context.Context) error {
	client := s.walletPool.GetClient(spamoor.SelectClientByIndex, 0, "")
	if client == nil {
		return fmt.Errorf("no client available")
	}
	
	// Update all deployer wallets with chain info
	for i, wallet := range s.deployerWallets {
		err := client.UpdateWallet(ctx, wallet)
		if err != nil {
			return fmt.Errorf("failed to update deployer wallet %d: %w", i, err)
		}

		s.logger.Infof("Initialized deployer wallet %d - Address: %s, Nonce: %d, Balance: %s ETH",
			i, 
			wallet.GetAddress().Hex(), 
			wallet.GetNonce(), 
			new(big.Int).Div(wallet.GetBalance(), big.NewInt(1e18)).String())
	}
	
	return nil
}

// initializeContractInstances creates contract instances for all deployed contracts
func (s *Scenario) initializeContractInstances(ctx context.Context) error {
	if len(s.contractInstances) > 0 {
		return nil // Already initialized
	}

	client := s.walletPool.GetClient(spamoor.SelectClientByIndex, 0, "")
	if client == nil {
		return fmt.Errorf("no client available")
	}

	for _, contractAddr := range s.allContracts {
		contractInstance, err := contract.NewContract(contractAddr, client.GetEthClient())
		if err != nil {
			return fmt.Errorf("failed to create contract instance for %s: %w", contractAddr.Hex(), err)
		}
		s.contractInstances[contractAddr] = contractInstance
	}

	s.logger.Infof("Initialized %d contract instances", len(s.contractInstances))
	return nil
}

func (s *Scenario) Run(ctx context.Context) error {
	s.logger.Infof("starting scenario: %s", ScenarioName)
	defer s.logger.Infof("scenario %s finished.", ScenarioName)

	// Initialize deployer wallets if needed
	if len(s.deployerWallets) == 0 {
		return fmt.Errorf("no deployer wallets available")
	}
	
	if err := s.initializeDeployerWallets(ctx); err != nil {
		return err
	}

	// Initialize contract instances
	if err := s.initializeContractInstances(ctx); err != nil {
		return err
	}

	// Start async logger
	s.loggerWg.Add(1)
	go s.asyncLogger(ctx)

	defer func() {
		close(s.logChannel)
		s.loggerWg.Wait()
	}()

	// Get initial block gas limit
	client := s.walletPool.GetClient(spamoor.SelectClientByIndex, 0, "")
	if client == nil {
		return fmt.Errorf("no client available")
	}
	
	s.blockGasLimit = s.getNetworkBlockGasLimit(ctx, client)
	// Calculate initial transfers per transaction
	s.currentTransfersPerTx = s.calculateOptimalTransfers(s.blockGasLimit)
	
	s.logger.Infof("Block gas limit: %d, starting with %d transfers per tx", 
		s.blockGasLimit, s.currentTransfersPerTx)

	// Setup child wallets with ETH and tokens
	s.logger.Info("Setting up child wallets...")
	if err := s.setupChildWallets(ctx); err != nil {
		return fmt.Errorf("failed to setup child wallets: %w", err)
	}

	// Run custom block-based loop
	return s.runBlockBasedLoop(ctx)
}

// setupChildWallets funds child wallets with ETH and distributes tokens to them
func (s *Scenario) setupChildWallets(ctx context.Context) error {
	// High fees for setup transactions to ensure quick inclusion
	setupBaseFee := uint64(100) // 100 gwei
	setupTipFee := uint64(20)   // 20 gwei
	
	// Get child wallets by iterating through wallet count
	walletCount := s.walletPool.GetWalletCount()
	if walletCount == 0 {
		return fmt.Errorf("no child wallets available")
	}
	
	childWallets := make([]*spamoor.Wallet, walletCount)
	for i := uint64(0); i < walletCount; i++ {
		wallet := s.walletPool.GetWallet(spamoor.SelectWalletByIndex, int(i))
		if wallet == nil {
			return fmt.Errorf("failed to get wallet at index %d", i)
		}
		childWallets[i] = wallet
	}
	
	s.logger.Infof("Setting up %d child wallets...", len(childWallets))
	
	// Phase 1: Fund child wallets with ETH
	s.logger.Info("Phase 1: Funding child wallets with ETH...")
	if err := s.fundChildWalletsWithETH(ctx, childWallets, setupBaseFee, setupTipFee); err != nil {
		return fmt.Errorf("failed to fund child wallets with ETH: %w", err)
	}
	
	// Phase 2: Distribute tokens to child wallets
	s.logger.Info("Phase 2: Distributing tokens to child wallets...")
	if err := s.distributeTokensToChildWallets(ctx, childWallets, setupBaseFee, setupTipFee); err != nil {
		return fmt.Errorf("failed to distribute tokens: %w", err)
	}
	
	s.logger.Info("Child wallet setup completed successfully")
	return nil
}

// fundChildWalletsWithETH sends ETH to all child wallets for gas
func (s *Scenario) fundChildWalletsWithETH(ctx context.Context, childWallets []*spamoor.Wallet, baseFee, tipFee uint64) error {
	client := s.walletPool.GetClient(spamoor.SelectClientByIndex, 0, "")
	if client == nil {
		return fmt.Errorf("no client available")
	}
	
	// Amount of ETH to send to each child wallet (enough for many transactions)
	ethAmount := new(big.Int).Mul(big.NewInt(10000), big.NewInt(1e18)) // 10,000 ETH per wallet
	
	// Get root wallet for funding
	rootWallet := s.walletPool.GetRootWallet()
	if rootWallet == nil {
		return fmt.Errorf("no root wallet available")
	}
	
	s.logger.Infof("Funding %d child wallets with %s ETH each from root wallet %s",
		len(childWallets), new(big.Float).Quo(new(big.Float).SetInt(ethAmount), big.NewFloat(1e18)).String(),
		rootWallet.GetWallet().GetAddress().Hex())
	
	// Send ETH to each child wallet
	pendingTxs := make(map[common.Hash]*types.Transaction)
	for i, childWallet := range childWallets {
		// Get suggested fees with our high setup values
		feeCap, tipCap, err := s.walletPool.GetTxPool().GetSuggestedFees(client, baseFee, tipFee)
		if err != nil {
			return fmt.Errorf("failed to get suggested fees: %w", err)
		}
		
		// Build ETH transfer transaction using root wallet's wallet
		childAddr := childWallet.GetAddress()
		rootWalletWrapper := rootWallet.GetWallet()
		tx, err := rootWalletWrapper.BuildDynamicFeeTx(&types.DynamicFeeTx{
			To:        &childAddr,
			Value:     ethAmount,
			GasFeeCap: feeCap,
			GasTipCap: tipCap,
			Gas:       21000, // Standard ETH transfer
		})
		if err != nil {
			return fmt.Errorf("failed to build ETH transfer for wallet %d: %w", i, err)
		}
		
		// Send transaction
		err = s.walletPool.GetTxPool().SendTransaction(ctx, rootWalletWrapper, tx, &spamoor.SendTransactionOptions{
			Client:      client,
			Rebroadcast: true,
		})
		if err != nil {
			return fmt.Errorf("failed to send ETH to wallet %d: %w", i, err)
		}
		
		pendingTxs[tx.Hash()] = tx
		
		if (i+1)%10 == 0 || i == len(childWallets)-1 {
			s.logger.Infof("Sent ETH funding to %d/%d child wallets", i+1, len(childWallets))
		}
	}
	
	// Wait for all ETH transfers to confirm
	s.logger.Info("Waiting for ETH funding transactions to confirm...")
	return s.waitForTransactions(ctx, client, pendingTxs, "ETH funding")
}

// distributeTokensToChildWallets sends tokens from deployer wallets to child wallets
func (s *Scenario) distributeTokensToChildWallets(ctx context.Context, childWallets []*spamoor.Wallet, baseFee, tipFee uint64) error {
	client := s.walletPool.GetClient(spamoor.SelectClientByIndex, 0, "")
	if client == nil {
		return fmt.Errorf("no client available")
	}
	
	// Amount of tokens to send to each child wallet
	tokenAmount := big.NewInt(1000000) // 1M tokens
	tokenAmount.Mul(tokenAmount, big.NewInt(1e18)) // Convert to smallest unit
	
	// Calculate how many child wallets each deployer should fund
	walletsPerDeployer := len(childWallets) / len(s.deployerWallets)
	if len(childWallets)%len(s.deployerWallets) > 0 {
		walletsPerDeployer++
	}
	
	s.logger.Infof("Distributing tokens: %d deployers will fund up to %d child wallets each",
		len(s.deployerWallets), walletsPerDeployer)
	
	pendingTxs := make(map[common.Hash]*types.Transaction)
	childIndex := 0
	
	// Distribute child wallets among deployer wallets
	for deployerIdx, deployerWallet := range s.deployerWallets {
		// Find contracts owned by this deployer
		var ownedContracts []common.Address
		for privateKey, contracts := range s.deployedContracts {
			// Check if this deployer wallet matches the private key
			deployerPrivKey := deployerWallet.GetPrivateKey()
			if deployerPrivKey != nil && deployerPrivKey.D.Text(16) == privateKey {
				ownedContracts = append(ownedContracts, contracts...)
			}
		}
		
		if len(ownedContracts) == 0 {
			continue // This deployer has no contracts
		}
		
		// Use the first contract from this deployer for token transfers
		contractAddr := ownedContracts[0]
		contractInstance := s.contractInstances[contractAddr]
		
		// Send tokens to assigned child wallets
		for i := 0; i < walletsPerDeployer && childIndex < len(childWallets); i++ {
			childWallet := childWallets[childIndex]
			childIndex++
			
			// Get suggested fees
			feeCap, tipCap, err := s.walletPool.GetTxPool().GetSuggestedFees(client, baseFee, tipFee)
			if err != nil {
				return fmt.Errorf("failed to get suggested fees: %w", err)
			}
			
			// Build token transfer transaction
			tx, err := deployerWallet.BuildBoundTx(ctx, &txbuilder.TxMetadata{
				To:        &contractAddr,
				GasFeeCap: uint256.MustFromBig(feeCap),
				GasTipCap: uint256.MustFromBig(tipCap),
				Gas:       100000, // Estimated gas for token transfer
			}, func(transactOpts *bind.TransactOpts) (*types.Transaction, error) {
				return contractInstance.Transfer(transactOpts, childWallet.GetAddress(), tokenAmount)
			})
			if err != nil {
				return fmt.Errorf("failed to build token transfer: %w", err)
			}
			
			// Send transaction
			err = s.walletPool.GetTxPool().SendTransaction(ctx, deployerWallet, tx, &spamoor.SendTransactionOptions{
				Client:      client,
				Rebroadcast: true,
			})
			if err != nil {
				return fmt.Errorf("failed to send tokens: %w", err)
			}
			
			pendingTxs[tx.Hash()] = tx
			
			if childIndex%10 == 0 || childIndex == len(childWallets) {
				s.logger.Infof("Sent tokens to %d/%d child wallets", childIndex, len(childWallets))
			}
		}
		
		// Break if all child wallets have been funded
		if childIndex >= len(childWallets) {
			break
		}
		
		// Log progress
		if (deployerIdx+1)%10 == 0 {
			s.logger.Debugf("Processed %d/%d deployer wallets", deployerIdx+1, len(s.deployerWallets))
		}
	}
	
	// Wait for all token transfers to confirm
	s.logger.Info("Waiting for token distribution transactions to confirm...")
	return s.waitForTransactions(ctx, client, pendingTxs, "token distribution")
}

// waitForTransactions waits for a set of transactions to be confirmed
func (s *Scenario) waitForTransactions(ctx context.Context, client *spamoor.Client, pendingTxs map[common.Hash]*types.Transaction, txType string) error {
	startTime := time.Now()
	confirmed := 0
	total := len(pendingTxs)
	
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	
	for len(pendingTxs) > 0 {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			for hash := range pendingTxs {
				receipt, err := client.GetEthClient().TransactionReceipt(ctx, hash)
				if err == nil && receipt != nil {
					if receipt.Status == types.ReceiptStatusSuccessful {
						confirmed++
						delete(pendingTxs, hash)
					} else {
						return fmt.Errorf("%s transaction failed: %s", txType, hash.Hex())
					}
				}
			}
			
			if confirmed > 0 {
				elapsed := time.Since(startTime)
				s.logger.Infof("%s progress: %d/%d confirmed (%.1fs elapsed)",
					txType, confirmed, total, elapsed.Seconds())
			}
			
			// Timeout after 5 minutes
			if time.Since(startTime) > 5*time.Minute {
				return fmt.Errorf("%s timeout: %d transactions still pending", txType, len(pendingTxs))
			}
		}
	}
	
	s.logger.Infof("%s completed: all %d transactions confirmed", txType, total)
	return nil
}

// runBlockBasedLoop implements the custom loop for one transaction per block
func (s *Scenario) runBlockBasedLoop(ctx context.Context) error {
	type pendingTxInfo struct {
		tx            *types.Transaction
		recipients    []common.Address
		contractAddr  common.Address
		sentAt        time.Time
		blockNumber   uint64
	}
	
	var txCount uint64
	var walletIndex uint64
	pendingTxs := make(map[common.Hash]*pendingTxInfo)
	pendingTxsMutex := sync.Mutex{}
	
	s.logger.Info("Starting block-based transaction loop...")
	
	// Start a goroutine to track transaction confirmations
	go func() {
		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop()
		
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				pendingTxsMutex.Lock()
				for hash, txInfo := range pendingTxs {
					client := s.walletPool.GetClient(spamoor.SelectClientByIndex, 0, "")
					if client == nil {
						continue
					}
					
					receipt, err := client.GetEthClient().TransactionReceipt(ctx, hash)
					if err == nil && receipt != nil {
						// Transaction confirmed
						if receipt.Status == types.ReceiptStatusSuccessful {
							// Send log entry for async processing
							s.logChannel <- &LogEntry{
								Contract:    txInfo.contractAddr,
								Recipients:  txInfo.recipients,
								GasUsed:     receipt.GasUsed,
								BlockNumber: receipt.BlockNumber.Uint64(),
								TxHash:      hash,
								Timestamp:   time.Now(),
							}
							
							// Update metrics
							atomic.AddUint64(&s.totalRecipients, uint64(len(txInfo.recipients)))
							atomic.AddUint64(&s.totalGasUsed, receipt.GasUsed)
							
							// Update contract stats
							s.updateContractStats(txInfo.contractAddr, len(txInfo.recipients))
							
							// Log confirmation with gas usage
							gasPercent := float64(receipt.GasUsed) * 100 / float64(s.blockGasLimit)
							s.logger.WithFields(logrus.Fields{
								"block":       receipt.BlockNumber.Uint64(),
								"tx":          hash.Hex()[:10],
								"recipients":  len(txInfo.recipients),
								"gas_used":    fmt.Sprintf("%dM", receipt.GasUsed/1_000_000),
								"gas_percent": fmt.Sprintf("%.1f%%", gasPercent),
								"duration":    fmt.Sprintf("%.1fs", time.Since(txInfo.sentAt).Seconds()),
							}).Info("Transaction confirmed")
						} else {
							s.logger.WithFields(logrus.Fields{
								"tx":    hash.Hex()[:10],
								"block": receipt.BlockNumber.Uint64(),
							}).Warn("Transaction failed")
						}
						
						delete(pendingTxs, hash)
					}
				}
				pendingTxsMutex.Unlock()
			}
		}
	}()
	
	// Main loop to send transactions continuously
	lastBlockNumber := uint64(0)
	blockTicker := time.NewTicker(BlockPollingInterval)
	defer blockTicker.Stop()
	
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-blockTicker.C:
			// Get client
			client := s.walletPool.GetClient(spamoor.SelectClientByIndex, 0, "")
			if client == nil {
				s.logger.Warn("no client available, waiting...")
				continue
			}

			// Check current block number
			header, err := client.GetEthClient().HeaderByNumber(ctx, nil)
			if err != nil {
				s.logger.Warnf("failed to get latest block header: %v", err)
				continue
			}
			
			currentBlockNumber := header.Number.Uint64()
			
			// Only send transaction if we're on a new block
			if currentBlockNumber <= lastBlockNumber {
				continue
			}
			
			lastBlockNumber = currentBlockNumber
			
			// Update gas limit periodically
			if txCount%10 == 0 {
				newGasLimit := s.getNetworkBlockGasLimit(ctx, client)
				if newGasLimit != s.blockGasLimit {
					s.blockGasLimit = newGasLimit
					s.currentTransfersPerTx = s.calculateOptimalTransfers(newGasLimit)
					s.logger.Infof("Block gas limit changed to %dM, adjusting transfers to %d",
						newGasLimit/1_000_000, s.currentTransfersPerTx)
				}
			}
			
			// Select next contract and child wallet
			contractAddr := s.selectNextContract()
			if contractAddr == (common.Address{}) {
				s.logger.Error("no contracts available")
				continue
			}
			
			// Get next child wallet in rotation
			childWallet := s.walletPool.GetWallet(spamoor.SelectWalletByIndex, int(walletIndex))
			if childWallet == nil {
				s.logger.Errorf("no child wallet available at index %d", walletIndex)
				continue
			}
			walletIndex = (walletIndex + 1) % uint64(s.walletPool.GetWalletCount())
			
			// Log that we're sending a transaction for this block
			targetGasPercent := DefaultTargetGasRatio * 100
			s.logger.WithFields(logrus.Fields{
				"block":       currentBlockNumber,
				"transfers":   s.currentTransfersPerTx,
				"target_gas":  fmt.Sprintf("%.0f%%", targetGasPercent),
				"wallet":      childWallet.GetAddress().Hex()[:10],
				"contract":    contractAddr.Hex()[:10],
			}).Info("Sending transaction for new block")
			
			// Build and send transaction without waiting
			tx, recipients, err := s.buildMultiTransferTransaction(ctx, contractAddr, childWallet, client)
			if err != nil {
				s.logger.Warnf("failed to build transaction: %v", err)
				continue
			}
			
			// Send transaction without waiting for confirmation
			err = s.walletPool.GetTxPool().SendTransaction(ctx, childWallet, tx, &spamoor.SendTransactionOptions{
				Client:      client,
				Rebroadcast: true,
			})
			if err != nil {
				s.logger.Warnf("failed to send transaction: %v", err)
				// Reset nonce on error
				childWallet.ResetPendingNonce(ctx, client)
				continue
			}
			
			txCount++
			pendingTxsMutex.Lock()
			pendingTxs[tx.Hash()] = &pendingTxInfo{
				tx:           tx,
				recipients:   recipients,
				contractAddr: contractAddr,
				sentAt:       time.Now(),
				blockNumber:  currentBlockNumber,
			}
			pendingCount := len(pendingTxs)
			pendingTxsMutex.Unlock()
			
			s.logger.WithFields(logrus.Fields{
				"tx":         tx.Hash().Hex()[:10],
				"block":      currentBlockNumber,
				"pending":    pendingCount,
				"total_sent": txCount,
			}).Info("Transaction submitted")
		}
	}
}


// buildMultiTransferTransaction creates a transaction containing multiple transfers
func (s *Scenario) buildMultiTransferTransaction(ctx context.Context, contractAddr common.Address, wallet *spamoor.Wallet, client *spamoor.Client) (*types.Transaction, []common.Address, error) {
	// Get suggested fees
	feeCap, tipCap, err := s.walletPool.GetTxPool().GetSuggestedFees(client, s.options.BaseFee, s.options.TipFee)
	if err != nil {
		return nil, nil, err
	}

	// Generate recipients
	recipients := make([]common.Address, s.currentTransfersPerTx)
	for i := range recipients {
		recipients[i] = s.generateRandomRecipient()
	}

	// Get the contract instance
	contractInstance, exists := s.contractInstances[contractAddr]
	if !exists {
		return nil, nil, fmt.Errorf("no contract instance found for %s", contractAddr.Hex())
	}

	// Calculate gas needed for batch transfer
	gasLimit := uint64(BaseTransactionGas + BatchFunctionOverhead + (GasPerBatchTransfer * uint64(len(recipients))))

	// Build transaction using BuildBoundTx with contract instance
	tx, err := wallet.BuildBoundTx(ctx, &txbuilder.TxMetadata{
		To:        &contractAddr,
		GasFeeCap: uint256.MustFromBig(feeCap),
		GasTipCap: uint256.MustFromBig(tipCap),
		Gas:       gasLimit,
	}, func(transactOpts *bind.TransactOpts) (*types.Transaction, error) {
		return contractInstance.BatchTransfer(transactOpts, recipients)
	})

	return tx, recipients, err
}


// updateContractStats updates the statistics for a contract when transfers are confirmed
func (s *Scenario) updateContractStats(contractAddr common.Address, recipientCount int) {
	s.contractStatsLock.Lock()
	defer s.contractStatsLock.Unlock()

	stats, exists := s.contractStats[contractAddr]
	if !exists {
		stats = &ContractBloatStats{
			UniqueRecipients: 0,
		}
		s.contractStats[contractAddr] = stats
	}
	stats.UniqueRecipients += recipientCount
}

// loadBloatingSummary loads the bloating summary from file or creates a new one
func (s *Scenario) loadBloatingSummary() (*BloatingSummary, error) {
	data, err := os.ReadFile(BloatingSummaryFileName)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist, return new summary
			return &BloatingSummary{
				Contracts:       make(map[string]*ContractBloatStats),
				TotalRecipients: 0,
			}, nil
		}
		return nil, fmt.Errorf("failed to read bloating summary: %w", err)
	}

	var summary BloatingSummary
	if err := json.Unmarshal(data, &summary); err != nil {
		return nil, fmt.Errorf("failed to unmarshal bloating summary: %w", err)
	}

	// Ensure contracts map is initialized
	if summary.Contracts == nil {
		summary.Contracts = make(map[string]*ContractBloatStats)
	}

	return &summary, nil
}

// saveBloatingSummary saves the bloating summary to file
func (s *Scenario) saveBloatingSummary(summary *BloatingSummary) error {
	data, err := json.MarshalIndent(summary, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal bloating summary: %w", err)
	}

	if err := os.WriteFile(BloatingSummaryFileName, data, 0644); err != nil {
		return fmt.Errorf("failed to write bloating summary: %w", err)
	}

	return nil
}

// updateAndSaveBloatingSummary updates the bloating summary with current stats and saves to file
func (s *Scenario) updateAndSaveBloatingSummary(blockNumber string) error {
	// Load existing summary
	summary, err := s.loadBloatingSummary()
	if err != nil {
		return err
	}

	// Update with current stats
	s.contractStatsLock.Lock()
	totalRecipients := 0
	for contractAddr, stats := range s.contractStats {
		contractHex := contractAddr.Hex()
		summary.Contracts[contractHex] = &ContractBloatStats{
			UniqueRecipients: stats.UniqueRecipients,
		}
		totalRecipients += stats.UniqueRecipients
	}
	s.contractStatsLock.Unlock()

	// Update summary metadata
	summary.TotalRecipients = totalRecipients
	summary.LastBlockNumber = blockNumber
	summary.LastBlockUpdate = time.Now()

	// Save to file
	return s.saveBloatingSummary(summary)
}

// asyncLogger handles asynchronous logging of transfer metrics
func (s *Scenario) asyncLogger(ctx context.Context) {
	defer s.loggerWg.Done()

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	logs := make([]*LogEntry, 0, 100)
	var lastBlockNumber uint64

	for {
		select {
		case <-ctx.Done():
			s.flushLogs(logs)
			return
		case log := <-s.logChannel:
			if log == nil {
				s.flushLogs(logs)
				return
			}
			logs = append(logs, log)
			if log.BlockNumber > lastBlockNumber {
				lastBlockNumber = log.BlockNumber
			}
		case <-ticker.C:
			s.flushLogs(logs)
			// Update bloating summary
			if lastBlockNumber > 0 {
				if err := s.updateAndSaveBloatingSummary(fmt.Sprintf("%d", lastBlockNumber)); err != nil {
					s.logger.Warnf("Failed to update bloating summary: %v", err)
				}
			}
			logs = logs[:0]
		}
	}
}

// flushLogs processes and logs accumulated transfer data
func (s *Scenario) flushLogs(logs []*LogEntry) {
	if len(logs) == 0 {
		return
	}

	// Calculate totals
	var totalTransfers int
	var totalGasUsed uint64
	contractTransfers := make(map[common.Address]int)
	blockStats := make(map[uint64]int)

	for _, log := range logs {
		totalTransfers += len(log.Recipients)
		totalGasUsed += log.GasUsed
		contractTransfers[log.Contract] += len(log.Recipients)
		blockStats[log.BlockNumber]++
	}

	// Get current metrics
	currentRecipients := atomic.LoadUint64(&s.totalRecipients)
	avgGasPerTransfer := float64(totalGasUsed) / float64(totalTransfers)
	avgGasPerTx := totalGasUsed / uint64(len(logs))
	
	// Calculate state growth estimate
	stateGrowthMB := float64(currentRecipients * EstimatedStateGrowthPerTransfer) / (1024 * 1024)

	s.logger.WithFields(logrus.Fields{
		"period_txs":        len(logs),
		"period_transfers":  totalTransfers,
		"avg_gas_per_tx":    fmt.Sprintf("%dM", avgGasPerTx/1_000_000),
		"avg_gas_per_xfer":  fmt.Sprintf("%.0f", avgGasPerTransfer),
		"total_recipients":  currentRecipients,
		"est_state_growth":  fmt.Sprintf("%.1fMB", stateGrowthMB),
		"blocks_with_txs":   len(blockStats),
	}).Info("ERC20 transfer summary")
}