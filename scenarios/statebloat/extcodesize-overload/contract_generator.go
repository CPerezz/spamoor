package extcodesizeoverload

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
)

type ContractGenerator struct {
	logger               *logrus.Entry
	targetAddresses      []common.Address
	contractsDir         string
	nextAddresses        map[int]common.Address
	addressesPerContract int
	factoryAddress       common.Address
	contractBytecodes    map[int][]byte
}

func NewContractGenerator(logger *logrus.Entry, targetAddresses []common.Address, addressesPerContract int) *ContractGenerator {
	return &ContractGenerator{
		logger:               logger,
		targetAddresses:      targetAddresses,
		contractsDir:         "scenarios/statebloat/extcodesize-overload/contracts",
		nextAddresses:        make(map[int]common.Address),
		addressesPerContract: addressesPerContract,
		contractBytecodes:    make(map[int][]byte),
	}
}

func (g *ContractGenerator) SetFactoryAddress(address common.Address) {
	g.factoryAddress = address
}

// CalculateCREATE2Address calculates the deterministic address for a contract
func (g *ContractGenerator) CalculateCREATE2Address(bytecode []byte, salt uint64) common.Address {
	// CREATE2 address = keccak256(0xff ++ factoryAddress ++ salt ++ keccak256(bytecode))[12:]
	data := []byte{0xff}
	data = append(data, g.factoryAddress.Bytes()...)
	
	// Convert salt to bytes32 (matching the factory contract's expectation)
	saltBytes := make([]byte, 32)
	big.NewInt(int64(salt)).FillBytes(saltBytes[24:]) // Put the value in the last 8 bytes
	data = append(data, saltBytes...)
	
	// Hash of bytecode
	bytecodeHash := crypto.Keccak256(bytecode)
	data = append(data, bytecodeHash...)
	
	// Final hash and extract address
	hash := crypto.Keccak256(data)
	return common.BytesToAddress(hash[12:])
}

func (g *ContractGenerator) GenerateContracts(numContracts int) error {
	// Clean contracts directory
	files, _ := filepath.Glob(filepath.Join(g.contractsDir, "*.sol"))
	for _, f := range files {
		os.Remove(f)
	}
	files, _ = filepath.Glob(filepath.Join(g.contractsDir, "*.bin"))
	for _, f := range files {
		os.Remove(f)
	}
	files, _ = filepath.Glob(filepath.Join(g.contractsDir, "*.abi"))
	for _, f := range files {
		os.Remove(f)
	}

	g.logger.Info("Generating contracts with storage-based NEXT addresses")
	
	// Generate all contracts (no NEXT addresses needed in bytecode)
	for i := 0; i < numContracts; i++ {
		if err := g.generateContract(i, numContracts); err != nil {
			return fmt.Errorf("failed to generate contract %d: %w", i, err)
		}
	}
	
	// Compile all contracts
	if err := g.compileContracts(); err != nil {
		return fmt.Errorf("compilation failed: %w", err)
	}
	
	// Calculate all CREATE2 addresses and store bytecode
	for i := 0; i < numContracts; i++ {
		bytecode, err := g.readBytecodeFromDisk(i)
		if err != nil {
			return fmt.Errorf("failed to read bytecode for contract %d: %w", i, err)
		}
		
		addr := g.CalculateCREATE2Address(bytecode, uint64(i))
		g.nextAddresses[i] = addr
		g.contractBytecodes[i] = bytecode
		g.logger.Infof("Contract %d will deploy at: %s", i, addr.Hex())
	}
	
	// Log the expected chain
	g.logger.Info("Expected chain after setNext calls:")
	for i := 0; i < numContracts; i++ {
		if i < numContracts-1 {
			g.logger.Infof("  Contract %d: %s -> %s", i, g.nextAddresses[i].Hex(), g.nextAddresses[i+1].Hex())
		} else {
			g.logger.Infof("  Contract %d: %s -> END", i, g.nextAddresses[i].Hex())
		}
	}
	
	return nil
}

func (g *ContractGenerator) generateContract(index int, totalContracts int) error {

	// Calculate address range for this contract using dynamic addresses per contract
	startIdx := index * g.addressesPerContract
	endIdx := startIdx + g.addressesPerContract
	
	// Ensure we don't go out of bounds
	if startIdx >= len(g.targetAddresses) {
		// This should not happen with proper address calculation
		return fmt.Errorf("contract %d has no addresses available (startIdx=%d >= total=%d)", 
			index, startIdx, len(g.targetAddresses))
	}
	
	if endIdx > len(g.targetAddresses) {
		g.logger.Warnf("Contract %d: endIdx %d exceeds address array length %d, truncating", 
			index, endIdx, len(g.targetAddresses))
		endIdx = len(g.targetAddresses)
	}

	contractAddresses := g.targetAddresses[startIdx:endIdx]
	
	// Log the actual number of addresses for this contract
	actualAddresses := len(contractAddresses)
	if actualAddresses < g.addressesPerContract {
		g.logger.Warnf("Contract %d: only got %d addresses, expected %d", 
			index, actualAddresses, g.addressesPerContract)
	}
	g.logger.Debugf("Contract %d: using addresses %d-%d (%d addresses)", 
		index, startIdx, endIdx-1, actualAddresses)

	// Generate contract source (no NEXT address needed)
	source := g.generateContractSource(index, "", contractAddresses)

	// Write to file
	filename := filepath.Join(g.contractsDir, fmt.Sprintf("ExtcodesizeCaller%d.sol", index))
	if err := os.WriteFile(filename, []byte(source), 0644); err != nil {
		return fmt.Errorf("failed to write contract source: %w", err)
	}

	return nil
}

