// hardhat ts test code for ERC721.sol
// Run test with `npx hardhat test`


import { ethers } from "hardhat";
import { expect } from "chai";
import { BigNumber, Contract } from "ethers";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";

describe("NFT", () => {
  let contract: Contract;
  let owner: SignerWithAddress;
  let addr1: SignerWithAddress;
  let addr2: SignerWithAddress;
  let addrs: SignerWithAddress[];

  beforeEach(async () => {
    const NFT = await ethers.getContractFactory("NFT");
    contract = await NFT.deploy();
    [owner, addr1, addr2, ...addrs] = await ethers.getSigners();
  });

  describe("Deployment", () => {
    it("Should set the right owner", async () => {
      expect(await contract.owner()).to.equal(owner.address);
    });
  });

  describe("Mint", () => {
    it("Should mint a token", async () => {
      await contract.mint(owner.address, 0);
      expect(await contract.ownerOf(0)).to.equal(owner.address);
    });
  });

  describe("Transfer", () => {
    it("Should transfer a token", async () => {
      await contract.mint(owner.address, 0);
      await contract.transferFrom(owner.address, addr1.address, 0);
      expect(await contract.ownerOf(0)).to.equal(addr1.address);
    });
  });
});



