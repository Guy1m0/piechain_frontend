// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth_arbitrage

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

// AMMABI is the input ABI used to generate the binding from.
const AMMABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token1_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token2_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"exchange1to2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"exchange2to1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate1_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate2_\",\"type\":\"uint256\"}],\"name\":\"setRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token1Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token2Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

var AMMParsedABI, _ = abi.JSON(strings.NewReader(AMMABI))

// AMMFuncSigs maps the 4-byte function signature to its string representation.
var AMMFuncSigs = map[string]string{
	"3ede9696": "exchange1to2(uint256)",
	"d030f9e5": "exchange2to1(uint256)",
	"cf854969": "rate1()",
	"f555b815": "rate2()",
	"46df2ccb": "setRate(uint256,uint256)",
	"d7cb416f": "token1Address()",
	"2456af82": "token2Address()",
}

// AMMBin is the compiled bytecode used for deploying new contracts.
var AMMBin = "0x608060405234801561001057600080fd5b506040516104473803806104478339818101604052604081101561003357600080fd5b508051602090910151600280546001600160a01b039384166001600160a01b031991821617909155600380549390921692169190911790556103cd8061007a6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063cf8549691161005b578063cf854969146100fa578063d030f9e514610102578063d7cb416f1461011f578063f555b815146101275761007d565b80632456af82146100825780633ede9696146100a657806346df2ccb146100d5575b600080fd5b61008a61012f565b604080516001600160a01b039092168252519081900360200190f35b6100c3600480360360208110156100bc57600080fd5b503561013e565b60408051918252519081900360200190f35b6100f8600480360360408110156100eb57600080fd5b5080359060200135610273565b005b6100c361027e565b6100c36004803603602081101561011857600080fd5b5035610284565b61008a610383565b6100c3610392565b6003546001600160a01b031681565b6002546003546000805460015491936001600160a01b0390811693169184919086028161016757fe5b604080516323b872dd60e01b81523360048201523060248201526044810189905290519290910492506001600160a01b038516916323b872dd916064808201926020929091908290030181600087803b1580156101c357600080fd5b505af11580156101d7573d6000803e3d6000fd5b505050506040513d60208110156101ed57600080fd5b50506040805163a9059cbb60e01b81523360048201526024810183905290516001600160a01b0384169163a9059cbb9160448083019260209291908290030181600087803b15801561023e57600080fd5b505af1158015610252573d6000803e3d6000fd5b505050506040513d602081101561026857600080fd5b509095945050505050565b600091909155600155565b60005481565b6002546003546001546000805490936001600160a01b0390811693169184918602816102ac57fe5b604080516323b872dd60e01b81523360048201523060248201526044810189905290519290910492506001600160a01b038416916323b872dd916064808201926020929091908290030181600087803b15801561030857600080fd5b505af115801561031c573d6000803e3d6000fd5b505050506040513d602081101561033257600080fd5b50506040805163a9059cbb60e01b81523360048201526024810183905290516001600160a01b0385169163a9059cbb9160448083019260209291908290030181600087803b15801561023e57600080fd5b6002546001600160a01b031681565b6001548156fea265627a7a72315820cb194bd228faf002d9a091edd76d2686facb67dfbd85a5705045eb076524c40964736f6c63430005110032"

