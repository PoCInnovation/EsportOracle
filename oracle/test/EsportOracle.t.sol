// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/esportOracle.sol";
import "../src/esportOracleTypes.sol";

contract EsportOracleTest is Test {
    using EsportOracleTypes for EsportOracleTypes.MatchRequest;
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
        user1 = makeAddr("Alice");
        user2 = makeAddr("Bob");
        user3 = makeAddr("Charlie");
        user4 = makeAddr("Dave");

        vm.deal(user1, 1 ether);
        vm.deal(user2, 1 ether);
        vm.deal(user3, 1 ether);
        vm.deal(user4, 1 ether);

        vm.prank(owner);
        oracle = new EsportOracle();
    }

    function testOwnership() public {
        assertEq(oracle._owner(), owner, "Le proprietaire initial doit etre l'adresse du deploiement");

        address newOwner = makeAddr("NouveauProprietaire");
        oracle.setOwner(newOwner);
        assertEq(oracle._owner(), newOwner, "Le proprietaire doit etre mis a jour");

        vm.prank(user1);
        vm.expectRevert("Not the contract owner");
        oracle.setOwner(user1);
    }

    function testPauseFunctionality() public {
        console.log("Test de la mise en pause du contrat");
        
        vm.prank(user1);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        console.log("Mise en pause du contrat par le proprietaire");
        oracle.pause();
        
        console.log("Verification qu'un nouvel utilisateur ne peut pas faire de staking");
        vm.startPrank(user2);
        bool success = false;
        try oracle.addFundToStaking{value: 0.001 ether}() {
            success = true;
        } catch {}
        vm.stopPrank();
        assertFalse(success, "Le staking devrait echouer quand le contrat est en pause");
        
        console.log("Verification qu'un utilisateur ne peut pas retirer ses fonds");
        vm.startPrank(user1);
        success = false;
        try oracle.withdrawStake() {
            success = true;
        } catch {}
        vm.stopPrank();
        assertFalse(success, "Le retrait devrait echouer quand le contrat est en pause");
        
        console.log("Verification qu'un noeud ne peut pas envoyer de donnees de match");
        EsportOracleTypes.Match[] memory matchData = prepareSampleMatch();
        
        vm.startPrank(user1);
        success = false;
        try oracle.handleNewMatches(matchData) {
            success = true;
        } catch {}
        vm.stopPrank();
        assertFalse(success, "L'envoi de donnees devrait echouer quand le contrat est en pause");
        
        console.log("Verification qu'un utilisateur normal ne peut pas desactiver la pause");
        vm.prank(user1);
        vm.expectRevert("Not the contract owner");
        oracle.unpause();
        
        console.log("Desactivation de la pause par le proprietaire");
        oracle.unpause();
        
        console.log("Verification que le staking fonctionne apres desactivation de la pause");
        vm.prank(user2);
        oracle.addFundToStaking{value: 0.001 ether}();
        assertEq(oracle._fundsStaked(user2), 0.001 ether, "Le staking doit fonctionner apres desactivation de la pause");
        
        console.log("Verification que l'envoi de donnees fonctionne apres desactivation de la pause");
        vm.prank(user1);
        oracle.handleNewMatches(matchData);
        
        console.log("Verification que le retrait fonctionne apres desactivation de la pause");
        uint256 balanceBefore = user2.balance;
        vm.prank(user2);
        oracle.withdrawStake();
        assertEq(user2.balance, balanceBefore + 0.001 ether, "Le retrait doit fonctionner apres desactivation de la pause");
    }
    
    function testPausePermissions() public {
        console.log("Test des permissions de mise en pause");
        
        console.log("Verification qu'un utilisateur normal ne peut pas mettre en pause");
        vm.prank(user1);
        vm.expectRevert("Not the contract owner");
        oracle.pause();
        
        console.log("Le proprietaire met en pause");
        oracle.pause();
        
        console.log("Verification qu'un utilisateur normal ne peut pas desactiver la pause");
        vm.prank(user1);
        vm.expectRevert("Not the contract owner");
        oracle.unpause();
        
        console.log("Le proprietaire desactive la pause");
        oracle.unpause();
        
        console.log("Transfert de propriete et verification des nouvelles permissions");
        oracle.setOwner(user3);
        
        console.log("Verification que l'ancien proprietaire ne peut plus mettre en pause");
        vm.expectRevert("Not the contract owner");
        oracle.pause();
        
        console.log("Le nouveau proprietaire met en pause");
        vm.prank(user3);
        oracle.pause();
        
        console.log("Le nouveau proprietaire desactive la pause");
        vm.prank(user3);
        oracle.unpause();
    }

    function testAddFundToStaking() public {
        console.log("Test d'ajout de fonds au staking");
        
        vm.prank(user1);
        vm.expectEmit(true, true, false, false);
        emit stakingSuccess(user1, 0.001 ether);
        emit newNodeAdded(user1);
        oracle.addFundToStaking{value: 0.001 ether}();

        address[] memory nodes = oracle.getListedNodes();
        assertEq(nodes.length, 1, "Un seul noeud doit etre ajoute");
        assertEq(nodes[0], user1, "L'adresse du noeud doit correspondre");
        assertEq(oracle._fundsStaked(user1), 0.001 ether, "Le montant stake doit etre 0.001 ether");
        
        console.log("Staking reussi pour", vm.toString(user1));
    }

    function testInvalidStakingAmount() public {
        console.log("Test de staking avec un montant invalide");
        
        vm.prank(user1);
        vm.expectRevert("amount must be exactly 0.001 ether");
        oracle.addFundToStaking{value: 0.002 ether}();
        
        console.log("Rejet correct du montant invalide");
    }

    function testDoubleStaking() public {
        console.log("Test de double staking par le meme utilisateur");
        
        vm.prank(user1);
        oracle.addFundToStaking{value: 0.001 ether}();
        console.log("Premier staking reussi pour", vm.toString(user1));

        vm.prank(user1);
        vm.expectRevert("Already staked");
        oracle.addFundToStaking{value: 0.001 ether}();
        
        console.log("Deuxieme staking correctement rejete");
    }

    function testMultipleStaking() public {
        console.log("Test de staking par plusieurs utilisateurs");
        
        vm.prank(user1);
        vm.expectEmit(true, true, false, false);
        emit stakingSuccess(user1, 0.001 ether);
        emit newNodeAdded(user1);
        oracle.addFundToStaking{value: 0.001 ether}();
        console.log("Staking reussi pour Alice");

        vm.prank(user2);
        vm.expectEmit(true, true, false, false);
        emit stakingSuccess(user2, 0.001 ether);
        emit newNodeAdded(user2);
        oracle.addFundToStaking{value: 0.001 ether}();
        console.log("Staking reussi pour Bob");

        address[] memory nodes = oracle.getListedNodes();
        assertEq(nodes.length, 2, "Deux noeuds seulement doivent etre ajoutes");
        assertEq(oracle._fundsStaked(user1), 0.001 ether, "Le montant stake d'Alice doit etre 0.001 ether");
        assertEq(oracle._fundsStaked(user2), 0.001 ether, "Le montant stake de Bob doit etre 0.001 ether");
    }

    function testQuorumWithEnoughNode() public {
        console.log("Test de quorum avec suffisamment de noeuds");
        
        EsportOracleTypes.Match[] memory matches = prepareSampleMatch();

        addFourTestNodes();
        
        console.log("Soumission du match par Alice");
        vm.prank(user1);
        oracle.handleNewMatches(matches);

        console.log("Soumission du match par Bob");
        vm.prank(user2);
        oracle.handleNewMatches(matches);

        console.log("Soumission du match par Charlie");
        vm.prank(user3);
        oracle.handleNewMatches(matches);
        
        console.log("Verification que le match a ete enregistre");
        EsportOracleTypes.Match memory dataNode = oracle.getMatchById(1);
        assertEq(dataNode._id, 1, "L'ID du match doit correspondre");

        assertEq(oracle.getPendingMatches().length, 0, "Il ne doit pas y avoir de matchs en attente");
    }

    function testQuorumWithEnoughSameMatch() public {
        console.log("Test de quorum avec differents matchs soumis");
        
        EsportOracleTypes.Match[] memory matches1 = prepareSampleMatch();
        
        EsportOracleTypes.Match[] memory matches2 = prepareSampleMatch();
        matches2[0]._id = 2;
        matches2[0]._winnerId = 2;

        addFourTestNodes();
        
        console.log("Alice soumet le match 1");
        vm.prank(user1);
        oracle.handleNewMatches(matches1);

        console.log("Bob soumet le match 1");
        vm.prank(user2);
        oracle.handleNewMatches(matches1);

        console.log("Charlie soumet le match 1");
        vm.prank(user3);
        oracle.handleNewMatches(matches1);

        console.log("Dave soumet le match 2");
        vm.prank(user4);
        oracle.handleNewMatches(matches2);

        console.log("Verification que seul le match avec quorum a ete enregistre");
        EsportOracleTypes.Match memory dataNode = oracle.getMatchById(1);
        assertEq(dataNode._id, 1, "Le match 1 doit etre enregistre");

        dataNode = oracle.getMatchById(2);
        assertEq(dataNode._id, 0, "Le match 2 ne doit pas etre enregistre (pas assez de votes)");

        assertEq(oracle.getPendingMatches().length, 1, "Il doit y avoir 1 match en attente");
    }

    function testUpdatingMatchAlreadyRegister() public {
        console.log("Test de mise a jour d'un match deja enregistre");
        
        EsportOracleTypes.Opponents[] memory opponents = new EsportOracleTypes.Opponents[](2);
        opponents[0] = EsportOracleTypes.Opponents({_acronym: "TA", _id: 1, _name: "Team A"});
        opponents[1] = EsportOracleTypes.Opponents({_acronym: "TB", _id: 2, _name: "Team B"});
        
        EsportOracleTypes.Games[] memory initialGames = new EsportOracleTypes.Games[](3);
        initialGames[0] = EsportOracleTypes.Games({_id: 1, _finished: false, _winnerId: 0});
        initialGames[1] = EsportOracleTypes.Games({_id: 2, _finished: false, _winnerId: 0});
        initialGames[2] = EsportOracleTypes.Games({_id: 3, _finished: false, _winnerId: 0});
        
        EsportOracleTypes.Result[] memory results = new EsportOracleTypes.Result[](2);
        results[0] = EsportOracleTypes.Result({_score: 3, _teamId: 1});
        results[1] = EsportOracleTypes.Result({_score: 1, _teamId: 2});
        
        EsportOracleTypes.Match[] memory initialMatch = new EsportOracleTypes.Match[](1);
        initialMatch[0] = EsportOracleTypes.Match({
            _id: 1, _opponents: opponents, _game: initialGames, _result: results, _winnerId: 1, _beginAt: block.timestamp
        });

        addFourTestNodes();
        
        console.log("Soumission initiale du match avec jeux non termines");
        vm.prank(user1);
        oracle.handleNewMatches(initialMatch);
        vm.prank(user2);
        oracle.handleNewMatches(initialMatch);
        vm.prank(user3);
        oracle.handleNewMatches(initialMatch);

        EsportOracleTypes.Games[] memory updatedGames = new EsportOracleTypes.Games[](3);
        updatedGames[0] = EsportOracleTypes.Games({_id: 1, _finished: true, _winnerId: 2});
        updatedGames[1] = EsportOracleTypes.Games({_id: 2, _finished: true, _winnerId: 2});
        updatedGames[2] = EsportOracleTypes.Games({_id: 3, _finished: true, _winnerId: 2});

        EsportOracleTypes.Match[] memory updatedMatch = new EsportOracleTypes.Match[](1);
        updatedMatch[0] = EsportOracleTypes.Match({
            _id: 1, _opponents: opponents, _game: updatedGames, _result: results, _winnerId: 1, _beginAt: block.timestamp
        });

        console.log("Soumission de la mise a jour du match avec jeux termines");
        vm.prank(user1);
        oracle.handleNewMatches(updatedMatch);
        vm.prank(user2);
        oracle.handleNewMatches(updatedMatch);
        vm.prank(user3);
        oracle.handleNewMatches(updatedMatch);

        console.log("Verification que le match a ete mis a jour correctement");
        EsportOracleTypes.Match memory dataNode = oracle.getMatchById(1);
        assertEq(dataNode._id, 1, "L'ID du match doit correspondre");
        assertEq(dataNode._game[0]._finished, true, "Le jeu doit etre marque comme termine");
        assertEq(dataNode._game[0]._winnerId, 2, "L'ID du gagnant doit etre mis a jour");

        assertEq(oracle.getPendingMatches().length, 0, "Il ne doit pas y avoir de matchs en attente");
    }

    function testBanAfterMultipleViolations() public {
        console.log("Test du systeme de bannissement apres violations repetees");
        
        addFourTestNodes();
        console.log("4 noeuds ajoutes au systeme");
        
        EsportOracleTypes.Opponents[] memory opponents = new EsportOracleTypes.Opponents[](2);
        opponents[0] = EsportOracleTypes.Opponents({_acronym: "TA", _id: 1, _name: "Team A"});
        opponents[1] = EsportOracleTypes.Opponents({_acronym: "TB", _id: 2, _name: "Team B"});
        
        EsportOracleTypes.Games[] memory games = new EsportOracleTypes.Games[](1);
        games[0] = EsportOracleTypes.Games({_id: 1, _finished: true, _winnerId: 1});
        
        EsportOracleTypes.Result[] memory results = new EsportOracleTypes.Result[](2);
        results[0] = EsportOracleTypes.Result({_score: 3, _teamId: 1});
        results[1] = EsportOracleTypes.Result({_score: 1, _teamId: 2});
        
        EsportOracleTypes.Match[] memory correctMatch1 = new EsportOracleTypes.Match[](1);
        correctMatch1[0] = EsportOracleTypes.Match({
            _id: 1, _opponents: opponents, _game: games, _result: results, _winnerId: 1, _beginAt: 1
        });
        
        EsportOracleTypes.Match[] memory correctMatch2 = new EsportOracleTypes.Match[](1);
        correctMatch2[0] = EsportOracleTypes.Match({
            _id: 2, _opponents: opponents, _game: games, _result: results, _winnerId: 1, _beginAt: 1
        });
        
        EsportOracleTypes.Match[] memory correctMatch3 = new EsportOracleTypes.Match[](1);
        correctMatch3[0] = EsportOracleTypes.Match({
            _id: 3, _opponents: opponents, _game: games, _result: results, _winnerId: 1, _beginAt: 1
        });
        
        EsportOracleTypes.Games[] memory incorrectGames = new EsportOracleTypes.Games[](1);
        incorrectGames[0] = EsportOracleTypes.Games({_id: 1, _finished: false, _winnerId: 2});
        
        EsportOracleTypes.Match[] memory incorrectMatch1 = new EsportOracleTypes.Match[](1);
        incorrectMatch1[0] = EsportOracleTypes.Match({
            _id: 1, _opponents: opponents, _game: incorrectGames, _result: results, _winnerId: 2, _beginAt: 1
        });
        
        EsportOracleTypes.Match[] memory incorrectMatch2 = new EsportOracleTypes.Match[](1);
        incorrectMatch2[0] = EsportOracleTypes.Match({
            _id: 2, _opponents: opponents, _game: incorrectGames, _result: results, _winnerId: 2, _beginAt: 1
        });
        
        EsportOracleTypes.Match[] memory incorrectMatch3 = new EsportOracleTypes.Match[](1);
        incorrectMatch3[0] = EsportOracleTypes.Match({
            _id: 3, _opponents: opponents, _game: incorrectGames, _result: results, _winnerId: 2, _beginAt: 1
        });
        
        vm.prank(user4);
        oracle.handleNewMatches(incorrectMatch1);
        
        vm.prank(user1);
        oracle.handleNewMatches(correctMatch1);
        
        vm.prank(user2);
        oracle.handleNewMatches(correctMatch1);
        
        vm.prank(user3);
        oracle.handleNewMatches(correctMatch1);
        
        (uint256 violations, bool banned) = oracle._nodeViolations(user4);
        assertEq(violations, 1, "After 1 violation, counter should be 1");
        assertEq(banned, false, "Node should not be banned yet");
        assertEq(oracle._fundsStaked(user4), 0.0009 ether, "Balance should be reduced after 1 violation");
        
        vm.prank(user4);
        oracle.handleNewMatches(incorrectMatch2);
        
        vm.prank(user1);
        oracle.handleNewMatches(correctMatch2);
        
        vm.prank(user2);
        oracle.handleNewMatches(correctMatch2);
        
        vm.prank(user3);
        oracle.handleNewMatches(correctMatch2);
        
        (violations, banned) = oracle._nodeViolations(user4);
        assertEq(violations, 2, "After 2 violations, counter should be 2");
        assertEq(banned, false, "Node should still not be banned");
        assertEq(oracle._fundsStaked(user4), 0.0008 ether, "Balance should be reduced after 2 violations");
        
        vm.prank(user4);
        oracle.handleNewMatches(incorrectMatch3);
        
        vm.prank(user1);
        oracle.handleNewMatches(correctMatch3);
        
        vm.prank(user2);
        oracle.handleNewMatches(correctMatch3);
        
        vm.recordLogs();
        
        vm.prank(user3);
        oracle.handleNewMatches(correctMatch3);
        
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
        
        (violations, banned) = oracle._nodeViolations(user4);
        assertEq(violations, 3, "After 3 violations, counter should be 3");
        assertEq(banned, true, "Node should be banned after 3 violations");
        assertEq(oracle._fundsStaked(user4), 0, "Balance should be 0 after ban");
        
        vm.prank(user4);
        vm.expectRevert("Node is banned");
        oracle.handleNewMatches(correctMatch1);
    }

    function testAdminFunctions() public {
        vm.prank(user1);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        vm.prank(user2);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        vm.expectEmit(true, false, false, false);
        emit NodeBanned(user1);

        oracle.banNode(user1);

        (uint256 violations, bool banned) = oracle._nodeViolations(user1);
        assertEq(banned, true, "Le noeud doit etre banni");
        assertEq(oracle._fundsStaked(user1), 0, "Les fonds doivent etre confisques");
        
        vm.prank(user1);
        vm.expectRevert("Node is banned");
        oracle.handleNewMatches(new EsportOracleTypes.Match[](1));
        
        oracle.rehabilitateNode(user1);
        
        (violations, banned) = oracle._nodeViolations(user1);
        assertEq(banned, false, "Le noeud ne doit plus etre banni apres rehabilitation");
        assertEq(violations, 0, "Le compteur de violations doit etre remis a zero");
        
        vm.prank(user1);
        vm.expectRevert("Node is not listed, please call addNewNode function to register a new node");
        oracle.handleNewMatches(new EsportOracleTypes.Match[](1));
    }

    function testWithdrawStake() public {
        vm.prank(user1);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        assertEq(oracle._fundsStaked(user1), 0.001 ether, "Le solde initial doit etre 0.001 ether");
        
        uint256 balanceBefore = user1.balance;
        
        vm.prank(user1);
        oracle.withdrawStake();
        
        assertEq(oracle._fundsStaked(user1), 0, "Le solde doit etre a 0 apres retrait");
        assertEq(user1.balance, balanceBefore + 0.001 ether, "L'ETH doit etre retourne a l'utilisateur");
        
        address[] memory nodes = oracle.getListedNodes();
        assertEq(nodes.length, 0, "La liste des noeuds doit etre vide apres retrait");
    }

    function testFundsRedistribution() public {
        vm.prank(user1);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        vm.prank(user2);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        vm.prank(user3);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        vm.prank(user4);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        oracle.banNode(user1);
        
        uint256 expectedShare = uint256(0.001 ether) / 3;
        
        assertEq(oracle._fundsStaked(user1), 0, "Le noeud banni doit avoir un solde de 0");
        assertEq(oracle._fundsStaked(user2), 0.001 ether + expectedShare, "Le noeud doit recevoir sa part");
        assertEq(oracle._fundsStaked(user3), 0.001 ether + expectedShare, "Le noeud doit recevoir sa part");
        assertEq(oracle._fundsStaked(user4), 0.001 ether + expectedShare, "Le noeud doit recevoir sa part");
    }

    function prepareSampleMatch() internal view returns (EsportOracleTypes.Match[] memory) {
        EsportOracleTypes.Opponents[] memory opponents = new EsportOracleTypes.Opponents[](2);
        opponents[0] = EsportOracleTypes.Opponents({_acronym: "TA", _id: 1, _name: "Team A"});
        opponents[1] = EsportOracleTypes.Opponents({_acronym: "TB", _id: 2, _name: "Team B"});
        
        EsportOracleTypes.Games[] memory games = new EsportOracleTypes.Games[](1);
        games[0] = EsportOracleTypes.Games({_id: 1, _finished: true, _winnerId: 1});
        
        EsportOracleTypes.Result[] memory results = new EsportOracleTypes.Result[](2);
        results[0] = EsportOracleTypes.Result({_score: 3, _teamId: 1});
        results[1] = EsportOracleTypes.Result({_score: 1, _teamId: 2});
        
        EsportOracleTypes.Match[] memory matches = new EsportOracleTypes.Match[](1);
        matches[0] = EsportOracleTypes.Match({
            _id: 1, _opponents: opponents, _game: games, _result: results, _winnerId: 1, _beginAt: block.timestamp
        });
        
        return matches;
    }

    function addFourTestNodes() internal {
        vm.prank(user1);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        vm.prank(user2);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        vm.prank(user3);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        vm.prank(user4);
        oracle.addFundToStaking{value: 0.001 ether}();
    }

    function testQuorumWithTwoNodes() public {
        console.log("Test de quorum avec seulement 2 noeuds");
        
        vm.prank(user1);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        vm.prank(user2);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        EsportOracleTypes.Match[] memory matches = prepareSampleMatch();
        
        vm.prank(user1);
        oracle.handleNewMatches(matches);
        
        vm.prank(user2);
        oracle.handleNewMatches(matches);
        
        EsportOracleTypes.Match memory dataNode = oracle.getMatchById(1);
        assertEq(dataNode._id, 1, "Le match doit etre enregistre avec 2 noeuds");
        assertEq(oracle.getPendingMatches().length, 0, "Il ne doit pas y avoir de matchs en attente");
    }

    function testQuorumWithThreeNodes() public {
        console.log("Test de quorum avec 3 noeuds (2 sur 3 suffisent)");
        
        vm.prank(user1);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        vm.prank(user2);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        vm.prank(user3);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        EsportOracleTypes.Match[] memory matches = prepareSampleMatch();
        
        vm.prank(user1);
        oracle.handleNewMatches(matches);
        
        vm.prank(user2);
        oracle.handleNewMatches(matches);
        
        EsportOracleTypes.Match memory dataNode = oracle.getMatchById(1);
        assertEq(dataNode._id, 1, "Le match doit etre enregistre avec 2 sur 3 noeuds");
        assertEq(oracle.getPendingMatches().length, 0, "Il ne doit pas y avoir de matchs en attente");
    }

    function testNodeCannotSubmitWhenBanned() public {
        console.log("Test qu'un noeud banni ne peut pas soumettre de donnees");
        
        addFourTestNodes();
        
        oracle.banNode(user1);
        
        EsportOracleTypes.Match[] memory matches = prepareSampleMatch();
        
        vm.prank(user1);
        vm.expectRevert("Node is banned");
        oracle.handleNewMatches(matches);
    }

    function testNodeCannotStakeWhenBanned() public {
        console.log("Test qu'un noeud banni ne peut pas faire de staking");
        
        addFourTestNodes();
        
        oracle.banNode(user1);
        
        vm.prank(user1);
        vm.expectRevert("Node is banned");
        oracle.addFundToStaking{value: 0.001 ether}();
    }

    function testRehabilitatedNodeCanStakeAgain() public {
        console.log("Test qu'un noeud rehabilite peut refaire du staking");
        
        vm.prank(user1);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        oracle.banNode(user1);
        
        oracle.rehabilitateNode(user1);
        
        vm.prank(user1);
        oracle.addFundToStaking{value: 0.001 ether}();
        
        address[] memory nodes = oracle.getListedNodes();
        assertEq(nodes.length, 1, "Le noeud rehabilite doit etre dans la liste");
    }

    function testWithdrawStakeWhenNotListed() public {
        console.log("Test de retrait quand le noeud n'est pas dans la liste");
        
        vm.expectRevert("No funds to withdraw");
        oracle.withdrawStake();
    }

    function testBanNodeWithZeroStake() public {
        console.log("Test de bannissement d'un noeud sans stake");
        
        vm.expectRevert();
        oracle.banNode(user1);
    }

    function testSetOwnerToZeroAddress() public {
        console.log("Test de definition du proprietaire a l'adresse zero");
        
        vm.expectRevert("New owner cannot be zero address");
        oracle.setOwner(address(0));
    }

    function testPunishNodeWithInsufficientFunds() public {
        console.log("Test de punition d'un noeud avec fonds insuffisants");
        
        addFourTestNodes();
        
        EsportOracleTypes.Opponents[] memory opponents = new EsportOracleTypes.Opponents[](2);
        opponents[0] = EsportOracleTypes.Opponents({_acronym: "TA", _id: 1, _name: "Team A"});
        opponents[1] = EsportOracleTypes.Opponents({_acronym: "TB", _id: 2, _name: "Team B"});
        
        EsportOracleTypes.Games[] memory games = new EsportOracleTypes.Games[](1);
        games[0] = EsportOracleTypes.Games({_id: 1, _finished: true, _winnerId: 1});
        
        EsportOracleTypes.Result[] memory results = new EsportOracleTypes.Result[](2);
        results[0] = EsportOracleTypes.Result({_score: 3, _teamId: 1});
        results[1] = EsportOracleTypes.Result({_score: 1, _teamId: 2});
        
        EsportOracleTypes.Match[] memory correctMatch = new EsportOracleTypes.Match[](1);
        correctMatch[0] = EsportOracleTypes.Match({
            _id: 1, _opponents: opponents, _game: games, _result: results, _winnerId: 1, _beginAt: 1
        });
        
        EsportOracleTypes.Match[] memory incorrectMatch = new EsportOracleTypes.Match[](1);
        incorrectMatch[0] = EsportOracleTypes.Match({
            _id: 1, _opponents: opponents, _game: games, _result: results, _winnerId: 2, _beginAt: 1
        });
        
        vm.prank(user1);
        oracle.handleNewMatches(correctMatch);
        vm.prank(user2);
        oracle.handleNewMatches(correctMatch);
        vm.prank(user3);
        oracle.handleNewMatches(correctMatch);
        
        vm.prank(user4);
        oracle.handleNewMatches(incorrectMatch);
        
        (uint256 violations, bool banned) = oracle._nodeViolations(user4);
        assertEq(violations, 0, "Le noeud ne doit pas avoir de violation dans ce scenario");
    }

    function testEmptyMatchArray() public {
        console.log("Test de soumission d'un tableau de matchs vide");
        
        addFourTestNodes();
        
        EsportOracleTypes.Match[] memory emptyMatches = new EsportOracleTypes.Match[](0);
        
        vm.prank(user1);
        vm.expectRevert("No match data provided");
        oracle.handleNewMatches(emptyMatches);
    }

}
