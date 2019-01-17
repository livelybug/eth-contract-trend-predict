// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package predict

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PredictABI is the input ABI used to generate the binding from.
const PredictABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"poolSum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"predict\",\"type\":\"uint256\"}],\"name\":\"resolve\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rangeMax\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"predict\",\"type\":\"uint256\"}],\"name\":\"bet\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minBet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"storeAdds\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"range\",\"type\":\"uint256\"},{\"name\":\"minimunt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// Predict is an auto generated Go binding around an Ethereum contract.
type Predict struct {
	PredictCaller     // Read-only binding to the contract
	PredictTransactor // Write-only binding to the contract
	PredictFilterer   // Log filterer for contract events
}

// PredictCaller is an auto generated read-only Go binding around an Ethereum contract.
type PredictCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredictTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PredictTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredictFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PredictFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredictSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PredictSession struct {
	Contract     *Predict          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PredictCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PredictCallerSession struct {
	Contract *PredictCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PredictTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PredictTransactorSession struct {
	Contract     *PredictTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PredictRaw is an auto generated low-level Go binding around an Ethereum contract.
type PredictRaw struct {
	Contract *Predict // Generic contract binding to access the raw methods on
}

// PredictCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PredictCallerRaw struct {
	Contract *PredictCaller // Generic read-only contract binding to access the raw methods on
}

// PredictTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PredictTransactorRaw struct {
	Contract *PredictTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPredict creates a new instance of Predict, bound to a specific deployed contract.
func NewPredict(address common.Address, backend bind.ContractBackend) (*Predict, error) {
	contract, err := bindPredict(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Predict{PredictCaller: PredictCaller{contract: contract}, PredictTransactor: PredictTransactor{contract: contract}, PredictFilterer: PredictFilterer{contract: contract}}, nil
}

// NewPredictCaller creates a new read-only instance of Predict, bound to a specific deployed contract.
func NewPredictCaller(address common.Address, caller bind.ContractCaller) (*PredictCaller, error) {
	contract, err := bindPredict(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PredictCaller{contract: contract}, nil
}

// NewPredictTransactor creates a new write-only instance of Predict, bound to a specific deployed contract.
func NewPredictTransactor(address common.Address, transactor bind.ContractTransactor) (*PredictTransactor, error) {
	contract, err := bindPredict(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PredictTransactor{contract: contract}, nil
}

// NewPredictFilterer creates a new log filterer instance of Predict, bound to a specific deployed contract.
func NewPredictFilterer(address common.Address, filterer bind.ContractFilterer) (*PredictFilterer, error) {
	contract, err := bindPredict(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PredictFilterer{contract: contract}, nil
}

// bindPredict binds a generic wrapper to an already deployed contract.
func bindPredict(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PredictABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Predict *PredictRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Predict.Contract.PredictCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Predict *PredictRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Predict.Contract.PredictTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Predict *PredictRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Predict.Contract.PredictTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Predict *PredictCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Predict.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Predict *PredictTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Predict.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Predict *PredictTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Predict.Contract.contract.Transact(opts, method, params...)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_Predict *PredictCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Predict.contract.Call(opts, out, "manager")
	return *ret0, err
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_Predict *PredictSession) Manager() (common.Address, error) {
	return _Predict.Contract.Manager(&_Predict.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_Predict *PredictCallerSession) Manager() (common.Address, error) {
	return _Predict.Contract.Manager(&_Predict.CallOpts)
}

// MinBet is a free data retrieval call binding the contract method 0x9619367d.
//
// Solidity: function minBet() constant returns(uint256)
func (_Predict *PredictCaller) MinBet(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Predict.contract.Call(opts, out, "minBet")
	return *ret0, err
}

// MinBet is a free data retrieval call binding the contract method 0x9619367d.
//
// Solidity: function minBet() constant returns(uint256)
func (_Predict *PredictSession) MinBet() (*big.Int, error) {
	return _Predict.Contract.MinBet(&_Predict.CallOpts)
}

// MinBet is a free data retrieval call binding the contract method 0x9619367d.
//
// Solidity: function minBet() constant returns(uint256)
func (_Predict *PredictCallerSession) MinBet() (*big.Int, error) {
	return _Predict.Contract.MinBet(&_Predict.CallOpts)
}

// PoolSum is a free data retrieval call binding the contract method 0x26e309d9.
//
// Solidity: function poolSum() constant returns(uint256)
func (_Predict *PredictCaller) PoolSum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Predict.contract.Call(opts, out, "poolSum")
	return *ret0, err
}

// PoolSum is a free data retrieval call binding the contract method 0x26e309d9.
//
// Solidity: function poolSum() constant returns(uint256)
func (_Predict *PredictSession) PoolSum() (*big.Int, error) {
	return _Predict.Contract.PoolSum(&_Predict.CallOpts)
}

// PoolSum is a free data retrieval call binding the contract method 0x26e309d9.
//
// Solidity: function poolSum() constant returns(uint256)
func (_Predict *PredictCallerSession) PoolSum() (*big.Int, error) {
	return _Predict.Contract.PoolSum(&_Predict.CallOpts)
}

// RangeMax is a free data retrieval call binding the contract method 0x5668972b.
//
// Solidity: function rangeMax() constant returns(uint256)
func (_Predict *PredictCaller) RangeMax(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Predict.contract.Call(opts, out, "rangeMax")
	return *ret0, err
}

// RangeMax is a free data retrieval call binding the contract method 0x5668972b.
//
// Solidity: function rangeMax() constant returns(uint256)
func (_Predict *PredictSession) RangeMax() (*big.Int, error) {
	return _Predict.Contract.RangeMax(&_Predict.CallOpts)
}

// RangeMax is a free data retrieval call binding the contract method 0x5668972b.
//
// Solidity: function rangeMax() constant returns(uint256)
func (_Predict *PredictCallerSession) RangeMax() (*big.Int, error) {
	return _Predict.Contract.RangeMax(&_Predict.CallOpts)
}

// StoreAdds is a free data retrieval call binding the contract method 0xa8a099ae.
//
// Solidity: function storeAdds(uint256 ) constant returns(address)
func (_Predict *PredictCaller) StoreAdds(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Predict.contract.Call(opts, out, "storeAdds", arg0)
	return *ret0, err
}

// StoreAdds is a free data retrieval call binding the contract method 0xa8a099ae.
//
// Solidity: function storeAdds(uint256 ) constant returns(address)
func (_Predict *PredictSession) StoreAdds(arg0 *big.Int) (common.Address, error) {
	return _Predict.Contract.StoreAdds(&_Predict.CallOpts, arg0)
}

// StoreAdds is a free data retrieval call binding the contract method 0xa8a099ae.
//
// Solidity: function storeAdds(uint256 ) constant returns(address)
func (_Predict *PredictCallerSession) StoreAdds(arg0 *big.Int) (common.Address, error) {
	return _Predict.Contract.StoreAdds(&_Predict.CallOpts, arg0)
}

// Bet is a paid mutator transaction binding the contract method 0x7365870b.
//
// Solidity: function bet(uint256 predict) returns()
func (_Predict *PredictTransactor) Bet(opts *bind.TransactOpts, predict *big.Int) (*types.Transaction, error) {
	return _Predict.contract.Transact(opts, "bet", predict)
}

// Bet is a paid mutator transaction binding the contract method 0x7365870b.
//
// Solidity: function bet(uint256 predict) returns()
func (_Predict *PredictSession) Bet(predict *big.Int) (*types.Transaction, error) {
	return _Predict.Contract.Bet(&_Predict.TransactOpts, predict)
}

// Bet is a paid mutator transaction binding the contract method 0x7365870b.
//
// Solidity: function bet(uint256 predict) returns()
func (_Predict *PredictTransactorSession) Bet(predict *big.Int) (*types.Transaction, error) {
	return _Predict.Contract.Bet(&_Predict.TransactOpts, predict)
}

// Resolve is a paid mutator transaction binding the contract method 0x4f896d4f.
//
// Solidity: function resolve(uint256 predict) returns()
func (_Predict *PredictTransactor) Resolve(opts *bind.TransactOpts, predict *big.Int) (*types.Transaction, error) {
	return _Predict.contract.Transact(opts, "resolve", predict)
}

// Resolve is a paid mutator transaction binding the contract method 0x4f896d4f.
//
// Solidity: function resolve(uint256 predict) returns()
func (_Predict *PredictSession) Resolve(predict *big.Int) (*types.Transaction, error) {
	return _Predict.Contract.Resolve(&_Predict.TransactOpts, predict)
}

// Resolve is a paid mutator transaction binding the contract method 0x4f896d4f.
//
// Solidity: function resolve(uint256 predict) returns()
func (_Predict *PredictTransactorSession) Resolve(predict *big.Int) (*types.Transaction, error) {
	return _Predict.Contract.Resolve(&_Predict.TransactOpts, predict)
}
