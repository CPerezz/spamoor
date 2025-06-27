# EOA Spam Scenario

This scenario deploys contracts that fund new EOA addresses by sending 1 wei to addresses generated through keccak256 hashing. It's designed to maximize gas usage per block by sending a single large transaction that consumes ~95% of the block's gas limit.

## How it works

1. **Initialization**: Deploys EOASpammer contracts for all wallets upfront
2. **Block-based execution**: Sends exactly ONE transaction per block
3. **Dynamic gas targeting**: Automatically adjusts iterations to consume ~95% of block gas limit
4. **Wallet rotation**: Uses a different wallet for each transaction
5. **Adaptive optimization**: Monitors actual gas usage and adjusts future iterations

## Contract Design

The EOASpammer contract is optimized for minimal gas usage:
- No event emissions
- No require statements  
- Uses assembly for transfers
- Unchecked arithmetic for loop counter
- Returns excess ETH to sender

## Configuration

- `--total-child-wallets`: Number of wallets to use (default: 50)
- `--min-iterations`: Minimum iterations per transaction (default: 100)
- `--max-iterations`: Maximum iterations per transaction (default: 100000)
- `--max-wallets`: Maximum concurrent wallets (default: 100)
- `--target-gas-ratio`: Target gas usage ratio of block limit (default: 0.95)
- `--gas-per-iteration`: Estimated gas cost per iteration (default: 5200)
- `--basefee`: Max fee per gas in gwei (default: 50)
- `--tipfee`: Max tip per gas in gwei (default: 2)

## Performance

- **Gas efficiency**: Each iteration costs ~5200 gas
- **Block utilization**: Targets 95% of block gas limit
- **Account creation**: Creates thousands of new EOAs per block
- **Adaptive scaling**: Automatically adjusts to network conditions

## Example Usage

```bash
# Run with 10 wallets targeting 95% gas utilization
./spamoor eoa-spam --privkey <key> --seed <seed> --rpchost <rpc> \
  --total-child-wallets 10 --target-gas-ratio 0.95

# Run with custom gas parameters
./spamoor eoa-spam --privkey <key> --seed <seed> --rpchost <rpc> \
  --min-iterations 1000 --max-iterations 50000 --gas-per-iteration 5000
```

## Logging

The scenario provides detailed metrics:
- **Per-transaction logs**: Shows gas used, utilization %, and accounts created
- **Batch summaries**: Average gas utilization and performance metrics
- **Async logging**: Non-blocking log collection for optimal performance
- **Real-time adjustments**: Logs when iterations are adjusted based on gas usage