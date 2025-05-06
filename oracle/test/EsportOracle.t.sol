// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/esportOracle.sol";

contract EsportOracleTest is Test {
    EsportOracle public oracle;
    address public owner;
    address public user1;
    address public user2;

    event newNodeAdded(address indexed addressAdded);

    function setUp() public {
        owner = address(this);
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");

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

    function testAddNewNode() public {
        vm.prank(user1);
        vm.expectEmit(true, false, false, false);
        emit newNodeAdded(user1);
        oracle.addNewNode();

        address[] memory nodes = oracle.getListedNodes();
        assertEq(nodes.length, 1, "Un seul noeud doit etre ajoute");
        assertEq(nodes[0], user1, "L'adresse du noeud doit correspondre");
    }

    function testMultipleNode() public {
        vm.prank(user1);
        vm.expectEmit(true, false, false, false);
        emit newNodeAdded(user1);
        oracle.addNewNode();

        vm.prank(user2);
        vm.expectEmit(true, false, false, false);
        emit newNodeAdded(user2);
        oracle.addNewNode();

        address[] memory nodes = oracle.getListedNodes();
        assertEq(nodes.length, 2, "Deux noeud seulement doivent etre ajoute");
        assertEq(nodes[0], user1, "L'adresse du noeud doit correspondre");
        assertEq(nodes[1], user2, "L'adresse du noeud doit correspondre");
    }

    function testAlreadyListedNode() public {
        vm.prank(user1);
        vm.expectEmit(true, false, false, false);
        emit newNodeAdded(user1);
        oracle.addNewNode();

        vm.prank(user1);
        vm.expectRevert("Node is already listed");
        oracle.addNewNode();

        address[] memory nodes = oracle.getListedNodes();
        assertEq(nodes.length, 1, "Un seul noeud doit etre ajoute");
    }

    function testAddNewNodeWithAddress0() public {
        vm.prank(address(0));
        vm.expectRevert("New node cannot be zero address");
        oracle.addNewNode();
    }

    function testAddNewNodeWithOwner() public {
        vm.prank(owner);
        vm.expectRevert("owner cannot be a node");
        oracle.addNewNode();
    }

    function testQuorumWithNotEnoughNode() public {
        vm.prank(user1);
        oracle.addNewNode();

        vm.prank(user2);
        oracle.addNewNode();

        vm.prank(owner);
        vm.expectRevert("require at least 3 nodes to reach quorum");
        oracle.isQuorumReached();
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

        address user3 = makeAddr("user3");

        vm.prank(user1);
        oracle.addNewNode();

        vm.prank(user2);
        oracle.addNewNode();

        vm.prank(user3);
        oracle.addNewNode();

        vm.prank(user1);
        oracle.addDataNode(matches);

        vm.prank(user2);
        oracle.addDataNode(matches);

        vm.prank(user3);
        oracle.addDataNode(matches);

        assertTrue(oracle.isQuorumReached(), "Le quorum doit etre atteint");
    }
}
