package randsstorebloater

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"

	"github.com/ethpandaops/spamoor/scenarios/statebloat/rand_sstore_bloater/contract"
	"github.com/ethpandaops/spamoor/scenario"
	"github.com/ethpandaops/spamoor/spamoor"
	"github.com/ethpandaops/spamoor/txbuilder"
)

//go:embed contract/SSTOREStorageBloater.abi
var contractABIBytes []byte

//go:embed contract/SSTOREStorageBloater.bin
var contractBytecodeHex []byte

// Constants for SSTORE operations
const (
	// Base Ethereum transaction cost
	BaseTxCost = uint64(21000)

	// Function call overhead (measured from actual transactions)
	// Includes: function selector, ABI decoding, contract loading, etc.
	FunctionCallOverhead = uint64(1556)

	// Gas cost per iteration (measured from actual transactions)
	// Includes: SSTORE (0→non-zero), MULMOD, loop overhead, stack operations
	// Based on observed usage with small buffer
	GasPerNewSlotIteration = uint64(23000)

	// Contract deployment and call overhead
	EstimatedDeployGas = uint64(500000) // Deployment gas for our contract

	// Safety margins and multipliers
	GasLimitSafetyMargin = 0.99 // Use 99% of block gas limit (EstimateGas ensures affordability)
	
	// Deployment tracking file
	DeploymentFileName = "deployments_sstore_bloating.json"
)

// BlockInfo stores block information for each storage round
type BlockInfo struct {
	BlockNumber uint64 `json:"block_number"`
	Timestamp   uint64 `json:"timestamp"`
}

// DeploymentData tracks a single contract deployment and its storage rounds
type DeploymentData struct {
	StorageRounds []BlockInfo `json:"storage_rounds"`
}

// DeploymentFile represents the entire deployment tracking file
type DeploymentFile map[string]*DeploymentData // key is contract address

type ScenarioOptions struct {
	BaseFee uint64 `yaml:"base_fee"`
	TipFee  uint64 `yaml:"tip_fee"`
}

type Scenario struct {
	options    ScenarioOptions
	logger     *logrus.Entry
	walletPool *spamoor.WalletPool

	// Contract state
	contractAddress  common.Address
	contractABI      abi.ABI
	contractInstance *contract.Contract // Generated contract binding
	isDeployed       bool
	deployMutex      sync.Mutex

	// Scenario state
	totalSlots      uint64 // Total number of slots created
	cycleCount      uint64 // Number of complete create/update cycles
	roundNumber     uint64 // Current round number for SSTORE bloating
	totalSlotsLock  sync.RWMutex

	// Adaptive gas tracking
	actualGasPerNewSlotIteration uint64          // Dynamically adjusted based on actual usage
	successfulSlotCounts         map[uint64]bool // Track successful slot counts to avoid retries
	gasTrackingLock              sync.RWMutex

	// Cached values
	chainID      *big.Int
	chainIDOnce  sync.Once
	chainIDError error

	// Deployment tracking
	deploymentBuffer     []BlockInfo
	deploymentBufferLock sync.Mutex
	deploymentStopChan   chan struct{}
	deploymentDone       sync.WaitGroup

	// Metrics tracking
	metricsStopChan chan struct{}
	metricsDone     sync.WaitGroup

	// Gas limit caching
	cachedGasLimit     uint64
	cachedGasLimitLock sync.RWMutex
	gasLimitStopChan   chan struct{}
	gasLimitDone       sync.WaitGroup
}

var ScenarioName = "rand_sstore_bloater"
var ScenarioDefaultOptions = ScenarioOptions{
	BaseFee: 10, // 10 gwei default
	TipFee:  2,  // 2 gwei default
}
var ScenarioDescriptor = scenario.Descriptor{
	Name:           ScenarioName,
	Description:    "Maximum state bloat via SSTORE operations using curve25519 prime dispersion",
	DefaultOptions: ScenarioDefaultOptions,
	NewScenario:    newScenario,
}

