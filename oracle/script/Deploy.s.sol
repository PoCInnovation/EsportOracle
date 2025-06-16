// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/esportOracle.sol";
import "../src/esportOracleClientRequester.sol";

contract DeploySellManager is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        EsportOracleRequester requester = new EsportOracleRequester();
        EsportOracleClientRequester esportOracle = new EsportOracleClientRequester(address(requester));

        vm.stopBroadcast();

        console.log("EsportOracle deployed at:", address(esportOracle));
    }
}
