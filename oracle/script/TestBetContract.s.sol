// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import "../src/BetContract.sol";
import "../test/BetContractTest.t.sol"; // Pour MockERC20

contract TestBetContract is Script {
    function run() external {
        // Adresses des contrats déployés (à mettre à jour après déploiement)
        address betContractAddress = vm.envAddress("BET_CONTRACT_ADDRESS");
        address tokenAddress = vm.envAddress("TOKEN_ADDRESS");
        
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        
        vm.startBroadcast(deployerPrivateKey);
        
        BetContract betContract = BetContract(payable(betContractAddress));
        MockERC20 token = MockERC20(tokenAddress);
        
        console.log("=== TESTING BET CONTRACT ===");
        console.log("BetContract:", betContractAddress);
        console.log("Token:", tokenAddress);
        console.log("Deployer:", deployer);
        
        // Test 1: Créer un pari
        console.log("\n1. Creating a bet...");
        uint256 deadline = block.timestamp + 1 hours;
        uint256 team1Id = 123; // Team A
        uint256 team2Id = 456; // Team B
        uint256 matchId = 789;
        
        betContract.createBet(
            "Team A vs Team B - LoL World Championship",
            team1Id,
            team2Id,
            deadline,
            matchId
        );
        
        uint256 betCount = betContract.getBetCount();
        console.log("Bet created! Total bets:", betCount);
        
        // Test 2: Vérifier le pari créé
        BetContract.Bet memory bet = betContract.getBet(0);
        console.log("Bet description:", bet.description);
        console.log("Team1 ID:", bet.team1Id);
        console.log("Team2 ID:", bet.team2Id);
        console.log("Match ID:", bet.matchId);
        console.log("Deadline:", bet.deadline);
        
        // Test 3: Approuver et placer des paris
        console.log("\n2. Placing bets...");
        uint256 betAmount = 100 * 10**18; // 100 tokens
        
        // Approuver le contrat pour dépenser nos tokens
        token.approve(address(betContract), betAmount * 2);
        
        // Parier sur l'équipe 1
        betContract.placeBet(0, 1, betAmount);
        console.log("Placed bet on team 1:", betAmount / 10**18, "tokens");
        
        // Vérifier le pari de l'utilisateur
        BetContract.UserBet memory userBet = betContract.getUserBet(deployer, 0);
        console.log("User bet amount:", userBet.amount / 10**18);
        console.log("User bet team:", userBet.teamChosen);
        
        // Vérifier les pools
        bet = betContract.getBet(0);
        console.log("Team1 pool:", bet.team1Pool / 10**18, "tokens");
        console.log("Team2 pool:", bet.team2Pool / 10**18, "tokens");
        
        // Test 4: Vérifier les fonctions utilitaires
        console.log("\n3. Testing utility functions...");
        address[] memory participants = betContract.getBetParticipants(0);
        console.log("Number of participants:", participants.length);
        console.log("First participant:", participants[0]);
        
        uint256 contractETHBalance = betContract.getContractETHBalance();
        console.log("Contract ETH balance:", contractETHBalance);
        
        uint256 matchRequestFee = betContract.matchRequestFee();
        console.log("Match request fee:", matchRequestFee);
        
        vm.stopBroadcast();
        
        console.log("\n=== TEST COMPLETED ===");
        console.log("All basic functions are working!");
        console.log("Ready for oracle integration testing.");
    }
}
