# Gas Burner

Send transactions that burn a specific amount of gas units. Note that the estimated gas units may not be 100% accurate.

## Usage

```bash
spamoor gasburnertx [flags]
```

## Configuration

### Base Settings (required)
- `--privkey` - Private key of the sending wallet
- `--rpchost` - RPC endpoint(s) to send transactions to

### Volume Control (either -c or -t required)
- `-c, --count` - Total number of transactions to send
- `-t, --throughput` - Transactions to send per slot
- `--max-pending` - Maximum number of pending transactions

### Transaction Settings
- `--basefee` - Max fee per gas in gwei (default: 20)
- `--tipfee` - Max tip per gas in gwei (default: 2)
- `--gas-units-to-burn` - Number of gas units for each transaction to cost (default: 2000000). Set to 0 to burn all available block gas.
- `--rebroadcast` - Seconds to wait before rebroadcasting (default: 120)

### Wallet Management
- `--max-wallets` - Maximum number of child wallets to use
- `--refill-amount` - ETH amount to fund each child wallet (default: 5)
- `--refill-balance` - Minimum ETH balance before refilling (default: 2)
- `--refill-interval` - Seconds between balance checks (default: 300)

### Client Settings
- `--client-group` - Client group to use for sending transactions

### Debug Options
- `-v, --verbose` - Enable verbose output
- `--log-txs` - Log all submitted transactions
- `--trace` - Enable tracing output

## Example

Send 100 transactions burning 5M gas each:
```bash
spamoor gasburnertx -p "<PRIVKEY>" -h http://rpc-host:8545 -c 100 --gas-units-to-burn 5000000
```

Send 2 transactions per slot burning 1M gas each:
```bash
spamoor gasburnertx -p "<PRIVKEY>" -h http://rpc-host:8545 -t 2 --gas-units-to-burn 1000000
```

Send transactions that burn all available block gas:
```bash
spamoor gasburnertx -p "<PRIVKEY>" -h http://rpc-host:8545 -t 1 --gas-units-to-burn 0
``` 