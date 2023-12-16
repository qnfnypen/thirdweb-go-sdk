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

// SplitMetaData contains all meta data concerning the Split contract.
var SplitMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"name\":\"\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ERC20PaymentReleased\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"indexed\":true,\"internalType\":\"contractIERC20Upgradeable\"},{\"type\":\"address\",\"name\":\"to\",\"indexed\":false,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amount\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"version\",\"indexed\":false,\"internalType\":\"uint8\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PayeeAdded\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"indexed\":false,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"shares\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentReceived\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"indexed\":false,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amount\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentReleased\",\"inputs\":[{\"type\":\"address\",\"name\":\"to\",\"indexed\":false,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amount\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"type\":\"bytes32\",\"name\":\"previousAdminRole\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"type\":\"bytes32\",\"name\":\"newAdminRole\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"sender\",\"indexed\":true,\"internalType\":\"address\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"sender\",\"indexed\":true,\"internalType\":\"address\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractType\",\"inputs\":[],\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"contractURI\",\"inputs\":[],\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractVersion\",\"inputs\":[],\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"distribute\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"contractIERC20Upgradeable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"distribute\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMember\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"uint256\",\"name\":\"index\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMemberCount\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}],\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"type\":\"address\",\"name\":\"_defaultAdmin\",\"internalType\":\"address\"},{\"type\":\"string\",\"name\":\"_contractURI\",\"internalType\":\"string\"},{\"type\":\"address[]\",\"name\":\"_trustedForwarders\",\"internalType\":\"address[]\"},{\"type\":\"address[]\",\"name\":\"_payees\",\"internalType\":\"address[]\"},{\"type\":\"uint256[]\",\"name\":\"_shares\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isTrustedForwarder\",\"inputs\":[{\"type\":\"address\",\"name\":\"forwarder\",\"internalType\":\"address\"}],\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"multicall\",\"inputs\":[{\"type\":\"bytes[]\",\"name\":\"data\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"type\":\"bytes[]\",\"name\":\"results\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"payee\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"index\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payeeCount\",\"inputs\":[],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releasable\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releasable\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"contractIERC20Upgradeable\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"release\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"release\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"contractIERC20Upgradeable\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"released\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"contractIERC20Upgradeable\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"released\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setContractURI\",\"inputs\":[{\"type\":\"string\",\"name\":\"_uri\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shares\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"type\":\"bytes4\",\"name\":\"interfaceId\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalReleased\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"contractIERC20Upgradeable\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalReleased\",\"inputs\":[],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"receive\",\"name\":\"\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"}]",
}

// SplitABI is the input ABI used to generate the binding from.
// Deprecated: Use SplitMetaData.ABI instead.
var SplitABI = SplitMetaData.ABI

// Split is an auto generated Go binding around an Ethereum contract.
type Split struct {
	SplitCaller     // Read-only binding to the contract
	SplitTransactor // Write-only binding to the contract
	SplitFilterer   // Log filterer for contract events
}