func newScenario(logger logrus.FieldLogger) scenario.Scenario {
	return &Scenario{
		logger:                       logger.WithField("scenario", ScenarioName),
		actualGasPerNewSlotIteration: GasPerNewSlotIteration, // Start with estimated values
		successfulSlotCounts:         make(map[uint64]bool),
		deploymentBuffer:             make([]BlockInfo, 0, 100),
		deploymentStopChan:           make(chan struct{}),
		metricsStopChan:              make(chan struct{}),
		gasLimitStopChan:             make(chan struct{}),
		cachedGasLimit:               30000000, // Default 30M gas limit
	}
}

func (s *Scenario) Flags(flags *pflag.FlagSet) error {
	flags.Uint64Var(&s.options.BaseFee, "basefee", ScenarioDefaultOptions.BaseFee, "Base fee per gas in gwei")
	flags.Uint64Var(&s.options.TipFee, "tipfee", ScenarioDefaultOptions.TipFee, "Tip fee per gas in gwei")
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

	// Use only one child wallet since we send one tx per block
	s.walletPool.SetWalletCount(1)

	// Parse contract ABI
	parsedABI, err := abi.JSON(strings.NewReader(string(contractABIBytes)))
	if err != nil {
		return fmt.Errorf("failed to parse contract ABI: %w", err)
	}
	s.contractABI = parsedABI

	return nil
}

func (s *Scenario) Config() string {
	yamlBytes, _ := yaml.Marshal(&s.options)
	return string(yamlBytes)
}

// loadDeploymentFile loads the deployment tracking file or creates an empty one
func loadDeploymentFile() (DeploymentFile, error) {
	data, err := os.ReadFile(DeploymentFileName)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist, return empty map
			return make(DeploymentFile), nil
		}
		return nil, fmt.Errorf("failed to read deployment file: %w", err)
	}

	var deployments DeploymentFile
	if err := json.Unmarshal(data, &deployments); err != nil {
		return nil, fmt.Errorf("failed to unmarshal deployment file: %w", err)
	}

	return deployments, nil
}

// saveDeploymentFile saves the deployment tracking file
func saveDeploymentFile(deployments DeploymentFile) error {
	data, err := json.MarshalIndent(deployments, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal deployment file: %w", err)
	}

	if err := os.WriteFile(DeploymentFileName, data, 0644); err != nil {
		return fmt.Errorf("failed to write deployment file: %w", err)
	}

	return nil
}

func (s *Scenario) getChainID(ctx context.Context) (*big.Int, error) {
	s.chainIDOnce.Do(func() {
		client := s.walletPool.GetClient(spamoor.SelectClientByIndex, 0, "")
		if client == nil {
			s.chainIDError = fmt.Errorf("no client available for chain ID")
			return
		}
		s.chainID, s.chainIDError = client.GetChainId(ctx)
	})
	return s.chainID, s.chainIDError
}

