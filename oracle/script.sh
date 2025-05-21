#!/bin/bash

CONTRACT_ADDRESS="0x2aB32CFC773515C177A306CcA850212f68A65FE0"
PRIVATE_KEY_1=$(grep '^PRIVATE_KEY_1=' .env | cut -d'=' -f2-)
PRIVATE_KEY_2=$(grep '^PRIVATE_KEY_2=' .env | cut -d'=' -f2-)
PRIVATE_KEY_3=$(grep '^PRIVATE_KEY_3=' .env | cut -d'=' -f2-)
RPC_URL=$(grep '^SEPOLIA_RPC_URL=' .env | cut -d'=' -f2-)
AMOUNT="1000000000000000"

MATCH_DATA="([(1,[\\\"TEAM1\\\",1,\\\"Team One\\\"],[\\\"TEAM2\\\",2,\\\"Team Two\\\"]],[(1,true,1)],[(3,1),(1,2)],1,1747900000))"

echo "Staking..."

cast send --private-key $PRIVATE_KEY_1 --rpc-url $RPC_URL --value $AMOUNT $CONTRACT_ADDRESS "addFundToStaking()" --gas-limit 200000
cast send --private-key $PRIVATE_KEY_2 --rpc-url $RPC_URL --value $AMOUNT $CONTRACT_ADDRESS "addFundToStaking()" --gas-limit 200000
cast send --private-key $PRIVATE_KEY_3 --rpc-url $RPC_URL --value $AMOUNT $CONTRACT_ADDRESS "addFundToStaking()" --gas-limit 200000

echo "Encoding match..."

ENCODED_DATA=$(cast abi-encode "f(uint256,string,uint256,string,uint256,bool,uint256,uint8,uint256,uint256,uint256)" \
                   "1" "TEAM1" "1" "Team One" "1" "true" "1" "3" "1" "1" "1747900000")

echo "Sending match..."

cast send --private-key $PRIVATE_KEY_1 --rpc-url $RPC_URL $CONTRACT_ADDRESS "handleNewMatches((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256)[])" "[(1,[(\"TEAM1\",1,\"Team One\"),(\"TEAM2\",2,\"Team Two\")],[(1,true,1)],[(3,1),(1,2)],1,1747900000)]" --gas-limit 1000000
sleep 2
cast send --private-key $PRIVATE_KEY_2 --rpc-url $RPC_URL $CONTRACT_ADDRESS "handleNewMatches((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256)[])" "[(1,[(\"TEAM1\",1,\"Team One\"),(\"TEAM2\",2,\"Team Two\")],[(1,true,1)],[(3,1),(1,2)],1,1747900000)]" --gas-limit 1000000
sleep 2
cast send --private-key $PRIVATE_KEY_3 --rpc-url $RPC_URL $CONTRACT_ADDRESS "handleNewMatches((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256)[])" "[(1,[(\"TEAM1\",1,\"Team One\"),(\"TEAM2\",2,\"Team Two\")],[(1,true,1)],[(3,1),(1,2)],1,1747900000)]" --gas-limit 1000000

echo "Checking match..."
sleep 5

cast call --rpc-url $RPC_URL $CONTRACT_ADDRESS "getMatchById(uint256)" 1

echo "Done."