// SplitCaller is an auto generated read-only Go binding around an Ethereum contract.
type SplitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SplitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SplitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SplitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SplitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SplitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SplitSession struct {
	Contract     *Split            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SplitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SplitCallerSession struct {
	Contract *SplitCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SplitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SplitTransactorSession struct {
	Contract     *SplitTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SplitRaw is an auto generated low-level Go binding around an Ethereum contract.
type SplitRaw struct {
	Contract *Split // Generic contract binding to access the raw methods on
}

// SplitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SplitCallerRaw struct {
	Contract *SplitCaller // Generic read-only contract binding to access the raw methods on
}

// SplitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SplitTransactorRaw struct {
	Contract *SplitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSplit creates a new instance of Split, bound to a specific deployed contract.
func NewSplit(address common.Address, backend bind.ContractBackend) (*Split, error) {
	contract, err := bindSplit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Split{SplitCaller: SplitCaller{contract: contract}, SplitTransactor: SplitTransactor{contract: contract}, SplitFilterer: SplitFilterer{contract: contract}}, nil
}

// NewSplitCaller creates a new read-only instance of Split, bound to a specific deployed contract.
func NewSplitCaller(address common.Address, caller bind.ContractCaller) (*SplitCaller, error) {
	contract, err := bindSplit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SplitCaller{contract: contract}, nil
}

// NewSplitTransactor creates a new write-only instance of Split, bound to a specific deployed contract.
func NewSplitTransactor(address common.Address, transactor bind.ContractTransactor) (*SplitTransactor, error) {
	contract, err := bindSplit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SplitTransactor{contract: contract}, nil
}

// NewSplitFilterer creates a new log filterer instance of Split, bound to a specific deployed contract.
func NewSplitFilterer(address common.Address, filterer bind.ContractFilterer) (*SplitFilterer, error) {
	contract, err := bindSplit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SplitFilterer{contract: contract}, nil
}

// bindSplit binds a generic wrapper to an already deployed contract.
func bindSplit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SplitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Split *SplitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Split.Contract.SplitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Split *SplitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Split.Contract.SplitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Split *SplitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Split.Contract.SplitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Split *SplitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Split.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Split *SplitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Split.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Split *SplitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Split.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Split *SplitCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Split.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Split *SplitSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Split.Contract.DEFAULTADMINROLE(&_Split.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Split *SplitCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Split.Contract.DEFAULTADMINROLE(&_Split.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() pure returns(bytes32)
func (_Split *SplitCaller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Split.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() pure returns(bytes32)
func (_Split *SplitSession) ContractType() ([32]byte, error) {
	return _Split.Contract.ContractType(&_Split.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() pure returns(bytes32)
func (_Split *SplitCallerSession) ContractType() ([32]byte, error) {
	return _Split.Contract.ContractType(&_Split.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Split *SplitCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Split.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Split *SplitSession) ContractURI() (string, error) {
	return _Split.Contract.ContractURI(&_Split.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Split *SplitCallerSession) ContractURI() (string, error) {
	return _Split.Contract.ContractURI(&_Split.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint8)
func (_Split *SplitCaller) ContractVersion(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Split.contract.Call(opts, &out, "contractVersion")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint8)
func (_Split *SplitSession) ContractVersion() (uint8, error) {
	return _Split.Contract.ContractVersion(&_Split.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint8)
func (_Split *SplitCallerSession) ContractVersion() (uint8, error) {
	return _Split.Contract.ContractVersion(&_Split.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Split *SplitCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Split.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Split *SplitSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Split.Contract.GetRoleAdmin(&_Split.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Split *SplitCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Split.Contract.GetRoleAdmin(&_Split.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Split *SplitCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Split.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Split *SplitSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Split.Contract.GetRoleMember(&_Split.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Split *SplitCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Split.Contract.GetRoleMember(&_Split.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Split *SplitCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Split.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Split *SplitSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Split.Contract.GetRoleMemberCount(&_Split.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Split *SplitCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Split.Contract.GetRoleMemberCount(&_Split.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Split *SplitCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Split.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Split *SplitSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Split.Contract.HasRole(&_Split.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Split *SplitCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Split.Contract.HasRole(&_Split.CallOpts, role, account)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Split *SplitCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _Split.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Split *SplitSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Split.Contract.IsTrustedForwarder(&_Split.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Split *SplitCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Split.Contract.IsTrustedForwarder(&_Split.CallOpts, forwarder)
}

// Payee is a free data retrieval call binding the contract method 0x8b83209b.
//
// Solidity: function payee(uint256 index) view returns(address)
func (_Split *SplitCaller) Payee(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Split.contract.Call(opts, &out, "payee", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Payee is a free data retrieval call binding the contract method 0x8b83209b.
//
// Solidity: function payee(uint256 index) view returns(address)
func (_Split *SplitSession) Payee(index *big.Int) (common.Address, error) {
	return _Split.Contract.Payee(&_Split.CallOpts, index)
}

// Payee is a free data retrieval call binding the contract method 0x8b83209b.
//
// Solidity: function payee(uint256 index) view returns(address)
func (_Split *SplitCallerSession) Payee(index *big.Int) (common.Address, error) {
	return _Split.Contract.Payee(&_Split.CallOpts, index)
}

// PayeeCount is a free data retrieval call binding the contract method 0x00dbe109.
//
// Solidity: function payeeCount() view returns(uint256)
func (_Split *SplitCaller) PayeeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Split.contract.Call(opts, &out, "payeeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PayeeCount is a free data retrieval call binding the contract method 0x00dbe109.
//
// Solidity: function payeeCount() view returns(uint256)
func (_Split *SplitSession) PayeeCount() (*big.Int, error) {
	return _Split.Contract.PayeeCount(&_Split.CallOpts)
}

// PayeeCount is a free data retrieval call binding the contract method 0x00dbe109.
//
// Solidity: function payeeCount() view returns(uint256)
func (_Split *SplitCallerSession) PayeeCount() (*big.Int, error) {
	return _Split.Contract.PayeeCount(&_Split.CallOpts)
}

// Releasable is a free data retrieval call binding the contract method 0xa3f8eace.
//
// Solidity: function releasable(address account) view returns(uint256)
func (_Split *SplitCaller) Releasable(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Split.contract.Call(opts, &out, "releasable", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Releasable is a free data retrieval call binding the contract method 0xa3f8eace.
//
// Solidity: function releasable(address account) view returns(uint256)
func (_Split *SplitSession) Releasable(account common.Address) (*big.Int, error) {
	return _Split.Contract.Releasable(&_Split.CallOpts, account)
}

// Releasable is a free data retrieval call binding the contract method 0xa3f8eace.
//
// Solidity: function releasable(address account) view returns(uint256)
func (_Split *SplitCallerSession) Releasable(account common.Address) (*big.Int, error) {
	return _Split.Contract.Releasable(&_Split.CallOpts, account)
}

// Releasable0 is a free data retrieval call binding the contract method 0xc45ac050.
//
// Solidity: function releasable(address token, address account) view returns(uint256)
func (_Split *SplitCaller) Releasable0(opts *bind.CallOpts, token common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Split.contract.Call(opts, &out, "releasable0", token, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Releasable0 is a free data retrieval call binding the contract method 0xc45ac050.
//
// Solidity: function releasable(address token, address account) view returns(uint256)
func (_Split *SplitSession) Releasable0(token common.Address, account common.Address) (*big.Int, error) {
	return _Split.Contract.Releasable0(&_Split.CallOpts, token, account)
}

// Releasable0 is a free data retrieval call binding the contract method 0xc45ac050.
//
// Solidity: function releasable(address token, address account) view returns(uint256)
func (_Split *SplitCallerSession) Releasable0(token common.Address, account common.Address) (*big.Int, error) {
	return _Split.Contract.Releasable0(&_Split.CallOpts, token, account)
}

// Released is a free data retrieval call binding the contract method 0x406072a9.
//
// Solidity: function released(address token, address account) view returns(uint256)
func (_Split *SplitCaller) Released(opts *bind.CallOpts, token common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Split.contract.Call(opts, &out, "released", token, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Released is a free data retrieval call binding the contract method 0x406072a9.
//
// Solidity: function released(address token, address account) view returns(uint256)
func (_Split *SplitSession) Released(token common.Address, account common.Address) (*big.Int, error) {
	return _Split.Contract.Released(&_Split.CallOpts, token, account)
}

// Released is a free data retrieval call binding the contract method 0x406072a9.
//
// Solidity: function released(address token, address account) view returns(uint256)
func (_Split *SplitCallerSession) Released(token common.Address, account common.Address) (*big.Int, error) {
	return _Split.Contract.Released(&_Split.CallOpts, token, account)
}

// Released0 is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released(address account) view returns(uint256)
func (_Split *SplitCaller) Released0(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Split.contract.Call(opts, &out, "released0", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Released0 is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released(address account) view returns(uint256)
func (_Split *SplitSession) Released0(account common.Address) (*big.Int, error) {
	return _Split.Contract.Released0(&_Split.CallOpts, account)
}

// Released0 is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released(address account) view returns(uint256)
func (_Split *SplitCallerSession) Released0(account common.Address) (*big.Int, error) {
	return _Split.Contract.Released0(&_Split.CallOpts, account)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address account) view returns(uint256)
func (_Split *SplitCaller) Shares(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Split.contract.Call(opts, &out, "shares", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address account) view returns(uint256)
func (_Split *SplitSession) Shares(account common.Address) (*big.Int, error) {
	return _Split.Contract.Shares(&_Split.CallOpts, account)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address account) view returns(uint256)
func (_Split *SplitCallerSession) Shares(account common.Address) (*big.Int, error) {
	return _Split.Contract.Shares(&_Split.CallOpts, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Split *SplitCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Split.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Split *SplitSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Split.Contract.SupportsInterface(&_Split.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Split *SplitCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Split.Contract.SupportsInterface(&_Split.CallOpts, interfaceId)
}

// TotalReleased is a free data retrieval call binding the contract method 0xd79779b2.
//
// Solidity: function totalReleased(address token) view returns(uint256)
func (_Split *SplitCaller) TotalReleased(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Split.contract.Call(opts, &out, "totalReleased", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReleased is a free data retrieval call binding the contract method 0xd79779b2.
//
// Solidity: function totalReleased(address token) view returns(uint256)
func (_Split *SplitSession) TotalReleased(token common.Address) (*big.Int, error) {
	return _Split.Contract.TotalReleased(&_Split.CallOpts, token)
}

// TotalReleased is a free data retrieval call binding the contract method 0xd79779b2.
//
// Solidity: function totalReleased(address token) view returns(uint256)
func (_Split *SplitCallerSession) TotalReleased(token common.Address) (*big.Int, error) {
	return _Split.Contract.TotalReleased(&_Split.CallOpts, token)
}

// TotalReleased0 is a free data retrieval call binding the contract method 0xe33b7de3.
//
// Solidity: function totalReleased() view returns(uint256)
func (_Split *SplitCaller) TotalReleased0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Split.contract.Call(opts, &out, "totalReleased0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReleased0 is a free data retrieval call binding the contract method 0xe33b7de3.
//
// Solidity: function totalReleased() view returns(uint256)
func (_Split *SplitSession) TotalReleased0() (*big.Int, error) {
	return _Split.Contract.TotalReleased0(&_Split.CallOpts)
}

// TotalReleased0 is a free data retrieval call binding the contract method 0xe33b7de3.
//
// Solidity: function totalReleased() view returns(uint256)
func (_Split *SplitCallerSession) TotalReleased0() (*big.Int, error) {
	return _Split.Contract.TotalReleased0(&_Split.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Split *SplitCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Split.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Split *SplitSession) TotalShares() (*big.Int, error) {
	return _Split.Contract.TotalShares(&_Split.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Split *SplitCallerSession) TotalShares() (*big.Int, error) {
	return _Split.Contract.TotalShares(&_Split.CallOpts)
}

// Distribute is a paid mutator transaction binding the contract method 0x63453ae1.
//
// Solidity: function distribute(address token) returns()
func (_Split *SplitTransactor) Distribute(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Split.contract.Transact(opts, "distribute", token)
}

// Distribute is a paid mutator transaction binding the contract method 0x63453ae1.
//
// Solidity: function distribute(address token) returns()
func (_Split *SplitSession) Distribute(token common.Address) (*types.Transaction, error) {
	return _Split.Contract.Distribute(&_Split.TransactOpts, token)
}

// Distribute is a paid mutator transaction binding the contract method 0x63453ae1.
//
// Solidity: function distribute(address token) returns()
func (_Split *SplitTransactorSession) Distribute(token common.Address) (*types.Transaction, error) {
	return _Split.Contract.Distribute(&_Split.TransactOpts, token)
}

// Distribute0 is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns()
func (_Split *SplitTransactor) Distribute0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Split.contract.Transact(opts, "distribute0")
}

// Distribute0 is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns()
func (_Split *SplitSession) Distribute0() (*types.Transaction, error) {
	return _Split.Contract.Distribute0(&_Split.TransactOpts)
}

// Distribute0 is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns()
func (_Split *SplitTransactorSession) Distribute0() (*types.Transaction, error) {
	return _Split.Contract.Distribute0(&_Split.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Split *SplitTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Split.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Split *SplitSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Split.Contract.GrantRole(&_Split.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Split *SplitTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Split.Contract.GrantRole(&_Split.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xb1a14437.
//
// Solidity: function initialize(address _defaultAdmin, string _contractURI, address[] _trustedForwarders, address[] _payees, uint256[] _shares) returns()
func (_Split *SplitTransactor) Initialize(opts *bind.TransactOpts, _defaultAdmin common.Address, _contractURI string, _trustedForwarders []common.Address, _payees []common.Address, _shares []*big.Int) (*types.Transaction, error) {
	return _Split.contract.Transact(opts, "initialize", _defaultAdmin, _contractURI, _trustedForwarders, _payees, _shares)
}

// Initialize is a paid mutator transaction binding the contract method 0xb1a14437.
//
// Solidity: function initialize(address _defaultAdmin, string _contractURI, address[] _trustedForwarders, address[] _payees, uint256[] _shares) returns()
func (_Split *SplitSession) Initialize(_defaultAdmin common.Address, _contractURI string, _trustedForwarders []common.Address, _payees []common.Address, _shares []*big.Int) (*types.Transaction, error) {
	return _Split.Contract.Initialize(&_Split.TransactOpts, _defaultAdmin, _contractURI, _trustedForwarders, _payees, _shares)
}

// Initialize is a paid mutator transaction binding the contract method 0xb1a14437.
//
// Solidity: function initialize(address _defaultAdmin, string _contractURI, address[] _trustedForwarders, address[] _payees, uint256[] _shares) returns()
func (_Split *SplitTransactorSession) Initialize(_defaultAdmin common.Address, _contractURI string, _trustedForwarders []common.Address, _payees []common.Address, _shares []*big.Int) (*types.Transaction, error) {
	return _Split.Contract.Initialize(&_Split.TransactOpts, _defaultAdmin, _contractURI, _trustedForwarders, _payees, _shares)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Split *SplitTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Split.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Split *SplitSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Split.Contract.Multicall(&_Split.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Split *SplitTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Split.Contract.Multicall(&_Split.TransactOpts, data)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address account) returns()
func (_Split *SplitTransactor) Release(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Split.contract.Transact(opts, "release", account)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address account) returns()
func (_Split *SplitSession) Release(account common.Address) (*types.Transaction, error) {
	return _Split.Contract.Release(&_Split.TransactOpts, account)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address account) returns()
func (_Split *SplitTransactorSession) Release(account common.Address) (*types.Transaction, error) {
	return _Split.Contract.Release(&_Split.TransactOpts, account)
}

// Release0 is a paid mutator transaction binding the contract method 0x48b75044.
//
// Solidity: function release(address token, address account) returns()
func (_Split *SplitTransactor) Release0(opts *bind.TransactOpts, token common.Address, account common.Address) (*types.Transaction, error) {
	return _Split.contract.Transact(opts, "release0", token, account)
}

// Release0 is a paid mutator transaction binding the contract method 0x48b75044.
//
// Solidity: function release(address token, address account) returns()
func (_Split *SplitSession) Release0(token common.Address, account common.Address) (*types.Transaction, error) {
	return _Split.Contract.Release0(&_Split.TransactOpts, token, account)
}

// Release0 is a paid mutator transaction binding the contract method 0x48b75044.
//
// Solidity: function release(address token, address account) returns()
func (_Split *SplitTransactorSession) Release0(token common.Address, account common.Address) (*types.Transaction, error) {
	return _Split.Contract.Release0(&_Split.TransactOpts, token, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Split *SplitTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Split.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Split *SplitSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Split.Contract.RenounceRole(&_Split.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Split *SplitTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Split.Contract.RenounceRole(&_Split.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Split *SplitTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Split.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Split *SplitSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Split.Contract.RevokeRole(&_Split.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Split *SplitTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Split.Contract.RevokeRole(&_Split.TransactOpts, role, account)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_Split *SplitTransactor) SetContractURI(opts *bind.TransactOpts, _uri string) (*types.Transaction, error) {
	return _Split.contract.Transact(opts, "setContractURI", _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_Split *SplitSession) SetContractURI(_uri string) (*types.Transaction, error) {
	return _Split.Contract.SetContractURI(&_Split.TransactOpts, _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_Split *SplitTransactorSession) SetContractURI(_uri string) (*types.Transaction, error) {
	return _Split.Contract.SetContractURI(&_Split.TransactOpts, _uri)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Split *SplitTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Split.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Split *SplitSession) Receive() (*types.Transaction, error) {
	return _Split.Contract.Receive(&_Split.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Split *SplitTransactorSession) Receive() (*types.Transaction, error) {
	return _Split.Contract.Receive(&_Split.TransactOpts)
}

// SplitERC20PaymentReleasedIterator is returned from FilterERC20PaymentReleased and is used to iterate over the raw logs and unpacked data for ERC20PaymentReleased events raised by the Split contract.
type SplitERC20PaymentReleasedIterator struct {
	Event *SplitERC20PaymentReleased // Event containing the contract specifics and raw log

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
func (it *SplitERC20PaymentReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SplitERC20PaymentReleased)
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
		it.Event = new(SplitERC20PaymentReleased)
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
func (it *SplitERC20PaymentReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SplitERC20PaymentReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SplitERC20PaymentReleased represents a ERC20PaymentReleased event raised by the Split contract.
type SplitERC20PaymentReleased struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterERC20PaymentReleased is a free log retrieval operation binding the contract event 0x3be5b7a71e84ed12875d241991c70855ac5817d847039e17a9d895c1ceb0f18a.
//
// Solidity: event ERC20PaymentReleased(address indexed token, address to, uint256 amount)
func (_Split *SplitFilterer) FilterERC20PaymentReleased(opts *bind.FilterOpts, token []common.Address) (*SplitERC20PaymentReleasedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Split.contract.FilterLogs(opts, "ERC20PaymentReleased", tokenRule)
	if err != nil {
		return nil, err
	}
	return &SplitERC20PaymentReleasedIterator{contract: _Split.contract, event: "ERC20PaymentReleased", logs: logs, sub: sub}, nil
}

// WatchERC20PaymentReleased is a free log subscription operation binding the contract event 0x3be5b7a71e84ed12875d241991c70855ac5817d847039e17a9d895c1ceb0f18a.
//
// Solidity: event ERC20PaymentReleased(address indexed token, address to, uint256 amount)
func (_Split *SplitFilterer) WatchERC20PaymentReleased(opts *bind.WatchOpts, sink chan<- *SplitERC20PaymentReleased, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Split.contract.WatchLogs(opts, "ERC20PaymentReleased", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SplitERC20PaymentReleased)
				if err := _Split.contract.UnpackLog(event, "ERC20PaymentReleased", log); err != nil {
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

// ParseERC20PaymentReleased is a log parse operation binding the contract event 0x3be5b7a71e84ed12875d241991c70855ac5817d847039e17a9d895c1ceb0f18a.
//
// Solidity: event ERC20PaymentReleased(address indexed token, address to, uint256 amount)
func (_Split *SplitFilterer) ParseERC20PaymentReleased(log types.Log) (*SplitERC20PaymentReleased, error) {
	event := new(SplitERC20PaymentReleased)
	if err := _Split.contract.UnpackLog(event, "ERC20PaymentReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SplitInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Split contract.
type SplitInitializedIterator struct {
	Event *SplitInitialized // Event containing the contract specifics and raw log

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
func (it *SplitInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SplitInitialized)
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
		it.Event = new(SplitInitialized)
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
func (it *SplitInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SplitInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SplitInitialized represents a Initialized event raised by the Split contract.
type SplitInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Split *SplitFilterer) FilterInitialized(opts *bind.FilterOpts) (*SplitInitializedIterator, error) {

	logs, sub, err := _Split.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SplitInitializedIterator{contract: _Split.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Split *SplitFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SplitInitialized) (event.Subscription, error) {

	logs, sub, err := _Split.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SplitInitialized)
				if err := _Split.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Split *SplitFilterer) ParseInitialized(log types.Log) (*SplitInitialized, error) {
	event := new(SplitInitialized)
	if err := _Split.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SplitPayeeAddedIterator is returned from FilterPayeeAdded and is used to iterate over the raw logs and unpacked data for PayeeAdded events raised by the Split contract.
type SplitPayeeAddedIterator struct {
	Event *SplitPayeeAdded // Event containing the contract specifics and raw log

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
func (it *SplitPayeeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SplitPayeeAdded)
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
		it.Event = new(SplitPayeeAdded)
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
func (it *SplitPayeeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SplitPayeeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SplitPayeeAdded represents a PayeeAdded event raised by the Split contract.
type SplitPayeeAdded struct {
	Account common.Address
	Shares  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPayeeAdded is a free log retrieval operation binding the contract event 0x40c340f65e17194d14ddddb073d3c9f888e3cb52b5aae0c6c7706b4fbc905fac.
//
// Solidity: event PayeeAdded(address account, uint256 shares)
func (_Split *SplitFilterer) FilterPayeeAdded(opts *bind.FilterOpts) (*SplitPayeeAddedIterator, error) {

	logs, sub, err := _Split.contract.FilterLogs(opts, "PayeeAdded")
	if err != nil {
		return nil, err
	}
	return &SplitPayeeAddedIterator{contract: _Split.contract, event: "PayeeAdded", logs: logs, sub: sub}, nil
}

// WatchPayeeAdded is a free log subscription operation binding the contract event 0x40c340f65e17194d14ddddb073d3c9f888e3cb52b5aae0c6c7706b4fbc905fac.
//
// Solidity: event PayeeAdded(address account, uint256 shares)
func (_Split *SplitFilterer) WatchPayeeAdded(opts *bind.WatchOpts, sink chan<- *SplitPayeeAdded) (event.Subscription, error) {

	logs, sub, err := _Split.contract.WatchLogs(opts, "PayeeAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SplitPayeeAdded)
				if err := _Split.contract.UnpackLog(event, "PayeeAdded", log); err != nil {
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

// ParsePayeeAdded is a log parse operation binding the contract event 0x40c340f65e17194d14ddddb073d3c9f888e3cb52b5aae0c6c7706b4fbc905fac.
//
// Solidity: event PayeeAdded(address account, uint256 shares)
func (_Split *SplitFilterer) ParsePayeeAdded(log types.Log) (*SplitPayeeAdded, error) {
	event := new(SplitPayeeAdded)
	if err := _Split.contract.UnpackLog(event, "PayeeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SplitPaymentReceivedIterator is returned from FilterPaymentReceived and is used to iterate over the raw logs and unpacked data for PaymentReceived events raised by the Split contract.
type SplitPaymentReceivedIterator struct {
	Event *SplitPaymentReceived // Event containing the contract specifics and raw log

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
func (it *SplitPaymentReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SplitPaymentReceived)
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
		it.Event = new(SplitPaymentReceived)
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
func (it *SplitPaymentReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SplitPaymentReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SplitPaymentReceived represents a PaymentReceived event raised by the Split contract.
type SplitPaymentReceived struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaymentReceived is a free log retrieval operation binding the contract event 0x6ef95f06320e7a25a04a175ca677b7052bdd97131872c2192525a629f51be770.
//
// Solidity: event PaymentReceived(address from, uint256 amount)
func (_Split *SplitFilterer) FilterPaymentReceived(opts *bind.FilterOpts) (*SplitPaymentReceivedIterator, error) {

	logs, sub, err := _Split.contract.FilterLogs(opts, "PaymentReceived")
	if err != nil {
		return nil, err
	}
	return &SplitPaymentReceivedIterator{contract: _Split.contract, event: "PaymentReceived", logs: logs, sub: sub}, nil
}

// WatchPaymentReceived is a free log subscription operation binding the contract event 0x6ef95f06320e7a25a04a175ca677b7052bdd97131872c2192525a629f51be770.
//
// Solidity: event PaymentReceived(address from, uint256 amount)
func (_Split *SplitFilterer) WatchPaymentReceived(opts *bind.WatchOpts, sink chan<- *SplitPaymentReceived) (event.Subscription, error) {

	logs, sub, err := _Split.contract.WatchLogs(opts, "PaymentReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SplitPaymentReceived)
				if err := _Split.contract.UnpackLog(event, "PaymentReceived", log); err != nil {
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

// ParsePaymentReceived is a log parse operation binding the contract event 0x6ef95f06320e7a25a04a175ca677b7052bdd97131872c2192525a629f51be770.
//
// Solidity: event PaymentReceived(address from, uint256 amount)
func (_Split *SplitFilterer) ParsePaymentReceived(log types.Log) (*SplitPaymentReceived, error) {
	event := new(SplitPaymentReceived)
	if err := _Split.contract.UnpackLog(event, "PaymentReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SplitPaymentReleasedIterator is returned from FilterPaymentReleased and is used to iterate over the raw logs and unpacked data for PaymentReleased events raised by the Split contract.
type SplitPaymentReleasedIterator struct {
	Event *SplitPaymentReleased // Event containing the contract specifics and raw log

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
func (it *SplitPaymentReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SplitPaymentReleased)
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
		it.Event = new(SplitPaymentReleased)
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
func (it *SplitPaymentReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SplitPaymentReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SplitPaymentReleased represents a PaymentReleased event raised by the Split contract.
type SplitPaymentReleased struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaymentReleased is a free log retrieval operation binding the contract event 0xdf20fd1e76bc69d672e4814fafb2c449bba3a5369d8359adf9e05e6fde87b056.
//
// Solidity: event PaymentReleased(address to, uint256 amount)
func (_Split *SplitFilterer) FilterPaymentReleased(opts *bind.FilterOpts) (*SplitPaymentReleasedIterator, error) {

	logs, sub, err := _Split.contract.FilterLogs(opts, "PaymentReleased")
	if err != nil {
		return nil, err
	}
	return &SplitPaymentReleasedIterator{contract: _Split.contract, event: "PaymentReleased", logs: logs, sub: sub}, nil
}

// WatchPaymentReleased is a free log subscription operation binding the contract event 0xdf20fd1e76bc69d672e4814fafb2c449bba3a5369d8359adf9e05e6fde87b056.
//
// Solidity: event PaymentReleased(address to, uint256 amount)
func (_Split *SplitFilterer) WatchPaymentReleased(opts *bind.WatchOpts, sink chan<- *SplitPaymentReleased) (event.Subscription, error) {

	logs, sub, err := _Split.contract.WatchLogs(opts, "PaymentReleased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SplitPaymentReleased)
				if err := _Split.contract.UnpackLog(event, "PaymentReleased", log); err != nil {
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

// ParsePaymentReleased is a log parse operation binding the contract event 0xdf20fd1e76bc69d672e4814fafb2c449bba3a5369d8359adf9e05e6fde87b056.
//
// Solidity: event PaymentReleased(address to, uint256 amount)
func (_Split *SplitFilterer) ParsePaymentReleased(log types.Log) (*SplitPaymentReleased, error) {
	event := new(SplitPaymentReleased)
	if err := _Split.contract.UnpackLog(event, "PaymentReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SplitRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Split contract.
type SplitRoleAdminChangedIterator struct {
	Event *SplitRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SplitRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SplitRoleAdminChanged)
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
		it.Event = new(SplitRoleAdminChanged)
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
func (it *SplitRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SplitRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SplitRoleAdminChanged represents a RoleAdminChanged event raised by the Split contract.
type SplitRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Split *SplitFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SplitRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Split.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SplitRoleAdminChangedIterator{contract: _Split.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Split *SplitFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SplitRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Split.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SplitRoleAdminChanged)
				if err := _Split.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Split *SplitFilterer) ParseRoleAdminChanged(log types.Log) (*SplitRoleAdminChanged, error) {
	event := new(SplitRoleAdminChanged)
	if err := _Split.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SplitRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Split contract.
type SplitRoleGrantedIterator struct {
	Event *SplitRoleGranted // Event containing the contract specifics and raw log

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
func (it *SplitRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SplitRoleGranted)
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
		it.Event = new(SplitRoleGranted)
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
func (it *SplitRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SplitRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SplitRoleGranted represents a RoleGranted event raised by the Split contract.
type SplitRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Split *SplitFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SplitRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Split.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SplitRoleGrantedIterator{contract: _Split.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Split *SplitFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SplitRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Split.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SplitRoleGranted)
				if err := _Split.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Split *SplitFilterer) ParseRoleGranted(log types.Log) (*SplitRoleGranted, error) {
	event := new(SplitRoleGranted)
	if err := _Split.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SplitRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Split contract.
type SplitRoleRevokedIterator struct {
	Event *SplitRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SplitRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SplitRoleRevoked)
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
		it.Event = new(SplitRoleRevoked)
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
func (it *SplitRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SplitRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SplitRoleRevoked represents a RoleRevoked event raised by the Split contract.
type SplitRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Split *SplitFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SplitRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Split.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SplitRoleRevokedIterator{contract: _Split.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Split *SplitFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SplitRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Split.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SplitRoleRevoked)
				if err := _Split.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Split *SplitFilterer) ParseRoleRevoked(log types.Log) (*SplitRoleRevoked, error) {
	event := new(SplitRoleRevoked)
	if err := _Split.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
