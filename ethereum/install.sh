set -x

git clone https://github.com/ethereum/go-ethereum.git
cd go-ethereum
git checkout v1.10.15
go build -o geth ./cmd/geth
sudo mv geth /usr/local/bin

cd ../

git clone https://github.com/Consensys/quorum.git
git checkout v22.1.0
cd quorum
go build -o quorum ./cmd/geth
sudo mv quorum /usr/local/bin

geth version
quorum version