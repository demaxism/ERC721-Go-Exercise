pragma solidity 0.4.24;

import "./zeppelin-solidity/contracts/token/ERC721/ERC721Token.sol";

contract StarNFT is ERC721Token {
    constructor() ERC721Token ("StarNFT", "SNFT") public {}

    uint count;
    
    function mintStar(string memory uri) public returns (uint) {
        uint _id = count;
        super._mint(msg.sender, _id);
        super._setTokenURI(_id, uri);
        count++;
        return _id;
    }
    
    function getOwner(uint id) public view returns (address) {
        return super.ownerOf(id);
    }

    function getURI(uint id) public view returns (string memory) {
        return super.tokenURI(id);
    }
}