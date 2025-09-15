// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "forge-std/console.sol";

import {EsportOracle} from "../src/esportOracle.sol";
import {EsportOracleTypes} from "../src/esportOracleTypes.sol";

contract OracleBroadcast is Script {
    using EsportOracleTypes for *;

    function run() external {
        uint256 pk1 = vm.envUint("EOA_PK1");
        uint256 pk2 = vm.envUint("EOA_PK2");

        vm.startBroadcast(pk1);
        EsportOracle oracle = new EsportOracle();
        console.log("Oracle:", address(oracle));
        vm.stopBroadcast();

        vm.startBroadcast(pk1);
        oracle.addFundToStaking{value: 0.001 ether}();
        vm.stopBroadcast();

        vm.startBroadcast(pk2);
        oracle.addFundToStaking{value: 0.001 ether}();
        vm.stopBroadcast();

        EsportOracleTypes.Match memory m;
        m._id = 42;
        m._winnerId = 7;
        m._beginAt = 1;
        m._opponents = new EsportOracleTypes.Opponents[](2);
        m._opponents[0] = EsportOracleTypes.Opponents({ _acronym: "A", _id: 1, _name: "Alpha" });
        m._opponents[1] = EsportOracleTypes.Opponents({ _acronym: "B", _id: 2, _name: "Beta" });
        m._game = new EsportOracleTypes.Games[](1);
        m._game[0] = EsportOracleTypes.Games({ _id: 1001, _finished: true, _winnerId: 7 });
        m._result = new EsportOracleTypes.Result[](2);
        m._result[0] = EsportOracleTypes.Result({ _score: 2, _teamId: 7 });
        m._result[1] = EsportOracleTypes.Result({ _score: 1, _teamId: 8 });

        EsportOracleTypes.Match[] memory arr = new EsportOracleTypes.Match[](1);
        arr[0] = m;

        vm.startBroadcast(pk1);
        oracle.handleNewMatches(arr);
        vm.stopBroadcast();

        vm.startBroadcast(pk2);
        oracle.handleNewMatches(arr);
        vm.stopBroadcast();

        EsportOracleTypes.Match memory stored = oracle.getMatchById(42);
        console.log("Stored match id:", stored._id);
        console.log("Winner:", stored._winnerId);
        require(stored._id == 42 && stored._winnerId == 7, "match not stored");

        console.log("Broadcast OK");
    }
}
