
## Prepare ERC721 contract and mint some tokens

Create a `.secret` file in this directory, with a twelve word phrase mnemonic of your Wallet, as deploying contracts comsumes gas. This will be used by truffle-config.js

### Deploy the contract in test network:
`truffle test test/deploy.js --network ropsten`

Contract StarNFT will be deployed in ropsten network.
It will generate a `contractAddr.txt` in this directory, with the contract address in it, which will be used by test/mint.js

### Mint a token:
`truffle test test/mint.js --network ropsten`

TokenId is an incremential integer start from 0.

## Launch the service

Start http server:
`go run main.go`

In browser, open `http://localhost:8080/<CONTRACT_ADDRESS>/<TOKENID> ` 

example: http://localhost:8080/0x77b9F65b1d9805a5DaB29dF9e66F6D5441fC003e/2

This will return the owner of the given token ID and the contract address.
