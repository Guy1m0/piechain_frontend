package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"

	"github.com/aungmawjj/piechain/cclib"
	"github.com/aungmawjj/piechain/contracts/eth_auction"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	password           = "password"
	auctionAddressFile = "auction_address"
)

var k0, k1 *bind.TransactOpts
var ethClient *ethclient.Client
var rpcClient *rpc.Client

func main() {
	var doSetup bool

	flag.BoolVar(&doSetup, "s", false, "do setup (i.e deploy and bid auction)")
	flag.Parse()

	connect()
	if doSetup {
		setup()
	}

	verifyAuctionProof()
}

func connect() {
	var err error
	k0, err = cclib.NewTransactor("../../keys/key0", password)
	check(err)
	k1, err = cclib.NewTransactor("../../keys/key1", password)
	check(err)

	ethClient, err = ethclient.Dial("http://localhost:8545")
	check(err)

	rpcClient, err = rpc.Dial("http://localhost:8545")
	check(err)
}

func setup() {
	auctionAddr, tx, _, err := eth_auction.DeployAuction(k0, ethClient)
	check(err)
	waitTx(tx, "deploy auction")

	ioutil.WriteFile(auctionAddressFile, auctionAddr.Bytes(), 0644)

	auction, err := eth_auction.NewAuction(auctionAddr, ethClient)
	check(err)

	k1.Value = big.NewInt(1000)
	tx, err = auction.Bid(k1)
	check(err)
	waitTx(tx, "bid auction")
}

func readAuctionAddress() common.Address {
	b, err := ioutil.ReadFile(auctionAddressFile)
	check(err)
	return common.BytesToAddress(b)
}

func verifyAuctionProof() {
	auctionAddr := readAuctionAddress()

	blockHeader, err := ethClient.HeaderByNumber(context.Background(), nil)
	check(err)

	result, err := GetProof(auctionAddr, []string{"0x1", "0x2"}, blockHeader.Number)
	check(err)

	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "  ")

	fmt.Println("BlockHeader:")
	e.Encode(blockHeader)

	fmt.Println("Proof:")
	e.Encode(result)

	ok, err := VerifyProof(blockHeader.Root, result)
	check(err)
	fmt.Println("verify proof", ok)

	// ok, err := VerifyAccountProof(result)
	// check(err)
	// fmt.Printf("account proof valid: %t\n", ok)

	// ok, err := VerifyStorageProof(result)
	// check(err)
	// fmt.Printf("storage proof valid: %t\n", ok)

	// accValid, err := ethstorageproof.VerifyEthAccountProof(&proof)
	// check(err)
	// fmt.Printf("account proof valid: %t\n", accValid)

	// for i, sp := range proof.StorageProof {
	// 	valid, err := ethstorageproof.VerifyEthStorageProof(&sp, proof.StorageHash)
	// 	check(err)

	// }
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func waitTx(tx *types.Transaction, label string) {
	fmt.Println(label + "...")
	success, err := cclib.WaitTx(ethClient, tx.Hash())
	check(err)
	printTxStatus(success)
}

func printTxStatus(success bool) {
	if success {
		fmt.Println("Transaction successful")
	} else {
		fmt.Println("Transaction failed")
	}
}
