# ERC20 Max Transfers Scenario

This scenario maximizes the number of ERC20 token transfers per block to unique recipient addresses, creating state bloat through new account storage entries.

## Overview

The scenario uses deployed StateBloatToken contracts from `deployments.json` to send ERC20 transfers. Each transfer sends 1 token to a unique, never-before-used address, creating state growth.

**Updated**: The StateBloatToken contract now includes an optimized `batchTransfer` function that allows sending tokens to multiple recipients in a single transaction, maximizing gas usage per block.

## Features

- **Single Deployer Wallet**: Uses only the deployer wallet which holds all tokens
- **Unique Recipients**: Generates random unique addresses for each transfer
- **Configurable Gas Fees**: Uses configured gas fees (default: 10 gwei base, 2 gwei tip)
- **Round-Robin Contract Usage**: Distributes transfers across multiple deployed contracts
- **Async Logging**: Separates logging from transaction sending for better performance
- **Bloating Summary Tracking**: Saves state growth metrics to `erc20_bloating_summary.json`

## Configuration

### Command Line Flags

- `--basefee`: Max fee per gas in gwei (default: 10)
- `--tipfee`: Max tip per gas in gwei (default: 2)
- `--contract`: Specific contract address to use (default: rotate through all)

## How It Works

1. **Initialization**:
   - Loads deployed contracts and private key from `deployments.json`
   - Sets up the deployer wallet (which holds all tokens)
   - Uses spamoor's RunTransactionScenario for transaction management

2. **Transfer Phase**:
   - Sends multiple transfers per transaction using batch transfer
   - Dynamically calculates optimal batch size based on block gas limit
   - Generates random unique recipient addresses
   - Uses round-robin contract selection
   - Leverages spamoor's built-in nonce handling and transaction rebroadcasting

3. **Logging & Tracking**:
   - Async logging of confirmed transfers
   - Updates contract statistics for unique recipients
   - Periodically saves bloating summary to JSON file

4. **Transaction Management**:
   - Uses BuildBoundTx pattern for proper nonce handling
   - Automatic transaction rebroadcasting on failure
   - One transaction per block (throughput: 1)

## State Growth Impact

Each successful transfer creates:
- New account entry for the recipient (~100 bytes)
- Token balance storage slot for the recipient
- Estimated state growth: 100 bytes per transfer

## Output

State bloat metrics are saved to `erc20_bloating_summary.json` with:
- Per-contract unique recipient counts
- Total recipients across all contracts
- Last block number and timestamp

## Requirements

- Deployed StateBloatToken contracts (via contract_deploy scenario)
- Deployer private key with full token supply
- Sufficient ETH for gas fees

## Example Usage

```bash
# Use default settings
./spamoor scenario --scenario erc20-max-transfers

# Custom gas fees
./spamoor scenario --scenario erc20-max-transfers --basefee 20 --tipfee 10

# Use specific contract only
./spamoor scenario --scenario erc20-max-transfers --contract 0xa513E6E4b8f2a923D98304ec87F64353C4D5C853
```
