package arbitrage

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

	KeyAccount = "account"

	Token1 = "token1"
)

func (sc *SmartContract) Setup(
	ctx contractapi.TransactionContextInterface, fljson string, loanHash string,
) error {
	err := ctx.GetStub().PutState(KeyFlashloan, []byte(fljson))
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(KeyLoanHash, []byte(loanHash))
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

func (sc *SmartContract) SetAccount(ctx contractapi.TransactionContextInterface, account string) error {
	return ctx.GetStub().PutState(KeyAccount, []byte(account))
}

func (sc *SmartContract) GetAccount(ctx contractapi.TransactionContextInterface) (string, error) {
	b, err := ctx.GetStub().GetState(KeyAccount)
	return string(b), err
}

func (sc *SmartContract) FeedLoan(ctx contractapi.TransactionContextInterface) error {
	flashloan, err := sc.GetFlashLoan(ctx)
	if err != nil {
		return err
	}
	myAccount, err := sc.GetAccount(ctx)
	if err != nil {
		return err
	}
	err = sc.transferToken(ctx, Token1, flashloan.Exchange, myAccount, flashloan.Loan)
	if err != nil {
		return err
	}
	return sc.SetStatus(ctx, 1)
}

func (sc *SmartContract) Execute(
	ctx contractapi.TransactionContextInterface, lenderSig, arbitSig string,
) error {
	myAccount, err := sc.GetAccount(ctx)
	if err != nil {
		return err
	}
	flashloan, err := sc.GetFlashLoan(ctx)
	if err != nil {
		return err
	}
	// loanHash, err := sc.GetLoanHash(ctx)
	// if err != nil {
	// 	return err
	// }

	// err = VerifySignature(loanHash, lenderSig, flashloan.Lender)
	// if err != nil {
	// 	return err
	// }

	// err = VerifySignature(loanHash, arbitSig, flashloan.Arbitrage)
	// if err != nil {
	// 	return err
	// }

	sc.transferToken(ctx, Token1, myAccount, flashloan.Arbitrage, flashloan.Loan)

	token2Amount, err := sc.exchange(ctx, "amm1", flashloan.Arbitrage, 1, flashloan.Loan)
	if err != nil {
		return err
	}
	token1Amount, err := sc.exchange(ctx, "amm2", flashloan.Arbitrage, 2, token2Amount)
	if err != nil {
		return err
	}
	if token1Amount < flashloan.Loan+flashloan.Intrest {
		return fmt.Errorf("not enough profit")
	}

	err = sc.transferToken(
		ctx, Token1, flashloan.Arbitrage, flashloan.Exchange,
		flashloan.Loan+flashloan.Intrest,
	)
	if err != nil {
		return err
	}

	return sc.SetStatus(ctx, 2)
}

func (sc *SmartContract) exchange(
	ctx contractapi.TransactionContextInterface,
	amm string, sender string, excType int, amount int64,
) (int64, error) {
	resp := ctx.GetStub().InvokeChaincode(amm, [][]byte{
		[]byte("Exchange"),
		[]byte(sender),
		[]byte(fmt.Sprint(excType)),
		[]byte(fmt.Sprint(amount)),
	}, "mychannel")
	if resp.Status != shim.OK {
		return 0, fmt.Errorf("failed to exchange %s", amm)
	}
	var result int64
	err := json.Unmarshal(resp.Payload, &result)
	return result, err
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
		return fmt.Errorf("failed to transfer %s", tokenName)
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
