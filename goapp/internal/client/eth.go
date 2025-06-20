package client

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

type MatchRequestedEvent struct {
	RequestId *big.Int
	MatchId   *big.Int
	Requester common.Address
	Fee       *big.Int
}

type EventHandler func(event MatchRequestedEvent)

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

func (e *EthereumClient) GetPublicAddress() common.Address {
	return crypto.PubkeyToAddress(e.privateKey.PublicKey)
}

func (e *EthereumClient) ListenToMatchRequested(handler EventHandler) error {

	matchRequestedABI := `[{"anonymous":false,"inputs":[{"indexed":true,"name":"requestId","type":"uint256"},{"indexed":true,"name":"matchId","type":"uint256"},{"indexed":true,"name":"requester","type":"address"},{"indexed":false,"name":"fee","type":"uint256"}],"name":"MatchRequested","type":"event"}]`

	contractABI, err := abi.JSON(strings.NewReader(matchRequestedABI))
	if err != nil {
		return fmt.Errorf("failed to parse ABI: %w", err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{e.contractAddress},
		Topics:    [][]common.Hash{{contractABI.Events["MatchRequested"].ID}},
	}

	logs := make(chan types.Log)
	ctx := context.Background()
	sub, err := e.client.SubscribeFilterLogs(ctx, query, logs)
	
	if err != nil {
		log.Printf("WebSocket subscription failed (%v), falling back to polling", err)
		go e.pollForEvents(query, contractABI, handler)
		return nil
	}

	go func() {
		defer sub.Unsubscribe()

		for {
			select {
			case err := <-sub.Err():
				log.Printf("Subscription error: %v", err)
				return
			case vLog := <-logs:
				event, err := e.decodeMatchRequestedEvent(vLog, contractABI)
				if err != nil {
					log.Printf("Error decoding event: %v", err)
					continue
				}

				handler(event)
			}
		}
	}()

	return nil
}

func (e *EthereumClient) pollForEvents(query ethereum.FilterQuery, contractABI abi.ABI, handler EventHandler) {
	lastBlock := uint64(0)
	
	currentBlock, err := e.client.BlockNumber(context.Background())
	if err != nil {
		log.Printf("Failed to get current block number: %v", err)
		return
	}
	lastBlock = currentBlock

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		currentBlock, err := e.client.BlockNumber(context.Background())
		if err != nil {
			log.Printf("Failed to get current block number: %v", err)
			continue
		}

		if currentBlock > lastBlock {
			query.FromBlock = big.NewInt(int64(lastBlock + 1))
			query.ToBlock = big.NewInt(int64(currentBlock))

			logs, err := e.client.FilterLogs(context.Background(), query)
			if err != nil {
				log.Printf("Failed to filter logs: %v", err)
				continue
			}

			for _, vLog := range logs {
				event, err := e.decodeMatchRequestedEvent(vLog, contractABI)
				if err != nil {
					log.Printf("Error decoding event: %v", err)
					continue
				}

				handler(event)
			}

			lastBlock = currentBlock
		}
	}
}

func (e *EthereumClient) decodeMatchRequestedEvent(vLog types.Log, contractABI abi.ABI) (MatchRequestedEvent, error) {
	var event MatchRequestedEvent

	err := contractABI.UnpackIntoInterface(&struct {
		Fee *big.Int
	}{Fee: event.Fee}, "MatchRequested", vLog.Data)
	if err != nil {
		return event, fmt.Errorf("failed to unpack event data: %w", err)
	}

	if len(vLog.Topics) >= 4 {
		event.RequestId = new(big.Int).SetBytes(vLog.Topics[1].Bytes())
		event.MatchId = new(big.Int).SetBytes(vLog.Topics[2].Bytes())
		event.Requester = common.BytesToAddress(vLog.Topics[3].Bytes())
	}

	if len(vLog.Data) >= 32 {
		event.Fee = new(big.Int).SetBytes(vLog.Data[:32])
	}

	return event, nil
}
