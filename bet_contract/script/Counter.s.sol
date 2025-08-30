// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {BetContract} from "../src/BetContract.sol";

contract BetContractScript is Script {
    BetContract public betContract;

    function setUp() public {}

    function run() public {
        vm.startBroadcast();

        // Remplacez ces adresses par les vraies si besoin
        address oracle = 0x000000000000000000000000000000000000dEaD;
        address token = 0x000000000000000000000000000000000000dEaD;

        betContract = new BetContract(oracle, token);

        vm.stopBroadcast();
    }
}
