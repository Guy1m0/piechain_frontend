package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	excKey    = "../../keys/key0"
	lenderKey = "../../keys/key1"
	arbitKey  = "../../keys/key2"
	password  = "password"
)

var ethClient *ethclient.Client

func main() {
	// TestFabricToken()
	// TestEthAMM()
	TestEthArbitrage()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func printTxStatus(success bool) {
	if success {
		fmt.Println("Transaction successful")
	} else {
		fmt.Println("Transaction failed")
	}
}
