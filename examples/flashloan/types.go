package flashloan

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
)

type Flashloan struct {
	LenderContract    string
	ArbitrageContract string

	Lender      string
	Arbitrageur string
	Exchange    string

	Loan    int64
	Intrest int64
}

func (fl *Flashloan) Hash() string {
	hash := fl.Hash32()
	return common.Bytes2Hex(hash[:])
}

func (fl *Flashloan) Hash32() [32]byte {
	h := sha3.New256()
	h.Write([]byte(fl.LenderContract))
	h.Write([]byte(fl.ArbitrageContract))
	h.Write([]byte(fl.Lender))
	h.Write([]byte(fl.Arbitrageur))
	h.Write([]byte(fl.Exchange))
	h.Write([]byte(fmt.Sprint(fl.Loan)))
	h.Write([]byte(fmt.Sprint(fl.Intrest)))
	hash := h.Sum(nil)
	ret := [32]byte{}
	copy(ret[:], hash)
	return ret
}

type CommitVote struct {
	LoanHash     string
	LenderSig    string
	ArbitrageSig string
}

type SetupInfo struct {
	Token1Address common.Address
	Token2Address common.Address

	Amm1Address common.Address
	Amm2Address common.Address

	FabricTokenName string

	Exchange    common.Address
	Lender      common.Address
	Arbitrageur common.Address
}
