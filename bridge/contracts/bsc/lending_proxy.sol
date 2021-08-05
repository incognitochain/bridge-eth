// SPDX-License-Identifier: MIT
pragma solidity ^0.6.12;
pragma experimental ABIEncoderV2;

import './IERC20.sol';

// To shield token again
interface Vault {
    function deposit(string calldata incognitoAddress) external;
    function depositERC20(address token, uint amount, string calldata incognitoAddress) external;
}

contract Agency {
    address factory;
    address constant public ETH_TOKEN = 0x0000000000000000000000000000000000000000;
    uint constant public APPROVE_MAX = 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff;

    modifier onlyFactory() {
        require(msg.sender == factory, "only factory");
        _;
    }
    
    constructor(address _factory) public {
        factory = _factory;
    }
    
    function _execute(uint ethAmount, address destination, bytes calldata data) internal {
        (bool success, ) = destination.call{value: ethAmount}(data);
        require(success, "Agency: call to external contract failed");
    }
    
    function executeWithEth(address[] calldata destinations, uint[] calldata amounts, bytes[] calldata datas) onlyFactory payable external {
        require(destinations.length == datas.length, "invalid input datas");
        uint checksum;
        for (uint i = 0; i < destinations.length; i++) {
            _execute(amounts[i], destinations[i], datas[i]);
            checksum += amounts[i];
        }

        require(checksum == msg.value, "invalid input amount");
    }

    function executeWithERC20(address[] calldata destinations, bytes[] calldata datas) onlyFactory external {
        require(destinations.length == datas.length, "invalid input datas");

        for (uint i = 0; i < destinations.length; i++) {
            _execute(0, destinations[i], datas[i]);
        }
    }
}

contract Verifier {
    // data signed storage
    mapping(bytes32 => bool) public sigDataUsed;

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
     * @dev verify sign data
     */
     function verifySignData(bytes memory data, bytes memory signData) internal returns(address){
        bytes32 hash = keccak256(data);
        require(!sigDataUsed[hash], "data signed already used");
        address verifier = sigToAddress(signData, hash);
        // reject when verifier equals zero
        require(verifier != address(0x0), "invalid signature");
        // mark data hash of sig as used
        sigDataUsed[hash] = true;

        return verifier;
     }
}

contract Factory is Verifier {
    string constant PREFIX_SIGNATURE = "Factory";
    mapping(address => address) public agencies;
    address constant public ETH_TOKEN = 0x0000000000000000000000000000000000000000;

    struct PreSignData {
        string prefix;
        address token;
        bytes timestamp;
    }

    function newPreSignData(address token, bytes calldata timestamp) pure internal returns (PreSignData memory) {
        PreSignData memory psd = PreSignData(PREFIX_SIGNATURE, token, timestamp);
        return psd;
    }

    function _order(address agency, bytes calldata data) internal {
        (bool success, ) = agency.call{value: msg.value}(data);
        require(success, "Factory: call to external contract failed");
    }

    /**
     * @dev check agency exist or not then create new or query based on result.
     */
    function _agency(bytes calldata signData, address token, bytes calldata data, bytes calldata timestamp) internal returns(address agency) {
        address verifier = verifySignData(abi.encode(newPreSignData(token, timestamp), data), signData);
        if (agencies[verifier] == address(0x0)) {
            agency = address(new Agency(address(this)));
            agencies[verifier] = agency;
        } else {
            agency = agencies[verifier];
        }
    }
 
    function orderAgency(bytes calldata signData, address token, bytes calldata data, bytes calldata timestamp) external payable returns(address, uint) {
        address agency = _agency(signData, token, data, timestamp);
        _order(agency, data);
        
        return (address(0x0), 0);
    }   

    function orderAgencyWithToken(bytes calldata signData, address token, bytes calldata data, bytes calldata timestamp) external payable returns(address, uint) {
        address agency = _agency(signData, token, data, timestamp);
        if (token != ETH_TOKEN) {
            IERC20 tokenInst = IERC20(token);
            tokenInst.transfer(agency, tokenInst.balanceOf(address(this)));
            require(checkSuccess(), "transfer token failed");
        }
        _order(agency, data);
        
        return (address(0x0), 0);
    }   
    
    
    /**
     * @dev Check if transfer() and transferFrom() of ERC20 succeeded or not
     * This check is needed to fix https://github.com/ethereum/solidity/issues/4116
     * This function is copied from https://github.com/AdExNetwork/adex-protocol-eth/blob/master/contracts/libs/SafeERC20.sol
     */
    function checkSuccess() private pure returns (bool) {
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