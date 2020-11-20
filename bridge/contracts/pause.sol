pragma solidity 0.6.6;

contract AdminPausable {
    
    /**
     * @dev Storage slot with the admin of the contract.
     * This is the keccak-256 hash of "eip1967.proxy.admin" subtracted by 1, and is
     * validated in the constructor.
     */
    bytes32 private constant _ADMIN_SLOT = 0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103;
    
    bool public paused;
    uint public expire;

    event Paused(address pauser);
    event Unpaused(address pauser);
    event Extend(uint ndays);

    modifier onlyAdmin() {
        require(msg.sender == _admin(), "not admin");
        _;
    }

    modifier isPaused() {
        require(paused, "not paused right now");
        _;
    }

    modifier isNotPaused() {
        require(!paused, "paused right now");
        _;
    }

    modifier isNotExpired() {
        require(block.timestamp < expire, "expired");
        _;
    }
    
    /**
     * @dev Returns the current admin.
     */
    function _admin() internal view returns (address adm) {
        bytes32 slot = _ADMIN_SLOT;
        // solhint-disable-next-line no-inline-assembly
        assembly {
            adm := sload(slot)
        }
    }

    function extend(uint n) public onlyAdmin isNotExpired {
        require(n < 366, "cannot extend for too long"); // To prevent overflow
        expire = expire + n * 1 days;
        emit Extend(n);
    }

    function pause() public onlyAdmin isNotPaused isNotExpired {
        paused = true;
        emit Paused(msg.sender);
    }

    function unpause() public onlyAdmin isPaused {
        paused = false;
        emit Unpaused(msg.sender);
    }
}