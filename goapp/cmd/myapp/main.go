package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

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

func main() {
	url := "https://api.pandascore.co/csgo/matches/upcoming?filter[future]=true&token=uqql5zPjp28p7kbgS_Uoc3TUS5wUV6fjTcTpuB1u82uCe7X6A2s"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Erreur lors de la création de la requête:", err)
		return
	}
	req.Header.Add("Authorization", "Bearer uqql5zPjp28p7kbgS_Uoc3TUS5wUV6fjTcTpuB1u82uCe7X6A2s")
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi de la requête:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse:", err)
		return
	}

	var matches []Match
	if err := json.Unmarshal(body, &matches); err != nil {
		fmt.Println("Erreur lors de l'unmarshal du JSON:", err)
		return
	}

	for _, match := range matches {
		matchOutput := map[string]interface{}{
			"begin_at":       match.BeginAt,
			"draw":           match.Draw,
			"forfeit":        match.Forfeit,
			"games":          match.Games,
			"id":             match.ID,
			"name":           match.Name,
			"number_of_games": match.NumberOfGames,
			"opponents":      match.Opponents,
			"results":        match.Results,
			"winner_id":      match.WinnerID,
		}

		prettyJSON, err := json.MarshalIndent(matchOutput, "", "  ")
		if err != nil {
			fmt.Println("Erreur lors de l'indentation du JSON:", err)
			return
		}

		fmt.Println(string(prettyJSON))
	}

	outputFile := "high_level_upcoming_matches.json"
	prettyJSON, err := json.MarshalIndent(matches, "", "  ")
	if err != nil {
		fmt.Println("Erreur lors de l'indentation du JSON:", err)
		return
	}
	if err := os.WriteFile(outputFile, prettyJSON, 0644); err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier:", err)
		return
	}
	fmt.Printf("Les données ont été sauvegardées dans le fichier %s\n", outputFile)
}
