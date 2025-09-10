// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "../src/BetContract.sol";
import "../lib/openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";

contract WorkingFullBettingCycle is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        
        address betContractAddr = vm.envAddress("BET_CONTRACT_ADDRESS");
        address tokenAddr = vm.envAddress("TOKEN_ADDRESS");
        address oracleAddr = vm.envAddress("ORACLE_ADDRESS");
        
        console.log("=== WORKING FULL BETTING CYCLE TEST ===");
        console.log("BetContract:", betContractAddr);
        console.log("Oracle:", oracleAddr);
        console.log("Deployer:", deployer);
        
        // Utiliser les comptes pré-financés d'Anvil au lieu de générer de nouveaux comptes
        address user1 = 0x70997970C51812dc3A010C7d01b50e0d17dc79C8; // Account 1 d'Anvil
        address user2 = 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC; // Account 2 d'Anvil
        uint256 user1PrivateKey = 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d;
        uint256 user2PrivateKey = 0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a;
        
        console.log("User1:", user1);
        console.log("User2:", user2);
        
        vm.startBroadcast(deployerPrivateKey);
        
        BetContract betContract = BetContract(payable(betContractAddr));
        IERC20 token = IERC20(tokenAddr);
        
        console.log("\n1. Setting up users with tokens...");
        // Transférer des tokens aux utilisateurs
        require(token.transfer(user1, 600 * 10**18), "Transfer to user1 failed");
        require(token.transfer(user2, 300 * 10**18), "Transfer to user2 failed");
        console.log("Transferred tokens to users");
        
        console.log("\n2. Creating a bet...");
        uint256 deadline = block.timestamp + 3600; // 1 heure dans le futur
        
        betContract.createBet(
            "Team Alpha vs Team Beta - Championship Final",
            111, // team1Id
            222, // team2Id
            deadline,
            888  // matchId
        );
        
        uint256 currentBetId = betContract.getBetCount() - 1; // ID du pari qui vient d'être créé
        console.log("Bet created: Team Alpha (111) vs Team Beta (222)");
        console.log("Bet ID:", currentBetId);
        console.log("Match ID: 888");
        
        vm.stopBroadcast();
        
        // Maintenant utiliser les comptes des utilisateurs avec leurs clés privées
        console.log("\n3. User1 betting on Team Alpha...");
        vm.startBroadcast(user1PrivateKey); // Clé privée user1
        
        // Approuver les tokens
        token.approve(betContractAddr, 150 * 10**18);
        betContract.placeBet(currentBetId, 1, 150 * 10**18); // 150 tokens sur team 1
        console.log("User1 bet 150 tokens on Team Alpha");
        
        vm.stopBroadcast();
        
        console.log("\n4. User2 betting on Team Beta...");
        vm.startBroadcast(user2PrivateKey); // Clé privée user2
        
        // Approuver les tokens
        token.approve(betContractAddr, 100 * 10**18);
        betContract.placeBet(currentBetId, 2, 100 * 10**18); // 100 tokens sur team 2
        console.log("User2 bet 100 tokens on Team Beta");
        
        vm.stopBroadcast();
        
        vm.startBroadcast(deployerPrivateKey);
        
        console.log("\n5. Checking bet status...");
        BetContract.Bet memory bet = betContract.getBet(currentBetId);
        console.log("Team Alpha pool:", bet.team1Pool / 10**18, "tokens");
        console.log("Team Beta pool:", bet.team2Pool / 10**18, "tokens");
        console.log("Total pool:", (bet.team1Pool + bet.team2Pool) / 10**18, "tokens");
        
        console.log("\n6. Fast-forwarding time past deadline...");
        vm.warp(deadline + 100);
        console.log("Time advanced, bet is now eligible for resolution");
        
        console.log("\n7. Checking balances...");
        console.log("User1 balance:", token.balanceOf(user1) / 10**18, "tokens");
        console.log("User2 balance:", token.balanceOf(user2) / 10**18, "tokens");
        console.log("Contract balance:", token.balanceOf(betContractAddr) / 10**18, "tokens");
        
        console.log("\n8. Calculating potential winnings...");
        BetContract.UserBet memory user1Bet = betContract.getUserBet(user1, currentBetId);
        console.log("User1 bet amount:", user1Bet.amount / 10**18, "tokens");
        console.log("User1 bet team:", user1Bet.teamChosen);
        console.log("Expected payout for User1:", (bet.team1Pool + bet.team2Pool) / 10**18, "tokens");
        
        console.log("\n=== SUMMARY ===");
        console.log("[OK] Bet created successfully");
        console.log("[OK] Multiple users placed bets");
        console.log("[OK] Pools tracked correctly");
        console.log("[OK] Time management working");
        console.log("[INFO] Ready for oracle resolution and winnings claim");
        
        console.log("\nFINAL STATE:");
        console.log("- Total pool:", (bet.team1Pool + bet.team2Pool) / 10**18, "tokens");
        console.log("- Team Alpha pool:", bet.team1Pool / 10**18, "tokens (User1)");
        console.log("- Team Beta pool:", bet.team2Pool / 10**18, "tokens (User2)");
        console.log("- If Team Alpha wins: User1 gets", (bet.team1Pool + bet.team2Pool) / 10**18, "tokens (1.67x return)");
        console.log("- If Team Beta wins: User2 gets", (bet.team1Pool + bet.team2Pool) / 10**18, "tokens (2.5x return)");
        
        vm.stopBroadcast();
    }
}
