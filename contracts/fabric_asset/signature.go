package asset

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func VerifySignature(hash, signature []byte) (address common.Address, err error) {
	pubkey, err := crypto.SigToPub(hash, signature)
	if err != nil {
		return address, err
	}
	return crypto.PubkeyToAddress(*pubkey), nil
}

func VerifyKnownSignatures(hash []byte, sigs [][]byte, addrs []common.Address) bool {
	set := make(map[string]bool, len(addrs))
	for _, addr := range addrs {
		set[addr.Hex()] = true
	}

	for _, sig := range sigs {
		addr, err := VerifySignature(hash, sig)
		if err != nil {
			return false
		}
		if !set[addr.Hex()] {
			return false
		}
	}
	return true
}
