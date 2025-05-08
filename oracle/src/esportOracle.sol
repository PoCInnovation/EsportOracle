// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract EsportOracle {
    address public _owner;

    struct Result {
        uint8 _score;
        uint256 _teamId;
    }

    struct Games {
        uint256 _id;
        bool _finished;
        uint256 _winnerId;
    }

    struct Opponents {
        string _acronym;
        uint256 _id;
        string _name;
    }

    struct Match {
        uint256 _id;
        Opponents[] _opponents;
        Games[] _game;
        Result[] _result;
        uint256 _winnerId;
        uint256 _beginAt;
    }

    mapping(uint256 => Match) public _matchMapping;
    address[] private listedNodes;
    mapping(bytes32 => uint8) public _matchVotes;
    mapping(bytes32 => address[]) public _addressByHash;
    bytes32[] public _pendingMatchesHashes;
    uint8 nbMatchSent;

    constructor() {
        _owner = msg.sender;
        nbMatchSent = 0;
    }

    /**
     * @notice event emitted when a new node is added
     * @param addressAdded The address of the owner of the new node
     */
    event newNodeAdded(address indexed addressAdded);

    modifier onlyOwner() {
        require(msg.sender == _owner, "Not the contract owner");
        _;
    }

    /**
     * @notice verify if the calling address is listed
     */
    modifier onlyListedNodes() {
        bool isListed = false;

        for (uint i = 0; i < listedNodes.length; i++) {
            if (listedNodes[i] == msg.sender) {
                isListed = true;
                break;
            }
        }
        require(isListed == true, "Node is not listed, please call ...");
        _;
    }

    /**
     * @notice verify if the calling address is already listed
     */
    modifier nodeAlreadyListed() {
        bool isListed = false;

        for (uint i = 0; i < listedNodes.length; i++) {
            if (listedNodes[i] == msg.sender) {
                isListed = true;
                break;
            }
        }
        require(isListed == false, "Node is already listed");
        _;
    }

    /**
     * @notice Set the new owner of the contract
     * @param newOwner The address of the new owner
     * @dev Only the current owner can call this function
     */
    function setOwner(address newOwner) public onlyOwner {
        require(newOwner != address(0), "New owner cannot be zero address");
        _owner = newOwner;
    }

    /**
     * @notice add match blockchain
     * @param newMatch a tab of Match
     */
    function addNewMatch(Match memory newMatch) internal {
        uint256 matchId = newMatch._id;

        _matchMapping[matchId]._id = matchId;
        _matchMapping[matchId]._winnerId = newMatch._winnerId;
        _matchMapping[matchId]._beginAt = newMatch._beginAt;

        for (uint256 j = 0; j < newMatch._opponents.length; j++) {
            Opponents memory opponent = newMatch._opponents[j];
            _matchMapping[matchId]._opponents.push(opponent);
        }

        for (uint256 j = 0; j < newMatch._game.length; j++) {
            Games memory game = newMatch._game[j];
            _matchMapping[matchId]._game.push(game);
        }

        for (uint256 j = 0; j < newMatch._result.length; j++) {
            Result memory result = newMatch._result[j];
            _matchMapping[matchId]._result.push(result);
        }
    }
    
    /**
     * @notice returns the match by id
     * @param matchId The id of the match
     * @return The match object
     * @dev This function retrieves a match by its ID from the mapping
     */
    function getMatchById(uint256 matchId) external view returns (Match memory) {
        return (_matchMapping[matchId]);
    }

    /**
     * @notice function to add a new node
     */
    function addNewNode() external nodeAlreadyListed {
        require(msg.sender != _owner, "owner cannot be a node");
        require(msg.sender != address(0), "New node cannot be zero address");
        listedNodes.push(msg.sender);
        emit newNodeAdded(msg.sender);
    }

    /**
     * @notice function called by listed nodes only, to register new matches
     * @param newMatch : a list of matches to register
     */
    function handleNewMatches(Match[] memory newMatch) external onlyListedNodes {
        require(newMatch.length > 0, "No match data provided");
        nbMatchSent++;
        for (uint256 i = 0; i < newMatch.length; i++) {
            bytes32 matchHash = keccak256(abi.encode(newMatch[i]));
            _matchVotes[matchHash]++;
            if (_matchVotes[matchHash] == 1) {
                _pendingMatchesHashes.push(matchHash);
                _addressByHash[matchHash].push(msg.sender);
            }
            if (_matchVotes[matchHash] == 3) {
                addNewMatch(newMatch[i]);
            }
        }
        if (nbMatchSent == listedNodes.length) {
            for (uint8 i = 0; i < _pendingMatchesHashes.length; i++) {
                delete(_matchVotes[_pendingMatchesHashes[i]]);
                delete(_addressByHash[_pendingMatchesHashes[i]]);
            }
            delete(_pendingMatchesHashes);
            nbMatchSent = 0;
        }
    }

    /**
     * @notice function to return the list pending match hash
     * @return The list of hashes
     */
    function getPendingMatches() external view returns (bytes32[] memory) {
        return (_pendingMatchesHashes);
    }

    /**
     * @notice function to return the list of nodes addresses
     * @return The list of addresses of the nodes
     */
    function getListedNodes() external view returns (address[] memory) {
        return listedNodes;
    }
}
