// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "../src/BetContract.sol";
import "../lib/openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";

contract SimpleTestBetting is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        
        address betContractAddr = vm.envAddress("BET_CONTRACT_ADDRESS");
        address tokenAddr = vm.envAddress("TOKEN_ADDRESS");
        
        console.log("=== SIMPLE BETTING TEST ===");
        console.log("Deployer:", deployer);
        console.log("Current timestamp:", block.timestamp);
        
        vm.startBroadcast(deployerPrivateKey);
        
        BetContract betContract = BetContract(payable(betContractAddr));
        IERC20 token = IERC20(tokenAddr);
        
        // Vérifier les soldes actuels
        console.log("Deployer token balance:", token.balanceOf(deployer));
        console.log("Deployer ETH balance:", deployer.balance);
        
        // Créer un pari avec une deadline dans le futur
        uint256 deadline = block.timestamp + 3600; // 1 heure dans le futur
        console.log("Creating bet with deadline:", deadline);
        
        try betContract.createBet(
            "Test Bet - Team A vs Team B",
            111, // team1Id
            222, // team2Id  
            deadline,
            888  // matchId
        ) {
            console.log("[OK] Bet created successfully");
            
            // Vérifier que le pari a été créé
            uint256 betCount = betContract.getBetCount();
            console.log("Total bets:", betCount);
            
            if (betCount > 0) {
                BetContract.Bet memory bet = betContract.getBet(betCount - 1);
                console.log("Latest bet deadline:", bet.deadline);
                console.log("Latest bet description:", bet.description);
            }
            
        } catch Error(string memory reason) {
            console.log("[ERROR] Create bet failed:", reason);
        } catch {
            console.log("[ERROR] Create bet failed: Unknown error");
        }
        
        vm.stopBroadcast();
        console.log("=== TEST COMPLETE ===");
    }
}
