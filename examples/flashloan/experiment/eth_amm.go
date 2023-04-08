package main

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/Guy1m0/piechain-frontend/cclib"
	"github.com/Guy1m0/piechain-frontend/contracts/eth_arbitrage"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	decimal     = 15
	decimalB, _ = big.NewInt(0).SetString("1"+strings.Repeat("0", decimal), 10)
)

var callOpts = &bind.CallOpts{}

func TestEthAMM() {
	excT, err := cclib.NewTransactor(excKey, password)
	check(err)
	arbT, err := cclib.NewTransactor(arbitKey, password)
	check(err)

	ethClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8545", "localhost"))
	check(err)

	_, _ = excT, arbT

	valueB, _ := big.NewInt(0).SetString("1"+strings.Repeat("0", decimal+10), 10)
	token1Addr, tx, _, err := eth_arbitrage.DeployERC20(excT, ethClient, valueB)
	check(err)
	waitTx(tx, "deploy token1")

	token2Addr, tx, _, err := eth_arbitrage.DeployERC20(excT, ethClient, valueB)
	check(err)
	waitTx(tx, "deploy token2")

	token1, err := eth_arbitrage.NewERC20(token1Addr, ethClient)
	check(err)

	token2, err := eth_arbitrage.NewERC20(token2Addr, ethClient)
	check(err)

	amm1Address, tx, _, err := eth_arbitrage.DeployAMM(excT, ethClient, token1Addr, token2Addr)
	check(err)
	waitTx(tx, "deploy amm1")

	valueB, _ = big.NewInt(0).SetString("1"+strings.Repeat("0", decimal+5), 10) // 100000

	transferToken(token1, excT, amm1Address, valueB)
	transferToken(token2, excT, amm1Address, valueB)

	valueB, _ = big.NewInt(0).SetString("1"+strings.Repeat("0", decimal+3), 10) // 1000
	transferToken(token1, excT, arbT.From, valueB)

	printTokenBalance(token1, excT.From, "token1", "exchange")
	printTokenBalance(token2, excT.From, "token2", "exchange")

	printTokenBalance(token1, amm1Address, "token1", "amm1")
	printTokenBalance(token2, amm1Address, "token2", "amm1")

	printTokenBalance(token1, arbT.From, "token1", "arbt")
	printTokenBalance(token2, arbT.From, "token2", "arbt")

	amm1, err := eth_arbitrage.NewAMM(amm1Address, ethClient)
	check(err)

	tx, err = amm1.SetRate(excT, big.NewInt(2), big.NewInt(3))
	check(err)
	waitTx(tx, "set amm1 rate")

	tx, err = token1.Approve(arbT, amm1Address, valueB)
	check(err)
	waitTx(tx, "approve token1 for amm1")

	tx, err = amm1.Exchange1to2(arbT, valueB)
	check(err)
	waitTx(tx, "exchange token1 to token2")

	printTokenBalance(token1, amm1Address, "token1", "amm1")
	printTokenBalance(token2, amm1Address, "token2", "amm1")

	printTokenBalance(token1, arbT.From, "token1", "arbt")
	printTokenBalance(token2, arbT.From, "token2", "arbt")
}

func transferToken(token *eth_arbitrage.ERC20, auth *bind.TransactOpts, to common.Address, amount *big.Int) {
	tx, err := token.Transfer(auth, to, amount)
	check(err)
	waitTx(tx, "transfer token")
}

func printTokenBalance(token *eth_arbitrage.ERC20, address common.Address, tokenName, accountName string) {
	valueB, err := token.BalanceOf(callOpts, address)
	check(err)
	fmt.Printf("%s %s balance: %s\n",
		tokenName, accountName,
		big.NewInt(0).Div(valueB, decimalB).String(),
	)
}

func waitTx(tx *types.Transaction, label string) {
	fmt.Println(label + "...")
	success, err := cclib.WaitTx(ethClient, tx.Hash())
	check(err)
	printTxStatus(success)
}
