// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/esportOracle.sol";

contract EsportOracleTest is Test {
    EsportOracle public oracle;
    address public owner;
    address public user1;
    address public user2;
    address public user3;
    address public user4;

    event newNodeAdded(address indexed addressAdded);
    event stakingSuccess(address indexed addressAdded, uint256 amount);
    event NodePunished(address indexed node, uint256 amount, uint256 violationsCount);
    event NodeBanned(address indexed node);

    function setUp() public {
        owner = address(this);
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");
        user3 = makeAddr("user3");
        user4 = makeAddr("user4");

        vm.deal(user1, 1 ether);
        vm.deal(user2, 1 ether);
        vm.deal(user3, 1 ether);
        vm.deal(user4, 1 ether);

        vm.prank(owner);
        oracle = new EsportOracle();
    }

    function testOwnership() public {
        assertEq(oracle._owner(), owner, "Le proprietaire initial doit etre l'adresse du deploiement");

        address newOwner = makeAddr("newOwner");
        oracle.setOwner(newOwner);
        assertEq(oracle._owner(), newOwner, "Le proprietaire doit etre mis a jour");

        vm.prank(user1);
        vm.expectRevert("Not the contract owner");
        oracle.setOwner(user1);
    }

    function testAddFundToStaking() public {
        vm.prank(user1);
        vm.expectEmit(true, true, false, false);
        emit stakingSuccess(user1, 0.001 ether);
        emit newNodeAdded(user1);
        oracle.addFundToStaking{value: 0.001 ether}();

        address[] memory nodes = oracle.getListedNodes();
        assertEq(nodes.length, 1, "Un seul noeud doit etre ajoute");
        assertEq(nodes[0], user1, "L'adresse du noeud doit correspondre");
        assertEq(oracle._fundsStaked(user1), 0.001 ether, "Le montant stake doit etre 0.001 ether");
    }

    function testInvalidStakingAmount() public {
        vm.prank(user1);
        vm.expectRevert("amount must be exactly 0.001 ether");
        oracle.addFundToStaking{value: 0.002 ether}();
    }

    function testDoubleStaking() public {
        vm.prank(user1);
        oracle.addFundToStaking{value: 0.001 ether}();

        vm.prank(user1);
        vm.expectRevert("Already staked");
        oracle.addFundToStaking{value: 0.001 ether}();
    }

    function testMultipleStaking() public {
        vm.prank(user1);
        vm.expectEmit(true, true, false, false);
        emit stakingSuccess(user1, 0.001 ether);
        emit newNodeAdded(user1);
        oracle.addFundToStaking{value: 0.001 ether}();

        vm.prank(user2);
        vm.expectEmit(true, true, false, false);
        emit stakingSuccess(user2, 0.001 ether);
        emit newNodeAdded(user2);
        oracle.addFundToStaking{value: 0.001 ether}();

        address[] memory nodes = oracle.getListedNodes();
        assertEq(nodes.length, 2, "Deux noeuds seulement doivent etre ajoutes");
        assertEq(nodes[0], user1, "L'adresse du noeud doit correspondre");
        assertEq(nodes[1], user2, "L'adresse du noeud doit correspondre");
        assertEq(oracle._fundsStaked(user1), 0.001 ether, "Le montant stake doit etre 0.001 ether");
        assertEq(oracle._fundsStaked(user2), 0.001 ether, "Le montant stake doit etre 0.001 ether");
    }

    function testQuorumWithEnoughNode() public {
        EsportOracle.Opponents[] memory opponents = new EsportOracle.Opponents[](2);
        opponents[0] = EsportOracle.Opponents({
            _acronym: "TA",
            _id: 1,
            _name: "Team A"
        });
        opponents[1] = EsportOracle.Opponents({
            _acronym: "TB",
            _id: 2,
            _name: "Team B"
        });
        EsportOracle.Games[] memory games = new EsportOracle.Games[](1);
        games[0] = EsportOracle.Games({
            _id: 1,
            _finished: true,
            _winnerId: 1
        });
        EsportOracle.Result[] memory results = new EsportOracle.Result[](2);
        results[0] = EsportOracle.Result({
            _score: 3,
            _teamId: 1
        });
        results[1] = EsportOracle.Result({
            _score: 1,
            _teamId: 2
        });
        EsportOracle.Match[] memory matches = new EsportOracle.Match[](1);
        matches[0] = EsportOracle.Match({
            _id: 1,
            _opponents: opponents,
            _game: games,
            _result: results,
            _winnerId: 1,
            _beginAt: block.timestamp
        });

        vm.prank(user1);
        oracle.addFundToStaking{value: 0.001 ether}();

        vm.prank(user2);
        oracle.addFundToStaking{value: 0.001 ether}();

        vm.prank(user3);
        oracle.addFundToStaking{value: 0.001 ether}();

        vm.prank(user4);
        oracle.addFundToStaking{value: 0.001 ether}();

        vm.prank(user1);
        oracle.handleNewMatches(matches);

        vm.prank(user2);
        oracle.handleNewMatches(matches);

        vm.prank(user3);
        oracle.handleNewMatches(matches);

        EsportOracle.Match memory dataNode = oracle.getMatchById(1);
        assertEq(dataNode._id, 1, "L'ID du match doit correspondre");

        assertEq(oracle.getPendingMatches().length, 0, "le nombre de match doit etre de 0");
    }

    function testQuorumWithEnoughSameMatch() public {
        EsportOracle.Opponents[] memory opponents = new EsportOracle.Opponents[](2);
        opponents[0] = EsportOracle.Opponents({
            _acronym: "TA",
            _id: 1,
            _name: "Team A"
        });
        opponents[1] = EsportOracle.Opponents({
            _acronym: "TB",
            _id: 2,
            _name: "Team B"
        });
        EsportOracle.Games[] memory games = new EsportOracle.Games[](1);
        games[0] = EsportOracle.Games({
            _id: 1,
            _finished: true,
            _winnerId: 1
        });
        EsportOracle.Result[] memory results = new EsportOracle.Result[](2);
        results[0] = EsportOracle.Result({
            _score: 3,
            _teamId: 1
        });
        results[1] = EsportOracle.Result({
            _score: 1,
            _teamId: 2
        });
        EsportOracle.Match[] memory matches = new EsportOracle.Match[](1);
        matches[0] = EsportOracle.Match({
            _id: 1,
            _opponents: opponents,
            _game: games,
            _result: results,
            _winnerId: 1,
            _beginAt: block.timestamp
        });
        EsportOracle.Match[] memory matches2 = new EsportOracle.Match[](1);
        matches2[0] = EsportOracle.Match({
            _id: 2,
            _opponents: opponents,
            _game: games,
            _result: results,
            _winnerId: 2,
            _beginAt: block.timestamp
        });

        vm.prank(user1);
        oracle.addFundToStaking{value: 0.001 ether}();

        vm.prank(user2);
        oracle.addFundToStaking{value: 0.001 ether}();

        vm.prank(user3);
        oracle.addFundToStaking{value: 0.001 ether}();

        vm.prank(user4);
        oracle.addFundToStaking{value: 0.001 ether}();

        vm.prank(user1);
        oracle.handleNewMatches(matches);

        vm.prank(user2);
        oracle.handleNewMatches(matches);

        vm.prank(user3);
        oracle.handleNewMatches(matches);

        vm.prank(user4);
        oracle.handleNewMatches(matches2);

        EsportOracle.Match memory dataNode = oracle.getMatchById(1);
        assertEq(dataNode._id, 1, "L'ID du match doit correspondre");

        dataNode = oracle.getMatchById(2);
        assertEq(dataNode._id, 0, "L'ID du match doit etre 0 car pas assez de votes");

        assertEq(oracle.getPendingMatches().length, 1, "Le nombre de match en attente doit etre de 1");
    }

    function testUpdatingMatchAlreadyRegister() public {
        EsportOracle.Opponents[] memory opponents = new EsportOracle.Opponents[](2);
        opponents[0] = EsportOracle.Opponents({
            _acronym: "TA",
            _id: 1,
            _name: "Team A"
        });
        opponents[1] = EsportOracle.Opponents({
            _acronym: "TB",
            _id: 2,
            _name: "Team B"
        });
        EsportOracle.Games[] memory games = new EsportOracle.Games[](3);
        games[0] = EsportOracle.Games({
            _id: 1,
            _finished: false,
            _winnerId: 0
        });
        games[1] = EsportOracle.Games({
            _id: 2,
            _finished: false,
            _winnerId: 0
        });
        games[2] = EsportOracle.Games({
            _id: 3,
            _finished: false,
            _winnerId: 0
        });
        EsportOracle.Result[] memory results = new EsportOracle.Result[](2);
        results[0] = EsportOracle.Result({
            _score: 3,
            _teamId: 1
        });
        results[1] = EsportOracle.Result({
            _score: 1,
            _teamId: 2
        });
        EsportOracle.Match[] memory matches = new EsportOracle.Match[](1);
        matches[0] = EsportOracle.Match({
            _id: 1,
            _opponents: opponents,
            _game: games,
            _result: results,
            _winnerId: 1,
            _beginAt: block.timestamp
        });

        vm.prank(user1);
        oracle.addFundToStaking{value: 0.001 ether}();

        vm.prank(user2);
        oracle.addFundToStaking{value: 0.001 ether}();

        vm.prank(user3);
        oracle.addFundToStaking{value: 0.001 ether}();

        vm.prank(user4);
        oracle.addFundToStaking{value: 0.001 ether}();

        vm.prank(user1);
        oracle.handleNewMatches(matches);

        vm.prank(user2);
        oracle.handleNewMatches(matches);

        vm.prank(user3);
        oracle.handleNewMatches(matches);

        EsportOracle.Games[] memory games2 = new EsportOracle.Games[](3);
        games2[0] = EsportOracle.Games({
            _id: 1,
            _finished: true,
            _winnerId: 2
        });
        games2[1] = EsportOracle.Games({
            _id: 2,
            _finished: true,
            _winnerId: 2
        });
        games2[2] = EsportOracle.Games({
            _id: 3,
            _finished: true,
            _winnerId: 2
        });

        EsportOracle.Match[] memory matches2 = new EsportOracle.Match[](1);
        matches2[0] = EsportOracle.Match({
            _id: 1,
            _opponents: opponents,
            _game: games2,
            _result: results,
            _winnerId: 1,
            _beginAt: block.timestamp
        });

        vm.prank(user1);
        oracle.handleNewMatches(matches2);

        vm.prank(user2);
        oracle.handleNewMatches(matches2);

        vm.prank(user3);
        oracle.handleNewMatches(matches2);

        EsportOracle.Match memory dataNode = oracle.getMatchById(1);

        assertEq(dataNode._id, 1, "L'ID du match doit correspondre");
        assertEq(dataNode._game[0]._finished, true, "La valeur de finish doit etre true");

        assertEq(oracle.getPendingMatches().length, 0, "le nombre de match doit etre de 0");
    }

    // Test of the ban system after multiple violations
    function testBanAfterMultipleViolations() public {
        // Add 4 nodes
        vm.prank(user1);
        oracle.addFundToStaking{value: 0.001 ether}();

        vm.prank(user2);
        oracle.addFundToStaking{value: 0.001 ether}();

        vm.prank(user3);
        oracle.addFundToStaking{value: 0.001 ether}();

        vm.prank(user4);
        oracle.addFundToStaking{value: 0.001 ether}();

        // Prepare matches for testing
        EsportOracle.Opponents[] memory opponents = new EsportOracle.Opponents[](2);
        opponents[0] = EsportOracle.Opponents({_acronym: "TA", _id: 1, _name: "Team A"});
        opponents[1] = EsportOracle.Opponents({_acronym: "TB", _id: 2, _name: "Team B"});
        
        EsportOracle.Games[] memory games = new EsportOracle.Games[](1);
        games[0] = EsportOracle.Games({_id: 1, _finished: true, _winnerId: 1});
        
        EsportOracle.Result[] memory results = new EsportOracle.Result[](2);
        results[0] = EsportOracle.Result({_score: 3, _teamId: 1});
        results[1] = EsportOracle.Result({_score: 1, _teamId: 2});
        
        // Correct matches with different IDs
        EsportOracle.Match[] memory correctMatch1 = new EsportOracle.Match[](1);
        correctMatch1[0] = EsportOracle.Match({
            _id: 1, _opponents: opponents, _game: games, _result: results, _winnerId: 1, _beginAt: 1
        });
        
        EsportOracle.Match[] memory correctMatch2 = new EsportOracle.Match[](1);
        correctMatch2[0] = EsportOracle.Match({
            _id: 2, _opponents: opponents, _game: games, _result: results, _winnerId: 1, _beginAt: 1
        });
        
        EsportOracle.Match[] memory correctMatch3 = new EsportOracle.Match[](1);
        correctMatch3[0] = EsportOracle.Match({
            _id: 3, _opponents: opponents, _game: games, _result: results, _winnerId: 1, _beginAt: 1
        });
        
        // Incorrect matches
        EsportOracle.Games[] memory incorrectGames = new EsportOracle.Games[](1);
        incorrectGames[0] = EsportOracle.Games({_id: 1, _finished: false, _winnerId: 2});
        
        EsportOracle.Match[] memory incorrectMatch1 = new EsportOracle.Match[](1);
        incorrectMatch1[0] = EsportOracle.Match({
            _id: 1, _opponents: opponents, _game: incorrectGames, _result: results, _winnerId: 2, _beginAt: 1
        });
        
        EsportOracle.Match[] memory incorrectMatch2 = new EsportOracle.Match[](1);
        incorrectMatch2[0] = EsportOracle.Match({
            _id: 2, _opponents: opponents, _game: incorrectGames, _result: results, _winnerId: 2, _beginAt: 1
        });
        
        EsportOracle.Match[] memory incorrectMatch3 = new EsportOracle.Match[](1);
        incorrectMatch3[0] = EsportOracle.Match({
            _id: 3, _opponents: opponents, _game: incorrectGames, _result: results, _winnerId: 2, _beginAt: 1
        });
        
        // First cycle - First violation
        vm.prank(user4);
        oracle.handleNewMatches(incorrectMatch1);
        
        vm.prank(user1);
        oracle.handleNewMatches(correctMatch1);
        
        vm.prank(user2);
        oracle.handleNewMatches(correctMatch1);
        
        vm.prank(user3);
        oracle.handleNewMatches(correctMatch1);
        
        // Verify first violation
        (uint256 violations, bool banned) = oracle._nodeViolations(user4);
        assertEq(violations, 1, "After 1 violation, counter should be 1");
        assertEq(banned, false, "Node should not be banned yet");
        assertEq(oracle._fundsStaked(user4), 0.0009 ether, "Balance should be reduced after 1 violation");
        
        // Second cycle - Second violation
        vm.prank(user4);
        oracle.handleNewMatches(incorrectMatch2);
        
        vm.prank(user1);
        oracle.handleNewMatches(correctMatch2);
        
        vm.prank(user2);
        oracle.handleNewMatches(correctMatch2);
        
        vm.prank(user3);
        oracle.handleNewMatches(correctMatch2);
        
        // Verify second violation
        (violations, banned) = oracle._nodeViolations(user4);
        assertEq(violations, 2, "After 2 violations, counter should be 2");
        assertEq(banned, false, "Node should still not be banned");
        assertEq(oracle._fundsStaked(user4), 0.0008 ether, "Balance should be reduced after 2 violations");
        
        // Third cycle - Third violation and expected ban
        vm.prank(user4);
        oracle.handleNewMatches(incorrectMatch3);
        
        vm.prank(user1);
        oracle.handleNewMatches(correctMatch3);
        
        vm.prank(user2);
        oracle.handleNewMatches(correctMatch3);
        
        // Set up event capture
        vm.recordLogs();
        
        vm.prank(user3);
        oracle.handleNewMatches(correctMatch3);
        
        // Verify NodeBanned event was emitted
        Vm.Log[] memory entries = vm.getRecordedLogs();
        bool eventFound = false;
        
        for(uint i = 0; i < entries.length; i++) {
            if(entries[i].topics[0] == keccak256("NodeBanned(address)")) {
                address eventNode = address(uint160(uint256(entries[i].topics[1])));
                assertEq(eventNode, user4, "NodeBanned event should contain user4's address");
                eventFound = true;
                break;
            }
        }
        
        assertTrue(eventFound, "NodeBanned event must be emitted");
        
        // Verify ban status
        (violations, banned) = oracle._nodeViolations(user4);
        assertEq(violations, 3, "After 3 violations, counter should be 3");
        assertEq(banned, true, "Node should be banned after 3 violations");
        assertEq(oracle._fundsStaked(user4), 0, "Balance should be 0 after ban");
        
        // Verify banned node cannot submit matches
        vm.prank(user4);
        vm.expectRevert("Node is banned");
        oracle.handleNewMatches(correctMatch1);
    }
    
    // Test des fonctions administratives
    function testAdminFunctions() public {
        // Ajouter 2 noeuds
        vm.prank(user1);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        vm.prank(user2);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        // Test de la fonction banNode
        vm.expectEmit(true, false, false, false);
        emit NodeBanned(user1);

        oracle.banNode(user1);

        // Verifier le bannissement
        (uint256 violations, bool banned) = oracle._nodeViolations(user1);
        assertEq(banned, true, "Le noeud doit etre banni");
        assertEq(oracle._fundsStaked(user1), 0, "Les fonds doivent etre confisques");
        
        // Verifier que le noeud banni ne peut plus soumettre de matches
        vm.prank(user1);
        vm.expectRevert("Node is banned");
        oracle.handleNewMatches(new EsportOracle.Match[](1));
        
        // Test de la fonction rehabilitateNode
        oracle.rehabilitateNode(user1);
        
        // Verifier la rehabilitation
        (violations, banned) = oracle._nodeViolations(user1);
        assertEq(banned, false, "Le noeud ne doit plus etre banni apres rehabilitation");
        assertEq(violations, 0, "Le compteur de violations doit etre remis a zero");
        
        // Le noeud rehabilite doit refaire un staking pour participer
        vm.prank(user1);
        vm.expectRevert("Node is not listed, please call addNewNode function to register a new node");
        oracle.handleNewMatches(new EsportOracle.Match[](1));
    }
    
    // Test de la fonction de retrait des fonds
    function testWithdrawStake() public {
        // Ajouter un noeud
        vm.prank(user1);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        // Verifier le solde initial
        assertEq(oracle._fundsStaked(user1), 0.001 ether, "Le solde initial doit etre 0.001 ether");
        
        // Enregistrer le solde ETH avant le retrait
        uint256 balanceBefore = user1.balance;
        
        // Retirer les fonds
        vm.prank(user1);
        oracle.withdrawStake();
        
        // Verifier que les fonds ont ete retires
        assertEq(oracle._fundsStaked(user1), 0, "Le solde doit etre a 0 apres retrait");
        assertEq(user1.balance, balanceBefore + 0.001 ether, "L'ETH doit etre retourne a l'utilisateur");
        
        // Verifier que le noeud a ete retire de la liste
        address[] memory nodes = oracle.getListedNodes();
        assertEq(nodes.length, 0, "La liste des noeuds doit etre vide apres retrait");
    }
    
    // Test de la redistribution des fonds apres bannissement
    function testFundsRedistribution() public {
        // Ajouter 4 noeuds
        vm.prank(user1);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        vm.prank(user2);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        vm.prank(user3);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        vm.prank(user4);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        // Bannir un noeud directement
        oracle.banNode(user1);
        
        // Verifier que les fonds ont ete redistribues aux autres noeuds
        uint256 expectedShare = uint256(0.001 ether) / 3;  // Divise par 3 noeuds restants
        
        assertEq(oracle._fundsStaked(user1), 0, "Le noeud banni doit avoir un solde de 0");
        assertEq(oracle._fundsStaked(user2), 0.001 ether + expectedShare, "Le noeud doit recevoir sa part");
        assertEq(oracle._fundsStaked(user3), 0.001 ether + expectedShare, "Le noeud doit recevoir sa part");
        assertEq(oracle._fundsStaked(user4), 0.001 ether + expectedShare, "Le noeud doit recevoir sa part");
    }
}
