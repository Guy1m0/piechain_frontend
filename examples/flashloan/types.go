package flashloan

import (
	"fmt"

	"golang.org/x/crypto/sha3"
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

func (fl *Flashloan) Hash() string {
	hash := fl.Hash32()
	return string(hash[:])
}

func (fl *Flashloan) Hash32() [32]byte {
	h := sha3.New256()
	h.Write([]byte(fl.LenderContract))
	h.Write([]byte(fl.ArbitrageContract))
	h.Write([]byte(fl.Lender))
	h.Write([]byte(fl.Arbitrage))
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
