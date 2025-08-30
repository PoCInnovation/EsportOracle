// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./esportOracleRequester.sol";
import "./esportOracleTypes.sol";

abstract contract EsportOracleClientRequester {
    address public immutable _owner;
    using EsportOracleTypes for EsportOracleTypes.MatchRequest;
    uint256 public requestId;
    InterfaceOracle public immutable oracle;

    constructor(address _oracle) {
        _owner = msg.sender;
        oracle = InterfaceOracle(_oracle);
    }

    modifier onlyOracle() {
        require(msg.sender == address(oracle), "Only the oracle contract can call this function");
        _;
    }


    /**
     * @notice Allow a user to request a match by providing the match ID
     * @param matchId Id of the match
     * @return requestId Id of the request
     */
    function receiveMatch(uint256 matchId) payable external returns (uint256) {
        require(oracle.isMatchRequested(matchId) == false, "Match already requested");
        requestId = oracle.requestMatch{value: msg.value}(matchId);
        return requestId;
    }

    /**
     * @notice Show the details of a match request by its ID
     * @param matchId Id of the request
     * @return MatchRequest struct containing the details of the request
    */
    function showMatch(uint256 matchId) external view returns(EsportOracleTypes.Match memory) {
        return (oracle.getMatchById(matchId));
    }

    /**
     * @notice Get a list of all match requests that have not been fulfilled
     * @return uint256[] Array of match IDs that have pending requests
     */
    function showPendingRequestedMatches() external view returns (uint256[] memory) {
        return (oracle.getPendingRequestedMatches());
    }

    function callMatchReceived(EsportOracleTypes.Match memory _match) external virtual onlyOracle {}
}

