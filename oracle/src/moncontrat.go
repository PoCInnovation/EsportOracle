// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Esp

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

// EspMetaData contains all meta data concerning the Esp contract.
var EspMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"getPerson\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setPerson\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506104278061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c80632a0d99af14610038578063552d2d5c1461004d575b5f5ffd5b61004b61004636600461018a565b610077565b005b61006061005b366004610256565b6100a4565b60405161006e929190610276565b60405180910390f35b6001600160a01b0383165f90815260208190526040902082815560010161009e8282610336565b50505050565b6001600160a01b0381165f908152602081905260408120805460019091018054606092919081906100d4906102b2565b80601f0160208091040260200160405190810160405280929190818152602001828054610100906102b2565b801561014b5780601f106101225761010080835404028352916020019161014b565b820191905f5260205f20905b81548152906001019060200180831161012e57829003601f168201915b5050505050905091509150915091565b80356001600160a01b0381168114610171575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f6060848603121561019c575f5ffd5b6101a58461015b565b925060208401359150604084013567ffffffffffffffff8111156101c7575f5ffd5b8401601f810186136101d7575f5ffd5b803567ffffffffffffffff8111156101f1576101f1610176565b604051601f8201601f19908116603f0116810167ffffffffffffffff8111828210171561022057610220610176565b604052818152828201602001881015610237575f5ffd5b816020840160208301375f602083830101528093505050509250925092565b5f60208284031215610266575f5ffd5b61026f8261015b565b9392505050565b828152604060208201525f82518060408401528060208501606085015e5f606082850101526060601f19601f8301168401019150509392505050565b600181811c908216806102c657607f821691505b6020821081036102e457634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561033157805f5260205f20601f840160051c8101602085101561030f5750805b601f840160051c820191505b8181101561032e575f815560010161031b565b50505b505050565b815167ffffffffffffffff81111561035057610350610176565b6103648161035e84546102b2565b846102ea565b6020601f821160018114610396575f831561037f5750848201515b5f19600385901b1c1916600184901b17845561032e565b5f84815260208120601f198516915b828110156103c557878501518255602094850194600190920191016103a5565b50848210156103e257868401515f19600387901b60f8161c191681555b50505050600190811b0190555056fea2646970667358221220a3d84eb5e13f06f6bd25bbf2d2351866b34a000d846fd61571d5531e8a26796c64736f6c634300081e0033",
}

// EspABI is the input ABI used to generate the binding from.
// Deprecated: Use EspMetaData.ABI instead.
var EspABI = EspMetaData.ABI

// EspBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EspMetaData.Bin instead.
var EspBin = EspMetaData.Bin

// DeployEsp deploys a new Ethereum contract, binding an instance of Esp to it.
func DeployEsp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Esp, error) {
	parsed, err := EspMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EspBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Esp{EspCaller: EspCaller{contract: contract}, EspTransactor: EspTransactor{contract: contract}, EspFilterer: EspFilterer{contract: contract}}, nil
}

// Esp is an auto generated Go binding around an Ethereum contract.
type Esp struct {
	EspCaller     // Read-only binding to the contract
	EspTransactor // Write-only binding to the contract
	EspFilterer   // Log filterer for contract events
}

