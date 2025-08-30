// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import "../src/BetContract.sol";
import "../src/esportOracle.sol";
import "../src/esportOracleTypes.sol";
import "../test/BetContractTest.t.sol"; // Pour MockERC20

contract TestOracleIntegration is Script {
    function run() external {
        // Adresses des contrats déployés
        address betContractAddress = vm.envAddress("BET_CONTRACT_ADDRESS");
        address tokenAddress = vm.envAddress("TOKEN_ADDRESS");
        address oracleAddress = vm.envAddress("ORACLE_ADDRESS");
        
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        
        vm.startBroadcast(deployerPrivateKey);
        
        BetContract betContract = BetContract(payable(betContractAddress));
        MockERC20 token = MockERC20(tokenAddress);
        EsportOracle oracle = EsportOracle(oracleAddress);
        
        console.log("=== TESTING ORACLE INTEGRATION ===");
        console.log("BetContract:", betContractAddress);
        console.log("Oracle:", oracleAddress);
        console.log("Deployer:", deployer);
        
        // Test 1: Créer un pari et attendre qu'il expire
        console.log("\n1. Creating a bet with short deadline...");
        uint256 deadline = block.timestamp + 60; // 1 minute
        uint256 team1Id = 100; // Team C
        uint256 team2Id = 200; // Team D
        uint256 matchId = 999;
        
        betContract.createBet(
            "Team C vs Team D - Test Match",
            team1Id,
            team2Id,
            deadline,
            matchId
        );
        
        console.log("Bet created with matchId:", matchId);
        
        // Test 2: Placer des paris avec différents comptes
        console.log("\n2. Placing bets from multiple users...");
        uint256 betAmount = 50 * 10**18; // 50 tokens
        
        // Parier avec le deployer (team 1)
        token.approve(address(betContract), betAmount);
        betContract.placeBet(1, 1, betAmount);
        console.log("Deployer bet 50 tokens on team 1");
        
        // Mint des tokens pour un autre compte et parier sur team 2
        address user2 = vm.addr(2); // Utilise la clé privée #2 d'Anvil
        token.mint(user2, betAmount);
        
        // Approuver et parier avec user2
        vm.stopBroadcast();
        vm.startBroadcast(2); // Utilise la clé privée #2
        token.approve(address(betContract), betAmount);
        betContract.placeBet(1, 2, betAmount);
        console.log("User2 bet 50 tokens on team 2");
        
        vm.stopBroadcast();
        vm.startBroadcast(deployerPrivateKey); // Retour au deployer
        
        // Test 3: Avancer le temps pour dépasser la deadline
        console.log("\n3. Fast-forwarding time past deadline...");
        vm.warp(deadline + 100); // Dépasser la deadline
        console.log("Time advanced past deadline");
        
        // Test 4: Simuler une résolution de pari par l'oracle
        console.log("\n4. Simulating oracle resolution...");
        
        // Créer les données du match (team1 gagne)
        EsportOracleTypes.Match memory matchData;
        matchData._id = matchId;
        matchData._winnerId = team1Id; // Team C gagne
        
        // Important: On doit simuler que l'oracle appelle resolveBet
        // Pour cela on va utiliser le même compte que celui qui a déployé l'oracle
        // car dans le déploiement, le deployer devient owner de l'oracle
        
        console.log("Match data prepared:");
        console.log("- Match ID:", matchData._id);
        console.log("- Winner ID:", matchData._winnerId);
        
        // Vérifier l'état avant résolution
        BetContract.Bet memory bet = betContract.getBet(1);
        console.log("Before resolution:");
        console.log("- Resolved:", bet.resolved);
        console.log("- Team1 Pool:", bet.team1Pool / 10**18);
        console.log("- Team2 Pool:", bet.team2Pool / 10**18);
        
        // L'oracle résout le pari
        // Note: Normalement c'est l'oracle qui appelle cette fonction
        // mais pour le test on simule l'appel
        try betContract.resolveBet(1, matchData) {
            console.log("Bet resolved successfully!");
        } catch Error(string memory reason) {
            console.log("Resolution failed:", reason);
        } catch {
            console.log("Resolution failed with unknown error");
        }
        
        // Vérifier l'état après résolution
        bet = betContract.getBet(1);
        console.log("After resolution:");
        console.log("- Resolved:", bet.resolved);
        console.log("- Winning Team:", bet.winningTeam);
        
        vm.stopBroadcast();
        
        console.log("\n=== ORACLE INTEGRATION TEST COMPLETED ===");
    }
}
