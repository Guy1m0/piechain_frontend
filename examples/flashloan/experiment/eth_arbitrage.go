package main

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/Guy1m0/piechain-frontend/cclib"
	"github.com/Guy1m0/piechain-frontend/contracts/eth_arbitrage"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestEthArbitrage() {
	excT, err := cclib.NewTransactor(excKey, password)
	check(err)
	lenderT, err := cclib.NewTransactor(lenderKey, password)
	check(err)
	arbT, err := cclib.NewTransactor(arbitKey, password)
	check(err)

	ethClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8545", "localhost"))
	check(err)

	valueB, _ := big.NewInt(0).SetString("1"+strings.Repeat("0", decimal+6), 10)
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

	amm2Address, tx, _, err := eth_arbitrage.DeployAMM(excT, ethClient, token1Addr, token2Addr)
	check(err)
	waitTx(tx, "deploy amm2")

	valueB, _ = big.NewInt(0).SetString("1"+strings.Repeat("0", decimal+5), 10) // 100000

	transferToken(token1, excT, amm1Address, valueB)
	transferToken(token2, excT, amm1Address, valueB)

	transferToken(token1, excT, amm2Address, valueB)
	transferToken(token2, excT, amm2Address, valueB)

	printTokenBalance(token1, excT.From, "token1", "exchange")
	printTokenBalance(token2, excT.From, "token2", "exchange")

	printTokenBalance(token1, amm1Address, "token1", "amm1")
	printTokenBalance(token2, amm1Address, "token2", "amm1")

	printTokenBalance(token1, amm2Address, "token1", "amm2")
	printTokenBalance(token2, amm2Address, "token2", "amm2")

	printTokenBalance(token1, arbT.From, "token1", "arbt")
	printTokenBalance(token2, arbT.From, "token2", "arbt")

	amm1, err := eth_arbitrage.NewAMM(amm1Address, ethClient)
	check(err)
	amm2, err := eth_arbitrage.NewAMM(amm2Address, ethClient)
	check(err)

	tx, err = amm1.SetRate(excT, big.NewInt(2), big.NewInt(3))
	check(err)
	waitTx(tx, "set amm1 rate")

	tx, err = amm2.SetRate(excT, big.NewInt(1), big.NewInt(1))
	check(err)
	waitTx(tx, "set amm2 rate")

	arbAddr, tx, _, err := eth_arbitrage.DeployArbitrage(
		arbT, ethClient,
		token1Addr, token2Addr, amm1Address, amm2Address,
	)
	check(err)
	waitTx(tx, "deploy arbitrage")

	arbitrage, err := eth_arbitrage.NewArbitrage(arbAddr, ethClient)
	check(err)

	loan, _ := big.NewInt(0).SetString("1"+strings.Repeat("0", decimal+3), 10)    // 1000
	intrest, _ := big.NewInt(0).SetString("2"+strings.Repeat("0", decimal+1), 10) // 20
	tx, err = arbitrage.Setup(arbT, lenderT.From, arbT.From, excT.From, loan, intrest, [32]byte{})
	check(err)
	waitTx(tx, "setup arbitrage")

	tx, err = token1.Transfer(excT, arbAddr, loan)
	check(err)
	waitTx(tx, "transfer loan to arbitrage")

	tx, err = arbitrage.Execute(
		arbT,
		0, [32]byte{}, [32]byte{},
		0, [32]byte{}, [32]byte{},
	)
	check(err)
	waitTx(tx, "execute arbitrage")

	printTokenBalance(token1, excT.From, "token1", "exchange")
	printTokenBalance(token2, excT.From, "token2", "exchange")

	printTokenBalance(token1, arbT.From, "token1", "arbt")
	printTokenBalance(token2, arbT.From, "token2", "arbt")

}
