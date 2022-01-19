// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth_arbitrage

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

// AMMMetaData contains all meta data concerning the AMM contract.
var AMMMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token1_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token2_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"exchange1to2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"exchange2to1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate1_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate2_\",\"type\":\"uint256\"}],\"name\":\"setRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token2Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3ede9696": "exchange1to2(uint256)",
		"d030f9e5": "exchange2to1(uint256)",
		"cf854969": "rate1()",
		"f555b815": "rate2()",
		"46df2ccb": "setRate(uint256,uint256)",
		"d7cb416f": "token1Address()",
		"2456af82": "token2Address()",
	},
	Bin: "0x608060405234801561001057600080fd5b506040516104d83803806104d883398101604081905261002f9161007c565b600280546001600160a01b039384166001600160a01b031991821617909155600380549290931691161790556100af565b80516001600160a01b038116811461007757600080fd5b919050565b6000806040838503121561008f57600080fd5b61009883610060565b91506100a660208401610060565b90509250929050565b61041a806100be6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063cf8549691161005b578063cf854969146100ee578063d030f9e5146100f7578063d7cb416f1461010a578063f555b8151461011d57600080fd5b80632456af82146100825780633ede9696146100b257806346df2ccb146100d3575b600080fd5b600354610095906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6100c56100c0366004610331565b610126565b6040519081526020016100a9565b6100ec6100e136600461034a565b600091909155600155565b005b6100c560005481565b6100c5610105366004610331565b610250565b600254610095906001600160a01b031681565b6100c560015481565b6002546003546000805460015491936001600160a01b03908116931691849190610150908761036c565b61015a9190610399565b6040516323b872dd60e01b8152336004820152306024820152604481018790529091506001600160a01b038416906323b872dd906064016020604051808303816000875af11580156101b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101d491906103bb565b5060405163a9059cbb60e01b8152336004820152602481018290526001600160a01b0383169063a9059cbb906044015b6020604051808303816000875af1158015610223573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061024791906103bb565b50949350505050565b6002546003546001546000805490936001600160a01b039081169316918491610279908761036c565b6102839190610399565b6040516323b872dd60e01b8152336004820152306024820152604481018790529091506001600160a01b038316906323b872dd906064016020604051808303816000875af11580156102d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102fd91906103bb565b5060405163a9059cbb60e01b8152336004820152602481018290526001600160a01b0384169063a9059cbb90604401610204565b60006020828403121561034357600080fd5b5035919050565b6000806040838503121561035d57600080fd5b50508035926020909101359150565b600081600019048311821515161561039457634e487b7160e01b600052601160045260246000fd5b500290565b6000826103b657634e487b7160e01b600052601260045260246000fd5b500490565b6000602082840312156103cd57600080fd5b815180151581146103dd57600080fd5b939250505056fea26469706673582212206f36fb0d3de5ff99b82982171ff332c898aac1463e2ac85a9e5c43b2750cb6a164736f6c634300080b0033",
}

// AMMABI is the input ABI used to generate the binding from.
// Deprecated: Use AMMMetaData.ABI instead.
var AMMABI = AMMMetaData.ABI

// Deprecated: Use AMMMetaData.Sigs instead.
// AMMFuncSigs maps the 4-byte function signature to its string representation.
var AMMFuncSigs = AMMMetaData.Sigs

// AMMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AMMMetaData.Bin instead.
var AMMBin = AMMMetaData.Bin

// DeployAMM deploys a new Ethereum contract, binding an instance of AMM to it.
func DeployAMM(auth *bind.TransactOpts, backend bind.ContractBackend, token1_ common.Address, token2_ common.Address) (common.Address, *types.Transaction, *AMM, error) {
	parsed, err := AMMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AMMBin), backend, token1_, token2_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AMM{AMMCaller: AMMCaller{contract: contract}, AMMTransactor: AMMTransactor{contract: contract}, AMMFilterer: AMMFilterer{contract: contract}}, nil
}

// AMM is an auto generated Go binding around an Ethereum contract.
type AMM struct {
	AMMCaller     // Read-only binding to the contract
	AMMTransactor // Write-only binding to the contract
	AMMFilterer   // Log filterer for contract events
}

