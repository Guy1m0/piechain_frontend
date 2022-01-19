// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth_lender

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

// ERC20MetaData contains all meta data concerning the ERC20 contract.
var ERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"loan\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"loan\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"loan\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b506040516106f63803806106f683398101604081905261002f916100e1565b610039338261003f565b50610120565b6001600160a01b0382166100995760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b80600260008282546100ab91906100fa565b90915550506001600160a01b038216600090815260208190526040812080548392906100d89084906100fa565b90915550505050565b6000602082840312156100f357600080fd5b5051919050565b6000821982111561011b57634e487b7160e01b600052601160045260246000fd5b500190565b6105c78061012f6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063095ea7b31461006757806318160ddd1461008f57806323b872dd146100a157806370a08231146100b4578063a9059cbb146100dd578063dd62ed3e146100f0575b600080fd5b61007a610075366004610491565b610129565b60405190151581526020015b60405180910390f35b6002545b604051908152602001610086565b61007a6100af3660046104bb565b61013f565b6100936100c23660046104f7565b6001600160a01b031660009081526020819052604090205490565b61007a6100eb366004610491565b6101f3565b6100936100fe366004610519565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6000610136338484610200565b50600192915050565b600061014c8484846102ef565b6001600160a01b0384166000908152600160209081526040808320338452909152902054828110156101d45760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e73666572206c6f616e206578636565647320616c6c6044820152656f77616e636560d01b60648201526084015b60405180910390fd5b6101e885336101e38685610562565b610200565b506001949350505050565b60006101363384846102ef565b6001600160a01b0383166102625760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016101cb565b6001600160a01b0382166102c35760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016101cb565b6001600160a01b0392831660009081526001602090815260408083209490951682529290925291902055565b6001600160a01b0383166103535760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016101cb565b6001600160a01b0382166103b55760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016101cb565b6001600160a01b0383166000908152602081905260409020548181101561042a5760405162461bcd60e51b8152602060048201526024808201527f45524332303a207472616e73666572206c6f616e20657863656564732062616c604482015263616e636560e01b60648201526084016101cb565b6104348282610562565b6001600160a01b03808616600090815260208190526040808220939093559085168152908120805484929061046a908490610579565b909155505050505050565b80356001600160a01b038116811461048c57600080fd5b919050565b600080604083850312156104a457600080fd5b6104ad83610475565b946020939093013593505050565b6000806000606084860312156104d057600080fd5b6104d984610475565b92506104e760208501610475565b9150604084013590509250925092565b60006020828403121561050957600080fd5b61051282610475565b9392505050565b6000806040838503121561052c57600080fd5b61053583610475565b915061054360208401610475565b90509250929050565b634e487b7160e01b600052601160045260246000fd5b6000828210156105745761057461054c565b500390565b6000821982111561058c5761058c61054c565b50019056fea2646970667358221220c0c67c71160b1051695817312afaa6ee997da5c311ac0d9df93963badef5f38964736f6c634300080b0033",
}

// ERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20MetaData.ABI instead.
var ERC20ABI = ERC20MetaData.ABI

// Deprecated: Use ERC20MetaData.Sigs instead.
// ERC20FuncSigs maps the 4-byte function signature to its string representation.
var ERC20FuncSigs = ERC20MetaData.Sigs

// ERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20MetaData.Bin instead.
var ERC20Bin = ERC20MetaData.Bin

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := ERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20Bin), backend, initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 loan) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, loan *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, loan)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 loan) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, loan *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, loan)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 loan) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, loan *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, loan)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 loan) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, loan *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", recipient, loan)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 loan) returns(bool)
func (_ERC20 *ERC20Session) Transfer(recipient common.Address, loan *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, loan)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 loan) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(recipient common.Address, loan *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, loan)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 loan) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, loan *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", sender, recipient, loan)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 loan) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(sender common.Address, recipient common.Address, loan *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, loan)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 loan) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, loan *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, loan)
}

