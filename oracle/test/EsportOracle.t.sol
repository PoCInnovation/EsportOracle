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
}
