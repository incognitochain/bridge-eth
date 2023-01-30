pragma solidity ^0.6.12;

import './trade_utils.sol';
import './IERC20.sol';

interface UniswapV2 {
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
  function getAmountsOut(uint amountIn, address[] calldata path) external view returns (uint[] memory amounts);
}

contract UniswapV2Trade is TradeUtils, Executor {
    // Variables
    UniswapV2 public uniswapV2;
    address public wETH;

    // Functions
    /**
     * @dev Contract constructor
     * @param _uniswapV2 uniswap routes contract address
     */
    constructor(UniswapV2 _uniswapV2) public {
        uniswapV2 = _uniswapV2;
        wETH = uniswapV2.WETH();
    }

    // Reciever function which allows transfer eth.
    receive() external payable {}

    function trade(address[] memory path, uint amountOutMin) public payable returns (address, uint) {
        require(path.length >= 2, "Proxy: invalid path");
        uint256 swapAmount = msg.value > 0 ? msg.value : balanceOf(IERC20(path[0]));
        uint[] memory amounts;
        bool isSwapForNative = false;
        if (msg.value == 0) {
            // approve
            approve(IERC20(path[0]), address(uniswapV2), swapAmount);
            if (path[path.length - 1] != wETH) { // token to token.
                amounts = tokenToToken(path, swapAmount, amountOutMin);
            } else {
                amounts = tokenToEth(path, swapAmount, amountOutMin);
                isSwapForNative = true;
            }
        } else {
            amounts = ethToToken(path, swapAmount, amountOutMin);
        }
        require(amounts.length >= 2, "Proxy: invalid response values");
        require(amounts[amounts.length - 1] >= amountOutMin && amounts[0] == swapAmount);
        // ETH_CONTRACT_ADDRESS is a address present for eth native
        return (isSwapForNative ? address(ETH_CONTRACT_ADDRESS) : path[path.length - 1], amounts[amounts.length - 1]);
    }

    function ethToToken(address[] memory path, uint srcQty, uint amountOutMin) internal returns (uint[] memory) {
        return uniswapV2.swapExactETHForTokens{value: srcQty}(amountOutMin, path, msg.sender, now + 600);
    }

    function tokenToEth(address[] memory path, uint srcQty, uint amountOutMin) internal returns (uint[] memory) {
        return uniswapV2.swapExactTokensForETH(srcQty, amountOutMin, path, msg.sender, now + 600);
    }

    function tokenToToken(address[] memory path, uint srcQty, uint amountOutMin) internal returns (uint[] memory) {
        return uniswapV2.swapExactTokensForTokens(srcQty, amountOutMin, path, msg.sender, now + 600);
    }

    /**
     * @dev Given an input asset amount and an array of token addresses, calculates all subsequent maximum output token.
     * @param srcToken source token contract address
     * @param srcQty amount of source tokens
     * @param destToken destination token contract address
     */
    function getAmountsOut(address srcToken, uint srcQty, address destToken) external view returns(uint[] memory) {
        address[] memory path = new address[](2);
        path[0] = srcToken;
        path[1] = destToken;
        return uniswapV2.getAmountsOut(srcQty, path);
    }
}
