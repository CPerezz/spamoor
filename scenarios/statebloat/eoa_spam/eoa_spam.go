package eoa_spam

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"

	"github.com/ethpandaops/spamoor/scenario"
	"github.com/ethpandaops/spamoor/scenarios/statebloat/eoa_spam/contract"
	"github.com/ethpandaops/spamoor/spamoor"
	"github.com/ethpandaops/spamoor/txbuilder"
)

// Helper functions for min/max
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

const ScenarioName = "eoa-spam"

const (
	// Gas cost per iteration in the spam contract (EOA creation cost + safety margin)
	// Based on actual analysis: keccak256(~3k) + external call(~21k) + account creation(~25k) + overhead(~2k)
	GasPerIteration = 51000  // Corrected from 26000 based on actual contract operations
	// Base gas for transaction overhead (includes contract overhead + refund logic)
	BaseTransactionGas = 60000  // Increased from 40000 for additional safety margin
	// Initial target gas ratio (start conservative)
	InitialTargetGasRatio = 0.35  // Reduced from 0.45 for more conservative start
	// Maximum target gas ratio
	MaxTargetGasRatio = 0.85  // Reduced from 0.90 for safety
	// Target gas ratio increment
	TargetGasRatioIncrement = 0.05  // Reduced from 0.10 for more gradual increases
	// Fallback block gas limit if network query fails
	FallbackBlockGasLimit = 30000000
	// Block polling interval
	BlockPollingInterval = 500 * time.Millisecond
	// Timeout for waiting for new block
	BlockMiningTimeout = 30 * time.Second
	// Minimum iterations to prevent transactions becoming too small
	MinViableIterations = 10
)

var ScenarioDefaultOptions = ScenarioOptions{
	TotalChildWallets: 50,
	MinIterations:     100,
	MaxIterations:     100000,
	MaxWallets:        100,
	BaseFee:           50,
	TipFee:            2,
	TargetGasRatio:    InitialTargetGasRatio,
	GasPerIteration:   GasPerIteration,
}

var ScenarioDescriptor = scenario.Descriptor{
	Name:           ScenarioName,
	Description:    "Deploy contracts that fund new EOAs with 1 wei using keccak256 address generation",
	DefaultOptions: ScenarioDefaultOptions,
	NewScenario:    newScenario,
}

func newScenario(logger logrus.FieldLogger) scenario.Scenario {
	return &Scenario{
		logger:     logger.WithField("scenario", ScenarioName),
		logChannel: make(chan *LogEntry, 1000),
	}
}

type ScenarioOptions struct {
	TotalChildWallets uint64  `yaml:"total_child_wallets"`
	MinIterations     uint64  `yaml:"min_iterations"`
	MaxIterations     uint64  `yaml:"max_iterations"`
	MaxWallets        uint64  `yaml:"max_wallets"`
	BaseFee           uint64  `yaml:"base_fee"`
	TipFee            uint64  `yaml:"tip_fee"`
	TargetGasRatio    float64 `yaml:"target_gas_ratio"`
	GasPerIteration   uint64  `yaml:"gas_per_iteration"`
}

type Scenario struct {
	options           ScenarioOptions
	logger            *logrus.Entry
	walletPool        *spamoor.WalletPool
	contractAddress   common.Address
	contractInstance  *contract.Contract
	isDeployed        bool
	deployMutex       sync.Mutex
	currentIterations uint64
	iterationsMtx     sync.RWMutex
	logChannel        chan *LogEntry
	loggerWg          sync.WaitGroup
	blockGasLimit     uint64
	gasLimitMtx       sync.RWMutex
	lastGasUsed       uint64
	walletIndex       int
	
	// Dynamic gas ratio management
	targetGasRatio       float64
	consecutiveSuccesses uint64
	consecutiveFailures  uint64
	pendingTxs          map[common.Hash]*PendingTx
	pendingTxsMtx       sync.RWMutex
	
	// Dynamic gas learning
	actualGasPerIteration float64
	gasEstimatesMtx       sync.RWMutex
	gasEstimateHistory    []GasEstimate
	failureBackoffLevel   int
}

