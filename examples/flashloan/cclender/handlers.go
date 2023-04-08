package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/Guy1m0/piechain-frontend/examples/flashloan"
)

func handleFlashloanSuccessful(payload []byte) {
	var floan flashloan.Flashloan
	err := json.Unmarshal(payload, &floan)
	check(err)

	lenderCode := flashloan.NewChaincode(floan.LenderContract)
	_, err = lenderCode.SubmitTransaction("EndLoan", "true")
	check(err)
	fmt.Println("Ending flash loan. Refund loan + intrest: ", floan.Loan+floan.Intrest)
	time.Sleep(3 * time.Second)
	confirmEndLoan(lenderCode)
}

func handleFlashloanFail(payload []byte) {
	var floan flashloan.Flashloan
	err := json.Unmarshal(payload, &floan)
	check(err)

	lenderCode := flashloan.NewChaincode(floan.LenderContract)
	_, err = lenderCode.SubmitTransaction("EndLoan", "false")
	check(err)
	fmt.Println("Ending flash loan. Refund loan: ", floan.Loan)
	time.Sleep(3 * time.Second)
	confirmEndLoan(lenderCode)
}

func confirmEndLoan(lenderCode *flashloan.Chaincode) {
	resp, err := lenderCode.EvaluateTransaction("GetStatus")
	check(err)

	status, err := strconv.Atoi(string(resp))
	check(err)

	if status == 3 {
		fmt.Println("ended loan")
	}
}
