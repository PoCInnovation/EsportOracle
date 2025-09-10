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

// EsportOracleClientMetaData contains all meta data concerning the EsportOracleClient contract.
var EsportOracleClientMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"_owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"callMatchReceived\",\"inputs\":[{\"name\":\"_match\",\"type\":\"tuple\",\"internalType\":\"structEsportOracleTypes.Match\",\"components\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_opponents\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Opponents[]\",\"components\":[{\"name\":\"_acronym\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"_game\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Games[]\",\"components\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_finished\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_winnerId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_result\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Result[]\",\"components\":[{\"name\":\"_score\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_teamId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_winnerId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_beginAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractInterfaceOracle\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveMatch\",\"inputs\":[{\"name\":\"matchId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"showMatch\",\"inputs\":[{\"name\":\"matchId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structEsportOracleTypes.Match\",\"components\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_opponents\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Opponents[]\",\"components\":[{\"name\":\"_acronym\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"_game\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Games[]\",\"components\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_finished\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_winnerId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_result\",\"type\":\"tuple[]\",\"internalType\":\"structEsportOracleTypes.Result[]\",\"components\":[{\"name\":\"_score\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_teamId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_winnerId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_beginAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"showPendingRequestedMatches\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"}]",
}

// EsportOracleClientABI is the input ABI used to generate the binding from.
// Deprecated: Use EsportOracleClientMetaData.ABI instead.
var EsportOracleClientABI = EsportOracleClientMetaData.ABI

// EsportOracleClient is an auto generated Go binding around an Ethereum contract.
type EsportOracleClient struct {
	EsportOracleClientCaller     // Read-only binding to the contract
	EsportOracleClientTransactor // Write-only binding to the contract
	EsportOracleClientFilterer   // Log filterer for contract events
}

// EsportOracleClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type EsportOracleClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EsportOracleClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EsportOracleClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EsportOracleClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EsportOracleClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EsportOracleClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EsportOracleClientSession struct {
	Contract     *EsportOracleClient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EsportOracleClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EsportOracleClientCallerSession struct {
	Contract *EsportOracleClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// EsportOracleClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EsportOracleClientTransactorSession struct {
	Contract     *EsportOracleClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// EsportOracleClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type EsportOracleClientRaw struct {
	Contract *EsportOracleClient // Generic contract binding to access the raw methods on
}

// EsportOracleClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EsportOracleClientCallerRaw struct {
	Contract *EsportOracleClientCaller // Generic read-only contract binding to access the raw methods on
}

// EsportOracleClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EsportOracleClientTransactorRaw struct {
	Contract *EsportOracleClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEsportOracleClient creates a new instance of EsportOracleClient, bound to a specific deployed contract.
func NewEsportOracleClient(address common.Address, backend bind.ContractBackend) (*EsportOracleClient, error) {
	contract, err := bindEsportOracleClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EsportOracleClient{EsportOracleClientCaller: EsportOracleClientCaller{contract: contract}, EsportOracleClientTransactor: EsportOracleClientTransactor{contract: contract}, EsportOracleClientFilterer: EsportOracleClientFilterer{contract: contract}}, nil
}

// NewEsportOracleClientCaller creates a new read-only instance of EsportOracleClient, bound to a specific deployed contract.
func NewEsportOracleClientCaller(address common.Address, caller bind.ContractCaller) (*EsportOracleClientCaller, error) {
	contract, err := bindEsportOracleClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EsportOracleClientCaller{contract: contract}, nil
}

// NewEsportOracleClientTransactor creates a new write-only instance of EsportOracleClient, bound to a specific deployed contract.
func NewEsportOracleClientTransactor(address common.Address, transactor bind.ContractTransactor) (*EsportOracleClientTransactor, error) {
	contract, err := bindEsportOracleClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EsportOracleClientTransactor{contract: contract}, nil
}

