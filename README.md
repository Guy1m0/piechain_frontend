# PIEChain

We develop a general cross-chain communication framework that uses the Kafka protocol for secure interaction.

![alt text](design.png)

## Entities
* The Kafka network provider, who maintains the Kafka network.
* Developers, who develops cross-chain services (CC-SVCs) using an event-
driven approach.
* End users, who deploy smart contracts and use the CC-SVCs for cross- chain operations.

## Setup

### Kafka
Start Kafka with `docker-compose`.
```bash
cd kafka
docker-compose up -d

# to stop after experiments
docker-compose down
```

### Fabric
Install Hyperledger Fabric.
```bash
cd fabric
./install.sh
```

### Ethereum
Install go-ethereum and quorum.
```bash
cd ethereum
sudo ./install.sh
```

## Usage

Register event handlers for different services.
```go
ccsvc, err = cclib.NewEventService(
	[]string{"zoolkeeper_nodes"},
	"service_name",
)
check(err)
// register event handlers
ccsvc.Register("event_1", handleEvent1) 
ccsvc.Register("event_2", handleEvent2)
err = ccsvc.Start() // start service
check(err)
```

Publish events with corresponding payloads
```go
b, _ := json.Marshal(Event1Body{
	Hash: hash,
	// ...
})
ccsvc.Publish("event_1", b)
```

## Example Applications

### Flash Loan 
[Demo](https://drive.google.com/file/d/1dBHgUrdAmx1COR_3nYZLI5uPjbjk2Sks/view?usp=sharing)

You can run the cross-chain flash loan demo by the following steps.


1. Run Fabric.
```bash
cd fabric-samples/test-network
./network.sh up createChannel

# to stop fabric after experiment
./network.sh down
```

2. Run ethereum.
```bash
cd ethereum/poa
./remove.sh
./init.sh
./start.sh

# to stop ethereum after experiment
./stop.sh
```
NOTE: Check `log.txt` file to ensure that ethereum is running.

3. Deploy contracts `fabric_erc20` and `lender` on fabric.
```bash
cd fabric-samples/test-network
./network.sh deployCC -ccn token1 -ccp ../../contracts/fabric_erc20/chaincode/ -ccl go
./network.sh deployCC -ccn lender -ccp ../../contracts/fabric_lender/chaincode/ -ccl go
```

4. Run `cclender` and `ccarbit` in two different terminals.
They are the cross-chain exchange services responsible for relaying messages between two blockchains.
```bash
cd examples/flashloan/cclender
go build .
./cclender

cd examples/flashloan/ccarbit
go build .
./ccarbit
```

5. Setup `tokens` and `amm` exchanges using `cli`
```bash
cd examples/flashloan/cli
go build .
# Deploy tokens and amm contracts and allocate initial coins.
./cli -c setup
```

6. Setup flashloan contract `arbitrage` on `ethereum`.
```bash
cd examples/flashloan/arbitrageur
go build .
# Deploy arbitrage contract on ethereum.
# Setup both lender and arbitrage contract and generate a signature.
./arbitrageur -c setup

```
You can check the results in `flash_loan.json` and `commit_vote.json` in the `flashloan` folder.

7. Register the new flashloan to the cc services.
```bash
# from arbitrageur directory
./arbitrageur -c register
```

8. Sign Commit Vote by `lender`.
```bash
cd examples/flashloan/lender
go build .
./lender -c sign
```

9. Initialize flash-loan with commit vote by `lender`.
```bash
# from lender directory
./lender -c initialize
```
After this step, cross-chain services will detect lender initialization on `fabric` and transfer loan to `arbitrage` contract on `ethereum`.

10. Execute `arbitrage` by `arbitrageur`.
```bash
./arbitrageur -c execute
```
If `execute` is successful, the `arbitrage` contract will transfer `loan + intrest` to the cross-chain exchange and the cross-chain services will transfer `loan + intrest` to the lender.
if `execute` is failed, the `lender` will still get refund of initial `loan` on fabric.

You can check the token balances on both blockchains using this command.
```bash
# from cli directory
./cli -c display
```

### Auction with three blockchains
[Demo](https://drive.google.com/file/d/16f1X0UOpoHOm3NqUJa_K66PA0JtSMzYq/view?usp=sharing)

You can run the auction example by the following steps.
1. Run Fabric.
```bash
cd fabric-samples/test-network
./network.sh up createChannel
```

2. Run Ethereum POA.
```bash
cd ethereum/poa
./remove.sh
./init.sh
./start.sh
```

3. Run Quorum Raft.
```bash
cd ethereum/raft
./remove.sh
./init.sh
./start.sh
```

4. Deploy `fabric_asset` on `fabric`.
```bash
./network.sh deployCC -ccn asset -ccp ../../contracts/fabric_asset/chaincode -ccl go
```

3. Run `relayer` crosschain service.
```bash
cd examples/auction/relayer
go build .
./relayer

4. Run `signer` crosschain services.
# from examples/auction/signer run signers with different keys for multiple instances (at least 2 for each auction blockchain)

cd examples/auction/signer
go build .

# on different terminals
./signer -p ethereum -eth localhost:8545 -key ../../keys/key1 -id 1
./signer -p ethereum -eth localhost:8545 -key ../../keys/key2 -id 2
./signer -p quorum -eth localhost:8546 -key ../../keys/key1 -id 1
./signer -p quorum -eth localhost:8546 -key ../../keys/key2 -id 2
```

3. Run the `scenario` script.
```bash
cd examples/auction/scenario
go run .
```

The scenario script will do the following steps.
1. Add a new asset on the fabric `asset` contract.
3. Deploy auction contracts on `ethereum` and `quorum`.
2. Create a new auction for this asset on on fabric.
3. Bid correspondingly on both `ethereum` and `quorum`.
4. End auction on `fabric` and print out the winner info and final asset owner on `fabric`.

### Auction with Ethereum POW
[Demo](https://drive.google.com/file/d/16f1X0UOpoHOm3NqUJa_K66PA0JtSMzYq/view?usp=sharing)

You can run the auction example by the following steps.
1. Run Fabric.
```bash
cd fabric-samples/test-network
./network.sh up createChannel
```

2. Run Ethereum POW.
```bash
cd ethereum/pow
./remove.sh
./init.sh
./start.sh
```

3. Deploy `fabric_asset_pow` on `fabric`.
```bash
./network.sh deployCC -ccn asset -ccp ../../contracts/fabric_asset_pow/chaincode -ccl go
```

4. Run `relayer` crosschain service.
```bash
cd examples/auction_pow/relayer
go build .
./relayer

5. Run `end_listener` crosschain services.
cd examples/auction_pow/end_listener
go build .
./end_listener
```

6. Run the `scenario` script.
```bash
cd examples/auction_pow/scenario
go run .
```

The scenario script will do the following steps.
1. Add a new asset on the fabric `asset` contract.
3. Deploy auction contract on `ethereum`.
2. Create a new auction for this asset on on fabric.
3. Bid on both `ethereum`.
4. End auction on `ethereum`.
5. `end_listener` will listen to ethereum auction ending and will publish `on_end_auction` event with POW headers and merkle proof for auction wiiner.
6. On receiving `on_end_auction` event, `relayer` will end auction on fabric.

