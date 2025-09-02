package backend

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

// Simple ABI for just the functions we need
const simpleBetABI = `[
	{
		"type": "function",
		"name": "getBetCount",
		"inputs": [],
		"outputs": [{"name": "", "type": "uint256", "internalType": "uint256"}],
		"stateMutability": "view"
	},
	{
		"type": "function",
		"name": "CurrentBets",
		"inputs": [{"name": "", "type": "uint256", "internalType": "uint256"}],
		"outputs": [
			{"name": "description", "type": "string", "internalType": "string"},
			{"name": "team1Id", "type": "uint256", "internalType": "uint256"},
			{"name": "team2Id", "type": "uint256", "internalType": "uint256"},
			{"name": "deadline", "type": "uint256", "internalType": "uint256"},
			{"name": "team1Pool", "type": "uint256", "internalType": "uint256"},
			{"name": "team2Pool", "type": "uint256", "internalType": "uint256"},
			{"name": "winningTeam", "type": "uint8", "internalType": "uint8"},
			{"name": "resolved", "type": "bool", "internalType": "bool"},
			{"name": "creator", "type": "address", "internalType": "address"},
			{"name": "matchId", "type": "uint256", "internalType": "uint256"}
		],
		"stateMutability": "view"
	},
	{
		"type": "function",
		"name": "userBets",
		"inputs": [
			{"name": "", "type": "address", "internalType": "address"},
			{"name": "", "type": "uint256", "internalType": "uint256"}
		],
		"outputs": [
			{"name": "amount", "type": "uint256", "internalType": "uint256"},
			{"name": "teamChosen", "type": "uint8", "internalType": "uint8"},
			{"name": "claimed", "type": "bool", "internalType": "bool"}
		],
		"stateMutability": "view"
	}
]`

func GetAllBetsSimple(w http.ResponseWriter, r *http.Request) {
	if ethereumRPCURL == "" || betContractAddress == "" {
		http.Error(w, "Ethereum RPC URL or Bet contract address not configured", http.StatusInternalServerError)
		return
	}

	client, err := ethclient.Dial(ethereumRPCURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to connect to Ethereum client: %v", err), http.StatusInternalServerError)
		return
	}
	defer client.Close()

	contractAbi, err := abi.JSON(strings.NewReader(simpleBetABI))
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse contract ABI: %v", err), http.StatusInternalServerError)
		return
	}

	contractAddress := common.HexToAddress(betContractAddress)

	// Get bet count using the contract directly
	betCountCallData, err := contractAbi.Pack("getBetCount")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to pack getBetCount call: %v", err), http.StatusInternalServerError)
		return
	}

	betCountResult, err := client.CallContract(r.Context(), ethereum.CallMsg{
		To:   &contractAddress,
		Data: betCountCallData,
	}, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to call getBetCount: %v", err), http.StatusInternalServerError)
		return
	}

	results, err := contractAbi.Unpack("getBetCount", betCountResult)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to unpack bet count: %v", err), http.StatusInternalServerError)
		return
	}

	betCount := results[0].(*big.Int)

	var bets []BetHistoryResponse

	// Get all bets using CurrentBets function (which is simpler than getBet)
	for i := int64(0); i < betCount.Int64(); i++ {
		betId := big.NewInt(i)
		
		// Call CurrentBets(i) instead of getBet(i)
		currentBetsCallData, err := contractAbi.Pack("CurrentBets", betId)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to pack CurrentBets call for bet %d: %v", i, err), http.StatusInternalServerError)
			return
		}

		currentBetsResult, err := client.CallContract(r.Context(), ethereum.CallMsg{
			To:   &contractAddress,
			Data: currentBetsCallData,
		}, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to call CurrentBets for bet %d: %v", i, err), http.StatusInternalServerError)
			return
		}

		// Unpack the results - CurrentBets returns individual fields, not a struct
		betResults, err := contractAbi.Unpack("CurrentBets", currentBetsResult)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to unpack bet data for bet %d: %v", i, err), http.StatusInternalServerError)
			return
		}

		if len(betResults) < 10 {
			http.Error(w, fmt.Sprintf("Unexpected number of fields in bet %d: got %d, expected 10", i, len(betResults)), http.StatusInternalServerError)
			return
		}

		bet := Bet{
			Description:  betResults[0].(string),
			Team1Id:      betResults[1].(*big.Int),
			Team2Id:      betResults[2].(*big.Int),
			Deadline:     betResults[3].(*big.Int),
			Team1Pool:    betResults[4].(*big.Int),
			Team2Pool:    betResults[5].(*big.Int),
			WinningTeam:  betResults[6].(uint8),
			Resolved:     betResults[7].(bool),
			Creator:      betResults[8].(common.Address).Hex(),
			MatchId:      betResults[9].(*big.Int),
		}

		bets = append(bets, BetHistoryResponse{
			BetId: betId,
			Bet:   bet,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bets)

	fmt.Println("Successfully fetched all bets from BetContract using simple approach")
}