func (g *ContractGenerator) generateContractSource(index int, nextAddr string, addresses []common.Address) string {
	var buf bytes.Buffer

	// Write pragma and contract header
	buf.WriteString("// SPDX-License-Identifier: MIT\n")
	buf.WriteString("pragma solidity ^0.8.0;\n\n")
	buf.WriteString(fmt.Sprintf("contract ExtcodesizeCaller%d {\n", index))
	buf.WriteString("    address public next;\n")
	buf.WriteString(fmt.Sprintf("    uint256 constant INDEX = %d;\n", index))
	buf.WriteString("\n")
	buf.WriteString("    function setNext(address _next) external {\n")
	buf.WriteString("        require(next == address(0), \"Next already set\");\n")
	buf.WriteString("        next = _next;\n")
	buf.WriteString("    }\n")
	buf.WriteString("\n")
	buf.WriteString("    function execute() external {\n")
	buf.WriteString("        uint256 size;\n")
	buf.WriteString("        assembly {\n")

	// Generate EXTCODESIZE calls that accumulate the sizes
	for i, addr := range addresses {
		if i == 0 {
			buf.WriteString(fmt.Sprintf("            size := extcodesize(%s)\n", addr.Hex()))
		} else {
			buf.WriteString(fmt.Sprintf("            size := add(size, extcodesize(%s))\n", addr.Hex()))
		}
	}

	buf.WriteString("        }\n\n")

	// Add call to next contract (using storage variable)
	buf.WriteString("        if (next != address(0)) {\n")
	buf.WriteString("            (bool success, bytes memory returnData) = next.call(abi.encodeWithSignature(\"execute()\"));\n")
	buf.WriteString("            if (!success) {\n")
	buf.WriteString("                // Try to decode the error message if it's another contract's failure\n")
	buf.WriteString("                if (returnData.length > 0) {\n")
	buf.WriteString("                    // Propagate the original error\n")
	buf.WriteString("                    assembly { revert(add(returnData, 32), mload(returnData)) }\n")
	buf.WriteString("                } else {\n")
	buf.WriteString(fmt.Sprintf("                    revert(\"Chain failed at contract %d calling next\");\n", index))
	buf.WriteString("                }\n")
	buf.WriteString("            }\n")
	buf.WriteString("        }\n")

	buf.WriteString("    }\n")
	buf.WriteString("}\n")

	return buf.String()
}

func (g *ContractGenerator) compileContracts() error {
	g.logger.Info("compiling contracts with solc")

	// Check if solc is available
	if _, err := exec.LookPath("solc"); err != nil {
		return fmt.Errorf("solc not found in PATH: %w", err)
	}

	// Compile each contract
	files, err := filepath.Glob(filepath.Join(g.contractsDir, "*.sol"))
	if err != nil {
		return fmt.Errorf("failed to list contract files: %w", err)
	}

	for _, file := range files {
		cmd := exec.Command("solc", 
			"--optimize",
			"--optimize-runs", "200",
			"--bin",
			"--abi",
			"--overwrite",
			"-o", g.contractsDir,
			file)

		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to compile %s: %w\nOutput: %s", file, err, output)
		}
	}

	g.logger.Info("contracts compiled successfully")
	return nil
}

func (g *ContractGenerator) compileSingleContract(index int) error {
	file := filepath.Join(g.contractsDir, fmt.Sprintf("ExtcodesizeCaller%d.sol", index))
	cmd := exec.Command("solc", 
		"--optimize",
		"--optimize-runs", "200",
		"--bin",
		"--abi",
		"--overwrite",
		"-o", g.contractsDir,
		file)
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to compile %s: %w\nOutput: %s", file, err, output)
	}
	return nil
}

func (g *ContractGenerator) GetContractBytecode(index int) ([]byte, error) {
	// First check if we have cached bytecode from address calculation
	if bytecode, ok := g.contractBytecodes[index]; ok && len(bytecode) > 0 {
		g.logger.Infof("contract %d bytecode size: %d bytes (addresses: %d) [cached]", index, len(bytecode), g.addressesPerContract)
		return bytecode, nil
	}

	// Fallback to reading from disk if not cached
	// Read bytecode
	binFile := filepath.Join(g.contractsDir, fmt.Sprintf("ExtcodesizeCaller%d.bin", index))
	binData, err := os.ReadFile(binFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read bytecode: %w", err)
	}

	// Convert hex string to bytes
	binStr := strings.TrimSpace(string(binData))
	bytecode, err := hex.DecodeString(binStr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode bytecode: %w", err)
	}

	g.logger.Infof("contract %d bytecode size: %d bytes (addresses: %d) [from disk]", index, len(bytecode), g.addressesPerContract)
	return bytecode, nil
}

// readBytecodeFromDisk reads bytecode directly from disk without checking cache
func (g *ContractGenerator) readBytecodeFromDisk(index int) ([]byte, error) {
	binFile := filepath.Join(g.contractsDir, fmt.Sprintf("ExtcodesizeCaller%d.bin", index))
	binData, err := os.ReadFile(binFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read bytecode: %w", err)
	}

	// Convert hex string to bytes
	binStr := strings.TrimSpace(string(binData))
	bytecode, err := hex.DecodeString(binStr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode bytecode: %w", err)
	}

	return bytecode, nil
}

// GetFactoryBytecode is no longer needed as we use the existing CREATE2Factory