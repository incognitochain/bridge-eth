// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract Proxy {
    function forward(address callee, bytes calldata message) payable external returns(address, uint) {
        (bool success, ) = callee.call{value: msg.value}(message);
        require(success, "Proxy: request to opensea failed");

        return (address(0x4cB607c24Ac252A0cE4b2e987eC4413dA0F1e3Ae), 0);
    }
}