func (s *Scenario) deployContract(ctx context.Context) error {
	s.deployMutex.Lock()
	defer s.deployMutex.Unlock()

	if s.isDeployed {
		return nil
	}

	s.logger.Info("Deploying SSTOREStorageBloater contract...")

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

	// Build deployment transaction
	var deployTx *types.Transaction
	var contractAddr common.Address
	var contractInst *contract.Contract
	
	tx, err := wallet.BuildBoundTx(ctx, &txbuilder.TxMetadata{
		GasFeeCap: uint256.MustFromBig(feeCap),
		GasTipCap: uint256.MustFromBig(tipCap),
		Gas:       EstimatedDeployGas,
	}, func(transactOpts *bind.TransactOpts) (*types.Transaction, error) {
		addr, dtx, cInst, err := contract.DeployContract(transactOpts, client.GetEthClient())
		if err != nil {
			return nil, err
		}
		contractAddr = addr
		deployTx = dtx
		contractInst = cInst
		return deployTx, nil
	})
	if err != nil {
		return fmt.Errorf("failed to build deployment transaction: %w", err)
	}

	s.logger.WithField("tx", tx.Hash().Hex()).Info("Contract deployment transaction sent")

	// Wait for deployment using transaction pool
	receipt, err := s.walletPool.GetTxPool().SendAndAwaitTransaction(ctx, wallet, tx, &spamoor.SendTransactionOptions{
		Client:      client,
		Rebroadcast: true,
	})
	if err != nil {
		return fmt.Errorf("failed to send deployment transaction: %w", err)
	}

	if receipt.Status != 1 {
		return fmt.Errorf("contract deployment failed")
	}

	s.contractAddress = contractAddr
	s.contractInstance = contractInst
	s.isDeployed = true

	// Create contract instance for future use
	contractInstance, err := contract.NewContract(s.contractAddress, client.GetEthClient())
	if err != nil {
		return fmt.Errorf("failed to create contract instance: %w", err)
	}
	s.contractInstance = contractInstance

	// Track deployment in JSON file
	deployments, err := loadDeploymentFile()
	if err != nil {
		s.logger.Warnf("failed to load deployment file: %v", err)
		deployments = make(DeploymentFile)
	}

	// Initialize deployment data for this contract
	deployments[contractAddr.Hex()] = &DeploymentData{
		StorageRounds: []BlockInfo{},
	}

	if err := saveDeploymentFile(deployments); err != nil {
		s.logger.Warnf("failed to save deployment file: %v", err)
	}

	s.logger.WithField("address", contractAddr.Hex()).Info("SSTOREStorageBloater contract deployed successfully")

	return nil
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

	// Start background workers
	s.startDeploymentTracker(ctx)
	s.startMetricsLogger(ctx)
	s.startGasLimitUpdater(ctx)

	// Cleanup on exit
	defer func() {
		// Stop background workers
		close(s.deploymentStopChan)
		close(s.metricsStopChan)
		close(s.gasLimitStopChan)
		s.deploymentDone.Wait()
		s.metricsDone.Wait()
		s.gasLimitDone.Wait()

		// Flush any remaining deployment data
		s.flushDeploymentBuffer()
	}()

	// Custom loop for single transaction per block
	return s.runSingleTxPerBlock(ctx)
}

// runSingleTxPerBlock implements the custom loop for one transaction per block
func (s *Scenario) runSingleTxPerBlock(ctx context.Context) error {
	// Get single wallet and client
	wallet := s.walletPool.GetWallet(spamoor.SelectWalletByIndex, 0)
	if wallet == nil {
		return fmt.Errorf("no wallet available")
	}

	var txCount uint64
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// Get client (may change due to failover)
		client := s.walletPool.GetClient(spamoor.SelectClientByIndex, 0, "")
		if client == nil {
			s.logger.Warn("no client available, waiting...")
			time.Sleep(5 * time.Second)
			continue
		}

		// Use cached gas limit
		s.cachedGasLimitLock.RLock()
		blockGasLimit := s.cachedGasLimit
		s.cachedGasLimitLock.RUnlock()

		targetGas := uint64(float64(blockGasLimit) * GasLimitSafetyMargin)

		// Send transaction and wait for confirmation
		tx, receipt, err := s.sendAndWaitForTransaction(ctx, client, wallet, targetGas, blockGasLimit)
		if err != nil {
			if strings.Contains(err.Error(), "context canceled") {
				return nil
			}
			s.logger.Warnf("failed to send transaction: %v", err)
			// Wait a bit before retry
			time.Sleep(2 * time.Second)
			continue
		}

		txCount++
		s.logger.WithFields(logrus.Fields{
			"tx":       tx.Hash().Hex(),
			"block":    receipt.BlockNumber,
			"gas_used": receipt.GasUsed,
			"total_tx": txCount,
		}).Info("Transaction confirmed")
	}
}

