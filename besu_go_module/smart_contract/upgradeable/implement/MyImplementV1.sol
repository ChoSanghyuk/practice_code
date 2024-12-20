// SPDX-License-Identifier: MIT
pragma solidity >=0.8.1 <0.9.0;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {UUPSUpgradeable} from  "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

/*
besu에서 UUPS 업그레이더블 패턴 배포 시 항상 실패하는 문제 발생.
=> Remix에서 배포 시도 시, solidity version 0.8.27 + cancun fork 적용 Evm에서 배포 가능
   현재 Besu(24.9.1)는 cancunBlock으로 genesis 설정 시, 적용 X. 추후 시도
*/
contract MyImplementV1 is Initializable, UUPSUpgradeable, OwnableUpgradeable {

    uint256 internal _value;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize() public initializer {
        __Ownable_init(msg.sender);
        __UUPSUpgradeable_init();
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyOwner{
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