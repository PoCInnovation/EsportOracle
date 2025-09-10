// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "../src/BetContract.sol";

contract TestBalance is Script {
    function run() external {
        // Utiliser la clé privée du premier compte Anvil
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        
        console.log("=== Test Balance ===");
        console.log("Deployer address:", deployer);
        console.log("Deployer ETH balance:", deployer.balance);
        console.log("Gas price:", tx.gasprice);
        
        // Récupérer les adresses des contrats depuis l'environnement
        address betContractAddr = vm.envAddress("BET_CONTRACT_ADDRESS");
        address tokenAddr = vm.envAddress("TOKEN_ADDRESS");
        address oracleAddr = vm.envAddress("ORACLE_ADDRESS");
        
        console.log("BetContract address:", betContractAddr);
        console.log("Token address:", tokenAddr);
        console.log("Oracle address:", oracleAddr);
        
        // Vérifier que les contrats existent
        console.log("BetContract code size:", betContractAddr.code.length);
        console.log("Token code size:", tokenAddr.code.length);
        console.log("Oracle code size:", oracleAddr.code.length);
        
        // Test simple : vérifier le solde ETH du BetContract
        console.log("BetContract ETH balance:", betContractAddr.balance);
        
        // Commencer une transaction pour tester
        vm.startBroadcast(deployerPrivateKey);
        
        BetContract betContract = BetContract(payable(betContractAddr));
        uint256 contractBalance = betContract.getContractETHBalance();
        console.log("Contract ETH balance via function:", contractBalance);
        
        vm.stopBroadcast();
        
        console.log("=== Test termine avec succes ===");
    }
}
