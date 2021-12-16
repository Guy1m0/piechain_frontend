// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth_lender

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

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"loan\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"loan\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"loan\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var ERC20ParsedABI, _ = abi.JSON(strings.NewReader(ERC20ABI))

// ERC20FuncSigs maps the 4-byte function signature to its string representation.
var ERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20Bin is the compiled bytecode used for deploying new contracts.
var ERC20Bin = "0x608060405234801561001057600080fd5b506040516106283803806106288339818101604052602081101561003357600080fd5b505161004833826001600160e01b0361004e16565b506100d3565b6001600160a01b0382166100a9576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b60028054820190556001600160a01b03909116600090815260208190526040902080549091019055565b610546806100e26000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063095ea7b31461006757806318160ddd146100a757806323b872dd146100c157806370a08231146100f7578063a9059cbb1461011d578063dd62ed3e14610149575b600080fd5b6100936004803603604081101561007d57600080fd5b506001600160a01b038135169060200135610177565b604080519115158252519081900360200190f35b6100af61018d565b60408051918252519081900360200190f35b610093600480360360608110156100d757600080fd5b506001600160a01b03813581169160208101359091169060400135610193565b6100af6004803603602081101561010d57600080fd5b50356001600160a01b031661021b565b6100936004803603604081101561013357600080fd5b506001600160a01b038135169060200135610236565b6100af6004803603604081101561015f57600080fd5b506001600160a01b0381358116916020013516610243565b600061018433848461026e565b50600192915050565b60025490565b60006101a0848484610324565b6001600160a01b0384166000908152600160209081526040808320338452909152902054828110156102035760405162461bcd60e51b815260040180806020018281038252602681526020018061047f6026913960400191505060405180910390fd5b610210853385840361026e565b506001949350505050565b6001600160a01b031660009081526020819052604090205490565b6000610184338484610324565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b0383166102b35760405162461bcd60e51b81526004018080602001828103825260248152602001806104ca6024913960400191505060405180910390fd5b6001600160a01b0382166102f85760405162461bcd60e51b815260040180806020018281038252602281526020018061045d6022913960400191505060405180910390fd5b6001600160a01b0392831660009081526001602090815260408083209490951682529290925291902055565b6001600160a01b0383166103695760405162461bcd60e51b81526004018080602001828103825260258152602001806104a56025913960400191505060405180910390fd5b6001600160a01b0382166103ae5760405162461bcd60e51b815260040180806020018281038252602381526020018061043a6023913960400191505060405180910390fd5b6001600160a01b038316600090815260208190526040902054818110156104065760405162461bcd60e51b81526004018080602001828103825260248152602001806104ee6024913960400191505060405180910390fd5b6001600160a01b03938416600090815260208190526040808220928490039092559290931682529190208054909101905556fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e73666572206c6f616e206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a207472616e73666572206c6f616e20657863656564732062616c616e6365a265627a7a72315820350f251c29b34e90f10fd47153c76e9924aa15b1c3afbf81e60d335ab41bb6ee64736f6c63430005110032"

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend, initialSupply)
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
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
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

// LenderABI is the input ABI used to generate the binding from.
const LenderABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"arbitrage\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"endLoan\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"vl\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"rl\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sl\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"va\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"ra\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sa\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"intrest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lender\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loanHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"lender_\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"arbitrage_\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"exchange_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"loan_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"intrest_\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"loanHash_\",\"type\":\"bytes32\"}],\"name\":\"setup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

var LenderParsedABI, _ = abi.JSON(strings.NewReader(LenderABI))

// LenderFuncSigs maps the 4-byte function signature to its string representation.
var LenderFuncSigs = map[string]string{
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
}

// LenderBin is the compiled bytecode used for deploying new contracts.
var LenderBin = "0x608060405234801561001057600080fd5b5060405161045f38038061045f8339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b03199092169190911790556103fa806100656000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063bcead63e11610071578063bcead63e14610146578063d285b7b41461014e578063d2f7265a14610156578063fa5f40341461015e578063fbd3f3ce146101a0578063ffbe7402146101bf576100a9565b80631e306772146100ae578063200d2ed2146100c8578063566b7de8146100d057806369c8d3381461011a5780639d76ea581461013e575b600080fd5b6100b66101c7565b60408051918252519081900360200190f35b6100b66101cd565b610118600480360360c08110156100e657600080fd5b506001600160a01b03813581169160208101358216916040820135169060608101359060808101359060a001356101d3565b005b610122610222565b604080516001600160a01b039092168252519081900360200190f35b610122610231565b610122610240565b6100b661024f565b610122610255565b610118600480360360c081101561017457600080fd5b5060ff813581169160208101359160408201359160608101359091169060808101359060a00135610264565b610118600480360360208110156101b657600080fd5b5035151561030f565b6100b66103bf565b60065481565b60075481565b600180546001600160a01b03199081166001600160a01b039889161790915560028054821696881696909617909555600380549095169390951692909217909255600491909155600555600655565b6002546001600160a01b031681565b6000546001600160a01b031681565b6001546001600160a01b031681565b60045481565b6003546001600160a01b031681565b6000805460015460035460048054604080516323b872dd60e01b81526001600160a01b0395861693810193909352928416602483015260448201529051919092169283926323b872dd9260648083019360209383900390910190829087803b1580156102cf57600080fd5b505af11580156102e3573d6000803e3d6000fd5b505050506040513d60208110156102f957600080fd5b5051156103065760016007555b50505050505050565b600454811561031d57600554015b60008054600354600154604080516323b872dd60e01b81526001600160a01b03938416600482015291831660248301526044820186905251919092169283926323b872dd9260648083019360209383900390910190829087803b15801561038357600080fd5b505af1158015610397573d6000803e3d6000fd5b505050506040513d60208110156103ad57600080fd5b5051156103ba5760026007555b505050565b6005548156fea265627a7a72315820fcdf5e6a108cc1c84b63d3810ca5167575954e19057c5f9ffd0bf558097f3a6064736f6c63430005110032"