// EspCaller is an auto generated read-only Go binding around an Ethereum contract.
type EspCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EspTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EspTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EspFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EspFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EspSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EspSession struct {
	Contract     *Esp              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EspCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EspCallerSession struct {
	Contract *EspCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EspTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EspTransactorSession struct {
	Contract     *EspTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EspRaw is an auto generated low-level Go binding around an Ethereum contract.
type EspRaw struct {
	Contract *Esp // Generic contract binding to access the raw methods on
}

// EspCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EspCallerRaw struct {
	Contract *EspCaller // Generic read-only contract binding to access the raw methods on
}

// EspTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EspTransactorRaw struct {
	Contract *EspTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEsp creates a new instance of Esp, bound to a specific deployed contract.
func NewEsp(address common.Address, backend bind.ContractBackend) (*Esp, error) {
	contract, err := bindEsp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Esp{EspCaller: EspCaller{contract: contract}, EspTransactor: EspTransactor{contract: contract}, EspFilterer: EspFilterer{contract: contract}}, nil
}

// NewEspCaller creates a new read-only instance of Esp, bound to a specific deployed contract.
func NewEspCaller(address common.Address, caller bind.ContractCaller) (*EspCaller, error) {
	contract, err := bindEsp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EspCaller{contract: contract}, nil
}

// NewEspTransactor creates a new write-only instance of Esp, bound to a specific deployed contract.
func NewEspTransactor(address common.Address, transactor bind.ContractTransactor) (*EspTransactor, error) {
	contract, err := bindEsp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EspTransactor{contract: contract}, nil
}

// NewEspFilterer creates a new log filterer instance of Esp, bound to a specific deployed contract.
func NewEspFilterer(address common.Address, filterer bind.ContractFilterer) (*EspFilterer, error) {
	contract, err := bindEsp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EspFilterer{contract: contract}, nil
}

// bindEsp binds a generic wrapper to an already deployed contract.
func bindEsp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EspMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Esp *EspRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Esp.Contract.EspCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Esp *EspRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Esp.Contract.EspTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Esp *EspRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Esp.Contract.EspTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Esp *EspCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Esp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Esp *EspTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Esp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Esp *EspTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Esp.Contract.contract.Transact(opts, method, params...)
}

// GetPerson is a free data retrieval call binding the contract method 0x552d2d5c.
//
// Solidity: function getPerson(address _to) view returns(uint256, string)
func (_Esp *EspCaller) GetPerson(opts *bind.CallOpts, _to common.Address) (*big.Int, string, error) {
	var out []interface{}
	err := _Esp.contract.Call(opts, &out, "getPerson", _to)

	if err != nil {
		return *new(*big.Int), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetPerson is a free data retrieval call binding the contract method 0x552d2d5c.
//
// Solidity: function getPerson(address _to) view returns(uint256, string)
func (_Esp *EspSession) GetPerson(_to common.Address) (*big.Int, string, error) {
	return _Esp.Contract.GetPerson(&_Esp.CallOpts, _to)
}

// GetPerson is a free data retrieval call binding the contract method 0x552d2d5c.
//
// Solidity: function getPerson(address _to) view returns(uint256, string)
func (_Esp *EspCallerSession) GetPerson(_to common.Address) (*big.Int, string, error) {
	return _Esp.Contract.GetPerson(&_Esp.CallOpts, _to)
}

// SetPerson is a paid mutator transaction binding the contract method 0x2a0d99af.
//
// Solidity: function setPerson(address _to, uint256 _id, string _name) returns()
func (_Esp *EspTransactor) SetPerson(opts *bind.TransactOpts, _to common.Address, _id *big.Int, _name string) (*types.Transaction, error) {
	return _Esp.contract.Transact(opts, "setPerson", _to, _id, _name)
}

// SetPerson is a paid mutator transaction binding the contract method 0x2a0d99af.
//
// Solidity: function setPerson(address _to, uint256 _id, string _name) returns()
func (_Esp *EspSession) SetPerson(_to common.Address, _id *big.Int, _name string) (*types.Transaction, error) {
	return _Esp.Contract.SetPerson(&_Esp.TransactOpts, _to, _id, _name)
}

// SetPerson is a paid mutator transaction binding the contract method 0x2a0d99af.
//
// Solidity: function setPerson(address _to, uint256 _id, string _name) returns()
func (_Esp *EspTransactorSession) SetPerson(_to common.Address, _id *big.Int, _name string) (*types.Transaction, error) {
	return _Esp.Contract.SetPerson(&_Esp.TransactOpts, _to, _id, _name)
}
