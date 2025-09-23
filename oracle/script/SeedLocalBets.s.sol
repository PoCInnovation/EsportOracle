// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

import "../src/BetContract.sol";
import "../test/BetContractTest.t.sol";

contract SeedLocalBets is Script {
    using stdJson for string;

    struct SeedBetConfig {
        string description;
        uint256 matchId;
        uint256 team1Id;
        uint256 team2Id;
        uint256 deadline;
        uint256 deadlineOffsetSeconds;
    }

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

        (SeedBetConfig[] memory betsToSeed, string memory source) = _loadSeedConfig();
        require(betsToSeed.length > 0, "SeedLocalBets: no seed data provided");

        uint256 defaultDeadlineOffset = vm.envOr("SEED_DEFAULT_DEADLINE_OFFSET", uint256(2 hours));
        uint256 bettorOneStake = vm.envOr("SEED_BETTOR_ONE_STAKE", uint256(250 ether));
        uint256 bettorTwoStake = vm.envOr("SEED_BETTOR_TWO_STAKE", uint256(300 ether));
        uint8 bettorOneTeam = uint8(vm.envOr("SEED_BETTOR_ONE_TEAM", uint256(1)));
        uint8 bettorTwoTeam = uint8(vm.envOr("SEED_BETTOR_TWO_TEAM", uint256(2)));

        console.log("=== Seeding local BetContract ===");
        console.log("Seed source:", bytes(source).length == 0 ? "default config" : source);
        console.log("BetContract:", betContractAddress);
        console.log("Token:", tokenAddress);
        console.log("Bettor #1:", bettorOne);
        console.log("Bettor #2:", bettorTwo);

        uint256 betCountBefore = betContract.getBetCount();

        vm.startBroadcast(deployerKey);
        token.mint(bettorOne, 3_000 ether);
        token.mint(bettorTwo, 3_000 ether);

        for (uint256 i = 0; i < betsToSeed.length; i++) {
            SeedBetConfig memory config = betsToSeed[i];
            uint256 deadline = _resolveDeadline(config, defaultDeadlineOffset);

            betContract.createBet(
                config.description,
                config.team1Id,
                config.team2Id,
                deadline,
                config.matchId
            );

            console.log("-- Created bet", i);
            console.log("   description:", config.description);
            console.log("   matchId:", config.matchId);
            console.log("   teams:", config.team1Id, config.team2Id);
            console.log("   deadline:", deadline);
        }
        vm.stopBroadcast();

        uint256 betCountAfter = betContract.getBetCount();
        require(
            betCountAfter == betCountBefore + betsToSeed.length,
            "SeedLocalBets: bet count mismatch"
        );

        uint256 firstNewBetId = betCountBefore;
        uint256 lastNewBetId = betCountAfter;

        if (_shouldPlaceStake(bettorOneStake, bettorOneTeam)) {
            vm.startBroadcast(bettorOneKey);
            token.approve(betContractAddress, type(uint256).max);
            for (uint256 i = firstNewBetId; i < lastNewBetId; i++) {
                betContract.placeBet(i, bettorOneTeam, bettorOneStake);
            }
            vm.stopBroadcast();
        }

        if (_shouldPlaceStake(bettorTwoStake, bettorTwoTeam)) {
            vm.startBroadcast(bettorTwoKey);
            token.approve(betContractAddress, type(uint256).max);
            for (uint256 i = firstNewBetId; i < lastNewBetId; i++) {
                betContract.placeBet(i, bettorTwoTeam, bettorTwoStake);
            }
            vm.stopBroadcast();
        }

        console.log("Seed complete. bets funded with sample stakes:", lastNewBetId - firstNewBetId);
    }

    function _loadSeedConfig() private view returns (SeedBetConfig[] memory seeds, string memory source) {
        string memory rawJson = vm.envOr("SEED_BETS_JSON", string(""));

        if (bytes(rawJson).length == 0) {
            string memory path = vm.envOr("SEED_BETS_FILE", string(""));
            if (bytes(path).length == 0) {
                path = "script/seed_data/sample_bets.json";
            }
            rawJson = vm.readFile(path);
            source = path;
        } else {
            source = "SEED_BETS_JSON";
        }

        require(bytes(rawJson).length != 0, "SeedLocalBets: empty config");

        seeds = _decodeSeedConfig(rawJson);
    }

    function _decodeSeedConfig(string memory rawJson) private view returns (SeedBetConfig[] memory seeds) {
        uint256 count = 0;
        while (rawJson.keyExists(string.concat("[", vm.toString(count), "]"))) {
            count++;
        }

        require(count > 0, "SeedLocalBets: invalid seed array");

        seeds = new SeedBetConfig[](count);
        for (uint256 i = 0; i < count; i++) {
            string memory prefix = string.concat("[", vm.toString(i), "]");

            seeds[i].description = rawJson.readStringOr(string.concat(prefix, ".description"), "");
            seeds[i].matchId = rawJson.readUint(string.concat(prefix, ".matchId"));
            seeds[i].team1Id = rawJson.readUint(string.concat(prefix, ".team1Id"));
            seeds[i].team2Id = rawJson.readUint(string.concat(prefix, ".team2Id"));
            seeds[i].deadline = rawJson.readUintOr(string.concat(prefix, ".deadline"), 0);
            seeds[i].deadlineOffsetSeconds = rawJson.readUintOr(
                string.concat(prefix, ".deadlineOffsetSeconds"),
                0
            );
        }
    }

    function _resolveDeadline(SeedBetConfig memory config, uint256 defaultOffset)
        private
        view
        returns (uint256)
    {
        if (config.deadline > block.timestamp) {
            return config.deadline;
        }

        uint256 offset = config.deadlineOffsetSeconds != 0 ? config.deadlineOffsetSeconds : defaultOffset;
        if (offset == 0) {
            offset = 2 hours;
        }

        return block.timestamp + offset;
    }

    function _shouldPlaceStake(uint256 stake, uint8 team) private pure returns (bool) {
        return stake > 0 && (team == 1 || team == 2);
    }
}
