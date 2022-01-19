// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth_auction

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

// AuctionMetaData contains all meta data concerning the Auction contract.
var AuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AuctionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"HighestBidIncreased\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ended\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"38af3eed": "beneficiary()",
		"1998aeef": "bid()",
		"fe67a54b": "endAuction()",
		"12fa6feb": "ended()",
		"d57bde79": "highestBid()",
		"91f90157": "highestBidder()",
		"3ccfd60b": "withdraw()",
	},
	Bin: "0x608060405234801561001057600080fd5b50600080546001600160a01b031916331790556103f1806100326000396000f3fe6080604052600436106100705760003560e01c80633ccfd60b1161004e5780633ccfd60b146100e657806391f90157146100fb578063d57bde791461011b578063fe67a54b1461013f57600080fd5b806312fa6feb146100755780631998aeef146100a457806338af3eed146100ae575b600080fd5b34801561008157600080fd5b5060045461008f9060ff1681565b60405190151581526020015b60405180910390f35b6100ac610154565b005b3480156100ba57600080fd5b506000546100ce906001600160a01b031681565b6040516001600160a01b03909116815260200161009b565b3480156100f257600080fd5b5061008f61023c565b34801561010757600080fd5b506001546100ce906001600160a01b031681565b34801561012757600080fd5b5061013160025481565b60405190815260200161009b565b34801561014b57600080fd5b506100ac6102a0565b60025434116101aa5760405162461bcd60e51b815260206004820152601e60248201527f546865726520616c7265616479206973206120686967686572206269642e000060448201526064015b60405180910390fd5b600254156101e5576002546001546001600160a01b0316600090815260036020526040812080549091906101df908490610395565b90915550505b600180546001600160a01b031916339081179091553460028190556040805192835260208301919091527ff4757a49b326036464bec6fe419a4ae38c8a02ce3e68bf0809674f6aab8ad300910160405180910390a1565b33600090815260036020526040812054801561029857336000818152600360205260408082208290555183156108fc0291849190818181858888f193505050506102985733600090815260036020526040812091909155919050565b600191505090565b60045460ff16156102ff5760405162461bcd60e51b815260206004820152602360248201527f61756374696f6e456e642068617320616c7265616479206265656e2063616c6c60448201526232b21760e91b60648201526084016101a1565b6004805460ff1916600190811790915554600254604080516001600160a01b03909316835260208301919091527fdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda910160405180910390a1600080546002546040516001600160a01b039092169281156108fc029290818181858888f19350505050158015610392573d6000803e3d6000fd5b50565b600082198211156103b657634e487b7160e01b600052601160045260246000fd5b50019056fea2646970667358221220adfe73743230318c18ca0d4bc5dd626087639b5118d14a9cbc86f6b78cf86ab164736f6c634300080b0033",
}

// AuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use AuctionMetaData.ABI instead.
var AuctionABI = AuctionMetaData.ABI

// Deprecated: Use AuctionMetaData.Sigs instead.
// AuctionFuncSigs maps the 4-byte function signature to its string representation.
var AuctionFuncSigs = AuctionMetaData.Sigs

// AuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AuctionMetaData.Bin instead.
var AuctionBin = AuctionMetaData.Bin

// DeployAuction deploys a new Ethereum contract, binding an instance of Auction to it.
func DeployAuction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Auction, error) {
	parsed, err := AuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AuctionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Auction{AuctionCaller: AuctionCaller{contract: contract}, AuctionTransactor: AuctionTransactor{contract: contract}, AuctionFilterer: AuctionFilterer{contract: contract}}, nil
}

// Auction is an auto generated Go binding around an Ethereum contract.
type Auction struct {
	AuctionCaller     // Read-only binding to the contract
	AuctionTransactor // Write-only binding to the contract
	AuctionFilterer   // Log filterer for contract events
}

// AuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuctionSession struct {
	Contract     *Auction          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuctionCallerSession struct {
	Contract *AuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuctionTransactorSession struct {
	Contract     *AuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuctionRaw struct {
	Contract *Auction // Generic contract binding to access the raw methods on
}

// AuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuctionCallerRaw struct {
	Contract *AuctionCaller // Generic read-only contract binding to access the raw methods on
}

// AuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuctionTransactorRaw struct {
	Contract *AuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuction creates a new instance of Auction, bound to a specific deployed contract.
func NewAuction(address common.Address, backend bind.ContractBackend) (*Auction, error) {
	contract, err := bindAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Auction{AuctionCaller: AuctionCaller{contract: contract}, AuctionTransactor: AuctionTransactor{contract: contract}, AuctionFilterer: AuctionFilterer{contract: contract}}, nil
}

// NewAuctionCaller creates a new read-only instance of Auction, bound to a specific deployed contract.
func NewAuctionCaller(address common.Address, caller bind.ContractCaller) (*AuctionCaller, error) {
	contract, err := bindAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionCaller{contract: contract}, nil
}

// NewAuctionTransactor creates a new write-only instance of Auction, bound to a specific deployed contract.
func NewAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*AuctionTransactor, error) {
	contract, err := bindAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionTransactor{contract: contract}, nil
}

// NewAuctionFilterer creates a new log filterer instance of Auction, bound to a specific deployed contract.
func NewAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*AuctionFilterer, error) {
	contract, err := bindAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuctionFilterer{contract: contract}, nil
}

