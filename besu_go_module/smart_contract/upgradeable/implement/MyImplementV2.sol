// SPDX-License-Identifier: MIT
pragma solidity >=0.8.1 <0.9.0;

import {MyImplementV1} from "./MyImplementV1.sol";


contract MyImplementV2 is MyImplementV1{

    constructor() {
        _disableInitializers();
    }

    function getVersion() external override pure returns(uint16) {
        return 2;
    }
}