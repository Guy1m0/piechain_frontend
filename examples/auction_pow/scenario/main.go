package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/Guy1m0/piechain-frontend/cclib"
	"github.com/Guy1m0/piechain-frontend/contracts/eth_auction"
	"github.com/Guy1m0/piechain-frontend/examples/auction_pow"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CreateAuctionRequest struct {
	AssetCC  []byte
	AssetID  []byte
	Platform string
}

var ethClient *ethclient.Client
var assetClient *auction_pow.AssetClient

func main() {
	var err error
	ethClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8545", "localhost"))
	check(err)
	assetClient = auction_pow.NewAssetClient()

	var asset *auction_pow.Asset
	var myAuction *auction_pow.Auction

	fmt.Println("[fabric] Adding asset")
	asset = addAsset("asset1")

	fmt.Println("[ethereum] Deploying auction")
	ethAddr := deployCrossChainAuction(ethClient)

	fmt.Println("[fabric] Creating auction")
	myAuction = startAuction(asset.ID, ethAddr)

	// fmt.Println("auction address", myAuction.AuctionAddr)

	fmt.Println("[ethereum] Bidding auction")
	bidAuction(ethClient, myAuction.AuctionAddr, "../../keys/key1", 500)
	bidAuction(ethClient, myAuction.AuctionAddr, "../../keys/key2", 1000)

	fmt.Println("[ethereum] End auction")
	endAuction(ethAddr, ethClient)

	time.Sleep(4 * time.Second)

	myAuction, err = assetClient.GetAuction(myAuction.ID)
	check(err)
	fmt.Println("[fabric] Auction ended")
	fmt.Printf("Highest Bidder: %s\n", myAuction.HighestBidder)
	fmt.Printf("Highest Bid: %d\n", myAuction.HighestBid)

	asset, err = assetClient.GetAsset(asset.ID)
	check(err)
	fmt.Printf("Asset owner: %s\n", asset.Owner)
}

func addAsset(id string) *auction_pow.Asset {
	auth, err := cclib.NewTransactor("../../keys/key0", "password")
	check(err)
	_, err = assetClient.AddAsset(id, auth.From.Hex())
	check(err)
	time.Sleep(3 * time.Second)
	asset, err := assetClient.GetAsset(id)
	check(err)
	fmt.Println("Asset added, owner: ", asset.Owner)
	return asset
}

func startAuction(assetID, auctionAddr string) *auction_pow.Auction {
	lastBlock, err := ethClient.HeaderByNumber(context.Background(), nil)
	check(err)

	args := auction_pow.StartAuctionArgs{
		AssetID:     assetID,
		AuctionAddr: auctionAddr,
		EthHeader:   lastBlock,
	}
	_, err = assetClient.StartAuction(args)
	check(err)
	time.Sleep(3 * time.Second)
	fmt.Println("Started auction for asset")

	auctionID, err := assetClient.GetLastAuctionID()
	check(err)
	fmt.Println("AuctionID: ", auctionID)

	a, err := assetClient.GetAuction(auctionID)
	check(err)
	return a
}

func endAuction(addrHex string, client *ethclient.Client) {
	addr := common.HexToAddress(addrHex)
	auctionSession := newAuctionSession(addr, client, "../../keys/key0")
	tx, err := auctionSession.EndAuction()
	check(err)
	success, err := cclib.WaitTx(client, tx.Hash())
	check(err)
	printTxStatus(success)
	if !success {
		panic("failed to end auction")
	}
}

func bidAuction(client *ethclient.Client, addrHex, keyfile string, value int64) {
	addr := common.HexToAddress(addrHex)
	auctionSession := newAuctionSession(addr, client, keyfile)
	auctionSession.TransactOpts.Value = big.NewInt(value)
	tx, err := auctionSession.Bid()
	check(err)
	success, err := cclib.WaitTx(client, tx.Hash())
	check(err)
	printTxStatus(success)
	if !success {
		panic("failed to bid auction")
	}
	auctionSession.TransactOpts.Value = big.NewInt(0)

	contract, err := eth_auction.NewAuction(addr, client)
	check(err)
	highestBidder, err := contract.HighestBidder(&bind.CallOpts{})
	check(err)
	fmt.Println("highest bidder:", highestBidder.Hex())

	highestBid, err := contract.HighestBid(&bind.CallOpts{})
	check(err)
	fmt.Println("highest bid:", highestBid)
}

func deployCrossChainAuction(client *ethclient.Client) string {
	auth, err := cclib.NewTransactor("../../keys/key0", "password")
	check(err)

	addr, tx, _, err := eth_auction.DeployAuction(auth, client)
	check(err)

	success, err := cclib.WaitTx(client, tx.Hash())
	check(err)

	printTxStatus(success)
	fmt.Printf("Auction contract address: %s\n", addr.Hex())

	return addr.Hex()
}

func newAuctionSession(
	addr common.Address, eth *ethclient.Client, keyfile string,
) *eth_auction.AuctionSession {
	auth, err := cclib.NewTransactor(keyfile, "password")
	check(err)

	cc, err := eth_auction.NewAuction(addr, eth)
	check(err)

	return &eth_auction.AuctionSession{
		Contract:     cc,
		TransactOpts: *auth,
	}
}

func printTxStatus(success bool) {
	if success {
		fmt.Println("Transaction successful")
	} else {
		fmt.Println("Transaction failed")
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
