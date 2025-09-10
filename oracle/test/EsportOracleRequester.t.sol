// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/esportOracleRequester.sol";
import "../src/esportOracleClientRequester.sol";
import "../src/esportOracleTypes.sol";

contract MockClient is EsportOracleClientRequester {
	using EsportOracleTypes for EsportOracleTypes.Match;

	uint256 public lastReceivedMatchId;
	uint256 public lastReceivedWinnerId;
	bool public received;

	event MatchReceived(uint256 indexed matchId, uint256 indexed winnerId);

	constructor(address oracle_) EsportOracleClientRequester(oracle_) {}

	function callMatchReceived(EsportOracleTypes.Match memory _match) external override onlyOracle {
		lastReceivedMatchId = _match._id;
		lastReceivedWinnerId = _match._winnerId;
		received = true;
		emit MatchReceived(_match._id, _match._winnerId);
	}
}

contract EsportOracleRequesterTest is Test {
	using EsportOracleTypes for EsportOracleTypes.Match;

	EsportOracleRequester public oracle;
	MockClient public client;

	address public node1;
	address public node2;
	address public node3;
	address public node4;

	uint256 constant STAKE = 0.001 ether;
	uint256 constant MIN_REQUEST_FEE = 0.0001 ether;

	function setUp() public {
		node1 = makeAddr("node1");
		node2 = makeAddr("node2");
		node3 = makeAddr("node3");
		node4 = makeAddr("node4");

		vm.deal(node1, 1 ether);
		vm.deal(node2, 1 ether);
		vm.deal(node3, 1 ether);
		vm.deal(node4, 1 ether);

		oracle = new EsportOracleRequester();
		client = new MockClient(address(oracle));
	}

	function testRequestMatchInsufficientFee() public {
		vm.expectRevert("Insufficient request fee");
		client.receiveMatch{value: 0}(1);
	}

	function testRequestMatchAndPreventDuplicateUntilFulfilled() public {
		client.receiveMatch{value: MIN_REQUEST_FEE}(1);

		uint256[] memory pendings = oracle.getPendingRequestedMatches();
		assertEq(pendings.length, 1);
		assertEq(pendings[0], 1);

		vm.expectRevert("Match already requested");
		client.receiveMatch{value: MIN_REQUEST_FEE}(1);
	}

	function testMultiplePendingRequestsAreListed() public {
		client.receiveMatch{value: MIN_REQUEST_FEE}(1);
		client.receiveMatch{value: MIN_REQUEST_FEE}(2);

		uint256[] memory pendings = oracle.getPendingRequestedMatches();
		assertEq(pendings.length, 2);
		assertEq(pendings[0], 1);
		assertEq(pendings[1], 2);
	}

	function testCallMatchOracleFulfillsAndCallbacks() public {
		uint256 matchId = 42;
		client.receiveMatch{value: MIN_REQUEST_FEE}(matchId);

		_addFourNodes();
		EsportOracleTypes.Match[] memory matches = _prepareSampleMatch(matchId, 7);

		vm.prank(node1); oracle.handleNewMatches(matches);
		vm.prank(node2); oracle.handleNewMatches(matches);
		vm.prank(node3); oracle.handleNewMatches(matches);

		EsportOracleTypes.Match memory accepted = oracle.getMatchById(matchId);
		assertEq(accepted._id, matchId);

		oracle.callMatchOracle(matchId, accepted);

		assertTrue(client.received());
		assertEq(client.lastReceivedMatchId(), matchId);
		assertEq(client.lastReceivedWinnerId(), 7);

		EsportOracleTypes.MatchRequest memory req = oracle.getRequestByMatchId(matchId);
		assertTrue(req.fulfilled);
	}

	function _addFourNodes() internal {
		vm.prank(node1); oracle.addFundToStaking{value: STAKE}();
		vm.prank(node2); oracle.addFundToStaking{value: STAKE}();
		vm.prank(node3); oracle.addFundToStaking{value: STAKE}();
		vm.prank(node4); oracle.addFundToStaking{value: STAKE}();
	}

	function _prepareSampleMatch(uint256 id, uint256 winnerId) internal view returns (EsportOracleTypes.Match[] memory) {
		EsportOracleTypes.Opponents[] memory opponents = new EsportOracleTypes.Opponents[](2);
		opponents[0] = EsportOracleTypes.Opponents({_acronym: "TA", _id: 1, _name: "Team A"});
		opponents[1] = EsportOracleTypes.Opponents({_acronym: "TB", _id: 2, _name: "Team B"});

		EsportOracleTypes.Games[] memory games = new EsportOracleTypes.Games[](1);
		games[0] = EsportOracleTypes.Games({_id: 1, _finished: true, _winnerId: winnerId});

		EsportOracleTypes.Result[] memory results = new EsportOracleTypes.Result[](2);
		results[0] = EsportOracleTypes.Result({_score: 1, _teamId: 1});
		results[1] = EsportOracleTypes.Result({_score: 0, _teamId: 2});

		EsportOracleTypes.Match[] memory matches = new EsportOracleTypes.Match[](1);
		matches[0] = EsportOracleTypes.Match({_id: id, _opponents: opponents, _game: games, _result: results, _winnerId: winnerId, _beginAt: block.timestamp});
		return matches;
	}

	function testRequestMatchWithExactFee() public {
		client.receiveMatch{value: MIN_REQUEST_FEE}(42);
		
		uint256[] memory pendings = oracle.getPendingRequestedMatches();
		assertEq(pendings.length, 1);
		assertEq(pendings[0], 42);
	}

	function testRequestMatchWithExcessFee() public {
		client.receiveMatch{value: MIN_REQUEST_FEE + 0.001 ether}(43);
		
		uint256[] memory pendings = oracle.getPendingRequestedMatches();
		assertEq(pendings.length, 1);
		assertEq(pendings[0], 43);
	}

	function testMultipleClientsRequestSameMatch() public {
		MockClient client2 = new MockClient(address(oracle));
		
		client.receiveMatch{value: MIN_REQUEST_FEE}(44);
		
		vm.expectRevert("Match already requested");
		client2.receiveMatch{value: MIN_REQUEST_FEE}(44);
	}

	function testGetMatchRequestForNonExistentRequest() public {
		EsportOracleTypes.MatchRequest memory req = oracle.getMatchRequest(999);
		assertEq(req.matchId, 0, "Should return empty request for non-existent ID");
		assertEq(req.requester, address(0), "Requester should be zero address");
		assertEq(req.fee, 0, "Fee should be 0");
		assertEq(req.fulfilled, false, "Should not be fulfilled");
	}

	function testMarkRequestsFulfilledMultipleTimes() public {
		client.receiveMatch{value: MIN_REQUEST_FEE}(45);
		
		oracle.markRequestsFulfilled(45);
		
		oracle.markRequestsFulfilled(45);
		
		EsportOracleTypes.MatchRequest memory req = oracle.getRequestByMatchId(45);
		assertTrue(req.fulfilled, "Request should remain fulfilled");
	}

	function testCallMatchOracleWithNonExistentMatch() public {
		client.receiveMatch{value: MIN_REQUEST_FEE}(46);
		
		EsportOracleTypes.Match memory nonExistentMatch;
		nonExistentMatch._id = 46;
		nonExistentMatch._winnerId = 7;
		
		vm.expectRevert("Match not validated by quorum yet");
		oracle.callMatchOracle(46, nonExistentMatch);
	}

	function testCallMatchOracleWithInvalidRequester() public {
		client.receiveMatch{value: MIN_REQUEST_FEE}(47);
		
		_addFourNodes();
		EsportOracleTypes.Match[] memory matches = _prepareSampleMatch(47, 7);
		
		vm.prank(node1); oracle.handleNewMatches(matches);
		vm.prank(node2); oracle.handleNewMatches(matches);
		vm.prank(node3); oracle.handleNewMatches(matches);
		
		EsportOracleTypes.Match memory accepted = oracle.getMatchById(47);
		
		EsportOracleTypes.Match memory differentMatch;
		differentMatch._id = 999;
		differentMatch._winnerId = 7;
		
		vm.expectRevert("Invalid address requester");
		oracle.callMatchOracle(999, differentMatch);
	}

	    function testRequestCounterIncrement() public {
        uint256 initialRequestId = client.receiveMatch{value: MIN_REQUEST_FEE}(48);
        uint256 secondRequestId = client.receiveMatch{value: MIN_REQUEST_FEE}(49);
        
        assertEq(initialRequestId, 1, "First request ID should be 1");
        assertEq(secondRequestId, 2, "Second request ID should be 2");
    }

	function testPendingRequestsAfterFulfillment() public {
		client.receiveMatch{value: MIN_REQUEST_FEE}(50);
		client.receiveMatch{value: MIN_REQUEST_FEE}(51);
		
		uint256[] memory pendings = oracle.getPendingRequestedMatches();
		assertEq(pendings.length, 2, "Should have 2 pending requests");
		
		oracle.markRequestsFulfilled(50);
		
		pendings = oracle.getPendingRequestedMatches();
		assertEq(pendings.length, 1, "Should have 1 pending request after fulfillment");
		assertEq(pendings[0], 51, "Remaining request should be for match 51");
	}

	    function testRequestMatchWhenPaused() public {
        oracle.pause();
        
        vm.expectRevert();
        client.receiveMatch{value: MIN_REQUEST_FEE}(52);
        
        oracle.unpause();
        
        client.receiveMatch{value: MIN_REQUEST_FEE}(52);
        uint256[] memory pendings = oracle.getPendingRequestedMatches();
        assertEq(pendings.length, 1);
    }

	function testGetRequestByMatchId() public {
		client.receiveMatch{value: MIN_REQUEST_FEE}(53);
		
		EsportOracleTypes.MatchRequest memory req = oracle.getRequestByMatchId(53);
		assertEq(req.matchId, 53, "Match ID should match");
		assertEq(req.requester, address(client), "Requester should be client address");
		assertEq(req.fee, MIN_REQUEST_FEE, "Fee should match");
		assertEq(req.fulfilled, false, "Should not be fulfilled initially");
	}
}


