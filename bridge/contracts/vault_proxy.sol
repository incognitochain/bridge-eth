pragma solidity 0.6.6;

import "./pause.sol";

/**
 * @dev Interface of the contract capable of checking if an instruction is
 * confirmed over at Incognito Chain
 */
interface Incognito {
    function instructionApproved(
        bool,
        bytes32,
        uint,
        bytes32[] calldata,
        bool[] calldata,
        bytes32,
        bytes32,
        uint[] calldata,
        uint8[] calldata,
        bytes32[] calldata,
        bytes32[] calldata
    ) external view returns (bool);
}

/**
 * @dev Interface of the previous Vault contract to query burn proof status
 */
interface Withdrawable {
    function isWithdrawed(bytes32)  external view returns (bool);
    function isSigDataUsed(bytes32)  external view returns (bool);
    function getDepositedBalance(address, address)  external view returns (uint);
    function updateAssets(address[] calldata, uint[] calldata) external returns (bool);
    function paused() external view returns (bool);
}

contract Delegator is AdminPausable {

    address constant public ETH_TOKEN = 0x0000000000000000000000000000000000000000;
    address public delegator;
    Withdrawable public preVault;
    Incognito public incognito;
    bool notEntered = true;

    /**
     * @dev Prevents a contract from calling itself, directly or indirectly.
     * Calling a `nonReentrant` function from another `nonReentrant`
     * function is not supported. It is possible to prevent this from happening
     * by making the `nonReentrant` function external, and make it call a
     * `private` function that does the actual work.
     */
    modifier nonReentrant() {
        // On the first call to nonReentrant, notEntered will be true
        require(notEntered, "can not reentrant");

        // Any calls to nonReentrant after this point will fail
        notEntered = false;
        _;
    }

    constructor(address _admin, address _delegator, Incognito _incognito, Withdrawable _prevault) public {
        require(_admin != address(0) && _delegator != address(0) && address(_incognito) != address(0));
        admin = _admin;
        expire = block.timestamp + 365 * 3 days;
        delegator = _delegator;
        incognito = _incognito;
        preVault = _prevault;
    }

    /**
     * @dev Payable receive function to receive Ether from oldVault when migrating
     */
    receive() external payable {}

    fallback() nonReentrant external payable {
        (bool success,) = delegator.delegatecall(msg.data);
        require(success, "delegate call reverted");

        // By storing the original value once again, a refund is triggered (see
        // https://eips.ethereum.org/EIPS/eip-2200)
        notEntered = true;
        assembly {
            // If the call succeeded return the return data from the delegate call
            let _returndataMemoryOffset := mload(0x40)
            let _returndatasize := returndatasize()
            // Update the pointer at 0x40 again to point at new free memory location so any theoretical allocation doesn't stomp our memory in this call
            mstore(0x40, add(_returndataMemoryOffset, _returndatasize))
            returndatacopy(_returndataMemoryOffset, 0x0, _returndatasize)
            return(_returndataMemoryOffset, returndatasize())
        }
    }
}