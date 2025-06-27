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

const ScenarioName = "eoa-spam"

const (
	// Gas cost per iteration in the spam contract
	GasPerIteration = 5200
	// Base gas for transaction overhead
	BaseTransactionGas = 30000
	// Default target gas ratio (95% of block limit)
	DefaultTargetGasRatio = 0.95
	// Fallback block gas limit if network query fails
	FallbackBlockGasLimit = 30000000
	// Block polling interval
	BlockPollingInterval = 500 * time.Millisecond
	// Timeout for waiting for new block
	BlockMiningTimeout = 30 * time.Second
)

var ScenarioDefaultOptions = ScenarioOptions{
	TotalChildWallets: 50,
	MinIterations:     100,
	MaxIterations:     100000,
	MaxWallets:        100,
	BaseFee:           50,
	TipFee:            2,
	TargetGasRatio:    DefaultTargetGasRatio,
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
}

type LogEntry struct {
	Wallet     common.Address
	Contract   common.Address
	Iterations uint64
	TxHash     common.Hash
	Success    bool
	Timestamp  time.Time
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

	// Initialize current iterations based on initial estimate
	s.currentIterations = s.options.MinIterations

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

// calculateIterations calculates the optimal number of iterations based on block gas limit
func (s *Scenario) calculateIterations(blockGasLimit uint64) uint64 {
	targetGas := uint64(float64(blockGasLimit) * s.options.TargetGasRatio)
	
	// Calculate iterations: (targetGas - baseGas) / gasPerIteration
	if targetGas <= BaseTransactionGas {
		return s.options.MinIterations
	}
	
	iterations := (targetGas - BaseTransactionGas) / s.options.GasPerIteration
	
	// Apply bounds
	if iterations < s.options.MinIterations {
		iterations = s.options.MinIterations
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
	
	s.logger.Infof("Initial block gas limit: %d, starting with %d iterations per tx", blockGasLimit, s.currentIterations)

	// Run with Throughput: 1 to ensure only one transaction per block
	return scenario.RunTransactionScenario(ctx, scenario.TransactionScenarioOptions{
		TotalCount: 0, // Run indefinitely
		Throughput: 1,  // Only one transaction per block
		MaxPending: 10, // Keep low to ensure sequential processing
		Timeout:    0,  // No timeout
		WalletPool: s.walletPool,
		Logger:     s.logger,
		ProcessNextTxFn: func(ctx context.Context, txIdx uint64, onComplete func()) (func(), error) {
			tx, client, wallet, err := s.processTransaction(ctx, txIdx, onComplete)
			logger := s.logger
			if client != nil {
				logger = logger.WithField("rpc", client.GetName())
			}
			if tx != nil {
				logger = logger.WithField("nonce", tx.Nonce())
			}
			if wallet != nil {
				logger = logger.WithField("wallet", s.walletPool.GetWalletName(wallet.GetAddress()))
			}

			return func() {
				if err != nil {
					logger.Warnf("could not send transaction: %v", err)
				} else {
					logger.Infof("sent tx #%6d: %v (iterations: %d)", txIdx+1, tx.Hash().String(), s.currentIterations)
				}
			}, err
		},
	})
}

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
				
				// Adjust iterations based on actual gas usage
				s.adjustIterations(gasUsed, gasLimit)
			}
		},
	})
	
	if err != nil {
		wallet.ResetPendingNonce(ctx, client)
	}
	
	return tx, client, wallet, err
}

// adjustIterations dynamically adjusts iterations based on actual gas usage
func (s *Scenario) adjustIterations(gasUsed uint64, gasLimit uint64) {
	s.iterationsMtx.Lock()
	defer s.iterationsMtx.Unlock()
	
	targetGas := uint64(float64(gasLimit) * s.options.TargetGasRatio)
	currentUtilization := float64(gasUsed) / float64(gasLimit)
	
	// If we're using less than 90% of target, increase iterations
	if currentUtilization < s.options.TargetGasRatio*0.9 {
		// Calculate how many more iterations we could fit
		unusedGas := targetGas - gasUsed
		additionalIterations := unusedGas / s.options.GasPerIteration
		
		newIterations := s.currentIterations + additionalIterations/2 // Conservative increase
		if newIterations > s.options.MaxIterations {
			newIterations = s.options.MaxIterations
		}
		
		if newIterations != s.currentIterations {
			s.logger.Debugf("Increasing iterations: %d -> %d (utilization: %.2f%%)", 
				s.currentIterations, newIterations, currentUtilization*100)
			s.currentIterations = newIterations
		}
	} else if currentUtilization > s.options.TargetGasRatio*1.05 {
		// If we're over target, decrease iterations
		excessGas := gasUsed - targetGas
		reduceIterations := excessGas / s.options.GasPerIteration
		
		newIterations := s.currentIterations - reduceIterations
		if newIterations < s.options.MinIterations {
			newIterations = s.options.MinIterations
		}
		
		if newIterations != s.currentIterations {
			s.logger.Debugf("Decreasing iterations: %d -> %d (utilization: %.2f%%)", 
				s.currentIterations, newIterations, currentUtilization*100)
			s.currentIterations = newIterations
		}
	}
	
	s.lastGasUsed = gasUsed
}


func (s *Scenario) sendSpamTransaction(ctx context.Context, wallet *spamoor.Wallet, client *spamoor.Client, iterations uint64) (*types.Transaction, error) {
	feeCap, tipCap, err := s.walletPool.GetTxPool().GetSuggestedFees(client, s.options.BaseFee, s.options.TipFee)
	if err != nil {
		return nil, err
	}
	
	// Calculate gas limit for this transaction
	gasLimit := BaseTransactionGas + (iterations * s.options.GasPerIteration)
	
	// Build transaction using BuildBoundTx pattern
	tx, err := wallet.BuildBoundTx(ctx, &txbuilder.TxMetadata{
		GasFeeCap: uint256.MustFromBig(feeCap),
		GasTipCap: uint256.MustFromBig(tipCap),
		Gas:       gasLimit,
		Value:     uint256.NewInt(iterations), // Send iterations wei
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
	if len(logs) == 0 {
		return
	}
	
	// Calculate average gas utilization
	s.gasLimitMtx.RLock()
	gasLimit := s.blockGasLimit
	s.gasLimitMtx.RUnlock()
	
	if gasLimit == 0 {
		gasLimit = FallbackBlockGasLimit
	}
	
	avgGasPerTx := s.lastGasUsed
	avgUtilization := float64(avgGasPerTx) / float64(gasLimit) * 100
	
	s.logger.WithFields(logrus.Fields{
		"transactions":    len(logs),
		"totalAccounts":   totalIterations,
		"avgIterations":   totalIterations / uint64(len(logs)),
		"avgGasPerTx":     avgGasPerTx,
		"avgUtilization":  fmt.Sprintf("%.2f%%", avgUtilization),
		"targetRatio":     fmt.Sprintf("%.0f%%", s.options.TargetGasRatio*100),
	}).Info("EOA spam batch completed")
	
	for _, log := range logs {
		s.logger.WithFields(logrus.Fields{
			"wallet":     log.Wallet.Hex(),
			"contract":   log.Contract.Hex(),
			"iterations": log.Iterations,
			"txHash":     log.TxHash.Hex(),
		}).Debug("EOA spam transaction")
	}
}