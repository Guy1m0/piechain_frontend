package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

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
	var floan flashloan.Flashloan
	err := c.Bind(&floan)
	if err != nil {
		return err
	}

	fmt.Printf("register new flashloan, loan: %d, intrest: %d\n", floan.Loan, floan.Intrest)
	go listenLenderInitialize(&floan)
	return c.NoContent(http.StatusOK)
}

func listenLenderInitialize(floan *flashloan.Flashloan) {
	lenderCode := flashloan.NewChaincode(floan.LenderContract)
	for {
		resp, err := lenderCode.EvaluateTransaction("GetStatus")
		check(err)

		status, err := strconv.Atoi(string(resp))
		check(err)

		if status == 2 {
			break
		}
		time.Sleep(1 * time.Second)
	}

	payload, _ := json.Marshal(floan)
	err := ccsvc.Publish(flashloan.OnInitializedLending, payload)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
