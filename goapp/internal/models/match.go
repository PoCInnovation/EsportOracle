package models

type Match struct {
    BeginAt      string    `json:"begin_at"`
    Draw         bool      `json:"draw"`
    Forfeit      bool      `json:"forfeit"`
    Games        []Game    `json:"games"`
    ID           int       `json:"id"`
    Name         string    `json:"name"`
    NumberOfGames int      `json:"number_of_games"`
    Opponents    []Opponent `json:"opponents"`
    Results      []Result  `json:"results"`
    WinnerID     int       `json:"winner_id"`
}

type Game struct {
    ID       int    `json:"id"`
    Finished bool   `json:"finished"`
    Winner   Winner `json:"winner"`
}

type Winner struct {
    ID   *int  `json:"id"`
    Type string `json:"type"`
}

type Opponent struct {
    Opponent Team `json:"opponent"`
}

type Team struct {
    Acronym  string `json:"acronym"`
    ID       int    `json:"id"`
    ImageURL string `json:"image_url"`
    Name     string `json:"name"`
    Slug     string `json:"slug"`
}

type Result struct {
    Score  int `json:"score"`
    TeamID int `json:"team_id"`
}