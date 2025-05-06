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

    constructor() {
        _owner = msg.sender;
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
    function addNewMatch(Match[] memory newMatch) external {
        for (uint8 i = 0; i < newMatch.length; i++) {
            uint256 matchId = newMatch[i]._id;

            _matchMapping[matchId]._id = matchId;
            _matchMapping[matchId]._winnerId = newMatch[i]._winnerId;
            _matchMapping[matchId]._beginAt = newMatch[i]._beginAt;
            
            for (uint256 j = 0; j < newMatch[i]._opponents.length; j++) {
                Opponents memory opponent = newMatch[i]._opponents[j];
                _matchMapping[matchId]._opponents.push(opponent);
            }

            for (uint256 j = 0; j < newMatch[i]._game.length; j++) {
                Games memory game = newMatch[i]._game[j];
                _matchMapping[matchId]._game.push(game);
            }
            
            for (uint256 j = 0; j < newMatch[i]._result.length; j++) {
                Result memory result = newMatch[i]._result[j];
                _matchMapping[matchId]._result.push(result);
            }
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
    function addNewNode() external {
        require(msg.sender != address(0), "New node cannot be zero address");
        listedNodes.push(msg.sender);
        emit newNodeAdded(msg.sender);
    }

    /**
     * @notice function to return the list of nodes addresses
     * @return The list of addresses of the nodes
     */
    function getListedNodes() external view returns (address[] memory) {
        return listedNodes;
    }
}
