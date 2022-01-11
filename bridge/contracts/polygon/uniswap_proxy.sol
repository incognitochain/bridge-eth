// SPDX-License-Identifier: GPL-2.0-or-later
pragma solidity =0.7.6;
import './IERC20.sol';
import '@uniswap/v3-periphery/contracts/libraries/Path.sol';
import '@uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol';

contract PancakeProxy {
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
