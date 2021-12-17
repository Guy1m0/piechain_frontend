package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/aungmawjj/piechain/cclib"
	"github.com/aungmawjj/piechain/contracts/eth_arbitrage"
	"github.com/aungmawjj/piechain/examples/flashloan"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	arbitKey = "../../keys/key3"
	password = "password"

	setupInfoFile  = "../setup_info.json"
	flashloanFile  = "../flash_loan.json"
	commitVoteFile = "../commit_vote.json"

	loan    = 10000
	intrest = 100
)

var (
	lenderCodeName = "lender"

	setupInfo flashloan.SetupInfo

	arbT      *bind.TransactOpts
	ethClient *ethclient.Client
)

func main() {
	command := flag.String("c", "", "command")
	flag.StringVar(&lenderCodeName, "lc", lenderCodeName, "lender code name")

	flag.Parse()

	flashloan.ReadJsonFile(setupInfoFile, &setupInfo)

	var err error
	arbT, err = cclib.NewTransactor(arbitKey, password)
	check(err)

	ethClient = flashloan.NewEthClient()

	switch *command {
	case "setup":
		setup()
	case "register":
		register()
	case "execute":
		execute()

	default:
		fmt.Println("command not found")
	}
}

func setup() {
	arbAddr, tx, _, err := eth_arbitrage.DeployArbitrage(
		arbT, ethClient,
		setupInfo.Token1Address, setupInfo.Token2Address,
		setupInfo.Amm1Address, setupInfo.Amm2Address,
	)
	check(err)
	flashloan.WaitTx(ethClient, tx, "deploy arbitrage")

	floan := flashloan.Flashloan{
		LenderContract:    lenderCodeName,
		ArbitrageContract: arbAddr.Hex(),

		Lender:      setupInfo.Lender.Hex(),
		Arbitrageur: setupInfo.Arbitrageur.Hex(),
		Exchange:    setupInfo.Exchange.Hex(),

		Loan:    loan,
		Intrest: intrest,
	}

	arbitrage, err := eth_arbitrage.NewArbitrage(arbAddr, ethClient)
	check(err)

	loanB := big.NewInt(0).Mul(big.NewInt(loan), flashloan.DecimalB)
	intrestB := big.NewInt(0).Mul(big.NewInt(intrest), flashloan.DecimalB)

	tx, err = arbitrage.Setup(arbT,
		setupInfo.Lender, setupInfo.Arbitrageur, setupInfo.Exchange,
		loanB, intrestB, floan.Hash32(),
	)
	check(err)
	flashloan.WaitTx(ethClient, tx, "setup arbitrage contract")

	lenderCode := flashloan.NewChaincode(lenderCodeName)

	fljson, _ := json.Marshal(floan)
	_, err = lenderCode.SubmitTransaction("Setup", string(fljson), floan.Hash())
	check(err)
	fmt.Println("setup lender contract...")
	time.Sleep(3 * time.Second)

	resp, err := lenderCode.EvaluateTransaction("GetStatus")
	check(err)
	status, err := strconv.Atoi(string(resp))
	check(err)
	if status == 1 {
		fmt.Println("setup successful")
	}

	fmt.Println("writing commit vote")
	arbS, err := cclib.NewSigner(arbitKey, password)
	check(err)

	loanHash := floan.Hash32()
	sig, err := arbS.Sign(loanHash[:])
	check(err)

	commitVote := flashloan.CommitVote{
		LoanHash:     floan.Hash(),
		ArbitrageSig: hex.EncodeToString(sig),
	}

	flashloan.WriteJsonFile(flashloanFile, floan)
	flashloan.WriteJsonFile(commitVoteFile, commitVote)
}

func register() {
	var floan flashloan.Flashloan
	flashloan.ReadJsonFile(flashloanFile, &floan)

	b, _ := json.Marshal(floan)

	http.Post("http://localhost:9000/flashloan",
		"application/json",
		bytes.NewReader(b),
	)

}

func execute() {
	var floan flashloan.Flashloan
	flashloan.ReadJsonFile(flashloanFile, &floan)

	arbitrage, err := eth_arbitrage.NewArbitrage(
		common.HexToAddress(floan.ArbitrageContract), ethClient,
	)
	check(err)
	tx, err := arbitrage.Execute(arbT,
		0, [32]byte{}, [32]byte{},
		0, [32]byte{}, [32]byte{},
	)
	check(err)
	fmt.Println("excute arbitrage...")
	success, err := cclib.WaitTx(ethClient, tx.Hash())
	check(err)
	flashloan.PrintTxStatus(success)
	if success {
		return
	}
	// not success
	tx, err = arbitrage.Rollback(arbT)
	check(err)
	flashloan.WaitTx(ethClient, tx, "rollback")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
