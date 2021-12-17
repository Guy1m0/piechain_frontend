package flashloan

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/aungmawjj/piechain/cclib"
	"github.com/aungmawjj/piechain/contracts/eth_arbitrage"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	Decimal     = 15
	DecimalB, _ = big.NewInt(0).SetString("1"+strings.Repeat("0", Decimal), 10)
)

func PrintFabricBalance(token *Chaincode, account string, label string) {
	b, err := token.EvaluateTransaction("GetBalance", account)
	check(err)
	fmt.Printf("fabtic %s %s balance: %s\n", token.GetName(), label, string(b))
}

func TransferToken(client *ethclient.Client, token *eth_arbitrage.ERC20, auth *bind.TransactOpts, to common.Address, amount int64) {
	tx, err := token.Transfer(auth, to, big.NewInt(0).Mul(big.NewInt(amount), DecimalB))
	check(err)
	WaitTx(client, tx, "transfer token")
}

func PrintTokenBalance(token *eth_arbitrage.ERC20, address common.Address, tokenName, accountName string) {
	valueB, err := token.BalanceOf(&bind.CallOpts{}, address)
	check(err)
	fmt.Printf("%s %s balance: %s\n",
		tokenName, accountName,
		big.NewInt(0).Div(valueB, DecimalB).String(),
	)
}

func WaitTx(client *ethclient.Client, tx *types.Transaction, label string) {
	fmt.Println(label + "...")
	success, err := cclib.WaitTx(client, tx.Hash())
	check(err)
	PrintTxStatus(success)
}

func PrintTxStatus(success bool) {
	if success {
		fmt.Println("Transaction successful")
	} else {
		fmt.Println("Transaction failed")
	}
}

func ReadJsonFile(filepath string, val interface{}) {
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()

	d := json.NewDecoder(f)
	err = d.Decode(val)
	check(err)
}

func WriteJsonFile(filepath string, val interface{}) {
	f, err := os.Create(filepath)
	check(err)
	defer f.Close()

	e := json.NewEncoder(f)
	e.SetIndent("", "  ")
	err = e.Encode(val)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
