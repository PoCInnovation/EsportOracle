package main

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "src/internal/client"
    "src/internal/service"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("Warning: .env file not found")
    }

    apiToken := os.Getenv("TOKEN")
    if apiToken == "" {
        log.Println("Warning: API token not set in environment variables")
		os.Exit(1)
    }

    pandaClient := client.NewPandaScoreClient(apiToken)
    matchService := service.NewMatchService(pandaClient)

    if err := matchService.FetchAndSaveMatches(); err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
}