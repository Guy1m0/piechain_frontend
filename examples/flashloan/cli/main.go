package main

import (
	"flag"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/aungmawjj/piechain/cclib"
	"github.com/aungmawjj/piechain/contracts/eth_arbitrage"
	"github.com/aungmawjj/piechain/examples/flashloan"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	rootKey   = "../../keys/key0"
	excKey    = "../../keys/key1"
	lenderKey = "../../keys/key2"
	arbitKey  = "../../keys/key3"
	password  = "password"

	fabricTokenName = "token1"

	setupInfoFile = "../setup_info.json"
)

var (
	rootT   *bind.TransactOpts
	excT    *bind.TransactOpts
	lenderT *bind.TransactOpts
	arbT    *bind.TransactOpts

	ethClient *ethclient.Client
)

func main() {
	var err error
	rootT, err = cclib.NewTransactor(rootKey, password)
	check(err)
	excT, err = cclib.NewTransactor(excKey, password)
	check(err)
	lenderT, err = cclib.NewTransactor(lenderKey, password)
	check(err)
	arbT, err = cclib.NewTransactor(arbitKey, password)
	check(err)

	ethClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8546", "localhost"))
	check(err)

	command := flag.String("c", "", "command")
	flag.Parse()

	switch *command {
	case "setup":
		setup()
	case "display":
		display()
	case "sameRate":
		sameRate()

	default:
		fmt.Println("command not found")
	}

}

func setup() {
	fmt.Println("Ethereum setup")
	supply, _ := big.NewInt(0).SetString("1"+strings.Repeat("0", flashloan.Decimal+10), 10)
	token1Addr, tx, _, err := eth_arbitrage.DeployERC20(rootT, ethClient, supply)
	check(err)
	flashloan.WaitTx(ethClient, tx, "deploy token1")

	token2Addr, tx, _, err := eth_arbitrage.DeployERC20(rootT, ethClient, supply)
	check(err)
	flashloan.WaitTx(ethClient, tx, "deploy token2")

	token1, err := eth_arbitrage.NewERC20(token1Addr, ethClient)
	check(err)

	token2, err := eth_arbitrage.NewERC20(token2Addr, ethClient)
	check(err)

	amm1Address, tx, _, err := eth_arbitrage.DeployAMM(rootT, ethClient, token1Addr, token2Addr)
	check(err)
	flashloan.WaitTx(ethClient, tx, "deploy amm1")

	amm2Address, tx, _, err := eth_arbitrage.DeployAMM(rootT, ethClient, token1Addr, token2Addr)
	check(err)
	flashloan.WaitTx(ethClient, tx, "deploy amm2")

	supply, _ = big.NewInt(0).SetString("1"+strings.Repeat("0", flashloan.Decimal+6), 10) // 1000000

	flashloan.TransferToken(ethClient, token1, rootT, excT.From, supply)
	flashloan.TransferToken(ethClient, token2, rootT, excT.From, supply)

	flashloan.TransferToken(ethClient, token1, rootT, amm1Address, supply)
	flashloan.TransferToken(ethClient, token2, rootT, amm1Address, supply)

	flashloan.TransferToken(ethClient, token1, rootT, amm2Address, supply)
	flashloan.TransferToken(ethClient, token2, rootT, amm2Address, supply)

	flashloan.PrintTokenBalance(token1, excT.From, "token1", "exchange")
	flashloan.PrintTokenBalance(token2, excT.From, "token2", "exchange")

	flashloan.PrintTokenBalance(token1, amm1Address, "token1", "amm1")
	flashloan.PrintTokenBalance(token2, amm1Address, "token2", "amm1")

	flashloan.PrintTokenBalance(token1, amm2Address, "token1", "amm2")
	flashloan.PrintTokenBalance(token2, amm2Address, "token2", "amm2")

	flashloan.PrintTokenBalance(token1, arbT.From, "token1", "arbt")
	flashloan.PrintTokenBalance(token2, arbT.From, "token2", "arbt")

	amm1, err := eth_arbitrage.NewAMM(amm1Address, ethClient)
	check(err)
	amm2, err := eth_arbitrage.NewAMM(amm2Address, ethClient)
	check(err)

	tx, err = amm1.SetRate(excT, big.NewInt(2), big.NewInt(3))
	check(err)
	flashloan.WaitTx(ethClient, tx, "set amm1 rate")

	tx, err = amm2.SetRate(excT, big.NewInt(1), big.NewInt(1))
	check(err)
	flashloan.WaitTx(ethClient, tx, "set amm2 rate")

	fmt.Println("Fabric setup")
	fabricToken := flashloan.NewChaincode(fabricTokenName)

	_, err = fabricToken.SubmitTransaction("SetBalance", excT.From.Hex(), "1000000")
	check(err)
	_, err = fabricToken.SubmitTransaction("SetBalance", lenderT.From.Hex(), "10000")
	check(err)
	fmt.Println(3 * time.Second)

	flashloan.PrintFabricBalance(fabricToken, excT.From.Hex(), "exchange")
	flashloan.PrintFabricBalance(fabricToken, lenderT.From.Hex(), "lender")

	flashloan.WriteJsonFile(setupInfoFile, flashloan.SetupInfo{
		Token1Address: token1Addr,
		Token2Address: token2Addr,

		Amm1Address: amm1Address,
		Amm2Address: amm2Address,

		FabricTokenName: fabricTokenName,

		Exchange:    excT.From,
		Lender:      lenderT.From,
		Arbitrageur: arbT.From,
	})
}

func display() {
	var setupInfo flashloan.SetupInfo
	flashloan.ReadJsonFile(setupInfoFile, &setupInfo)

	token1, err := eth_arbitrage.NewERC20(setupInfo.Token1Address, ethClient)
	check(err)

	token2, err := eth_arbitrage.NewERC20(setupInfo.Token2Address, ethClient)
	check(err)

	flashloan.PrintTokenBalance(token1, excT.From, "eth token1", "exchange")
	flashloan.PrintTokenBalance(token2, excT.From, "eth token2", "exchange")

	flashloan.PrintTokenBalance(token1, arbT.From, "eth token1", "arbitrageur")
	flashloan.PrintTokenBalance(token2, arbT.From, "eth token2", "arbitrageur")

	fmt.Println("Fabric setup")
	fabricToken := flashloan.NewChaincode(fabricTokenName)

	flashloan.PrintFabricBalance(fabricToken, excT.From.Hex(), "exchange")
	flashloan.PrintFabricBalance(fabricToken, lenderT.From.Hex(), "lender")
}

func sameRate() {

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
