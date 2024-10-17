// SPDX-License-Identifier: MIT
pragma solidity >=0.8.1 <0.9.0;

import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

/*
UUPS proxies are implemented using an {ERC1967Proxy}. Note that this proxy is not by itself upgradeable. It is the role of the implementation to include, alongside the contract’s logic, all the code necessary to update the implementation’s address that is stored at a specific slot in the proxy’s storage space. This is where the {UUPSUpgradeable} contract comes in. Inheriting from it (and overriding the {xref-UUPSUpgradeable-_authorizeUpgrade-address-}[_authorizeUpgrade] function with the relevant access control mechanism) will turn your contract into a UUPS compliant implementation.
*/

/*
Requirements:
     *
     * - If `data` is empty, `msg.value` must be zero.
*/

/*
UUPS와 Beacon은 같이 사용하지 않음. 서로의 목적이 다르기 때문.
UUPS는 implement에서 업그레이드 로직을 관리하기 위함. Beacon은 한번의 작업으로 모든 프록시들의 Implemet 주소를 바꾸기 위함
*/

contract MyProxy is ERC1967Proxy {
    
    constructor(address implementation, bytes memory data) ERC1967Proxy(implementation, data) {}

    function getImplementation() external view returns (address) {
        return _implementation();
    }

    receive() external payable {
        // Do Nothing
    }
}