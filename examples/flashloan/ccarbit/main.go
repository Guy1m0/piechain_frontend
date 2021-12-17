package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/aungmawjj/piechain/cclib"
	"github.com/aungmawjj/piechain/examples/flashloan"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	zkNodes = "localhost:2181"

	ccsvc *cclib.CCService

	ethClient *ethclient.Client
)

func main() {
	flag.StringVar(&zkNodes, "zk", zkNodes, "comma separated zoolkeeper nodes")
	flag.Parse()

	var err error

	ethClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8546", "localhost"))
	check(err)

	ccsvc, err = cclib.NewEventService(strings.Split(zkNodes, ","), "ccarbit")
	check(err)
	ccsvc.Register(flashloan.OnInitializedLending, handleFlashloanInitialized)
	err = ccsvc.Start()
	check(err)

	select {}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
