package service

import (
    "encoding/json"
    "fmt"
    "os"

    "src/internal/client"
)

type MatchService struct {
    client *client.PandaScoreClient
}

func NewMatchService(client *client.PandaScoreClient) *MatchService {
    return &MatchService{
        client: client,
    }
}

func (s *MatchService) FetchAndSaveMatches() error {
    matches, err := s.client.GetUpcomingMatches()
    if err != nil {
        return fmt.Errorf("error fetching matches: %w", err)
    }

    for _, match := range matches {
        prettyJSON, err := json.MarshalIndent(match, "", "  ")
        if err != nil {
            return fmt.Errorf("error indenting JSON: %w", err)
        }
        fmt.Println(string(prettyJSON))
    }

    outputFile := "high_level_upcoming_matches.json"
    prettyJSON, err := json.MarshalIndent(matches, "", "  ")
    if err != nil {
        return fmt.Errorf("error indenting JSON for file: %w", err)
    }
    
    if err := os.WriteFile(outputFile, prettyJSON, 0644); err != nil {
        return fmt.Errorf("error writing to file: %w", err)
    }
    
    fmt.Printf("Data has been saved to %s\n", outputFile)
    return nil
}