const StarNFT = artifacts.require('./StarNFT');
const fs = require('fs');
let _inst;
let contractAddr = "0x77b9F65b1d9805a5DaB29dF9e66F6D5441fC003e";

contract("StarNFT", function(accounts) {

    let tokenId;
    try {
        contractAddr = fs.readFileSync("contractAddr.txt").toString().trim();
    }
    catch (err) { console.log("Read contract address failed:" + err) }
    
    const tokenUri = "I created an awesome token!";

    it ("should mint a new non-fungible token", async () => {
        console.log("Contract address:" + contractAddr);
        let token = await StarNFT.at(contractAddr);
        tokenId = await token.mintStar(tokenUri);
        console.log("New NFT minted, id:" + JSON.stringify(tokenId));
    });

});