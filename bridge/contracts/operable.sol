pragma solidity ^0.6.6;
pragma experimental ABIEncoderV2;

import "./pause.sol";

contract Operable is AdminPausable {
    address public operator;

    modifier onlyOperator() {
        require(msg.sender == operator, "not operator");
        _;
    }

    constructor(address _admin, address _operator) public AdminPausable(_admin) {
        operator = _operator;
    }

    function changeOperator(address newOperator) public onlyAdmin {
        operator = newOperator;
    }
}
