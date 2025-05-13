package main

import (
	"log"
	"strings"
	"math/big"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

var Infura = "https://hoodi.infura.io/v3/ed18ee4128b54ddaa05e76fd10786bd6"

/*func main() {
	client, err := ethclient.Dial(Infura)
	if err != nil {
		log.Fatalf("error connect account %d", err)
	}
	context, err := abi.JSON(strings.NewReader()) //avoir l'addresse du Contract
	if err != nil {
		log.Fatalf("Failed to parse ABI")
	}
	result, err := parsedABI.Pack("getBalance", nil)
	if err != nil {
		log.Fatalf("Failed to pack")
	}
	fmt.Println(result)
	//utiliser JSON pour récuperer les données
	//après use to 

}*/

func main() {
	client, err := ethclient.Dial(Infura)
	if err != nil {
		log.Fatal(err)
	}
	//charger l'abi du contract et l'addresse.
	contractAddress := "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"getPerson\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setPerson\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	parsed, err := abi.JSON(strings.NewReader(contractAddress))
	if err != nil {
		log.Fatal(err)
	}
	//créer une instance de ce contract en Go de l'ABI
	contract := bind.NewBoundContract(contractAddress, parsed, client, )
	privateKey , err := crypto.HexToECDSA("cle_prive")
	//interagir avec le contract, donc besoin d'une cle prive et ensuite d'appeller la fonction.
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}
	tx, err := contract.setPerson(auth, big.NewInt(12345), "Value")
	if err != nil {
		log.Fatal(err)
	}
}
