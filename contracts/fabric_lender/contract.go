package lender

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Flashloan struct {
	LenderContract    string
	ArbitrageContract string

	Lender    string
	Arbitrage string
	Exchange  string

	Loan    int64
	Intrest int64
}

type SmartContract struct {
	contractapi.Contract
}

const (
	KeyFlashloan = "flashloan"
	KeyLoanHash  = "loanHash"
	KeyStatus    = "status"

	Token1 = "token1"
)

func (sc *SmartContract) Setup(
	ctx contractapi.TransactionContextInterface, fljson string, loanHash string,
) error {
	err := ctx.GetStub().PutState(KeyFlashloan, []byte(fljson))
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(KeyLoanHash, []byte(loanHash))
	if err != nil {
		return err
	}
	return sc.SetStatus(ctx, 1)
}

func (sc *SmartContract) GetLoanHash(ctx contractapi.TransactionContextInterface) (string, error) {
	b, err := ctx.GetStub().GetState(KeyLoanHash)
	if err != nil {
		return "", nil
	}
	return string(b), nil
}

func (sc *SmartContract) GetFlashLoan(ctx contractapi.TransactionContextInterface) (*Flashloan, error) {
	b, err := ctx.GetStub().GetState(KeyFlashloan)
	if err != nil {
		return nil, err
	}
	var result Flashloan
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (sc *SmartContract) SetStatus(ctx contractapi.TransactionContextInterface, status int) error {
	return ctx.GetStub().PutState(KeyStatus, []byte(strconv.Itoa(status)))
}

func (sc *SmartContract) GetStatus(ctx contractapi.TransactionContextInterface) (int, error) {
	b, err := ctx.GetStub().GetState(KeyStatus)
	if err != nil {
		return 0, err
	}
	if b == nil {
		return 0, nil
	}
	return strconv.Atoi(string(b))
}

func (sc *SmartContract) Initialize(
	ctx contractapi.TransactionContextInterface, lenderSig, arbSig string,
) error {
	floan, err := sc.GetFlashLoan(ctx)
	if err != nil {
		return err
	}

	// // verify sig
	// hash, err := sc.GetLoanHash(ctx)
	// if err != nil {
	// 	return err
	// }

	// err = VerifySignature(hash, lenderSig, floan.Lender)
	// if err != nil {
	// 	return err
	// }
	// err = VerifySignature(hash, arbSig, floan.Arbitrage)
	// if err != nil {
	// 	return err
	// }

	err = sc.transferToken(ctx, Token1, floan.Lender, floan.Exchange, floan.Loan)
	if err != nil {
		return err
	}
	return sc.SetStatus(ctx, 2)
}

func (sc *SmartContract) EndLoan(
	ctx contractapi.TransactionContextInterface, success bool,
) error {
	floan, err := sc.GetFlashLoan(ctx)
	if err != nil {
		return err
	}
	refund := floan.Loan
	if success {
		refund += floan.Intrest
	}
	err = sc.transferToken(ctx, Token1, floan.Exchange, floan.Lender, refund)
	if err != nil {
		return err
	}
	return sc.SetStatus(ctx, 3)
}

func (sc *SmartContract) transferToken(
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

func VerifySignature(hash_, signature_, address_ string) error {
	hash, err := hex.DecodeString(hash_)
	if err != nil {
		return nil
	}
	signature, err := hex.DecodeString(signature_)
	if err != nil {
		return nil
	}
	address := common.HexToAddress(address_)
	pubkey, err := crypto.SigToPub(hash, signature)
	if err != nil {
		return err
	}
	recover := crypto.PubkeyToAddress(*pubkey)
	if address.Hex() != recover.Hex() {
		return fmt.Errorf("invalid signer address")
	}
	return nil
}
