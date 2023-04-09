package main

import (
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	"encoding/json"
	"log"
	"net/http"

	"github.com/Guy1m0/piechain-frontend/cclib"
	"github.com/Guy1m0/piechain-frontend/contracts/eth_auction"
	"github.com/Guy1m0/piechain-frontend/examples/auction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/cors"
)

// Define request and response structures
type APIResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type CreateAuctionRequest struct {
	AssetCC  []byte
	AssetID  []byte
	Platform string
}

var ethClient *ethclient.Client
var quorumClient *ethclient.Client
var assetClient *auction.AssetClient

var asset *auction.Asset
var myAuction *auction.Auction

// Add a function to start the HTTP server
func startHTTPServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})

	mux.HandleFunc("/api/add-asset", handleAddAsset)
	mux.HandleFunc("/api/start-auction", handleStartAuction)

	// Enable CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)

	go func() {
		fmt.Println("Starting server on port 6789")
		err := http.ListenAndServe(":6789", handler)
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
}

// Add handler functions for each API endpoint
func handleAddAsset(w http.ResponseWriter, r *http.Request) {
	assetID := r.URL.Query().Get("asset_id")
	if assetID == "" {
		http.Error(w, "Missing asset_id parameter", http.StatusBadRequest)
		return
	}
	fmt.Println("Receive asset ID:", assetID)

	asset := addAsset(assetID)

	type responseStruct struct {
		Owner string `json:"Owner"`
	}

	// Set the content type to JSON
	//w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Create a sample response
	response := responseStruct{
		Owner: asset.Owner,
	}

	// Send the response as JSON
	fmt.Println("Sending response:", response)
	json.NewEncoder(w).Encode(response)

}

func handleStartAuction(w http.ResponseWriter, r *http.Request) {
	assetID := r.URL.Query().Get("asset_id")
	if assetID == "" {
		http.Error(w, "Missing asset_id parameter", http.StatusBadRequest)
		return
	}

	type responseStruct struct {
		Quorum string `json:"Quorum"`
		Eth    string `json:"Eth"`
	}

	asset := addAsset(assetID)
	fmt.Println("Adding asset", assetID)

	fmt.Println("Starting auction")
	fmt.Println("[ethereum] Deploying auction")
	ethAddr := deployCrossChainAuction(ethClient)
	// Make ethAddr different from quorumAddr
	//time.Sleep(3 * time.Second

	fmt.Println("[quorum] Deploying auction")
	quorumAddr := deployCrossChainAuction(quorumClient)

	fmt.Println("[fabric] Creating auction")
	myAuction = startAuction(asset.ID, ethAddr, quorumAddr)

	// Set the content type to JSON
	//w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Create a sample response
	response := responseStruct{
		Quorum: quorumAddr,
		Eth:    ethAddr,
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	// Initialize system

	// Start two clients for users to send/sign tx
	var err error
	ethClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8545", "localhost"))
	check(err)

	quorumClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8546", "localhost"))
	check(err)

	assetClient = auction.NewAssetClient()

	//var asset *auction.Asset
	//var myAuction *auction.Auction

	//fmt.Println("Start to listening")
	startHTTPServer()
	//fmt.Println("[fabric] Adding asset")

	// Wait for an interrupt signal to gracefully shut down the server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	// Start the server

	//fmt.Println("[fabric] Adding asset")–
	/*
		fmt.Println("[fabric] Adding asset")
		asset = addAsset("asset1")

		fmt.Println("[ethereum] Bidding auction")
		bidAuction(ethClient, myAuction.EthAddr, "../../keys/key1", 500)

		fmt.Println("[quorum] Bidding auction")
		bidAuction(quorumClient, myAuction.QuorumAddr, "../../keys/key2", 1000)

		fmt.Println("[fabric] Ending auction")
		endAuction(myAuction)
	*/
}

func addAsset(id string) *auction.Asset {
	// auth is the admin account to add asset
	auth, err := cclib.NewTransactor("../../keys/key0", "password")
	check(err)
	_, err = assetClient.AddAsset(id, auth.From.Hex())
	check(err)
	//time.Sleep(3 * time.Second)
	asset, err := assetClient.GetAsset(id)
	check(err)
	fmt.Println("Asset added, owner: ", asset.Owner)
	return asset
}

// more like a bridge?
func startAuction(assetID, ethAddr, quorumAddr string) *auction.Auction {
	args := auction.StartAuctionArgs{
		AssetID:    assetID,
		EthAddr:    ethAddr,
		QuorumAddr: quorumAddr,
	}
	_, err := assetClient.StartAuction(args)
	check(err)
	//time.Sleep(3 * time.Second)
	fmt.Println("Started auction for asset")

	auctionID, err := assetClient.GetLastAuctionID()
	check(err)
	//fmt.Println("AuctionID: ", auctionID)

	a, err := assetClient.GetAuction(auctionID)
	check(err)
	return a
}

func endAuction(a *auction.Auction) {
	_, err := assetClient.EndAuction(a.AssetID)
	check(err)

	for {
		time.Sleep(1 * time.Second)
		a, err = assetClient.GetAuction(a.ID)
		check(err)
		if a.Status == "Ended" {

			fmt.Println("Auction Ended")
			fmt.Println("Highest Bidder: ", a.HighestBidder)
			fmt.Println("Highest Bid: ", a.HighestBid)
			fmt.Println("Highest Bid Platform: ", a.HighestBidPlatform)

			asset, err := assetClient.GetAsset(a.AssetID)
			check(err)
			fmt.Println("New Asset Owner: ", asset.Owner)

			break
		}
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

	highestBidder, err := auctionSession.HighestBidder()
	check(err)
	fmt.Println("highest bidder:", highestBidder.Hex())

	highestBid, err := auctionSession.HighestBid()
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
