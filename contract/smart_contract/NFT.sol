pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract NonFungibleToken is ERC721 {
    constructor(string memory name_, string memory symbol_) ERC721(name_, symbol_){
        
    }
}