package main

import (
	"flag"
	"strings"

	"github.com/aungmawjj/piechain/cclib"
	"github.com/aungmawjj/piechain/examples/auction_pow"
)

var (
	zkNodes = "localhost:2181"

	assetClient *auction_pow.AssetClient
	ccsvc       *cclib.CCService
)

func main() {
	flag.StringVar(&zkNodes, "zk", zkNodes, "comma separated zoolkeeper nodes")
	flag.Parse()

	assetClient = auction_pow.NewAssetClient()

	var err error

	ccsvc, err = cclib.NewEventService(strings.Split(zkNodes, ","), "relayer")
	check(err)
	ccsvc.Register(auction_pow.TopicOnEndAuction, handleOnEndAuction)
	err = ccsvc.Start(true)
	check(err)

	runAuctionListener()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
