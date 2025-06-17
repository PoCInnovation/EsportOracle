// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

library lib {
    struct MatchRequest {
        uint256 matchId;
        address requester;
        uint256 fee;
        bool fulfilled;
    }
}