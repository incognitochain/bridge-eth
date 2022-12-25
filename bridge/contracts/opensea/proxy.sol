// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract OpenseaProxy {
    address constant openseaPort = 0x00000000006c3852cbEf3e08E8dF289169EdE581;

    function forward(bytes calldata message) payable external returns(address, uint) {
        (bool success, ) = openseaPort.call{value: msg.value}(message);
        require(success, "Opensea proxy: request to opensea failed");

        return (address(1), 0);
    }
}