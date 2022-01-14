// SPDX-License-Identifier: GPL-2.0-or-later
pragma solidity =0.7.6;
pragma abicoder v2;
import "../IERC20.sol";
import "@uniswap/v3-periphery/contracts/libraries/Path.sol";
import "@uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol";

interface ISwapRouter2 is ISwapRouter {
	function WETH9() external returns(address);
}

interface Wmatic is IERC20 {
	function withdraw(uint256 amount) external;
}

contract UniswapProxy {
	using Path for bytes;
    // Variables
    address constant public ETH_CONTRACT_ADDRESS = 0x0000000000000000000000000000000000000000;
    uint constant public MAX = uint(-1);
    ISwapRouter2 public swaprouter02;
	Wmatic public wmatic;

    /**
     * @dev Contract constructor
     * @param _swaproute02 uniswap routes contract address
     */
    constructor(ISwapRouter2 _swaproute02) payable {
        swaprouter02 = _swaproute02;
        wmatic = Wmatic(swaprouter02.WETH9());
    }

    function tradeInputSingle(ISwapRouter2.ExactInputSingleParams calldata params, bool isNative) external payable returns(address, uint) {
		checkApproved(IERC20(params.tokenIn), params.amountIn);
		uint amountOut = swaprouter02.exactInputSingle{value: msg.value}(params);
		require(amountOut >= params.amountOutMinimum, "lower than expected output");
		address returnToken = withdrawMatic(params.tokenOut, amountOut, isNative);
		return (returnToken, amountOut);
	}

	function tradeInput(ISwapRouter2.ExactInputParams calldata params, bool isNative) external payable returns(address, uint) {
		(address tokenIn,,) = params.path.decodeFirstPool();
		checkApproved(IERC20(tokenIn), params.amountIn);
		uint amountOut = swaprouter02.exactInput{value: msg.value}(params);
		bytes memory tempPath = params.path;
		address returnToken;
		while (true) {
            bool hasMultiplePools = tempPath.hasMultiplePools();
			// decide whether to continue or terminate
            if (hasMultiplePools) {
                tempPath = tempPath.skipToken();
            } else {
                (,returnToken,) = tempPath.decodeFirstPool();
                break;
            }
		}
		returnToken = withdrawMatic(returnToken, amountOut, isNative);
        return (returnToken, amountOut);
	}

	function checkApproved(IERC20 srcToken, uint256 amount) internal {
		if (msg.value == 0 && srcToken.allowance(address(this), address(swaprouter02)) < amount) {
                srcToken.approve(address(swaprouter02), MAX);
		}
	}

	function withdrawMatic(address tokenOut, uint amountOut, bool isNative) internal returns(address returnToken) {
		if (tokenOut == address(wmatic) && isNative) {
			// convert wmatic to matic
			// recipient in params must be this contract
			wmatic.withdraw(amountOut);
			returnToken = ETH_CONTRACT_ADDRESS;
			transfer(returnToken, amountOut);
		} else {
			returnToken = tokenOut;
		}
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
