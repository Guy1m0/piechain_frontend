package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/aungmawjj/piechain/cclib"
	"github.com/aungmawjj/piechain/contracts/eth_erc20"
	"github.com/aungmawjj/piechain/contracts/eth_lender"
	"github.com/aungmawjj/piechain/examples/flashloan"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	excKey    = "../../keys/key0"
	lenderKey = "../../keys/key1"
	arbitKey  = "../../keys/key2"
	password  = "password"

	token1Name = "token1"
	token2Name = "token2"
	amm1Name   = "amm1"
	amm2Name   = "amm2"
	arbitName  = "arbitrage"

	amm1Account = "amm1"
	amm2Account = "amm2"
)

var ethClient *ethclient.Client

func main() {
	// testToken()
	// testAMM()
	testFullApp()
}

func testFullApp() {
	var err error

	ethClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8546", "localhost"))
	check(err)

	excT, err := cclib.NewTransactor(excKey, password)
	check(err)

	fmt.Println("deploy erc20 token")
	tokenAddr, tx, _, err := eth_erc20.DeployERC20(excT, ethClient, big.NewInt(1_000_000_000))
	check(err)

	success, err := cclib.WaitTx(ethClient, tx.Hash())
	check(err)
	printTxStatus(success)

	fmt.Println("transfer to lender")
	token, err := eth_erc20.NewERC20(tokenAddr, ethClient)
	check(err)

	lenderT, err := cclib.NewTransactor(lenderKey, password)
	check(err)

	arbitT, err := cclib.NewTransactor(arbitKey, password)
	check(err)

	tx, err = token.Transfer(excT, lenderT.From, big.NewInt(10_000))
	check(err)

	success, err = cclib.WaitTx(ethClient, tx.Hash())
	check(err)
	printTxStatus(success)

	balbig, err := token.BalanceOf(&bind.CallOpts{}, lenderT.From)
	check(err)
	fmt.Println("lender balance: ", balbig.Int64())

	token1 := flashloan.NewChaincode(token1Name)
	token2 := flashloan.NewChaincode(token2Name)
	amm1 := flashloan.NewChaincode(amm1Name)
	amm2 := flashloan.NewChaincode(amm2Name)

	_, err = amm1.SubmitTransaction("SetAccount", amm1Account)
	check(err)
	_, err = amm1.SubmitTransaction("SetRate", "1.5")
	check(err)

	_, err = amm2.SubmitTransaction("SetAccount", amm2Account)
	check(err)
	_, err = amm2.SubmitTransaction("SetRate", "1")
	check(err)

	_, err = token1.SubmitTransaction("SetBalance", amm1Account, "10000000")
	check(err)
	_, err = token2.SubmitTransaction("SetBalance", amm1Account, "15000000")
	check(err)

	_, err = token1.SubmitTransaction("SetBalance", amm2Account, "11000000")
	check(err)
	_, err = token2.SubmitTransaction("SetBalance", amm2Account, "10000000")
	check(err)

	_, err = token1.SubmitTransaction("SetBalance", excT.From.Hex(), "10000000")
	check(err)

	_, err = token1.SubmitTransaction("SetBalance", arbitT.From.Hex(), "0")
	check(err)
	time.Sleep(3 * time.Second)

	printBalance(token1, amm1Account)
	printBalance(token2, amm1Account)

	printBalance(token1, amm2Account)
	printBalance(token2, amm2Account)

	fmt.Println("exchange balance")
	printBalance(token1, excT.From.Hex())

	fmt.Println("arbitrage balance")
	printBalance(token1, arbitT.From.Hex())

	printAMMRate(amm1)
	printAMMRate(amm2)

	fmt.Println("setup")

	lenderAddr, tx, _, err := eth_lender.DeployLender(arbitT, ethClient, tokenAddr)
	check(err)
	success, err = cclib.WaitTx(ethClient, tx.Hash())
	check(err)

	printTxStatus(success)

	floan := flashloan.Flashloan{
		LenderContract:    lenderAddr.Hex(),
		ArbitrageContract: arbitName,
		Lender:            lenderT.From.Hex(),
		Arbitrage:         arbitT.From.Hex(),
		Exchange:          excT.From.Hex(),
		Loan:              10000,
		Intrest:           20,
	}
	fljson, _ := json.Marshal(floan)

	lenderContract, err := eth_lender.NewLender(lenderAddr, ethClient)
	check(err)

	tokenAddr_, err := lenderContract.TokenAddress(&bind.CallOpts{})
	check(err)
	fmt.Println("token address: ", tokenAddr_.Hex())

	tx, err = lenderContract.Setup(
		arbitT, lenderT.From, arbitT.From, excT.From,
		big.NewInt(floan.Loan), big.NewInt(floan.Intrest), [32]byte{},
	)
	check(err)
	success, err = cclib.WaitTx(ethClient, tx.Hash())
	check(err)

	printTxStatus(success)

	arbitrage := flashloan.NewChaincode(arbitName)

	_, err = arbitrage.SubmitTransaction("SetAccount", arbitName)
	check(err)

	_, err = arbitrage.SubmitTransaction("Setup", string(fljson), floan.Hash())
	check(err)
	time.Sleep(3 * time.Second)

	fmt.Println("initialize")

	tx, err = token.Approve(lenderT, lenderAddr, big.NewInt(floan.Loan))
	check(err)
	success, err = cclib.WaitTx(ethClient, tx.Hash())
	check(err)
	printTxStatus(success)

	allowance, err := token.Allowance(&bind.CallOpts{}, lenderT.From, lenderAddr)
	check(err)
	fmt.Println("Allowance: ", allowance.Int64())

	tx, err = lenderContract.Initialize(lenderT, 0, [32]byte{}, [32]byte{}, 0, [32]byte{}, [32]byte{})
	check(err)
	success, err = cclib.WaitTx(ethClient, tx.Hash())
	check(err)
	printTxStatus(success)

	balbig, err = token.BalanceOf(&bind.CallOpts{}, lenderT.From)
	check(err)
	fmt.Println("lender balance: ", balbig.Int64())

	_, err = arbitrage.SubmitTransaction("FeedLoan")
	check(err)
	time.Sleep(3 * time.Second)

	fmt.Println("exchange balance")
	printBalance(token1, excT.From.Hex())

	fmt.Println("arbitrage balance")
	printBalance(token1, arbitT.From.Hex())

	printBalance(token1, arbitName)

	_, err = arbitrage.SubmitTransaction("Execute", "", "")
	check(err)
	time.Sleep(3 * time.Second)

	fmt.Println("exchange balance")
	printBalance(token1, excT.From.Hex())

	fmt.Println("arbitrage balance")
	printBalance(token1, arbitT.From.Hex())

	fmt.Println("end loan")

	tx, err = token.Approve(excT, lenderAddr, big.NewInt(floan.Loan+floan.Intrest))
	check(err)
	success, err = cclib.WaitTx(ethClient, tx.Hash())
	check(err)
	printTxStatus(success)

	allowance, err = token.Allowance(&bind.CallOpts{}, excT.From, lenderAddr)
	check(err)
	fmt.Println("Allowance: ", allowance.Int64())

	tx, err = lenderContract.EndLoan(excT, true)
	check(err)
	success, err = cclib.WaitTx(ethClient, tx.Hash())
	check(err)
	printTxStatus(success)

	balbig, err = token.BalanceOf(&bind.CallOpts{}, lenderT.From)
	check(err)
	fmt.Println("lender balance: ", balbig.Int64())

	fmt.Println("arbitrage balance")
	printBalance(token1, arbitT.From.Hex())

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

func printTxStatus(success bool) {
	if success {
		fmt.Println("Transaction successful")
	} else {
		fmt.Println("Transaction failed")
	}
}
