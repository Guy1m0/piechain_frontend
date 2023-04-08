package main

import (
	"encoding/json"
	"time"

	"github.com/Guy1m0/piechain-frontend/examples/auction"
)

func runAuctionListener() {
	for {
		a := listenNewAuction()
		go onNewAuction(a)
	}
}

func onNewAuction(a *auction.Auction) {
	auctionResults[a.ID] = &auction.FinalizeAuctionArgs{
		AuctionID: a.ID,
	}
	onAuctionEnding(listenAuctionEnding(a))
}

func onAuctionEnding(a *auction.Auction) {
	b, _ := json.Marshal(a)
	ccsvc.Publish(auction.AuctionEndingEvent, b)
}

func listenNewAuction() *auction.Auction {
	lastID, err := assetClient.GetLastAuctionID()
	check(err)

	for {
		time.Sleep(1 * time.Second)
		auctionID, err := assetClient.GetLastAuctionID()
		check(err)
		if auctionID > lastID {
			auction, err := assetClient.GetAuction(auctionID)
			check(err)
			return auction
		}
	}
}

func listenAuctionEnding(a *auction.Auction) *auction.Auction {
	for {
		time.Sleep(1 * time.Second)
		auction, err := assetClient.GetAuction(a.ID)
		check(err)
		if auction.Status == "Ending" {
			return auction
		}
	}
}
