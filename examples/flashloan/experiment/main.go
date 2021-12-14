package main

import (
	"fmt"
	"time"

	"github.com/aungmawjj/piechain/examples/flashloan"
)

func main() {
	testToken()
	testAMM()
}

func testToken() {
	fmt.Println("testing token")
	token1 := flashloan.NewChaincode("token1")

	account1 := "account1"
	account2 := "account2"

	var err error
	_, err = token1.SubmitTransaction("SetBalance", account1, "2000")
	check(err)
	_, err = token1.SubmitTransaction("SetBalance", account2, "1000")
	check(err)
	time.Sleep(3 * time.Second)

	printBalance(token1, account1)
	printBalance(token1, account2)

	fmt.Println("transfer balance 200")
	_, err = token1.SubmitTransaction("Transfer", account1, account2, "200")
	check(err)
	time.Sleep(3 * time.Second)

	printBalance(token1, account1)
	printBalance(token1, account2)
}

func testAMM() {
	fmt.Println("testing amm exchange")

	token1 := flashloan.NewChaincode("token1")
	token2 := flashloan.NewChaincode("token2")
	amm1 := flashloan.NewChaincode("amm1")

	alice := "alice"
	amm1Acc := "amm1"

	var err error

	_, err = token1.SubmitTransaction("SetBalance", amm1Acc, "10000")
	check(err)
	_, err = token2.SubmitTransaction("SetBalance", amm1Acc, "10000")
	check(err)
	_, err = token1.SubmitTransaction("SetBalance", alice, "2000")
	check(err)

	_, err = amm1.SubmitTransaction("SetAccount", amm1Acc)
	check(err)
	_, err = amm1.SubmitTransaction("SetRate", "1.4")
	check(err)

	time.Sleep(3 * time.Second)

	printBalance(token1, amm1Acc)
	printBalance(token2, amm1Acc)
	printBalance(token1, alice)
	printBalance(token2, alice)

	printAMMRate(amm1)

	fmt.Println("exchange 1000 token1 -> token2")
	_, err = amm1.SubmitTransaction("Exchange", alice, "1", "1000")
	check(err)

	time.Sleep(3 * time.Second)

	printBalance(token1, amm1Acc)
	printBalance(token2, amm1Acc)
	printBalance(token1, alice)
	printBalance(token2, alice)
}

func printBalance(token *flashloan.Chaincode, account string) {
	b, err := token.EvaluateTransaction("GetBalance", account)
	check(err)
	fmt.Printf("%s %s balance: %s\n", token.GetName(), account, string(b))
}

func printAMMRate(amm *flashloan.Chaincode) {
	b, err := amm.EvaluateTransaction("GetRate")
	check(err)
	fmt.Printf("%s exchange rate token1 -> token2 : %s\n", amm.GetName(), string(b))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
