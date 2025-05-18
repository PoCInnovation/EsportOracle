package internal

import (
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

var Infura = "Url_Client_ETH"

func store_Match(match_info models.Match) {
	client, err := ethclient.Dial(Infura)
	if err != nil {
		log.Fatal("Erreur par rapport la clé privé: ", err)
	}
	private_Key, err := crypto.HexToECDSA("clé_privé_ici")
	if err != nil {
		log.Fatal("Erreur dans le chargement de la clé privée: ", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(private_Key, big.NewInt(1))
	if err != nil {
		log.Fatal("Erreur authentification: ", err)
	}
	contractAddress := common.HexToAddress("0xAdressContract")
	instance, err := contract.NewContract(contractAddress, client)
	if err != nil {
		log.Fatal("Erreur du contract :", err)
	}
	resuls := convert_Results(match_info.Results)
	games := convert_Games(match_info.Games)
	opponents := convert_Opponents(match_info.Opponents)
	beginAt := big.NewInt(0)
	if res_time, err := time.Parse(time.RFC3339, match_info.BeginAt); err == nil {
		beginAt = big.NewInt((res_time.Unix()))
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