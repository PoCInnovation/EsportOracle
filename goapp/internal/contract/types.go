package contract

import (
    "math/big"
	"github.com/ethereum/go-ethereum/common"
)

// EsportOracleTypesGames is an auto generated low-level Go binding around an user-defined struct.
type EsportOracleTypesGames struct {
	Id       *big.Int
	Finished bool
	WinnerId *big.Int
}

// EsportOracleTypesMatch is an auto generated low-level Go binding around an user-defined struct.
type EsportOracleTypesMatch struct {
	Id        *big.Int
	Opponents []EsportOracleTypesOpponents
	Game      []EsportOracleTypesGames
	Result    []EsportOracleTypesResult
	WinnerId  *big.Int
	BeginAt   *big.Int
}

// EsportOracleTypesOpponents is an auto generated low-level Go binding around an user-defined struct.
type EsportOracleTypesOpponents struct {
	Acronym string
	Id      *big.Int
	Name    string
}

// EsportOracleTypesResult is an auto generated low-level Go binding around an user-defined struct.
type EsportOracleTypesResult struct {
	Score  uint8
	TeamId *big.Int
}
// EsportOracleTypesMatchRequest is an auto generated low-level Go binding around an user-defined struct.
type EsportOracleTypesMatchRequest struct {
	MatchId   *big.Int
	Requester common.Address
	Fee       *big.Int
	Fulfilled bool
}

