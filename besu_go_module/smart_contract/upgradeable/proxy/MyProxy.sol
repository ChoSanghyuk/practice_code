// SPDX-License-Identifier: MIT
pragma solidity >=0.8.1 <0.9.0;

import "@openzeppelin/contracts/proxy/Proxy.sol";
import "@openzeppelin/contracts/proxy/ERC1967Upgrade.sol";

contract MyProxy is Proxy {
    
    constructor(address newImplementation) {
        _upgradeTo(newImplementation);
        _changeAdmin(msg.sender);
    }

    modifier ifAdmin() {
        if (msg.sender == _getAdmin()){
            _;
        } else {
            _fallback();
        }
    }

    function getAdmin() public view returns (address) {
        return _getAdmin();
    }

    function changeAdmin(address newAdmin) external ifAdmin {
        _changeAdmin(newAdmin)
    }

    function upgradeTo(address newImplementation) external ifAdmin {
        _upgradeTo(newImplementation);
    }

    function getImplementation() external view returns (address) {
        return _implementation();
    }

    function _implementation() internal view virtual override returns (address) {
        return _getImplementation();
    }
}