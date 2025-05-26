package client

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"src/internal/contract"
)

type EthereumClient struct {
	client          *ethclient.Client
	contractAddress common.Address
	privateKey      *ecdsa.PrivateKey
	chainID         *big.Int
	contract        *contract.Contract
}

func NewEthereumClient(rpcURL, contractAddr, privateKeyHex, chainIDString string) (*EthereumClient, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("ethclient.Dial(rpcURL): %w", err)
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("crypto.HexToECDSA(privateKeyHex): %w", err)
	}

	i, err := strconv.Atoi(chainIDString)
	if err != nil {
		return nil, fmt.Errorf("strconv.Atoi(chainID): %w", err)
	}
	chainID := big.NewInt(int64(i))

	contractAddress := common.HexToAddress(contractAddr)
	contractInstance, err := contract.NewContract(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("contract.NewContract(): %w", err)
	}

	return &EthereumClient{
		client:          client,
		contractAddress: contractAddress,
		privateKey:      privateKey,
		chainID:         chainID,
		contract:        contractInstance,
	}, nil
}

func (e *EthereumClient) SendMatchesToContract(matches []contract.EsportOracleMatch) error {
	auth, err := bind.NewKeyedTransactorWithChainID(e.privateKey, e.chainID)
	if err != nil {
		return fmt.Errorf("bind.NewKeyedTransactorWithChainID(): %w", err)
	}

	log.Printf("Sending %d matches to contract at %s", len(matches), e.contractAddress.Hex())

	tx, err := e.contract.HandleNewMatches(auth, matches)
	if err != nil {
		return fmt.Errorf("contract.HandleNewMatches(): %w", err)
	}

	log.Printf("Transaction sent successfully: %s", tx.Hash().Hex())
	return nil
}

func (e *EthereumClient) GetPendingMatches() ([][32]byte, error) {
	return e.contract.GetPendingMatches(&bind.CallOpts{})
}

func (e *EthereumClient) GetListedNodes() ([]common.Address, error) {
	return e.contract.GetListedNodes(&bind.CallOpts{})
}

func (e *EthereumClient) GetMatchById(matchId *big.Int) (contract.EsportOracleMatch, error) {
	return e.contract.GetMatchById(&bind.CallOpts{}, matchId)
}

func (e *EthereumClient) AddFundsToStaking(amount *big.Int) error {
	auth, err := bind.NewKeyedTransactorWithChainID(e.privateKey, e.chainID)
	if err != nil {
		return fmt.Errorf("bind.NewKeyedTransactorWithChainID(): %w", err)
	}

	auth.Value = amount

	tx, err := e.contract.AddFundToStaking(auth)
	if err != nil {
		return fmt.Errorf("e.contract.AddFundToStaking(): %w", err)
	}

	log.Printf("Staking transaction sent: %s", tx.Hash().Hex())
	return nil
}

func (e *EthereumClient) WithdrawStake() error {
	auth, err := bind.NewKeyedTransactorWithChainID(e.privateKey, e.chainID)
	if err != nil {
		return fmt.Errorf("bind.NewKeyedTransactorWithChainID(): %w", err)
	}

	tx, err := e.contract.WithdrawStake(auth)
	if err != nil {
		return fmt.Errorf("contract.WithdrawStake(): %w", err)
	}

	log.Printf("Withdraw transaction sent: %s", tx.Hash().Hex())
	return nil
}

func (e *EthereumClient) GetContractOwner() (common.Address, error) {
	return e.contract.Owner(&bind.CallOpts{})
}

func (e *EthereumClient) GetFundsStaked(address common.Address) (*big.Int, error) {
	return e.contract.FundsStaked(&bind.CallOpts{}, address)
}

func (e *EthereumClient) GetNodeViolations(address common.Address) (struct {
	IncorrectMatches *big.Int
	IsBanned         bool
}, error) {
	return e.contract.NodeViolations(&bind.CallOpts{}, address)
}

func (e *EthereumClient) Close() {
	if e.client != nil {
		e.client.Close()
	}
}
