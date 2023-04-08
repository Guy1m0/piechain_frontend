package main

import (
	"encoding/json"
	"time"

	"github.com/Guy1m0/piechain-frontend/contracts/eth_arbitrage"
	"github.com/Guy1m0/piechain-frontend/examples/flashloan"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func handleFlashloanInitialized(payload []byte) {
	var floan flashloan.Flashloan
	err := json.Unmarshal(payload, &floan)
	check(err)

	// transfer token to arbitrage contract
	var setupInfo flashloan.SetupInfo
	flashloan.ReadJsonFile(setupInfoFile, &setupInfo)

	token1, err := eth_arbitrage.NewERC20(setupInfo.Token1Address, ethClient)
	check(err)

	flashloan.TransferToken(
		ethClient, token1, excT, common.HexToAddress(floan.ArbitrageContract), floan.Loan,
	)

	status := listenArbitrageEnd(&floan)
	if status == 2 {
		ccsvc.Publish(flashloan.OnLoanSuccessful, payload)
	} else {
		ccsvc.Publish(flashloan.OnLoanFail, payload)
	}
}

func listenArbitrageEnd(floan *flashloan.Flashloan) uint8 {
	addr := common.HexToAddress(floan.ArbitrageContract)
	arbitrage, err := eth_arbitrage.NewArbitrage(addr, ethClient)
	check(err)

	for {
		status, err := arbitrage.Status(&bind.CallOpts{})
		check(err)
		if status >= 2 {
			return status
		}
		time.Sleep(1 * time.Second)
	}
}
