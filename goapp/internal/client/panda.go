package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"src/internal/models"
)

type PandaScoreClient struct {
	APIToken string
	BaseURL  string
}

func NewPandaScoreClient(apiToken string) *PandaScoreClient {
	return &PandaScoreClient{
		APIToken: apiToken,
		BaseURL:  "https://api.pandascore.co",
	}
}

func (c *PandaScoreClient) GetMatchByID(matchID int) (*models.Match, error) {
	url := fmt.Sprintf("%s/matches/%d", c.BaseURL, matchID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest(\"GET\", url, nil): %w", err)
	}

	req.Header.Add("Authorization", "Bearer "+c.APIToken)
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ttp.DefaultClient.Do(req): %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll(res.Body): %w", err)
	}

	var match models.Match
	if err := json.Unmarshal(body, &match); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(body, &match): %w", err)
	}

	return &match, nil
}

func (c *PandaScoreClient) GetUpcomingMatches() ([]models.Match, error) {
	url := fmt.Sprintf("%s/csgo/matches/upcoming?filter[future]=true", c.BaseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest(\"GET\", url, nil): %w", err)
	}

	req.Header.Add("Authorization", "Bearer "+c.APIToken)
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ttp.DefaultClient.Do(req): %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll(res.Body): %w", err)
	}

	var matches []models.Match
	if err := json.Unmarshal(body, &matches); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(body, &matches): %w", err)
	}

	return matches, nil
}
