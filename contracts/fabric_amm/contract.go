package amm

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

const (
	KeyExchangeRate = "exchange_rate"
	KeyMyAccount    = "my_account"
)

func (sc *SmartContract) SetRate(ctx contractapi.TransactionContextInterface, rate float64) error {
	b, _ := json.Marshal(rate)
	return ctx.GetStub().PutState(KeyExchangeRate, b)
}

func (sc *SmartContract) GetRate(ctx contractapi.TransactionContextInterface) (float64, error) {
	b, err := ctx.GetStub().GetState(KeyExchangeRate)
	if err != nil {
		return 0, nil
	}
	var rate float64
	err = json.Unmarshal(b, &rate)
	return rate, err
}

func (sc *SmartContract) SetAccount(ctx contractapi.TransactionContextInterface, account string) error {
	return ctx.GetStub().PutState(KeyMyAccount, []byte(account))
}

func (sc *SmartContract) GetAccount(ctx contractapi.TransactionContextInterface) (string, error) {
	b, err := ctx.GetStub().GetState(KeyMyAccount)
	return string(b), err
}

func (sc *SmartContract) Exchange(
	ctx contractapi.TransactionContextInterface, sender string, exctype int, amount int64,
) (int64, error) {
	myAccount, err := sc.GetAccount(ctx)
	if err != nil {
		return 0, err
	}

	rate, err := sc.GetRate(ctx)
	if err != nil {
		return 0, err
	}
	tokenFrom, tokenTo, resultAmount := sc.getExchangeInfo(exctype, amount, rate)

	err = sc.TransferToken(ctx, tokenFrom, sender, myAccount, amount)
	if err != nil {
		return 0, err
	}
	err = sc.TransferToken(ctx, tokenTo, myAccount, sender, resultAmount)
	if err != nil {
		return 0, err
	}
	return resultAmount, nil
}

func (sc *SmartContract) TransferToken(
	ctx contractapi.TransactionContextInterface,
	tokenName, sender, receiver string, amount int64,
) error {
	resp := ctx.GetStub().InvokeChaincode(tokenName, [][]byte{
		[]byte("Transfer"),
		[]byte(sender),
		[]byte(receiver),
		[]byte(fmt.Sprint(amount)),
	}, "mychannel")
	if resp.Status != shim.OK {
		return fmt.Errorf("failed to transfer %s, %s", tokenName, resp.Message)
	}
	return nil
}

func (sc *SmartContract) getExchangeInfo(
	excType int, amount int64, rate float64,
) (tokenFrom string, tokenTo string, resultAmount int64) {
	if excType == 1 {
		tokenFrom = "token1"
		tokenTo = "token2"
		resultAmount = int64(float64(amount) * rate)
	} else {
		tokenFrom = "token2"
		tokenTo = "token1"
		resultAmount = int64(float64(amount) / rate)
	}
	return
}
