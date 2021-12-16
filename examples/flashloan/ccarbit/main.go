package main

import (
	"flag"
	"strings"

	"github.com/aungmawjj/piechain/cclib"
	"github.com/aungmawjj/piechain/examples/flashloan"
)

var (
	zkNodes = "localhost:2181"

	ccsvc *cclib.CCService
)

func main() {
	flag.StringVar(&zkNodes, "zk", zkNodes, "comma separated zoolkeeper nodes")
	flag.Parse()

	var err error

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
