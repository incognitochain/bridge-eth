pragma solidity ^0.6.6;
pragma experimental ABIEncoderV2;

import "./IERC20.sol";
import "./pause.sol";

contract Locker is AdminPausable {
    address constant public ETH_TOKEN = 0x0000000000000000000000000000000000000000;

    bool public notEntered = true;
    address public vault;
    address public nextLocker;

    /**
     * @dev Prevents a contract from calling itself, directly or indirectly.
     * Calling a `nonReentrant` function from another `nonReentrant`
     * function is not supported. It is possible to prevent this from happening
     * by making the `nonReentrant` function external, and make it call a
     * `private` function that does the actual work.
     */
    modifier nonReentrant() {
        // On the first call to nonReentrant, notEntered will be true
        require(notEntered, "reentranted");

        // Any calls to nonReentrant after this point will fail
        notEntered = false;

        _;

        // By storing the original value once again, a refund is triggered (see
        // https://eips.ethereum.org/EIPS/eip-2200)
        notEntered = true;
    }

    modifier onlyValidVault() {
        require(msg.sender == vault && vault != address(0), "not vault");
        _;
    }

    constructor(address _admin, address _vault) public AdminPausable(_admin) {
        vault = _vault;
    }

    /**
     * @dev Save the address of the Vault that is able to move assets out of this Locker
     * @notice Can only be called by admin
     * @param newVault: address of the Vault
     */
    function changeVault(address newVault) public onlyAdmin nonReentrant {
        vault = newVault;
    }

    /**
     * @dev Set the next locker that will inherit the funds from this locker
     * @notice This only works when the contract is Paused
     * @notice Can only be called by admin
     * @param _nextLocker: address of the next locker contract
     */
    function setNextLocker(address payable _nextLocker) public onlyAdmin isPaused nonReentrant {
        require(_nextLocker != address(0), "null nextLocker");
        nextLocker = _nextLocker;
    }

    /**
     * @dev Move assets to the next Locker
     * @notice This should only be used when we are migrating to a new Locker
     * For normal usage (withdrawing), use `give` since it checks for transfer's results
     * @notice This only works when the contract is Paused
     * @notice Can only be called by admin
     * @param assets: list of address of the assets to move
     */
    function unlock(address[] memory assets) public onlyAdmin isPaused nonReentrant {
        require(nextLocker != address(0), "null nextLocker");
        for (uint i = 0; i < assets.length; i++) {
            // No need to check for result; if it fails, move on
            if (assets[i] == ETH_TOKEN) {
                nextLocker.call{value: address(this).balance}("");
            } else {
                uint bal = IERC20(assets[i]).balanceOf(address(this));
                if (bal > 0) {
                    IERC20(assets[i]).transfer(nextLocker, bal);
                }
            }
        }
    }

    /**
     * @dev Transfer some amount of asset to a specified address
     * @notice This should be used for normal operation (withdrawing).
     * For migrating funds, consider using `unlock` since it can transfer many assets
     * in the same transaction.
     * @notice This only works when the contract is not Paused
     * @notice Can only be called by Vault
     * @param to: address to transfer asset to
     * @param token: asset to transfer (0x0 for ETH)
     * @param amount: amount to transfer
     */
    function give(address to, address token, uint amount) public onlyValidVault isNotPaused nonReentrant {
        if (token == ETH_TOKEN) {
            require(address(this).balance >= amount, "not enough eth");
            (bool success, ) = to.call{value: amount}("");
            require(success, "failed to give eth");
        } else {
            uint bal = IERC20(token).balanceOf(address(this));
            require(bal >= amount, "not enough erc20");
            IERC20(token).transfer(to, amount);
            require(checkSuccess(), "failed to give erc20");
        }
    }

    /**
     * @dev Payable receive function to receive Ether
     */
    receive() external payable {}

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