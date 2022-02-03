package auction

import (
	"strconv"

	"golang.org/x/crypto/sha3"
)

type Asset struct {
	ID               string
	Owner            string
	PendingAuctionID int
}

type Auction struct {
	ID         int
	AssetID    string
	EthAddr    string
	QuorumAddr string

	Status string

	HighestBid         int
	HighestBidder      string
	HighestBidPlatform string
}

type StartAuctionArgs struct {
	AssetID    string
	EthAddr    string
	QuorumAddr string
	Signature  string
}

func (args *StartAuctionArgs) Hash() []byte {
	h := sha3.New256()
	h.Write([]byte(args.AssetID))
	h.Write([]byte(args.EthAddr))
	h.Write([]byte(args.QuorumAddr))
	return h.Sum(nil)
}

type AuctionResult struct {
	Platform      string
	HostAuctionID int
	AuctionAddr   string
	HighestBid    int
	HighestBidder string
}

func (args *AuctionResult) Hash() []byte {
	h := sha3.New256()
	h.Write([]byte(args.AuctionAddr))
	h.Write([]byte(args.Platform))
	h.Write([]byte(strconv.Itoa(args.HighestBid)))
	h.Write([]byte(args.HighestBidder))
	return h.Sum(nil)
}

type CrossChainAuctionResult struct {
	AuctionResult
	Signatures [][]byte
}

type FinalizeAuctionArgs struct {
	AuctionID    int
	EthResult    CrossChainAuctionResult
	QuorumResult CrossChainAuctionResult
}