// AMMCaller is an auto generated read-only Go binding around an Ethereum contract.
type AMMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AMMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AMMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AMMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AMMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AMMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AMMSession struct {
	Contract     *AMM              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AMMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AMMCallerSession struct {
	Contract *AMMCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AMMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AMMTransactorSession struct {
	Contract     *AMMTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AMMRaw is an auto generated low-level Go binding around an Ethereum contract.
type AMMRaw struct {
	Contract *AMM // Generic contract binding to access the raw methods on
}

// AMMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AMMCallerRaw struct {
	Contract *AMMCaller // Generic read-only contract binding to access the raw methods on
}

// AMMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AMMTransactorRaw struct {
	Contract *AMMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAMM creates a new instance of AMM, bound to a specific deployed contract.
func NewAMM(address common.Address, backend bind.ContractBackend) (*AMM, error) {
	contract, err := bindAMM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AMM{AMMCaller: AMMCaller{contract: contract}, AMMTransactor: AMMTransactor{contract: contract}, AMMFilterer: AMMFilterer{contract: contract}}, nil
}

// NewAMMCaller creates a new read-only instance of AMM, bound to a specific deployed contract.
func NewAMMCaller(address common.Address, caller bind.ContractCaller) (*AMMCaller, error) {
	contract, err := bindAMM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AMMCaller{contract: contract}, nil
}

// NewAMMTransactor creates a new write-only instance of AMM, bound to a specific deployed contract.
func NewAMMTransactor(address common.Address, transactor bind.ContractTransactor) (*AMMTransactor, error) {
	contract, err := bindAMM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AMMTransactor{contract: contract}, nil
}

// NewAMMFilterer creates a new log filterer instance of AMM, bound to a specific deployed contract.
func NewAMMFilterer(address common.Address, filterer bind.ContractFilterer) (*AMMFilterer, error) {
	contract, err := bindAMM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AMMFilterer{contract: contract}, nil
}

// bindAMM binds a generic wrapper to an already deployed contract.
func bindAMM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AMMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AMM *AMMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AMM.Contract.AMMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AMM *AMMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AMM.Contract.AMMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AMM *AMMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AMM.Contract.AMMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AMM *AMMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AMM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AMM *AMMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AMM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AMM *AMMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AMM.Contract.contract.Transact(opts, method, params...)
}

// Rate1 is a free data retrieval call binding the contract method 0xcf854969.
//
// Solidity: function rate1() view returns(uint256)
func (_AMM *AMMCaller) Rate1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "rate1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate1 is a free data retrieval call binding the contract method 0xcf854969.
//
// Solidity: function rate1() view returns(uint256)
func (_AMM *AMMSession) Rate1() (*big.Int, error) {
	return _AMM.Contract.Rate1(&_AMM.CallOpts)
}

// Rate1 is a free data retrieval call binding the contract method 0xcf854969.
//
// Solidity: function rate1() view returns(uint256)
func (_AMM *AMMCallerSession) Rate1() (*big.Int, error) {
	return _AMM.Contract.Rate1(&_AMM.CallOpts)
}

// Rate2 is a free data retrieval call binding the contract method 0xf555b815.
//
// Solidity: function rate2() view returns(uint256)
func (_AMM *AMMCaller) Rate2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "rate2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate2 is a free data retrieval call binding the contract method 0xf555b815.
//
// Solidity: function rate2() view returns(uint256)
func (_AMM *AMMSession) Rate2() (*big.Int, error) {
	return _AMM.Contract.Rate2(&_AMM.CallOpts)
}

// Rate2 is a free data retrieval call binding the contract method 0xf555b815.
//
// Solidity: function rate2() view returns(uint256)
func (_AMM *AMMCallerSession) Rate2() (*big.Int, error) {
	return _AMM.Contract.Rate2(&_AMM.CallOpts)
}

// Token1Address is a free data retrieval call binding the contract method 0xd7cb416f.
//
// Solidity: function token1Address() view returns(address)
func (_AMM *AMMCaller) Token1Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "token1Address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1Address is a free data retrieval call binding the contract method 0xd7cb416f.
//
// Solidity: function token1Address() view returns(address)
func (_AMM *AMMSession) Token1Address() (common.Address, error) {
	return _AMM.Contract.Token1Address(&_AMM.CallOpts)
}

// Token1Address is a free data retrieval call binding the contract method 0xd7cb416f.
//
// Solidity: function token1Address() view returns(address)
func (_AMM *AMMCallerSession) Token1Address() (common.Address, error) {
	return _AMM.Contract.Token1Address(&_AMM.CallOpts)
}

// Token2Address is a free data retrieval call binding the contract method 0x2456af82.
//
// Solidity: function token2Address() view returns(address)
func (_AMM *AMMCaller) Token2Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "token2Address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token2Address is a free data retrieval call binding the contract method 0x2456af82.
//
// Solidity: function token2Address() view returns(address)
func (_AMM *AMMSession) Token2Address() (common.Address, error) {
	return _AMM.Contract.Token2Address(&_AMM.CallOpts)
}

// Token2Address is a free data retrieval call binding the contract method 0x2456af82.
//
// Solidity: function token2Address() view returns(address)
func (_AMM *AMMCallerSession) Token2Address() (common.Address, error) {
	return _AMM.Contract.Token2Address(&_AMM.CallOpts)
}

// Exchange1to2 is a paid mutator transaction binding the contract method 0x3ede9696.
//
// Solidity: function exchange1to2(uint256 amount) returns(uint256)
func (_AMM *AMMTransactor) Exchange1to2(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _AMM.contract.Transact(opts, "exchange1to2", amount)
}

// Exchange1to2 is a paid mutator transaction binding the contract method 0x3ede9696.
//
// Solidity: function exchange1to2(uint256 amount) returns(uint256)
func (_AMM *AMMSession) Exchange1to2(amount *big.Int) (*types.Transaction, error) {
	return _AMM.Contract.Exchange1to2(&_AMM.TransactOpts, amount)
}

// Exchange1to2 is a paid mutator transaction binding the contract method 0x3ede9696.
//
// Solidity: function exchange1to2(uint256 amount) returns(uint256)
func (_AMM *AMMTransactorSession) Exchange1to2(amount *big.Int) (*types.Transaction, error) {
	return _AMM.Contract.Exchange1to2(&_AMM.TransactOpts, amount)
}

// Exchange2to1 is a paid mutator transaction binding the contract method 0xd030f9e5.
//
// Solidity: function exchange2to1(uint256 amount) returns(uint256)
func (_AMM *AMMTransactor) Exchange2to1(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _AMM.contract.Transact(opts, "exchange2to1", amount)
}

// Exchange2to1 is a paid mutator transaction binding the contract method 0xd030f9e5.
//
// Solidity: function exchange2to1(uint256 amount) returns(uint256)
func (_AMM *AMMSession) Exchange2to1(amount *big.Int) (*types.Transaction, error) {
	return _AMM.Contract.Exchange2to1(&_AMM.TransactOpts, amount)
}

// Exchange2to1 is a paid mutator transaction binding the contract method 0xd030f9e5.
//
// Solidity: function exchange2to1(uint256 amount) returns(uint256)
func (_AMM *AMMTransactorSession) Exchange2to1(amount *big.Int) (*types.Transaction, error) {
	return _AMM.Contract.Exchange2to1(&_AMM.TransactOpts, amount)
}

// SetRate is a paid mutator transaction binding the contract method 0x46df2ccb.
//
// Solidity: function setRate(uint256 rate1_, uint256 rate2_) returns()
func (_AMM *AMMTransactor) SetRate(opts *bind.TransactOpts, rate1_ *big.Int, rate2_ *big.Int) (*types.Transaction, error) {
	return _AMM.contract.Transact(opts, "setRate", rate1_, rate2_)
}

// SetRate is a paid mutator transaction binding the contract method 0x46df2ccb.
//
// Solidity: function setRate(uint256 rate1_, uint256 rate2_) returns()
func (_AMM *AMMSession) SetRate(rate1_ *big.Int, rate2_ *big.Int) (*types.Transaction, error) {
	return _AMM.Contract.SetRate(&_AMM.TransactOpts, rate1_, rate2_)
}

// SetRate is a paid mutator transaction binding the contract method 0x46df2ccb.
//
// Solidity: function setRate(uint256 rate1_, uint256 rate2_) returns()
func (_AMM *AMMTransactorSession) SetRate(rate1_ *big.Int, rate2_ *big.Int) (*types.Transaction, error) {
	return _AMM.Contract.SetRate(&_AMM.TransactOpts, rate1_, rate2_)
}

// ArbitrageMetaData contains all meta data concerning the Arbitrage contract.
var ArbitrageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token1_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token2_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"amm1_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"amm2_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"amm1Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amm2Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arbitrageur\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"vl\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"rl\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sl\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"va\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"ra\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sa\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"intrest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"loan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"loanHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lender_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"arbitrageur_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchange_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"loan_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"intrest_\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"loanHash_\",\"type\":\"bytes32\"}],\"name\":\"setup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token2Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"472fa96c": "amm1Address()",
		"87700c1d": "amm2Address()",
		"b3115c5c": "arbitrageur()",
		"d2f7265a": "exchange()",
		"99d94b0f": "execute(uint8,bytes32,bytes32,uint8,bytes32,bytes32)",
		"ffbe7402": "intrest()",
		"bcead63e": "lender()",
		"d285b7b4": "loan()",
		"1e306772": "loanHash()",
		"9afd9d78": "rollback()",
		"566b7de8": "setup(address,address,address,uint256,uint256,bytes32)",
		"200d2ed2": "status()",
		"d7cb416f": "token1Address()",
		"2456af82": "token2Address()",
	},
	Bin: "0x608060405234801561001057600080fd5b50604051610b27380380610b2783398101604081905261002f9161009e565b600080546001600160a01b039586166001600160a01b03199182161790915560018054948616948216949094179093556002805492851692841692909217909155600380549190931691161790556100f2565b80516001600160a01b038116811461009957600080fd5b919050565b600080600080608085870312156100b457600080fd5b6100bd85610082565b93506100cb60208601610082565b92506100d960408601610082565b91506100e760608601610082565b905092959194509250565b610a26806101016000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80639afd9d781161008c578063d285b7b411610066578063d285b7b414610228578063d2f7265a14610231578063d7cb416f14610244578063ffbe74021461025757600080fd5b80639afd9d78146101fa578063b3115c5c14610202578063bcead63e1461021557600080fd5b8063472fa96c116100c8578063472fa96c14610155578063566b7de81461016857806387700c1d146101d457806399d94b0f146101e757600080fd5b80631e306772146100ef578063200d2ed21461010b5780632456af821461012a575b600080fd5b6100f860095481565b6040519081526020015b60405180910390f35b600a546101189060ff1681565b60405160ff9091168152602001610102565b60015461013d906001600160a01b031681565b6040516001600160a01b039091168152602001610102565b60025461013d906001600160a01b031681565b6101d26101763660046108a1565b600480546001600160a01b03199081166001600160a01b039889161790915560058054821696881696909617909555600680549095169390951692909217909255600992909255600755600855600a805460ff19166001179055565b005b60035461013d906001600160a01b031681565b6101d26101f5366004610911565b610260565b6101d2610791565b60055461013d906001600160a01b031681565b60045461013d906001600160a01b031681565b6100f860075481565b60065461013d906001600160a01b031681565b60005461013d906001600160a01b031681565b6100f860085481565b600954604080516000808252602082018084529390935260ff891691810191909152606081018790526080810186905260019060a0016020604051602081039080840390855afa1580156102b8573d6000803e3d6000fd5b5050506020604051035190506000600160095486868660405160008152602001604052604051610304949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015610326573d6000803e3d6000fd5b5050604051601f1901516004549092506001600160a01b0384811691161490506103975760405162461bcd60e51b815260206004820152601860248201527f696e76616c6964206c656e646572207369676e6174757265000000000000000060448201526064015b60405180910390fd5b6005546001600160a01b038281169116146103f45760405162461bcd60e51b815260206004820152601d60248201527f696e76616c6964206172626974726167657572207369676e6174757265000000604482015260640161038e565b60025460035460005460015460075460405163095ea7b360e01b81526001600160a01b039586166004820181905260248201929092529094938416939283169290911690829063095ea7b3906044016020604051808303816000875af1158015610462573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104869190610969565b50600754604051631f6f4b4b60e11b81526000916001600160a01b03871691633ede9696916104bb9160040190815260200190565b6020604051808303816000875af11580156104da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104fe9190610992565b60035460405163095ea7b360e01b81526001600160a01b0391821660048201526024810183905291925083169063095ea7b3906044016020604051808303816000875af1158015610553573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105779190610969565b5060405163d030f9e560e01b8152600481018290526000906001600160a01b0386169063d030f9e5906024016020604051808303816000875af11580156105c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e69190610992565b90506008546007546105f891906109c1565b811161063a5760405162461bcd60e51b81526020600482015260116024820152701b9bdd08195b9bdd59da081c1c9bd99a5d607a1b604482015260640161038e565b6006546008546007546001600160a01b038088169363a9059cbb9391169161066291906109c1565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af11580156106ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106d19190610969565b506005546008546007546001600160a01b038088169363a9059cbb939116916106fa90866109d9565b61070491906109d9565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af115801561074f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107739190610969565b5050600a805460ff1916600217905550505050505050505050505050565b6006546001600160a01b03163314806107b457506005546001600160a01b031633145b6107f65760405162461bcd60e51b8152602060048201526013602482015272191bdb89dd081a185d9948185c1c1c9bdd985b606a1b604482015260640161038e565b60005460065460075460405163a9059cbb60e01b81526001600160a01b0392831660048201526024810191909152911690819063a9059cbb906044016020604051808303816000875af1158015610851573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108759190610969565b5050600a805460ff191681179055565b80356001600160a01b038116811461089c57600080fd5b919050565b60008060008060008060c087890312156108ba57600080fd5b6108c387610885565b95506108d160208801610885565b94506108df60408801610885565b9350606087013592506080870135915060a087013590509295509295509295565b803560ff8116811461089c57600080fd5b60008060008060008060c0878903121561092a57600080fd5b61093387610900565b9550602087013594506040870135935061094f60608801610900565b92506080870135915060a087013590509295509295509295565b60006020828403121561097b57600080fd5b8151801515811461098b57600080fd5b9392505050565b6000602082840312156109a457600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b600082198211156109d4576109d46109ab565b500190565b6000828210156109eb576109eb6109ab565b50039056fea2646970667358221220abe3b417486219c87dfe5e679ae89fe666a5a314904688a0b518163d3edb658164736f6c634300080b0033",
}

