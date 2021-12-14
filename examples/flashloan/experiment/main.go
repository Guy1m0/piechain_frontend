package main

import (
	"fmt"
	"time"

	"github.com/aungmawjj/piechain/examples/flashloan"
)

func main() {
	testToken()
}

func testToken() {
	token1 := flashloan.NewChaincode("token1")

	account1 := "account1"
	account2 := "account2"

	var err error
	fmt.Println("initialize token")
	_, err = token1.SubmitTransaction("SetBalance", account1, "2000")
	check(err)
	_, err = token1.SubmitTransaction("SetBalance", account2, "1000")
	check(err)
	time.Sleep(3 * time.Second)

	printBalance(token1, account1)
	printBalance(token1, account2)

	fmt.Println("transfer balance...")
	_, err = token1.SubmitTransaction("Transfer", account1, account2, "200")
	check(err)
	time.Sleep(3 * time.Second)

	printBalance(token1, account1)
	printBalance(token1, account2)
}

func printBalance(token *flashloan.Chaincode, account string) {
	b, err := token.EvaluateTransaction("GetBalance", account)
	check(err)
	fmt.Printf("%s %s balance: %s\n", token.GetName(), account, string(b))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
