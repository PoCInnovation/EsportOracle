// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

library EsportOracleTypes {
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

    struct NodeViolation {
        uint256 incorrectMatches;
        bool isBanned;
    }

    struct MatchRequest {
        uint256 matchId;
        address requester;
        uint256 fee;
        bool fulfilled;
    }
}