// ArbitrageABI is the input ABI used to generate the binding from.
// Deprecated: Use ArbitrageMetaData.ABI instead.
var ArbitrageABI = ArbitrageMetaData.ABI

// Deprecated: Use ArbitrageMetaData.Sigs instead.
// ArbitrageFuncSigs maps the 4-byte function signature to its string representation.
var ArbitrageFuncSigs = ArbitrageMetaData.Sigs

// ArbitrageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ArbitrageMetaData.Bin instead.
var ArbitrageBin = ArbitrageMetaData.Bin

// DeployArbitrage deploys a new Ethereum contract, binding an instance of Arbitrage to it.
func DeployArbitrage(auth *bind.TransactOpts, backend bind.ContractBackend, token1_ common.Address, token2_ common.Address, amm1_ common.Address, amm2_ common.Address) (common.Address, *types.Transaction, *Arbitrage, error) {
	parsed, err := ArbitrageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ArbitrageBin), backend, token1_, token2_, amm1_, amm2_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Arbitrage{ArbitrageCaller: ArbitrageCaller{contract: contract}, ArbitrageTransactor: ArbitrageTransactor{contract: contract}, ArbitrageFilterer: ArbitrageFilterer{contract: contract}}, nil
}

// Arbitrage is an auto generated Go binding around an Ethereum contract.
type Arbitrage struct {
	ArbitrageCaller     // Read-only binding to the contract
	ArbitrageTransactor // Write-only binding to the contract
	ArbitrageFilterer   // Log filterer for contract events
}

// ArbitrageCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbitrageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbitrageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbitrageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbitrageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbitrageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbitrageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbitrageSession struct {
	Contract     *Arbitrage        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbitrageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbitrageCallerSession struct {
	Contract *ArbitrageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ArbitrageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbitrageTransactorSession struct {
	Contract     *ArbitrageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ArbitrageRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbitrageRaw struct {
	Contract *Arbitrage // Generic contract binding to access the raw methods on
}

// ArbitrageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbitrageCallerRaw struct {
	Contract *ArbitrageCaller // Generic read-only contract binding to access the raw methods on
}

// ArbitrageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbitrageTransactorRaw struct {
	Contract *ArbitrageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbitrage creates a new instance of Arbitrage, bound to a specific deployed contract.
func NewArbitrage(address common.Address, backend bind.ContractBackend) (*Arbitrage, error) {
	contract, err := bindArbitrage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Arbitrage{ArbitrageCaller: ArbitrageCaller{contract: contract}, ArbitrageTransactor: ArbitrageTransactor{contract: contract}, ArbitrageFilterer: ArbitrageFilterer{contract: contract}}, nil
}

// NewArbitrageCaller creates a new read-only instance of Arbitrage, bound to a specific deployed contract.
func NewArbitrageCaller(address common.Address, caller bind.ContractCaller) (*ArbitrageCaller, error) {
	contract, err := bindArbitrage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrageCaller{contract: contract}, nil
}

// NewArbitrageTransactor creates a new write-only instance of Arbitrage, bound to a specific deployed contract.
func NewArbitrageTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbitrageTransactor, error) {
	contract, err := bindArbitrage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrageTransactor{contract: contract}, nil
}

