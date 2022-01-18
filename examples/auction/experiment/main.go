package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/aungmawjj/piechain/cclib"
	"github.com/aungmawjj/piechain/contracts/eth_auction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/trie"
)

func main() {
	storage_proof()
}

const password = "password"

var ethClient *ethclient.Client
var rpcClient *rpc.Client

func storage_proof() {
	k0, err := cclib.NewTransactor("../../keys/key0", password)
	check(err)
	k1, err := cclib.NewTransactor("../../keys/key1", password)
	check(err)

	ethClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8545", "localhost"))
	check(err)

	rpcClient, err = rpc.Dial(fmt.Sprintf("http://%s:8545", "localhost"))
	check(err)

	auctionAddr, tx, _, err := eth_auction.DeployAuction(k0, ethClient)
	check(err)
	waitTx(tx, "deploy auction")

	auction, err := eth_auction.NewAuction(auctionAddr, ethClient)
	check(err)

	k1.Value = big.NewInt(1000)
	tx, err = auction.Bid(k1)
	check(err)
	waitTx(tx, "bid auction")

	var proof StorageProof
	err = rpcClient.Call(
		&proof,
		"eth_getProof",
		fmt.Sprintf("0x%x", auctionAddr),
		[]string{"0x1", "0x2"},
		"latest",
	)
	check(err)

	fmt.Println("Storage proof:")
	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "  ")
	e.Encode(proof)

	for i, sp := range proof.StorageProof {
		ok, err := VerifyEthStorageProof(&sp, proof.StorageHash)
		check(err)
		fmt.Printf("Verify result %d: %t\n", i, ok)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func waitTx(tx *types.Transaction, label string) {
	fmt.Println(label + "...")
	success, err := cclib.WaitTx(ethClient, tx.Hash())
	check(err)
	printTxStatus(success)
}

func printTxStatus(success bool) {
	if success {
		fmt.Println("Transaction successful")
	} else {
		fmt.Println("Transaction failed")
	}
}

// VerifyEthStorageProof verifies an Ethereum storage proof against the StateRoot.
// It does not verify the account proof against the Ethereum StateHash.
func VerifyEthStorageProof(proof *StorageResult, storageHash common.Hash) (bool, error) {
	var err error
	var value []byte

	if len(proof.Value) != 0 {
		value, err = rlp.EncodeToBytes(proof.Value)
		if err != nil {
			return false, err
		}
	}
	return VerifyProof(storageHash, proof.Key, value, proof.Proof)
}

// VerifyProof verifies that the path generated from key, following the nodes
// in proof leads to a leaf with value, where the hashes are correct up to the
// rootHash.
// WARNING: When the value is not found, `eth_getProof` will return "0x0" at
// the StorageProof `value` field.  In order to verify the proof of non
// existence, you must set `value` to nil, *not* the RLP encoding of 0 or null
// (which would be 0x80).
func VerifyProof(rootHash common.Hash, key []byte, value []byte, proof [][]byte) (bool, error) {
	proofDB := NewMemDB()
	for _, node := range proof {
		key := crypto.Keccak256(node)
		proofDB.Put(key, node)
	}
	path := crypto.Keccak256(key)

	res, nodes, err := trie.VerifyProof(rootHash, path, proofDB)
	if err != nil {
		return false, err
	}

	fmt.Println("value", value)
	fmt.Println("res", res)

	fmt.Println("nodes", nodes)
	return bytes.Equal(value, res), nil
}
