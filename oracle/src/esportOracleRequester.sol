// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./esportOracle.sol";

contract EsportOracleRequester is EsportOracle {
    // Struct to represent a match request
    struct MatchRequest {
        uint256 matchId;
        address requester;
        uint256 fee;
        bool fulfilled;
    }

    // Counter for request IDs
    uint256 private _requestCounter;

    // Mapping to store match requests
    mapping(uint256 => MatchRequest) public _matchRequests;

    // Minimum fee required to request a match
    uint256 public constant MIN_REQUEST_FEE = 0.0001 ether;

    // Event emitted when a match is requested
    event MatchRequested(uint256 indexed requestId, uint256 indexed matchId, address indexed requester, uint256 fee);

    constructor() EsportOracle() {
        _requestCounter = 0;
    }

    /**
     * @notice Allow a user to request a match by providing the match ID
     * @param matchId Id of the match
     * @return requestId Id of the request
     */
    function requestMatch(uint256 matchId) external payable whenNotPaused returns (uint256) {
        require(msg.value >= MIN_REQUEST_FEE, "Insufficient request fee");
        require(_matchMapping[matchId]._id == 0, "Match already registered");

        _requestCounter++;
        uint256 requestId = _requestCounter;

        _matchRequests[requestId] = MatchRequest({
            matchId: matchId,
            requester: msg.sender,
            fee: msg.value,
            fulfilled: false
        });

        emit MatchRequested(requestId, matchId, msg.sender, msg.value);

        return requestId;
    }

    /**
     * @notice Get the details of a match request by its ID
     * @param requestId Id of the request
     * @return MatchRequest struct containing the details of the request
     */
    function getMatchRequest(uint256 requestId) external view returns (MatchRequest memory) {
        return _matchRequests[requestId];
    }

    /**
     * @notice Check if a match has been requested
     * @param matchId Id of the match to check
     * @return bool True if the match has been requested, false otherwise
     */
    function isMatchRequested(uint256 matchId) external view returns (bool) {
        for (uint256 i = 1; i <= _requestCounter; i++) {
            if (_matchRequests[i].matchId == matchId && !_matchRequests[i].fulfilled) {
                return true;
            }
        }
        return false;
    }

    /**
     * @notice Mark a match request as fulfilled
     * @param matchId Id of the match to mark as fulfilled
     */
    function markRequestsFulfilled(uint256 matchId) external onlyOwner {
        for (uint256 i = 1; i <= _requestCounter; i++) {
            if (_matchRequests[i].matchId == matchId && !_matchRequests[i].fulfilled) {
                _matchRequests[i].fulfilled = true;
            }
        }
    }
    
    /**
     * @notice Get a list of all match requests that have not been fulfilled
     * @return uint256[] Array of match IDs that have pending requests
     */
    function getPendingRequestedMatches() external view returns (uint256[] memory) {
        uint256 count = 0;

        for (uint256 i = 1; i <= _requestCounter; i++) {
            if (!_matchRequests[i].fulfilled) {
                count++;
            }
        }
        
        uint256[] memory pendingMatches = new uint256[](count);
        uint256 index = 0;
        
        for (uint256 i = 1; i <= _requestCounter; i++) {
            if (!_matchRequests[i].fulfilled) {
                pendingMatches[index] = _matchRequests[i].matchId;
                index++;
            }
        }
        
        return pendingMatches;
    }
}