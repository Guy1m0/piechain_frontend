package main

import (
	"encoding/json"

	"github.com/aungmawjj/piechain/contracts/eth_auction"
	"github.com/aungmawjj/piechain/examples/auction"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func handleAuctionEnding(payload []byte) {
	var a auction.Auction
	err := json.Unmarshal(payload, &a)
	check(err)

	var addr string
	if platform == "ethereum" {
		addr = a.EthAddr
	} else {
		addr = a.QuorumAddr
	}
	result := auction.AuctionResult{
		Platform:      platform,
		HostAuctionID: a.ID,
		AuctionAddr:   addr,
	}

	contract, err := eth_auction.NewAuction(common.HexToAddress(addr), ethClient)
	check(err)
	opts := &bind.CallOpts{}

	highestBid, err := contract.HighestBid(opts)
	check(err)

	highestBidder, err := contract.HighestBidder(opts)
	check(err)

	result.HighestBid = int(highestBid.Int64())
	result.HighestBidder = highestBidder.Hex()

	signed := auction.SignedAuctionResult{
		AuctionResult: result,
	}
	signed.Signature, err = signer.Sign(result.Hash())
	check(err)

	b, _ := json.Marshal(signed)
	eventService.Publish(auction.SignedAuctionResultEvent, b)
}
