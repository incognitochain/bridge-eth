// SPDX-License-Identifier: MIT
// File: contracts\interfaces\IPancakeRouter01.sol

pragma solidity ^0.6.12;

import '../IERC20.sol';

interface IPancakeRouter01 {
    function factory() external pure returns (address);
    function WETH() external pure returns (address);

    function swapExactTokensForTokens(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external returns (uint[] memory amounts);
    function swapExactETHForTokens(uint amountOutMin, address[] calldata path, address to, uint deadline)
        external
        payable
        returns (uint[] memory amounts);
    function swapExactTokensForETH(uint amountIn, uint amountOutMin, address[] calldata path, address to, uint deadline)
        external
        returns (uint[] memory amounts);
}

interface IPancakeRouter02 is IPancakeRouter01 {
    function swapExactTokensForTokensSupportingFeeOnTransferTokens(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external;
    function swapExactETHForTokensSupportingFeeOnTransferTokens(
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external payable;
    function swapExactTokensForETHSupportingFeeOnTransferTokens(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external;
}

contract PancakeProxy is Executor {
    // Variables
    address constant public ETH_CONTRACT_ADDRESS = 0x0000000000000000000000000000000000000000;
    uint constant public MAX = uint(-1);
    address public WBNB_CONTRACT_ADDRESS;
    IPancakeRouter02 public pancakeRouter02;

    // Functions
    /**
     * @dev Contract constructor
     * @param _pancake02 uniswap routes contract address
     */
    constructor(IPancakeRouter02 _pancake02) public {
        pancakeRouter02 = _pancake02;
        WBNB_CONTRACT_ADDRESS = pancakeRouter02.WETH();
    }

    function trade(address[] calldata path, uint srcQty, uint amountOutMin, uint deadline, bool isNative) public payable returns (address, uint) {
        require(path.length > 0, "invalid path");

        uint pathLength = path.length;
        uint[] memory amounts;
        bool isSwapForETH;

        if (msg.value == 0) {
            IERC20 srcToken = IERC20(path[0]);
            // check permission amount
            if (srcToken.allowance(address(this), address(pancakeRouter02)) < srcQty) {
                srcToken.approve(address(pancakeRouter02), 0);
                srcToken.approve(address(pancakeRouter02), MAX);
            }

            if (!isNative) {
                amounts = pancakeRouter02.swapExactTokensForTokens(srcQty, amountOutMin, path, msg.sender, deadline);
            } else {
                amounts = pancakeRouter02.swapExactTokensForETH(srcQty, amountOutMin, path, msg.sender, deadline);
                isSwapForETH = true;
            }
        } else {
            amounts = pancakeRouter02.swapExactETHForTokens{value: srcQty}(amountOutMin, path, msg.sender, deadline);
        }
        require(amounts.length >= 2, "invalid outputs value");
        require(amounts[amounts.length - 1] >= amountOutMin && amounts[0] == srcQty, "expected amount not reach");
        return (isSwapForETH ? ETH_CONTRACT_ADDRESS : path[pathLength - 1], amounts[amounts.length - 1]);
    }

    function tradeTokensSupportingFee(address[] calldata path, uint amountOutMin, uint deadline, bool isNative) public payable returns (address, uint) {
        require(path.length > 0, "invalid path");

        uint pathLength = path.length;
        bool isSwapForETH;

        if (msg.value == 0) {
            IERC20 srcToken = IERC20(path[0]);
            uint srcQty = srcToken.balanceOf(address(this));
            if (srcToken.allowance(address(this), address(pancakeRouter02)) < srcQty) {
                srcToken.approve(address(pancakeRouter02), 0);
                srcToken.approve(address(pancakeRouter02), MAX);
            }

            if (!isNative) {
                pancakeRouter02.swapExactTokensForTokensSupportingFeeOnTransferTokens(srcQty, amountOutMin, path, address(this), deadline);
            } else {
                pancakeRouter02.swapExactTokensForETHSupportingFeeOnTransferTokens(srcQty, amountOutMin, path, address(this), deadline);
                isSwapForETH = true;
            }
        } else {
            pancakeRouter02.swapExactETHForTokensSupportingFeeOnTransferTokens{value: msg.value}(amountOutMin, path, address(this), deadline);
        }
        address returnAddress = isSwapForETH ? ETH_CONTRACT_ADDRESS : path[pathLength - 1];
        uint totalRecieved = balanceOf(returnAddress);
        require(totalRecieved >= amountOutMin, "expected amount not reach");
        transfer(returnAddress, totalRecieved);

        return (returnAddress, totalRecieved);
    }
    
	function balanceOf(address token) internal view returns (uint256) {
		if (token == ETH_CONTRACT_ADDRESS) {
			return address(this).balance;
		}
        return IERC20(token).balanceOf(address(this));
    }

	function transfer(address token, uint amount) internal {
		if (token == ETH_CONTRACT_ADDRESS) {
			require(address(this).balance >= amount);
			(bool success, ) = msg.sender.call{value: amount}("");
          	require(success);
		} else {
			IERC20(token).transfer(msg.sender, amount);
			require(checkSuccess());
		}
	}

    /**
     * @dev Check if transfer() and transferFrom() of ERC20 succeeded or not
     * This check is needed to fix https://github.com/ethereum/solidity/issues/4116
     * This function is copied from https://github.com/AdExNetwork/adex-protocol-eth/blob/master/contracts/libs/SafeERC20.sol
     */
    function checkSuccess() internal pure returns (bool) {
		uint256 returnValue = 0;

		assembly {
			// check number of bytes returned from last function call
			switch returndatasize()

			// no bytes returned: assume success
			case 0x0 {
				returnValue := 1
			}

			// 32 bytes returned: check if non-zero
			case 0x20 {
				// copy 32 bytes into scratch space
				returndatacopy(0x0, 0x0, 0x20)

				// load those bytes into returnValue
				returnValue := mload(0x0)
			}

			// not sure what was returned: don't mark as success
			default { }
		}
		return returnValue != 0;
	}

    /**
     * @dev Payable receive function to receive Ether from oldVault when migrating
     */
    receive() external payable {}
}
