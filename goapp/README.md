## Deployment Guide

**Complete guide to deploying and interacting with the `EsportOracle` smart contract.**

## Documentation

https://book.getfoundry.sh/

## Usage

### 1. Launch a Local Anvil Node

```shell
$ anvil
``` 
-   **Private key**: used to sign transactions

### 2. Deploy the Smart Contract with Foundry

```shell
$ forge create "path/to/Contract.sol:ContractName" --private-key <your_private_key> --broadcast
```

This command will return the address of the deployed contract, which you will need for later interactions.

### 3. Call a Contract Function

```shell
$ cast send --private-key <your_private_key> --value <amount_in_ETH> "CONTRACT_ADDRESS" "addFundToStaking()" --gas-limit <your_gas_limit>
```

### 4. Set up .env

```shell
$ mv .env.sample .env
```

.env consists of:
-   **Chain-ID**:  chain ID provided by anvil
-   **Private-Key**: your private key (without 0x)
-   **Contract Adress**:  address of the deployed contract (without 0x)
-   **Client_ETH**: Ethereum client endpoint (e.g., http://localhost:8545 or Infura URL)
-   **Pandascore API token**:  your Pandascore API token
    Register for an account at: https://www.pandascore.co/

### 5. Run the Go Application and send match to the smart contract

```shell
$ go run cmd/myapp/main.go
```

### 6. View the Pending Matches Sent to the Smart Contract

```shell
$ cast call "your_private_key" "getPendingMatches()(bytes32[])"
```

### Help

```shell
$ forge --help
$ anvil --help
$ cast --help
```