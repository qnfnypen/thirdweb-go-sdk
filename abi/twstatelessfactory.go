// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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
)

// TWStatelessFactoryMetaData contains all meta data concerning the TWStatelessFactory contract.
var TWStatelessFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustedForwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"deployer\",\"type\":\"address\"}],\"name\":\"ProxyDeployed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"}],\"name\":\"deployProxyByImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"deployedProxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TWStatelessFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use TWStatelessFactoryMetaData.ABI instead.
var TWStatelessFactoryABI = TWStatelessFactoryMetaData.ABI

// TWStatelessFactory is an auto generated Go binding around an Ethereum contract.
type TWStatelessFactory struct {
	TWStatelessFactoryCaller     // Read-only binding to the contract
	TWStatelessFactoryTransactor // Write-only binding to the contract
	TWStatelessFactoryFilterer   // Log filterer for contract events
}

// TWStatelessFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TWStatelessFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TWStatelessFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TWStatelessFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TWStatelessFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TWStatelessFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TWStatelessFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TWStatelessFactorySession struct {
	Contract     *TWStatelessFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TWStatelessFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TWStatelessFactoryCallerSession struct {
	Contract *TWStatelessFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TWStatelessFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TWStatelessFactoryTransactorSession struct {
	Contract     *TWStatelessFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TWStatelessFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TWStatelessFactoryRaw struct {
	Contract *TWStatelessFactory // Generic contract binding to access the raw methods on
}

// TWStatelessFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TWStatelessFactoryCallerRaw struct {
	Contract *TWStatelessFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// TWStatelessFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TWStatelessFactoryTransactorRaw struct {
	Contract *TWStatelessFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTWStatelessFactory creates a new instance of TWStatelessFactory, bound to a specific deployed contract.
func NewTWStatelessFactory(address common.Address, backend bind.ContractBackend) (*TWStatelessFactory, error) {
	contract, err := bindTWStatelessFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TWStatelessFactory{TWStatelessFactoryCaller: TWStatelessFactoryCaller{contract: contract}, TWStatelessFactoryTransactor: TWStatelessFactoryTransactor{contract: contract}, TWStatelessFactoryFilterer: TWStatelessFactoryFilterer{contract: contract}}, nil
}

// NewTWStatelessFactoryCaller creates a new read-only instance of TWStatelessFactory, bound to a specific deployed contract.
func NewTWStatelessFactoryCaller(address common.Address, caller bind.ContractCaller) (*TWStatelessFactoryCaller, error) {
	contract, err := bindTWStatelessFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TWStatelessFactoryCaller{contract: contract}, nil
}

// NewTWStatelessFactoryTransactor creates a new write-only instance of TWStatelessFactory, bound to a specific deployed contract.
func NewTWStatelessFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*TWStatelessFactoryTransactor, error) {
	contract, err := bindTWStatelessFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TWStatelessFactoryTransactor{contract: contract}, nil
}

// NewTWStatelessFactoryFilterer creates a new log filterer instance of TWStatelessFactory, bound to a specific deployed contract.
func NewTWStatelessFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*TWStatelessFactoryFilterer, error) {
	contract, err := bindTWStatelessFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TWStatelessFactoryFilterer{contract: contract}, nil
}

// bindTWStatelessFactory binds a generic wrapper to an already deployed contract.
func bindTWStatelessFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TWStatelessFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TWStatelessFactory *TWStatelessFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TWStatelessFactory.Contract.TWStatelessFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TWStatelessFactory *TWStatelessFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TWStatelessFactory.Contract.TWStatelessFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TWStatelessFactory *TWStatelessFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TWStatelessFactory.Contract.TWStatelessFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TWStatelessFactory *TWStatelessFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TWStatelessFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TWStatelessFactory *TWStatelessFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TWStatelessFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TWStatelessFactory *TWStatelessFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TWStatelessFactory.Contract.contract.Transact(opts, method, params...)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_TWStatelessFactory *TWStatelessFactoryCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _TWStatelessFactory.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_TWStatelessFactory *TWStatelessFactorySession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _TWStatelessFactory.Contract.IsTrustedForwarder(&_TWStatelessFactory.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_TWStatelessFactory *TWStatelessFactoryCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _TWStatelessFactory.Contract.IsTrustedForwarder(&_TWStatelessFactory.CallOpts, forwarder)
}

// DeployProxyByImplementation is a paid mutator transaction binding the contract method 0x11b804ab.
//
// Solidity: function deployProxyByImplementation(address _implementation, bytes _data, bytes32 _salt) returns(address deployedProxy)
func (_TWStatelessFactory *TWStatelessFactoryTransactor) DeployProxyByImplementation(opts *bind.TransactOpts, _implementation common.Address, _data []byte, _salt [32]byte) (*types.Transaction, error) {
	return _TWStatelessFactory.contract.Transact(opts, "deployProxyByImplementation", _implementation, _data, _salt)
}

