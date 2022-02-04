#!/bin/sh

set -x

nohup geth \
--datadir data --networkid 15 --nodiscover \
--unlock 0x2b50a4fe2a6300116edbcf2632d079a12abf5f2d --password password.txt \
--http --http.api eth,web3,net --http.addr "localhost" --http.corsdomain "*" \
--allow-insecure-unlock --gcmode archive \
--rpc.txfeecap 0 --mine >& log.txt 2>&1 &
