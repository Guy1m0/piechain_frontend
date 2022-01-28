package asset_pow

import (
	"bytes"
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
	ctx contractapi.TransactionContextInterface, argsJson string,
) error {
	var args StartAuctionArgs
	err := json.Unmarshal([]byte(argsJson), &args)
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
		ID:            lastID + 1,
		AssetID:       args.AssetID,
		AuctionAddr:   args.AssetID,
		BaseEthHeader: args.EthHeader,
		Status:        "Started",
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
	ctx contractapi.TransactionContextInterface, argsJson string,
) error {

	var args EndAuctionArgs
	err := json.Unmarshal([]byte(argsJson), &args)
	if err != nil {
		return err
	}

	auction, err := cc.getAuctionInternal(ctx, args.AuctionID)
	if err != nil {
		return err
	}

	if err := VerifyEthHeaders(args.EthHeaders, auction.BaseEthHeader); err != nil {
		return fmt.Errorf("invalid eth headers %w", err)
	}

	if !bytes.Equal(args.ProvableResult.Address.Bytes(), common.FromHex(auction.AuctionAddr)) {
		return fmt.Errorf("unknown auction address")
	}

	if args.ProvableResult.StorageProof[0].Key != "0x1" {
		return fmt.Errorf("invalid proof key")
	}

	if args.ProvableResult.StorageProof[1].Key != "0x2" {
		return fmt.Errorf("invalid proof key")
	}

	lastHeader := args.EthHeaders[len(args.EthHeaders)-1]
	ok, err := VerifyProof(lastHeader.Root, args.ProvableResult)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("verify merkle proof failed")
	}

	// TODO: set auction winer
	auction.ProvableResult = args.ProvableResult
	auction.EthHeaders = args.EthHeaders
	auction.HighestBidder = common.Bytes2Hex(args.ProvableResult.StorageProof[0].Value.Bytes())
	auction.HighestBid = int(args.ProvableResult.StorageProof[1].Value.Int64())
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
	return cc.setAsset(ctx, asset)
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

func (cc *SmartContract) getAuctionInternal(
	ctx contractapi.TransactionContextInterface, auctionID int,
) (*Auction, error) {
	b, err := cc.GetAuction(ctx, auctionID)
	if err != nil {
		return nil, err
	}
	var auction Auction
	err = json.Unmarshal([]byte(b), &auction)
	return &auction, err

}

func (cc *SmartContract) GetAuction(
	ctx contractapi.TransactionContextInterface, auctionID int,
) (string, error) {
	b, err := ctx.GetStub().GetState(cc.makeAuctionKey(auctionID))
	if err != nil {
		return "", err
	}
	if b == nil {
		return "", fmt.Errorf("auction not found")
	}
	return string(b), nil
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
