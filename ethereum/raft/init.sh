#!/bin/sh

set -x

quorum --datadir data init genesis.json
cp data/nodekey data/geth