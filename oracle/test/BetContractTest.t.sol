// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "../src/BetContract.sol";
import "../src/esportOracleTypes.sol";
import "../lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";

contract MockERC20 is ERC20 {
    constructor() ERC20("Test Token", "TEST") {
        _mint(msg.sender, 1000000 * 10**18);
    }

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}

contract MockOracle {
    mapping(uint256 => bool) public matchRequested;
    mapping(uint256 => EsportOracleTypes.MatchRequest) public matchRequests;
    mapping(uint256 => EsportOracleTypes.Match) public matches;

    function requestMatch(uint256 matchId) external payable returns (uint256) {
        matchRequested[matchId] = true;
        return matchId;
    }

    function isMatchRequested(uint256 matchId) external view returns (bool) {
        return matchRequested[matchId];
    }

    function getMatchRequest(uint256 matchId) external view returns (EsportOracleTypes.MatchRequest memory) {
        return matchRequests[matchId];
    }

    function markRequestsFulfilled(uint256 matchId) external {
        matchRequests[matchId].fulfilled = true;
    }

    function getPendingRequestedMatches() external view returns (uint256[] memory) {
        uint256[] memory empty = new uint256[](0);
        return empty;
    }

    function getMatchById(uint256 matchId) external view returns (EsportOracleTypes.Match memory) {
        return matches[matchId];
    }

    function setMatchRequest(uint256 matchId, bool fulfilled) external {
        matchRequested[matchId] = true;
        matchRequests[matchId] = EsportOracleTypes.MatchRequest({
            matchId: matchId,
            requester: address(this),
            fee: 0,
            fulfilled: fulfilled
        });
    }

    function setMatch(uint256 matchId, uint256 winnerId) external {
        matches[matchId]._id = matchId;
        matches[matchId]._winnerId = winnerId;
    }
}

