// SPDX-License-Identifier: MIT
pragma solidity >=0.8.1 <0.9.0;

import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {UUPSUpgradeable} from  "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

contract MyImplementV1 is Initializable, AccessControlUpgradeable, UUPSUpgradeable {

    bytes32 public constant CUSTOM_OWNER_ROLE = keccak256("CUSTOM_OWNER_ROLE"); // OwnableUpgradealbe.sol 사용도 가능
    uint256 internal _value;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize() public initializer {
        __AccessControl_init();

        _grantRole(CUSTOM_OWNER_ROLE, msg.sender);
        __UUPSUpgradeable_init();
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyRole(CUSTOM_OWNER_ROLE){
        // Only the owner can authorize an upgrade
    }

    function getValue() external view returns(uint256){
        return _value;
    }

    function setValue(uint256 newValue) external {
        _value = newValue;
    }

    function getVersion() external virtual pure returns(uint16) {
        return 1;
    }
}