// DeployLender deploys a new Ethereum contract, binding an instance of Lender to it.
func DeployLender(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAddress_ common.Address) (common.Address, *types.Transaction, *Lender, error) {
	parsed, err := abi.JSON(strings.NewReader(LenderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LenderBin), backend, tokenAddress_)
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
func (_Lender *LenderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_Lender *LenderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
// Solidity: function arbitrage() constant returns(address)
func (_Lender *LenderCaller) Arbitrage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Lender.contract.Call(opts, out, "arbitrage")
	return *ret0, err
}

// Arbitrage is a free data retrieval call binding the contract method 0x69c8d338.
//
// Solidity: function arbitrage() constant returns(address)
func (_Lender *LenderSession) Arbitrage() (common.Address, error) {
	return _Lender.Contract.Arbitrage(&_Lender.CallOpts)
}

// Arbitrage is a free data retrieval call binding the contract method 0x69c8d338.
//
// Solidity: function arbitrage() constant returns(address)
func (_Lender *LenderCallerSession) Arbitrage() (common.Address, error) {
	return _Lender.Contract.Arbitrage(&_Lender.CallOpts)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() constant returns(address)
func (_Lender *LenderCaller) Exchange(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Lender.contract.Call(opts, out, "exchange")
	return *ret0, err
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() constant returns(address)
func (_Lender *LenderSession) Exchange() (common.Address, error) {
	return _Lender.Contract.Exchange(&_Lender.CallOpts)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() constant returns(address)
func (_Lender *LenderCallerSession) Exchange() (common.Address, error) {
	return _Lender.Contract.Exchange(&_Lender.CallOpts)
}

// Intrest is a free data retrieval call binding the contract method 0xffbe7402.
//
// Solidity: function intrest() constant returns(uint256)
func (_Lender *LenderCaller) Intrest(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Lender.contract.Call(opts, out, "intrest")
	return *ret0, err
}

// Intrest is a free data retrieval call binding the contract method 0xffbe7402.
//
// Solidity: function intrest() constant returns(uint256)
func (_Lender *LenderSession) Intrest() (*big.Int, error) {
	return _Lender.Contract.Intrest(&_Lender.CallOpts)
}

// Intrest is a free data retrieval call binding the contract method 0xffbe7402.
//
// Solidity: function intrest() constant returns(uint256)
func (_Lender *LenderCallerSession) Intrest() (*big.Int, error) {
	return _Lender.Contract.Intrest(&_Lender.CallOpts)
}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() constant returns(address)
func (_Lender *LenderCaller) Lender(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Lender.contract.Call(opts, out, "lender")
	return *ret0, err
}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() constant returns(address)
func (_Lender *LenderSession) Lender() (common.Address, error) {
	return _Lender.Contract.Lender(&_Lender.CallOpts)
}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() constant returns(address)
func (_Lender *LenderCallerSession) Lender() (common.Address, error) {
	return _Lender.Contract.Lender(&_Lender.CallOpts)
}

// Loan is a free data retrieval call binding the contract method 0xd285b7b4.
//
// Solidity: function loan() constant returns(uint256)
func (_Lender *LenderCaller) Loan(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Lender.contract.Call(opts, out, "loan")
	return *ret0, err
}

// Loan is a free data retrieval call binding the contract method 0xd285b7b4.
//
// Solidity: function loan() constant returns(uint256)
func (_Lender *LenderSession) Loan() (*big.Int, error) {
	return _Lender.Contract.Loan(&_Lender.CallOpts)
}

// Loan is a free data retrieval call binding the contract method 0xd285b7b4.
//
// Solidity: function loan() constant returns(uint256)
func (_Lender *LenderCallerSession) Loan() (*big.Int, error) {
	return _Lender.Contract.Loan(&_Lender.CallOpts)
}

// LoanHash is a free data retrieval call binding the contract method 0x1e306772.
//
// Solidity: function loanHash() constant returns(bytes32)
func (_Lender *LenderCaller) LoanHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Lender.contract.Call(opts, out, "loanHash")
	return *ret0, err
}

// LoanHash is a free data retrieval call binding the contract method 0x1e306772.
//
// Solidity: function loanHash() constant returns(bytes32)
func (_Lender *LenderSession) LoanHash() ([32]byte, error) {
	return _Lender.Contract.LoanHash(&_Lender.CallOpts)
}

// LoanHash is a free data retrieval call binding the contract method 0x1e306772.
//
// Solidity: function loanHash() constant returns(bytes32)
func (_Lender *LenderCallerSession) LoanHash() ([32]byte, error) {
	return _Lender.Contract.LoanHash(&_Lender.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(uint256)
func (_Lender *LenderCaller) Status(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Lender.contract.Call(opts, out, "status")
	return *ret0, err
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(uint256)
func (_Lender *LenderSession) Status() (*big.Int, error) {
	return _Lender.Contract.Status(&_Lender.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(uint256)
func (_Lender *LenderCallerSession) Status() (*big.Int, error) {
	return _Lender.Contract.Status(&_Lender.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_Lender *LenderCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Lender.contract.Call(opts, out, "tokenAddress")
	return *ret0, err
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_Lender *LenderSession) TokenAddress() (common.Address, error) {
	return _Lender.Contract.TokenAddress(&_Lender.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
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
