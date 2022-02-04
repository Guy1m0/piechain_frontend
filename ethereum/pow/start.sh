#!/bin/sh

set -x

nohup geth \
--datadir data --networkid 15 --nodiscover \
--http --http.api eth,web3,net --http.addr "localhost" --http.corsdomain "*" \
--rpc.txfeecap 0 --mine --miner.threads 1 >& log.txt 2>&1 &