// DeployProxyByImplementation is a paid mutator transaction binding the contract method 0x11b804ab.
//
// Solidity: function deployProxyByImplementation(address _implementation, bytes _data, bytes32 _salt) returns(address deployedProxy)
func (_TWStatelessFactory *TWStatelessFactorySession) DeployProxyByImplementation(_implementation common.Address, _data []byte, _salt [32]byte) (*types.Transaction, error) {
	return _TWStatelessFactory.Contract.DeployProxyByImplementation(&_TWStatelessFactory.TransactOpts, _implementation, _data, _salt)
}

// DeployProxyByImplementation is a paid mutator transaction binding the contract method 0x11b804ab.
//
// Solidity: function deployProxyByImplementation(address _implementation, bytes _data, bytes32 _salt) returns(address deployedProxy)
func (_TWStatelessFactory *TWStatelessFactoryTransactorSession) DeployProxyByImplementation(_implementation common.Address, _data []byte, _salt [32]byte) (*types.Transaction, error) {
	return _TWStatelessFactory.Contract.DeployProxyByImplementation(&_TWStatelessFactory.TransactOpts, _implementation, _data, _salt)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_TWStatelessFactory *TWStatelessFactoryTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _TWStatelessFactory.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_TWStatelessFactory *TWStatelessFactorySession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _TWStatelessFactory.Contract.Multicall(&_TWStatelessFactory.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_TWStatelessFactory *TWStatelessFactoryTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _TWStatelessFactory.Contract.Multicall(&_TWStatelessFactory.TransactOpts, data)
}

// TWStatelessFactoryProxyDeployedIterator is returned from FilterProxyDeployed and is used to iterate over the raw logs and unpacked data for ProxyDeployed events raised by the TWStatelessFactory contract.
type TWStatelessFactoryProxyDeployedIterator struct {
	Event *TWStatelessFactoryProxyDeployed // Event containing the contract specifics and raw log

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
func (it *TWStatelessFactoryProxyDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TWStatelessFactoryProxyDeployed)
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
		it.Event = new(TWStatelessFactoryProxyDeployed)
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
func (it *TWStatelessFactoryProxyDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TWStatelessFactoryProxyDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TWStatelessFactoryProxyDeployed represents a ProxyDeployed event raised by the TWStatelessFactory contract.
type TWStatelessFactoryProxyDeployed struct {
	Implementation common.Address
	Proxy          common.Address
	Deployer       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterProxyDeployed is a free log retrieval operation binding the contract event 0x9e0862c4ebff2150fbbfd3f8547483f55bdec0c34fd977d3fccaa55d6c4ce784.
//
// Solidity: event ProxyDeployed(address indexed implementation, address proxy, address indexed deployer)
func (_TWStatelessFactory *TWStatelessFactoryFilterer) FilterProxyDeployed(opts *bind.FilterOpts, implementation []common.Address, deployer []common.Address) (*TWStatelessFactoryProxyDeployedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	var deployerRule []interface{}
	for _, deployerItem := range deployer {
		deployerRule = append(deployerRule, deployerItem)
	}

	logs, sub, err := _TWStatelessFactory.contract.FilterLogs(opts, "ProxyDeployed", implementationRule, deployerRule)
	if err != nil {
		return nil, err
	}
	return &TWStatelessFactoryProxyDeployedIterator{contract: _TWStatelessFactory.contract, event: "ProxyDeployed", logs: logs, sub: sub}, nil
}

// WatchProxyDeployed is a free log subscription operation binding the contract event 0x9e0862c4ebff2150fbbfd3f8547483f55bdec0c34fd977d3fccaa55d6c4ce784.
//
// Solidity: event ProxyDeployed(address indexed implementation, address proxy, address indexed deployer)
func (_TWStatelessFactory *TWStatelessFactoryFilterer) WatchProxyDeployed(opts *bind.WatchOpts, sink chan<- *TWStatelessFactoryProxyDeployed, implementation []common.Address, deployer []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	var deployerRule []interface{}
	for _, deployerItem := range deployer {
		deployerRule = append(deployerRule, deployerItem)
	}

	logs, sub, err := _TWStatelessFactory.contract.WatchLogs(opts, "ProxyDeployed", implementationRule, deployerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TWStatelessFactoryProxyDeployed)
				if err := _TWStatelessFactory.contract.UnpackLog(event, "ProxyDeployed", log); err != nil {
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

// ParseProxyDeployed is a log parse operation binding the contract event 0x9e0862c4ebff2150fbbfd3f8547483f55bdec0c34fd977d3fccaa55d6c4ce784.
//
// Solidity: event ProxyDeployed(address indexed implementation, address proxy, address indexed deployer)
func (_TWStatelessFactory *TWStatelessFactoryFilterer) ParseProxyDeployed(log types.Log) (*TWStatelessFactoryProxyDeployed, error) {
	event := new(TWStatelessFactoryProxyDeployed)
	if err := _TWStatelessFactory.contract.UnpackLog(event, "ProxyDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