// NewArbitrageFilterer creates a new log filterer instance of Arbitrage, bound to a specific deployed contract.
func NewArbitrageFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbitrageFilterer, error) {
	contract, err := bindArbitrage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbitrageFilterer{contract: contract}, nil
}

// bindArbitrage binds a generic wrapper to an already deployed contract.
func bindArbitrage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbitrageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arbitrage *ArbitrageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arbitrage.Contract.ArbitrageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arbitrage *ArbitrageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arbitrage.Contract.ArbitrageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arbitrage *ArbitrageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arbitrage.Contract.ArbitrageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arbitrage *ArbitrageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arbitrage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arbitrage *ArbitrageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arbitrage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arbitrage *ArbitrageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arbitrage.Contract.contract.Transact(opts, method, params...)
}

// Amm1Address is a free data retrieval call binding the contract method 0x472fa96c.
//
// Solidity: function amm1Address() view returns(address)
func (_Arbitrage *ArbitrageCaller) Amm1Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Arbitrage.contract.Call(opts, &out, "amm1Address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Amm1Address is a free data retrieval call binding the contract method 0x472fa96c.
//
// Solidity: function amm1Address() view returns(address)
func (_Arbitrage *ArbitrageSession) Amm1Address() (common.Address, error) {
	return _Arbitrage.Contract.Amm1Address(&_Arbitrage.CallOpts)
}

// Amm1Address is a free data retrieval call binding the contract method 0x472fa96c.
//
// Solidity: function amm1Address() view returns(address)
func (_Arbitrage *ArbitrageCallerSession) Amm1Address() (common.Address, error) {
	return _Arbitrage.Contract.Amm1Address(&_Arbitrage.CallOpts)
}

// Amm2Address is a free data retrieval call binding the contract method 0x87700c1d.
//
// Solidity: function amm2Address() view returns(address)
func (_Arbitrage *ArbitrageCaller) Amm2Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Arbitrage.contract.Call(opts, &out, "amm2Address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Amm2Address is a free data retrieval call binding the contract method 0x87700c1d.
//
// Solidity: function amm2Address() view returns(address)
func (_Arbitrage *ArbitrageSession) Amm2Address() (common.Address, error) {
	return _Arbitrage.Contract.Amm2Address(&_Arbitrage.CallOpts)
}

// Amm2Address is a free data retrieval call binding the contract method 0x87700c1d.
//
// Solidity: function amm2Address() view returns(address)
func (_Arbitrage *ArbitrageCallerSession) Amm2Address() (common.Address, error) {
	return _Arbitrage.Contract.Amm2Address(&_Arbitrage.CallOpts)
}

// Arbitrageur is a free data retrieval call binding the contract method 0xb3115c5c.
//
// Solidity: function arbitrageur() view returns(address)
func (_Arbitrage *ArbitrageCaller) Arbitrageur(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Arbitrage.contract.Call(opts, &out, "arbitrageur")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Arbitrageur is a free data retrieval call binding the contract method 0xb3115c5c.
//
// Solidity: function arbitrageur() view returns(address)
func (_Arbitrage *ArbitrageSession) Arbitrageur() (common.Address, error) {
	return _Arbitrage.Contract.Arbitrageur(&_Arbitrage.CallOpts)
}

// Arbitrageur is a free data retrieval call binding the contract method 0xb3115c5c.
//
// Solidity: function arbitrageur() view returns(address)
func (_Arbitrage *ArbitrageCallerSession) Arbitrageur() (common.Address, error) {
	return _Arbitrage.Contract.Arbitrageur(&_Arbitrage.CallOpts)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_Arbitrage *ArbitrageCaller) Exchange(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Arbitrage.contract.Call(opts, &out, "exchange")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_Arbitrage *ArbitrageSession) Exchange() (common.Address, error) {
	return _Arbitrage.Contract.Exchange(&_Arbitrage.CallOpts)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_Arbitrage *ArbitrageCallerSession) Exchange() (common.Address, error) {
	return _Arbitrage.Contract.Exchange(&_Arbitrage.CallOpts)
}

// Intrest is a free data retrieval call binding the contract method 0xffbe7402.
//
// Solidity: function intrest() view returns(uint256)
func (_Arbitrage *ArbitrageCaller) Intrest(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Arbitrage.contract.Call(opts, &out, "intrest")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Intrest is a free data retrieval call binding the contract method 0xffbe7402.
//
// Solidity: function intrest() view returns(uint256)
func (_Arbitrage *ArbitrageSession) Intrest() (*big.Int, error) {
	return _Arbitrage.Contract.Intrest(&_Arbitrage.CallOpts)
}

// Intrest is a free data retrieval call binding the contract method 0xffbe7402.
//
// Solidity: function intrest() view returns(uint256)
func (_Arbitrage *ArbitrageCallerSession) Intrest() (*big.Int, error) {
	return _Arbitrage.Contract.Intrest(&_Arbitrage.CallOpts)
}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() view returns(address)
func (_Arbitrage *ArbitrageCaller) Lender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Arbitrage.contract.Call(opts, &out, "lender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() view returns(address)
func (_Arbitrage *ArbitrageSession) Lender() (common.Address, error) {
	return _Arbitrage.Contract.Lender(&_Arbitrage.CallOpts)
}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() view returns(address)
func (_Arbitrage *ArbitrageCallerSession) Lender() (common.Address, error) {
	return _Arbitrage.Contract.Lender(&_Arbitrage.CallOpts)
}

// Loan is a free data retrieval call binding the contract method 0xd285b7b4.
//
// Solidity: function loan() view returns(uint256)
func (_Arbitrage *ArbitrageCaller) Loan(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Arbitrage.contract.Call(opts, &out, "loan")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Loan is a free data retrieval call binding the contract method 0xd285b7b4.
//
// Solidity: function loan() view returns(uint256)
func (_Arbitrage *ArbitrageSession) Loan() (*big.Int, error) {
	return _Arbitrage.Contract.Loan(&_Arbitrage.CallOpts)
}

// Loan is a free data retrieval call binding the contract method 0xd285b7b4.
//
// Solidity: function loan() view returns(uint256)
func (_Arbitrage *ArbitrageCallerSession) Loan() (*big.Int, error) {
	return _Arbitrage.Contract.Loan(&_Arbitrage.CallOpts)
}

// LoanHash is a free data retrieval call binding the contract method 0x1e306772.
//
// Solidity: function loanHash() view returns(bytes32)
func (_Arbitrage *ArbitrageCaller) LoanHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Arbitrage.contract.Call(opts, &out, "loanHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LoanHash is a free data retrieval call binding the contract method 0x1e306772.
//
// Solidity: function loanHash() view returns(bytes32)
func (_Arbitrage *ArbitrageSession) LoanHash() ([32]byte, error) {
	return _Arbitrage.Contract.LoanHash(&_Arbitrage.CallOpts)
}

// LoanHash is a free data retrieval call binding the contract method 0x1e306772.
//
// Solidity: function loanHash() view returns(bytes32)
func (_Arbitrage *ArbitrageCallerSession) LoanHash() ([32]byte, error) {
	return _Arbitrage.Contract.LoanHash(&_Arbitrage.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_Arbitrage *ArbitrageCaller) Status(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Arbitrage.contract.Call(opts, &out, "status")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_Arbitrage *ArbitrageSession) Status() (uint8, error) {
	return _Arbitrage.Contract.Status(&_Arbitrage.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_Arbitrage *ArbitrageCallerSession) Status() (uint8, error) {
	return _Arbitrage.Contract.Status(&_Arbitrage.CallOpts)
}

// Token1Address is a free data retrieval call binding the contract method 0xd7cb416f.
//
// Solidity: function token1Address() view returns(address)
func (_Arbitrage *ArbitrageCaller) Token1Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Arbitrage.contract.Call(opts, &out, "token1Address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1Address is a free data retrieval call binding the contract method 0xd7cb416f.
//
// Solidity: function token1Address() view returns(address)
func (_Arbitrage *ArbitrageSession) Token1Address() (common.Address, error) {
	return _Arbitrage.Contract.Token1Address(&_Arbitrage.CallOpts)
}

// Token1Address is a free data retrieval call binding the contract method 0xd7cb416f.
//
// Solidity: function token1Address() view returns(address)
func (_Arbitrage *ArbitrageCallerSession) Token1Address() (common.Address, error) {
	return _Arbitrage.Contract.Token1Address(&_Arbitrage.CallOpts)
}

// Token2Address is a free data retrieval call binding the contract method 0x2456af82.
//
// Solidity: function token2Address() view returns(address)
func (_Arbitrage *ArbitrageCaller) Token2Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Arbitrage.contract.Call(opts, &out, "token2Address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token2Address is a free data retrieval call binding the contract method 0x2456af82.
//
// Solidity: function token2Address() view returns(address)
func (_Arbitrage *ArbitrageSession) Token2Address() (common.Address, error) {
	return _Arbitrage.Contract.Token2Address(&_Arbitrage.CallOpts)
}

// Token2Address is a free data retrieval call binding the contract method 0x2456af82.
//
// Solidity: function token2Address() view returns(address)
func (_Arbitrage *ArbitrageCallerSession) Token2Address() (common.Address, error) {
	return _Arbitrage.Contract.Token2Address(&_Arbitrage.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x99d94b0f.
//
// Solidity: function execute(uint8 vl, bytes32 rl, bytes32 sl, uint8 va, bytes32 ra, bytes32 sa) returns()
func (_Arbitrage *ArbitrageTransactor) Execute(opts *bind.TransactOpts, vl uint8, rl [32]byte, sl [32]byte, va uint8, ra [32]byte, sa [32]byte) (*types.Transaction, error) {
	return _Arbitrage.contract.Transact(opts, "execute", vl, rl, sl, va, ra, sa)
}

// Execute is a paid mutator transaction binding the contract method 0x99d94b0f.
//
// Solidity: function execute(uint8 vl, bytes32 rl, bytes32 sl, uint8 va, bytes32 ra, bytes32 sa) returns()
func (_Arbitrage *ArbitrageSession) Execute(vl uint8, rl [32]byte, sl [32]byte, va uint8, ra [32]byte, sa [32]byte) (*types.Transaction, error) {
	return _Arbitrage.Contract.Execute(&_Arbitrage.TransactOpts, vl, rl, sl, va, ra, sa)
}

// Execute is a paid mutator transaction binding the contract method 0x99d94b0f.
//
// Solidity: function execute(uint8 vl, bytes32 rl, bytes32 sl, uint8 va, bytes32 ra, bytes32 sa) returns()
func (_Arbitrage *ArbitrageTransactorSession) Execute(vl uint8, rl [32]byte, sl [32]byte, va uint8, ra [32]byte, sa [32]byte) (*types.Transaction, error) {
	return _Arbitrage.Contract.Execute(&_Arbitrage.TransactOpts, vl, rl, sl, va, ra, sa)
}

// Rollback is a paid mutator transaction binding the contract method 0x9afd9d78.
//
// Solidity: function rollback() returns()
func (_Arbitrage *ArbitrageTransactor) Rollback(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arbitrage.contract.Transact(opts, "rollback")
}

// Rollback is a paid mutator transaction binding the contract method 0x9afd9d78.
//
// Solidity: function rollback() returns()
func (_Arbitrage *ArbitrageSession) Rollback() (*types.Transaction, error) {
	return _Arbitrage.Contract.Rollback(&_Arbitrage.TransactOpts)
}

// Rollback is a paid mutator transaction binding the contract method 0x9afd9d78.
//
// Solidity: function rollback() returns()
func (_Arbitrage *ArbitrageTransactorSession) Rollback() (*types.Transaction, error) {
	return _Arbitrage.Contract.Rollback(&_Arbitrage.TransactOpts)
}

// Setup is a paid mutator transaction binding the contract method 0x566b7de8.
//
// Solidity: function setup(address lender_, address arbitrageur_, address exchange_, uint256 loan_, uint256 intrest_, bytes32 loanHash_) returns()
func (_Arbitrage *ArbitrageTransactor) Setup(opts *bind.TransactOpts, lender_ common.Address, arbitrageur_ common.Address, exchange_ common.Address, loan_ *big.Int, intrest_ *big.Int, loanHash_ [32]byte) (*types.Transaction, error) {
	return _Arbitrage.contract.Transact(opts, "setup", lender_, arbitrageur_, exchange_, loan_, intrest_, loanHash_)
}

// Setup is a paid mutator transaction binding the contract method 0x566b7de8.
//
// Solidity: function setup(address lender_, address arbitrageur_, address exchange_, uint256 loan_, uint256 intrest_, bytes32 loanHash_) returns()
func (_Arbitrage *ArbitrageSession) Setup(lender_ common.Address, arbitrageur_ common.Address, exchange_ common.Address, loan_ *big.Int, intrest_ *big.Int, loanHash_ [32]byte) (*types.Transaction, error) {
	return _Arbitrage.Contract.Setup(&_Arbitrage.TransactOpts, lender_, arbitrageur_, exchange_, loan_, intrest_, loanHash_)
}

// Setup is a paid mutator transaction binding the contract method 0x566b7de8.
//
// Solidity: function setup(address lender_, address arbitrageur_, address exchange_, uint256 loan_, uint256 intrest_, bytes32 loanHash_) returns()
func (_Arbitrage *ArbitrageTransactorSession) Setup(lender_ common.Address, arbitrageur_ common.Address, exchange_ common.Address, loan_ *big.Int, intrest_ *big.Int, loanHash_ [32]byte) (*types.Transaction, error) {
	return _Arbitrage.Contract.Setup(&_Arbitrage.TransactOpts, lender_, arbitrageur_, exchange_, loan_, intrest_, loanHash_)
}

// ERC20MetaData contains all meta data concerning the ERC20 contract.
var ERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b506040516106fb3803806106fb83398101604081905261002f916100e1565b610039338261003f565b50610120565b6001600160a01b0382166100995760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b80600260008282546100ab91906100fa565b90915550506001600160a01b038216600090815260208190526040812080548392906100d89084906100fa565b90915550505050565b6000602082840312156100f357600080fd5b5051919050565b6000821982111561011b57634e487b7160e01b600052601160045260246000fd5b500190565b6105cc8061012f6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063095ea7b31461006757806318160ddd1461008f57806323b872dd146100a157806370a08231146100b4578063a9059cbb146100dd578063dd62ed3e146100f0575b600080fd5b61007a610075366004610496565b610129565b60405190151581526020015b60405180910390f35b6002545b604051908152602001610086565b61007a6100af3660046104c0565b61013f565b6100936100c23660046104fc565b6001600160a01b031660009081526020819052604090205490565b61007a6100eb366004610496565b6101f5565b6100936100fe36600461051e565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6000610136338484610202565b50600192915050565b600061014c8484846102f1565b6001600160a01b0384166000908152600160209081526040808320338452909152902054828110156101d65760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b60648201526084015b60405180910390fd5b6101ea85336101e58685610567565b610202565b506001949350505050565b60006101363384846102f1565b6001600160a01b0383166102645760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016101cd565b6001600160a01b0382166102c55760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016101cd565b6001600160a01b0392831660009081526001602090815260408083209490951682529290925291902055565b6001600160a01b0383166103555760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016101cd565b6001600160a01b0382166103b75760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016101cd565b6001600160a01b0383166000908152602081905260409020548181101561042f5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016101cd565b6104398282610567565b6001600160a01b03808616600090815260208190526040808220939093559085168152908120805484929061046f90849061057e565b909155505050505050565b80356001600160a01b038116811461049157600080fd5b919050565b600080604083850312156104a957600080fd5b6104b28361047a565b946020939093013593505050565b6000806000606084860312156104d557600080fd5b6104de8461047a565b92506104ec6020850161047a565b9150604084013590509250925092565b60006020828403121561050e57600080fd5b6105178261047a565b9392505050565b6000806040838503121561053157600080fd5b61053a8361047a565b91506105486020840161047a565b90509250929050565b634e487b7160e01b600052601160045260246000fd5b60008282101561057957610579610551565b500390565b6000821982111561059157610591610551565b50019056fea2646970667358221220167eedd3d3d0b737f6f7e88bb85ddd4e3d312531186d7c2154f8c4e95e3eb20b64736f6c634300080b0033",
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
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}
