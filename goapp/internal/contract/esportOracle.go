// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)


// EsportOracleMetaData contains all meta data concerning the EsportOracle contract.
var EsportOracleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MAX_VIOLATIONS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PUNISHMENT_AMOUNT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_addressByHash\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_fundsStaked\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_matchMapping\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_winnerId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_beginAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_matchVotes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_nodeViolations\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"incorrectMatches\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isBanned\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_pendingMatchesHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addFundToStaking\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"banNode\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkQorum\",\"inputs\":[{\"name\":\"matchData\",\"type\":\"tuple\",\"internalType\":\"structEsportOracleTypes.Match\",\"components\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_opponents\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Opponents[]\",\"components\":[{\"name\":\"_acronym\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"_game\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Games[]\",\"components\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_finished\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_winnerId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_result\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Result[]\",\"components\":[{\"name\":\"_score\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_teamId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_winnerId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_beginAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getListedNodes\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMatchById\",\"inputs\":[{\"name\":\"matchId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structEsportOracleTypes.Match\",\"components\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_opponents\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Opponents[]\",\"components\":[{\"name\":\"_acronym\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"_game\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Games[]\",\"components\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_finished\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_winnerId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_result\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Result[]\",\"components\":[{\"name\":\"_score\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_teamId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_winnerId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_beginAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingMatches\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"handleNewMatches\",\"inputs\":[{\"name\":\"newMatch\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Match[]\",\"components\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_opponents\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Opponents[]\",\"components\":[{\"name\":\"_acronym\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"_game\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Games[]\",\"components\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_finished\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_winnerId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_result\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Result[]\",\"components\":[{\"name\":\"_score\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_teamId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_winnerId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_beginAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rehabilitateNode\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOwner\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawStake\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"NodeBanned\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodePunished\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"violationsCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"newNodeAdded\",\"inputs\":[{\"name\":\"addressAdded\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"stakingSuccess\",\"inputs\":[{\"name\":\"addressAdded\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]}]",
}

// EsportOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use EsportOracleMetaData.ABI instead.
var EsportOracleABI = EsportOracleMetaData.ABI

// EsportOracle is an auto generated Go binding around an Ethereum contract.
type EsportOracle struct {
	EsportOracleCaller     // Read-only binding to the contract
	EsportOracleTransactor // Write-only binding to the contract
	EsportOracleFilterer   // Log filterer for contract events
}

// EsportOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type EsportOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EsportOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EsportOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EsportOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EsportOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EsportOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EsportOracleSession struct {
	Contract     *EsportOracle     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EsportOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EsportOracleCallerSession struct {
	Contract *EsportOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// EsportOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EsportOracleTransactorSession struct {
	Contract     *EsportOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EsportOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type EsportOracleRaw struct {
	Contract *EsportOracle // Generic contract binding to access the raw methods on
}

// EsportOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EsportOracleCallerRaw struct {
	Contract *EsportOracleCaller // Generic read-only contract binding to access the raw methods on
}

// EsportOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EsportOracleTransactorRaw struct {
	Contract *EsportOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEsportOracle creates a new instance of EsportOracle, bound to a specific deployed contract.
func NewEsportOracle(address common.Address, backend bind.ContractBackend) (*EsportOracle, error) {
	contract, err := bindEsportOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EsportOracle{EsportOracleCaller: EsportOracleCaller{contract: contract}, EsportOracleTransactor: EsportOracleTransactor{contract: contract}, EsportOracleFilterer: EsportOracleFilterer{contract: contract}}, nil
}

// NewEsportOracleCaller creates a new read-only instance of EsportOracle, bound to a specific deployed contract.
func NewEsportOracleCaller(address common.Address, caller bind.ContractCaller) (*EsportOracleCaller, error) {
	contract, err := bindEsportOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EsportOracleCaller{contract: contract}, nil
}

// NewEsportOracleTransactor creates a new write-only instance of EsportOracle, bound to a specific deployed contract.
func NewEsportOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*EsportOracleTransactor, error) {
	contract, err := bindEsportOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EsportOracleTransactor{contract: contract}, nil
}

// NewEsportOracleFilterer creates a new log filterer instance of EsportOracle, bound to a specific deployed contract.
func NewEsportOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*EsportOracleFilterer, error) {
	contract, err := bindEsportOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EsportOracleFilterer{contract: contract}, nil
}

// bindEsportOracle binds a generic wrapper to an already deployed contract.
func bindEsportOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EsportOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EsportOracle *EsportOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EsportOracle.Contract.EsportOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EsportOracle *EsportOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EsportOracle.Contract.EsportOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EsportOracle *EsportOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EsportOracle.Contract.EsportOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EsportOracle *EsportOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EsportOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EsportOracle *EsportOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EsportOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EsportOracle *EsportOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EsportOracle.Contract.contract.Transact(opts, method, params...)
}

