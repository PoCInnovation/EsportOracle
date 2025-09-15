// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./esportOracleClientRequester.sol";
import "./esportOracleTypes.sol";

contract MockBettingClient is EsportOracleClientRequester {
    using EsportOracleTypes for EsportOracleTypes.Match;

    event MatchDelivered(uint256 indexed matchId, uint256 winnerId);

    uint256 private _lastMatchId;
    uint256 private _lastWinnerId;

    constructor(address oracle_) EsportOracleClientRequester(oracle_) {}

    function lastMatchId() external view returns (uint256) { return _lastMatchId; }
    function lastWinnerId() external view returns (uint256) { return _lastWinnerId; }

    function callMatchReceived(EsportOracleTypes.Match memory _match) external override onlyOracle {
        _lastMatchId = _match._id;
        _lastWinnerId = _match._winnerId;
        emit MatchDelivered(_lastMatchId, _lastWinnerId);
    }
}


