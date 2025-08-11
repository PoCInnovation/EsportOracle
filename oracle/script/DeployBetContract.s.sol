// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import "../src/BetContract.sol";
import "../src/esportOracle.sol";
import "../test/BetContractTest.t.sol"; // Pour MockERC20

contract DeployBetContract is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        
        vm.startBroadcast(deployerPrivateKey);
        
        // 1. Déployer un token ERC20 pour les tests
        MockERC20 token = new MockERC20();
        console.log("MockERC20 deployed at:", address(token));
        
        // 2. Déployer l'oracle (nécessaire pour BetContract)
        EsportOracle oracle = new EsportOracle();
        console.log("EsportOracle deployed at:", address(oracle));
        
        // 3. Déployer le contrat de paris
        BetContract betContract = new BetContract(address(oracle), address(token));
        console.log("BetContract deployed at:", address(betContract));
        
        // 4. Configurer des tokens pour les tests
        address deployer = vm.addr(deployerPrivateKey);
        token.mint(deployer, 10000 * 10**18); // 10,000 tokens pour les tests
        console.log("Minted 10,000 tokens to deployer:", deployer);
        
        // 5. Déposer des ETH dans le contrat pour les fees oracle
        betContract.depositForFees{value: 0.1 ether}();
        console.log("Deposited 0.1 ETH for oracle fees");
        
        vm.stopBroadcast();
        
        console.log("\n=== DEPLOYMENT SUMMARY ===");
        console.log("Token Address:", address(token));
        console.log("Oracle Address:", address(oracle));
        console.log("BetContract Address:", address(betContract));
        console.log("Deployer Address:", deployer);
        console.log("Contract ETH Balance:", betContract.getContractETHBalance());
    }
}