// LenderMetaData contains all meta data concerning the Lender contract.
var LenderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"arbitrage\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"endLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"vl\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"rl\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sl\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"va\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"ra\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sa\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"intrest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lender\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"loan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"loanHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"lender_\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"arbitrage_\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"exchange_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"loan_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"intrest_\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"loanHash_\",\"type\":\"bytes32\"}],\"name\":\"setup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"69c8d338": "arbitrage()",
		"fbd3f3ce": "endLoan(bool)",
		"d2f7265a": "exchange()",
		"fa5f4034": "initialize(uint8,bytes32,bytes32,uint8,bytes32,bytes32)",
		"ffbe7402": "intrest()",
		"bcead63e": "lender()",
		"d285b7b4": "loan()",
		"1e306772": "loanHash()",
		"566b7de8": "setup(address,address,address,uint256,uint256,bytes32)",
		"200d2ed2": "status()",
		"9d76ea58": "tokenAddress()",
	},
	Bin: "0x608060405234801561001057600080fd5b506040516106c23803806106c283398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b61062f806100936000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063bcead63e11610071578063bcead63e14610170578063d285b7b414610183578063d2f7265a1461018c578063fa5f40341461019f578063fbd3f3ce146101b2578063ffbe7402146101c557600080fd5b80631e306772146100ae578063200d2ed2146100ca578063566b7de8146100d357806369c8d338146101325780639d76ea581461015d575b600080fd5b6100b760065481565b6040519081526020015b60405180910390f35b6100b760075481565b6101306100e13660046104b9565b600180546001600160a01b03199081166001600160a01b039889161790915560028054821696881696909617909555600380549095169390951692909217909255600491909155600555600655565b005b600254610145906001600160a01b031681565b6040516001600160a01b0390911681526020016100c1565b600054610145906001600160a01b031681565b600154610145906001600160a01b031681565b6100b760045481565b600354610145906001600160a01b031681565b6101306101ad366004610529565b6101ce565b6101306101c0366004610592565b6103ee565b6100b760055481565b6006546040805160008152602081018083529290925260ff881690820152606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610222573d6000803e3d6000fd5b5050604051601f1901516001546001600160a01b0390811691161490506102905760405162461bcd60e51b815260206004820152601860248201527f496e76616c6964204c656e646572205369676e6174757265000000000000000060448201526064015b60405180910390fd5b6006546040805160008152602081018083529290925260ff851690820152606081018390526080810182905260019060a0016020604051602081039080840390855afa1580156102e4573d6000803e3d6000fd5b5050604051601f1901516002546001600160a01b03908116911614905061034d5760405162461bcd60e51b815260206004820152601b60248201527f496e76616c696420417262697472616765205369676e617475726500000000006044820152606401610287565b600054600154600354600480546040516323b872dd60e01b81526001600160a01b03948516928101929092529183166024820152604481019190915291169081906323b872dd906064016020604051808303816000875af11580156103b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103da91906105b6565b156103e55760016007555b50505050505050565b60045481156104075760055461040490826105d3565b90505b6000546003546001546040516323b872dd60e01b81526001600160a01b03928316600482015290821660248201526044810184905291169081906323b872dd906064016020604051808303816000875af1158015610469573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061048d91906105b6565b156104985760026007555b505050565b80356001600160a01b03811681146104b457600080fd5b919050565b60008060008060008060c087890312156104d257600080fd5b6104db8761049d565b95506104e96020880161049d565b94506104f76040880161049d565b9350606087013592506080870135915060a087013590509295509295509295565b803560ff811681146104b457600080fd5b60008060008060008060c0878903121561054257600080fd5b61054b87610518565b9550602087013594506040870135935061056760608801610518565b92506080870135915060a087013590509295509295509295565b801515811461058f57600080fd5b50565b6000602082840312156105a457600080fd5b81356105af81610581565b9392505050565b6000602082840312156105c857600080fd5b81516105af81610581565b600082198211156105f457634e487b7160e01b600052601160045260246000fd5b50019056fea2646970667358221220ef34e221d768bb84dadbd96a98a50f016eaf57e2d93deed1b74a56c887439ced64736f6c634300080b0033",
}

// LenderABI is the input ABI used to generate the binding from.
// Deprecated: Use LenderMetaData.ABI instead.
var LenderABI = LenderMetaData.ABI

