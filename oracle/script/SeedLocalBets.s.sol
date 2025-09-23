// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import "forge-std/console.sol";

import "../src/BetContract.sol";
import "../test/BetContractTest.t.sol";

contract SeedLocalBets is Script {
    function run() external {
        address betContractAddress = vm.envAddress("BET_CONTRACT_ADDRESS");
        address tokenAddress = vm.envAddress("TOKEN_ADDRESS");

        uint256 deployerKey = vm.envUint("PRIVATE_KEY");
        uint256 bettorOneKey = vm.envUint("BETTOR1_PRIVATE_KEY");
        uint256 bettorTwoKey = vm.envUint("BETTOR2_PRIVATE_KEY");

        BetContract betContract = BetContract(payable(betContractAddress));
        MockERC20 token = MockERC20(tokenAddress);

        address bettorOne = vm.addr(bettorOneKey);
        address bettorTwo = vm.addr(bettorTwoKey);

        console.log("=== Seeding local BetContract ===");
        console.log("BetContract:", betContractAddress);
        console.log("Token:", tokenAddress);
        console.log("Bettor #1:", bettorOne);
        console.log("Bettor #2:", bettorTwo);

        uint256 deadlineOne = block.timestamp + 2 hours;
        uint256 deadlineTwo = block.timestamp + 4 hours;

        // These IDs come from PandaScore upcoming matches so the frontend can
        // cross-reference live data
        uint256 matchOneId = vm.envOr("SEED_MATCH_ONE_ID", uint256(1239584));
        uint256 matchTwoId = vm.envOr("SEED_MATCH_TWO_ID", uint256(1239454));

        uint256 matchOneTeam1Id = vm.envOr("SEED_MATCH_ONE_TEAM1", uint256(136846));
        uint256 matchOneTeam2Id = vm.envOr("SEED_MATCH_ONE_TEAM2", uint256(136516));

        uint256 matchTwoTeam1Id = vm.envOr("SEED_MATCH_TWO_TEAM1", uint256(127014));
        uint256 matchTwoTeam2Id = vm.envOr("SEED_MATCH_TWO_TEAM2", uint256(131010));

        // Mint tokens and create bets as deployer
        vm.startBroadcast(deployerKey);
        token.mint(bettorOne, 3_000 ether);
        token.mint(bettorTwo, 3_000 ether);

        betContract.createBet(
            "CS2 - Prestige vs Basement Boys",
            matchOneTeam1Id,
            matchOneTeam2Id,
            deadlineOne,
            matchOneId
        );

        betContract.createBet(
            "CS2 - Sinners vs Sashi",
            matchTwoTeam1Id,
            matchTwoTeam2Id,
            deadlineTwo,
            matchTwoId
        );
        vm.stopBroadcast();

        // Bettor #1 stakes on team 1 for both bets
        vm.startBroadcast(bettorOneKey);
        token.approve(betContractAddress, type(uint256).max);
        betContract.placeBet(0, 1, 250 ether);
        betContract.placeBet(1, 1, 150 ether);
        vm.stopBroadcast();

        // Bettor #2 stakes on team 2 for both bets
        vm.startBroadcast(bettorTwoKey);
        token.approve(betContractAddress, type(uint256).max);
        betContract.placeBet(0, 2, 300 ether);
        betContract.placeBet(1, 2, 200 ether);
        vm.stopBroadcast();

        console.log("Seed complete. Two bets funded with sample stakes.");
    }
}
