# EsportOracle Smart Contract Schema

## Overview

The EsportOracle is a decentralized smart contract that enables secure collection, validation, and storage of esports data through a consensus-based oracle mechanism. It implements a stake-based validation system where nodes must stake ETH to participate in data validation.

## Contract Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                    EsportOracle Contract                    │
├─────────────────────────────────────────────────────────────┤
│  Governance Layer                                           │
│  ├── Owner Management                                       │
│  └── Pause/Unpause Functionality                            │
├─────────────────────────────────────────────────────────────┤
│  Node Management Layer                                      │
│  ├── Staking System (0.001 ETH)                             │
│  ├── Node Registration                                      │
│  └── Violation Tracking                                     │
├─────────────────────────────────────────────────────────────┤
│  Consensus Layer                                            │
│  ├── Match Data Submission                                  │
│  ├── Voting Mechanism                                       │
│  └── Quorum Validation                                      │
├─────────────────────────────────────────────────────────────┤
│  Punishment System                                          │
│  ├── Violation Detection                                    │
│  ├── Progressive Penalties                                  │
│  └── Fund Redistribution                                    │
├─────────────────────────────────────────────────────────────┤
│  Data Storage Layer                                         │
│  ├── Validated Matches                                      │
│  ├── Pending Matches                                        │
│  └── Node States                                            │
└─────────────────────────────────────────────────────────────┘
```

## Data Structures

### Core Match Structure

```solidity
struct Match {
    uint256 _id;              // Unique match identifier
    Opponents[] _opponents;   // Participating teams
    Games[] _game;            // Individual games/maps
    Result[] _result;         // Final scores
    uint256 _winnerId;        // Winning team ID
    uint256 _beginAt;         // Match start timestamp
}
```

### Supporting Structures

```solidity
struct Opponents {
    string _acronym;    // Team acronym (e.g., "NaVi")
    uint256 _id;        // Unique team identifier
    string _name;       // Full team name
}

struct Games {
    uint256 _id;        // Game/map identifier
    bool _finished;     // Completion status
    uint256 _winnerId;  // Game winner ID
}

struct Result {
    uint8 _score;       // Team score
    uint256 _teamId;    // Team identifier
}

struct NodeViolation {
    uint256 incorrectMatches;   // Count of incorrect submissions
    bool isBanned;              // Ban status
}
```

## State Variables

### Core Storage

| Variable | Type | Access | Description |
|----------|------|--------|-------------|
| `_owner` | `address` | `public` | Contract owner address |
| `_matchMapping` | `mapping(uint256 => Match)` | `public` | Validated match storage |
| `listedNodes` | `address[]` | `private` | Active validator nodes |
| `_fundsStaked` | `mapping(address => uint256)` | `public` | Staked amounts per node |
| `_nodeViolations` | `mapping(address => NodeViolation)` | `public` | Violation tracking |

### Consensus Tracking

| Variable | Type | Access | Description |
|----------|------|--------|-------------|
| `_matchVotes` | `mapping(bytes32 => uint8)` | `public` | Vote count per match hash |
| `_addressByHash` | `mapping(bytes32 => address[])` | `public` | Voters per match hash |
| `_pendingMatchesHashes` | `bytes32[]` | `public` | Pending validation hashes |

## Constants

```solidity
uint256 public constant MAX_VIOLATIONS = 3;
uint256 public constant PUNISHMENT_AMOUNT = 0.0001 ether;
uint256 private constant STAKING_AMOUNT = 0.001 ether;
```

## Function Categories

### Governance Functions

#### Owner Management
```solidity
function setOwner(address newOwner) external onlyOwner
```
- **Purpose**: Transfer contract ownership
- **Access**: Owner only

#### Pause Control
```solidity
function pause() external onlyOwner
function unpause() external onlyOwner
```
- **Purpose**: Emergency contract control
- **Access**: Owner only

### Staking Functions

#### Node Registration
```solidity
function addFundToStaking() external payable
```
- **Purpose**: Become a validator node
- **Requirements**: Exactly 0.001 ETH
- **Side Effects**: Adds to `listedNodes`
- **Events**: `stakingSuccess`, `newNodeAdded`

#### Stake Withdrawal
```solidity
function withdrawStake() external
```
- **Purpose**: Exit validator system
- **Requirements**: Not banned, has staked funds
- **Side Effects**: Removes from `listedNodes`

### Match Management Functions

#### Data Submission
```solidity
function handleNewMatches(Match[] memory newMatch) external
```
- **Purpose**: Submit match data for validation
- **Access**: Listed nodes only
- **Process**:
  1. Hash match data
  2. Record vote
  3. Check quorum
  4. Validate or punish

#### Data Retrieval
```solidity
function getMatchById(uint256 matchId) external view returns (Match memory)
function getPendingMatches() external view returns (bytes32[] memory)
```
- **Purpose**: Access validated/pending match data
- **Access**: Public

### Punishment Functions

#### Manual Ban
```solidity
function banNode(address node) external onlyOwner
```
- **Purpose**: Immediate node removal
- **Side Effects**: Fund confiscation and redistribution
- **Events**: `NodeBanned`

#### Rehabilitation
```solidity
function rehabilitateNode(address node) external onlyOwner
```
- **Purpose**: Restore banned node
- **Side Effects**: Reset violation counter

### View Functions

```solidity
function getListedNodes() external view returns (address[] memory)
function qorumIsReached(uint8 nbVote) private view returns (bool)
```

## Consensus Mechanism

### Quorum Formula
```solidity
return (listedNodes.length / 2) < nbVote && nbVote > 2;
```

### Validation Flow
```
┌─────────────────┐
│  Node Submits   │
│   Match Data    │
└─────────┬───────┘
          │
          ▼
