// SPDX-License-Identifier: MIT
pragma solidity >=0.8.1 <0.9.0;

import "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/utils/Address.sol";

contract MyBeacon is IBeacon, AccessControl {

    bytes32 public constant IMPLEMENTATION_ROLE = keccak256("IMPLEMENTATION_ROLE")

    address private _implementation;

    event Upgraded(address indexed implementation);

    constructor(address newImplementation) {
        _setImplementation(newImplementation);
        _grantRole(DEFAULT_ADMIN_ROLE, _msgSender());
        _grantRole(IMPLEMENTATION_ROLE, _msgSender());
    }

    function implementation() public view virtual override returns (address) {
        return _implementation;
    }

    function upgradeTo(address newImplementation) public virtual onlyRole(IMPLEMENTATION_ROLE) {
        _setImplementation(newImplementation);
        emit Upgraded(newImplementation);
    }

    function _setImplementation(address newImplementation) private {
        require(Address.isContract(newImplementation), "Implementation is not a contract")
        _implementation = newImplementation;
    }


}