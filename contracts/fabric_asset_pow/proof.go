package asset_pow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

func VerifyProof(rootHash common.Hash, proof *gethclient.AccountResult) (bool, error) {
	ok, err := VerifyAccountProof(rootHash, proof)
	if !ok || err != nil {
		return ok, err
	}
	for _, sp := range proof.StorageProof {
		ok, err := VerifyStorageProof(proof.StorageHash, &sp)
		if !ok || err != nil {
			return ok, err
		}
	}
	return true, nil
}

func VerifyAccountProof(rootHash common.Hash, proof *gethclient.AccountResult) (bool, error) {
	key := proof.Address.Bytes()
	value := []interface{}{
		proof.Nonce, proof.Balance, proof.StorageHash, proof.CodeHash,
	}
	return VerifyTrieProof(rootHash, proof.AccountProof, key, value)
}

func VerifyStorageProof(storageHash common.Hash, proof *gethclient.StorageResult) (bool, error) {
	key := common.LeftPadBytes(common.FromHex(proof.Key), 32)
	value := proof.Value
	return VerifyTrieProof(storageHash, proof.Proof, key, value)
}

func VerifyTrieProof(rootHash common.Hash, proof []string, key []byte, value interface{}) (bool, error) {
	proofDB := memorydb.New()
	for _, p := range proof {
		n := common.FromHex(p)
		k := crypto.Keccak256(n)
		proofDB.Put(k, n)
	}
	path := crypto.Keccak256(key)
	v, err := trie.VerifyProof(rootHash, path, proofDB)
	if err != nil {
		return false, err
	}
	ev, err := rlp.EncodeToBytes(value)
	if err != nil {
		return false, err
	}
	return bytes.Equal(ev, v), nil
}

func VerifyEthHeaders(headers []*types.Header, base *types.Header) error {
	parent := base
	for _, header := range headers {
		err := VerifyOneEthHeader(header, parent)
		if err != nil {
			return err
		}
		parent = header
	}
	return nil
}

func VerifyOneEthHeader(header, parent *types.Header) error {
	if !bytes.Equal(header.ParentHash.Bytes(), parent.Hash().Bytes()) {
		return fmt.Errorf("invalid parent")
	}
	if header.Time <= parent.Time {
		return fmt.Errorf("old block time")
	}
	expectedDifficulty := ethash.CalcDifficulty(getChainConfig(), header.Time, parent)
	if header.Difficulty.Cmp(expectedDifficulty) != 0 {
		return fmt.Errorf("invalid difficulty: have %v, want %v", header.Difficulty, expectedDifficulty)
	}
	// Verify that the block number is parent's +1
	if diff := new(big.Int).Sub(header.Number, parent.Number); diff.Cmp(big.NewInt(1)) != 0 {
		return consensus.ErrInvalidNumber
	}
	return nil
}

func getChainConfig() *params.ChainConfig {
	s := `{
		"chainId": 15,
		"homesteadBlock": 0,
		"eip150Block": 0,
		"eip150Hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
		"eip155Block": 0,
		"eip158Block": 0,
		"byzantiumBlock": 0,
		"constantinopleBlock": 0,
		"petersburgBlock": 0,
		"ethash": {}
	  }
	`
	var config params.ChainConfig
	json.Unmarshal([]byte(s), &config)
	return &config
}
