package backend

import (
	"encoding/json"
	"context"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"net/url"
	"time"

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

type pandaTeam struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

const pandaBase = "https://api.pandascore.co"
//Rajouter /csgo et bien préciser les IDs présents.

func fetchTeamsInfo(ctx context.Context, ids []int64) (map[int64]pandaTeam, error) {
	result := make(map[int64]pandaTeam)
	if len(ids) == 0 {
		return result, nil
	}
	if PandaScoreAPIToken == "" {
		return result, nil
	}

	idsStr := make([]string, 0, len(ids))
	seen := make(map[int64]struct{})
	for _, id := range ids {
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		idsStr = append(idsStr, fmt.Sprintf("%d", id))
	}

	u, _ := url.Parse(pandaBase + "/teams")
	q := u.Query()
	q.Set("filter[id]", strings.Join(idsStr, ","))
	q.Set("per_page", "100")

	u.RawQuery = q.Encode()
	fmt.Printf("[PS] URL: %s\n", u.String()) 


	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	req.Header.Set("Authorization", "Bearer "+PandaScoreAPIToken)
	req.Header.Set("Accept", "application/json")

	httpClient := &http.Client{Timeout: 8 * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Printf("[PS] HTTP error: %v\n", err)
		return result, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("pandascore teams HTTP %d", resp.StatusCode)
	}

	var teams []pandaTeam
	if err := json.NewDecoder(resp.Body).Decode(&teams); err != nil {
		fmt.Printf("[PS] Decode error: %v\n", err)
		return result, err
	}

	for _, t := range teams {
		result[t.ID] = t
	}
	return result, nil
}

func biToString(x *big.Int) string {
	if x == nil {
		return "0"
	}
	return x.String()
}

func biToInt64(x *big.Int) int64 {
	if x == nil {
		return 0
	}
	if x.IsInt64() {
		return x.Int64()
	}
	return new(big.Int).Set(x).Int64()
}

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
	var allTeamIDs []int64

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
			Description: betResults[0].(string),
			Team1Id:     biToInt64(betResults[1].(*big.Int)),
			Team2Id:     biToInt64(betResults[2].(*big.Int)),
			Deadline:    biToInt64(betResults[3].(*big.Int)),
			Team1Pool:   biToString(betResults[4].(*big.Int)),
			Team2Pool:   biToString(betResults[5].(*big.Int)),
			WinningTeam: int(betResults[6].(uint8)),
			Resolved:    betResults[7].(bool),
			Creator:     betResults[8].(common.Address).Hex(),
			MatchId:     biToInt64(betResults[9].(*big.Int)),
		}

		//Collect IDs
		allTeamIDs = append(allTeamIDs, bet.Team1Id, bet.Team2Id)

		bets = append(bets, BetHistoryResponse{
			BetId: int64(i),
			Bet:   bet,
		})
	}

	teamsByID, err := fetchTeamsInfo(r.Context(), allTeamIDs)
	if err != nil {
		fmt.Printf("fetchTeamsInfo error: %v\n", err)
	}

	fmt.Printf("[PS] Enrichment map size: %d\n", len(teamsByID))
		for i := range bets {
			if t1, ok := teamsByID[bets[i].Bet.Team1Id]; ok {
				bets[i].Bet.Team1Name = t1.Name
				bets[i].Bet.Team1Logo = t1.ImageURL
			}
			if t2, ok := teamsByID[bets[i].Bet.Team2Id]; ok {
				bets[i].Bet.Team2Name = t2.Name
				bets[i].Bet.Team2Logo = t2.ImageURL
			}
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
	var allTeamIDs []int64

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
			Description: betResults[0].(string),
			Team1Id:     biToInt64(betResults[1].(*big.Int)),
			Team2Id:     biToInt64(betResults[2].(*big.Int)),
			Deadline:    biToInt64(betResults[3].(*big.Int)),
			Team1Pool:   biToString(betResults[4].(*big.Int)),
			Team2Pool:   biToString(betResults[5].(*big.Int)),
			WinningTeam: int(betResults[6].(uint8)),
			Resolved:    betResults[7].(bool),
			Creator:     betResults[8].(common.Address).Hex(),
			MatchId:     biToInt64(betResults[9].(*big.Int)),
		}

		var userBet *UserBet
        if userBetResults != nil {
            amount := userBetResults[0].(*big.Int)
            teamChosen := userBetResults[1].(uint8)
            claimed := userBetResults[2].(bool)

            if amount != nil && amount.Sign() != 0 {
                userBet = &UserBet{
                    Amount:     biToString(amount),
                    TeamChosen: int(teamChosen),              
                    Claimed:    claimed,
                }
            }
        }

		allTeamIDs = append(allTeamIDs, bet.Team1Id, bet.Team2Id)


		userBets = append(userBets, BetHistoryResponse{
			BetId:   int64(i),
			Bet:     bet,
			UserBet: userBet,
		})
	}

	teamsByID, err := fetchTeamsInfo(r.Context(), allTeamIDs)
	if err != nil {
		fmt.Printf("fetchTeamsInfo error: %v\n", err)
	}
	for i := range userBets {
		if t1, ok := teamsByID[userBets[i].Bet.Team1Id]; ok {
			userBets[i].Bet.Team1Name = t1.Name
			userBets[i].Bet.Team1Logo = t1.ImageURL
		}
		if t2, ok := teamsByID[userBets[i].Bet.Team2Id]; ok {
			userBets[i].Bet.Team2Name = t2.Name
			userBets[i].Bet.Team2Logo = t2.ImageURL
		}
	}
	fmt.Printf("[USER] returning %d bets for %s\n", len(userBets), userAddress)
	

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userBets)

	fmt.Printf("Successfully fetched bet history for user %s using simple approach\n", userAddress)
}