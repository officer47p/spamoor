// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.19;

contract LoadTest {
    event DataReceived(bytes data);

    function processCallData(bytes memory data) public {
        emit DataReceived(data);
    }
}