// Deprecated: Use LenderMetaData.Sigs instead.
// LenderFuncSigs maps the 4-byte function signature to its string representation.
var LenderFuncSigs = LenderMetaData.Sigs

// LenderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LenderMetaData.Bin instead.
var LenderBin = LenderMetaData.Bin

// DeployLender deploys a new Ethereum contract, binding an instance of Lender to it.
func DeployLender(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAddress_ common.Address) (common.Address, *types.Transaction, *Lender, error) {
	parsed, err := LenderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LenderBin), backend, tokenAddress_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Lender{LenderCaller: LenderCaller{contract: contract}, LenderTransactor: LenderTransactor{contract: contract}, LenderFilterer: LenderFilterer{contract: contract}}, nil
}

// Lender is an auto generated Go binding around an Ethereum contract.
type Lender struct {
	LenderCaller     // Read-only binding to the contract
	LenderTransactor // Write-only binding to the contract
	LenderFilterer   // Log filterer for contract events
}

// LenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type LenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LenderSession struct {
	Contract     *Lender           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LenderCallerSession struct {
	Contract *LenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LenderTransactorSession struct {
	Contract     *LenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type LenderRaw struct {
	Contract *Lender // Generic contract binding to access the raw methods on
}

// LenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LenderCallerRaw struct {
	Contract *LenderCaller // Generic read-only contract binding to access the raw methods on
}

// LenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LenderTransactorRaw struct {
	Contract *LenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLender creates a new instance of Lender, bound to a specific deployed contract.
func NewLender(address common.Address, backend bind.ContractBackend) (*Lender, error) {
	contract, err := bindLender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lender{LenderCaller: LenderCaller{contract: contract}, LenderTransactor: LenderTransactor{contract: contract}, LenderFilterer: LenderFilterer{contract: contract}}, nil
}

// NewLenderCaller creates a new read-only instance of Lender, bound to a specific deployed contract.
func NewLenderCaller(address common.Address, caller bind.ContractCaller) (*LenderCaller, error) {
	contract, err := bindLender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LenderCaller{contract: contract}, nil
}

// NewLenderTransactor creates a new write-only instance of Lender, bound to a specific deployed contract.
func NewLenderTransactor(address common.Address, transactor bind.ContractTransactor) (*LenderTransactor, error) {
	contract, err := bindLender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LenderTransactor{contract: contract}, nil
}

// NewLenderFilterer creates a new log filterer instance of Lender, bound to a specific deployed contract.
func NewLenderFilterer(address common.Address, filterer bind.ContractFilterer) (*LenderFilterer, error) {
	contract, err := bindLender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LenderFilterer{contract: contract}, nil
}

// bindLender binds a generic wrapper to an already deployed contract.
func bindLender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LenderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lender *LenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lender.Contract.LenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lender *LenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lender.Contract.LenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lender *LenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lender.Contract.LenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lender *LenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lender *LenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lender *LenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lender.Contract.contract.Transact(opts, method, params...)
}