type LogEntry struct {
	Wallet     common.Address
	Contract   common.Address
	Iterations uint64
	TxHash     common.Hash
	Success    bool
	Timestamp  time.Time
}

type PendingTx struct {
	Iterations  uint64
	BlockNumber uint64
	Timestamp   time.Time
}

type GasEstimate struct {
	Iterations    uint64
	EstimatedGas  uint64
	ActualGas     uint64
	Success       bool
	Timestamp     time.Time
}


func (s *Scenario) Flags(flags *pflag.FlagSet) error {
	flags.Uint64Var(&s.options.TotalChildWallets, "total-child-wallets", ScenarioDefaultOptions.TotalChildWallets, "Number of child wallets to use")
	flags.Uint64Var(&s.options.MinIterations, "min-iterations", ScenarioDefaultOptions.MinIterations, "Minimum number of iterations per transaction")
	flags.Uint64Var(&s.options.MaxIterations, "max-iterations", ScenarioDefaultOptions.MaxIterations, "Maximum iterations per transaction")
	flags.Uint64Var(&s.options.MaxWallets, "max-wallets", ScenarioDefaultOptions.MaxWallets, "Maximum concurrent wallets")
	flags.Uint64Var(&s.options.BaseFee, "basefee", ScenarioDefaultOptions.BaseFee, "Max fee per gas to use in transactions (in gwei)")
	flags.Uint64Var(&s.options.TipFee, "tipfee", ScenarioDefaultOptions.TipFee, "Max tip per gas to use in transactions (in gwei)")
	flags.Float64Var(&s.options.TargetGasRatio, "target-gas-ratio", ScenarioDefaultOptions.TargetGasRatio, "Target gas usage ratio of block limit (0.0-1.0)")
	flags.Uint64Var(&s.options.GasPerIteration, "gas-per-iteration", ScenarioDefaultOptions.GasPerIteration, "Estimated gas cost per iteration")
	return nil
}

func (s *Scenario) Config(config interface{}) error {
	optBytes, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %v", err)
	}
	err = yaml.Unmarshal(optBytes, &s.options)
	if err != nil {
		return fmt.Errorf("failed to parse config: %v", err)
	}
	
	if s.options.BaseFee == 0 {
		s.options.BaseFee = ScenarioDefaultOptions.BaseFee
	}
	if s.options.TipFee == 0 {
		s.options.TipFee = ScenarioDefaultOptions.TipFee
	}
	
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

	if s.options.TotalChildWallets > s.options.MaxWallets {
		s.options.TotalChildWallets = s.options.MaxWallets
	}

	if s.options.MaxWallets > 0 {
		s.walletPool.SetWalletCount(s.options.MaxWallets)
	} else {
		s.walletPool.SetWalletCount(s.options.TotalChildWallets)
	}

	// Initialize dynamic gas management
	s.targetGasRatio = s.options.TargetGasRatio
	s.pendingTxs = make(map[common.Hash]*PendingTx)
	s.currentIterations = 100 // Start with minimum
	
	// Initialize dynamic gas learning
	s.actualGasPerIteration = float64(GasPerIteration) // Start with constant estimate
	s.gasEstimateHistory = make([]GasEstimate, 0, 50)  // Keep last 50 estimates
	s.failureBackoffLevel = 0

	return nil
}

