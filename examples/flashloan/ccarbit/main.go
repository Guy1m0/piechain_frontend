package main

import (
	"flag"
	"strings"

	"github.com/aungmawjj/piechain/cclib"
	"github.com/aungmawjj/piechain/examples/flashloan"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	zkNodes = "localhost:2181"

	excKey   = "../../keys/key1"
	password = "password"

	setupInfoFile = "../setup_info.json"

	ccsvc *cclib.CCService

	ethClient *ethclient.Client
	excT      *bind.TransactOpts
)

func main() {
	flag.StringVar(&zkNodes, "zk", zkNodes, "comma separated zoolkeeper nodes")
	flag.Parse()

	var err error

	ethClient = flashloan.NewEthClient()
	excT, err = cclib.NewTransactor(excKey, password)
	check(err)

	ccsvc, err = cclib.NewEventService(strings.Split(zkNodes, ","), "ccarbit")
	check(err)
	ccsvc.Register(flashloan.OnInitializedLending, handleFlashloanInitialized)
	err = ccsvc.Start(true)
	check(err)

	select {}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