// sendAndWaitForTransaction sends a transaction and waits for it to be mined
func (s *Scenario) sendAndWaitForTransaction(ctx context.Context, client *spamoor.Client, wallet *spamoor.Wallet, targetGas uint64, blockGasLimit uint64) (*types.Transaction, *types.Receipt, error) {
	// Build and send the transaction
	tx, err := s.buildCreateSlotsTransaction(ctx, client, wallet, targetGas, blockGasLimit)
	if err != nil {
		return nil, nil, err
	}

	// Send transaction and wait for confirmation
	receipt, err := s.walletPool.GetTxPool().SendAndAwaitTransaction(ctx, wallet, tx, &spamoor.SendTransactionOptions{
		Client:      client,
		Rebroadcast: true,
	})
	if err != nil {
		// Reset nonce on error
		wallet.ResetPendingNonce(ctx, client)
		return nil, nil, err
	}

	if receipt.Status != 1 {
		return tx, receipt, fmt.Errorf("transaction failed")
	}

	// Update metrics from receipt
	// Extract slots created from tx data
	slotsCreated := s.extractSlotsFromTx(tx)
	s.updateMetricsFromReceipt(receipt, slotsCreated, blockGasLimit)

	return tx, receipt, nil
}

// extractSlotsFromTx extracts the number of slots from transaction data
func (s *Scenario) extractSlotsFromTx(tx *types.Transaction) uint64 {
	if len(tx.Data()) < 36 { // 4 bytes selector + 32 bytes uint256
		return 0
	}
	
	// Skip function selector (first 4 bytes) and read the uint256 parameter
	slotsBig := new(big.Int).SetBytes(tx.Data()[4:36])
	return slotsBig.Uint64()
}

// buildCreateSlotsTransaction builds a single createSlots transaction
func (s *Scenario) buildCreateSlotsTransaction(ctx context.Context, client *spamoor.Client, wallet *spamoor.Wallet, targetGas uint64, blockGasLimit uint64) (*types.Transaction, error) {
	// Get suggested fees first
	feeCap, tipCap, err := s.walletPool.GetTxPool().GetSuggestedFees(client, s.options.BaseFee, s.options.TipFee)
	if err != nil {
		return nil, fmt.Errorf("failed to get suggested fees: %w", err)
	}

	// Calculate initial slots to create based on gas limits
	s.gasTrackingLock.RLock()
	gasPerSlot := s.actualGasPerNewSlotIteration
	s.gasTrackingLock.RUnlock()

	availableGas := targetGas - BaseTxCost - FunctionCallOverhead
	slotsToCreate := availableGas / gasPerSlot

	if slotsToCreate == 0 {
		return nil, fmt.Errorf("not enough gas to create any slots")
	}

	// Cap slots to reasonable maximum
	maxSlotsPerTx := uint64(5200)
	if slotsToCreate > maxSlotsPerTx {
		slotsToCreate = maxSlotsPerTx
	}

	// Encode the transaction data for gas estimation
	data, err := s.contractABI.Pack("createSlots", big.NewInt(int64(slotsToCreate)))
	if err != nil {
		return nil, fmt.Errorf("failed to pack transaction data: %w", err)
	}

	// Estimate gas for the actual transaction
	msg := ethereum.CallMsg{
		From:     wallet.GetAddress(),
		To:       &s.contractAddress,
		GasPrice: feeCap,
		Value:    big.NewInt(0),
		Data:     data,
	}

	estimatedGas, err := client.GetEthClient().EstimateGas(ctx, msg)
	if err != nil {
		return nil, fmt.Errorf("failed to estimate gas: %w", err)
	}

	// Add 10% buffer to the estimate
	adjustedGas := uint64(float64(estimatedGas) * 1.1)
	
	// Check wallet balance
	walletBalance := wallet.GetBalance()
	txCost := new(big.Int).Mul(feeCap, big.NewInt(int64(adjustedGas)))
	
	if walletBalance.Cmp(txCost) < 0 {
		// Try to reduce slots to fit wallet balance
		maxAffordableGas := new(big.Int).Div(walletBalance, feeCap)
		if !maxAffordableGas.IsUint64() {
			return nil, fmt.Errorf("wallet balance too low")
		}
		
		// Calculate slots we can afford (with safety margin)
		affordableGas := uint64(float64(maxAffordableGas.Uint64()) * 0.9)
		if affordableGas < estimatedGas/uint64(slotsToCreate)*100 { // Need at least 100 slots
			return nil, fmt.Errorf("insufficient balance for minimum transaction")
		}
		
		// Adjust slots proportionally
		slotsToCreate = (affordableGas - BaseTxCost - FunctionCallOverhead) / gasPerSlot
		adjustedGas = affordableGas
		
		s.logger.Warnf("Reduced slots from original to %d due to wallet balance", slotsToCreate)
	}

	s.logger.Debugf("Gas estimation: estimated=%d, adjusted=%d, slots=%d",
		estimatedGas, adjustedGas, slotsToCreate)

	// Build transaction using BuildBoundTx
	tx, err := wallet.BuildBoundTx(ctx, &txbuilder.TxMetadata{
		To:        &s.contractAddress,
		GasFeeCap: uint256.MustFromBig(feeCap),
		GasTipCap: uint256.MustFromBig(tipCap),
		Gas:       adjustedGas,
	}, func(transactOpts *bind.TransactOpts) (*types.Transaction, error) {
		return s.contractInstance.CreateSlots(transactOpts, big.NewInt(int64(slotsToCreate)))
	})
	if err != nil {
		return nil, fmt.Errorf("failed to build transaction: %w", err)
	}

	return tx, nil
}



