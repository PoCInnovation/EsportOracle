package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"src/internal/client"
	"src/internal/service"
)

func main() {
	fmt.Println("Esport Oracle starting...")

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	pandaToken := getEnvOrExit("PANDASCORE_API_TOKEN", "API token for PandaScore is required")
	rpcURL := getEnvOrExit("CLIENT_ETH", "Ethereum RPC URL is required")
	contractAddr := getEnvOrExit("CONTRACT_ADDRESS", "Contract address is required")
	privateKey := getEnvOrExit("PRIVATE_KEY", "Private key is required")
	chainID := getEnvOrExit("CHAIN_ID", "Chain ID is required")

	cronSchedule := os.Getenv("CRON_SCHEDULE")
	if cronSchedule == "" {
		cronSchedule = "*/15 * * * *"
	}

	pandaClient := client.NewPandaScoreClient(pandaToken)

	ethereumClient, err := client.NewEthereumClient(rpcURL, contractAddr, privateKey, chainID)
	if err != nil {
		log.Fatalf("client.NewEthereumClient(rpcURL, contractAddr, privateKey): %v", err)
	}

	matchService := service.NewMatchService(pandaClient, ethereumClient)

	fmt.Println("Running initial update")
	if err := matchService.RunOnce(); err != nil {
		log.Printf("Warning: matchService.RunOnce(): %v", err)
	}

	fmt.Println("Starting scheduler...")
	if err := matchService.StartScheduler(cronSchedule); err != nil {
		log.Fatalf("matchService.StartScheduler(cronSchedule): %v", err)
	}

	fmt.Println("Service running !")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	fmt.Println("Shutdown signal received stopping...")
	matchService.StopScheduler()

	fmt.Println("Waiting...")
	time.Sleep(2 * time.Second)

	fmt.Println("Esport Oracle stopped")
}

func getEnvOrExit(key, errorMsg string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("os.Getenv(%s): %s", key, errorMsg)
	}
	return value
}
