pragma solidity ^0.6.6;

contract Reentrancer {
    // Data for fallback function
    address[] public fallbackUnlockAssets;

    address public fallbackGiveTo;
    address public fallbackGiveToken;
    uint public fallbackGiveAmount;

    uint public fallbackState;
    uint public called;

    constructor() public {
        called = 0;
    }

    function pauseLocker(address locker) public {
        bytes memory payload = abi.encodeWithSignature("pause()");
        (bool success, ) = locker.call(payload);
        require(success);
    }

    function setNextLocker(address locker, address nextLocker) public {
        bytes memory payload = abi.encodeWithSignature("setNextLocker(address)", nextLocker);
        (bool success, ) = locker.call(payload);
        require(success);
    }

    function setFallbackUnlockData(address[] memory unlockAssets) public {
        fallbackUnlockAssets = unlockAssets;
    }

    function setFallbackGiveData(address giveTo, address giveToken, uint  giveAmount) public {
        fallbackGiveTo = giveTo;
        fallbackGiveToken = giveToken;
        fallbackGiveAmount = giveAmount;
    }

    function reUnlock(
        uint _fallbackState,
        uint _startCall,
        address locker,
        address[] memory unlockAssets
    ) public {
        fallbackState = _fallbackState;
        called = _startCall;
        bytes memory payload = abi.encodeWithSignature("unlock(address[])", unlockAssets);
        (bool success, bytes memory result) = locker.call(payload);
        require(success, string(result));
    }

    function reGive(
        uint _fallbackState,
        uint _startCall,
        address locker,
        address giveTo,
        address giveToken,
        uint giveAmount
    ) public {
        fallbackState = _fallbackState;
        called = _startCall;
        bytes memory payload = abi.encodeWithSignature("give(address,address,uint256)", giveTo, giveToken, giveAmount);
        (bool success, bytes memory result) = locker.call(payload);
        require(success, string(result));
    }

    fallback() external payable {
        if (fallbackState == 0) {
            if (called > 0) {
                return;
            }

            bytes memory payload = abi.encodeWithSignature("unlock(address[])", fallbackUnlockAssets);
            called = 1;
            (bool success, ) = msg.sender.call(payload);
            require(success);
        } else {
            if (called > 0) {
                return;
            }

            bytes memory payload = abi.encodeWithSignature("give(address,address,uint256)", fallbackGiveTo, fallbackGiveToken, fallbackGiveAmount);
            called = 1;
            (bool success, ) = msg.sender.call(payload);
            require(success);
        }
    }
}