// updateMetricsFromReceipt updates scenario metrics from a confirmed transaction
func (s *Scenario) updateMetricsFromReceipt(receipt *types.Receipt, slotsCreated uint64, blockGasLimit uint64) {
	if receipt.Status != 1 {
		s.logger.Warnf("transaction failed with status %d", receipt.Status)
		return
	}

	// Update gas tracking
	s.gasTrackingLock.Lock()
	totalOverhead := BaseTxCost + FunctionCallOverhead
	actualGasPerSlotIteration := (receipt.GasUsed - totalOverhead) / slotsCreated

	// Update gas estimate using exponential moving average
	newEstimate := uint64(float64(s.actualGasPerNewSlotIteration)*0.7 + float64(actualGasPerSlotIteration)*0.3)
	
	// Ensure minimum safe estimate
	minSafeEstimate := uint64(float64(actualGasPerSlotIteration) * 1.05)
	if newEstimate < minSafeEstimate {
		newEstimate = minSafeEstimate
	}
	
	s.actualGasPerNewSlotIteration = newEstimate
	s.successfulSlotCounts[slotsCreated] = true
	s.gasTrackingLock.Unlock()

	// Update total slots
	s.totalSlotsLock.Lock()
	s.totalSlots += slotsCreated
	totalSlots := s.totalSlots
	s.totalSlotsLock.Unlock()

	// Add block info to deployment buffer
	s.deploymentBufferLock.Lock()
	s.deploymentBuffer = append(s.deploymentBuffer, BlockInfo{
		BlockNumber: receipt.BlockNumber.Uint64(),
		Timestamp:   uint64(time.Now().Unix()),
	})
	s.deploymentBufferLock.Unlock()

	// Log metrics
	mbWrittenThisTx := float64(slotsCreated*64) / (1024 * 1024)
	blockUtilization := float64(receipt.GasUsed) / float64(blockGasLimit) * 100

	s.logger.WithFields(logrus.Fields{
		"block_number":      receipt.BlockNumber,
		"gas_used":          receipt.GasUsed,
		"slots_created":     slotsCreated,
		"gas_per_slot":      actualGasPerSlotIteration,
		"total_slots":       totalSlots,
		"mb_written":        mbWrittenThisTx,
		"block_utilization": fmt.Sprintf("%.2f%%", blockUtilization),
	}).Info("SSTORE bloating round summary")
}

// startDeploymentTracker starts a background goroutine to periodically save deployment data
func (s *Scenario) startDeploymentTracker(ctx context.Context) {
	s.deploymentDone.Add(1)
	go func() {
		defer s.deploymentDone.Done()
		
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		
		for {
			select {
			case <-ctx.Done():
				return
			case <-s.deploymentStopChan:
				return
			case <-ticker.C:
				s.flushDeploymentBuffer()
			}
		}
	}()
}

