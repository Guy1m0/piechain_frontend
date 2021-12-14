package main

import (
	"flag"
	"strings"

	"github.com/aungmawjj/piechain/cclib"
	"github.com/aungmawjj/piechain/examples/auction"
)

var (
	zkNodes = "localhost:2181"

	assetClient *auction.AssetClient
	ccsvc       *cclib.CCService

	auctionResults map[int]*auction.FinalizeAuctionArgs
)

func main() {
	flag.StringVar(&zkNodes, "zk", zkNodes, "comma separated zoolkeeper nodes")
	flag.Parse()

	var err error
	assetClient = auction.NewAssetClient()
	ccsvc, err = cclib.NewEventService(strings.Split(zkNodes, ","), "relayer")
	check(err)

	auctionResults = make(map[int]*auction.FinalizeAuctionArgs)

	ccsvc.Register(auction.SignedAuctionResultEvent, handleSignedAuctionResult)

	err = ccsvc.Start()
	check(err)

	runAuctionListener()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
