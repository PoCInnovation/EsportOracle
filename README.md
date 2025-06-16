# Esport Oracle

Esport Oracle is a decentralized oracle developed by PoC Innovation that bridges Counter-Strike 2 esports data with blockchain smart contracts. It securely and verifiably transfers match results, player statistics, tournament details and other competitive data onto the blockchain, enabling new possibilities for decentralized betting and other CS2-related applications while ensuring full transparency and accessibility for all stakeholders in the esports ecosystem.

## 🏗️ Architecture

The project consists of two main components:

### 1. Smart Contract Oracle (`/oracle`)
- **EsportOracle.sol**: Main oracle contract with staking mechanism and consensus system
- **EsportOracleRequester.sol**: Extension contract for handling match data requests
- Built with **Solidity ^0.8.20** and **Foundry** framework
- Implements a decentralized consensus mechanism with node validation
- Features a punishment system for malicious actors

### 2. Go Data Provider (`/goapp`)
- **Go application** that fetches data from PandaScore API 
- Automatically submits match data to the smart contract
- Runs on a configurable cron schedule (default: every 15 minutes)
- Handles data transformation from API format to contract-compatible format

## 🔧 How does it work?

1. **Node Registration**: Data providers stake 0.001 ETH to become oracle nodes
2. **Data Collection**: Nodes fetch CS2 match data from PandaScore API
3. **Consensus Mechanism**: Multiple nodes submit the same data, requiring quorum (>50% + minimum 3 votes)
4. **Data Validation**: Smart contract validates and stores match results on-chain
5. **Punishment System**: Nodes submitting incorrect data are penalized or banned after 3 violations

## 🚀 Getting Started

### Prerequisites

