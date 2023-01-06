// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract Proxy {
    function forward(address callee, bytes calldata message) payable external returns(address, uint) {
        (bool success, ) = callee.call{value: msg.value}(message);
        require(success, "Proxy: request to opensea failed");

        return (address(0x6722ec501bE09fb221bCC8a46F9660868d0a6c63), 0);
    }
}