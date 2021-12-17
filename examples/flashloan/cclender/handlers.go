package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aungmawjj/piechain/examples/flashloan"
)

func handleFlashloanSuccessful(payload []byte) {
	var floan flashloan.Flashloan
	err := json.Unmarshal(payload, &floan)
	check(err)

	lenderCode := flashloan.NewChaincode(floan.LenderContract)
	_, err = lenderCode.EvaluateTransaction("EndLoan", "true")
	check(err)
	time.Sleep(3 * time.Second)
	fmt.Println("Ended flash loan. Refund loan + intrest: ", floan.Loan+floan.Intrest)
}

func handleFlashloanFail(payload []byte) {
	var floan flashloan.Flashloan
	err := json.Unmarshal(payload, &floan)
	check(err)

	lenderCode := flashloan.NewChaincode(floan.LenderContract)
	_, err = lenderCode.EvaluateTransaction("EndLoan", "false")
	check(err)
	time.Sleep(3 * time.Second)
	fmt.Println("Ended flash loan. Refund loan: ", floan.Loan)
}
