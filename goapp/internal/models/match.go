package models

import (
	"math/big"
	"time"
)

type Match struct {
	ID        int        `json:"id"`
	Opponents []Opponent `json:"opponents"`
	Games     []Game     `json:"games"`
	Results   []Result   `json:"results"`
	WinnerID  int        `json:"winner_id"`
	BeginAt   string     `json:"begin_at"`
}

type Game struct {
	ID       int    `json:"id"`
	Finished bool   `json:"finished"`
	Winner   Winner `json:"winner"`
}

type Winner struct {
	ID *int `json:"id"`
}

type Opponent struct {
	Opponent Team `json:"opponent"`
}

type Team struct {
	Acronym string `json:"acronym"`
	ID      int    `json:"id"`
	Name    string `json:"name"`
}

type Result struct {
	Score  int `json:"score"`
	TeamID int `json:"team_id"`
}

type ContractMatch struct {
	ID        *big.Int         `json:"_id"`
	Opponents []ContractTeam   `json:"_opponents"`
	Games     []ContractGame   `json:"_game"`
	Results   []ContractResult `json:"_result"`
	WinnerID  *big.Int         `json:"_winnerId"`
	BeginAt   *big.Int         `json:"_beginAt"`
}

type ContractTeam struct {
	Acronym string   `json:"_acronym"`
	ID      *big.Int `json:"_id"`
	Name    string   `json:"_name"`
}

type ContractGame struct {
	ID       *big.Int `json:"_id"`
	Finished bool     `json:"_finished"`
	WinnerID *big.Int `json:"_winnerId"`
}

type ContractResult struct {
	Score  uint8    `json:"_score"`
	TeamID *big.Int `json:"_teamId"`
}

func (m *Match) ToContractMatch() ContractMatch {
	contractMatch := ContractMatch{
		ID:       big.NewInt(int64(m.ID)),
		WinnerID: big.NewInt(int64(m.WinnerID)),
	}

	if t, err := time.Parse(time.RFC3339, m.BeginAt); err == nil {
		contractMatch.BeginAt = big.NewInt(t.Unix())
	} else {
		contractMatch.BeginAt = big.NewInt(0)
	}

	for _, opp := range m.Opponents {
		team := opp.Opponent
		contractMatch.Opponents = append(contractMatch.Opponents, ContractTeam{
			Acronym: team.Acronym,
			ID:      big.NewInt(int64(team.ID)),
			Name:    team.Name,
		})
	}

	for _, game := range m.Games {
		winnerId := big.NewInt(0)
		if game.Winner.ID != nil {
			winnerId = big.NewInt(int64(*game.Winner.ID))
		}

		contractMatch.Games = append(contractMatch.Games, ContractGame{
			ID:       big.NewInt(int64(game.ID)),
			Finished: game.Finished,
			WinnerID: winnerId,
		})
	}

	for _, result := range m.Results {
		contractMatch.Results = append(contractMatch.Results, ContractResult{
			Score:  uint8(result.Score),
			TeamID: big.NewInt(int64(result.TeamID)),
		})
	}
	return contractMatch
}
