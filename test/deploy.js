const StarNFT = artifacts.require('./StarNFT');
const fs = require('fs');

contract("StarNFT", function(accounts) {
    
    it ("should deploy a new token:", async () => {
        let token = await StarNFT.new();
        let contractAddr = token.address;
        console.log("Token deployed with contract address:" + contractAddr);
        fs.writeFileSync("contractAddr.txt",contractAddr);
    });
});