package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/aungmawjj/piechain/cclib"
	"github.com/aungmawjj/piechain/examples/flashloan"
)

const (
	lenderKey = "../../keys/key2"
	password  = "password"

	setupInfoFile  = "../setup_info.json"
	flashloanFile  = "../flash_loan.json"
	commitVoteFile = "../commit_vote.json"
)

var (
	lenderS *cclib.Signer
)

func main() {
	var err error

	lenderS, err = cclib.NewSigner(lenderKey, password)
	check(err)

	command := flag.String("c", "", "command")
	flag.Parse()

	switch *command {
	case "sign":
		sign()
	case "initialize":
		initialize()

	default:
		fmt.Println("command not found")
	}
}

func sign() {
	var commitVote flashloan.CommitVote
	flashloan.ReadJsonFile(commitVoteFile, &commitVote)

	hash, err := hex.DecodeString(commitVote.LoanHash)
	check(err)
	sig, err := lenderS.Sign(hash)
	check(err)

	commitVote.LenderSig = hex.EncodeToString(sig)

	flashloan.WriteJsonFile(commitVoteFile, commitVote)
}

func initialize() {
	var setupInfo flashloan.SetupInfo
	flashloan.ReadJsonFile(setupInfoFile, &setupInfo)

	var floan flashloan.Flashloan
	flashloan.ReadJsonFile(flashloanFile, &floan)

	var commitVote flashloan.CommitVote
	flashloan.ReadJsonFile(commitVoteFile, &commitVote)

	lenderCode := flashloan.NewChaincode(floan.LenderContract)

	_, err := lenderCode.SubmitTransaction("Initialize", commitVote.LenderSig, commitVote.ArbitrageSig)
	check(err)
	fmt.Println("initialize lender contract...")
	time.Sleep(3 * time.Second)

	resp, err := lenderCode.EvaluateTransaction("GetStatus")
	check(err)
	status, err := strconv.Atoi(string(resp))
	check(err)
	if status == 2 {
		fmt.Println("initialize successful")
	}
	fabricToken := flashloan.NewChaincode(setupInfo.FabricTokenName)
	flashloan.PrintFabricBalance(fabricToken, lenderS.Address().Hex(), "lender")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
