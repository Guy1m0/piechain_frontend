package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/aungmawjj/piechain/cclib"
	"github.com/aungmawjj/piechain/examples/flashloan"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	zkNodes = "localhost:2181"

	ccsvc *cclib.CCService
)

func main() {
	flag.StringVar(&zkNodes, "zk", zkNodes, "comma separated zoolkeeper nodes")
	flag.Parse()

	var err error

	ccsvc, err = cclib.NewEventService(strings.Split(zkNodes, ","), "cclender")
	check(err)
	ccsvc.Register(flashloan.OnLoanSuccessful, handleFlashloanSuccessful)
	ccsvc.Register(flashloan.OnLoanFail, handleFlashloanFail)
	err = ccsvc.Start()
	check(err)

	e := echo.New()
	e.Use(middleware.Recover())
	e.POST("/flashloan", createFlashloan)

	log.Fatal(e.Start(":9000"))
}

func createFlashloan(c echo.Context) error {
	var req flashloan.Flashloan
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	// go listen init
	return c.NoContent(http.StatusOK)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
