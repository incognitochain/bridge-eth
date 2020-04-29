pragma solidity ^0.6.6;

import "./trade_utils.sol";
import "./IERC20.sol";

interface WETH {
    function withdraw(uint256 wad) external;
}

contract ZRXTrade is TradeUtils {
    
    // Reciever function which allows transfer eth.
    receive() external payable {}

    address zeroProxy;
    address wETH;

    constructor(
        address _wETH,
        address _zeroProxy,
        address payable _incognitoSmartContract
    ) public {
        zeroProxy = _zeroProxy;
        incognitoSmartContract = _incognitoSmartContract;
        wETH = _wETH;
    }

    function trade(
        IERC20 srcToken,
        uint256 amount,
        IERC20 destToken,
        bytes memory callDataHex,
        address _forwarder
    ) public payable isIncognitoSmartContract returns (address, uint256) {
        // do approve if srcToken is not ETH
        approve(srcToken, zeroProxy, amount);

        // trigger 0x forwarder.
        (bool success, ) = _forwarder.call{value: msg.value}(
            callDataHex
        );
        require(success);

        uint256 sentAmount;
        // if destToken is ETH_CONTRACT_ADDRESS then unwrap WETH back to ETH
        if (destToken == ETH_CONTRACT_ADDRESS) {
            sentAmount = getBalance(ETH_CONTRACT_ADDRESS);
            withdrawWrapETH(sentAmount);
            transfer(ETH_CONTRACT_ADDRESS, sentAmount);
        } else {
            sentAmount = getBalance(destToken);
            transfer(destToken, sentAmount);
        }
        return (address(destToken), sentAmount);
    }

    function getBalance(IERC20 token) internal view returns (uint256) {
        if (token == ETH_CONTRACT_ADDRESS) {
            return balanceOf(IERC20(wETH));
        }
        return balanceOf(token);
    }

    function withdrawWrapETH(uint256 amount) public {
        WETH(wETH).withdraw(amount);
    }
}
