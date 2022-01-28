package auction_pow

import (
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

func (this *StartAuctionArgs) Hash() []byte {
	h := sha3.New256()
	h.Write([]byte(this.AssetID))
	h.Write([]byte(this.EthAddr))
	h.Write([]byte(this.QuorumAddr))
	return h.Sum(nil)
}

type AuctionResult struct {
	HostAuctionID int
	AuctionAddr   string
	HighestBid    int
	HighestBidder string
}
