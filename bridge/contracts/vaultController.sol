pragma solidity ^0.6.6;

contract vaultController {

    // contain the address of contract
    mapping (bytes32 => address) registry;
    address public admin;

    modifier onlyAdmin() {
        require(msg.sender == admin, "Only addmin do change on this contract!");
        _;
    }

    event SetAddress(bytes32, address);
    event newAdmin(address);

    constructor(address _admin) public {
        admin = _admin;
    }

    /**
     * @dev Return address of contract based on given name
     */
    function lookup(bytes32 key) public view returns(address) {
        return registry[key];
    }

    function setAddress(bytes32 key, address contractAddress) external onlyAdmin returns(bool) {
        registry[key] = contractAddress;

        emit SetAddress(key, contractAddress);
        return true;
    }

    function getVault() public view returns (address) {
        return lookup("Vault");
    }

    function transferAdminship(address _newAdmin) public onlyAdmin returns (bool) {
        admin = _newAdmin;

        emit newAdmin(admin);
        return true;
    }
}