// SPDX-License-Identifier: MIT
pragma solidity >=0.8.1 <0.9.0;

import "@openzeppelin/contracts-upgradealbe/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradealbe/proxy/utils/Initializable.sol";

contract MyImplementV1 is Initializable, AccessControlUpgradeable {

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