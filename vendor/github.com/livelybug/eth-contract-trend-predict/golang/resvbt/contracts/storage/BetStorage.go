// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package storage

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

// StorageABI is the input ABI used to generate the binding from.
const StorageABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"bettor\",\"type\":\"address\"}],\"name\":\"addBettor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bettors\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wager\",\"type\":\"uint256\"}],\"name\":\"addBetPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"betMapping\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"bettor\",\"type\":\"address\"},{\"name\":\"wager\",\"type\":\"uint256\"}],\"name\":\"addBetMapping\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBettorsLen\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"range\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"betPool\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// BetMapping is a free data retrieval call binding the contract method 0xa3f763bf.
//
// Solidity: function betMapping(address ) constant returns(uint256)
func (_Storage *StorageCaller) BetMapping(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "betMapping", arg0)
	return *ret0, err
}

// BetMapping is a free data retrieval call binding the contract method 0xa3f763bf.
//
// Solidity: function betMapping(address ) constant returns(uint256)
func (_Storage *StorageSession) BetMapping(arg0 common.Address) (*big.Int, error) {
	return _Storage.Contract.BetMapping(&_Storage.CallOpts, arg0)
}

// BetMapping is a free data retrieval call binding the contract method 0xa3f763bf.
//
// Solidity: function betMapping(address ) constant returns(uint256)
func (_Storage *StorageCallerSession) BetMapping(arg0 common.Address) (*big.Int, error) {
	return _Storage.Contract.BetMapping(&_Storage.CallOpts, arg0)
}

// BetPool is a free data retrieval call binding the contract method 0xec90bfc7.
//
// Solidity: function betPool() constant returns(uint256)
func (_Storage *StorageCaller) BetPool(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "betPool")
	return *ret0, err
}

// BetPool is a free data retrieval call binding the contract method 0xec90bfc7.
//
// Solidity: function betPool() constant returns(uint256)
func (_Storage *StorageSession) BetPool() (*big.Int, error) {
	return _Storage.Contract.BetPool(&_Storage.CallOpts)
}

// BetPool is a free data retrieval call binding the contract method 0xec90bfc7.
//
// Solidity: function betPool() constant returns(uint256)
func (_Storage *StorageCallerSession) BetPool() (*big.Int, error) {
	return _Storage.Contract.BetPool(&_Storage.CallOpts)
}

// Bettors is a free data retrieval call binding the contract method 0x3482b644.
//
// Solidity: function bettors(uint256 ) constant returns(address)
func (_Storage *StorageCaller) Bettors(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "bettors", arg0)
	return *ret0, err
}

// Bettors is a free data retrieval call binding the contract method 0x3482b644.
//
// Solidity: function bettors(uint256 ) constant returns(address)
func (_Storage *StorageSession) Bettors(arg0 *big.Int) (common.Address, error) {
	return _Storage.Contract.Bettors(&_Storage.CallOpts, arg0)
}

// Bettors is a free data retrieval call binding the contract method 0x3482b644.
//
// Solidity: function bettors(uint256 ) constant returns(address)
func (_Storage *StorageCallerSession) Bettors(arg0 *big.Int) (common.Address, error) {
	return _Storage.Contract.Bettors(&_Storage.CallOpts, arg0)
}

// GetBettorsLen is a free data retrieval call binding the contract method 0xe0d613ef.
//
// Solidity: function getBettorsLen() constant returns(uint256)
func (_Storage *StorageCaller) GetBettorsLen(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "getBettorsLen")
	return *ret0, err
}

// GetBettorsLen is a free data retrieval call binding the contract method 0xe0d613ef.
//
// Solidity: function getBettorsLen() constant returns(uint256)
func (_Storage *StorageSession) GetBettorsLen() (*big.Int, error) {
	return _Storage.Contract.GetBettorsLen(&_Storage.CallOpts)
}

// GetBettorsLen is a free data retrieval call binding the contract method 0xe0d613ef.
//
// Solidity: function getBettorsLen() constant returns(uint256)
func (_Storage *StorageCallerSession) GetBettorsLen() (*big.Int, error) {
	return _Storage.Contract.GetBettorsLen(&_Storage.CallOpts)
}

// Range is a free data retrieval call binding the contract method 0xe97206a9.
//
// Solidity: function range() constant returns(uint256)
func (_Storage *StorageCaller) Range(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "range")
	return *ret0, err
}

// Range is a free data retrieval call binding the contract method 0xe97206a9.
//
// Solidity: function range() constant returns(uint256)
func (_Storage *StorageSession) Range() (*big.Int, error) {
	return _Storage.Contract.Range(&_Storage.CallOpts)
}

// Range is a free data retrieval call binding the contract method 0xe97206a9.
//
// Solidity: function range() constant returns(uint256)
func (_Storage *StorageCallerSession) Range() (*big.Int, error) {
	return _Storage.Contract.Range(&_Storage.CallOpts)
}

// AddBetMapping is a paid mutator transaction binding the contract method 0xa8dbc83d.
//
// Solidity: function addBetMapping(address bettor, uint256 wager) returns()
func (_Storage *StorageTransactor) AddBetMapping(opts *bind.TransactOpts, bettor common.Address, wager *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addBetMapping", bettor, wager)
}

// AddBetMapping is a paid mutator transaction binding the contract method 0xa8dbc83d.
//
// Solidity: function addBetMapping(address bettor, uint256 wager) returns()
func (_Storage *StorageSession) AddBetMapping(bettor common.Address, wager *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddBetMapping(&_Storage.TransactOpts, bettor, wager)
}

// AddBetMapping is a paid mutator transaction binding the contract method 0xa8dbc83d.
//
// Solidity: function addBetMapping(address bettor, uint256 wager) returns()
func (_Storage *StorageTransactorSession) AddBetMapping(bettor common.Address, wager *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddBetMapping(&_Storage.TransactOpts, bettor, wager)
}

// AddBetPool is a paid mutator transaction binding the contract method 0x8b2bd927.
//
// Solidity: function addBetPool(uint256 wager) returns()
func (_Storage *StorageTransactor) AddBetPool(opts *bind.TransactOpts, wager *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addBetPool", wager)
}

// AddBetPool is a paid mutator transaction binding the contract method 0x8b2bd927.
//
// Solidity: function addBetPool(uint256 wager) returns()
func (_Storage *StorageSession) AddBetPool(wager *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddBetPool(&_Storage.TransactOpts, wager)
}

// AddBetPool is a paid mutator transaction binding the contract method 0x8b2bd927.
//
// Solidity: function addBetPool(uint256 wager) returns()
func (_Storage *StorageTransactorSession) AddBetPool(wager *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddBetPool(&_Storage.TransactOpts, wager)
}

// AddBettor is a paid mutator transaction binding the contract method 0x2497cbc3.
//
// Solidity: function addBettor(address bettor) returns()
func (_Storage *StorageTransactor) AddBettor(opts *bind.TransactOpts, bettor common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addBettor", bettor)
}

// AddBettor is a paid mutator transaction binding the contract method 0x2497cbc3.
//
// Solidity: function addBettor(address bettor) returns()
func (_Storage *StorageSession) AddBettor(bettor common.Address) (*types.Transaction, error) {
	return _Storage.Contract.AddBettor(&_Storage.TransactOpts, bettor)
}

// AddBettor is a paid mutator transaction binding the contract method 0x2497cbc3.
//
// Solidity: function addBettor(address bettor) returns()
func (_Storage *StorageTransactorSession) AddBettor(bettor common.Address) (*types.Transaction, error) {
	return _Storage.Contract.AddBettor(&_Storage.TransactOpts, bettor)
}
