// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./esportOracle.sol";

import "./esportOracleTypes.sol";

import "./esportOracleClientRequester.sol";

interface InterfaceOracle {
    function requestMatch(uint256 matchId) external payable returns (uint256);
    function getMatchRequest(uint256 requestId) external view returns (EsportOracleTypes.MatchRequest memory);
    function isMatchRequested(uint256 matchId) external view returns (bool);
    function markRequestsFulfilled(uint256 matchId) external;
    function getPendingRequestedMatches() external view returns (uint256[] memory);
    function getMatchById(uint256 matchId) external view returns (EsportOracleTypes.Match memory);
}

contract EsportOracleRequester is EsportOracle {
    // Use to library MatchRequest
    using EsportOracleTypes for EsportOracleTypes.MatchRequest;

    // Counter for request IDs
    uint256 private _requestCounter;

    // Mapping to store match requests
    mapping(uint256 => EsportOracleTypes.MatchRequest) public _matchRequests;

    // Mapping to store match retrieved by matchRequest
    mapping(uint256 => uint256) private _matchIdToRequestId;

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

        _matchRequests[requestId] = EsportOracleTypes.MatchRequest({
            matchId: matchId,
            requester: msg.sender,
            fee: msg.value,
            fulfilled: false
        });

        _matchIdToRequestId[matchId] = requestId;

        emit MatchRequested(requestId, matchId, msg.sender, msg.value);

        return requestId;
    }

    /**
     * @notice Get the details of a match request by its ID
     * @param matchId Id of the request
     * @return MatchRequest struct containing the details of the request
     */
    function getMatchRequest(uint256 matchId) external view returns (EsportOracleTypes.MatchRequest memory) {
        for (uint256 i = 1; i <= _requestCounter; i++) {
            if (_matchRequests[i].matchId == matchId) {
                return _matchRequests[i];
            }
        }
        return EsportOracleTypes.MatchRequest(0, address(0), 0, false);
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
    function markRequestsFulfilled(uint256 matchId) public {
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

    function getRequestByMatchId(uint256 matchId) public view returns (EsportOracleTypes.MatchRequest memory) {
        uint256 requestId = _matchIdToRequestId[matchId];
        return _matchRequests[requestId];
    }

    function callMatchOracle(uint256 matchId, EsportOracleTypes.Match memory matchData) external {
        EsportOracleTypes.MatchRequest memory match_request = getRequestByMatchId(matchId);
        require(match_request.requester != address(0), "Invalid address requester");
        require(_matchMapping[matchId]._id != 0, "Match not validated by quorum yet");
        EsportOracleClientRequester(match_request.requester).callMatchReceived(matchData);
        markRequestsFulfilled(matchId);
    }
}
