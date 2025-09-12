package backend

import "math/big"

// Bet represents a single bet's on-chain data in a convenient Go form.
type Bet struct {
    Description string
    Team1Id     *big.Int
    Team2Id     *big.Int
    Deadline    *big.Int
    Team1Pool   *big.Int
    Team2Pool   *big.Int
    WinningTeam uint8
    Resolved    bool
    Creator     string
    MatchId     *big.Int
}

// UserBet represents a user's participation in a given bet.
type UserBet struct {
    Amount     *big.Int
    TeamChosen uint8
    Claimed    bool
}

// BetHistoryResponse wraps a bet with its ID and optionally the user's bet info.
type BetHistoryResponse struct {
    BetId   *big.Int `json:"betId"`
    Bet     Bet      `json:"bet"`
    UserBet *UserBet `json:"userBet,omitempty"`
}
