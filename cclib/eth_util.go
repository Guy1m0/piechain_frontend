package cclib

import (
	"context"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func WaitTx(client *ethclient.Client, hash common.Hash) (bool, error) {
	for {
		r, err := client.TransactionReceipt(context.Background(), hash)
		if err == ethereum.NotFound {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		if err != nil {
			return false, err
		}
		return r.Status == types.ReceiptStatusSuccessful, nil
	}
}

func NewTransactor(keyfile, password string) (*bind.TransactOpts, error) {
	f, err := os.Open(keyfile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	auth, err := bind.NewTransactorWithChainID(f, password, big.NewInt(15))
	if err != nil {
		return nil, err
	}
	auth.GasLimit = 100000000
	return auth, nil
}
