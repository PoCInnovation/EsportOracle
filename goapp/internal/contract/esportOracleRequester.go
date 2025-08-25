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

// EsportOracleRequesterMetaData contains all meta data concerning the EsportOracleRequester contract.
var EsportOracleRequesterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MAX_VIOLATIONS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_REQUEST_FEE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PUNISHMENT_AMOUNT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_addressByHash\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_fundsStaked\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_matchMapping\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_winnerId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_beginAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_matchRequests\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"matchId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requester\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_matchVotes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_nodeViolations\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"incorrectMatches\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isBanned\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_pendingMatchesHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addFundToStaking\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"banNode\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"callMatchOracle\",\"inputs\":[{\"name\":\"matchId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"matchData\",\"type\":\"tuple\",\"internalType\":\"structEsportOracleTypes.Match\",\"components\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_opponents\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Opponents[]\",\"components\":[{\"name\":\"_acronym\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"_game\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Games[]\",\"components\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_finished\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_winnerId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_result\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Result[]\",\"components\":[{\"name\":\"_score\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_teamId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_winnerId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_beginAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkQorum\",\"inputs\":[{\"name\":\"matchData\",\"type\":\"tuple\",\"internalType\":\"structEsportOracleTypes.Match\",\"components\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_opponents\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Opponents[]\",\"components\":[{\"name\":\"_acronym\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"_game\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Games[]\",\"components\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_finished\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_winnerId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_result\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Result[]\",\"components\":[{\"name\":\"_score\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_teamId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_winnerId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_beginAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getListedNodes\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMatchById\",\"inputs\":[{\"name\":\"matchId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structEsportOracleTypes.Match\",\"components\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_opponents\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Opponents[]\",\"components\":[{\"name\":\"_acronym\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"_game\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Games[]\",\"components\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_finished\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_winnerId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_result\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Result[]\",\"components\":[{\"name\":\"_score\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_teamId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_winnerId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_beginAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMatchRequest\",\"inputs\":[{\"name\":\"matchId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structEsportOracleTypes.MatchRequest\",\"components\":[{\"name\":\"matchId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requester\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingMatches\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingRequestedMatches\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequestByMatchId\",\"inputs\":[{\"name\":\"matchId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structEsportOracleTypes.MatchRequest\",\"components\":[{\"name\":\"matchId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requester\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"handleNewMatches\",\"inputs\":[{\"name\":\"newMatch\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Match[]\",\"components\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_opponents\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Opponents[]\",\"components\":[{\"name\":\"_acronym\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"_game\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Games[]\",\"components\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_finished\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_winnerId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_result\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Result[]\",\"components\":[{\"name\":\"_score\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_teamId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_winnerId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_beginAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isMatchRequested\",\"inputs\":[{\"name\":\"matchId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"markRequestsFulfilled\",\"inputs\":[{\"name\":\"matchId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rehabilitateNode\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestMatch\",\"inputs\":[{\"name\":\"matchId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setOwner\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawStake\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"MatchRequested\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"matchId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"requester\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeBanned\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodePunished\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"violationsCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"newNodeAdded\",\"inputs\":[{\"name\":\"addressAdded\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"stakingSuccess\",\"inputs\":[{\"name\":\"addressAdded\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]}]",
}

// EsportOracleRequesterABI is the input ABI used to generate the binding from.
// Deprecated: Use EsportOracleRequesterMetaData.ABI instead.
var EsportOracleRequesterABI = EsportOracleRequesterMetaData.ABI

// EsportOracleRequester is an auto generated Go binding around an Ethereum contract.
type EsportOracleRequester struct {
	EsportOracleRequesterCaller     // Read-only binding to the contract
	EsportOracleRequesterTransactor // Write-only binding to the contract
	EsportOracleRequesterFilterer   // Log filterer for contract events
}

// EsportOracleRequesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type EsportOracleRequesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EsportOracleRequesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EsportOracleRequesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EsportOracleRequesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EsportOracleRequesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EsportOracleRequesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EsportOracleRequesterSession struct {
	Contract     *EsportOracleRequester // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EsportOracleRequesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EsportOracleRequesterCallerSession struct {
	Contract *EsportOracleRequesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// EsportOracleRequesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EsportOracleRequesterTransactorSession struct {
	Contract     *EsportOracleRequesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// EsportOracleRequesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type EsportOracleRequesterRaw struct {
	Contract *EsportOracleRequester // Generic contract binding to access the raw methods on
}

// EsportOracleRequesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EsportOracleRequesterCallerRaw struct {
	Contract *EsportOracleRequesterCaller // Generic read-only contract binding to access the raw methods on
}

// EsportOracleRequesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EsportOracleRequesterTransactorRaw struct {
	Contract *EsportOracleRequesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEsportOracleRequester creates a new instance of EsportOracleRequester, bound to a specific deployed contract.
func NewEsportOracleRequester(address common.Address, backend bind.ContractBackend) (*EsportOracleRequester, error) {
	contract, err := bindEsportOracleRequester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EsportOracleRequester{EsportOracleRequesterCaller: EsportOracleRequesterCaller{contract: contract}, EsportOracleRequesterTransactor: EsportOracleRequesterTransactor{contract: contract}, EsportOracleRequesterFilterer: EsportOracleRequesterFilterer{contract: contract}}, nil
}

// NewEsportOracleRequesterCaller creates a new read-only instance of EsportOracleRequester, bound to a specific deployed contract.
func NewEsportOracleRequesterCaller(address common.Address, caller bind.ContractCaller) (*EsportOracleRequesterCaller, error) {
	contract, err := bindEsportOracleRequester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EsportOracleRequesterCaller{contract: contract}, nil
}

// NewEsportOracleRequesterTransactor creates a new write-only instance of EsportOracleRequester, bound to a specific deployed contract.
func NewEsportOracleRequesterTransactor(address common.Address, transactor bind.ContractTransactor) (*EsportOracleRequesterTransactor, error) {
	contract, err := bindEsportOracleRequester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EsportOracleRequesterTransactor{contract: contract}, nil
}

// NewEsportOracleRequesterFilterer creates a new log filterer instance of EsportOracleRequester, bound to a specific deployed contract.
func NewEsportOracleRequesterFilterer(address common.Address, filterer bind.ContractFilterer) (*EsportOracleRequesterFilterer, error) {
	contract, err := bindEsportOracleRequester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EsportOracleRequesterFilterer{contract: contract}, nil
}

// bindEsportOracleRequester binds a generic wrapper to an already deployed contract.
func bindEsportOracleRequester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EsportOracleRequesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EsportOracleRequester *EsportOracleRequesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EsportOracleRequester.Contract.EsportOracleRequesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EsportOracleRequester *EsportOracleRequesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.EsportOracleRequesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EsportOracleRequester *EsportOracleRequesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.EsportOracleRequesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EsportOracleRequester *EsportOracleRequesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EsportOracleRequester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EsportOracleRequester *EsportOracleRequesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EsportOracleRequester *EsportOracleRequesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.contract.Transact(opts, method, params...)
}

