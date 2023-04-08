package main

import (
	"flag"
	"strings"

	"github.com/Guy1m0/piechain-frontend/cclib"
	"github.com/Guy1m0/piechain-frontend/examples/auction"
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

	auctionResults = make(map[int]*auction.FinalizeAuctionArgs)
	assetClient = auction.NewAssetClient()

	var err error

	ccsvc, err = cclib.NewEventService(strings.Split(zkNodes, ","), "relayer")
	check(err)
	ccsvc.Register(auction.SignedAuctionResultEvent, handleSignedAuctionResult)
	err = ccsvc.Start(true)
	check(err)

	runAuctionListener()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
