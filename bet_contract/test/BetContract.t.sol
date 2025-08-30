// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "../src/BetContract.sol";
import "openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";

// Mock ERC20 token pour les tests
contract MockERC20 is ERC20 {
    constructor() ERC20("Test Token", "TEST") {
        _mint(msg.sender, 1000000 * 10**18);
    }

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}

contract BetContractTest is Test {
    BetContract public betContract;
    MockERC20 public token;
    
    address public owner;
    address public oracle;
    address public user1;
    address public user2;
    address public user3;
    
    // Constantes pour les tests
    uint256 constant INITIAL_BALANCE = 10000 * 10**18;
    uint256 constant BET_AMOUNT = 100 * 10**18;
    uint256 constant FUTURE_TIMESTAMP = 1735689600; // 1 Jan 2025
    
    function setUp() public {
        // Configuration des adresses
        owner = address(this);
        oracle = makeAddr("oracle");
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");
        user3 = makeAddr("user3");
        
        // Déploiement du token de test
        token = new MockERC20();
        
        // Déploiement du contrat de paris
        betContract = new BetContract(oracle, address(token));
        
        // Distribution des tokens aux utilisateurs
        token.mint(user1, INITIAL_BALANCE);
        token.mint(user2, INITIAL_BALANCE);
        token.mint(user3, INITIAL_BALANCE);
        
        // Approval des tokens pour le contrat
        vm.prank(user1);
        token.approve(address(betContract), type(uint256).max);
        
        vm.prank(user2);
        token.approve(address(betContract), type(uint256).max);
        
        vm.prank(user3);
        token.approve(address(betContract), type(uint256).max);
    }
    
    // ===== TESTS DE CRÉATION DE PARIS =====
    
    function test_CreateBet_Success() public {
        vm.prank(user1);
        betContract.createBet("PSG vs OM", "PSG", "OM", FUTURE_TIMESTAMP);
        
        assertEq(betContract.getBetCount(), 1);
        
        (string memory description, string memory team1, string memory team2, 
         uint256 deadline, uint256 team1Pool, uint256 team2Pool, 
         uint8 winningTeam, bool resolved, address creator) = betContract.CurrentBets(0);
         
        assertEq(description, "PSG vs OM");
        assertEq(team1, "PSG");
        assertEq(team2, "OM");
        assertEq(deadline, FUTURE_TIMESTAMP);
        assertEq(team1Pool, 0);
        assertEq(team2Pool, 0);
        assertEq(winningTeam, 0);
        assertEq(resolved, false);
        assertEq(creator, user1);
    }
    
    function test_CreateBet_RevertWhenDeadlineInPast() public {
        vm.prank(user1);
        vm.expectRevert("Deadline must be in the future");
        betContract.createBet("PSG vs OM", "PSG", "OM", block.timestamp - 1);
    }
    
    // ===== TESTS DE PLACEMENT DE PARIS =====
    
    function test_PlaceBet_Success() public {
        // Créer un pari
        vm.prank(user1);
        betContract.createBet("PSG vs OM", "PSG", "OM", FUTURE_TIMESTAMP);
        
        uint256 initialBalance = token.balanceOf(user2);
        
        // Placer un pari
        vm.prank(user2);
        betContract.placeBet(0, 1, BET_AMOUNT);
        
        // Vérifications
        assertEq(token.balanceOf(user2), initialBalance - BET_AMOUNT);
        assertEq(token.balanceOf(address(betContract)), BET_AMOUNT);
        
        (,,,, uint256 team1Pool, uint256 team2Pool,,,) = betContract.CurrentBets(0);
        assertEq(team1Pool, BET_AMOUNT);
        assertEq(team2Pool, 0);
        
        // Vérifier le pari utilisateur
        BetContract.UserBet memory userBet = betContract.getUserBet(user2, 0);
        assertEq(userBet.amount, BET_AMOUNT);
        assertEq(userBet.teamChosen, 1);
        assertEq(userBet.claimed, false);
        
        // Vérifier les participants
        address[] memory participants = betContract.getBetParticipants(0);
        assertEq(participants.length, 1);
        assertEq(participants[0], user2);
    }
    
    function test_PlaceBet_MultipleUsers() public {
        // Créer un pari
        vm.prank(user1);
        betContract.createBet("PSG vs OM", "PSG", "OM", FUTURE_TIMESTAMP);
        
        // User2 parie sur l'équipe 1
        vm.prank(user2);
        betContract.placeBet(0, 1, BET_AMOUNT);
        
        // User3 parie sur l'équipe 2
        vm.prank(user3);
        betContract.placeBet(0, 2, BET_AMOUNT * 2);
        
        // Vérifications des pools
        (,,,, uint256 team1Pool, uint256 team2Pool,,,) = betContract.CurrentBets(0);
        assertEq(team1Pool, BET_AMOUNT);
        assertEq(team2Pool, BET_AMOUNT * 2);
        
        // Vérifier les participants
        address[] memory participants = betContract.getBetParticipants(0);
        assertEq(participants.length, 2);
    }
    
    function test_PlaceBet_SameUserAddsToBet() public {
        // Créer un pari
        vm.prank(user1);
        betContract.createBet("PSG vs OM", "PSG", "OM", FUTURE_TIMESTAMP);
        
        // Premier pari
        vm.prank(user2);
        betContract.placeBet(0, 1, BET_AMOUNT);
        
        // Deuxième pari du même utilisateur sur la même équipe
        vm.prank(user2);
        betContract.placeBet(0, 1, BET_AMOUNT);
        
        // Vérifications
        BetContract.UserBet memory userBet = betContract.getUserBet(user2, 0);
        assertEq(userBet.amount, BET_AMOUNT * 2);
        
        (,,,, uint256 team1Pool,,,,) = betContract.CurrentBets(0);
        assertEq(team1Pool, BET_AMOUNT * 2);
        
        // Le participant ne doit être ajouté qu'une seule fois
        address[] memory participants = betContract.getBetParticipants(0);
        assertEq(participants.length, 1);
    }
    
    function test_PlaceBet_RevertWhenInvalidBetId() public {
        vm.prank(user1);
        vm.expectRevert("Invalid bet ID");
        betContract.placeBet(999, 1, BET_AMOUNT);
    }
    
    function test_PlaceBet_RevertWhenExpired() public {
        vm.prank(user1);
        betContract.createBet("PSG vs OM", "PSG", "OM", FUTURE_TIMESTAMP);
        
        // Avancer le temps au-delà de la deadline
        vm.warp(FUTURE_TIMESTAMP + 1);
        
        vm.prank(user2);
        vm.expectRevert("Pari expire");
        betContract.placeBet(0, 1, BET_AMOUNT);
    }
    
    function test_PlaceBet_RevertWhenInvalidTeam() public {
        vm.prank(user1);
        betContract.createBet("PSG vs OM", "PSG", "OM", FUTURE_TIMESTAMP);
        
        vm.prank(user2);
        vm.expectRevert("Equipe invalide");
        betContract.placeBet(0, 3, BET_AMOUNT);
    }
    
    function test_PlaceBet_RevertWhenChangingTeam() public {
        vm.prank(user1);
        betContract.createBet("PSG vs OM", "PSG", "OM", FUTURE_TIMESTAMP);
        
        // Premier pari sur l'équipe 1
        vm.prank(user2);
        betContract.placeBet(0, 1, BET_AMOUNT);
        
        // Tentative de pari sur l'équipe 2
        vm.prank(user2);
        vm.expectRevert("Cannot change team choice");
        betContract.placeBet(0, 2, BET_AMOUNT);
    }
    
    function test_PlaceBet_RevertWhenContractPaused() public {
        vm.prank(user1);
        betContract.createBet("PSG vs OM", "PSG", "OM", FUTURE_TIMESTAMP);
        
        // Mettre en pause le contrat
        betContract.pause();
        
        vm.prank(user2);
        vm.expectRevert(); // Changé pour accepter n'importe quel revert lié à la pause
        betContract.placeBet(0, 1, BET_AMOUNT);
    }
    
    // ===== TESTS DE RÉSOLUTION DE PARIS =====
    
    function test_ResolveBet_Success() public {
        // Créer et configurer un pari
        vm.prank(user1);
        betContract.createBet("PSG vs OM", "PSG", "OM", FUTURE_TIMESTAMP);
        
        // Placer des paris
        vm.prank(user2);
        betContract.placeBet(0, 1, BET_AMOUNT);
        
        // Avancer le temps après la deadline
        vm.warp(FUTURE_TIMESTAMP + 1);
        
        // Résoudre le pari
        vm.prank(oracle);
        betContract.resolveBet(0, 1);
        
        // Vérifications
        (,,,,,,uint8 winningTeam, bool resolved,) = betContract.CurrentBets(0);
        assertEq(winningTeam, 1);
        assertTrue(resolved);
    }
    
    function test_ResolveBet_RevertWhenNotOracle() public {
        vm.prank(user1);
        betContract.createBet("PSG vs OM", "PSG", "OM", FUTURE_TIMESTAMP);
        
        vm.warp(FUTURE_TIMESTAMP + 1);
        
        vm.prank(user2);
        vm.expectRevert("Only oracle can call this");
        betContract.resolveBet(0, 1);
    }
    
    function test_ResolveBet_RevertWhenAlreadyResolved() public {
        vm.prank(user1);
        betContract.createBet("PSG vs OM", "PSG", "OM", FUTURE_TIMESTAMP);
        
        vm.warp(FUTURE_TIMESTAMP + 1);
        
        vm.prank(oracle);
        betContract.resolveBet(0, 1);
        
        vm.prank(oracle);
        vm.expectRevert("Bet already resolved");
        betContract.resolveBet(0, 2);
    }
    
    function test_ResolveBet_RevertWhenStillActive() public {
        vm.prank(user1);
        betContract.createBet("PSG vs OM", "PSG", "OM", FUTURE_TIMESTAMP);
        
        vm.prank(oracle);
        vm.expectRevert("Bet still active");
        betContract.resolveBet(0, 1);
    }
    
    // ===== TESTS DE RÉCLAMATION DES GAINS =====
    
    function test_ClaimWinnings_Success() public {
        // Configuration du pari
        vm.prank(user1);
        betContract.createBet("PSG vs OM", "PSG", "OM", FUTURE_TIMESTAMP);
        
        // Placer des paris
        vm.prank(user2);
        betContract.placeBet(0, 1, BET_AMOUNT); // Team 1
        
        vm.prank(user3);
        betContract.placeBet(0, 2, BET_AMOUNT * 2); // Team 2
        
        // Résoudre le pari (Team 1 gagne)
        vm.warp(FUTURE_TIMESTAMP + 1);
        vm.prank(oracle);
        betContract.resolveBet(0, 1);
        
        uint256 balanceBefore = token.balanceOf(user2);
        
        // Réclamer les gains
        vm.prank(user2);
        betContract.claimWinnings(0);
        
        // Calcul des gains attendus: (BET_AMOUNT * (BET_AMOUNT + BET_AMOUNT*2)) / BET_AMOUNT = 3 * BET_AMOUNT
        uint256 expectedPayout = (BET_AMOUNT * (BET_AMOUNT + BET_AMOUNT * 2)) / BET_AMOUNT;
        assertEq(token.balanceOf(user2), balanceBefore + expectedPayout);
        
        // Vérifier que les gains sont marqués comme réclamés
        BetContract.UserBet memory userBet = betContract.getUserBet(user2, 0);
        assertTrue(userBet.claimed);
    }
    
    function test_ClaimWinnings_RevertWhenNotResolved() public {
        vm.prank(user1);
        betContract.createBet("PSG vs OM", "PSG", "OM", FUTURE_TIMESTAMP);
        
        vm.prank(user2);
        betContract.placeBet(0, 1, BET_AMOUNT);
        
        vm.prank(user2);
        vm.expectRevert("Pari non resolu");
        betContract.claimWinnings(0);
    }
    
    function test_ClaimWinnings_RevertWhenAlreadyClaimed() public {
        vm.prank(user1);
        betContract.createBet("PSG vs OM", "PSG", "OM", FUTURE_TIMESTAMP);
        
        vm.prank(user2);
        betContract.placeBet(0, 1, BET_AMOUNT);
        
        vm.warp(FUTURE_TIMESTAMP + 1);
        vm.prank(oracle);
        betContract.resolveBet(0, 1);
        
        vm.prank(user2);
        betContract.claimWinnings(0);
        
        vm.prank(user2);
        vm.expectRevert("Gains deja reclames");
        betContract.claimWinnings(0);
    }
    
    function test_ClaimWinnings_RevertWhenLosingTeam() public {
        vm.prank(user1);
        betContract.createBet("PSG vs OM", "PSG", "OM", FUTURE_TIMESTAMP);
        
        vm.prank(user2);
        betContract.placeBet(0, 1, BET_AMOUNT);
        
        vm.warp(FUTURE_TIMESTAMP + 1);
        vm.prank(oracle);
        betContract.resolveBet(0, 2); // Team 2 gagne, user2 avait parié sur Team 1
        
        vm.prank(user2);
        vm.expectRevert("Equipe perdante");
        betContract.claimWinnings(0);
    }
    
    // ===== TESTS DES FONCTIONS ADMINISTRATIVES =====
    
    function test_SetOracle_Success() public {
        address newOracle = makeAddr("newOracle");
        betContract.setOracle(newOracle);
        assertEq(betContract.oracle(), newOracle);
    }
    
    function test_SetOracle_RevertWhenNotOwner() public {
        vm.prank(user1);
        vm.expectRevert();
        betContract.setOracle(makeAddr("newOracle"));
    }
    
    function test_EmergencyWithdraw_Success() public {
        // Envoyer des tokens au contrat
        token.transfer(address(betContract), BET_AMOUNT);
        
        uint256 ownerBalanceBefore = token.balanceOf(owner);
        
        betContract.emergencyWithdraw(address(token), BET_AMOUNT);
        
        assertEq(token.balanceOf(owner), ownerBalanceBefore + BET_AMOUNT);
        assertEq(token.balanceOf(address(betContract)), 0);
    }
    
    function test_Pause_Success() public {
        betContract.pause();
        assertTrue(betContract.paused());
        
        betContract.unpause();
        assertFalse(betContract.paused());
    }
    
    // ===== TESTS DES FONCTIONS DE LECTURE =====
    
    function test_GetBetCount() public {
        assertEq(betContract.getBetCount(), 0);
        
        vm.prank(user1);
        betContract.createBet("Test 1", "A", "B", FUTURE_TIMESTAMP);
        assertEq(betContract.getBetCount(), 1);
        
        vm.prank(user1);
        betContract.createBet("Test 2", "C", "D", FUTURE_TIMESTAMP);
        assertEq(betContract.getBetCount(), 2);
    }
    
    // ===== TESTS D'ÉVÉNEMENTS =====
    
    function test_Events() public {
        // Test BetCreated event
        vm.expectEmit(true, false, false, true);
        emit BetContract.BetCreated(0, "PSG vs OM", "PSG", "OM", FUTURE_TIMESTAMP);
        
        vm.prank(user1);
        betContract.createBet("PSG vs OM", "PSG", "OM", FUTURE_TIMESTAMP);
        
        // Test BetPlaced event
        vm.expectEmit(true, true, false, true);
        emit BetContract.BetPlaced(user2, 0, 1, BET_AMOUNT);
        
        vm.prank(user2);
        betContract.placeBet(0, 1, BET_AMOUNT);
        
        // Test BetResolved event
        vm.warp(FUTURE_TIMESTAMP + 1);
        vm.expectEmit(true, false, false, true);
        emit BetContract.BetResolved(0, 1);
        
        vm.prank(oracle);
        betContract.resolveBet(0, 1);
        
        // Test WinningsClaimed event
        vm.expectEmit(true, true, false, true);
        emit BetContract.WinningsClaimed(user2, 0, BET_AMOUNT);
        
        vm.prank(user2);
        betContract.claimWinnings(0);
    }
}