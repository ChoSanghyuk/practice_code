// SPDX-License-Identifier: MIT
pragma solidity >=0.8.1 <0.9.0;

import "@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol";

contract MyBeaconProxy is BeaconProxy {

    constructor(address newBeacon) BeaconProxy(newBeacon, bytes("")) {
        _changeAdmin(msg.sender)
    }

    modifier ifAdmin() {
        if (msg.sender == _getAdmin()) {
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

    function getBeacon() external view returns (address) {
        return _getBeacon();
    }

    function upgradeToBeacon(address newBeacon) external ifAdmin {
        _setBeacon(newBeacon, bytes(""));
    }

    function getImplementation() external view returns (address) {
        return _implementation();
    }
}