pragma solidity 0.6.6;

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

contract UniswapV2Adapter is TradeUtils {
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

    function trade(address[] calldata paths, uint srcQty , uint amountOutMin, uint timeout) external payable returns (address, uint) {
        require(paths.length > 1);
        address destToken = paths[paths.length - 1];
        require(paths[0] != destToken);
        uint[] memory amounts;
        if (paths[0] != wETH) {
            // approve
            approve(IERC20(paths[0]), address(uniswapV2), srcQty);
            if (destToken != wETH) { // token to token.
                amounts = tokenToToken(paths, srcQty, amountOutMin, timeout);
            } else {
                amounts = tokenToEth(paths, srcQty, amountOutMin, timeout);
                destToken = address(ETH_CONTRACT_ADDRESS);
            }
        } else {
            amounts = ethToToken(paths, srcQty, amountOutMin, timeout);
        }
        require(amounts.length > 1);
        require(amounts[amounts.length - 1] >= amountOutMin && amounts[0] == srcQty);
        return (destToken, amounts[amounts.length - 1]);
    }

    function ethToToken(address[] memory path, uint srcQty, uint amountOutMin, uint timeout) internal returns (uint[] memory) {
        return uniswapV2.swapExactETHForTokens{value: srcQty}(amountOutMin, path, msg.sender, timeout);
    }

    function tokenToEth(address[] memory path, uint srcQty, uint amountOutMin, uint timeout) internal returns (uint[] memory) {
        return uniswapV2.swapExactTokensForETH(srcQty, amountOutMin, path, msg.sender, timeout);
    }

    function tokenToToken(address[] memory path, uint srcQty, uint amountOutMin, uint timeout) internal returns (uint[] memory) {
        return uniswapV2.swapExactTokensForTokens(srcQty, amountOutMin, path, msg.sender, timeout);
    }

    /**
     * @dev Given an input asset amount and an array of token addresses, calculates all subsequent maximum output token.
     * @param paths the path from source toke to dest token
     * @param srcQty amount of source tokens
     */
    function getAmountsOut(address[] calldata paths, uint srcQty) external view returns(uint[] memory) {
        return uniswapV2.getAmountsOut(srcQty, paths);
    }
}
