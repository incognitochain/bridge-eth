pragma solidity ^0.6.6;

import './IERC20.sol';

contract TradeUtilsCompound {
	IERC20 constant public ETH_CONTRACT_ADDRESS = IERC20(0x0000000000000000000000000000000000000000);
	address payable constant public incognitoSmartContract = 0x8C4B922E2f7d6d1ABA41d79C47F497F6F54e0Af8;

	modifier isIncognitoSmartContract {
	    require(msg.sender == incognitoSmartContract);
	    _;
	}

	function balanceOf(IERC20 token) internal view returns (uint256) {
		if (token == ETH_CONTRACT_ADDRESS) {
			return address(this).balance;
		}
        return token.balanceOf(address(this));
    }

	function transfer(IERC20 token, uint amount) internal {
		if (token == ETH_CONTRACT_ADDRESS) {
			require(address(this).balance >= amount);
			incognitoSmartContract.transfer(amount);
		} else {
			token.transfer(incognitoSmartContract, amount);
			require(checkSuccess());
		}
	}

	function approve(IERC20 token, address proxy, uint amount) internal {
		if (token != ETH_CONTRACT_ADDRESS) {
			token.approve(proxy, 0);
			require(checkSuccess());
			token.approve(proxy, amount);
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
}

interface CTokenInterface {
    function redeem(uint redeemTokens) external returns (uint);
    function redeemUnderlying(uint redeemAmount) external returns (uint);
    function borrow(uint borrowAmount) external returns (uint);
}

interface CEther is CTokenInterface {
    function mint() external payable;
    function repayBorrow() external payable;
    function repayBorrowBehalf(address borrower) external;
    function liquidateBorrow(address borrower, CTokenInterface cTokenCollateral) external payable;
}

interface CErc20 is CTokenInterface {
    function underlying() external returns(address);
    function mint(uint mintAmount) external returns (uint);
    function repayBorrow(uint repayAmount) external returns (uint);
    function repayBorrowBehalf(address borrower, uint repayAmount) external returns (uint);
    function liquidateBorrow(address borrower, uint repayAmount, CTokenInterface cTokenCollateral) external returns (uint);
}

interface Comptroller {
    function enterMarkets(address[] calldata cTokens) external returns (uint[] memory);
    function exitMarket(address cTokenAddress) external returns (uint);
}

contract CompoundAgentLogic is TradeUtilsCompound {

    Comptroller public constant comptroller = Comptroller(0x1f5D7F3CaAC149fE41b8bd62A3673FE6eC0AB73b);
    CEther public constant cEther = CEther(0xf92FbE0D3C0dcDAE407923b2Ac17eC223b1084E4);

    constructor() public {}

    // receive function which allows transfer eth.
    receive() external payable {}

    /**
     * @dev Call mint func to compound 
     * @param cToken: cToken of compound
     * @param amount: total to mint
     * @return bool: token address, amount recieved
     */
    function mint(address cToken, uint amount) external payable returns (address, uint) {
        uint amountRecieved = balanceOf(IERC20(cToken));
        if(cToken == address(cEther)) {
            CEther(cToken).mint{value: msg.value}();
        } else {
            approve(IERC20(CErc20(cToken).underlying()), cToken, amount);
            require(CErc20(cToken).mint(amount) == 0);
        }
        amountRecieved = balanceOf(IERC20(cToken)) - amountRecieved ;
        transfer(IERC20(cToken), amountRecieved);
        return (cToken, amountRecieved);
    }

    /**
     * @dev Call borrow func to compound 
     * @param cToken: cToken of compound
     * @param amount: total to borrow
     * @param addToMarkets: add tokens to market as collateral
     * @return bool: token address, amount recieved
     */
    function borrow(address cToken, uint amount, address[] calldata addToMarkets) external returns (address, uint) {
        if(addToMarkets.length > 0) {
            comptroller.enterMarkets(addToMarkets);
        }
        require(CTokenInterface(cToken).borrow(amount) == 0);
        uint amountAfter = 0;
        if (cToken == address(cEther)) {
            amountAfter = balanceOf(ETH_CONTRACT_ADDRESS);
            transfer(ETH_CONTRACT_ADDRESS, amountAfter);
            return (address(ETH_CONTRACT_ADDRESS), amountAfter); 
        } else {
            amountAfter = balanceOf(IERC20(CErc20(cToken).underlying()));
            transfer(IERC20(CErc20(cToken).underlying()), amountAfter);
            return (CErc20(cToken).underlying(), amountAfter);
        }
    }

    /**
     * @dev Call borrow func to compound 
     * @param cToken: cToken of compound
     * @param amount: total to borrow
     * @param addToMarkets: add token to market as collateral
     * @return bool: token address, amount recieved
     */
    function borrowByMultiCollateral(address cToken, uint amount, address[] calldata addToMarkets) external returns (address[] memory, uint[] memory) {
        if(addToMarkets.length > 0) {
            comptroller.enterMarkets(addToMarkets);
        }
        address[] memory addressAfter = new address[](1);
        require(CTokenInterface(cToken).borrow(amount) == 0);
        uint amountAfter = 0;
        if (cToken == address(cEther)) {
            amountAfter = balanceOf(ETH_CONTRACT_ADDRESS);
            transfer(ETH_CONTRACT_ADDRESS, amountAfter);
            addressAfter[0] = address(ETH_CONTRACT_ADDRESS);
        } else {
            amountAfter = balanceOf(IERC20(CErc20(cToken).underlying()));
            transfer(IERC20(CErc20(cToken).underlying()), amountAfter);
            addressAfter[0] = CErc20(cToken).underlying();
        }
        uint[] memory balAfter = new uint[](1);
        balAfter[0] = amountAfter;
        return (addressAfter,  balAfter);
    }

    /**
     * @dev Call mint func to compound 
     * @param cToken: cToken of compound
     * @param amount: total to redeem
     * @param isredeemUnderlying: whether redeem cToken or native coin
     * @param exitToMarkets: remove coin from market
     * @return bool: token address, amount recieved
     */
    function redeem(address cToken, uint amount, bool isredeemUnderlying, address exitToMarkets) external returns (address, uint) {
        if(exitToMarkets != address(0x0)) {
            comptroller.exitMarket(exitToMarkets);   
        }
        if(isredeemUnderlying) {
            require(CTokenInterface(cToken).redeemUnderlying(amount) == 0);
        } else {
            require(CTokenInterface(cToken).redeem(amount) == 0);
        }
        uint amountAfter = 0;
        if (cToken == address(cEther)) {
            amountAfter = balanceOf(ETH_CONTRACT_ADDRESS);
            transfer(ETH_CONTRACT_ADDRESS, amountAfter);
            return (address(ETH_CONTRACT_ADDRESS), amountAfter); 
        } else {
            amountAfter = balanceOf(IERC20(CErc20(cToken).underlying()));
            transfer(IERC20(CErc20(cToken).underlying()), amountAfter);
            return (CErc20(cToken).underlying(), amountAfter);
        }
    }

    /**
     * @dev Call repayBorrow func to compound 
     * @param cToken: cToken of compound
     * @param amount: total to repay
     * @return bool: token address, amount (default 0)
     */
    function repayBorrow(address cToken, uint amount) external payable returns (address, uint) {
        if(cToken == address(cEther)) {
            CEther(cToken).repayBorrow{value: msg.value}();
        } else {
            approve(IERC20(CErc20(cToken).underlying()), cToken, amount);
            require(CErc20(cToken).repayBorrow(amount) == 0);
        }

        return (cToken, 0);
    }

    // TODO:
    // function repayBorrowBehalf() external payable onlyProxyCompound return (address, uint) {}

    /**
     * @dev Call liquidateBorrow func to compound 
     * @param cToken: cToken of compound
     * @param borrower: the address can't pay this borrow
     * @param repayAmount: total to repay
     * @param cTokenCollateral: the address of ctoken 
     * @return bool: token address, amount recieved
     */
    function liquidateBorrow(address cToken, address borrower, uint repayAmount, address cTokenCollateral) external payable returns (address, uint) {
        uint amountRecieved = balanceOf(IERC20(cTokenCollateral));
        if(cToken == address(cEther)) {
            CEther(cToken).liquidateBorrow{value: msg.value}(borrower, CTokenInterface(cTokenCollateral));
        } else {
            approve(IERC20(CErc20(cToken).underlying()), cToken, repayAmount);
            require(CErc20(cToken).liquidateBorrow(borrower, repayAmount, CTokenInterface(cTokenCollateral)) == 0);
        }
        amountRecieved = balanceOf(IERC20(cTokenCollateral)) - amountRecieved;
        transfer(IERC20(cTokenCollateral), amountRecieved);
        return (cTokenCollateral, amountRecieved);
    }
}

contract CompoundAgent {
    
    address public proxyCompound;
    address public compoundAgentLogic;
    
    modifier onlyProxyCompound () {
        require(msg.sender == proxyCompound, "Only proxy compound allowed to call this agent!");
        _;
    }
    
    constructor(address _proxyCompound, address _compoundAgentLogic) public {
        proxyCompound = _proxyCompound;
        compoundAgentLogic = _compoundAgentLogic;
    }
    
    // Receive function which allows transfer eth.
    receive() external payable {}
    
    /**
     * @notice External method to delegate execution to another contract
     * @dev It returns to the external caller whatever the implementation returns or forwards reverts
     * @param data The raw data to delegatecall
     * @return The returned bytes from the delegatecall
     */
    function delegateCall(bytes calldata data) external payable onlyProxyCompound returns (bytes memory) {
        (bool success, bytes memory returnData) = compoundAgentLogic.delegatecall(data);
        assembly {
            if eq(success, 0) {
                revert(add(returnData, 0x20), returndatasize())
            }
        }
        
        return returnData;
    }
    
    /**
     * @dev update proxy compound  
     */
    function updateProxy(address _proxyCompound) external onlyProxyCompound returns (address, uint) {
        proxyCompound = _proxyCompound;
        return (address(0x0), 0);
    }
    
    /**
     * @dev update compoundAgentLogic compound  
     */
    function updateAgentLogic(address _compoundAgentLogic) external onlyProxyCompound returns (address, uint) {
        compoundAgentLogic = _compoundAgentLogic;
        return (address(0x0), 0);
    }
}

import './trade_utils.sol';
contract CompoundProxy is TradeUtils {
     
    mapping(bytes32 => bool) public sigDataUsed;
    mapping(address => address) public agents;
    
    address public admin;
    address public compoundAgentLogic;

    constructor(address payable _incognitoSmartContract, address  _compoundAgentLogic, address _admin) public {
        incognitoSmartContract = _incognitoSmartContract;
        compoundAgentLogic = _compoundAgentLogic;
        admin = _admin;
    }
    
    modifier onlyAdmin() {
        require(msg.sender == admin, "Only admin can call this function!");
        _;
    }
    
    event UpdateVaultCompound(address);
    event UpdateAgentLogic(address);
    
    /**
     * @dev Checks if a caller already has agent or not
     * @notice First, we check inside the storage of this contract itself. If the
     * hash has been used before, we return the result. Otherwise, we query
     * previous vault recursively until the first compound proxy
     * @param caller: address of caller
     * @return address: agent contract address of this caller
     */
    function isAgentExist(address caller) internal returns(address) {
        if(agents[caller] == address(0x0)) {
            agents[caller] = address(new CompoundAgent(address(this), compoundAgentLogic));
        }
        return agents[caller];
    }
    
    /**
     * @dev Checks if a sig data has been used before
     * @notice First, we check inside the storage of this contract itself. If the
     * hash has been used before, we return the result. Otherwise, we query
     * previous vault recursively until the first Vault (prevVault address is 0x0)
     * @param hash: of the sig data
     * @return bool: whether the sig data has been used or not
     */
    function isSigDataUsed(bytes32 hash) public view returns(bool) {
        if (sigDataUsed[hash]) {
            return true;
        }
        return false;
    }
    
    /**
     * @dev Call agent contract to interact with compound
     * @param srcToken: The address of token
     * @param callData: The function and param to foward to agent
     * @param timestamp: Random data for signing 
     * @param signData: The signature of owner
     * @return bool: whether the sig data has been used or not
     */
    function execute(
        IERC20 srcToken,
        uint amount,
        bytes memory callData,
        bytes memory timestamp,
        bytes memory signData
    ) public payable isIncognitoSmartContract returns(address, uint)  {
        //verify ower signs data from input
        address verifier = verifySignData(abi.encodePacked(srcToken, callData, timestamp), signData);
        address agent = isAgentExist(verifier);
        if(srcToken != ETH_CONTRACT_ADDRESS) {
            srcToken.transfer(agent, amount);
            checkSuccess();
        }
        (bool success, bytes memory result) = agent.call{value: msg.value}(callData);
        require(success);
        bytes memory decodeResult = abi.decode(result, (bytes));
        
        return abi.decode(decodeResult, (address, uint));
    }

    /**
     * @dev : borrow by sending many coins to 
     */
    function executeMulti(
        IERC20[] memory srcTokens,
        uint[] memory amounts,
        bytes memory callData,
        bytes memory timestamp,
        bytes memory signData
    ) public payable isIncognitoSmartContract returns(address[] memory, uint[] memory) {
        require(srcTokens.length == amounts.length);
        //verify ower signs data from input
        address verifier = verifySignData(abi.encodePacked(callData, timestamp), signData);
        address agent = isAgentExist(verifier);
        for(uint i = 0; i < srcTokens.length; i++) {
            if(srcTokens[i] != ETH_CONTRACT_ADDRESS) {
                srcTokens[i].transfer(agent, amounts[i]);
                checkSuccess();
            }
        }
        (bool success, bytes memory result) = agent.call{value: msg.value}(callData);
        require(success);
        bytes memory decodeResult = abi.decode(result, (bytes));
        
        return abi.decode(decodeResult, (address[], uint[]));
    }

    /**
     * @dev generate address from signature data and hash.
     */
    function sigToAddress(bytes memory signData, bytes32 hash) public pure returns (address) {
        bytes32 s;
        bytes32 r;
        uint8 v;
        assembly {
            r := mload(add(signData, 0x20))
            s := mload(add(signData, 0x40))
        }
        v = uint8(signData[64]) + 27;
        return ecrecover(hash, v, r, s);
    }

    /**
     * @dev return agent contract address.
     */
    function getAgentAddress(address caller) public view returns (address) {
        return agents[caller];
    }

    /**
     * @dev verify sign data
     */
    function verifySignData(bytes memory data, bytes memory signData) internal returns(address){
        bytes32 hash = keccak256(data);
        require(!isSigDataUsed(hash));
        address verifier = sigToAddress(signData, hash);
        // mark data hash of sig as used
        sigDataUsed[hash] = true;
        
        return verifier;
    }
    
    /**
     * @dev update vault address
     */
    function updateVault(address payable _incognitoSmartContract) external onlyAdmin {
        incognitoSmartContract = _incognitoSmartContract;
        
        emit UpdateVaultCompound(_incognitoSmartContract);
    }
    
    /**
     * @dev update compoundAgentLogic compound  
     */
    function updateAgentLogic(address _compoundAgentLogic) external onlyAdmin {
        compoundAgentLogic = _compoundAgentLogic;
        
        emit UpdateAgentLogic(_compoundAgentLogic);
    }
}