// MAXVIOLATIONS is a free data retrieval call binding the contract method 0xe8819ff7.
//
// Solidity: function MAX_VIOLATIONS() view returns(uint256)
func (_EsportOracle *EsportOracleCaller) MAXVIOLATIONS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EsportOracle.contract.Call(opts, &out, "MAX_VIOLATIONS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXVIOLATIONS is a free data retrieval call binding the contract method 0xe8819ff7.
//
// Solidity: function MAX_VIOLATIONS() view returns(uint256)
func (_EsportOracle *EsportOracleSession) MAXVIOLATIONS() (*big.Int, error) {
	return _EsportOracle.Contract.MAXVIOLATIONS(&_EsportOracle.CallOpts)
}

// MAXVIOLATIONS is a free data retrieval call binding the contract method 0xe8819ff7.
//
// Solidity: function MAX_VIOLATIONS() view returns(uint256)
func (_EsportOracle *EsportOracleCallerSession) MAXVIOLATIONS() (*big.Int, error) {
	return _EsportOracle.Contract.MAXVIOLATIONS(&_EsportOracle.CallOpts)
}

// PUNISHMENTAMOUNT is a free data retrieval call binding the contract method 0xf21b121c.
//
// Solidity: function PUNISHMENT_AMOUNT() view returns(uint256)
func (_EsportOracle *EsportOracleCaller) PUNISHMENTAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EsportOracle.contract.Call(opts, &out, "PUNISHMENT_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUNISHMENTAMOUNT is a free data retrieval call binding the contract method 0xf21b121c.
//
// Solidity: function PUNISHMENT_AMOUNT() view returns(uint256)
func (_EsportOracle *EsportOracleSession) PUNISHMENTAMOUNT() (*big.Int, error) {
	return _EsportOracle.Contract.PUNISHMENTAMOUNT(&_EsportOracle.CallOpts)
}

// PUNISHMENTAMOUNT is a free data retrieval call binding the contract method 0xf21b121c.
//
// Solidity: function PUNISHMENT_AMOUNT() view returns(uint256)
func (_EsportOracle *EsportOracleCallerSession) PUNISHMENTAMOUNT() (*big.Int, error) {
	return _EsportOracle.Contract.PUNISHMENTAMOUNT(&_EsportOracle.CallOpts)
}

// AddressByHash is a free data retrieval call binding the contract method 0xbb27ea26.
//
// Solidity: function _addressByHash(bytes32 , uint256 ) view returns(address)
func (_EsportOracle *EsportOracleCaller) AddressByHash(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EsportOracle.contract.Call(opts, &out, "_addressByHash", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressByHash is a free data retrieval call binding the contract method 0xbb27ea26.
//
// Solidity: function _addressByHash(bytes32 , uint256 ) view returns(address)
func (_EsportOracle *EsportOracleSession) AddressByHash(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _EsportOracle.Contract.AddressByHash(&_EsportOracle.CallOpts, arg0, arg1)
}

// AddressByHash is a free data retrieval call binding the contract method 0xbb27ea26.
//
// Solidity: function _addressByHash(bytes32 , uint256 ) view returns(address)
func (_EsportOracle *EsportOracleCallerSession) AddressByHash(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _EsportOracle.Contract.AddressByHash(&_EsportOracle.CallOpts, arg0, arg1)
}

// FundsStaked is a free data retrieval call binding the contract method 0x652832f1.
//
// Solidity: function _fundsStaked(address ) view returns(uint256)
func (_EsportOracle *EsportOracleCaller) FundsStaked(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EsportOracle.contract.Call(opts, &out, "_fundsStaked", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundsStaked is a free data retrieval call binding the contract method 0x652832f1.
//
// Solidity: function _fundsStaked(address ) view returns(uint256)
func (_EsportOracle *EsportOracleSession) FundsStaked(arg0 common.Address) (*big.Int, error) {
	return _EsportOracle.Contract.FundsStaked(&_EsportOracle.CallOpts, arg0)
}

// FundsStaked is a free data retrieval call binding the contract method 0x652832f1.
//
// Solidity: function _fundsStaked(address ) view returns(uint256)
func (_EsportOracle *EsportOracleCallerSession) FundsStaked(arg0 common.Address) (*big.Int, error) {
	return _EsportOracle.Contract.FundsStaked(&_EsportOracle.CallOpts, arg0)
}

// MatchMapping is a free data retrieval call binding the contract method 0x42d150c9.
//
// Solidity: function _matchMapping(uint256 ) view returns(uint256 _id, uint256 _winnerId, uint256 _beginAt)
func (_EsportOracle *EsportOracleCaller) MatchMapping(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id       *big.Int
	WinnerId *big.Int
	BeginAt  *big.Int
}, error) {
	var out []interface{}
	err := _EsportOracle.contract.Call(opts, &out, "_matchMapping", arg0)

	outstruct := new(struct {
		Id       *big.Int
		WinnerId *big.Int
		BeginAt  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.WinnerId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BeginAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MatchMapping is a free data retrieval call binding the contract method 0x42d150c9.
//
// Solidity: function _matchMapping(uint256 ) view returns(uint256 _id, uint256 _winnerId, uint256 _beginAt)
func (_EsportOracle *EsportOracleSession) MatchMapping(arg0 *big.Int) (struct {
	Id       *big.Int
	WinnerId *big.Int
	BeginAt  *big.Int
}, error) {
	return _EsportOracle.Contract.MatchMapping(&_EsportOracle.CallOpts, arg0)
}

// MatchMapping is a free data retrieval call binding the contract method 0x42d150c9.
//
// Solidity: function _matchMapping(uint256 ) view returns(uint256 _id, uint256 _winnerId, uint256 _beginAt)
func (_EsportOracle *EsportOracleCallerSession) MatchMapping(arg0 *big.Int) (struct {
	Id       *big.Int
	WinnerId *big.Int
	BeginAt  *big.Int
}, error) {
	return _EsportOracle.Contract.MatchMapping(&_EsportOracle.CallOpts, arg0)
}

// MatchVotes is a free data retrieval call binding the contract method 0x1902762f.
//
// Solidity: function _matchVotes(bytes32 ) view returns(uint8)
func (_EsportOracle *EsportOracleCaller) MatchVotes(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _EsportOracle.contract.Call(opts, &out, "_matchVotes", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MatchVotes is a free data retrieval call binding the contract method 0x1902762f.
//
// Solidity: function _matchVotes(bytes32 ) view returns(uint8)
func (_EsportOracle *EsportOracleSession) MatchVotes(arg0 [32]byte) (uint8, error) {
	return _EsportOracle.Contract.MatchVotes(&_EsportOracle.CallOpts, arg0)
}

// MatchVotes is a free data retrieval call binding the contract method 0x1902762f.
//
// Solidity: function _matchVotes(bytes32 ) view returns(uint8)
func (_EsportOracle *EsportOracleCallerSession) MatchVotes(arg0 [32]byte) (uint8, error) {
	return _EsportOracle.Contract.MatchVotes(&_EsportOracle.CallOpts, arg0)
}

// NodeViolations is a free data retrieval call binding the contract method 0x91cac59e.
//
// Solidity: function _nodeViolations(address ) view returns(uint256 incorrectMatches, bool isBanned)
func (_EsportOracle *EsportOracleCaller) NodeViolations(opts *bind.CallOpts, arg0 common.Address) (struct {
	IncorrectMatches *big.Int
	IsBanned         bool
}, error) {
	var out []interface{}
	err := _EsportOracle.contract.Call(opts, &out, "_nodeViolations", arg0)

	outstruct := new(struct {
		IncorrectMatches *big.Int
		IsBanned         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IncorrectMatches = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IsBanned = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// NodeViolations is a free data retrieval call binding the contract method 0x91cac59e.
//
// Solidity: function _nodeViolations(address ) view returns(uint256 incorrectMatches, bool isBanned)
func (_EsportOracle *EsportOracleSession) NodeViolations(arg0 common.Address) (struct {
	IncorrectMatches *big.Int
	IsBanned         bool
}, error) {
	return _EsportOracle.Contract.NodeViolations(&_EsportOracle.CallOpts, arg0)
}

// NodeViolations is a free data retrieval call binding the contract method 0x91cac59e.
//
// Solidity: function _nodeViolations(address ) view returns(uint256 incorrectMatches, bool isBanned)
func (_EsportOracle *EsportOracleCallerSession) NodeViolations(arg0 common.Address) (struct {
	IncorrectMatches *big.Int
	IsBanned         bool
}, error) {
	return _EsportOracle.Contract.NodeViolations(&_EsportOracle.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_EsportOracle *EsportOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EsportOracle.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_EsportOracle *EsportOracleSession) Owner() (common.Address, error) {
	return _EsportOracle.Contract.Owner(&_EsportOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_EsportOracle *EsportOracleCallerSession) Owner() (common.Address, error) {
	return _EsportOracle.Contract.Owner(&_EsportOracle.CallOpts)
}

// PendingMatchesHashes is a free data retrieval call binding the contract method 0xda83225a.
//
// Solidity: function _pendingMatchesHashes(uint256 ) view returns(bytes32)
func (_EsportOracle *EsportOracleCaller) PendingMatchesHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _EsportOracle.contract.Call(opts, &out, "_pendingMatchesHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PendingMatchesHashes is a free data retrieval call binding the contract method 0xda83225a.
//
// Solidity: function _pendingMatchesHashes(uint256 ) view returns(bytes32)
func (_EsportOracle *EsportOracleSession) PendingMatchesHashes(arg0 *big.Int) ([32]byte, error) {
	return _EsportOracle.Contract.PendingMatchesHashes(&_EsportOracle.CallOpts, arg0)
}

// PendingMatchesHashes is a free data retrieval call binding the contract method 0xda83225a.
//
// Solidity: function _pendingMatchesHashes(uint256 ) view returns(bytes32)
func (_EsportOracle *EsportOracleCallerSession) PendingMatchesHashes(arg0 *big.Int) ([32]byte, error) {
	return _EsportOracle.Contract.PendingMatchesHashes(&_EsportOracle.CallOpts, arg0)
}

// CheckQorum is a free data retrieval call binding the contract method 0x31ecedd1.
//
// Solidity: function checkQorum((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256) matchData) view returns(bool)
func (_EsportOracle *EsportOracleCaller) CheckQorum(opts *bind.CallOpts, matchData EsportOracleTypesMatch) (bool, error) {
	var out []interface{}
	err := _EsportOracle.contract.Call(opts, &out, "checkQorum", matchData)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckQorum is a free data retrieval call binding the contract method 0x31ecedd1.
//
// Solidity: function checkQorum((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256) matchData) view returns(bool)
func (_EsportOracle *EsportOracleSession) CheckQorum(matchData EsportOracleTypesMatch) (bool, error) {
	return _EsportOracle.Contract.CheckQorum(&_EsportOracle.CallOpts, matchData)
}

// CheckQorum is a free data retrieval call binding the contract method 0x31ecedd1.
//
// Solidity: function checkQorum((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256) matchData) view returns(bool)
func (_EsportOracle *EsportOracleCallerSession) CheckQorum(matchData EsportOracleTypesMatch) (bool, error) {
	return _EsportOracle.Contract.CheckQorum(&_EsportOracle.CallOpts, matchData)
}

// GetListedNodes is a free data retrieval call binding the contract method 0x50344f48.
//
// Solidity: function getListedNodes() view returns(address[])
func (_EsportOracle *EsportOracleCaller) GetListedNodes(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EsportOracle.contract.Call(opts, &out, "getListedNodes")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetListedNodes is a free data retrieval call binding the contract method 0x50344f48.
//
// Solidity: function getListedNodes() view returns(address[])
func (_EsportOracle *EsportOracleSession) GetListedNodes() ([]common.Address, error) {
	return _EsportOracle.Contract.GetListedNodes(&_EsportOracle.CallOpts)
}

// GetListedNodes is a free data retrieval call binding the contract method 0x50344f48.
//
// Solidity: function getListedNodes() view returns(address[])
func (_EsportOracle *EsportOracleCallerSession) GetListedNodes() ([]common.Address, error) {
	return _EsportOracle.Contract.GetListedNodes(&_EsportOracle.CallOpts)
}

// GetMatchById is a free data retrieval call binding the contract method 0xf3c9fd0e.
//
// Solidity: function getMatchById(uint256 matchId) view returns((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256))
func (_EsportOracle *EsportOracleCaller) GetMatchById(opts *bind.CallOpts, matchId *big.Int) (EsportOracleTypesMatch, error) {
	var out []interface{}
	err := _EsportOracle.contract.Call(opts, &out, "getMatchById", matchId)

	if err != nil {
		return *new(EsportOracleTypesMatch), err
	}

	out0 := *abi.ConvertType(out[0], new(EsportOracleTypesMatch)).(*EsportOracleTypesMatch)

	return out0, err

}

// GetMatchById is a free data retrieval call binding the contract method 0xf3c9fd0e.
//
// Solidity: function getMatchById(uint256 matchId) view returns((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256))
func (_EsportOracle *EsportOracleSession) GetMatchById(matchId *big.Int) (EsportOracleTypesMatch, error) {
	return _EsportOracle.Contract.GetMatchById(&_EsportOracle.CallOpts, matchId)
}

// GetMatchById is a free data retrieval call binding the contract method 0xf3c9fd0e.
//
// Solidity: function getMatchById(uint256 matchId) view returns((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256))
func (_EsportOracle *EsportOracleCallerSession) GetMatchById(matchId *big.Int) (EsportOracleTypesMatch, error) {
	return _EsportOracle.Contract.GetMatchById(&_EsportOracle.CallOpts, matchId)
}

// GetPendingMatches is a free data retrieval call binding the contract method 0x5f29d4b1.
//
// Solidity: function getPendingMatches() view returns(bytes32[])
func (_EsportOracle *EsportOracleCaller) GetPendingMatches(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _EsportOracle.contract.Call(opts, &out, "getPendingMatches")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetPendingMatches is a free data retrieval call binding the contract method 0x5f29d4b1.
//
// Solidity: function getPendingMatches() view returns(bytes32[])
func (_EsportOracle *EsportOracleSession) GetPendingMatches() ([][32]byte, error) {
	return _EsportOracle.Contract.GetPendingMatches(&_EsportOracle.CallOpts)
}

// GetPendingMatches is a free data retrieval call binding the contract method 0x5f29d4b1.
//
// Solidity: function getPendingMatches() view returns(bytes32[])
func (_EsportOracle *EsportOracleCallerSession) GetPendingMatches() ([][32]byte, error) {
	return _EsportOracle.Contract.GetPendingMatches(&_EsportOracle.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EsportOracle *EsportOracleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EsportOracle.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EsportOracle *EsportOracleSession) Paused() (bool, error) {
	return _EsportOracle.Contract.Paused(&_EsportOracle.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EsportOracle *EsportOracleCallerSession) Paused() (bool, error) {
	return _EsportOracle.Contract.Paused(&_EsportOracle.CallOpts)
}

// AddFundToStaking is a paid mutator transaction binding the contract method 0x5c0a126d.
//
// Solidity: function addFundToStaking() payable returns()
func (_EsportOracle *EsportOracleTransactor) AddFundToStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EsportOracle.contract.Transact(opts, "addFundToStaking")
}

// AddFundToStaking is a paid mutator transaction binding the contract method 0x5c0a126d.
//
// Solidity: function addFundToStaking() payable returns()
func (_EsportOracle *EsportOracleSession) AddFundToStaking() (*types.Transaction, error) {
	return _EsportOracle.Contract.AddFundToStaking(&_EsportOracle.TransactOpts)
}

// AddFundToStaking is a paid mutator transaction binding the contract method 0x5c0a126d.
//
// Solidity: function addFundToStaking() payable returns()
func (_EsportOracle *EsportOracleTransactorSession) AddFundToStaking() (*types.Transaction, error) {
	return _EsportOracle.Contract.AddFundToStaking(&_EsportOracle.TransactOpts)
}

// BanNode is a paid mutator transaction binding the contract method 0x3faba59e.
//
// Solidity: function banNode(address node) returns()
func (_EsportOracle *EsportOracleTransactor) BanNode(opts *bind.TransactOpts, node common.Address) (*types.Transaction, error) {
	return _EsportOracle.contract.Transact(opts, "banNode", node)
}

// BanNode is a paid mutator transaction binding the contract method 0x3faba59e.
//
// Solidity: function banNode(address node) returns()
func (_EsportOracle *EsportOracleSession) BanNode(node common.Address) (*types.Transaction, error) {
	return _EsportOracle.Contract.BanNode(&_EsportOracle.TransactOpts, node)
}

// BanNode is a paid mutator transaction binding the contract method 0x3faba59e.
//
// Solidity: function banNode(address node) returns()
func (_EsportOracle *EsportOracleTransactorSession) BanNode(node common.Address) (*types.Transaction, error) {
	return _EsportOracle.Contract.BanNode(&_EsportOracle.TransactOpts, node)
}

// HandleNewMatches is a paid mutator transaction binding the contract method 0x3b23e7d7.
//
// Solidity: function handleNewMatches((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256)[] newMatch) returns()
func (_EsportOracle *EsportOracleTransactor) HandleNewMatches(opts *bind.TransactOpts, newMatch []EsportOracleTypesMatch) (*types.Transaction, error) {
	return _EsportOracle.contract.Transact(opts, "handleNewMatches", newMatch)
}

// HandleNewMatches is a paid mutator transaction binding the contract method 0x3b23e7d7.
//
// Solidity: function handleNewMatches((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256)[] newMatch) returns()
func (_EsportOracle *EsportOracleSession) HandleNewMatches(newMatch []EsportOracleTypesMatch) (*types.Transaction, error) {
	return _EsportOracle.Contract.HandleNewMatches(&_EsportOracle.TransactOpts, newMatch)
}

// HandleNewMatches is a paid mutator transaction binding the contract method 0x3b23e7d7.
//
// Solidity: function handleNewMatches((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256)[] newMatch) returns()
func (_EsportOracle *EsportOracleTransactorSession) HandleNewMatches(newMatch []EsportOracleTypesMatch) (*types.Transaction, error) {
	return _EsportOracle.Contract.HandleNewMatches(&_EsportOracle.TransactOpts, newMatch)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_EsportOracle *EsportOracleTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EsportOracle.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_EsportOracle *EsportOracleSession) Pause() (*types.Transaction, error) {
	return _EsportOracle.Contract.Pause(&_EsportOracle.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_EsportOracle *EsportOracleTransactorSession) Pause() (*types.Transaction, error) {
	return _EsportOracle.Contract.Pause(&_EsportOracle.TransactOpts)
}

// RehabilitateNode is a paid mutator transaction binding the contract method 0x148b56d6.
//
// Solidity: function rehabilitateNode(address node) returns()
func (_EsportOracle *EsportOracleTransactor) RehabilitateNode(opts *bind.TransactOpts, node common.Address) (*types.Transaction, error) {
	return _EsportOracle.contract.Transact(opts, "rehabilitateNode", node)
}

// RehabilitateNode is a paid mutator transaction binding the contract method 0x148b56d6.
//
// Solidity: function rehabilitateNode(address node) returns()
func (_EsportOracle *EsportOracleSession) RehabilitateNode(node common.Address) (*types.Transaction, error) {
	return _EsportOracle.Contract.RehabilitateNode(&_EsportOracle.TransactOpts, node)
}

// RehabilitateNode is a paid mutator transaction binding the contract method 0x148b56d6.
//
// Solidity: function rehabilitateNode(address node) returns()
func (_EsportOracle *EsportOracleTransactorSession) RehabilitateNode(node common.Address) (*types.Transaction, error) {
	return _EsportOracle.Contract.RehabilitateNode(&_EsportOracle.TransactOpts, node)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_EsportOracle *EsportOracleTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EsportOracle.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_EsportOracle *EsportOracleSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _EsportOracle.Contract.SetOwner(&_EsportOracle.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_EsportOracle *EsportOracleTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _EsportOracle.Contract.SetOwner(&_EsportOracle.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_EsportOracle *EsportOracleTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EsportOracle.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_EsportOracle *EsportOracleSession) Unpause() (*types.Transaction, error) {
	return _EsportOracle.Contract.Unpause(&_EsportOracle.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_EsportOracle *EsportOracleTransactorSession) Unpause() (*types.Transaction, error) {
	return _EsportOracle.Contract.Unpause(&_EsportOracle.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_EsportOracle *EsportOracleTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EsportOracle.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_EsportOracle *EsportOracleSession) WithdrawStake() (*types.Transaction, error) {
	return _EsportOracle.Contract.WithdrawStake(&_EsportOracle.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_EsportOracle *EsportOracleTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _EsportOracle.Contract.WithdrawStake(&_EsportOracle.TransactOpts)
}

// EsportOracleNodeBannedIterator is returned from FilterNodeBanned and is used to iterate over the raw logs and unpacked data for NodeBanned events raised by the EsportOracle contract.
type EsportOracleNodeBannedIterator struct {
	Event *EsportOracleNodeBanned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EsportOracleNodeBannedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsportOracleNodeBanned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EsportOracleNodeBanned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EsportOracleNodeBannedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsportOracleNodeBannedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsportOracleNodeBanned represents a NodeBanned event raised by the EsportOracle contract.
type EsportOracleNodeBanned struct {
	Node common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNodeBanned is a free log retrieval operation binding the contract event 0xf1a151967feb9c87e5aa6cd9cdee148015dab9c012ef309968cb91c279e094b4.
//
// Solidity: event NodeBanned(address indexed node)
func (_EsportOracle *EsportOracleFilterer) FilterNodeBanned(opts *bind.FilterOpts, node []common.Address) (*EsportOracleNodeBannedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _EsportOracle.contract.FilterLogs(opts, "NodeBanned", nodeRule)
	if err != nil {
		return nil, err
	}
	return &EsportOracleNodeBannedIterator{contract: _EsportOracle.contract, event: "NodeBanned", logs: logs, sub: sub}, nil
}

// WatchNodeBanned is a free log subscription operation binding the contract event 0xf1a151967feb9c87e5aa6cd9cdee148015dab9c012ef309968cb91c279e094b4.
//
// Solidity: event NodeBanned(address indexed node)
func (_EsportOracle *EsportOracleFilterer) WatchNodeBanned(opts *bind.WatchOpts, sink chan<- *EsportOracleNodeBanned, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _EsportOracle.contract.WatchLogs(opts, "NodeBanned", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsportOracleNodeBanned)
				if err := _EsportOracle.contract.UnpackLog(event, "NodeBanned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNodeBanned is a log parse operation binding the contract event 0xf1a151967feb9c87e5aa6cd9cdee148015dab9c012ef309968cb91c279e094b4.
//
// Solidity: event NodeBanned(address indexed node)
func (_EsportOracle *EsportOracleFilterer) ParseNodeBanned(log types.Log) (*EsportOracleNodeBanned, error) {
	event := new(EsportOracleNodeBanned)
	if err := _EsportOracle.contract.UnpackLog(event, "NodeBanned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EsportOracleNodePunishedIterator is returned from FilterNodePunished and is used to iterate over the raw logs and unpacked data for NodePunished events raised by the EsportOracle contract.
type EsportOracleNodePunishedIterator struct {
	Event *EsportOracleNodePunished // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EsportOracleNodePunishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsportOracleNodePunished)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EsportOracleNodePunished)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EsportOracleNodePunishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsportOracleNodePunishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsportOracleNodePunished represents a NodePunished event raised by the EsportOracle contract.
type EsportOracleNodePunished struct {
	Node            common.Address
	Amount          *big.Int
	ViolationsCount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNodePunished is a free log retrieval operation binding the contract event 0xbb5e0de9b41d911fe88799197a82d842eff4fdea524cdfc0becff3857f0ee3ac.
//
// Solidity: event NodePunished(address indexed node, uint256 amount, uint256 violationsCount)
func (_EsportOracle *EsportOracleFilterer) FilterNodePunished(opts *bind.FilterOpts, node []common.Address) (*EsportOracleNodePunishedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _EsportOracle.contract.FilterLogs(opts, "NodePunished", nodeRule)
	if err != nil {
		return nil, err
	}
	return &EsportOracleNodePunishedIterator{contract: _EsportOracle.contract, event: "NodePunished", logs: logs, sub: sub}, nil
}

// WatchNodePunished is a free log subscription operation binding the contract event 0xbb5e0de9b41d911fe88799197a82d842eff4fdea524cdfc0becff3857f0ee3ac.
//
// Solidity: event NodePunished(address indexed node, uint256 amount, uint256 violationsCount)
func (_EsportOracle *EsportOracleFilterer) WatchNodePunished(opts *bind.WatchOpts, sink chan<- *EsportOracleNodePunished, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _EsportOracle.contract.WatchLogs(opts, "NodePunished", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsportOracleNodePunished)
				if err := _EsportOracle.contract.UnpackLog(event, "NodePunished", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNodePunished is a log parse operation binding the contract event 0xbb5e0de9b41d911fe88799197a82d842eff4fdea524cdfc0becff3857f0ee3ac.
//
// Solidity: event NodePunished(address indexed node, uint256 amount, uint256 violationsCount)
func (_EsportOracle *EsportOracleFilterer) ParseNodePunished(log types.Log) (*EsportOracleNodePunished, error) {
	event := new(EsportOracleNodePunished)
	if err := _EsportOracle.contract.UnpackLog(event, "NodePunished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EsportOraclePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EsportOracle contract.
type EsportOraclePausedIterator struct {
	Event *EsportOraclePaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EsportOraclePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsportOraclePaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EsportOraclePaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EsportOraclePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsportOraclePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsportOraclePaused represents a Paused event raised by the EsportOracle contract.
type EsportOraclePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EsportOracle *EsportOracleFilterer) FilterPaused(opts *bind.FilterOpts) (*EsportOraclePausedIterator, error) {

	logs, sub, err := _EsportOracle.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EsportOraclePausedIterator{contract: _EsportOracle.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EsportOracle *EsportOracleFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EsportOraclePaused) (event.Subscription, error) {

	logs, sub, err := _EsportOracle.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsportOraclePaused)
				if err := _EsportOracle.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EsportOracle *EsportOracleFilterer) ParsePaused(log types.Log) (*EsportOraclePaused, error) {
	event := new(EsportOraclePaused)
	if err := _EsportOracle.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EsportOracleUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EsportOracle contract.
type EsportOracleUnpausedIterator struct {
	Event *EsportOracleUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EsportOracleUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsportOracleUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EsportOracleUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EsportOracleUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsportOracleUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsportOracleUnpaused represents a Unpaused event raised by the EsportOracle contract.
type EsportOracleUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EsportOracle *EsportOracleFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EsportOracleUnpausedIterator, error) {

	logs, sub, err := _EsportOracle.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EsportOracleUnpausedIterator{contract: _EsportOracle.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EsportOracle *EsportOracleFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EsportOracleUnpaused) (event.Subscription, error) {

	logs, sub, err := _EsportOracle.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsportOracleUnpaused)
				if err := _EsportOracle.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EsportOracle *EsportOracleFilterer) ParseUnpaused(log types.Log) (*EsportOracleUnpaused, error) {
	event := new(EsportOracleUnpaused)
	if err := _EsportOracle.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EsportOracleNewNodeAddedIterator is returned from FilterNewNodeAdded and is used to iterate over the raw logs and unpacked data for NewNodeAdded events raised by the EsportOracle contract.
type EsportOracleNewNodeAddedIterator struct {
	Event *EsportOracleNewNodeAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EsportOracleNewNodeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsportOracleNewNodeAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EsportOracleNewNodeAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EsportOracleNewNodeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsportOracleNewNodeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsportOracleNewNodeAdded represents a NewNodeAdded event raised by the EsportOracle contract.
type EsportOracleNewNodeAdded struct {
	AddressAdded common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewNodeAdded is a free log retrieval operation binding the contract event 0xc473661237ac48848af775cb3c9f5bfdf90ed59dbe739ebac0186bc6e61d1984.
//
// Solidity: event newNodeAdded(address indexed addressAdded)
func (_EsportOracle *EsportOracleFilterer) FilterNewNodeAdded(opts *bind.FilterOpts, addressAdded []common.Address) (*EsportOracleNewNodeAddedIterator, error) {

	var addressAddedRule []interface{}
	for _, addressAddedItem := range addressAdded {
		addressAddedRule = append(addressAddedRule, addressAddedItem)
	}

	logs, sub, err := _EsportOracle.contract.FilterLogs(opts, "newNodeAdded", addressAddedRule)
	if err != nil {
		return nil, err
	}
	return &EsportOracleNewNodeAddedIterator{contract: _EsportOracle.contract, event: "newNodeAdded", logs: logs, sub: sub}, nil
}

// WatchNewNodeAdded is a free log subscription operation binding the contract event 0xc473661237ac48848af775cb3c9f5bfdf90ed59dbe739ebac0186bc6e61d1984.
//
// Solidity: event newNodeAdded(address indexed addressAdded)
func (_EsportOracle *EsportOracleFilterer) WatchNewNodeAdded(opts *bind.WatchOpts, sink chan<- *EsportOracleNewNodeAdded, addressAdded []common.Address) (event.Subscription, error) {

	var addressAddedRule []interface{}
	for _, addressAddedItem := range addressAdded {
		addressAddedRule = append(addressAddedRule, addressAddedItem)
	}

	logs, sub, err := _EsportOracle.contract.WatchLogs(opts, "newNodeAdded", addressAddedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsportOracleNewNodeAdded)
				if err := _EsportOracle.contract.UnpackLog(event, "newNodeAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewNodeAdded is a log parse operation binding the contract event 0xc473661237ac48848af775cb3c9f5bfdf90ed59dbe739ebac0186bc6e61d1984.
//
// Solidity: event newNodeAdded(address indexed addressAdded)
func (_EsportOracle *EsportOracleFilterer) ParseNewNodeAdded(log types.Log) (*EsportOracleNewNodeAdded, error) {
	event := new(EsportOracleNewNodeAdded)
	if err := _EsportOracle.contract.UnpackLog(event, "newNodeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EsportOracleStakingSuccessIterator is returned from FilterStakingSuccess and is used to iterate over the raw logs and unpacked data for StakingSuccess events raised by the EsportOracle contract.
type EsportOracleStakingSuccessIterator struct {
	Event *EsportOracleStakingSuccess // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EsportOracleStakingSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsportOracleStakingSuccess)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EsportOracleStakingSuccess)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EsportOracleStakingSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsportOracleStakingSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsportOracleStakingSuccess represents a StakingSuccess event raised by the EsportOracle contract.
type EsportOracleStakingSuccess struct {
	AddressAdded common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakingSuccess is a free log retrieval operation binding the contract event 0xc60400ca3fd7ce7a123251117104ceac79276329b4f7e0413700342999d3cc1d.
//
// Solidity: event stakingSuccess(address indexed addressAdded, uint256 amount)
func (_EsportOracle *EsportOracleFilterer) FilterStakingSuccess(opts *bind.FilterOpts, addressAdded []common.Address) (*EsportOracleStakingSuccessIterator, error) {

	var addressAddedRule []interface{}
	for _, addressAddedItem := range addressAdded {
		addressAddedRule = append(addressAddedRule, addressAddedItem)
	}

	logs, sub, err := _EsportOracle.contract.FilterLogs(opts, "stakingSuccess", addressAddedRule)
	if err != nil {
		return nil, err
	}
	return &EsportOracleStakingSuccessIterator{contract: _EsportOracle.contract, event: "stakingSuccess", logs: logs, sub: sub}, nil
}

// WatchStakingSuccess is a free log subscription operation binding the contract event 0xc60400ca3fd7ce7a123251117104ceac79276329b4f7e0413700342999d3cc1d.
//
// Solidity: event stakingSuccess(address indexed addressAdded, uint256 amount)
func (_EsportOracle *EsportOracleFilterer) WatchStakingSuccess(opts *bind.WatchOpts, sink chan<- *EsportOracleStakingSuccess, addressAdded []common.Address) (event.Subscription, error) {

	var addressAddedRule []interface{}
	for _, addressAddedItem := range addressAdded {
		addressAddedRule = append(addressAddedRule, addressAddedItem)
	}

	logs, sub, err := _EsportOracle.contract.WatchLogs(opts, "stakingSuccess", addressAddedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsportOracleStakingSuccess)
				if err := _EsportOracle.contract.UnpackLog(event, "stakingSuccess", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakingSuccess is a log parse operation binding the contract event 0xc60400ca3fd7ce7a123251117104ceac79276329b4f7e0413700342999d3cc1d.
//
// Solidity: event stakingSuccess(address indexed addressAdded, uint256 amount)
func (_EsportOracle *EsportOracleFilterer) ParseStakingSuccess(log types.Log) (*EsportOracleStakingSuccess, error) {
	event := new(EsportOracleStakingSuccess)
	if err := _EsportOracle.contract.UnpackLog(event, "stakingSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
