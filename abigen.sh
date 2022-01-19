#!/bin/sh


echo Compiling solidity and generating golang...

set -x

cd contracts/eth_arbitrage
abigen --sol arbitrage.sol --pkg eth_arbitrage --type Arbitrage --out arbitrage_gen.go
cd ../..

cd contracts/eth_auction
abigen --sol auction.sol --pkg eth_auction --type Auction --out auction_gen.go
cd ../..

cd contracts/eth_lender
abigen --sol Lender.sol  --pkg eth_lender --type Lender --out lender_gen.go
cd ../..