func GetUserBetHistorySimple(w http.ResponseWriter, r *http.Request) {
	userAddress := mux.Vars(r)["userAddress"]
	
	if !common.IsHexAddress(userAddress) {
		http.Error(w, "Invalid user address", http.StatusBadRequest)
		return
	}

	if ethereumRPCURL == "" || betContractAddress == "" {
		http.Error(w, "Ethereum RPC URL or Bet contract address not configured", http.StatusInternalServerError)
		return
	}

	client, err := ethclient.Dial(ethereumRPCURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to connect to Ethereum client: %v", err), http.StatusInternalServerError)
		return
	}
	defer client.Close()

	contractAbi, err := abi.JSON(strings.NewReader(simpleBetABI))
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse contract ABI: %v", err), http.StatusInternalServerError)
		return
	}

	contractAddress := common.HexToAddress(betContractAddress)
	userAddr := common.HexToAddress(userAddress)

	// Get bet count
	betCountCallData, err := contractAbi.Pack("getBetCount")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to pack getBetCount call: %v", err), http.StatusInternalServerError)
		return
	}

	betCountResult, err := client.CallContract(r.Context(), ethereum.CallMsg{
		To:   &contractAddress,
		Data: betCountCallData,
	}, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to call getBetCount: %v", err), http.StatusInternalServerError)
		return
	}

	results, err := contractAbi.Unpack("getBetCount", betCountResult)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to unpack bet count: %v", err), http.StatusInternalServerError)
		return
	}

	betCount := results[0].(*big.Int)

	var userBets []BetHistoryResponse

	// Check all bets for this user
	for i := int64(0); i < betCount.Int64(); i++ {
		betId := big.NewInt(i)
		
		// First check if user has a bet for this ID
		userBetsCallData, err := contractAbi.Pack("userBets", userAddr, betId)
		if err != nil {
			continue // Skip this bet if we can't pack the call
		}

		userBetsResult, err := client.CallContract(r.Context(), ethereum.CallMsg{
			To:   &contractAddress,
			Data: userBetsCallData,
		}, nil)
		if err != nil {
			continue // Skip this bet if call fails
		}

		userBetResults, err := contractAbi.Unpack("userBets", userBetsResult)
		if err != nil {
			continue // Skip this bet if unpacking fails
		}

		if len(userBetResults) < 3 {
			continue // Skip if not enough fields
		}

		amount := userBetResults[0].(*big.Int)
		
		// Skip if user didn't bet (amount is 0)
		if amount.Cmp(big.NewInt(0)) == 0 {
			continue
		}

		// Get bet details
		currentBetsCallData, err := contractAbi.Pack("CurrentBets", betId)
		if err != nil {
			continue // Skip this bet if we can't pack the call
		}

		currentBetsResult, err := client.CallContract(r.Context(), ethereum.CallMsg{
			To:   &contractAddress,
			Data: currentBetsCallData,
		}, nil)
		if err != nil {
			continue // Skip this bet if call fails
		}

		betResults, err := contractAbi.Unpack("CurrentBets", currentBetsResult)
		if err != nil {
			continue // Skip this bet if unpacking fails
		}

		if len(betResults) < 10 {
			continue // Skip if not enough fields
		}

		bet := Bet{
			Description:  betResults[0].(string),
			Team1Id:      betResults[1].(*big.Int),
			Team2Id:      betResults[2].(*big.Int),
			Deadline:     betResults[3].(*big.Int),
			Team1Pool:    betResults[4].(*big.Int),
			Team2Pool:    betResults[5].(*big.Int),
			WinningTeam:  betResults[6].(uint8),
			Resolved:     betResults[7].(bool),
			Creator:      betResults[8].(common.Address).Hex(),
			MatchId:      betResults[9].(*big.Int),
		}

		userBet := UserBet{
			Amount:     amount,
			TeamChosen: userBetResults[1].(uint8),
			Claimed:    userBetResults[2].(bool),
		}

		userBets = append(userBets, BetHistoryResponse{
			BetId:   betId,
			Bet:     bet,
			UserBet: &userBet,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userBets)

	fmt.Printf("Successfully fetched bet history for user %s using simple approach\n", userAddress)
}