// flushDeploymentBuffer saves buffered deployment data to file
func (s *Scenario) flushDeploymentBuffer() {
	s.deploymentBufferLock.Lock()
	if len(s.deploymentBuffer) == 0 {
		s.deploymentBufferLock.Unlock()
		return
	}
	
	// Copy and clear buffer
	bufferCopy := make([]BlockInfo, len(s.deploymentBuffer))
	copy(bufferCopy, s.deploymentBuffer)
	s.deploymentBuffer = s.deploymentBuffer[:0]
	s.deploymentBufferLock.Unlock()

	// Load existing deployments
	deployments, err := loadDeploymentFile()
	if err != nil {
		s.logger.Warnf("failed to load deployment file: %v", err)
		deployments = make(DeploymentFile)
	}

	// Update deployment data
	contractAddr := s.contractAddress.Hex()
	if deploymentData, exists := deployments[contractAddr]; exists {
		deploymentData.StorageRounds = append(deploymentData.StorageRounds, bufferCopy...)
	} else {
		deployments[contractAddr] = &DeploymentData{
			StorageRounds: bufferCopy,
		}
	}

	// Save to file
	if err := saveDeploymentFile(deployments); err != nil {
		s.logger.Warnf("failed to save deployment file: %v", err)
	}
}

// startMetricsLogger starts a background goroutine to log overall metrics
func (s *Scenario) startMetricsLogger(ctx context.Context) {
	s.metricsDone.Add(1)
	go func() {
		defer s.metricsDone.Done()
		
		ticker := time.NewTicker(60 * time.Second)
		defer ticker.Stop()
		
		for {
			select {
			case <-ctx.Done():
				return
			case <-s.metricsStopChan:
				return
			case <-ticker.C:
				s.totalSlotsLock.RLock()
				totalSlots := s.totalSlots
				s.totalSlotsLock.RUnlock()
				
				s.gasTrackingLock.RLock()
				gasPerSlot := s.actualGasPerNewSlotIteration
				s.gasTrackingLock.RUnlock()
				
				totalMB := float64(totalSlots*64) / (1024 * 1024)
				s.logger.WithFields(logrus.Fields{
					"total_slots":          totalSlots,
					"total_mb_written":     totalMB,
					"current_gas_per_slot": gasPerSlot,
				}).Info("SSTORE bloater overall metrics")
			}
		}
	}()
}

// startGasLimitUpdater starts a background goroutine to periodically update cached gas limit
func (s *Scenario) startGasLimitUpdater(ctx context.Context) {
	s.gasLimitDone.Add(1)
	go func() {
		defer s.gasLimitDone.Done()
		
		// Update immediately on start
		s.updateGasLimit(ctx)
		
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		
		for {
			select {
			case <-ctx.Done():
				return
			case <-s.gasLimitStopChan:
				return
			case <-ticker.C:
				s.updateGasLimit(ctx)
			}
		}
	}()
}

// updateGasLimit fetches the current block gas limit and updates the cache
func (s *Scenario) updateGasLimit(ctx context.Context) {
	client := s.walletPool.GetClient(spamoor.SelectClientByIndex, 0, "")
	if client == nil {
		s.logger.Warn("no client available for gas limit update")
		return
	}

	latestBlock, err := client.GetEthClient().BlockByNumber(ctx, nil)
	if err != nil {
		s.logger.Warnf("failed to get latest block for gas limit: %v", err)
		return
	}

	newGasLimit := latestBlock.GasLimit()
	
	s.cachedGasLimitLock.Lock()
	if s.cachedGasLimit != newGasLimit {
		s.logger.Infof("Updated cached gas limit from %d to %d", s.cachedGasLimit, newGasLimit)
		s.cachedGasLimit = newGasLimit
	}
	s.cachedGasLimitLock.Unlock()
}
