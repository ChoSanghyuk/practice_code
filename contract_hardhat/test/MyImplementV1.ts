import { expect } from "chai";
import { ignition, ethers } from "hardhat";

import MyProxyV1Module from "../ignition/modules/MyImplementModule";


describe("Demo Proxy", function () {
  describe("Proxy interaction", async function () {
    it("Should be interactable via proxy", async function () {
      const [, otherAccount] = await ethers.getSigners();

      const { impl } = await ignition.deploy(MyProxyV1Module);

      expect(await impl.connect(otherAccount).getVersion()).to.equal(1); // 아니 빨간줄이 있어도 실행이 되고, 성공을 하네.
    });
  });

});


/*
import { expect } from "chai";
import { ethers } from "ethers"; // Use this for utility functions like ethers.utils
import { ethers as hardhatEthers, upgrades } from "hardhat"; // For Hardhat-specific ethers and upgrades

import { MyImplementV1 } from "../typechain-types"; // adjust if your typechain directory is different


describe("MyImplementV1 Contract", function () {
  let myContract: MyImplementV1;
  let owner: any;
  let otherAccount: any;
  const CUSTOM_OWNER_ROLE = ethers.keccak256(
    ethers.toUtf8Bytes("CUSTOM_OWNER_ROLE")
  );

  beforeEach(async function () {
    [owner, otherAccount] = await hardhatEthers.getSigners();

    // Deploying the contract using OpenZeppelin's upgradeable factory
    const MyImplementV1 = await hardhatEthers.getContractFactory("MyImplementV1");

        // myContract = (await upgrades.deployProxy(MyImplementV1, {
        //   initializer: "initialize",
        // })) as MyImplementV1;
        try {
    myContract = (await upgrades.deployProxy(MyImplementV1, {
        initializer: "initialize",
    })) as unknown as MyImplementV1; // Double-cast here
    }catch (error) {
        console.error("Error in beforeEach:", error);
    }

  });

  it("should initialize with correct role for owner", async function () {
    const hasRole = await myContract.hasRole(CUSTOM_OWNER_ROLE, owner.address);
    expect(hasRole).to.be.true;
  });

  it("should allow the owner to set a value", async function () {
    await myContract.setValue(42);
    const value = await myContract.getValue();
    expect(value).to.equal(42);
  });

  it("should return version 1", async function () {
    const version = await myContract.getVersion();
    expect(version).to.equal(1);
  });

  
  // 아직 여기는 해결되지 않았음
  it("should only allow owner to authorize an upgrade", async function () {

    // Check that only an account with the CUSTOM_OWNER_ROLE can upgrade
    const NewImplementation = await hardhatEthers.getContractFactory("MyImplementV1");
    await expect(
      upgrades.upgradeProxy(await myContract.getAddress(), NewImplementation, {
        // from: otherAccount.address,
      })
    ).to.be.revertedWith("AccessControl: account");

    // Upgrade should succeed when called by the owner
    await expect(upgrades.upgradeProxy(await myContract.getAddress(), NewImplementation))
      .to.emit(myContract, "Upgraded");
  });
  

});
*/