// DeployAMM deploys a new Ethereum contract, binding an instance of AMM to it.
func DeployAMM(auth *bind.TransactOpts, backend bind.ContractBackend, token1_ common.Address, token2_ common.Address) (common.Address, *types.Transaction, *AMM, error) {
	parsed, err := abi.JSON(strings.NewReader(AMMABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AMMBin), backend, token1_, token2_)
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
func (_AMM *AMMRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_AMM *AMMCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
// Solidity: function rate1() constant returns(uint256)
func (_AMM *AMMCaller) Rate1(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AMM.contract.Call(opts, out, "rate1")
	return *ret0, err
}

// Rate1 is a free data retrieval call binding the contract method 0xcf854969.
//
// Solidity: function rate1() constant returns(uint256)
func (_AMM *AMMSession) Rate1() (*big.Int, error) {
	return _AMM.Contract.Rate1(&_AMM.CallOpts)
}

// Rate1 is a free data retrieval call binding the contract method 0xcf854969.
//
// Solidity: function rate1() constant returns(uint256)
func (_AMM *AMMCallerSession) Rate1() (*big.Int, error) {
	return _AMM.Contract.Rate1(&_AMM.CallOpts)
}

// Rate2 is a free data retrieval call binding the contract method 0xf555b815.
//
// Solidity: function rate2() constant returns(uint256)
func (_AMM *AMMCaller) Rate2(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AMM.contract.Call(opts, out, "rate2")
	return *ret0, err
}

// Rate2 is a free data retrieval call binding the contract method 0xf555b815.
//
// Solidity: function rate2() constant returns(uint256)
func (_AMM *AMMSession) Rate2() (*big.Int, error) {
	return _AMM.Contract.Rate2(&_AMM.CallOpts)
}

// Rate2 is a free data retrieval call binding the contract method 0xf555b815.
//
// Solidity: function rate2() constant returns(uint256)
func (_AMM *AMMCallerSession) Rate2() (*big.Int, error) {
	return _AMM.Contract.Rate2(&_AMM.CallOpts)
}

// Token1Address is a free data retrieval call binding the contract method 0xd7cb416f.
//
// Solidity: function token1Address() constant returns(address)
func (_AMM *AMMCaller) Token1Address(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AMM.contract.Call(opts, out, "token1Address")
	return *ret0, err
}

// Token1Address is a free data retrieval call binding the contract method 0xd7cb416f.
//
// Solidity: function token1Address() constant returns(address)
func (_AMM *AMMSession) Token1Address() (common.Address, error) {
	return _AMM.Contract.Token1Address(&_AMM.CallOpts)
}

// Token1Address is a free data retrieval call binding the contract method 0xd7cb416f.
//
// Solidity: function token1Address() constant returns(address)
func (_AMM *AMMCallerSession) Token1Address() (common.Address, error) {
	return _AMM.Contract.Token1Address(&_AMM.CallOpts)
}

// Token2Address is a free data retrieval call binding the contract method 0x2456af82.
//
// Solidity: function token2Address() constant returns(address)
func (_AMM *AMMCaller) Token2Address(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AMM.contract.Call(opts, out, "token2Address")
	return *ret0, err
}

// Token2Address is a free data retrieval call binding the contract method 0x2456af82.
//
// Solidity: function token2Address() constant returns(address)
func (_AMM *AMMSession) Token2Address() (common.Address, error) {
	return _AMM.Contract.Token2Address(&_AMM.CallOpts)
}

// Token2Address is a free data retrieval call binding the contract method 0x2456af82.
//
// Solidity: function token2Address() constant returns(address)
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

// ArbitrageABI is the input ABI used to generate the binding from.
const ArbitrageABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token1_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token2_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"amm1_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"amm2_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"amm1Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"amm2Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"arbitrageur\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"vl\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"rl\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sl\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"va\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"ra\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sa\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"intrest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loanHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"rollback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"lender_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"arbitrageur_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchange_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"loan_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"intrest_\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"loanHash_\",\"type\":\"bytes32\"}],\"name\":\"setup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token1Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token2Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

var ArbitrageParsedABI, _ = abi.JSON(strings.NewReader(ArbitrageABI))

// ArbitrageFuncSigs maps the 4-byte function signature to its string representation.
var ArbitrageFuncSigs = map[string]string{
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
}

// ArbitrageBin is the compiled bytecode used for deploying new contracts.
var ArbitrageBin = "0x608060405234801561001057600080fd5b506040516107e83803806107e88339818101604052608081101561003357600080fd5b50805160208201516040830151606090930151600080546001600160a01b039485166001600160a01b0319918216179091556001805493851693821693909317909255600280549484169483169490941790935560038054929093169116179055610745806100a36000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80639afd9d781161008c578063d285b7b411610066578063d285b7b4146101ff578063d2f7265a14610207578063d7cb416f1461020f578063ffbe740214610217576100ea565b80639afd9d78146101e7578063b3115c5c146101ef578063bcead63e146101f7576100ea565b8063472fa96c116100c8578063472fa96c1461014b578063566b7de81461015357806387700c1d1461019d57806399d94b0f146101a5576100ea565b80631e306772146100ef578063200d2ed2146101095780632456af8214610127575b600080fd5b6100f761021f565b60408051918252519081900360200190f35b610111610225565b6040805160ff9092168252519081900360200190f35b61012f61022e565b604080516001600160a01b039092168252519081900360200190f35b61012f61023d565b61019b600480360360c081101561016957600080fd5b506001600160a01b03813581169160208101358216916040820135169060608101359060808101359060a0013561024c565b005b61012f6102a8565b61019b600480360360c08110156101bb57600080fd5b5060ff813581169160208101359160408201359160608101359091169060808101359060a001356102b7565b61019b610650565b61012f6106c8565b61012f6106d7565b6100f76106e6565b61012f6106ec565b61012f6106fb565b6100f761070a565b60095481565b600a5460ff1681565b6001546001600160a01b031681565b6002546001600160a01b031681565b600480546001600160a01b03199081166001600160a01b039889161790915560058054821696881696909617909555600680549095169390951692909217909255600992909255600755600855600a805460ff19166001179055565b6003546001600160a01b031681565b600254600354600080546001546007546040805163095ea7b360e01b81526001600160a01b039788166004820181905260248201939093529051919695861695938416949390921692849263095ea7b39260448083019360209383900390910190829087803b15801561032957600080fd5b505af115801561033d573d6000803e3d6000fd5b505050506040513d602081101561035357600080fd5b505060075460408051631f6f4b4b60e11b81526004810192909252516000916001600160a01b03871691633ede96969160248082019260209290919082900301818787803b1580156103a457600080fd5b505af11580156103b8573d6000803e3d6000fd5b505050506040513d60208110156103ce57600080fd5b50516003546040805163095ea7b360e01b81526001600160a01b0392831660048201526024810184905290519293509084169163095ea7b3916044808201926020929091908290030181600087803b15801561042957600080fd5b505af115801561043d573d6000803e3d6000fd5b505050506040513d602081101561045357600080fd5b50506040805163d030f9e560e01b81526004810183905290516000916001600160a01b0387169163d030f9e59160248082019260209290919082900301818787803b1580156104a157600080fd5b505af11580156104b5573d6000803e3d6000fd5b505050506040513d60208110156104cb57600080fd5b505160085460075491925001811161051e576040805162461bcd60e51b81526020600482015260116024820152701b9bdd08195b9bdd59da081c1c9bd99a5d607a1b604482015290519081900360640190fd5b6006546008546007546040805163a9059cbb60e01b81526001600160a01b03948516600482015291909201602482015290519186169163a9059cbb916044808201926020929091908290030181600087803b15801561057c57600080fd5b505af1158015610590573d6000803e3d6000fd5b505050506040513d60208110156105a657600080fd5b50506005546008546007546040805163a9059cbb60e01b81526001600160a01b03948516600482015291850392909203602482015290519186169163a9059cbb916044808201926020929091908290030181600087803b15801561060957600080fd5b505af115801561061d573d6000803e3d6000fd5b505050506040513d602081101561063357600080fd5b5050600a805460ff19166002179055505050505050505050505050565b6006546001600160a01b031633148061067357506005546001600160a01b031633145b6106ba576040805162461bcd60e51b8152602060048201526013602482015272191bdb89dd081a185d9948185c1c1c9bdd985b606a1b604482015290519081900360640190fd5b600a805460ff191681179055565b6005546001600160a01b031681565b6004546001600160a01b031681565b60075481565b6006546001600160a01b031681565b6000546001600160a01b031681565b6008548156fea265627a7a7231582005aa4a0d23b04f0b06718dbf5c5e491cbf7d6831b562b2cc08a0bfbcc048be3d64736f6c63430005110032"

// DeployArbitrage deploys a new Ethereum contract, binding an instance of Arbitrage to it.
func DeployArbitrage(auth *bind.TransactOpts, backend bind.ContractBackend, token1_ common.Address, token2_ common.Address, amm1_ common.Address, amm2_ common.Address) (common.Address, *types.Transaction, *Arbitrage, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbitrageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbitrageBin), backend, token1_, token2_, amm1_, amm2_)
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
func (_Arbitrage *ArbitrageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_Arbitrage *ArbitrageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
// Solidity: function amm1Address() constant returns(address)
func (_Arbitrage *ArbitrageCaller) Amm1Address(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Arbitrage.contract.Call(opts, out, "amm1Address")
	return *ret0, err
}

// Amm1Address is a free data retrieval call binding the contract method 0x472fa96c.
//
// Solidity: function amm1Address() constant returns(address)
func (_Arbitrage *ArbitrageSession) Amm1Address() (common.Address, error) {
	return _Arbitrage.Contract.Amm1Address(&_Arbitrage.CallOpts)
}

// Amm1Address is a free data retrieval call binding the contract method 0x472fa96c.
//
// Solidity: function amm1Address() constant returns(address)
func (_Arbitrage *ArbitrageCallerSession) Amm1Address() (common.Address, error) {
	return _Arbitrage.Contract.Amm1Address(&_Arbitrage.CallOpts)
}

// Amm2Address is a free data retrieval call binding the contract method 0x87700c1d.
//
// Solidity: function amm2Address() constant returns(address)
func (_Arbitrage *ArbitrageCaller) Amm2Address(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Arbitrage.contract.Call(opts, out, "amm2Address")
	return *ret0, err
}

// Amm2Address is a free data retrieval call binding the contract method 0x87700c1d.
//
// Solidity: function amm2Address() constant returns(address)
func (_Arbitrage *ArbitrageSession) Amm2Address() (common.Address, error) {
	return _Arbitrage.Contract.Amm2Address(&_Arbitrage.CallOpts)
}

// Amm2Address is a free data retrieval call binding the contract method 0x87700c1d.
//
// Solidity: function amm2Address() constant returns(address)
func (_Arbitrage *ArbitrageCallerSession) Amm2Address() (common.Address, error) {
	return _Arbitrage.Contract.Amm2Address(&_Arbitrage.CallOpts)
}

// Arbitrageur is a free data retrieval call binding the contract method 0xb3115c5c.
//
// Solidity: function arbitrageur() constant returns(address)
func (_Arbitrage *ArbitrageCaller) Arbitrageur(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Arbitrage.contract.Call(opts, out, "arbitrageur")
	return *ret0, err
}

// Arbitrageur is a free data retrieval call binding the contract method 0xb3115c5c.
//
// Solidity: function arbitrageur() constant returns(address)
func (_Arbitrage *ArbitrageSession) Arbitrageur() (common.Address, error) {
	return _Arbitrage.Contract.Arbitrageur(&_Arbitrage.CallOpts)
}

// Arbitrageur is a free data retrieval call binding the contract method 0xb3115c5c.
//
// Solidity: function arbitrageur() constant returns(address)
func (_Arbitrage *ArbitrageCallerSession) Arbitrageur() (common.Address, error) {
	return _Arbitrage.Contract.Arbitrageur(&_Arbitrage.CallOpts)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() constant returns(address)
func (_Arbitrage *ArbitrageCaller) Exchange(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Arbitrage.contract.Call(opts, out, "exchange")
	return *ret0, err
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() constant returns(address)
func (_Arbitrage *ArbitrageSession) Exchange() (common.Address, error) {
	return _Arbitrage.Contract.Exchange(&_Arbitrage.CallOpts)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() constant returns(address)
func (_Arbitrage *ArbitrageCallerSession) Exchange() (common.Address, error) {
	return _Arbitrage.Contract.Exchange(&_Arbitrage.CallOpts)
}

// Intrest is a free data retrieval call binding the contract method 0xffbe7402.
//
// Solidity: function intrest() constant returns(uint256)
func (_Arbitrage *ArbitrageCaller) Intrest(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Arbitrage.contract.Call(opts, out, "intrest")
	return *ret0, err
}

// Intrest is a free data retrieval call binding the contract method 0xffbe7402.
//
// Solidity: function intrest() constant returns(uint256)
func (_Arbitrage *ArbitrageSession) Intrest() (*big.Int, error) {
	return _Arbitrage.Contract.Intrest(&_Arbitrage.CallOpts)
}

// Intrest is a free data retrieval call binding the contract method 0xffbe7402.
//
// Solidity: function intrest() constant returns(uint256)
func (_Arbitrage *ArbitrageCallerSession) Intrest() (*big.Int, error) {
	return _Arbitrage.Contract.Intrest(&_Arbitrage.CallOpts)
}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() constant returns(address)
func (_Arbitrage *ArbitrageCaller) Lender(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Arbitrage.contract.Call(opts, out, "lender")
	return *ret0, err
}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() constant returns(address)
func (_Arbitrage *ArbitrageSession) Lender() (common.Address, error) {
	return _Arbitrage.Contract.Lender(&_Arbitrage.CallOpts)
}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() constant returns(address)
func (_Arbitrage *ArbitrageCallerSession) Lender() (common.Address, error) {
	return _Arbitrage.Contract.Lender(&_Arbitrage.CallOpts)
}

// Loan is a free data retrieval call binding the contract method 0xd285b7b4.
//
// Solidity: function loan() constant returns(uint256)
func (_Arbitrage *ArbitrageCaller) Loan(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Arbitrage.contract.Call(opts, out, "loan")
	return *ret0, err
}

// Loan is a free data retrieval call binding the contract method 0xd285b7b4.
//
// Solidity: function loan() constant returns(uint256)
func (_Arbitrage *ArbitrageSession) Loan() (*big.Int, error) {
	return _Arbitrage.Contract.Loan(&_Arbitrage.CallOpts)
}

// Loan is a free data retrieval call binding the contract method 0xd285b7b4.
//
// Solidity: function loan() constant returns(uint256)
func (_Arbitrage *ArbitrageCallerSession) Loan() (*big.Int, error) {
	return _Arbitrage.Contract.Loan(&_Arbitrage.CallOpts)
}

// LoanHash is a free data retrieval call binding the contract method 0x1e306772.
//
// Solidity: function loanHash() constant returns(bytes32)
func (_Arbitrage *ArbitrageCaller) LoanHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Arbitrage.contract.Call(opts, out, "loanHash")
	return *ret0, err
}

// LoanHash is a free data retrieval call binding the contract method 0x1e306772.
//
// Solidity: function loanHash() constant returns(bytes32)
func (_Arbitrage *ArbitrageSession) LoanHash() ([32]byte, error) {
	return _Arbitrage.Contract.LoanHash(&_Arbitrage.CallOpts)
}

// LoanHash is a free data retrieval call binding the contract method 0x1e306772.
//
// Solidity: function loanHash() constant returns(bytes32)
func (_Arbitrage *ArbitrageCallerSession) LoanHash() ([32]byte, error) {
	return _Arbitrage.Contract.LoanHash(&_Arbitrage.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(uint8)
func (_Arbitrage *ArbitrageCaller) Status(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Arbitrage.contract.Call(opts, out, "status")
	return *ret0, err
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(uint8)
func (_Arbitrage *ArbitrageSession) Status() (uint8, error) {
	return _Arbitrage.Contract.Status(&_Arbitrage.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(uint8)
func (_Arbitrage *ArbitrageCallerSession) Status() (uint8, error) {
	return _Arbitrage.Contract.Status(&_Arbitrage.CallOpts)
}

// Token1Address is a free data retrieval call binding the contract method 0xd7cb416f.
//
// Solidity: function token1Address() constant returns(address)
func (_Arbitrage *ArbitrageCaller) Token1Address(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Arbitrage.contract.Call(opts, out, "token1Address")
	return *ret0, err
}

// Token1Address is a free data retrieval call binding the contract method 0xd7cb416f.
//
// Solidity: function token1Address() constant returns(address)
func (_Arbitrage *ArbitrageSession) Token1Address() (common.Address, error) {
	return _Arbitrage.Contract.Token1Address(&_Arbitrage.CallOpts)
}

// Token1Address is a free data retrieval call binding the contract method 0xd7cb416f.
//
// Solidity: function token1Address() constant returns(address)
func (_Arbitrage *ArbitrageCallerSession) Token1Address() (common.Address, error) {
	return _Arbitrage.Contract.Token1Address(&_Arbitrage.CallOpts)
}

// Token2Address is a free data retrieval call binding the contract method 0x2456af82.
//
// Solidity: function token2Address() constant returns(address)
func (_Arbitrage *ArbitrageCaller) Token2Address(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Arbitrage.contract.Call(opts, out, "token2Address")
	return *ret0, err
}

// Token2Address is a free data retrieval call binding the contract method 0x2456af82.
//
// Solidity: function token2Address() constant returns(address)
func (_Arbitrage *ArbitrageSession) Token2Address() (common.Address, error) {
	return _Arbitrage.Contract.Token2Address(&_Arbitrage.CallOpts)
}

// Token2Address is a free data retrieval call binding the contract method 0x2456af82.
//
// Solidity: function token2Address() constant returns(address)
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

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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
var ERC20Bin = "0x608060405234801561001057600080fd5b5060405161062c38038061062c8339818101604052602081101561003357600080fd5b505161004833826001600160e01b0361004e16565b506100d3565b6001600160a01b0382166100a9576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b60028054820190556001600160a01b03909116600090815260208190526040902080549091019055565b61054a806100e26000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063095ea7b31461006757806318160ddd146100a757806323b872dd146100c157806370a08231146100f7578063a9059cbb1461011d578063dd62ed3e14610149575b600080fd5b6100936004803603604081101561007d57600080fd5b506001600160a01b038135169060200135610177565b604080519115158252519081900360200190f35b6100af61018d565b60408051918252519081900360200190f35b610093600480360360608110156100d757600080fd5b506001600160a01b03813581169160208101359091169060400135610193565b6100af6004803603602081101561010d57600080fd5b50356001600160a01b031661021b565b6100936004803603604081101561013357600080fd5b506001600160a01b038135169060200135610236565b6100af6004803603604081101561015f57600080fd5b506001600160a01b0381358116916020013516610243565b600061018433848461026e565b50600192915050565b60025490565b60006101a0848484610324565b6001600160a01b0384166000908152600160209081526040808320338452909152902054828110156102035760405162461bcd60e51b81526004018080602001828103825260288152602001806104a56028913960400191505060405180910390fd5b610210853385840361026e565b506001949350505050565b6001600160a01b031660009081526020819052604090205490565b6000610184338484610324565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b0383166102b35760405162461bcd60e51b81526004018080602001828103825260248152602001806104f26024913960400191505060405180910390fd5b6001600160a01b0382166102f85760405162461bcd60e51b815260040180806020018281038252602281526020018061045d6022913960400191505060405180910390fd5b6001600160a01b0392831660009081526001602090815260408083209490951682529290925291902055565b6001600160a01b0383166103695760405162461bcd60e51b81526004018080602001828103825260258152602001806104cd6025913960400191505060405180910390fd5b6001600160a01b0382166103ae5760405162461bcd60e51b815260040180806020018281038252602381526020018061043a6023913960400191505060405180910390fd5b6001600160a01b038316600090815260208190526040902054818110156104065760405162461bcd60e51b815260040180806020018281038252602681526020018061047f6026913960400191505060405180910390fd5b6001600160a01b03938416600090815260208190526040808220928490039092559290931682529190208054909101905556fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f2061646472657373a265627a7a72315820a99946582a2ff0bd8537f55b5b9bbdf38b35d92c9292101945caefe702d961ec64736f6c63430005110032"

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