// Arbitrage is a free data retrieval call binding the contract method 0x69c8d338.
//
// Solidity: function arbitrage() view returns(address)
func (_Lender *LenderCaller) Arbitrage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lender.contract.Call(opts, &out, "arbitrage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Arbitrage is a free data retrieval call binding the contract method 0x69c8d338.
//
// Solidity: function arbitrage() view returns(address)
func (_Lender *LenderSession) Arbitrage() (common.Address, error) {
	return _Lender.Contract.Arbitrage(&_Lender.CallOpts)
}

// Arbitrage is a free data retrieval call binding the contract method 0x69c8d338.
//
// Solidity: function arbitrage() view returns(address)
func (_Lender *LenderCallerSession) Arbitrage() (common.Address, error) {
	return _Lender.Contract.Arbitrage(&_Lender.CallOpts)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_Lender *LenderCaller) Exchange(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lender.contract.Call(opts, &out, "exchange")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_Lender *LenderSession) Exchange() (common.Address, error) {
	return _Lender.Contract.Exchange(&_Lender.CallOpts)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_Lender *LenderCallerSession) Exchange() (common.Address, error) {
	return _Lender.Contract.Exchange(&_Lender.CallOpts)
}

// Intrest is a free data retrieval call binding the contract method 0xffbe7402.
//
// Solidity: function intrest() view returns(uint256)
func (_Lender *LenderCaller) Intrest(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lender.contract.Call(opts, &out, "intrest")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Intrest is a free data retrieval call binding the contract method 0xffbe7402.
//
// Solidity: function intrest() view returns(uint256)
func (_Lender *LenderSession) Intrest() (*big.Int, error) {
	return _Lender.Contract.Intrest(&_Lender.CallOpts)
}

// Intrest is a free data retrieval call binding the contract method 0xffbe7402.
//
// Solidity: function intrest() view returns(uint256)
func (_Lender *LenderCallerSession) Intrest() (*big.Int, error) {
	return _Lender.Contract.Intrest(&_Lender.CallOpts)
}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() view returns(address)
func (_Lender *LenderCaller) Lender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lender.contract.Call(opts, &out, "lender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() view returns(address)
func (_Lender *LenderSession) Lender() (common.Address, error) {
	return _Lender.Contract.Lender(&_Lender.CallOpts)
}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() view returns(address)
func (_Lender *LenderCallerSession) Lender() (common.Address, error) {
	return _Lender.Contract.Lender(&_Lender.CallOpts)
}

// Loan is a free data retrieval call binding the contract method 0xd285b7b4.
//
// Solidity: function loan() view returns(uint256)
func (_Lender *LenderCaller) Loan(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lender.contract.Call(opts, &out, "loan")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Loan is a free data retrieval call binding the contract method 0xd285b7b4.
//
// Solidity: function loan() view returns(uint256)
func (_Lender *LenderSession) Loan() (*big.Int, error) {
	return _Lender.Contract.Loan(&_Lender.CallOpts)
}

// Loan is a free data retrieval call binding the contract method 0xd285b7b4.
//
// Solidity: function loan() view returns(uint256)
func (_Lender *LenderCallerSession) Loan() (*big.Int, error) {
	return _Lender.Contract.Loan(&_Lender.CallOpts)
}

// LoanHash is a free data retrieval call binding the contract method 0x1e306772.
//
// Solidity: function loanHash() view returns(bytes32)
func (_Lender *LenderCaller) LoanHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lender.contract.Call(opts, &out, "loanHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LoanHash is a free data retrieval call binding the contract method 0x1e306772.
//
// Solidity: function loanHash() view returns(bytes32)
func (_Lender *LenderSession) LoanHash() ([32]byte, error) {
	return _Lender.Contract.LoanHash(&_Lender.CallOpts)
}

// LoanHash is a free data retrieval call binding the contract method 0x1e306772.
//
// Solidity: function loanHash() view returns(bytes32)
func (_Lender *LenderCallerSession) LoanHash() ([32]byte, error) {
	return _Lender.Contract.LoanHash(&_Lender.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint256)
func (_Lender *LenderCaller) Status(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lender.contract.Call(opts, &out, "status")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint256)
func (_Lender *LenderSession) Status() (*big.Int, error) {
	return _Lender.Contract.Status(&_Lender.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint256)
func (_Lender *LenderCallerSession) Status() (*big.Int, error) {
	return _Lender.Contract.Status(&_Lender.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_Lender *LenderCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lender.contract.Call(opts, &out, "tokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_Lender *LenderSession) TokenAddress() (common.Address, error) {
	return _Lender.Contract.TokenAddress(&_Lender.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_Lender *LenderCallerSession) TokenAddress() (common.Address, error) {
	return _Lender.Contract.TokenAddress(&_Lender.CallOpts)
}

// EndLoan is a paid mutator transaction binding the contract method 0xfbd3f3ce.
//
// Solidity: function endLoan(bool success) returns()
func (_Lender *LenderTransactor) EndLoan(opts *bind.TransactOpts, success bool) (*types.Transaction, error) {
	return _Lender.contract.Transact(opts, "endLoan", success)
}

// EndLoan is a paid mutator transaction binding the contract method 0xfbd3f3ce.
//
// Solidity: function endLoan(bool success) returns()
func (_Lender *LenderSession) EndLoan(success bool) (*types.Transaction, error) {
	return _Lender.Contract.EndLoan(&_Lender.TransactOpts, success)
}

// EndLoan is a paid mutator transaction binding the contract method 0xfbd3f3ce.
//
// Solidity: function endLoan(bool success) returns()
func (_Lender *LenderTransactorSession) EndLoan(success bool) (*types.Transaction, error) {
	return _Lender.Contract.EndLoan(&_Lender.TransactOpts, success)
}

// Initialize is a paid mutator transaction binding the contract method 0xfa5f4034.
//
// Solidity: function initialize(uint8 vl, bytes32 rl, bytes32 sl, uint8 va, bytes32 ra, bytes32 sa) returns()
func (_Lender *LenderTransactor) Initialize(opts *bind.TransactOpts, vl uint8, rl [32]byte, sl [32]byte, va uint8, ra [32]byte, sa [32]byte) (*types.Transaction, error) {
	return _Lender.contract.Transact(opts, "initialize", vl, rl, sl, va, ra, sa)
}

// Initialize is a paid mutator transaction binding the contract method 0xfa5f4034.
//
// Solidity: function initialize(uint8 vl, bytes32 rl, bytes32 sl, uint8 va, bytes32 ra, bytes32 sa) returns()
func (_Lender *LenderSession) Initialize(vl uint8, rl [32]byte, sl [32]byte, va uint8, ra [32]byte, sa [32]byte) (*types.Transaction, error) {
	return _Lender.Contract.Initialize(&_Lender.TransactOpts, vl, rl, sl, va, ra, sa)
}

// Initialize is a paid mutator transaction binding the contract method 0xfa5f4034.
//
// Solidity: function initialize(uint8 vl, bytes32 rl, bytes32 sl, uint8 va, bytes32 ra, bytes32 sa) returns()
func (_Lender *LenderTransactorSession) Initialize(vl uint8, rl [32]byte, sl [32]byte, va uint8, ra [32]byte, sa [32]byte) (*types.Transaction, error) {
	return _Lender.Contract.Initialize(&_Lender.TransactOpts, vl, rl, sl, va, ra, sa)
}

// Setup is a paid mutator transaction binding the contract method 0x566b7de8.
//
// Solidity: function setup(address lender_, address arbitrage_, address exchange_, uint256 loan_, uint256 intrest_, bytes32 loanHash_) returns()
func (_Lender *LenderTransactor) Setup(opts *bind.TransactOpts, lender_ common.Address, arbitrage_ common.Address, exchange_ common.Address, loan_ *big.Int, intrest_ *big.Int, loanHash_ [32]byte) (*types.Transaction, error) {
	return _Lender.contract.Transact(opts, "setup", lender_, arbitrage_, exchange_, loan_, intrest_, loanHash_)
}

// Setup is a paid mutator transaction binding the contract method 0x566b7de8.
//
// Solidity: function setup(address lender_, address arbitrage_, address exchange_, uint256 loan_, uint256 intrest_, bytes32 loanHash_) returns()
func (_Lender *LenderSession) Setup(lender_ common.Address, arbitrage_ common.Address, exchange_ common.Address, loan_ *big.Int, intrest_ *big.Int, loanHash_ [32]byte) (*types.Transaction, error) {
	return _Lender.Contract.Setup(&_Lender.TransactOpts, lender_, arbitrage_, exchange_, loan_, intrest_, loanHash_)
}

// Setup is a paid mutator transaction binding the contract method 0x566b7de8.
//
// Solidity: function setup(address lender_, address arbitrage_, address exchange_, uint256 loan_, uint256 intrest_, bytes32 loanHash_) returns()
func (_Lender *LenderTransactorSession) Setup(lender_ common.Address, arbitrage_ common.Address, exchange_ common.Address, loan_ *big.Int, intrest_ *big.Int, loanHash_ [32]byte) (*types.Transaction, error) {
	return _Lender.Contract.Setup(&_Lender.TransactOpts, lender_, arbitrage_, exchange_, loan_, intrest_, loanHash_)
}
