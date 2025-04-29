package models

type Match struct {
    ID           int       `json:"id"`
    Opponents    []Opponent `json:"opponents"`
    Games        []Game    `json:"games"`
    Results      []Result  `json:"results"`
    WinnerID     int       `json:"winner_id"`
    BeginAt      string    `json:"begin_at"`
}

type Game struct {
    ID       int    `json:"id"`
    Finished bool   `json:"finished"`
    Winner   Winner `json:"winner"`
}

type Winner struct {
    ID   *int  `json:"id"`
}

type Opponent struct {
    Opponent Team `json:"opponent"`
}

type Team struct {
    Acronym  string `json:"acronym"`
    ID       int    `json:"id"`
    Name     string `json:"name"`
}

type Result struct {
    Score  int `json:"score"`
    TeamID int `json:"team_id"`
}