// deployContract deploys the EOASpammer contract using the first child wallet
func (s *Scenario) deployContract(ctx context.Context) error {
	s.deployMutex.Lock()
	defer s.deployMutex.Unlock()

	if s.isDeployed {
		return nil
	}

	s.logger.Info("Deploying EOASpammer contract...")

	client := s.walletPool.GetClient(spamoor.SelectClientByIndex, 0, "")
	if client == nil {
		return fmt.Errorf("no client available")
	}

	// Use first child wallet for deployment
	wallet := s.walletPool.GetWallet(spamoor.SelectWalletByIndex, 0)
	if wallet == nil {
		return fmt.Errorf("no wallet available")
	}

	// Get suggested fees from the network
	feeCap, tipCap, err := s.walletPool.GetTxPool().GetSuggestedFees(client, s.options.BaseFee, s.options.TipFee)
	if err != nil {
		return fmt.Errorf("failed to get suggested fees: %w", err)
	}

	// Build deployment transaction using BuildBoundTx
	var contractAddr common.Address
	var contractInst *contract.Contract
	
	tx, err := wallet.BuildBoundTx(ctx, &txbuilder.TxMetadata{
		GasFeeCap: uint256.MustFromBig(feeCap),
		GasTipCap: uint256.MustFromBig(tipCap),
		Gas:       3000000, // Standard deployment gas
	}, func(transactOpts *bind.TransactOpts) (*types.Transaction, error) {
		addr, deployTx, cInst, err := contract.DeployContract(transactOpts, client.GetEthClient())
		if err != nil {
			return nil, err
		}
		contractAddr = addr
		contractInst = cInst
		return deployTx, nil
	})
	if err != nil {
		return fmt.Errorf("failed to build deployment transaction: %w", err)
	}

	s.logger.WithField("tx", tx.Hash().Hex()).Info("Contract deployment transaction sent")

	// Send and wait for deployment
	err = s.walletPool.GetTxPool().SendTransaction(ctx, wallet, tx, &spamoor.SendTransactionOptions{
		Client:      client,
		Rebroadcast: true,
		OnConfirm: func(tx *types.Transaction, receipt *types.Receipt) {
			if receipt.Status == types.ReceiptStatusSuccessful {
				s.logger.WithField("address", contractAddr.Hex()).Info("EOASpammer contract deployed successfully")
			}
		},
	})
	if err != nil {
		return fmt.Errorf("failed to send deployment transaction: %w", err)
	}

	// Store contract info
	s.contractAddress = contractAddr
	s.contractInstance = contractInst
	s.isDeployed = true

	// Wait a bit for contract to be fully deployed
	time.Sleep(3 * time.Second)

	return nil
}

// getNetworkBlockGasLimit queries the network for the current block gas limit
func (s *Scenario) getNetworkBlockGasLimit(ctx context.Context, client *spamoor.Client) uint64 {
	// Create a timeout context for the entire operation
	timeoutCtx, cancel := context.WithTimeout(ctx, BlockMiningTimeout)
	defer cancel()

	// Get the current block
	block, err := client.GetEthClient().BlockByNumber(timeoutCtx, nil)
	if err != nil {
		s.logger.Warnf("failed to get current block: %v, using fallback: %d", err, FallbackBlockGasLimit)
		return FallbackBlockGasLimit
	}

	gasLimit := block.GasLimit()
	s.logger.Debugf("network block gas limit: %d", gasLimit)
	
	// Update cached value
	s.gasLimitMtx.Lock()
	s.blockGasLimit = gasLimit
	s.gasLimitMtx.Unlock()
	
	return gasLimit
}

// updateGasEstimate updates the dynamic gas learning with actual transaction results
func (s *Scenario) updateGasEstimate(iterations, estimatedGas, actualGas uint64, success bool) {
	s.gasEstimatesMtx.Lock()
	defer s.gasEstimatesMtx.Unlock()
	
	estimate := GasEstimate{
		Iterations:   iterations,
		EstimatedGas: estimatedGas,
		ActualGas:    actualGas,
		Success:      success,
		Timestamp:    time.Now(),
	}
	
	// Add to history
	s.gasEstimateHistory = append(s.gasEstimateHistory, estimate)
	
	// Keep only last 50 estimates
	if len(s.gasEstimateHistory) > 50 {
		s.gasEstimateHistory = s.gasEstimateHistory[1:]
	}
	
	// Update running average of actual gas per iteration (only successful txs)
	if success && iterations > 0 {
		baseGas := uint64(BaseTransactionGas)
		if actualGas > baseGas {
			actualGasPerIteration := float64(actualGas-baseGas) / float64(iterations)
			// Exponential moving average with alpha = 0.1
			s.actualGasPerIteration = 0.9*s.actualGasPerIteration + 0.1*actualGasPerIteration
		}
	}
}

