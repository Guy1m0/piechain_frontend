package auction_pow

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
)

type Asset struct {
	ID               string
	Owner            string
	PendingAuctionID int
}

type Auction struct {
	ID      int
	AssetID string
	Status  string

	BaseEthHeader *types.Header
	AuctionAddr   string

	HighestBid    int
	HighestBidder string

	ProvableResult *gethclient.AccountResult
	EthHeaders     []*types.Header
}

type StartAuctionArgs struct {
	AssetID     string
	AuctionAddr string
	EthHeader   *types.Header
}

type EndAuctionArgs struct {
	AuctionID int

	ProvableResult *gethclient.AccountResult
	EthHeaders     []*types.Header
}