// MAXVIOLATIONS is a free data retrieval call binding the contract method 0xe8819ff7.
//
// Solidity: function MAX_VIOLATIONS() view returns(uint256)
func (_EsportOracleRequester *EsportOracleRequesterCaller) MAXVIOLATIONS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EsportOracleRequester.contract.Call(opts, &out, "MAX_VIOLATIONS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXVIOLATIONS is a free data retrieval call binding the contract method 0xe8819ff7.
//
// Solidity: function MAX_VIOLATIONS() view returns(uint256)
func (_EsportOracleRequester *EsportOracleRequesterSession) MAXVIOLATIONS() (*big.Int, error) {
	return _EsportOracleRequester.Contract.MAXVIOLATIONS(&_EsportOracleRequester.CallOpts)
}

// MAXVIOLATIONS is a free data retrieval call binding the contract method 0xe8819ff7.
//
// Solidity: function MAX_VIOLATIONS() view returns(uint256)
func (_EsportOracleRequester *EsportOracleRequesterCallerSession) MAXVIOLATIONS() (*big.Int, error) {
	return _EsportOracleRequester.Contract.MAXVIOLATIONS(&_EsportOracleRequester.CallOpts)
}

// MINREQUESTFEE is a free data retrieval call binding the contract method 0xb3ed78b8.
//
// Solidity: function MIN_REQUEST_FEE() view returns(uint256)
func (_EsportOracleRequester *EsportOracleRequesterCaller) MINREQUESTFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EsportOracleRequester.contract.Call(opts, &out, "MIN_REQUEST_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINREQUESTFEE is a free data retrieval call binding the contract method 0xb3ed78b8.
//
// Solidity: function MIN_REQUEST_FEE() view returns(uint256)
func (_EsportOracleRequester *EsportOracleRequesterSession) MINREQUESTFEE() (*big.Int, error) {
	return _EsportOracleRequester.Contract.MINREQUESTFEE(&_EsportOracleRequester.CallOpts)
}

// MINREQUESTFEE is a free data retrieval call binding the contract method 0xb3ed78b8.
//
// Solidity: function MIN_REQUEST_FEE() view returns(uint256)
func (_EsportOracleRequester *EsportOracleRequesterCallerSession) MINREQUESTFEE() (*big.Int, error) {
	return _EsportOracleRequester.Contract.MINREQUESTFEE(&_EsportOracleRequester.CallOpts)
}

// PUNISHMENTAMOUNT is a free data retrieval call binding the contract method 0xf21b121c.
//
// Solidity: function PUNISHMENT_AMOUNT() view returns(uint256)
func (_EsportOracleRequester *EsportOracleRequesterCaller) PUNISHMENTAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EsportOracleRequester.contract.Call(opts, &out, "PUNISHMENT_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUNISHMENTAMOUNT is a free data retrieval call binding the contract method 0xf21b121c.
//
// Solidity: function PUNISHMENT_AMOUNT() view returns(uint256)
func (_EsportOracleRequester *EsportOracleRequesterSession) PUNISHMENTAMOUNT() (*big.Int, error) {
	return _EsportOracleRequester.Contract.PUNISHMENTAMOUNT(&_EsportOracleRequester.CallOpts)
}

// PUNISHMENTAMOUNT is a free data retrieval call binding the contract method 0xf21b121c.
//
// Solidity: function PUNISHMENT_AMOUNT() view returns(uint256)
func (_EsportOracleRequester *EsportOracleRequesterCallerSession) PUNISHMENTAMOUNT() (*big.Int, error) {
	return _EsportOracleRequester.Contract.PUNISHMENTAMOUNT(&_EsportOracleRequester.CallOpts)
}

// AddressByHash is a free data retrieval call binding the contract method 0xbb27ea26.
//
// Solidity: function _addressByHash(bytes32 , uint256 ) view returns(address)
func (_EsportOracleRequester *EsportOracleRequesterCaller) AddressByHash(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EsportOracleRequester.contract.Call(opts, &out, "_addressByHash", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressByHash is a free data retrieval call binding the contract method 0xbb27ea26.
//
// Solidity: function _addressByHash(bytes32 , uint256 ) view returns(address)
func (_EsportOracleRequester *EsportOracleRequesterSession) AddressByHash(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _EsportOracleRequester.Contract.AddressByHash(&_EsportOracleRequester.CallOpts, arg0, arg1)
}

// AddressByHash is a free data retrieval call binding the contract method 0xbb27ea26.
//
// Solidity: function _addressByHash(bytes32 , uint256 ) view returns(address)
func (_EsportOracleRequester *EsportOracleRequesterCallerSession) AddressByHash(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _EsportOracleRequester.Contract.AddressByHash(&_EsportOracleRequester.CallOpts, arg0, arg1)
}

// FundsStaked is a free data retrieval call binding the contract method 0x652832f1.
//
// Solidity: function _fundsStaked(address ) view returns(uint256)
func (_EsportOracleRequester *EsportOracleRequesterCaller) FundsStaked(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EsportOracleRequester.contract.Call(opts, &out, "_fundsStaked", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundsStaked is a free data retrieval call binding the contract method 0x652832f1.
//
// Solidity: function _fundsStaked(address ) view returns(uint256)
func (_EsportOracleRequester *EsportOracleRequesterSession) FundsStaked(arg0 common.Address) (*big.Int, error) {
	return _EsportOracleRequester.Contract.FundsStaked(&_EsportOracleRequester.CallOpts, arg0)
}

// FundsStaked is a free data retrieval call binding the contract method 0x652832f1.
//
// Solidity: function _fundsStaked(address ) view returns(uint256)
func (_EsportOracleRequester *EsportOracleRequesterCallerSession) FundsStaked(arg0 common.Address) (*big.Int, error) {
	return _EsportOracleRequester.Contract.FundsStaked(&_EsportOracleRequester.CallOpts, arg0)
}

// MatchMapping is a free data retrieval call binding the contract method 0x42d150c9.
//
// Solidity: function _matchMapping(uint256 ) view returns(uint256 _id, uint256 _winnerId, uint256 _beginAt)
func (_EsportOracleRequester *EsportOracleRequesterCaller) MatchMapping(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id       *big.Int
	WinnerId *big.Int
	BeginAt  *big.Int
}, error) {
	var out []interface{}
	err := _EsportOracleRequester.contract.Call(opts, &out, "_matchMapping", arg0)

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
func (_EsportOracleRequester *EsportOracleRequesterSession) MatchMapping(arg0 *big.Int) (struct {
	Id       *big.Int
	WinnerId *big.Int
	BeginAt  *big.Int
}, error) {
	return _EsportOracleRequester.Contract.MatchMapping(&_EsportOracleRequester.CallOpts, arg0)
}

// MatchMapping is a free data retrieval call binding the contract method 0x42d150c9.
//
// Solidity: function _matchMapping(uint256 ) view returns(uint256 _id, uint256 _winnerId, uint256 _beginAt)
func (_EsportOracleRequester *EsportOracleRequesterCallerSession) MatchMapping(arg0 *big.Int) (struct {
	Id       *big.Int
	WinnerId *big.Int
	BeginAt  *big.Int
}, error) {
	return _EsportOracleRequester.Contract.MatchMapping(&_EsportOracleRequester.CallOpts, arg0)
}

// MatchRequests is a free data retrieval call binding the contract method 0xbd61d67e.
//
// Solidity: function _matchRequests(uint256 ) view returns(uint256 matchId, address requester, uint256 fee, bool fulfilled)
func (_EsportOracleRequester *EsportOracleRequesterCaller) MatchRequests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	MatchId   *big.Int
	Requester common.Address
	Fee       *big.Int
	Fulfilled bool
}, error) {
	var out []interface{}
	err := _EsportOracleRequester.contract.Call(opts, &out, "_matchRequests", arg0)

	outstruct := new(struct {
		MatchId   *big.Int
		Requester common.Address
		Fee       *big.Int
		Fulfilled bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MatchId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Requester = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Fulfilled = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// MatchRequests is a free data retrieval call binding the contract method 0xbd61d67e.
//
// Solidity: function _matchRequests(uint256 ) view returns(uint256 matchId, address requester, uint256 fee, bool fulfilled)
func (_EsportOracleRequester *EsportOracleRequesterSession) MatchRequests(arg0 *big.Int) (struct {
	MatchId   *big.Int
	Requester common.Address
	Fee       *big.Int
	Fulfilled bool
}, error) {
	return _EsportOracleRequester.Contract.MatchRequests(&_EsportOracleRequester.CallOpts, arg0)
}

// MatchRequests is a free data retrieval call binding the contract method 0xbd61d67e.
//
// Solidity: function _matchRequests(uint256 ) view returns(uint256 matchId, address requester, uint256 fee, bool fulfilled)
func (_EsportOracleRequester *EsportOracleRequesterCallerSession) MatchRequests(arg0 *big.Int) (struct {
	MatchId   *big.Int
	Requester common.Address
	Fee       *big.Int
	Fulfilled bool
}, error) {
	return _EsportOracleRequester.Contract.MatchRequests(&_EsportOracleRequester.CallOpts, arg0)
}

// MatchVotes is a free data retrieval call binding the contract method 0x1902762f.
//
// Solidity: function _matchVotes(bytes32 ) view returns(uint8)
func (_EsportOracleRequester *EsportOracleRequesterCaller) MatchVotes(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _EsportOracleRequester.contract.Call(opts, &out, "_matchVotes", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MatchVotes is a free data retrieval call binding the contract method 0x1902762f.
//
// Solidity: function _matchVotes(bytes32 ) view returns(uint8)
func (_EsportOracleRequester *EsportOracleRequesterSession) MatchVotes(arg0 [32]byte) (uint8, error) {
	return _EsportOracleRequester.Contract.MatchVotes(&_EsportOracleRequester.CallOpts, arg0)
}

// MatchVotes is a free data retrieval call binding the contract method 0x1902762f.
//
// Solidity: function _matchVotes(bytes32 ) view returns(uint8)
func (_EsportOracleRequester *EsportOracleRequesterCallerSession) MatchVotes(arg0 [32]byte) (uint8, error) {
	return _EsportOracleRequester.Contract.MatchVotes(&_EsportOracleRequester.CallOpts, arg0)
}

// NodeViolations is a free data retrieval call binding the contract method 0x91cac59e.
//
// Solidity: function _nodeViolations(address ) view returns(uint256 incorrectMatches, bool isBanned)
func (_EsportOracleRequester *EsportOracleRequesterCaller) NodeViolations(opts *bind.CallOpts, arg0 common.Address) (struct {
	IncorrectMatches *big.Int
	IsBanned         bool
}, error) {
	var out []interface{}
	err := _EsportOracleRequester.contract.Call(opts, &out, "_nodeViolations", arg0)

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
func (_EsportOracleRequester *EsportOracleRequesterSession) NodeViolations(arg0 common.Address) (struct {
	IncorrectMatches *big.Int
	IsBanned         bool
}, error) {
	return _EsportOracleRequester.Contract.NodeViolations(&_EsportOracleRequester.CallOpts, arg0)
}

// NodeViolations is a free data retrieval call binding the contract method 0x91cac59e.
//
// Solidity: function _nodeViolations(address ) view returns(uint256 incorrectMatches, bool isBanned)
func (_EsportOracleRequester *EsportOracleRequesterCallerSession) NodeViolations(arg0 common.Address) (struct {
	IncorrectMatches *big.Int
	IsBanned         bool
}, error) {
	return _EsportOracleRequester.Contract.NodeViolations(&_EsportOracleRequester.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_EsportOracleRequester *EsportOracleRequesterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EsportOracleRequester.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_EsportOracleRequester *EsportOracleRequesterSession) Owner() (common.Address, error) {
	return _EsportOracleRequester.Contract.Owner(&_EsportOracleRequester.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_EsportOracleRequester *EsportOracleRequesterCallerSession) Owner() (common.Address, error) {
	return _EsportOracleRequester.Contract.Owner(&_EsportOracleRequester.CallOpts)
}

// PendingMatchesHashes is a free data retrieval call binding the contract method 0xda83225a.
//
// Solidity: function _pendingMatchesHashes(uint256 ) view returns(bytes32)
func (_EsportOracleRequester *EsportOracleRequesterCaller) PendingMatchesHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _EsportOracleRequester.contract.Call(opts, &out, "_pendingMatchesHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PendingMatchesHashes is a free data retrieval call binding the contract method 0xda83225a.
//
// Solidity: function _pendingMatchesHashes(uint256 ) view returns(bytes32)
func (_EsportOracleRequester *EsportOracleRequesterSession) PendingMatchesHashes(arg0 *big.Int) ([32]byte, error) {
	return _EsportOracleRequester.Contract.PendingMatchesHashes(&_EsportOracleRequester.CallOpts, arg0)
}

// PendingMatchesHashes is a free data retrieval call binding the contract method 0xda83225a.
//
// Solidity: function _pendingMatchesHashes(uint256 ) view returns(bytes32)
func (_EsportOracleRequester *EsportOracleRequesterCallerSession) PendingMatchesHashes(arg0 *big.Int) ([32]byte, error) {
	return _EsportOracleRequester.Contract.PendingMatchesHashes(&_EsportOracleRequester.CallOpts, arg0)
}

// CheckQorum is a free data retrieval call binding the contract method 0x31ecedd1.
//
// Solidity: function checkQorum((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256) matchData) view returns(bool)
func (_EsportOracleRequester *EsportOracleRequesterCaller) CheckQorum(opts *bind.CallOpts, matchData EsportOracleTypesMatch) (bool, error) {
	var out []interface{}
	err := _EsportOracleRequester.contract.Call(opts, &out, "checkQorum", matchData)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckQorum is a free data retrieval call binding the contract method 0x31ecedd1.
//
// Solidity: function checkQorum((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256) matchData) view returns(bool)
func (_EsportOracleRequester *EsportOracleRequesterSession) CheckQorum(matchData EsportOracleTypesMatch) (bool, error) {
	return _EsportOracleRequester.Contract.CheckQorum(&_EsportOracleRequester.CallOpts, matchData)
}

// CheckQorum is a free data retrieval call binding the contract method 0x31ecedd1.
//
// Solidity: function checkQorum((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256) matchData) view returns(bool)
func (_EsportOracleRequester *EsportOracleRequesterCallerSession) CheckQorum(matchData EsportOracleTypesMatch) (bool, error) {
	return _EsportOracleRequester.Contract.CheckQorum(&_EsportOracleRequester.CallOpts, matchData)
}

// GetListedNodes is a free data retrieval call binding the contract method 0x50344f48.
//
// Solidity: function getListedNodes() view returns(address[])
func (_EsportOracleRequester *EsportOracleRequesterCaller) GetListedNodes(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EsportOracleRequester.contract.Call(opts, &out, "getListedNodes")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetListedNodes is a free data retrieval call binding the contract method 0x50344f48.
//
// Solidity: function getListedNodes() view returns(address[])
func (_EsportOracleRequester *EsportOracleRequesterSession) GetListedNodes() ([]common.Address, error) {
	return _EsportOracleRequester.Contract.GetListedNodes(&_EsportOracleRequester.CallOpts)
}

// GetListedNodes is a free data retrieval call binding the contract method 0x50344f48.
//
// Solidity: function getListedNodes() view returns(address[])
func (_EsportOracleRequester *EsportOracleRequesterCallerSession) GetListedNodes() ([]common.Address, error) {
	return _EsportOracleRequester.Contract.GetListedNodes(&_EsportOracleRequester.CallOpts)
}

// GetMatchById is a free data retrieval call binding the contract method 0xf3c9fd0e.
//
// Solidity: function getMatchById(uint256 matchId) view returns((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256))
func (_EsportOracleRequester *EsportOracleRequesterCaller) GetMatchById(opts *bind.CallOpts, matchId *big.Int) (EsportOracleTypesMatch, error) {
	var out []interface{}
	err := _EsportOracleRequester.contract.Call(opts, &out, "getMatchById", matchId)

	if err != nil {
		return *new(EsportOracleTypesMatch), err
	}

	out0 := *abi.ConvertType(out[0], new(EsportOracleTypesMatch)).(*EsportOracleTypesMatch)

	return out0, err

}

// GetMatchById is a free data retrieval call binding the contract method 0xf3c9fd0e.
//
// Solidity: function getMatchById(uint256 matchId) view returns((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256))
func (_EsportOracleRequester *EsportOracleRequesterSession) GetMatchById(matchId *big.Int) (EsportOracleTypesMatch, error) {
	return _EsportOracleRequester.Contract.GetMatchById(&_EsportOracleRequester.CallOpts, matchId)
}

// GetMatchById is a free data retrieval call binding the contract method 0xf3c9fd0e.
//
// Solidity: function getMatchById(uint256 matchId) view returns((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256))
func (_EsportOracleRequester *EsportOracleRequesterCallerSession) GetMatchById(matchId *big.Int) (EsportOracleTypesMatch, error) {
	return _EsportOracleRequester.Contract.GetMatchById(&_EsportOracleRequester.CallOpts, matchId)
}

// GetMatchRequest is a free data retrieval call binding the contract method 0x72b1aed2.
//
// Solidity: function getMatchRequest(uint256 matchId) view returns((uint256,address,uint256,bool))
func (_EsportOracleRequester *EsportOracleRequesterCaller) GetMatchRequest(opts *bind.CallOpts, matchId *big.Int) (EsportOracleTypesMatchRequest, error) {
	var out []interface{}
	err := _EsportOracleRequester.contract.Call(opts, &out, "getMatchRequest", matchId)

	if err != nil {
		return *new(EsportOracleTypesMatchRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(EsportOracleTypesMatchRequest)).(*EsportOracleTypesMatchRequest)

	return out0, err

}

// GetMatchRequest is a free data retrieval call binding the contract method 0x72b1aed2.
//
// Solidity: function getMatchRequest(uint256 matchId) view returns((uint256,address,uint256,bool))
func (_EsportOracleRequester *EsportOracleRequesterSession) GetMatchRequest(matchId *big.Int) (EsportOracleTypesMatchRequest, error) {
	return _EsportOracleRequester.Contract.GetMatchRequest(&_EsportOracleRequester.CallOpts, matchId)
}

// GetMatchRequest is a free data retrieval call binding the contract method 0x72b1aed2.
//
// Solidity: function getMatchRequest(uint256 matchId) view returns((uint256,address,uint256,bool))
func (_EsportOracleRequester *EsportOracleRequesterCallerSession) GetMatchRequest(matchId *big.Int) (EsportOracleTypesMatchRequest, error) {
	return _EsportOracleRequester.Contract.GetMatchRequest(&_EsportOracleRequester.CallOpts, matchId)
}

// GetPendingMatches is a free data retrieval call binding the contract method 0x5f29d4b1.
//
// Solidity: function getPendingMatches() view returns(bytes32[])
func (_EsportOracleRequester *EsportOracleRequesterCaller) GetPendingMatches(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _EsportOracleRequester.contract.Call(opts, &out, "getPendingMatches")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetPendingMatches is a free data retrieval call binding the contract method 0x5f29d4b1.
//
// Solidity: function getPendingMatches() view returns(bytes32[])
func (_EsportOracleRequester *EsportOracleRequesterSession) GetPendingMatches() ([][32]byte, error) {
	return _EsportOracleRequester.Contract.GetPendingMatches(&_EsportOracleRequester.CallOpts)
}

// GetPendingMatches is a free data retrieval call binding the contract method 0x5f29d4b1.
//
// Solidity: function getPendingMatches() view returns(bytes32[])
func (_EsportOracleRequester *EsportOracleRequesterCallerSession) GetPendingMatches() ([][32]byte, error) {
	return _EsportOracleRequester.Contract.GetPendingMatches(&_EsportOracleRequester.CallOpts)
}

// GetPendingRequestedMatches is a free data retrieval call binding the contract method 0x2f881c1b.
//
// Solidity: function getPendingRequestedMatches() view returns(uint256[])
func (_EsportOracleRequester *EsportOracleRequesterCaller) GetPendingRequestedMatches(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _EsportOracleRequester.contract.Call(opts, &out, "getPendingRequestedMatches")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPendingRequestedMatches is a free data retrieval call binding the contract method 0x2f881c1b.
//
// Solidity: function getPendingRequestedMatches() view returns(uint256[])
func (_EsportOracleRequester *EsportOracleRequesterSession) GetPendingRequestedMatches() ([]*big.Int, error) {
	return _EsportOracleRequester.Contract.GetPendingRequestedMatches(&_EsportOracleRequester.CallOpts)
}

// GetPendingRequestedMatches is a free data retrieval call binding the contract method 0x2f881c1b.
//
// Solidity: function getPendingRequestedMatches() view returns(uint256[])
func (_EsportOracleRequester *EsportOracleRequesterCallerSession) GetPendingRequestedMatches() ([]*big.Int, error) {
	return _EsportOracleRequester.Contract.GetPendingRequestedMatches(&_EsportOracleRequester.CallOpts)
}

// GetRequestByMatchId is a free data retrieval call binding the contract method 0x2cbe0c0e.
//
// Solidity: function getRequestByMatchId(uint256 matchId) view returns((uint256,address,uint256,bool))
func (_EsportOracleRequester *EsportOracleRequesterCaller) GetRequestByMatchId(opts *bind.CallOpts, matchId *big.Int) (EsportOracleTypesMatchRequest, error) {
	var out []interface{}
	err := _EsportOracleRequester.contract.Call(opts, &out, "getRequestByMatchId", matchId)

	if err != nil {
		return *new(EsportOracleTypesMatchRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(EsportOracleTypesMatchRequest)).(*EsportOracleTypesMatchRequest)

	return out0, err

}

// GetRequestByMatchId is a free data retrieval call binding the contract method 0x2cbe0c0e.
//
// Solidity: function getRequestByMatchId(uint256 matchId) view returns((uint256,address,uint256,bool))
func (_EsportOracleRequester *EsportOracleRequesterSession) GetRequestByMatchId(matchId *big.Int) (EsportOracleTypesMatchRequest, error) {
	return _EsportOracleRequester.Contract.GetRequestByMatchId(&_EsportOracleRequester.CallOpts, matchId)
}

// GetRequestByMatchId is a free data retrieval call binding the contract method 0x2cbe0c0e.
//
// Solidity: function getRequestByMatchId(uint256 matchId) view returns((uint256,address,uint256,bool))
func (_EsportOracleRequester *EsportOracleRequesterCallerSession) GetRequestByMatchId(matchId *big.Int) (EsportOracleTypesMatchRequest, error) {
	return _EsportOracleRequester.Contract.GetRequestByMatchId(&_EsportOracleRequester.CallOpts, matchId)
}

// IsMatchRequested is a free data retrieval call binding the contract method 0xb3b8bf05.
//
// Solidity: function isMatchRequested(uint256 matchId) view returns(bool)
func (_EsportOracleRequester *EsportOracleRequesterCaller) IsMatchRequested(opts *bind.CallOpts, matchId *big.Int) (bool, error) {
	var out []interface{}
	err := _EsportOracleRequester.contract.Call(opts, &out, "isMatchRequested", matchId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMatchRequested is a free data retrieval call binding the contract method 0xb3b8bf05.
//
// Solidity: function isMatchRequested(uint256 matchId) view returns(bool)
func (_EsportOracleRequester *EsportOracleRequesterSession) IsMatchRequested(matchId *big.Int) (bool, error) {
	return _EsportOracleRequester.Contract.IsMatchRequested(&_EsportOracleRequester.CallOpts, matchId)
}

// IsMatchRequested is a free data retrieval call binding the contract method 0xb3b8bf05.
//
// Solidity: function isMatchRequested(uint256 matchId) view returns(bool)
func (_EsportOracleRequester *EsportOracleRequesterCallerSession) IsMatchRequested(matchId *big.Int) (bool, error) {
	return _EsportOracleRequester.Contract.IsMatchRequested(&_EsportOracleRequester.CallOpts, matchId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EsportOracleRequester *EsportOracleRequesterCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EsportOracleRequester.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EsportOracleRequester *EsportOracleRequesterSession) Paused() (bool, error) {
	return _EsportOracleRequester.Contract.Paused(&_EsportOracleRequester.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EsportOracleRequester *EsportOracleRequesterCallerSession) Paused() (bool, error) {
	return _EsportOracleRequester.Contract.Paused(&_EsportOracleRequester.CallOpts)
}

// AddFundToStaking is a paid mutator transaction binding the contract method 0x5c0a126d.
//
// Solidity: function addFundToStaking() payable returns()
func (_EsportOracleRequester *EsportOracleRequesterTransactor) AddFundToStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EsportOracleRequester.contract.Transact(opts, "addFundToStaking")
}

// AddFundToStaking is a paid mutator transaction binding the contract method 0x5c0a126d.
//
// Solidity: function addFundToStaking() payable returns()
func (_EsportOracleRequester *EsportOracleRequesterSession) AddFundToStaking() (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.AddFundToStaking(&_EsportOracleRequester.TransactOpts)
}

// AddFundToStaking is a paid mutator transaction binding the contract method 0x5c0a126d.
//
// Solidity: function addFundToStaking() payable returns()
func (_EsportOracleRequester *EsportOracleRequesterTransactorSession) AddFundToStaking() (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.AddFundToStaking(&_EsportOracleRequester.TransactOpts)
}

// BanNode is a paid mutator transaction binding the contract method 0x3faba59e.
//
// Solidity: function banNode(address node) returns()
func (_EsportOracleRequester *EsportOracleRequesterTransactor) BanNode(opts *bind.TransactOpts, node common.Address) (*types.Transaction, error) {
	return _EsportOracleRequester.contract.Transact(opts, "banNode", node)
}

// BanNode is a paid mutator transaction binding the contract method 0x3faba59e.
//
// Solidity: function banNode(address node) returns()
func (_EsportOracleRequester *EsportOracleRequesterSession) BanNode(node common.Address) (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.BanNode(&_EsportOracleRequester.TransactOpts, node)
}

// BanNode is a paid mutator transaction binding the contract method 0x3faba59e.
//
// Solidity: function banNode(address node) returns()
func (_EsportOracleRequester *EsportOracleRequesterTransactorSession) BanNode(node common.Address) (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.BanNode(&_EsportOracleRequester.TransactOpts, node)
}

// CallMatchOracle is a paid mutator transaction binding the contract method 0x89770753.
//
// Solidity: function callMatchOracle(uint256 matchId, (uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256) matchData) returns()
func (_EsportOracleRequester *EsportOracleRequesterTransactor) CallMatchOracle(opts *bind.TransactOpts, matchId *big.Int, matchData EsportOracleTypesMatch) (*types.Transaction, error) {
	return _EsportOracleRequester.contract.Transact(opts, "callMatchOracle", matchId, matchData)
}

// CallMatchOracle is a paid mutator transaction binding the contract method 0x89770753.
//
// Solidity: function callMatchOracle(uint256 matchId, (uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256) matchData) returns()
func (_EsportOracleRequester *EsportOracleRequesterSession) CallMatchOracle(matchId *big.Int, matchData EsportOracleTypesMatch) (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.CallMatchOracle(&_EsportOracleRequester.TransactOpts, matchId, matchData)
}

// CallMatchOracle is a paid mutator transaction binding the contract method 0x89770753.
//
// Solidity: function callMatchOracle(uint256 matchId, (uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256) matchData) returns()
func (_EsportOracleRequester *EsportOracleRequesterTransactorSession) CallMatchOracle(matchId *big.Int, matchData EsportOracleTypesMatch) (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.CallMatchOracle(&_EsportOracleRequester.TransactOpts, matchId, matchData)
}

// HandleNewMatches is a paid mutator transaction binding the contract method 0x3b23e7d7.
//
// Solidity: function handleNewMatches((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256)[] newMatch) returns()
func (_EsportOracleRequester *EsportOracleRequesterTransactor) HandleNewMatches(opts *bind.TransactOpts, newMatch []EsportOracleTypesMatch) (*types.Transaction, error) {
	return _EsportOracleRequester.contract.Transact(opts, "handleNewMatches", newMatch)
}

// HandleNewMatches is a paid mutator transaction binding the contract method 0x3b23e7d7.
//
// Solidity: function handleNewMatches((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256)[] newMatch) returns()
func (_EsportOracleRequester *EsportOracleRequesterSession) HandleNewMatches(newMatch []EsportOracleTypesMatch) (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.HandleNewMatches(&_EsportOracleRequester.TransactOpts, newMatch)
}

// HandleNewMatches is a paid mutator transaction binding the contract method 0x3b23e7d7.
//
// Solidity: function handleNewMatches((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256)[] newMatch) returns()
func (_EsportOracleRequester *EsportOracleRequesterTransactorSession) HandleNewMatches(newMatch []EsportOracleTypesMatch) (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.HandleNewMatches(&_EsportOracleRequester.TransactOpts, newMatch)
}

// MarkRequestsFulfilled is a paid mutator transaction binding the contract method 0xf8112862.
//
// Solidity: function markRequestsFulfilled(uint256 matchId) returns()
func (_EsportOracleRequester *EsportOracleRequesterTransactor) MarkRequestsFulfilled(opts *bind.TransactOpts, matchId *big.Int) (*types.Transaction, error) {
	return _EsportOracleRequester.contract.Transact(opts, "markRequestsFulfilled", matchId)
}

// MarkRequestsFulfilled is a paid mutator transaction binding the contract method 0xf8112862.
//
// Solidity: function markRequestsFulfilled(uint256 matchId) returns()
func (_EsportOracleRequester *EsportOracleRequesterSession) MarkRequestsFulfilled(matchId *big.Int) (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.MarkRequestsFulfilled(&_EsportOracleRequester.TransactOpts, matchId)
}

// MarkRequestsFulfilled is a paid mutator transaction binding the contract method 0xf8112862.
//
// Solidity: function markRequestsFulfilled(uint256 matchId) returns()
func (_EsportOracleRequester *EsportOracleRequesterTransactorSession) MarkRequestsFulfilled(matchId *big.Int) (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.MarkRequestsFulfilled(&_EsportOracleRequester.TransactOpts, matchId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_EsportOracleRequester *EsportOracleRequesterTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EsportOracleRequester.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_EsportOracleRequester *EsportOracleRequesterSession) Pause() (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.Pause(&_EsportOracleRequester.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_EsportOracleRequester *EsportOracleRequesterTransactorSession) Pause() (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.Pause(&_EsportOracleRequester.TransactOpts)
}

// RehabilitateNode is a paid mutator transaction binding the contract method 0x148b56d6.
//
// Solidity: function rehabilitateNode(address node) returns()
func (_EsportOracleRequester *EsportOracleRequesterTransactor) RehabilitateNode(opts *bind.TransactOpts, node common.Address) (*types.Transaction, error) {
	return _EsportOracleRequester.contract.Transact(opts, "rehabilitateNode", node)
}

// RehabilitateNode is a paid mutator transaction binding the contract method 0x148b56d6.
//
// Solidity: function rehabilitateNode(address node) returns()
func (_EsportOracleRequester *EsportOracleRequesterSession) RehabilitateNode(node common.Address) (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.RehabilitateNode(&_EsportOracleRequester.TransactOpts, node)
}

// RehabilitateNode is a paid mutator transaction binding the contract method 0x148b56d6.
//
// Solidity: function rehabilitateNode(address node) returns()
func (_EsportOracleRequester *EsportOracleRequesterTransactorSession) RehabilitateNode(node common.Address) (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.RehabilitateNode(&_EsportOracleRequester.TransactOpts, node)
}

// RequestMatch is a paid mutator transaction binding the contract method 0xfcaf2c40.
//
// Solidity: function requestMatch(uint256 matchId) payable returns(uint256)
func (_EsportOracleRequester *EsportOracleRequesterTransactor) RequestMatch(opts *bind.TransactOpts, matchId *big.Int) (*types.Transaction, error) {
	return _EsportOracleRequester.contract.Transact(opts, "requestMatch", matchId)
}

// RequestMatch is a paid mutator transaction binding the contract method 0xfcaf2c40.
//
// Solidity: function requestMatch(uint256 matchId) payable returns(uint256)
func (_EsportOracleRequester *EsportOracleRequesterSession) RequestMatch(matchId *big.Int) (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.RequestMatch(&_EsportOracleRequester.TransactOpts, matchId)
}

// RequestMatch is a paid mutator transaction binding the contract method 0xfcaf2c40.
//
// Solidity: function requestMatch(uint256 matchId) payable returns(uint256)
func (_EsportOracleRequester *EsportOracleRequesterTransactorSession) RequestMatch(matchId *big.Int) (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.RequestMatch(&_EsportOracleRequester.TransactOpts, matchId)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_EsportOracleRequester *EsportOracleRequesterTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EsportOracleRequester.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_EsportOracleRequester *EsportOracleRequesterSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.SetOwner(&_EsportOracleRequester.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_EsportOracleRequester *EsportOracleRequesterTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.SetOwner(&_EsportOracleRequester.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_EsportOracleRequester *EsportOracleRequesterTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EsportOracleRequester.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_EsportOracleRequester *EsportOracleRequesterSession) Unpause() (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.Unpause(&_EsportOracleRequester.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_EsportOracleRequester *EsportOracleRequesterTransactorSession) Unpause() (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.Unpause(&_EsportOracleRequester.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_EsportOracleRequester *EsportOracleRequesterTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EsportOracleRequester.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_EsportOracleRequester *EsportOracleRequesterSession) WithdrawStake() (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.WithdrawStake(&_EsportOracleRequester.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_EsportOracleRequester *EsportOracleRequesterTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _EsportOracleRequester.Contract.WithdrawStake(&_EsportOracleRequester.TransactOpts)
}

// EsportOracleRequesterMatchRequestedIterator is returned from FilterMatchRequested and is used to iterate over the raw logs and unpacked data for MatchRequested events raised by the EsportOracleRequester contract.
type EsportOracleRequesterMatchRequestedIterator struct {
	Event *EsportOracleRequesterMatchRequested // Event containing the contract specifics and raw log

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
func (it *EsportOracleRequesterMatchRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsportOracleRequesterMatchRequested)
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
		it.Event = new(EsportOracleRequesterMatchRequested)
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
func (it *EsportOracleRequesterMatchRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsportOracleRequesterMatchRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsportOracleRequesterMatchRequested represents a MatchRequested event raised by the EsportOracleRequester contract.
type EsportOracleRequesterMatchRequested struct {
	RequestId *big.Int
	MatchId   *big.Int
	Requester common.Address
	Fee       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMatchRequested is a free log retrieval operation binding the contract event 0x5a32146e161b34ba01d8cc6d09c6f37a13e52667cf666ac798e61b32876f1df4.
//
// Solidity: event MatchRequested(uint256 indexed requestId, uint256 indexed matchId, address indexed requester, uint256 fee)
func (_EsportOracleRequester *EsportOracleRequesterFilterer) FilterMatchRequested(opts *bind.FilterOpts, requestId []*big.Int, matchId []*big.Int, requester []common.Address) (*EsportOracleRequesterMatchRequestedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var matchIdRule []interface{}
	for _, matchIdItem := range matchId {
		matchIdRule = append(matchIdRule, matchIdItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _EsportOracleRequester.contract.FilterLogs(opts, "MatchRequested", requestIdRule, matchIdRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &EsportOracleRequesterMatchRequestedIterator{contract: _EsportOracleRequester.contract, event: "MatchRequested", logs: logs, sub: sub}, nil
}

// WatchMatchRequested is a free log subscription operation binding the contract event 0x5a32146e161b34ba01d8cc6d09c6f37a13e52667cf666ac798e61b32876f1df4.
//
// Solidity: event MatchRequested(uint256 indexed requestId, uint256 indexed matchId, address indexed requester, uint256 fee)
func (_EsportOracleRequester *EsportOracleRequesterFilterer) WatchMatchRequested(opts *bind.WatchOpts, sink chan<- *EsportOracleRequesterMatchRequested, requestId []*big.Int, matchId []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var matchIdRule []interface{}
	for _, matchIdItem := range matchId {
		matchIdRule = append(matchIdRule, matchIdItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _EsportOracleRequester.contract.WatchLogs(opts, "MatchRequested", requestIdRule, matchIdRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsportOracleRequesterMatchRequested)
				if err := _EsportOracleRequester.contract.UnpackLog(event, "MatchRequested", log); err != nil {
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

// ParseMatchRequested is a log parse operation binding the contract event 0x5a32146e161b34ba01d8cc6d09c6f37a13e52667cf666ac798e61b32876f1df4.
//
// Solidity: event MatchRequested(uint256 indexed requestId, uint256 indexed matchId, address indexed requester, uint256 fee)
func (_EsportOracleRequester *EsportOracleRequesterFilterer) ParseMatchRequested(log types.Log) (*EsportOracleRequesterMatchRequested, error) {
	event := new(EsportOracleRequesterMatchRequested)
	if err := _EsportOracleRequester.contract.UnpackLog(event, "MatchRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EsportOracleRequesterNodeBannedIterator is returned from FilterNodeBanned and is used to iterate over the raw logs and unpacked data for NodeBanned events raised by the EsportOracleRequester contract.
type EsportOracleRequesterNodeBannedIterator struct {
	Event *EsportOracleRequesterNodeBanned // Event containing the contract specifics and raw log

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
func (it *EsportOracleRequesterNodeBannedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsportOracleRequesterNodeBanned)
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
		it.Event = new(EsportOracleRequesterNodeBanned)
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
func (it *EsportOracleRequesterNodeBannedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsportOracleRequesterNodeBannedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsportOracleRequesterNodeBanned represents a NodeBanned event raised by the EsportOracleRequester contract.
type EsportOracleRequesterNodeBanned struct {
	Node common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNodeBanned is a free log retrieval operation binding the contract event 0xf1a151967feb9c87e5aa6cd9cdee148015dab9c012ef309968cb91c279e094b4.
//
// Solidity: event NodeBanned(address indexed node)
func (_EsportOracleRequester *EsportOracleRequesterFilterer) FilterNodeBanned(opts *bind.FilterOpts, node []common.Address) (*EsportOracleRequesterNodeBannedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _EsportOracleRequester.contract.FilterLogs(opts, "NodeBanned", nodeRule)
	if err != nil {
		return nil, err
	}
	return &EsportOracleRequesterNodeBannedIterator{contract: _EsportOracleRequester.contract, event: "NodeBanned", logs: logs, sub: sub}, nil
}

// WatchNodeBanned is a free log subscription operation binding the contract event 0xf1a151967feb9c87e5aa6cd9cdee148015dab9c012ef309968cb91c279e094b4.
//
// Solidity: event NodeBanned(address indexed node)
func (_EsportOracleRequester *EsportOracleRequesterFilterer) WatchNodeBanned(opts *bind.WatchOpts, sink chan<- *EsportOracleRequesterNodeBanned, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _EsportOracleRequester.contract.WatchLogs(opts, "NodeBanned", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsportOracleRequesterNodeBanned)
				if err := _EsportOracleRequester.contract.UnpackLog(event, "NodeBanned", log); err != nil {
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
func (_EsportOracleRequester *EsportOracleRequesterFilterer) ParseNodeBanned(log types.Log) (*EsportOracleRequesterNodeBanned, error) {
	event := new(EsportOracleRequesterNodeBanned)
	if err := _EsportOracleRequester.contract.UnpackLog(event, "NodeBanned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EsportOracleRequesterNodePunishedIterator is returned from FilterNodePunished and is used to iterate over the raw logs and unpacked data for NodePunished events raised by the EsportOracleRequester contract.
type EsportOracleRequesterNodePunishedIterator struct {
	Event *EsportOracleRequesterNodePunished // Event containing the contract specifics and raw log

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
func (it *EsportOracleRequesterNodePunishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsportOracleRequesterNodePunished)
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
		it.Event = new(EsportOracleRequesterNodePunished)
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
func (it *EsportOracleRequesterNodePunishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsportOracleRequesterNodePunishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsportOracleRequesterNodePunished represents a NodePunished event raised by the EsportOracleRequester contract.
type EsportOracleRequesterNodePunished struct {
	Node            common.Address
	Amount          *big.Int
	ViolationsCount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNodePunished is a free log retrieval operation binding the contract event 0xbb5e0de9b41d911fe88799197a82d842eff4fdea524cdfc0becff3857f0ee3ac.
//
// Solidity: event NodePunished(address indexed node, uint256 amount, uint256 violationsCount)
func (_EsportOracleRequester *EsportOracleRequesterFilterer) FilterNodePunished(opts *bind.FilterOpts, node []common.Address) (*EsportOracleRequesterNodePunishedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _EsportOracleRequester.contract.FilterLogs(opts, "NodePunished", nodeRule)
	if err != nil {
		return nil, err
	}
	return &EsportOracleRequesterNodePunishedIterator{contract: _EsportOracleRequester.contract, event: "NodePunished", logs: logs, sub: sub}, nil
}

// WatchNodePunished is a free log subscription operation binding the contract event 0xbb5e0de9b41d911fe88799197a82d842eff4fdea524cdfc0becff3857f0ee3ac.
//
// Solidity: event NodePunished(address indexed node, uint256 amount, uint256 violationsCount)
func (_EsportOracleRequester *EsportOracleRequesterFilterer) WatchNodePunished(opts *bind.WatchOpts, sink chan<- *EsportOracleRequesterNodePunished, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _EsportOracleRequester.contract.WatchLogs(opts, "NodePunished", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsportOracleRequesterNodePunished)
				if err := _EsportOracleRequester.contract.UnpackLog(event, "NodePunished", log); err != nil {
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
func (_EsportOracleRequester *EsportOracleRequesterFilterer) ParseNodePunished(log types.Log) (*EsportOracleRequesterNodePunished, error) {
	event := new(EsportOracleRequesterNodePunished)
	if err := _EsportOracleRequester.contract.UnpackLog(event, "NodePunished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EsportOracleRequesterPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EsportOracleRequester contract.
type EsportOracleRequesterPausedIterator struct {
	Event *EsportOracleRequesterPaused // Event containing the contract specifics and raw log

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
func (it *EsportOracleRequesterPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsportOracleRequesterPaused)
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
		it.Event = new(EsportOracleRequesterPaused)
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
func (it *EsportOracleRequesterPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsportOracleRequesterPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsportOracleRequesterPaused represents a Paused event raised by the EsportOracleRequester contract.
type EsportOracleRequesterPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EsportOracleRequester *EsportOracleRequesterFilterer) FilterPaused(opts *bind.FilterOpts) (*EsportOracleRequesterPausedIterator, error) {

	logs, sub, err := _EsportOracleRequester.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EsportOracleRequesterPausedIterator{contract: _EsportOracleRequester.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EsportOracleRequester *EsportOracleRequesterFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EsportOracleRequesterPaused) (event.Subscription, error) {

	logs, sub, err := _EsportOracleRequester.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsportOracleRequesterPaused)
				if err := _EsportOracleRequester.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_EsportOracleRequester *EsportOracleRequesterFilterer) ParsePaused(log types.Log) (*EsportOracleRequesterPaused, error) {
	event := new(EsportOracleRequesterPaused)
	if err := _EsportOracleRequester.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EsportOracleRequesterUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EsportOracleRequester contract.
type EsportOracleRequesterUnpausedIterator struct {
	Event *EsportOracleRequesterUnpaused // Event containing the contract specifics and raw log

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
func (it *EsportOracleRequesterUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsportOracleRequesterUnpaused)
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
		it.Event = new(EsportOracleRequesterUnpaused)
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
func (it *EsportOracleRequesterUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsportOracleRequesterUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsportOracleRequesterUnpaused represents a Unpaused event raised by the EsportOracleRequester contract.
type EsportOracleRequesterUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EsportOracleRequester *EsportOracleRequesterFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EsportOracleRequesterUnpausedIterator, error) {

	logs, sub, err := _EsportOracleRequester.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EsportOracleRequesterUnpausedIterator{contract: _EsportOracleRequester.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EsportOracleRequester *EsportOracleRequesterFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EsportOracleRequesterUnpaused) (event.Subscription, error) {

	logs, sub, err := _EsportOracleRequester.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsportOracleRequesterUnpaused)
				if err := _EsportOracleRequester.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_EsportOracleRequester *EsportOracleRequesterFilterer) ParseUnpaused(log types.Log) (*EsportOracleRequesterUnpaused, error) {
	event := new(EsportOracleRequesterUnpaused)
	if err := _EsportOracleRequester.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EsportOracleRequesterNewNodeAddedIterator is returned from FilterNewNodeAdded and is used to iterate over the raw logs and unpacked data for NewNodeAdded events raised by the EsportOracleRequester contract.
type EsportOracleRequesterNewNodeAddedIterator struct {
	Event *EsportOracleRequesterNewNodeAdded // Event containing the contract specifics and raw log

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
func (it *EsportOracleRequesterNewNodeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsportOracleRequesterNewNodeAdded)
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
		it.Event = new(EsportOracleRequesterNewNodeAdded)
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
func (it *EsportOracleRequesterNewNodeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsportOracleRequesterNewNodeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsportOracleRequesterNewNodeAdded represents a NewNodeAdded event raised by the EsportOracleRequester contract.
type EsportOracleRequesterNewNodeAdded struct {
	AddressAdded common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewNodeAdded is a free log retrieval operation binding the contract event 0xc473661237ac48848af775cb3c9f5bfdf90ed59dbe739ebac0186bc6e61d1984.
//
// Solidity: event newNodeAdded(address indexed addressAdded)
func (_EsportOracleRequester *EsportOracleRequesterFilterer) FilterNewNodeAdded(opts *bind.FilterOpts, addressAdded []common.Address) (*EsportOracleRequesterNewNodeAddedIterator, error) {

	var addressAddedRule []interface{}
	for _, addressAddedItem := range addressAdded {
		addressAddedRule = append(addressAddedRule, addressAddedItem)
	}

	logs, sub, err := _EsportOracleRequester.contract.FilterLogs(opts, "newNodeAdded", addressAddedRule)
	if err != nil {
		return nil, err
	}
	return &EsportOracleRequesterNewNodeAddedIterator{contract: _EsportOracleRequester.contract, event: "newNodeAdded", logs: logs, sub: sub}, nil
}

// WatchNewNodeAdded is a free log subscription operation binding the contract event 0xc473661237ac48848af775cb3c9f5bfdf90ed59dbe739ebac0186bc6e61d1984.
//
// Solidity: event newNodeAdded(address indexed addressAdded)
func (_EsportOracleRequester *EsportOracleRequesterFilterer) WatchNewNodeAdded(opts *bind.WatchOpts, sink chan<- *EsportOracleRequesterNewNodeAdded, addressAdded []common.Address) (event.Subscription, error) {

	var addressAddedRule []interface{}
	for _, addressAddedItem := range addressAdded {
		addressAddedRule = append(addressAddedRule, addressAddedItem)
	}

	logs, sub, err := _EsportOracleRequester.contract.WatchLogs(opts, "newNodeAdded", addressAddedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsportOracleRequesterNewNodeAdded)
				if err := _EsportOracleRequester.contract.UnpackLog(event, "newNodeAdded", log); err != nil {
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
func (_EsportOracleRequester *EsportOracleRequesterFilterer) ParseNewNodeAdded(log types.Log) (*EsportOracleRequesterNewNodeAdded, error) {
	event := new(EsportOracleRequesterNewNodeAdded)
	if err := _EsportOracleRequester.contract.UnpackLog(event, "newNodeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EsportOracleRequesterStakingSuccessIterator is returned from FilterStakingSuccess and is used to iterate over the raw logs and unpacked data for StakingSuccess events raised by the EsportOracleRequester contract.
type EsportOracleRequesterStakingSuccessIterator struct {
	Event *EsportOracleRequesterStakingSuccess // Event containing the contract specifics and raw log

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
func (it *EsportOracleRequesterStakingSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsportOracleRequesterStakingSuccess)
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
		it.Event = new(EsportOracleRequesterStakingSuccess)
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
func (it *EsportOracleRequesterStakingSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsportOracleRequesterStakingSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsportOracleRequesterStakingSuccess represents a StakingSuccess event raised by the EsportOracleRequester contract.
type EsportOracleRequesterStakingSuccess struct {
	AddressAdded common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakingSuccess is a free log retrieval operation binding the contract event 0xc60400ca3fd7ce7a123251117104ceac79276329b4f7e0413700342999d3cc1d.
//
// Solidity: event stakingSuccess(address indexed addressAdded, uint256 amount)
func (_EsportOracleRequester *EsportOracleRequesterFilterer) FilterStakingSuccess(opts *bind.FilterOpts, addressAdded []common.Address) (*EsportOracleRequesterStakingSuccessIterator, error) {

	var addressAddedRule []interface{}
	for _, addressAddedItem := range addressAdded {
		addressAddedRule = append(addressAddedRule, addressAddedItem)
	}

	logs, sub, err := _EsportOracleRequester.contract.FilterLogs(opts, "stakingSuccess", addressAddedRule)
	if err != nil {
		return nil, err
	}
	return &EsportOracleRequesterStakingSuccessIterator{contract: _EsportOracleRequester.contract, event: "stakingSuccess", logs: logs, sub: sub}, nil
}

// WatchStakingSuccess is a free log subscription operation binding the contract event 0xc60400ca3fd7ce7a123251117104ceac79276329b4f7e0413700342999d3cc1d.
//
// Solidity: event stakingSuccess(address indexed addressAdded, uint256 amount)
func (_EsportOracleRequester *EsportOracleRequesterFilterer) WatchStakingSuccess(opts *bind.WatchOpts, sink chan<- *EsportOracleRequesterStakingSuccess, addressAdded []common.Address) (event.Subscription, error) {

	var addressAddedRule []interface{}
	for _, addressAddedItem := range addressAdded {
		addressAddedRule = append(addressAddedRule, addressAddedItem)
	}

	logs, sub, err := _EsportOracleRequester.contract.WatchLogs(opts, "stakingSuccess", addressAddedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsportOracleRequesterStakingSuccess)
				if err := _EsportOracleRequester.contract.UnpackLog(event, "stakingSuccess", log); err != nil {
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
func (_EsportOracleRequester *EsportOracleRequesterFilterer) ParseStakingSuccess(log types.Log) (*EsportOracleRequesterStakingSuccess, error) {
	event := new(EsportOracleRequesterStakingSuccess)
	if err := _EsportOracleRequester.contract.UnpackLog(event, "stakingSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
