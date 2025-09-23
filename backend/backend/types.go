package backend

import "math/big"

// Bet represents a single bet's on-chain data in a convenient Go form.
type Bet struct {
	Description string   `json:"description"`
	Team1Id     *big.Int `json:"team1Id"`
	Team2Id     *big.Int `json:"team2Id"`
	Deadline    *big.Int `json:"deadline"`
	Team1Pool   *big.Int `json:"team1Pool"`
	Team2Pool   *big.Int `json:"team2Pool"`
	WinningTeam uint8    `json:"winningTeam"`
	Resolved    bool     `json:"resolved"`
	Creator     string   `json:"creator"`
	MatchId     *big.Int `json:"matchId"`
}

// UserBet represents a user's participation in a given bet.
type UserBet struct {
	Amount     *big.Int `json:"amount"`
	TeamChosen uint8    `json:"teamChosen"`
	Claimed    bool     `json:"claimed"`
}

// BetHistoryResponse wraps a bet with its ID and optionally the user's bet info.
type BetHistoryResponse struct {
	BetId   *big.Int `json:"betId"`
	Bet     Bet      `json:"bet"`
	UserBet *UserBet `json:"userBet,omitempty"`
}
