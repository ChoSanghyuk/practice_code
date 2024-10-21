import {
    time,
    loadFixture,
  } from "@nomicfoundation/hardhat-toolbox/network-helpers";
  import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
  import { expect } from "chai";
  import hre from "hardhat";
  const { ethers, upgrades } = require("hardhat");

  
  describe("ImplementV1", function () {

    it("Deploy", async function () {
        const MyUpgradeableContract = await ethers.getContractFactory("MyImplementV1");
        const proxy = await upgrades.deployProxy(MyUpgradeableContract);
        await proxy.deployed();
        
        console.log("Proxy deployed to:", proxy.address);
    });
  })