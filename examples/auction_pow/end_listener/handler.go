package main

import (
	"context"
	"encoding/json"
	"math/big"
	"time"

	"github.com/aungmawjj/piechain/contracts/eth_auction"
	"github.com/aungmawjj/piechain/examples/auction_pow"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func onBindAuction(data []byte) {
	var auction auction_pow.Auction
	err := json.Unmarshal(data, &auction)
	check(err)
	err = listenAuctionEnd(&auction)
	check(err)

	result := getEndAuctionResults(&auction)
	payload, _ := json.Marshal(result)
	err = ccsvc.Publish(auction_pow.TopicOnEndAuction, payload)
	check(err)
}

func listenAuctionEnd(auction *auction_pow.Auction) error {
	contract, err := eth_auction.NewAuction(common.HexToAddress(auction.AuctionAddr), ethClient)
	if err != nil {
		return err
	}
	opts := &bind.CallOpts{}

	time.Sleep(3 * time.Second)
	for {
		time.Sleep(1 * time.Second)
		ended, _ := contract.Ended(opts)
		if ended {
			return nil
		}
	}
}

func getEndAuctionResults(auction *auction_pow.Auction) *auction_pow.EndAuctionArgs {
	lastBlock, err := ethClient.HeaderByNumber(context.Background(), nil)
	check(err)

	var result auction_pow.EndAuctionArgs
	result.ProvableResult, err = GetProof(
		common.HexToAddress(auction.AuctionAddr), []string{"0x1", "0x2"}, lastBlock.Number,
	)
	check(err)

	n := auction.BaseEthHeader.Number
	for n.Cmp(lastBlock.Number) == -1 {
		n = big.NewInt(0).Add(n, big.NewInt(1))
		h, err := ethClient.HeaderByNumber(context.Background(), n)
		check(err)
		result.EthHeaders = append(result.EthHeaders, h)
	}

	return &result
}