// getDynamicGasPerIteration returns the current best estimate of gas per iteration
func (s *Scenario) getDynamicGasPerIteration() uint64 {
	s.gasEstimatesMtx.RLock()
	defer s.gasEstimatesMtx.RUnlock()
	
	// Use learned estimate with safety margin
	gasPerIteration := uint64(s.actualGasPerIteration * 1.05) // 5% safety margin
	
	// Bound it to reasonable values
	if gasPerIteration < 30000 {
		gasPerIteration = 30000
	} else if gasPerIteration > 80000 {
		gasPerIteration = 80000
	}
	
	return gasPerIteration
}

// calculateIterations calculates the optimal number of iterations based on block gas limit
func (s *Scenario) calculateIterations(blockGasLimit uint64) uint64 {
	// Use current dynamic target gas ratio
	targetGas := uint64(float64(blockGasLimit) * s.targetGasRatio)
	
	// Calculate iterations: (targetGas - baseGas) / gasPerIteration
	if targetGas <= BaseTransactionGas {
		return max(s.options.MinIterations, MinViableIterations)
	}
	
	// Use dynamic gas per iteration estimate
	gasPerIteration := s.getDynamicGasPerIteration()
	iterations := (targetGas - BaseTransactionGas) / gasPerIteration
	
	// Apply failure backoff
	if s.failureBackoffLevel > 0 {
		// Reduce iterations based on failure level
		reductionFactor := 1.0 - (float64(s.failureBackoffLevel) * 0.25)
		if reductionFactor < 0.25 {
			reductionFactor = 0.25
		}
		iterations = uint64(float64(iterations) * reductionFactor)
	}
	
	// Apply bounds
	minIterations := max(s.options.MinIterations, MinViableIterations)
	if iterations < minIterations {
		iterations = minIterations
	} else if iterations > s.options.MaxIterations {
		iterations = s.options.MaxIterations
	}
	
	return iterations
}

func (s *Scenario) Run(ctx context.Context) error {
	s.logger.Infof("starting scenario: %s", ScenarioName)
	defer s.logger.Infof("scenario %s finished.", ScenarioName)

	// Deploy the contract if not already deployed
	if !s.isDeployed {
		if err := s.deployContract(ctx); err != nil {
			return fmt.Errorf("failed to deploy contract: %w", err)
		}
	}

	s.loggerWg.Add(1)
	go s.asyncLogger(ctx)

	defer func() {
		close(s.logChannel)
		s.loggerWg.Wait()
	}()

	// Get initial block gas limit
	client := s.walletPool.GetClient(spamoor.SelectClientRoundRobin, 0, "")
	if client == nil {
		return fmt.Errorf("no client available")
	}
	
	blockGasLimit := s.getNetworkBlockGasLimit(ctx, client)
	s.currentIterations = s.calculateIterations(blockGasLimit)
	
	s.logger.Infof("Initial configuration:")
	s.logger.Infof("  Block gas limit: %d", blockGasLimit)
	s.logger.Infof("  Starting target gas ratio: %.0f%%", s.targetGasRatio*100)
	s.logger.Infof("  Initial iterations per tx: %d (creates %d EOAs)", s.currentIterations, s.currentIterations)
	s.logger.Infof("  Gas per iteration estimate: %d (corrected from 26k)", s.getDynamicGasPerIteration())
	s.logger.Infof("  Estimated gas per tx: %d", BaseTransactionGas + (s.currentIterations * s.getDynamicGasPerIteration()))
	s.logger.Infof("  Child wallets: %d", s.options.TotalChildWallets)

	// Run custom block-based loop for better gas utilization
	return s.runBlockBasedLoop(ctx)
}

