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

// EsportOracleGames is an auto generated low-level Go binding around an user-defined struct.
type EsportOracleGames struct {
	Id       *big.Int
	Finished bool
	WinnerId *big.Int
}

// EsportOracleMatch is an auto generated low-level Go binding around an user-defined struct.
type EsportOracleMatch struct {
	Id        *big.Int
	Opponents []EsportOracleOpponents
	Game      []EsportOracleGames
	Result    []EsportOracleResult
	WinnerId  *big.Int
	BeginAt   *big.Int
}

// EsportOracleOpponents is an auto generated low-level Go binding around an user-defined struct.
type EsportOracleOpponents struct {
	Acronym string
	Id      *big.Int
	Name    string
}

// EsportOracleResult is an auto generated low-level Go binding around an user-defined struct.
type EsportOracleResult struct {
	Score  uint8
	TeamId *big.Int
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"NodeBanned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"violationsCount\",\"type\":\"uint256\"}],\"name\":\"NodePunished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addressAdded\",\"type\":\"address\"}],\"name\":\"newNodeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addressAdded\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stakingSuccess\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_VIOLATIONS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUNISHMENT_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_addressByHash\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_fundsStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_matchMapping\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_winnerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_beginAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_matchVotes\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_nodeViolations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"incorrectMatches\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isBanned\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_pendingMatchesHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addFundToStaking\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"banNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getListedNodes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"matchId\",\"type\":\"uint256\"}],\"name\":\"getMatchById\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"_acronym\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"internalType\":\"structEsportOracle.Opponents[]\",\"name\":\"_opponents\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_finished\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_winnerId\",\"type\":\"uint256\"}],\"internalType\":\"structEsportOracle.Games[]\",\"name\":\"_game\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"_score\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_teamId\",\"type\":\"uint256\"}],\"internalType\":\"structEsportOracle.Result[]\",\"name\":\"_result\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_winnerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_beginAt\",\"type\":\"uint256\"}],\"internalType\":\"structEsportOracle.Match\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingMatches\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"_acronym\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"internalType\":\"structEsportOracle.Opponents[]\",\"name\":\"_opponents\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_finished\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_winnerId\",\"type\":\"uint256\"}],\"internalType\":\"structEsportOracle.Games[]\",\"name\":\"_game\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"_score\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_teamId\",\"type\":\"uint256\"}],\"internalType\":\"structEsportOracle.Result[]\",\"name\":\"_result\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_winnerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_beginAt\",\"type\":\"uint256\"}],\"internalType\":\"structEsportOracle.Match[]\",\"name\":\"newMatch\",\"type\":\"tuple[]\"}],\"name\":\"handleNewMatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"rehabilitateNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// MAXVIOLATIONS is a free data retrieval call binding the contract method 0xe8819ff7.
//
// Solidity: function MAX_VIOLATIONS() view returns(uint256)
func (_Contract *ContractCaller) MAXVIOLATIONS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "MAX_VIOLATIONS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXVIOLATIONS is a free data retrieval call binding the contract method 0xe8819ff7.
//
// Solidity: function MAX_VIOLATIONS() view returns(uint256)
func (_Contract *ContractSession) MAXVIOLATIONS() (*big.Int, error) {
	return _Contract.Contract.MAXVIOLATIONS(&_Contract.CallOpts)
}

// MAXVIOLATIONS is a free data retrieval call binding the contract method 0xe8819ff7.
//
// Solidity: function MAX_VIOLATIONS() view returns(uint256)
func (_Contract *ContractCallerSession) MAXVIOLATIONS() (*big.Int, error) {
	return _Contract.Contract.MAXVIOLATIONS(&_Contract.CallOpts)
}

// PUNISHMENTAMOUNT is a free data retrieval call binding the contract method 0xf21b121c.
//
// Solidity: function PUNISHMENT_AMOUNT() view returns(uint256)
func (_Contract *ContractCaller) PUNISHMENTAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "PUNISHMENT_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUNISHMENTAMOUNT is a free data retrieval call binding the contract method 0xf21b121c.
//
// Solidity: function PUNISHMENT_AMOUNT() view returns(uint256)
func (_Contract *ContractSession) PUNISHMENTAMOUNT() (*big.Int, error) {
	return _Contract.Contract.PUNISHMENTAMOUNT(&_Contract.CallOpts)
}

// PUNISHMENTAMOUNT is a free data retrieval call binding the contract method 0xf21b121c.
//
// Solidity: function PUNISHMENT_AMOUNT() view returns(uint256)
func (_Contract *ContractCallerSession) PUNISHMENTAMOUNT() (*big.Int, error) {
	return _Contract.Contract.PUNISHMENTAMOUNT(&_Contract.CallOpts)
}

// AddressByHash is a free data retrieval call binding the contract method 0xbb27ea26.
//
// Solidity: function _addressByHash(bytes32 , uint256 ) view returns(address)
func (_Contract *ContractCaller) AddressByHash(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_addressByHash", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressByHash is a free data retrieval call binding the contract method 0xbb27ea26.
//
// Solidity: function _addressByHash(bytes32 , uint256 ) view returns(address)
func (_Contract *ContractSession) AddressByHash(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _Contract.Contract.AddressByHash(&_Contract.CallOpts, arg0, arg1)
}

// AddressByHash is a free data retrieval call binding the contract method 0xbb27ea26.
//
// Solidity: function _addressByHash(bytes32 , uint256 ) view returns(address)
func (_Contract *ContractCallerSession) AddressByHash(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _Contract.Contract.AddressByHash(&_Contract.CallOpts, arg0, arg1)
}

// FundsStaked is a free data retrieval call binding the contract method 0x652832f1.
//
// Solidity: function _fundsStaked(address ) view returns(uint256)
func (_Contract *ContractCaller) FundsStaked(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_fundsStaked", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundsStaked is a free data retrieval call binding the contract method 0x652832f1.
//
// Solidity: function _fundsStaked(address ) view returns(uint256)
func (_Contract *ContractSession) FundsStaked(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.FundsStaked(&_Contract.CallOpts, arg0)
}

// FundsStaked is a free data retrieval call binding the contract method 0x652832f1.
//
// Solidity: function _fundsStaked(address ) view returns(uint256)
func (_Contract *ContractCallerSession) FundsStaked(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.FundsStaked(&_Contract.CallOpts, arg0)
}

// MatchMapping is a free data retrieval call binding the contract method 0x42d150c9.
//
// Solidity: function _matchMapping(uint256 ) view returns(uint256 _id, uint256 _winnerId, uint256 _beginAt)
func (_Contract *ContractCaller) MatchMapping(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id       *big.Int
	WinnerId *big.Int
	BeginAt  *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_matchMapping", arg0)

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
func (_Contract *ContractSession) MatchMapping(arg0 *big.Int) (struct {
	Id       *big.Int
	WinnerId *big.Int
	BeginAt  *big.Int
}, error) {
	return _Contract.Contract.MatchMapping(&_Contract.CallOpts, arg0)
}

// MatchMapping is a free data retrieval call binding the contract method 0x42d150c9.
//
// Solidity: function _matchMapping(uint256 ) view returns(uint256 _id, uint256 _winnerId, uint256 _beginAt)
func (_Contract *ContractCallerSession) MatchMapping(arg0 *big.Int) (struct {
	Id       *big.Int
	WinnerId *big.Int
	BeginAt  *big.Int
}, error) {
	return _Contract.Contract.MatchMapping(&_Contract.CallOpts, arg0)
}

// MatchVotes is a free data retrieval call binding the contract method 0x1902762f.
//
// Solidity: function _matchVotes(bytes32 ) view returns(uint8)
func (_Contract *ContractCaller) MatchVotes(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_matchVotes", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MatchVotes is a free data retrieval call binding the contract method 0x1902762f.
//
// Solidity: function _matchVotes(bytes32 ) view returns(uint8)
func (_Contract *ContractSession) MatchVotes(arg0 [32]byte) (uint8, error) {
	return _Contract.Contract.MatchVotes(&_Contract.CallOpts, arg0)
}

// MatchVotes is a free data retrieval call binding the contract method 0x1902762f.
//
// Solidity: function _matchVotes(bytes32 ) view returns(uint8)
func (_Contract *ContractCallerSession) MatchVotes(arg0 [32]byte) (uint8, error) {
	return _Contract.Contract.MatchVotes(&_Contract.CallOpts, arg0)
}

// NodeViolations is a free data retrieval call binding the contract method 0x91cac59e.
//
// Solidity: function _nodeViolations(address ) view returns(uint256 incorrectMatches, bool isBanned)
func (_Contract *ContractCaller) NodeViolations(opts *bind.CallOpts, arg0 common.Address) (struct {
	IncorrectMatches *big.Int
	IsBanned         bool
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_nodeViolations", arg0)

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
func (_Contract *ContractSession) NodeViolations(arg0 common.Address) (struct {
	IncorrectMatches *big.Int
	IsBanned         bool
}, error) {
	return _Contract.Contract.NodeViolations(&_Contract.CallOpts, arg0)
}

// NodeViolations is a free data retrieval call binding the contract method 0x91cac59e.
//
// Solidity: function _nodeViolations(address ) view returns(uint256 incorrectMatches, bool isBanned)
func (_Contract *ContractCallerSession) NodeViolations(arg0 common.Address) (struct {
	IncorrectMatches *big.Int
	IsBanned         bool
}, error) {
	return _Contract.Contract.NodeViolations(&_Contract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// PendingMatchesHashes is a free data retrieval call binding the contract method 0xda83225a.
//
// Solidity: function _pendingMatchesHashes(uint256 ) view returns(bytes32)
func (_Contract *ContractCaller) PendingMatchesHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_pendingMatchesHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PendingMatchesHashes is a free data retrieval call binding the contract method 0xda83225a.
//
// Solidity: function _pendingMatchesHashes(uint256 ) view returns(bytes32)
func (_Contract *ContractSession) PendingMatchesHashes(arg0 *big.Int) ([32]byte, error) {
	return _Contract.Contract.PendingMatchesHashes(&_Contract.CallOpts, arg0)
}

// PendingMatchesHashes is a free data retrieval call binding the contract method 0xda83225a.
//
// Solidity: function _pendingMatchesHashes(uint256 ) view returns(bytes32)
func (_Contract *ContractCallerSession) PendingMatchesHashes(arg0 *big.Int) ([32]byte, error) {
	return _Contract.Contract.PendingMatchesHashes(&_Contract.CallOpts, arg0)
}

// GetListedNodes is a free data retrieval call binding the contract method 0x50344f48.
//
// Solidity: function getListedNodes() view returns(address[])
func (_Contract *ContractCaller) GetListedNodes(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getListedNodes")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetListedNodes is a free data retrieval call binding the contract method 0x50344f48.
//
// Solidity: function getListedNodes() view returns(address[])
func (_Contract *ContractSession) GetListedNodes() ([]common.Address, error) {
	return _Contract.Contract.GetListedNodes(&_Contract.CallOpts)
}

// GetListedNodes is a free data retrieval call binding the contract method 0x50344f48.
//
// Solidity: function getListedNodes() view returns(address[])
func (_Contract *ContractCallerSession) GetListedNodes() ([]common.Address, error) {
	return _Contract.Contract.GetListedNodes(&_Contract.CallOpts)
}

// GetMatchById is a free data retrieval call binding the contract method 0xf3c9fd0e.
//
// Solidity: function getMatchById(uint256 matchId) view returns((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256))
func (_Contract *ContractCaller) GetMatchById(opts *bind.CallOpts, matchId *big.Int) (EsportOracleMatch, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getMatchById", matchId)

	if err != nil {
		return *new(EsportOracleMatch), err
	}

	out0 := *abi.ConvertType(out[0], new(EsportOracleMatch)).(*EsportOracleMatch)

	return out0, err

}

// GetMatchById is a free data retrieval call binding the contract method 0xf3c9fd0e.
//
// Solidity: function getMatchById(uint256 matchId) view returns((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256))
func (_Contract *ContractSession) GetMatchById(matchId *big.Int) (EsportOracleMatch, error) {
	return _Contract.Contract.GetMatchById(&_Contract.CallOpts, matchId)
}

// GetMatchById is a free data retrieval call binding the contract method 0xf3c9fd0e.
//
// Solidity: function getMatchById(uint256 matchId) view returns((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256))
func (_Contract *ContractCallerSession) GetMatchById(matchId *big.Int) (EsportOracleMatch, error) {
	return _Contract.Contract.GetMatchById(&_Contract.CallOpts, matchId)
}

// GetPendingMatches is a free data retrieval call binding the contract method 0x5f29d4b1.
//
// Solidity: function getPendingMatches() view returns(bytes32[])
func (_Contract *ContractCaller) GetPendingMatches(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPendingMatches")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetPendingMatches is a free data retrieval call binding the contract method 0x5f29d4b1.
//
// Solidity: function getPendingMatches() view returns(bytes32[])
func (_Contract *ContractSession) GetPendingMatches() ([][32]byte, error) {
	return _Contract.Contract.GetPendingMatches(&_Contract.CallOpts)
}

// GetPendingMatches is a free data retrieval call binding the contract method 0x5f29d4b1.
//
// Solidity: function getPendingMatches() view returns(bytes32[])
func (_Contract *ContractCallerSession) GetPendingMatches() ([][32]byte, error) {
	return _Contract.Contract.GetPendingMatches(&_Contract.CallOpts)
}

// AddFundToStaking is a paid mutator transaction binding the contract method 0x5c0a126d.
//
// Solidity: function addFundToStaking() payable returns()
func (_Contract *ContractTransactor) AddFundToStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addFundToStaking")
}

// AddFundToStaking is a paid mutator transaction binding the contract method 0x5c0a126d.
//
// Solidity: function addFundToStaking() payable returns()
func (_Contract *ContractSession) AddFundToStaking() (*types.Transaction, error) {
	return _Contract.Contract.AddFundToStaking(&_Contract.TransactOpts)
}

// AddFundToStaking is a paid mutator transaction binding the contract method 0x5c0a126d.
//
// Solidity: function addFundToStaking() payable returns()
func (_Contract *ContractTransactorSession) AddFundToStaking() (*types.Transaction, error) {
	return _Contract.Contract.AddFundToStaking(&_Contract.TransactOpts)
}

// BanNode is a paid mutator transaction binding the contract method 0x3faba59e.
//
// Solidity: function banNode(address node) returns()
func (_Contract *ContractTransactor) BanNode(opts *bind.TransactOpts, node common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "banNode", node)
}

// BanNode is a paid mutator transaction binding the contract method 0x3faba59e.
//
// Solidity: function banNode(address node) returns()
func (_Contract *ContractSession) BanNode(node common.Address) (*types.Transaction, error) {
	return _Contract.Contract.BanNode(&_Contract.TransactOpts, node)
}

// BanNode is a paid mutator transaction binding the contract method 0x3faba59e.
//
// Solidity: function banNode(address node) returns()
func (_Contract *ContractTransactorSession) BanNode(node common.Address) (*types.Transaction, error) {
	return _Contract.Contract.BanNode(&_Contract.TransactOpts, node)
}

// HandleNewMatches is a paid mutator transaction binding the contract method 0x3b23e7d7.
//
// Solidity: function handleNewMatches((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256)[] newMatch) returns()
func (_Contract *ContractTransactor) HandleNewMatches(opts *bind.TransactOpts, newMatch []EsportOracleMatch) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "handleNewMatches", newMatch)
}

// HandleNewMatches is a paid mutator transaction binding the contract method 0x3b23e7d7.
//
// Solidity: function handleNewMatches((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256)[] newMatch) returns()
func (_Contract *ContractSession) HandleNewMatches(newMatch []EsportOracleMatch) (*types.Transaction, error) {
	return _Contract.Contract.HandleNewMatches(&_Contract.TransactOpts, newMatch)
}

// HandleNewMatches is a paid mutator transaction binding the contract method 0x3b23e7d7.
//
// Solidity: function handleNewMatches((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256)[] newMatch) returns()
func (_Contract *ContractTransactorSession) HandleNewMatches(newMatch []EsportOracleMatch) (*types.Transaction, error) {
	return _Contract.Contract.HandleNewMatches(&_Contract.TransactOpts, newMatch)
}

// RehabilitateNode is a paid mutator transaction binding the contract method 0x148b56d6.
//
// Solidity: function rehabilitateNode(address node) returns()
func (_Contract *ContractTransactor) RehabilitateNode(opts *bind.TransactOpts, node common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "rehabilitateNode", node)
}

// RehabilitateNode is a paid mutator transaction binding the contract method 0x148b56d6.
//
// Solidity: function rehabilitateNode(address node) returns()
func (_Contract *ContractSession) RehabilitateNode(node common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RehabilitateNode(&_Contract.TransactOpts, node)
}

// RehabilitateNode is a paid mutator transaction binding the contract method 0x148b56d6.
//
// Solidity: function rehabilitateNode(address node) returns()
func (_Contract *ContractTransactorSession) RehabilitateNode(node common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RehabilitateNode(&_Contract.TransactOpts, node)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Contract *ContractTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Contract *ContractSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetOwner(&_Contract.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Contract *ContractTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetOwner(&_Contract.TransactOpts, newOwner)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Contract *ContractTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Contract *ContractSession) WithdrawStake() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawStake(&_Contract.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Contract *ContractTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawStake(&_Contract.TransactOpts)
}

// ContractNodeBannedIterator is returned from FilterNodeBanned and is used to iterate over the raw logs and unpacked data for NodeBanned events raised by the Contract contract.
type ContractNodeBannedIterator struct {
	Event *ContractNodeBanned // Event containing the contract specifics and raw log

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
func (it *ContractNodeBannedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNodeBanned)
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
		it.Event = new(ContractNodeBanned)
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
func (it *ContractNodeBannedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNodeBannedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNodeBanned represents a NodeBanned event raised by the Contract contract.
type ContractNodeBanned struct {
	Node common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNodeBanned is a free log retrieval operation binding the contract event 0xf1a151967feb9c87e5aa6cd9cdee148015dab9c012ef309968cb91c279e094b4.
//
// Solidity: event NodeBanned(address indexed node)
func (_Contract *ContractFilterer) FilterNodeBanned(opts *bind.FilterOpts, node []common.Address) (*ContractNodeBannedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NodeBanned", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ContractNodeBannedIterator{contract: _Contract.contract, event: "NodeBanned", logs: logs, sub: sub}, nil
}

// WatchNodeBanned is a free log subscription operation binding the contract event 0xf1a151967feb9c87e5aa6cd9cdee148015dab9c012ef309968cb91c279e094b4.
//
// Solidity: event NodeBanned(address indexed node)
func (_Contract *ContractFilterer) WatchNodeBanned(opts *bind.WatchOpts, sink chan<- *ContractNodeBanned, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NodeBanned", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNodeBanned)
				if err := _Contract.contract.UnpackLog(event, "NodeBanned", log); err != nil {
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
func (_Contract *ContractFilterer) ParseNodeBanned(log types.Log) (*ContractNodeBanned, error) {
	event := new(ContractNodeBanned)
	if err := _Contract.contract.UnpackLog(event, "NodeBanned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNodePunishedIterator is returned from FilterNodePunished and is used to iterate over the raw logs and unpacked data for NodePunished events raised by the Contract contract.
type ContractNodePunishedIterator struct {
	Event *ContractNodePunished // Event containing the contract specifics and raw log

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
func (it *ContractNodePunishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNodePunished)
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
		it.Event = new(ContractNodePunished)
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
func (it *ContractNodePunishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNodePunishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNodePunished represents a NodePunished event raised by the Contract contract.
type ContractNodePunished struct {
	Node            common.Address
	Amount          *big.Int
	ViolationsCount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNodePunished is a free log retrieval operation binding the contract event 0xbb5e0de9b41d911fe88799197a82d842eff4fdea524cdfc0becff3857f0ee3ac.
//
// Solidity: event NodePunished(address indexed node, uint256 amount, uint256 violationsCount)
func (_Contract *ContractFilterer) FilterNodePunished(opts *bind.FilterOpts, node []common.Address) (*ContractNodePunishedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NodePunished", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ContractNodePunishedIterator{contract: _Contract.contract, event: "NodePunished", logs: logs, sub: sub}, nil
}

// WatchNodePunished is a free log subscription operation binding the contract event 0xbb5e0de9b41d911fe88799197a82d842eff4fdea524cdfc0becff3857f0ee3ac.
//
// Solidity: event NodePunished(address indexed node, uint256 amount, uint256 violationsCount)
func (_Contract *ContractFilterer) WatchNodePunished(opts *bind.WatchOpts, sink chan<- *ContractNodePunished, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NodePunished", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNodePunished)
				if err := _Contract.contract.UnpackLog(event, "NodePunished", log); err != nil {
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
func (_Contract *ContractFilterer) ParseNodePunished(log types.Log) (*ContractNodePunished, error) {
	event := new(ContractNodePunished)
	if err := _Contract.contract.UnpackLog(event, "NodePunished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewNodeAddedIterator is returned from FilterNewNodeAdded and is used to iterate over the raw logs and unpacked data for NewNodeAdded events raised by the Contract contract.
type ContractNewNodeAddedIterator struct {
	Event *ContractNewNodeAdded // Event containing the contract specifics and raw log

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
func (it *ContractNewNodeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewNodeAdded)
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
		it.Event = new(ContractNewNodeAdded)
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
func (it *ContractNewNodeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewNodeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewNodeAdded represents a NewNodeAdded event raised by the Contract contract.
type ContractNewNodeAdded struct {
	AddressAdded common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewNodeAdded is a free log retrieval operation binding the contract event 0xc473661237ac48848af775cb3c9f5bfdf90ed59dbe739ebac0186bc6e61d1984.
//
// Solidity: event newNodeAdded(address indexed addressAdded)
func (_Contract *ContractFilterer) FilterNewNodeAdded(opts *bind.FilterOpts, addressAdded []common.Address) (*ContractNewNodeAddedIterator, error) {

	var addressAddedRule []interface{}
	for _, addressAddedItem := range addressAdded {
		addressAddedRule = append(addressAddedRule, addressAddedItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "newNodeAdded", addressAddedRule)
	if err != nil {
		return nil, err
	}
	return &ContractNewNodeAddedIterator{contract: _Contract.contract, event: "newNodeAdded", logs: logs, sub: sub}, nil
}

// WatchNewNodeAdded is a free log subscription operation binding the contract event 0xc473661237ac48848af775cb3c9f5bfdf90ed59dbe739ebac0186bc6e61d1984.
//
// Solidity: event newNodeAdded(address indexed addressAdded)
func (_Contract *ContractFilterer) WatchNewNodeAdded(opts *bind.WatchOpts, sink chan<- *ContractNewNodeAdded, addressAdded []common.Address) (event.Subscription, error) {

	var addressAddedRule []interface{}
	for _, addressAddedItem := range addressAdded {
		addressAddedRule = append(addressAddedRule, addressAddedItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "newNodeAdded", addressAddedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewNodeAdded)
				if err := _Contract.contract.UnpackLog(event, "newNodeAdded", log); err != nil {
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
func (_Contract *ContractFilterer) ParseNewNodeAdded(log types.Log) (*ContractNewNodeAdded, error) {
	event := new(ContractNewNodeAdded)
	if err := _Contract.contract.UnpackLog(event, "newNodeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStakingSuccessIterator is returned from FilterStakingSuccess and is used to iterate over the raw logs and unpacked data for StakingSuccess events raised by the Contract contract.
type ContractStakingSuccessIterator struct {
	Event *ContractStakingSuccess // Event containing the contract specifics and raw log

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
func (it *ContractStakingSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStakingSuccess)
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
		it.Event = new(ContractStakingSuccess)
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
func (it *ContractStakingSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStakingSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStakingSuccess represents a StakingSuccess event raised by the Contract contract.
type ContractStakingSuccess struct {
	AddressAdded common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakingSuccess is a free log retrieval operation binding the contract event 0xc60400ca3fd7ce7a123251117104ceac79276329b4f7e0413700342999d3cc1d.
//
// Solidity: event stakingSuccess(address indexed addressAdded, uint256 amount)
func (_Contract *ContractFilterer) FilterStakingSuccess(opts *bind.FilterOpts, addressAdded []common.Address) (*ContractStakingSuccessIterator, error) {

	var addressAddedRule []interface{}
	for _, addressAddedItem := range addressAdded {
		addressAddedRule = append(addressAddedRule, addressAddedItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "stakingSuccess", addressAddedRule)
	if err != nil {
		return nil, err
	}
	return &ContractStakingSuccessIterator{contract: _Contract.contract, event: "stakingSuccess", logs: logs, sub: sub}, nil
}

// WatchStakingSuccess is a free log subscription operation binding the contract event 0xc60400ca3fd7ce7a123251117104ceac79276329b4f7e0413700342999d3cc1d.
//
// Solidity: event stakingSuccess(address indexed addressAdded, uint256 amount)
func (_Contract *ContractFilterer) WatchStakingSuccess(opts *bind.WatchOpts, sink chan<- *ContractStakingSuccess, addressAdded []common.Address) (event.Subscription, error) {

	var addressAddedRule []interface{}
	for _, addressAddedItem := range addressAdded {
		addressAddedRule = append(addressAddedRule, addressAddedItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "stakingSuccess", addressAddedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStakingSuccess)
				if err := _Contract.contract.UnpackLog(event, "stakingSuccess", log); err != nil {
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
func (_Contract *ContractFilterer) ParseStakingSuccess(log types.Log) (*ContractStakingSuccess, error) {
	event := new(ContractStakingSuccess)
	if err := _Contract.contract.UnpackLog(event, "stakingSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