┌─────────────────┐
│   Hash Data &   │
│   Record Vote   │
└─────────┬───────┘
          │
          ▼
┌─────────────────┐     NO     ┌─────────────────┐
│ Quorum Reached? │───────────▶│  Wait for More  │
└─────────┬───────┘            │     Votes       │
          │ YES                └─────────────────┘
          ▼
┌─────────────────┐
│   Validate &    │
│  Store Match    │
└─────────┬───────┘
          │
          ▼
┌─────────────────┐
│ Punish Minority │
│     Voters      │
└─────────────────┘
```

## Punishment System

### Violation Levels

| Violations | Action | Penalty | Status |
|------------|--------|---------|--------|
| 1-2 | Fine | 0.0001 ETH deduction | Active |
| 3+ | Ban | Full stake confiscation | Banned |

### Redistribution Logic
```solidity
// Equal distribution among correct voters
uint256 amountPerNode = confiscatedAmount / correctVoters.length;
for (uint i = 0; i < correctVoters.length; i++) {
    _fundsStaked[correctVoters[i]] += amountPerNode;
}
```

## Events

```solidity
event newNodeAdded(address indexed addressAdded);
event stakingSuccess(address indexed addressAdded, uint256 amount);
event NodePunished(address indexed node, uint256 amount, uint256 violationsCount);
event NodeBanned(address indexed node);
```

## Security Features

### Access Modifiers

```solidity
modifier onlyOwner()        // Owner-only functions
modifier onlyListedNodes()  // Validator-only functions  
modifier notBanned()        // Exclude banned nodes
modifier nodeAlreadyStake() // Prevent double staking
modifier whenNotPaused()    // Pause mechanism
```

### Attack Vectors & Mitigations

| Attack | Mitigation |
|--------|------------|
| **Sybil Attack** | Staking requirement (0.001 ETH) |
| **Data Manipulation** | Consensus voting + punishment |
| **Front-running** | Hash-based voting |
| **Griefing** | Progressive punishment system |
| **Centralization** | Quorum requirements |

## Integration Patterns

### External Systems Integration

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│  PandaScore API │───▶│   Go Oracle     │───▶│ Smart Contract  │
│                 │    │   Application   │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                              │                        │
                              ▼                        ▼
                       ┌─────────────────┐    ┌─────────────────┐
                       │   Data Format   │    │   Blockchain    │
                       │  Transformation │    │    Storage      │
                       └─────────────────┘    └─────────────────┘
```

### Typical Usage Flow

1. **Node Registration**: Stake 0.001 ETH via `addFundToStaking()`
2. **Data Collection**: Go app fetches from PandaScore API
3. **Data Submission**: Call `handleNewMatches()` with formatted data
4. **Consensus**: Multiple nodes vote on same data
5. **Validation**: Quorum reached, data stored permanently
6. **Punishment**: Minority voters penalized automatically

## Deployment Information

### Constructor
```solidity
constructor() {
    _owner = msg.sender;
    nbMatchSent = 0;
}
```

### Dependencies
- OpenZeppelin Pausable contract
- Solidity ^0.8.20

### Network Compatibility
- Ethereum Mainnet
- Ethereum Testnets
- Local development (Anvil/Hardhat)