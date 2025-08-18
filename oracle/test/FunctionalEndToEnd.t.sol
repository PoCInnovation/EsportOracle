// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/BetContract.sol";
import "../src/esportOracleRequester.sol";
import "../src/esportOracleTypes.sol";
import "../lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";

contract MockERC20E2E is ERC20 {
	constructor() ERC20("Functional Test Token", "FTT") {
		_mint(msg.sender, 10_000_000 ether);
	}

	function mint(address to, uint256 amount) external {
		_mint(to, amount);
	}
}

contract FunctionalEndToEndTest is Test {
	using EsportOracleTypes for EsportOracleTypes.Match;

	EsportOracleRequester public oracle;
	BetContract public betContract;
	MockERC20E2E public token;

	address public node1;
	address public node2;
	address public node3;
	address public node4;

	address public user1;
	address public user2;

	uint256 constant STAKE = 0.001 ether;
	uint256 constant REQUEST_FEE = 0.001 ether;

	function setUp() public {
		node1 = makeAddr("node1");
		node2 = makeAddr("node2");
		node3 = makeAddr("node3");
		node4 = makeAddr("node4");

		user1 = makeAddr("user1");
		user2 = makeAddr("user2");

		vm.deal(node1, 1 ether);
		vm.deal(node2, 1 ether);
		vm.deal(node3, 1 ether);
		vm.deal(node4, 1 ether);

		oracle = new EsportOracleRequester();
		token = new MockERC20E2E();
		betContract = new BetContract(address(oracle), address(token));

		token.mint(user1, 1_000 ether);
		token.mint(user2, 1_000 ether);
		vm.prank(user1); token.approve(address(betContract), type(uint256).max);
		vm.prank(user2); token.approve(address(betContract), type(uint256).max);
	}

	function _addFourNodes() internal {
		vm.prank(node1); oracle.addFundToStaking{value: STAKE}();
		vm.prank(node2); oracle.addFundToStaking{value: STAKE}();
		vm.prank(node3); oracle.addFundToStaking{value: STAKE}();
		vm.prank(node4); oracle.addFundToStaking{value: STAKE}();
	}

	function _prepareMatch(uint256 id, uint256 winnerId) internal view returns (EsportOracleTypes.Match[] memory) {
		EsportOracleTypes.Opponents[] memory opponents = new EsportOracleTypes.Opponents[](2);
		opponents[0] = EsportOracleTypes.Opponents({_acronym: "T1", _id: 111, _name: "Team One"});
		opponents[1] = EsportOracleTypes.Opponents({_acronym: "T2", _id: 222, _name: "Team Two"});

		EsportOracleTypes.Games[] memory games = new EsportOracleTypes.Games[](1);
		games[0] = EsportOracleTypes.Games({_id: 1, _finished: true, _winnerId: winnerId});

		EsportOracleTypes.Result[] memory results = new EsportOracleTypes.Result[](2);
		results[0] = EsportOracleTypes.Result({_score: 1, _teamId: 111});
		results[1] = EsportOracleTypes.Result({_score: 0, _teamId: 222});

		EsportOracleTypes.Match[] memory matches = new EsportOracleTypes.Match[](1);
		matches[0] = EsportOracleTypes.Match({_id: id, _opponents: opponents, _game: games, _result: results, _winnerId: winnerId, _beginAt: block.timestamp});
		return matches;
	}

	function test_endToEnd_BetResolutionViaOracle() public {
		_addFourNodes();

		uint256 deadline = block.timestamp + 1 hours;
		uint256 team1Id = 111;
		uint256 team2Id = 222;
		uint256 matchId = 9999;

		betContract.createBet("Functional E2E Match", team1Id, team2Id, deadline, matchId);

		vm.prank(user1); betContract.placeBet(0, 1, 200 ether);
		vm.prank(user2); betContract.placeBet(0, 2, 100 ether);

		vm.deal(address(this), REQUEST_FEE);
		betContract.requestMatchIfNeeded{value: REQUEST_FEE}(matchId);

		EsportOracleTypes.Match[] memory matches = _prepareMatch(matchId, team1Id);
		vm.prank(node1); oracle.handleNewMatches(matches);
		vm.prank(node2); oracle.handleNewMatches(matches);
		vm.prank(node3); oracle.handleNewMatches(matches);

		EsportOracleTypes.Match memory accepted = oracle.getMatchById(matchId);
		assertEq(accepted._id, matchId, "Match must be validated by quorum");
		vm.deal(address(betContract), REQUEST_FEE);
		oracle.callMatchOracle(matchId, accepted);

		vm.warp(deadline + 1);
		vm.prank(address(oracle));
		betContract.resolveBet(0, accepted);

		uint256 user1Before = token.balanceOf(user1);
		vm.prank(user1); betContract.claimWinnings(0);
		uint256 user1After = token.balanceOf(user1);
		assertEq(user1After - user1Before, 300 ether, "Winner should receive full pool proportionally");

		vm.prank(user2);
		vm.expectRevert("Equipe perdante");
		betContract.claimWinnings(0);
	}
}


