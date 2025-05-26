package models

import (
	"math/big"
	"src/internal/contract"
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

func (m *Match) ToContractMatch() contract.EsportOracleMatch {
	contractMatch := contract.EsportOracleMatch{
		Id:       big.NewInt(int64(m.ID)),
		WinnerId: big.NewInt(int64(m.WinnerID)),
	}

	if t, err := time.Parse(time.RFC3339, m.BeginAt); err == nil {
		contractMatch.BeginAt = big.NewInt(t.Unix())
	} else {
		contractMatch.BeginAt = big.NewInt(0)
	}

	for _, opp := range m.Opponents {
		team := opp.Opponent
		contractMatch.Opponents = append(contractMatch.Opponents, contract.EsportOracleOpponents{
			Acronym: team.Acronym,
			Id:      big.NewInt(int64(team.ID)),
			Name:    team.Name,
		})
	}

	for _, game := range m.Games {
		winnerId := big.NewInt(0)
		if game.Winner.ID != nil {
			winnerId = big.NewInt(int64(*game.Winner.ID))
		}

		contractMatch.Game = append(contractMatch.Game, contract.EsportOracleGames{
			Id:       big.NewInt(int64(game.ID)),
			Finished: game.Finished,
			WinnerId: winnerId,
		})
	}

	for _, result := range m.Results {
		contractMatch.Result = append(contractMatch.Result, contract.EsportOracleResult{
			Score:  uint8(result.Score),
			TeamId: big.NewInt(int64(result.TeamID)),
		})
	}
	return contractMatch
}