// NewEsportOracleClientFilterer creates a new log filterer instance of EsportOracleClient, bound to a specific deployed contract.
func NewEsportOracleClientFilterer(address common.Address, filterer bind.ContractFilterer) (*EsportOracleClientFilterer, error) {
	contract, err := bindEsportOracleClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EsportOracleClientFilterer{contract: contract}, nil
}

// bindEsportOracleClient binds a generic wrapper to an already deployed contract.
func bindEsportOracleClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EsportOracleClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EsportOracleClient *EsportOracleClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EsportOracleClient.Contract.EsportOracleClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EsportOracleClient *EsportOracleClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EsportOracleClient.Contract.EsportOracleClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EsportOracleClient *EsportOracleClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EsportOracleClient.Contract.EsportOracleClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EsportOracleClient *EsportOracleClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EsportOracleClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EsportOracleClient *EsportOracleClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EsportOracleClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EsportOracleClient *EsportOracleClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EsportOracleClient.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_EsportOracleClient *EsportOracleClientCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EsportOracleClient.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_EsportOracleClient *EsportOracleClientSession) Owner() (common.Address, error) {
	return _EsportOracleClient.Contract.Owner(&_EsportOracleClient.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_EsportOracleClient *EsportOracleClientCallerSession) Owner() (common.Address, error) {
	return _EsportOracleClient.Contract.Owner(&_EsportOracleClient.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_EsportOracleClient *EsportOracleClientCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EsportOracleClient.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_EsportOracleClient *EsportOracleClientSession) Oracle() (common.Address, error) {
	return _EsportOracleClient.Contract.Oracle(&_EsportOracleClient.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_EsportOracleClient *EsportOracleClientCallerSession) Oracle() (common.Address, error) {
	return _EsportOracleClient.Contract.Oracle(&_EsportOracleClient.CallOpts)
}

// RequestId is a free data retrieval call binding the contract method 0x006d6cae.
//
// Solidity: function requestId() view returns(uint256)
func (_EsportOracleClient *EsportOracleClientCaller) RequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EsportOracleClient.contract.Call(opts, &out, "requestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestId is a free data retrieval call binding the contract method 0x006d6cae.
//
// Solidity: function requestId() view returns(uint256)
func (_EsportOracleClient *EsportOracleClientSession) RequestId() (*big.Int, error) {
	return _EsportOracleClient.Contract.RequestId(&_EsportOracleClient.CallOpts)
}

// RequestId is a free data retrieval call binding the contract method 0x006d6cae.
//
// Solidity: function requestId() view returns(uint256)
func (_EsportOracleClient *EsportOracleClientCallerSession) RequestId() (*big.Int, error) {
	return _EsportOracleClient.Contract.RequestId(&_EsportOracleClient.CallOpts)
}

// ShowMatch is a free data retrieval call binding the contract method 0x8630bc76.
//
// Solidity: function showMatch(uint256 matchId) view returns((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256))
func (_EsportOracleClient *EsportOracleClientCaller) ShowMatch(opts *bind.CallOpts, matchId *big.Int) (EsportOracleTypesMatch, error) {
	var out []interface{}
	err := _EsportOracleClient.contract.Call(opts, &out, "showMatch", matchId)

	if err != nil {
		return *new(EsportOracleTypesMatch), err
	}

	out0 := *abi.ConvertType(out[0], new(EsportOracleTypesMatch)).(*EsportOracleTypesMatch)

	return out0, err

}

// ShowMatch is a free data retrieval call binding the contract method 0x8630bc76.
//
// Solidity: function showMatch(uint256 matchId) view returns((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256))
func (_EsportOracleClient *EsportOracleClientSession) ShowMatch(matchId *big.Int) (EsportOracleTypesMatch, error) {
	return _EsportOracleClient.Contract.ShowMatch(&_EsportOracleClient.CallOpts, matchId)
}

// ShowMatch is a free data retrieval call binding the contract method 0x8630bc76.
//
// Solidity: function showMatch(uint256 matchId) view returns((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256))
func (_EsportOracleClient *EsportOracleClientCallerSession) ShowMatch(matchId *big.Int) (EsportOracleTypesMatch, error) {
	return _EsportOracleClient.Contract.ShowMatch(&_EsportOracleClient.CallOpts, matchId)
}

// ShowPendingRequestedMatches is a free data retrieval call binding the contract method 0x7fd8511d.
//
// Solidity: function showPendingRequestedMatches() view returns(uint256[])
func (_EsportOracleClient *EsportOracleClientCaller) ShowPendingRequestedMatches(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _EsportOracleClient.contract.Call(opts, &out, "showPendingRequestedMatches")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ShowPendingRequestedMatches is a free data retrieval call binding the contract method 0x7fd8511d.
//
// Solidity: function showPendingRequestedMatches() view returns(uint256[])
func (_EsportOracleClient *EsportOracleClientSession) ShowPendingRequestedMatches() ([]*big.Int, error) {
	return _EsportOracleClient.Contract.ShowPendingRequestedMatches(&_EsportOracleClient.CallOpts)
}

// ShowPendingRequestedMatches is a free data retrieval call binding the contract method 0x7fd8511d.
//
// Solidity: function showPendingRequestedMatches() view returns(uint256[])
func (_EsportOracleClient *EsportOracleClientCallerSession) ShowPendingRequestedMatches() ([]*big.Int, error) {
	return _EsportOracleClient.Contract.ShowPendingRequestedMatches(&_EsportOracleClient.CallOpts)
}

// CallMatchReceived is a paid mutator transaction binding the contract method 0x95288977.
//
// Solidity: function callMatchReceived((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256) _match) returns()
func (_EsportOracleClient *EsportOracleClientTransactor) CallMatchReceived(opts *bind.TransactOpts, _match EsportOracleTypesMatch) (*types.Transaction, error) {
	return _EsportOracleClient.contract.Transact(opts, "callMatchReceived", _match)
}

// CallMatchReceived is a paid mutator transaction binding the contract method 0x95288977.
//
// Solidity: function callMatchReceived((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256) _match) returns()
func (_EsportOracleClient *EsportOracleClientSession) CallMatchReceived(_match EsportOracleTypesMatch) (*types.Transaction, error) {
	return _EsportOracleClient.Contract.CallMatchReceived(&_EsportOracleClient.TransactOpts, _match)
}

// CallMatchReceived is a paid mutator transaction binding the contract method 0x95288977.
//
// Solidity: function callMatchReceived((uint256,(string,uint256,string)[],(uint256,bool,uint256)[],(uint8,uint256)[],uint256,uint256) _match) returns()
func (_EsportOracleClient *EsportOracleClientTransactorSession) CallMatchReceived(_match EsportOracleTypesMatch) (*types.Transaction, error) {
	return _EsportOracleClient.Contract.CallMatchReceived(&_EsportOracleClient.TransactOpts, _match)
}

// ReceiveMatch is a paid mutator transaction binding the contract method 0x9f6503e0.
//
// Solidity: function receiveMatch(uint256 matchId) payable returns(uint256)
func (_EsportOracleClient *EsportOracleClientTransactor) ReceiveMatch(opts *bind.TransactOpts, matchId *big.Int) (*types.Transaction, error) {
	return _EsportOracleClient.contract.Transact(opts, "receiveMatch", matchId)
}

// ReceiveMatch is a paid mutator transaction binding the contract method 0x9f6503e0.
//
// Solidity: function receiveMatch(uint256 matchId) payable returns(uint256)
func (_EsportOracleClient *EsportOracleClientSession) ReceiveMatch(matchId *big.Int) (*types.Transaction, error) {
	return _EsportOracleClient.Contract.ReceiveMatch(&_EsportOracleClient.TransactOpts, matchId)
}

// ReceiveMatch is a paid mutator transaction binding the contract method 0x9f6503e0.
//
// Solidity: function receiveMatch(uint256 matchId) payable returns(uint256)
func (_EsportOracleClient *EsportOracleClientTransactorSession) ReceiveMatch(matchId *big.Int) (*types.Transaction, error) {
	return _EsportOracleClient.Contract.ReceiveMatch(&_EsportOracleClient.TransactOpts, matchId)
}