// processTransaction is kept for compatibility but not used in the block-based loop
func (s *Scenario) processTransaction(ctx context.Context, txIdx uint64, onComplete func()) (*types.Transaction, *spamoor.Client, *spamoor.Wallet, error) {
	transactionSubmitted := false
	defer func() {
		if !transactionSubmitted {
			onComplete()
		}
	}()

	// Select wallet in round-robin fashion
	walletIdx := s.walletIndex % int(s.options.TotalChildWallets)
	s.walletIndex++
	
	wallet := s.walletPool.GetWallet(spamoor.SelectWalletByIndex, walletIdx)
	if wallet == nil {
		return nil, nil, nil, fmt.Errorf("wallet %d not available", walletIdx)
	}

	client := s.walletPool.GetClient(spamoor.SelectClientRoundRobin, 0, "")
	if client == nil {
		return nil, nil, wallet, fmt.Errorf("no client available")
	}

	// Check and update block gas limit periodically
	if txIdx%10 == 0 {
		newGasLimit := s.getNetworkBlockGasLimit(ctx, client)
		if newGasLimit != s.blockGasLimit {
			s.currentIterations = s.calculateIterations(newGasLimit)
			s.logger.Infof("Block gas limit changed: %d -> %d, adjusting iterations to %d", 
				s.blockGasLimit, newGasLimit, s.currentIterations)
		}
	}

	// Check if contract is deployed
	if !s.isDeployed {
		return nil, client, wallet, fmt.Errorf("contract not deployed")
	}
	
	// Get current iterations
	s.iterationsMtx.RLock()
	currentIterations := s.currentIterations
	s.iterationsMtx.RUnlock()
	
	// Send spam transaction
	tx, err := s.sendSpamTransaction(ctx, wallet, client, currentIterations)
	if err != nil {
		return nil, client, wallet, err
	}
	
	transactionSubmitted = true
	err = s.walletPool.GetTxPool().SendTransaction(ctx, wallet, tx, &spamoor.SendTransactionOptions{
		Client:      client,
		Rebroadcast: true,
		OnComplete: func(tx *types.Transaction, receipt *types.Receipt, err error) {
			onComplete()
		},
		OnConfirm: func(tx *types.Transaction, receipt *types.Receipt) {
			if receipt != nil && receipt.Status == types.ReceiptStatusSuccessful {
				gasUsed := receipt.GasUsed
				gasLimit := s.blockGasLimit
				if gasLimit == 0 {
					gasLimit = FallbackBlockGasLimit
				}
				utilization := float64(gasUsed) / float64(gasLimit) * 100
				
				s.logChannel <- &LogEntry{
					Wallet:     wallet.GetAddress(),
					Contract:   s.contractAddress,
					Iterations: currentIterations,
					TxHash:     tx.Hash(),
					Success:    true,
					Timestamp:  time.Now(),
				}
				
				// Log gas utilization
				s.logger.WithFields(logrus.Fields{
					"tx":          tx.Hash().Hex(),
					"gasUsed":     gasUsed,
					"gasLimit":    gasLimit,
					"utilization": fmt.Sprintf("%.2f%%", utilization),
					"iterations":  currentIterations,
					"accounts":    currentIterations,
				}).Info("EOA spam transaction confirmed")
			}
		},
	})
	
	if err != nil {
		wallet.ResetPendingNonce(ctx, client)
	}
	
	return tx, client, wallet, err
}

// cleanupPendingTxs removes old pending transactions that likely won't be confirmed
func (s *Scenario) cleanupPendingTxs() {
	s.pendingTxsMtx.Lock()
	defer s.pendingTxsMtx.Unlock()
	
	now := time.Now()
	for hash, info := range s.pendingTxs {
		if now.Sub(info.Timestamp) > 5*time.Minute {
			delete(s.pendingTxs, hash)
		}
	}
}


