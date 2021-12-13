package main

import (
	"flag"
	"strings"

	"github.com/aungmawjj/piechain/cclib"
	"github.com/aungmawjj/piechain/examples/auction"
)

var (
	zkNodes = "localhost:2181"

	assetClient  *auction.AssetClient
	eventService *cclib.EventService

	auctionResults map[int]*auction.FinalizeAuctionArgs
)

func main() {
	flag.StringVar(&zkNodes, "zk", zkNodes, "comma separated zoolkeeper nodes")
	flag.Parse()

	var err error
	assetClient = auction.NewAssetClient()
	eventService, err = cclib.NewEventService(strings.Split(zkNodes, ","), "relayer")
	check(err)

	eventService.Register(auction.SignedAuctionResultEvent, handleSignedAuctionResult)

	err = eventService.Start()
	check(err)

	runAuctionListener()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
