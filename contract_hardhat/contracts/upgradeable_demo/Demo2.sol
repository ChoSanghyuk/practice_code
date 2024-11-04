// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import {Demo} from "./Demo.sol";

// A contrived example of a contract that can be upgraded
contract DemoV2 is Demo {
  string public name;

  function version() public override pure returns (string memory) {
    return "2.0.0";
  }

  function setName(string memory _name) public {
    name = _name;
  }
}