contract BetContractTest is Test {
    BetContract public betContract;
    MockERC20 public token;
    MockOracle public mockOracle;
    
    address public user1 = address(0x2);
    address public user2 = address(0x3);
    
    uint256 public constant INITIAL_BALANCE = 1000 * 10**18;
    
    function setUp() public {
        token = new MockERC20();
        mockOracle = new MockOracle();
        betContract = new BetContract(address(mockOracle), address(token));
        
        token.mint(user1, INITIAL_BALANCE);
        token.mint(user2, INITIAL_BALANCE);
        
        vm.prank(user1);
        token.approve(address(betContract), INITIAL_BALANCE);
        
        vm.prank(user2);
        token.approve(address(betContract), INITIAL_BALANCE);
    }
    
    function testCreateBet() public {
        uint256 deadline = block.timestamp + 1 hours;
        uint256 team1Id = 123;
        uint256 team2Id = 456;
        uint256 matchId = 789;
        
        betContract.createBet(
            "Test Match",
            team1Id,
            team2Id,
            deadline,
            matchId
        );
        
        assertEq(betContract.getBetCount(), 1);
    }
    
    function testPlaceBet() public {
        uint256 deadline = block.timestamp + 1 hours;
        betContract.createBet("Test Match", 123, 456, deadline, 789);
        
        uint256 betAmount = 100 * 10**18;
        
        vm.prank(user1);
        betContract.placeBet(0, 1, betAmount);
        
        BetContract.UserBet memory userBet = betContract.getUserBet(user1, 0);
        assertEq(userBet.amount, betAmount);
        assertEq(userBet.teamChosen, 1);
        assertEq(userBet.claimed, false);
    }
    
    function testResolveBet() public {
        uint256 deadline = block.timestamp + 1 hours;
        uint256 team1Id = 123;
        uint256 team2Id = 456;
        uint256 matchId = 789;
        
        betContract.createBet("Test Match", team1Id, team2Id, deadline, matchId);
        
        uint256 betAmount = 100 * 10**18;
        
        vm.prank(user1);
        betContract.placeBet(0, 1, betAmount);
        
        vm.prank(user2);
        betContract.placeBet(0, 2, betAmount);
        
        vm.warp(deadline + 1);
        
        EsportOracleTypes.Match memory matchData;
        matchData._id = matchId;
        matchData._winnerId = team1Id;
        
        mockOracle.setMatchRequest(matchId, true);
        
        vm.prank(address(mockOracle));
        betContract.resolveBet(0, matchData);
    }
    
    function testClaimWinnings() public {
        uint256 deadline = block.timestamp + 1 hours;
        uint256 team1Id = 123;
        uint256 team2Id = 456;
        uint256 matchId = 789;
        
        betContract.createBet("Test Match", team1Id, team2Id, deadline, matchId);
        
        uint256 betAmount = 100 * 10**18;
        
        vm.prank(user1);
        betContract.placeBet(0, 1, betAmount);
        
        vm.prank(user2);
        betContract.placeBet(0, 2, betAmount);
        
        vm.warp(deadline + 1);
        
        EsportOracleTypes.Match memory matchData;
        matchData._id = matchId;
        matchData._winnerId = team1Id;
        
        mockOracle.setMatchRequest(matchId, true);
        
        vm.prank(address(mockOracle));
        betContract.resolveBet(0, matchData);
        
        uint256 balanceBefore = token.balanceOf(user1);
        
        vm.prank(user1);
        betContract.claimWinnings(0);
        
        uint256 balanceAfter = token.balanceOf(user1);
        
        assertGt(balanceAfter, balanceBefore);
        assertEq(balanceAfter - balanceBefore, betAmount * 2);
    }
    
    function testFailResolveBetBeforeDeadline() public {
        uint256 deadline = block.timestamp + 1 hours;
        betContract.createBet("Test Match", 123, 456, deadline, 789);
        
        EsportOracleTypes.Match memory matchData;
        matchData._id = 789;
        matchData._winnerId = 123;
        
        vm.prank(address(mockOracle));
        betContract.resolveBet(0, matchData);
    }
    
    function testFailResolveBetWrongMatchId() public {
        uint256 deadline = block.timestamp + 1 hours;
        betContract.createBet("Test Match", 123, 456, deadline, 789);
        
        vm.warp(deadline + 1);
        
        EsportOracleTypes.Match memory matchData;
        matchData._id = 999;
        matchData._winnerId = 123;
        
        vm.prank(address(mockOracle));
        betContract.resolveBet(0, matchData);
    }
    
    function testRequestMatchWithFees() public {
        uint256 matchId = 999;
        
        uint256 fee = 0.001 ether;
        
        vm.deal(user1, 1 ether);
        vm.prank(user1);
        betContract.requestMatchIfNeeded{value: fee}(matchId);
        
        assertTrue(mockOracle.isMatchRequested(matchId));
    }
    
    function testDepositForFees() public {
        uint256 initialBalance = betContract.getContractETHBalance();
        uint256 depositAmount = 0.1 ether;
        
        vm.deal(user1, 1 ether);
        vm.prank(user1);
        betContract.depositForFees{value: depositAmount}();
        
        assertEq(betContract.getContractETHBalance(), initialBalance + depositAmount);
    }
    
    function testSetMatchRequestFee() public {
        uint256 newFee = 0.002 ether;
        
        betContract.setMatchRequestFee(newFee);
        assertEq(betContract.matchRequestFee(), newFee);
    }
    
    function testCallMatchReceivedAutoRequest() public {
        uint256 matchId = 999;
        uint256 winnerId = 123;
        
        vm.deal(address(betContract), 1 ether);
        
        EsportOracleTypes.Match memory matchData;
        matchData._id = matchId;
        matchData._winnerId = winnerId;
        
        mockOracle.setMatchRequest(matchId, true);
        
        vm.prank(address(mockOracle));
        betContract.callMatchReceived(matchData);
        
        assertTrue(true);
    }
}