func (s *Scenario) sendSpamTransaction(ctx context.Context, wallet *spamoor.Wallet, client *spamoor.Client, iterations uint64) (*types.Transaction, error) {
	feeCap, tipCap, err := s.walletPool.GetTxPool().GetSuggestedFees(client, s.options.BaseFee, s.options.TipFee)
	if err != nil {
		return nil, err
	}
	
	// Calculate gas limit for this transaction using dynamic estimate
	gasPerIteration := s.getDynamicGasPerIteration()
	gasLimit := BaseTransactionGas + (iterations * gasPerIteration)
	
	// Safety check: never exceed 95% of block gas limit
	s.gasLimitMtx.RLock()
	blockGasLimit := s.blockGasLimit
	s.gasLimitMtx.RUnlock()
	
	if blockGasLimit == 0 {
		blockGasLimit = FallbackBlockGasLimit
	}
	
	maxSafeGas := uint64(float64(blockGasLimit) * 0.95)
	if gasLimit > maxSafeGas {
		// Recalculate iterations to fit within safe limit
		if maxSafeGas <= BaseTransactionGas {
			iterations = max(s.options.MinIterations, MinViableIterations)
		} else {
			iterations = (maxSafeGas - BaseTransactionGas) / gasPerIteration
			if iterations < MinViableIterations {
				iterations = MinViableIterations
			}
		}
		gasLimit = BaseTransactionGas + (iterations * gasPerIteration)
		s.logger.Debugf("Reduced iterations to %d to fit within gas limit (gasPerIteration: %d)", iterations, gasPerIteration)
	}
	
	// Build transaction using BuildBoundTx pattern
	tx, err := wallet.BuildBoundTx(ctx, &txbuilder.TxMetadata{
		GasFeeCap: uint256.MustFromBig(feeCap),
		GasTipCap: uint256.MustFromBig(tipCap),
		Gas:       gasLimit,
		Value:     uint256.NewInt(iterations), // Send iterations wei (1 wei per EOA)
	}, func(transactOpts *bind.TransactOpts) (*types.Transaction, error) {
		return s.contractInstance.SpamFund(transactOpts, new(big.Int).SetUint64(iterations))
	})
	
	return tx, err
}

func (s *Scenario) asyncLogger(ctx context.Context) {
	defer s.loggerWg.Done()
	
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	
	logs := make([]*LogEntry, 0, 100)
	totalIterations := uint64(0)
	
	for {
		select {
		case <-ctx.Done():
			s.flushLogs(logs, totalIterations)
			return
		case log := <-s.logChannel:
			if log == nil {
				s.flushLogs(logs, totalIterations)
				return
			}
			logs = append(logs, log)
			totalIterations += log.Iterations
		case <-ticker.C:
			s.flushLogs(logs, totalIterations)
			logs = logs[:0]
			totalIterations = 0
		}
	}
}

func (s *Scenario) flushLogs(logs []*LogEntry, totalIterations uint64) {
	// This method is now less important since we log per block
	// Keep it minimal for debugging purposes only
	if len(logs) == 0 {
		return
	}
	
	for _, log := range logs {
		s.logger.WithFields(logrus.Fields{
			"wallet":     log.Wallet.Hex(),
			"contract":   log.Contract.Hex(),
			"iterations": log.Iterations,
			"txHash":     log.TxHash.Hex(),
		}).Debug("EOA spam transaction detail")
	}
}

