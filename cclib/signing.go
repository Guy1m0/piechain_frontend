package cclib

import (
	"io"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Signer struct {
	privkey *keystore.Key
}

func NewSigner(keyin io.Reader, password string) (*Signer, error) {
	keyjson, err := ioutil.ReadAll(keyin)
	if err != nil {
		return nil, err
	}
	privkey, err := keystore.DecryptKey(keyjson, password)
	if err != nil {
		return nil, err
	}
	return &Signer{
		privkey: privkey,
	}, nil
}

func (s *Signer) Address() common.Address {
	return s.privkey.Address
}

func (s *Signer) Sign(hash []byte) ([]byte, error) {
	return crypto.Sign(hash, s.privkey.PrivateKey)
}

func VerifySignature(hash, signature []byte) (address common.Address, err error) {
	pubkey, err := crypto.SigToPub(hash, signature)
	if err != nil {
		return address, err
	}
	return crypto.PubkeyToAddress(*pubkey), nil
}

func VerifyKnownSignatures(hash []byte, sigs [][]byte, addrs []common.Address) bool {
	set := make(map[string]struct{}, len(addrs))
	for _, addr := range addrs {
		set[addr.Hex()] = struct{}{}
	}
	for _, sig := range sigs {
		addr, err := VerifySignature(hash, sig)
		if err != nil {
			return false
		}
		_, ok := set[addr.Hex()]
		if !ok {
			return false
		}
	}
	return true
}
