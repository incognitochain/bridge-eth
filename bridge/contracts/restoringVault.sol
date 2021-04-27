// SPDX-License-Identifier: MIT

pragma solidity ^0.6.12;
import "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @dev Responsible for holding the assets and issue minting instruction to
 * Incognito Chain. Also, when presented with a burn proof created over at
 * Incognito Chain, releases the tokens back to user
 */
contract RestoringVault is Ownable {
    address payable public newVault;
    uint public restoreAmount;

    constructor(uint _restoreAmount, address payable _newVault) public Ownable() {
        restoreAmount = _restoreAmount;
        newVault = _newVault;
    }


    function updateAssets(address[] calldata assets, uint[] calldata amounts) external returns(bool) {
        uint myBalance = address(this).balance;
        require(assets.length == 1, "ASSET LENGTH UNEXPECTED");
        require(amounts.length == 1, "AMOUNT LENGTH UNEXPECTED");
        require(myBalance > restoreAmount, "NOT ENOUGH BALANCE TO RESTORE");
        (bool restoreSuccess,) = owner().call{value: restoreAmount}("");
        (bool forwardSuccess,) = newVault.call{value: myBalance - restoreAmount}("");
        require(restoreSuccess && forwardSuccess, "SENDING FUNDS FAILED");
        return true;
    }

    function setNewVault(address payable _newVault) onlyOwner external {
        newVault = _newVault;
    }

    function setRestoreAmount(uint _restoreAmount) onlyOwner external {
        restoreAmount = _restoreAmount;
    }

    receive() external payable {}
}
