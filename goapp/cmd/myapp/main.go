package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	"src/internal/client"
	"src/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Esport Oracle starting...")

	envFile := ".env"
    if len(os.Args) > 1 {
        envFile = ".env." + os.Args[1]
    }
    
    err := godotenv.Load(envFile)
    if err != nil {
        log.Fatal("Error loading file .env:", err)
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

	// Check if node is already registered and has funds staked
	fmt.Println("Checking node registration status...")
	publicKey := ethereumClient.GetPublicAddress()
	fundsStaked, err := ethereumClient.GetFundsStaked(publicKey)
	if err != nil {
		log.Printf("Warning: Could not check staked funds: %v", err)
	} else if fundsStaked.Cmp(big.NewInt(0)) == 0 {
		fmt.Println("Node not registered. Registering node by staking funds...")
		stakeAmount := big.NewInt(1000000000000000) // 0.001 ETH in wei
		if err := ethereumClient.AddFundsToStaking(stakeAmount); err != nil {
			log.Printf("Warning: Failed to stake funds and register node: %v", err)
			log.Println("Node may not be registered for oracle participation")
		} else {
			fmt.Println("Successfully staked funds and registered as oracle node")
		}
	} else {
		fmt.Printf("Node already registered with %s wei staked\n", fundsStaked.String())
	}

	err = ethereumClient.ListenToMatchRequested(func(event client.MatchRequestedEvent) {
		fmt.Printf("   Event MatchRequested called\n")
		fmt.Printf("   Request ID: %s\n", event.RequestId.String())
		fmt.Printf("   Match ID: %s\n", event.MatchId.String())
		fmt.Printf("   Requester: %s\n", event.Requester.Hex())
		fmt.Printf("   Fee: %s\n", event.Fee.String())

		fmt.Println("   Event MatchRequested processed successfully")
		
		if err := matchService.GetMatchByIDAndSendToContract(int(event.MatchId.Int64())); err != nil {
			fmt.Printf("   Error processing match by ID: %v\n", err)
		} else {
			fmt.Printf("   Match successfully fetched and sent to contract\n")
			fmt.Println("	Requester's address (should be contract):", event.Requester.Hex())
			match, err := ethereumClient.GetMatchById(event.MatchId)
			if err != nil {
				log.Printf("   Error getting match by ID %s to call CallMatchReceived: %v\n", event.MatchId.String(), err)
			} else {
				if err := ethereumClient.CallOnMatchReceived(contractAddr, event.Requester, event.MatchId, match); err != nil {
					log.Printf("   Error calling CallMatchReceived on client %s: %v\n", event.Requester.Hex(), err)
				} else {
					fmt.Printf("   CallMatchReceived successfully called on client %s\n", event.Requester.Hex())
				}
			}
		}
		
	})

	if err != nil {
		log.Fatalf("Failed to start event listener: %v", err)
	}

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
