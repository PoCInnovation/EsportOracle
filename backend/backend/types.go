package backend

// Bet represents a single bet's on-chain data in a convenient Go form.
type Bet struct {
    Description string `json:"description"`
	Team1Id     int64  `json:"team1Id"`
	Team2Id     int64  `json:"team2Id"`
	Deadline    int64  `json:"deadline"`
	Team1Pool   string `json:"team1Pool"`
	Team2Pool   string `json:"team2Pool"`
	WinningTeam int    `json:"winningTeam"`
	Resolved    bool   `json:"resolved"`
	Creator     string `json:"creator"`
	MatchId     int64  `json:"matchId"`

	Team1Name string `json:"team1Name,omitempty"`
	Team2Name string `json:"team2Name,omitempty"`
	Team1Logo string `json:"team1Logo,omitempty"`
	Team2Logo string `json:"team2Logo,omitempty"`
}

// UserBet represents a user's participation in a given bet.
type UserBet struct {
    Amount     string `json:"amount"`
	TeamChosen int    `json:"teamChosen"`
	Claimed    bool   `json:"claimed"`
}

// BetHistoryResponse wraps a bet with its ID and optionally the user's bet info.
type BetHistoryResponse struct {
    BetId   int64 `json:"betId"`
    Bet     Bet      `json:"bet"`
    UserBet *UserBet `json:"userBet,omitempty"`
}
