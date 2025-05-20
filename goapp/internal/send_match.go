package internal

import (
	"os"
	"github.com/joho/godotenv"
	"fmt"
	"log"
	"math/big"
	"time"
	"src/internal/models"
	"src/internal/build"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Config_t struct {
	Url_Client	string
	Private_Key string
	Contract_Address string
	Chain_Id *big.Int
}

func load_Env() Config_t {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erreur chargement .env:", err)
	}
	chain_ID := os.Getenv("CHAIN_ID")
	private_Key := os.Getenv("PRIVATE_KEY")
	contractAddress := os.Getenv("CONTRACT_ADDRESS")
	client_ETH := os.Getenv("CLIENT_ETH")
	new_chain_ID, err := new(big.Int).SetString(chain_ID, 10)
	if !err {
		log.Fatal("Chain_ID invalid")
	}
	return Config_t{
		Url_Client: client_ETH,
		Private_Key: private_Key,
		Contract_Address: contractAddress,
		Chain_Id: new_chain_ID,
	}
}

func store_Match(match_info models.Match) {
	config := load_Env()
	client, err := ethclient.Dial(config.Url_Client)
	if err != nil {
		log.Fatal("Erreur par rapport la clé privé: ", err)
	}
	private_Key, err := crypto.HexToECDSA(config.Private_Key)
	if err != nil {
		log.Fatal("Erreur dans le chargement de la clé privée: ", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(private_Key, config.Chain_Id)
	if err != nil {
		log.Fatal("Erreur authentification: ", err)
	}
	contractAddress := common.HexToAddress(config.Contract_Address)
	instance, err := contract.NewContract(contractAddress, client)
	if err != nil {
		log.Fatal("Erreur du contract :", err)
	}
	resuls := convert_Results(match_info.Results)
	games := convert_Games(match_info.Games)
	opponents := convert_Opponents(match_info.Opponents)
	beginAt := big.NewInt(0)
	if t, err := time.Parse(time.RFC3339, match_info.BeginAt); err == nil {
		beginAt = big.NewInt((t.Unix()))
	}
	match := contract.EsportOracleMatch{
		Id: big.NewInt(int64(match_info.ID)),
		Opponents: opponents,
		Game: games,
		Result: resuls,
		WinnerId: big.NewInt(int64(match_info.WinnerID)),
		BeginAt: beginAt,
	}
	res, err := instance.HandleNewMatches(auth, []contract.EsportOracleMatch{match})
	if err != nil {
		log.Fatal("Erreur dans l'envoi: ", err)
	}
	fmt.Println("Transaction validé :", res)
}

func convert_Games(models_games []models.Game) []contract.EsportOracleGames {
	var games []contract.EsportOracleGames
	for _, gam := range models_games {
		games = append(games, contract.EsportOracleGames{
			Id: big.NewInt(int64(gam.ID)),
			Finished: gam.Finished,
			WinnerId: big.NewInt(int64(*gam.Winner.ID)),
		})
	}
	return games
}

func convert_Opponents(models_opponents []models.Opponent) []contract.EsportOracleOpponents {
	var opponents []contract.EsportOracleOpponents
	for _, opp := range models_opponents {
		opponents = append(opponents, contract.EsportOracleOpponents{
			Acronym: opp.Opponent.Acronym,
			Id: big.NewInt(int64(opp.Opponent.ID)),
			Name: opp.Opponent.Name,
		})
	}
	return opponents
}

func convert_Results(models_results []models.Result) []contract.EsportOracleResult {
	var results []contract.EsportOracleResult
	for _, res := range models_results {
		results = append(results, contract.EsportOracleResult{
			Score: uint8(res.Score),
			TeamId: big.NewInt(int64(res.TeamID)),
		})
	}
	return results
}