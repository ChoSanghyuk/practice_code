// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.1 <0.9.0;

contract ClientReceipt {
    event Deposit(
        address indexed from,
        bytes32 indexed id,
        uint value
    );

    function deposit(bytes32 id) public payable {
        emit Deposit(msg.sender, id, msg.value);
    }
}