- **Node.js** (v16+)
- **Go** (v1.24+)
- **Foundry** (for smart contract development)
- **Docker & Docker Compose** (for deployment)
- **PandaScore API Token** ([Get one here](https://pandascore.co/))

### Installation

1. **Clone the repository**
```bash
git clone https://github.com/PoCInnovation/esport-oracle.git
cd esport-oracle
```

2. **Set up the smart contract**
```bash
cd oracle
forge install
forge build
```

3. **Configure the Go application**
```bash
cd ../goapp
cp .env.sample .env
```

Edit the `.env` file with your configuration:
```env
CHAIN_ID=1337
PRIVATE_KEY=your_private_key_without_0x
CONTRACT_ADDRESS=deployed_contract_address_without_0x
CLIENT_ETH=http://localhost:8545
PANDASCORE_API_TOKEN=your_pandascore_api_token
CRON_SCHEDULE=*/15 * * * *
```

4. **Install Go dependencies**
```bash
go mod download
```

### Quick Deployment with Docker

For a complete local deployment:

```bash
cd oracle
chmod +x deploy_oracle.sh
./deploy_oracle.sh
```

This will:
- Start a local Anvil blockchain node
- Deploy the EsportOracle contract
- Provide you with the contract address for configuration

### Manual Deployment

#### 1. Deploy Smart Contract

```bash
cd oracle

# Start local blockchain
anvil

# Deploy contract (in another terminal)
forge create "src/esportOracle.sol:EsportOracle" \
  --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
  --rpc-url http://localhost:8545
```

#### 2. Run the Go Oracle Node

```bash
cd goapp
go run cmd/myapp/main.go
```

## 📖 Usage

### For Oracle Node Operators

1. **Stake ETH to become a node**:
```bash
cast send --private-key <your_private_key> \
  --value 0.001ether \
  "<CONTRACT_ADDRESS>" \
  "addFundToStaking()"
```

2. **Start the Go application**:
```bash
go run cmd/myapp/main.go
```

The application will automatically:
- Fetch match data every 15 minutes (configurable)
- Submit data to the smart contract
- Handle consensus and validation

### For Smart Contract Interactions

**Get pending matches**:
```bash
cast call "<CONTRACT_ADDRESS>" "getPendingMatches()(bytes32[])"
```

**Get match by ID**:
```bash
cast call "<CONTRACT_ADDRESS>" "getMatchById(uint256)(tuple)" <MATCH_ID>
```

**Get listed nodes**:
```bash
cast call "<CONTRACT_ADDRESS>" "getListedNodes()(address[])"
```

**Withdraw stake**:
```bash
cast send --private-key <your_private_key> \
  "<CONTRACT_ADDRESS>" \
  "withdrawStake()"
```

### For DApp Developers

You can integrate the oracle data into your applications:

```solidity
interface IEsportOracle {
    struct Match {
        uint256 _id;
        Opponents[] _opponents;
        Games[] _game;
        Result[] _result;
        uint256 _winnerId;
        uint256 _beginAt;
    }
    
    function getMatchById(uint256 matchId) external view returns (Match memory);
    function getPendingMatches() external view returns (bytes32[] memory);
}
```

## 🧪 Testing

### Smart Contract Tests

```bash
cd oracle
forge test -vvv
```

### Go Application Tests (need to be implemented)

```bash
cd goapp
go test ./...
```

## 📁 Project Structure

```
esport-oracle/
├── oracle/                    # Smart contract components
│   ├── src/
│   │   ├── esportOracle.sol          # Main oracle contract
│   │   └── esportOracleRequester.sol # Request handling extension
│   ├── test/                         # Foundry tests
│   ├── script/                       # Deployment scripts
│   └── docker-compose.yml           # Local deployment
├── goapp/                     # Go data provider
│   ├── cmd/myapp/            # Application entry point
│   ├── internal/
│   │   ├── client/           # API and blockchain clients
│   │   ├── contract/         # Generated contract bindings
│   │   ├── models/           # Data models
│   │   └── service/          # Business logic
│   └── .env.sample          # Environment configuration template
├── .github/                  # GitHub workflows and templates
└── README.md                # This file
```

## 🔐 Security Features

- **Staking Mechanism**: Nodes must stake ETH to participate
- **Consensus System**: Requires majority agreement for data validation
- **Punishment System**: Progressive penalties for incorrect submissions
- **Node Banning**: Automatic banning after 3 violations
- **Owner Controls**: Administrative functions for emergency situations
- **Pausable Contract**: Can be paused in case of emergency

## 🌐 API Integration

The oracle currently integrates with:
- **PandaScore API**: For CS2 match data, statistics, and tournament information

### Data Sources

- Match results and scores
- Team information and rosters
- Tournament brackets and schedules
- Player statistics
- Live match updates

## 🔧 Configuration

### Smart Contract Parameters

- **Staking Amount**: 0.001 ETH (fixed)
- **Maximum Violations**: 3 violations before ban
- **Punishment Amount**: 0.0001 ETH per violation
- **Quorum Requirement**: >50% of nodes + minimum 3 votes

### Go Application Configuration

- **Update Frequency**: Configurable via `CRON_SCHEDULE`
- **API Rate Limiting**: Handled automatically
- **Error Handling**: Automatic retry with exponential backoff
- **Data Validation**: Schema validation before submission

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](./CONTRIBUTING.md) for details.

### Development Workflow

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

### Code Standards

- **Solidity**: Follow Foundry conventions
- **Go**: Use `gofmt` and follow Go best practices
- **Commits**: Use conventional commit messages
- **Testing**: Maintain test coverage

## 📄 License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## 🔗 Links

- **Documentation**: [Foundry Book](https://book.getfoundry.sh/)
- **PandaScore API**: [Documentation](https://pandascore.co/docs)
- **Support**: For questions, contact our team or open an issue

## Get involved

You're invited to join this project! Check out the [contributing guide](./CONTRIBUTING.md).

If you're interested in how the project is organized at a higher level, please contact the current project manager.

## Our PoC team ❤️

Developers
| [<img src="https://avatars.githubusercontent.com/u/146731380?s=400&u=dcf3f96207f69a4ad4f9317791e3502d0bc18757&v=4" width=85><br><sub>[Jules Lordet]</sub>](https://github.com/L3yserEpitech) | [<img src="https://avatars.githubusercontent.com/u/141178010?v=4" width=85><br><sub>[Gregoire Caseaux]</sub>](https://github.com/Nezketsu) | [<img src="https://avatars.githubusercontent.com/u/146000349?v=4" width=85><br><sub>[Grégor Sternat]</sub>](https://github.com/gregorsternat) | [<img src="https://avatars.githubusercontent.com/u/181359842?v=4" width=85><br><sub>[Aaron Aniambossou]</sub>](https://github.com/Aa-D-Wlter) 
| :---: | :---: | :---: | :---: |

Manager
| [<img src="https://avatars.githubusercontent.com/u/78302154?v=4" width=85><br><sub>[Lucas Leclerc]</sub>](https://github.com/Intermarch3) | [<img src="https://avatars.githubusercontent.com/u/117595009?v=4" width=85><br><sub>[Sacha Dujardin]</sub>](https://github.com/Sacharbon)
| :---: | :---: |

<h2 align=center>
Organization
</h2>

<p align='center'>
    <a href="https://www.linkedin.com/company/pocinnovation/mycompany/">
        <img src="https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white" alt="LinkedIn logo">
    </a>
    <a href="https://www.instagram.com/pocinnovation/">
        <img src="https://img.shields.io/badge/Instagram-E4405F?style=for-the-badge&logo=instagram&logoColor=white" alt="Instagram logo"
>
    </a>
    <a href="https://twitter.com/PoCInnovation">
        <img src="https://img.shields.io/badge/Twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white" alt="Twitter logo">
    </a>
    <a href="https://discord.com/invite/Yqq2ADGDS7">
        <img src="https://img.shields.io/badge/Discord-7289DA?style=for-the-badge&logo=discord&logoColor=white" alt="Discord logo">
    </a>
</p>
<p align=center>
    <a href="https://www.poc-innovation.fr/">
        <img src="https://img.shields.io/badge/WebSite-1a2b6d?style=for-the-badge&logo=GitHub Sponsors&logoColor=white" alt="Website logo">
    </a>
</p>

> 🚀 Don't hesitate to follow us on our different networks, and put a star 🌟 on `PoC's` repositories

> Made with ❤️ by PoC