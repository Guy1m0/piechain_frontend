package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/aungmawjj/piechain/cclib"
	"github.com/aungmawjj/piechain/examples/auction_pow"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	zkNodes     = "localhost:2181"
	ethEndpoint = "localhost:8545"

	ccsvc     *cclib.CCService
	ethClient *ethclient.Client
	rpcClient *rpc.Client
)

func main() {
	flag.StringVar(&zkNodes, "zk", zkNodes, "comma separated zoolkeeper nodes")
	flag.StringVar(&ethEndpoint, "eth", ethEndpoint, "eth endpoint")
	flag.Parse()

	var err error
	rpcClient, err = rpc.Dial(fmt.Sprintf("http://%s", ethEndpoint))
	check(err)

	ethClient, err = ethclient.Dial(fmt.Sprintf("http://%s", ethEndpoint))
	check(err)

	ccsvc, err = cclib.NewEventService(strings.Split(zkNodes, ","), "end_listener")
	check(err)

	ccsvc.Register(auction_pow.TopicOnBindAuction, onBindAuction)

	err = ccsvc.Start()
	check(err)

	select {}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
