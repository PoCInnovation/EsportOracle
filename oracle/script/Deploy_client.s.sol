// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/esportOracle.sol";
import "../src/esportOracleClientRequester.sol";


contract DeployClient is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        EsportOracleRequester requester = new EsportOracleRequester();

        vm.stopBroadcast();

        console.log("Oracle deployed at:", address(requester));
    }
}