// bindAuction binds a generic wrapper to an already deployed contract.
func bindAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction *AuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auction.Contract.AuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction *AuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.Contract.AuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction *AuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auction.Contract.AuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction *AuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction *AuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction *AuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auction.Contract.contract.Transact(opts, method, params...)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_Auction *AuctionCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "beneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_Auction *AuctionSession) Beneficiary() (common.Address, error) {
	return _Auction.Contract.Beneficiary(&_Auction.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_Auction *AuctionCallerSession) Beneficiary() (common.Address, error) {
	return _Auction.Contract.Beneficiary(&_Auction.CallOpts)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Auction *AuctionCaller) Ended(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "ended")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Auction *AuctionSession) Ended() (bool, error) {
	return _Auction.Contract.Ended(&_Auction.CallOpts)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Auction *AuctionCallerSession) Ended() (bool, error) {
	return _Auction.Contract.Ended(&_Auction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Auction *AuctionCaller) HighestBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "highestBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Auction *AuctionSession) HighestBid() (*big.Int, error) {
	return _Auction.Contract.HighestBid(&_Auction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Auction *AuctionCallerSession) HighestBid() (*big.Int, error) {
	return _Auction.Contract.HighestBid(&_Auction.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Auction *AuctionCaller) HighestBidder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "highestBidder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Auction *AuctionSession) HighestBidder() (common.Address, error) {
	return _Auction.Contract.HighestBidder(&_Auction.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Auction *AuctionCallerSession) HighestBidder() (common.Address, error) {
	return _Auction.Contract.HighestBidder(&_Auction.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Auction *AuctionTransactor) Bid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "bid")
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Auction *AuctionSession) Bid() (*types.Transaction, error) {
	return _Auction.Contract.Bid(&_Auction.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Auction *AuctionTransactorSession) Bid() (*types.Transaction, error) {
	return _Auction.Contract.Bid(&_Auction.TransactOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_Auction *AuctionTransactor) EndAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "endAuction")
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_Auction *AuctionSession) EndAuction() (*types.Transaction, error) {
	return _Auction.Contract.EndAuction(&_Auction.TransactOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_Auction *AuctionTransactorSession) EndAuction() (*types.Transaction, error) {
	return _Auction.Contract.EndAuction(&_Auction.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_Auction *AuctionTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_Auction *AuctionSession) Withdraw() (*types.Transaction, error) {
	return _Auction.Contract.Withdraw(&_Auction.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_Auction *AuctionTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Auction.Contract.Withdraw(&_Auction.TransactOpts)
}

// AuctionAuctionEndedIterator is returned from FilterAuctionEnded and is used to iterate over the raw logs and unpacked data for AuctionEnded events raised by the Auction contract.
type AuctionAuctionEndedIterator struct {
	Event *AuctionAuctionEnded // Event containing the contract specifics and raw log

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
func (it *AuctionAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionAuctionEnded)
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
		it.Event = new(AuctionAuctionEnded)
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
func (it *AuctionAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionAuctionEnded represents a AuctionEnded event raised by the Auction contract.
type AuctionAuctionEnded struct {
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuctionEnded is a free log retrieval operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address winner, uint256 amount)
func (_Auction *AuctionFilterer) FilterAuctionEnded(opts *bind.FilterOpts) (*AuctionAuctionEndedIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return &AuctionAuctionEndedIterator{contract: _Auction.contract, event: "AuctionEnded", logs: logs, sub: sub}, nil
}

// WatchAuctionEnded is a free log subscription operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address winner, uint256 amount)
func (_Auction *AuctionFilterer) WatchAuctionEnded(opts *bind.WatchOpts, sink chan<- *AuctionAuctionEnded) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionAuctionEnded)
				if err := _Auction.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
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

// ParseAuctionEnded is a log parse operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address winner, uint256 amount)
func (_Auction *AuctionFilterer) ParseAuctionEnded(log types.Log) (*AuctionAuctionEnded, error) {
	event := new(AuctionAuctionEnded)
	if err := _Auction.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionHighestBidIncreasedIterator is returned from FilterHighestBidIncreased and is used to iterate over the raw logs and unpacked data for HighestBidIncreased events raised by the Auction contract.
type AuctionHighestBidIncreasedIterator struct {
	Event *AuctionHighestBidIncreased // Event containing the contract specifics and raw log

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
func (it *AuctionHighestBidIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionHighestBidIncreased)
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
		it.Event = new(AuctionHighestBidIncreased)
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
func (it *AuctionHighestBidIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionHighestBidIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionHighestBidIncreased represents a HighestBidIncreased event raised by the Auction contract.
type AuctionHighestBidIncreased struct {
	Bidder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHighestBidIncreased is a free log retrieval operation binding the contract event 0xf4757a49b326036464bec6fe419a4ae38c8a02ce3e68bf0809674f6aab8ad300.
//
// Solidity: event HighestBidIncreased(address bidder, uint256 amount)
func (_Auction *AuctionFilterer) FilterHighestBidIncreased(opts *bind.FilterOpts) (*AuctionHighestBidIncreasedIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return &AuctionHighestBidIncreasedIterator{contract: _Auction.contract, event: "HighestBidIncreased", logs: logs, sub: sub}, nil
}

// WatchHighestBidIncreased is a free log subscription operation binding the contract event 0xf4757a49b326036464bec6fe419a4ae38c8a02ce3e68bf0809674f6aab8ad300.
//
// Solidity: event HighestBidIncreased(address bidder, uint256 amount)
func (_Auction *AuctionFilterer) WatchHighestBidIncreased(opts *bind.WatchOpts, sink chan<- *AuctionHighestBidIncreased) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionHighestBidIncreased)
				if err := _Auction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
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

// ParseHighestBidIncreased is a log parse operation binding the contract event 0xf4757a49b326036464bec6fe419a4ae38c8a02ce3e68bf0809674f6aab8ad300.
//
// Solidity: event HighestBidIncreased(address bidder, uint256 amount)
func (_Auction *AuctionFilterer) ParseHighestBidIncreased(log types.Log) (*AuctionHighestBidIncreased, error) {
	event := new(AuctionHighestBidIncreased)
	if err := _Auction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