// runBlockBasedLoop runs a custom loop that sends one transaction per block
func (s *Scenario) runBlockBasedLoop(ctx context.Context) error {
	lastBlockNumber := uint64(0)
	blockTicker := time.NewTicker(BlockPollingInterval)
	defer blockTicker.Stop()
	
	txCount := uint64(0)
	totalGasUsed := uint64(0)
	totalEOAsCreated := uint64(0)
	blockCount := uint64(0)
	
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-blockTicker.C:
			// Check for new block
			client := s.walletPool.GetClient(spamoor.SelectClientRoundRobin, 0, "")
			if client == nil {
				continue
			}
			
			block, err := client.GetEthClient().BlockByNumber(ctx, nil)
			if err != nil {
				s.logger.Warnf("failed to get latest block: %v", err)
				continue
			}
			
			currentBlockNumber := block.NumberU64()
			if currentBlockNumber <= lastBlockNumber {
				continue // No new block yet
			}
			
			// New block detected
			lastBlockNumber = currentBlockNumber
			blockCount++
			
			// Update block gas limit periodically
			if blockCount%10 == 0 {
				newGasLimit := s.getNetworkBlockGasLimit(ctx, client)
				if newGasLimit != s.blockGasLimit {
					s.currentIterations = s.calculateIterations(newGasLimit)
					s.logger.Infof("Block gas limit changed: %d -> %d, adjusting iterations to %d", 
						s.blockGasLimit, newGasLimit, s.currentIterations)
				}
			}
			
			// Select wallet in round-robin fashion
			walletIdx := s.walletIndex % int(s.options.TotalChildWallets)
			s.walletIndex++
			
			wallet := s.walletPool.GetWallet(spamoor.SelectWalletByIndex, walletIdx)
			if wallet == nil {
				s.logger.Warnf("wallet %d not available", walletIdx)
				continue
			}
			
			// Get current iterations
			s.iterationsMtx.RLock()
			currentIterations := s.currentIterations
			s.iterationsMtx.RUnlock()
			
			// Log that we're sending a transaction for this block
			s.logger.Infof("Block %d: sending tx with %d iterations (%.0f%% target)", 
				currentBlockNumber, currentIterations, s.targetGasRatio*100)
			
			// Send spam transaction - MUST send every block
			tx, err := s.sendSpamTransaction(ctx, wallet, client, currentIterations)
			if err != nil {
				// If failed, try with minimum iterations
				s.logger.Warnf("Block %d: failed to send with %d iterations: %v, retrying with minimum", 
					currentBlockNumber, currentIterations, err)
				wallet.ResetPendingNonce(ctx, client)
				
				// Retry with minimum iterations
				minIterations := s.options.MinIterations
				tx, err = s.sendSpamTransaction(ctx, wallet, client, minIterations)
				if err != nil {
					s.logger.Errorf("Block %d: failed to send even with minimum iterations: %v", 
						currentBlockNumber, err)
					continue
				}
				currentIterations = minIterations
			}
			
			txCount++
			
			// Log transaction sent
			s.logger.Infof("Block %d: tx %s sent (%d EOAs)", 
				currentBlockNumber, tx.Hash().Hex()[:16], currentIterations)
			
			// Store pending tx info
			s.pendingTxsMtx.Lock()
			s.pendingTxs[tx.Hash()] = &PendingTx{
				Iterations:  currentIterations,
				BlockNumber: currentBlockNumber,
				Timestamp:   time.Now(),
			}
			s.pendingTxsMtx.Unlock()
			
			// Send transaction without waiting for confirmation
			err = s.walletPool.GetTxPool().SendTransaction(ctx, wallet, tx, &spamoor.SendTransactionOptions{
				Client:      client,
				Rebroadcast: true,
				OnConfirm: func(tx *types.Transaction, receipt *types.Receipt) {
					// Get pending tx info
					s.pendingTxsMtx.RLock()
					pendingInfo, exists := s.pendingTxs[tx.Hash()]
					s.pendingTxsMtx.RUnlock()
					
					if !exists {
						return // Unknown tx
					}
					
					gasLimit := s.blockGasLimit
					if gasLimit == 0 {
						gasLimit = FallbackBlockGasLimit
					}
					
					if receipt != nil {
						gasUsed := receipt.GasUsed
						utilization := float64(gasUsed) / float64(gasLimit) * 100
						txGasLimit := tx.Gas()
						
						// Update gas learning with this transaction's results
						s.updateGasEstimate(pendingInfo.Iterations, txGasLimit, gasUsed, receipt.Status == types.ReceiptStatusSuccessful)
						
						if receipt.Status == types.ReceiptStatusSuccessful {
							// Success
							s.logger.Infof("Block %d: tx %s confirmed (%.1f%% gas used, %d iterations)", 
								pendingInfo.BlockNumber, tx.Hash().Hex()[:16], utilization, pendingInfo.Iterations)
							
							totalGasUsed += gasUsed
							totalEOAsCreated += pendingInfo.Iterations
							
							// Update success counter and reduce failure backoff
							s.consecutiveSuccesses++
							s.consecutiveFailures = 0
							if s.failureBackoffLevel > 0 {
								s.failureBackoffLevel--
							}
							
							// Check if we should increase target ratio
							if s.consecutiveSuccesses >= 5 && s.targetGasRatio < MaxTargetGasRatio {
								s.targetGasRatio += TargetGasRatioIncrement
								if s.targetGasRatio > MaxTargetGasRatio {
									s.targetGasRatio = MaxTargetGasRatio
								}
								s.logger.Infof("Increasing target gas ratio to %.0f%% (learned gas/iter: %.0f)", 
									s.targetGasRatio*100, s.actualGasPerIteration)
								s.consecutiveSuccesses = 0
								// Recalculate iterations for new ratio
								s.currentIterations = s.calculateIterations(gasLimit)
							}
						} else {
							// Failed - check if out of gas
							isOutOfGas := false
							
							// Check if we used all available gas (strong indicator of out-of-gas)
							if gasUsed == txGasLimit {
								isOutOfGas = true
								s.logger.Warnf("Block %d: tx %s OUT OF GAS (used all %d gas, %d iterations)", 
									pendingInfo.BlockNumber, tx.Hash().Hex()[:16], gasUsed, pendingInfo.Iterations)
							} else {
								s.logger.Warnf("Block %d: tx %s failed (gas used: %d of %d, %d iterations)", 
									pendingInfo.BlockNumber, tx.Hash().Hex()[:16], gasUsed, txGasLimit, pendingInfo.Iterations)
							}
							
							// Update failure counter
							s.consecutiveFailures++
							s.consecutiveSuccesses = 0
							
							// More aggressive reduction for out-of-gas
							if isOutOfGas {
								// Increase failure backoff and reduce target ratio
								s.failureBackoffLevel = min(s.failureBackoffLevel+1, 3)
								s.targetGasRatio -= 0.15
								if s.targetGasRatio < 0.25 {
									s.targetGasRatio = 0.25 // Never go below 25%
								}
								s.logger.Warnf("Out of gas detected, reducing target gas ratio to %.0f%% (backoff level: %d)", 
									s.targetGasRatio*100, s.failureBackoffLevel)
								s.consecutiveFailures = 0
								// Recalculate iterations for new ratio
								s.currentIterations = s.calculateIterations(gasLimit)
							} else if s.consecutiveFailures >= 2 {
								// Other failures - reduce less aggressively
								s.failureBackoffLevel = min(s.failureBackoffLevel+1, 2)
								s.targetGasRatio -= 0.08
								if s.targetGasRatio < 0.25 {
									s.targetGasRatio = 0.25
								}
								s.logger.Warnf("Multiple failures, reducing target gas ratio to %.0f%% (backoff level: %d)", 
									s.targetGasRatio*100, s.failureBackoffLevel)
								s.consecutiveFailures = 0
								// Recalculate iterations for new ratio
								s.currentIterations = s.calculateIterations(gasLimit)
							}
						}
						
						// Clean up pending tx
						s.pendingTxsMtx.Lock()
						delete(s.pendingTxs, tx.Hash())
						s.pendingTxsMtx.Unlock()
					}
				},
			})
			
			if err != nil {
				s.logger.Warnf("failed to submit transaction: %v", err)
				wallet.ResetPendingNonce(ctx, client)
			}
			
			// Clean up old pending txs periodically
			if txCount%50 == 0 {
				s.cleanupPendingTxs()
			}
		}
	}
}