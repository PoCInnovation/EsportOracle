package service

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"

	"src/internal/client"
	"src/internal/models"
)

type MatchService struct {
	pandaClient    *client.PandaScoreClient
	ethereumClient *client.EthereumClient
	scheduler      *cron.Cron
}

func NewMatchService(pandaClient *client.PandaScoreClient, ethereumClient *client.EthereumClient) *MatchService {
	return &MatchService{
		pandaClient:    pandaClient,
		ethereumClient: ethereumClient,
		scheduler:      cron.New(),
	}
}

func (s *MatchService) FetchAndSaveMatches() error {
	matches, err := s.pandaClient.GetUpcomingMatches()
	if err != nil {
		return fmt.Errorf(".pandaClient.GetUpcomingMatches(): %w", err)
	}

	log.Printf("Fetched %d matches\n", len(matches))

	outputFile := "upcoming_matches.json"
	prettyJSON, err := json.MarshalIndent(matches, "", "  ")
	if err != nil {
		return fmt.Errorf("json.MarshalIndent(matches, \"\", \"  \"): %w", err)
	}

	if err := os.WriteFile(outputFile, prettyJSON, 0644); err != nil {
		return fmt.Errorf("os.WriteFile(outputFile, prettyJSON, 0644): %w", err)
	}

	log.Printf("Raw data has been saved to %s\n", outputFile)
	return nil
}

func (s *MatchService) FetchAndUpdateOracle() error {
	matches, err := s.pandaClient.GetUpcomingMatches()
	if err != nil {
		return fmt.Errorf("s.pandaClient.GetUpcomingMatches(): %w", err)
	}

	log.Printf("Fetched %d matches from PandaScore API\n", len(matches))

	var contractMatches []models.ContractMatch
	for _, match := range matches {
		contractMatch := match.ToContractMatch()
		contractMatches = append(contractMatches, contractMatch)

		log.Printf("Match #%d: %d opponents, %d games, %d results, begins at %s\n",
			match.ID,
			len(match.Opponents),
			len(match.Games),
			len(match.Results),
			match.BeginAt)
	}

	contractFile := "contract_matches.json"
	contractJSON, err := json.MarshalIndent(contractMatches, "", "  ")
	if err != nil {
		log.Printf("Warning: %v", err)
	} else {
		if err := os.WriteFile(contractFile, contractJSON, 0644); err != nil {
			log.Printf("Warning: %v", err)
		} else {
			log.Printf("Contract data saved to %s\n", contractFile)
		}
	}

	if len(contractMatches) > 0 {
		if err := s.ethereumClient.SendMatchesToContract(contractMatches); err != nil {
			return fmt.Errorf("s.ethereumClient.SendMatchesToContract(contractMatches): %w", err)
		}
		log.Printf("Successfully sent %d matches to the contract\n", len(contractMatches))
	} else {
		log.Println("No matches to send to the contract")
	}

	return nil
}

func (s *MatchService) StartScheduler(cronExpression string) error {
	if cronExpression == "" {
		cronExpression = "*/15 * * * *"
	}

	_, err := s.scheduler.AddFunc(cronExpression, func() {
		log.Println("Running scheduled task: FetchAndUpdateOracle")
		if err := s.FetchAndUpdateOracle(); err != nil {
			log.Printf("s.FetchAndUpdateOracle(): %v\n", err)
		}
	})

	if err != nil {
		return fmt.Errorf("s.scheduler.AddFunc(): %w", err)
	}

	s.scheduler.Start()
	log.Printf("Scheduler started: %s\n", cronExpression)
	return nil
}

func (s *MatchService) StopScheduler() {
	if s.scheduler != nil {
		s.scheduler.Stop()
		log.Println("Scheduler stopped")
	}
}

func (s *MatchService) RunOnce() error {
	log.Println("Running one time update")
	return s.FetchAndUpdateOracle()
}
