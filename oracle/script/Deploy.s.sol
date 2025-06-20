// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/esportOracleRequester.sol";
import "../src/esportOracle.sol";
import "../src/esportOracleClientRequester.sol";

contract DeploySellManager is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        EsportOracleRequester esportOracleRequester = new EsportOracleRequester();

        vm.stopBroadcast();

        console.log("EsportOracleRequester deployed at:", address(esportOracleRequester));
    }
}
