// SPDX-License-Identifier: MIT
pragma solidity >=0.8.1 <0.9.0;

import {AccessControlUpgradeable} "@openzeppelin/contracts-upgradealbe/access/AccessControlUpgradeable.sol";
import {Initializable} "@openzeppelin/contracts-upgradealbe/proxy/utils/Initializable.sol";
import {UUPSUpgradeable} "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol"

contract MyImplementV1 is Initializable, AccessControlUpgradeable, UUPSUpgradeable {

    bytes32 public constant CUSTOM_ROLE = keccak256("CUSTOM_ROLE");

    constructor() {
        _disableInitializers();
    }

    function initialize() public initializer {
        _AccessControl_init();

        _grantRole(CUSTOM_ROLE, msg.sender);
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }


}