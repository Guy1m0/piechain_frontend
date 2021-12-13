package fabric_asset

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

const (
	KeyAssets        = "assets"
	KeyAuctions      = "auctions"
	KeyLastAuctionID = "lastAuction"
)

func (cc *SmartContract) AddAsset(
	ctx contractapi.TransactionContextInterface, id, owner string,
) error {
	asset := Asset{
		ID:    id,
		Owner: owner,
	}
	return cc.setAsset(ctx, &asset)
}

func (cc *SmartContract) StartAuction(
	ctx contractapi.TransactionContextInterface, argjson string,
) error {
	var args StartAuctionArgs
	err := json.Unmarshal([]byte(argjson), &args)
	if err != nil {
		return err
	}

	asset, err := cc.GetAsset(ctx, args.AssetID)
	if err != nil {
		return err
	}
	if asset.PendingAuctionID > 0 {
		return fmt.Errorf("pending auction on asset")
	}

	lastID, err := cc.GetLastAuctionID(ctx)
	if err != nil {
		return err
	}
	auction := Auction{
		ID:         lastID + 1,
		AssetID:    args.AssetID,
		EthAddr:    args.EthAddr,
		QuorumAddr: args.QuorumAddr,
		Status:     "Started",
	}
	err = cc.setAuction(ctx, &auction)
	if err != nil {
		return err
	}
	err = cc.setLastAuctionID(ctx, auction.ID)
	if err != nil {
		return err
	}

	asset.PendingAuctionID = auction.ID
	return cc.setAsset(ctx, asset)
}

func (cc *SmartContract) EndAuction(
	ctx contractapi.TransactionContextInterface, assetID string,
) error {
	asset, err := cc.GetAsset(ctx, assetID)
	if err != nil {
		return err
	}
	auction, err := cc.GetAuction(ctx, asset.PendingAuctionID)
	if err != nil {
		return err
	}
	auction.Status = "Ending"
	return cc.setAuction(ctx, auction)
}

func (cc *SmartContract) FinalizeAuction(
	ctx contractapi.TransactionContextInterface, argjson string,
) error {
	var args FinalizeAuctionArgs
	err := json.Unmarshal([]byte(argjson), &args)
	if err != nil {
		return err
	}

	auction, err := cc.GetAuction(ctx, args.AuctionID)
	if err != nil {
		return err
	}

	addrs, min := cc.ethAddrs()
	if !cc.verifyAuctionResult(addrs, min, args.EthResult) {
		return fmt.Errorf("invalid ethereum result")
	}
	addrs, min = cc.quorumAddrs()
	if !cc.verifyAuctionResult(addrs, min, args.EthResult) {
		return fmt.Errorf("invalid quorum result")
	}

	if args.EthResult.AuctionAddr != auction.EthAddr {
		return fmt.Errorf("invalid ethereum address")
	}
	if args.QuorumResult.AuctionAddr != auction.QuorumAddr {
		return fmt.Errorf("invalid quorum address")
	}

	if args.EthResult.HighestBid >= args.QuorumResult.HighestBid {
		auction.HighestBid = args.EthResult.HighestBid
		auction.HighestBidder = args.EthResult.HighestBidder
		auction.HighestBidPlatform = "ethereum"
	} else {
		auction.HighestBid = args.EthResult.HighestBid
		auction.HighestBidder = args.EthResult.HighestBidder
		auction.HighestBidPlatform = "quorum"
	}

	auction.Status = "Ended"
	err = cc.setAuction(ctx, auction)
	if err != nil {
		return err
	}

	asset, err := cc.GetAsset(ctx, auction.AssetID)
	if err != nil {
		return err
	}

	asset.Owner = auction.HighestBidder
	asset.PendingAuctionID = 0
	err = cc.setAsset(ctx, asset)
	if err != nil {
		return err
	}
	return nil
}

func (cc *SmartContract) ethAddrs() (addrs []string, min int) {
	return []string{
		"",
		"",
		"",
		"",
	}, 2
}

func (cc *SmartContract) quorumAddrs() (addrs []string, min int) {
	return []string{
		"",
		"",
		"",
		"",
	}, 2
}

func (cc *SmartContract) verifyAuctionResult(
	trustedAddrs []string, majority int, result CrossChainAuctionResult,
) bool {
	if len(result.Signatures) < majority {
		return false
	}
	addrs := make([]common.Address, 0, len(trustedAddrs))
	for _, addr := range trustedAddrs {
		addrs = append(addrs, common.HexToAddress(addr))
	}
	return VerifyKnownSignatures(result.Hash(), result.Signatures, addrs)
}

func (cc *SmartContract) GetAsset(
	ctx contractapi.TransactionContextInterface, assetID string,
) (*Asset, error) {
	var asset Asset
	b, err := ctx.GetStub().GetState(cc.makeAssetKey(assetID))
	if err != nil {
		return nil, err
	}
	if b == nil {
		return nil, fmt.Errorf("asset not found")
	}
	err = json.Unmarshal(b, &asset)
	return &asset, err
}

func (cc *SmartContract) GetAuction(
	ctx contractapi.TransactionContextInterface, auctionID int,
) (*Auction, error) {
	b, err := ctx.GetStub().GetState(cc.makeAuctionKey(auctionID))
	if err != nil {
		return nil, err
	}
	if b == nil {
		return nil, fmt.Errorf("auction not found")
	}
	var auction Auction
	err = json.Unmarshal(b, &auction)
	return &auction, err
}

func (cc *SmartContract) GetLastAuctionID(
	ctx contractapi.TransactionContextInterface,
) (int, error) {
	b, err := ctx.GetStub().GetState(KeyLastAuctionID)
	if err != nil {
		return 0, err
	}
	var count int
	json.Unmarshal(b, &count)
	return count, nil
}

func (cc *SmartContract) setAsset(
	ctx contractapi.TransactionContextInterface, asset *Asset,
) error {
	b, _ := json.Marshal(asset)
	err := ctx.GetStub().PutState(cc.makeAssetKey(asset.ID), b)
	if err != nil {
		return fmt.Errorf("set asset error: %v", err)
	}
	return nil
}

func (cc *SmartContract) setAuction(
	ctx contractapi.TransactionContextInterface, auction *Auction,
) error {
	b, _ := json.Marshal(auction)
	err := ctx.GetStub().PutState(cc.makeAuctionKey(auction.ID), b)
	if err != nil {
		return fmt.Errorf("set auction error: %v", err)
	}
	return nil
}

func (cc *SmartContract) setLastAuctionID(
	ctx contractapi.TransactionContextInterface, id int,
) error {
	b, _ := json.Marshal(id)
	return ctx.GetStub().PutState(KeyLastAuctionID, b)
}

func (cc *SmartContract) makeAssetKey(assetID string) string {
	return fmt.Sprintf("%s_%s", KeyAssets, assetID)
}

func (cc *SmartContract) makeAuctionKey(auctionID int) string {
	return fmt.Sprintf("%s_%d", KeyAuctions, auctionID)
}
