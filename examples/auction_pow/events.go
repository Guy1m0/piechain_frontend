package auction_pow

var (
	SignedAuctionResultEvent = "auction.signed_result"
	AuctionEndingEvent       = "auction.auction_ending"
)

type SignedAuctionResult struct {
	AuctionResult
	Signature []byte
}
