// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20VotesCompUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol";

/**
 * @dev Interface of the contract capable of checking if an instruction is
 * confirmed over at Incognito Chain
 */
interface Incognito {
    function instructionApproved(
        bool,
        bytes32,
        uint,
        bytes32[] calldata,
        bool[] calldata,
        bytes32,
        bytes32,
        uint[] calldata,
        uint8[] calldata,
        bytes32[] calldata,
        bytes32[] calldata
    ) external view returns (bool);
}

contract PrvVoting is ERC20VotesCompUpgradeable {

    IERC20Upgradeable constant private OLD_PRV = IERC20Upgradeable(address(0x0));
    uint256 constant private BEACON_HEIGHT = 0; // todo: update on testnet/mainnet
    uint constant private MINT_METADATA = 170;

    /**
     * @dev START Storage variables
     */
    mapping(bytes32 => bool) private withdrawed;
    /**
     * @dev Storage slot with the incognito proxy.
     * This is the keccak-256 hash of "eip1967.proxy.incognito." subtracted by 1
     */
    bytes32 private constant _INCOGNITO_SLOT = 0x62135fc083646fdb4e1a9d700e351b886a4a5a39da980650269edd1ade91ffd2;

    /**
     * @dev END Storage variables
     */

    struct BurnInstData {
        uint8 meta; // type of the instruction
        uint8 shard; // ID of the Incognito shard containing the instruction, must be 1
        address token; // ETH address of the token contract (0x0 for ETH)
        address payable to; // ETH address of the receiver of the token
        uint amount; // burned amount (on Incognito)
        bytes32 itx; // Incognito's burning tx
    }

    /**
     * @dev Emitted when function burn called with amount of value to shield to incognito chain
     *
     * Note that `value` may be zero.
     */
    event Deposit(address token, string incognitoAddress, uint amount);
    event Migrate(uint256 amount);

    function initialize(string memory name_, string memory symbol_) external initializer {
        __ERC20_init(name_, symbol_);
    }

    /**
     * @dev Returns the current incognito proxy.
     */
    function _incognito() internal view returns (address icg) {
        bytes32 slot = _INCOGNITO_SLOT;
        // solhint-disable-next-line no-inline-assembly
        assembly {
            icg := sload(slot)
        }
    }

    /**
         * @dev Parses a burn instruction and returns the components
     * @param inst: the full instruction, containing both metadata and body
     */
    function _parseBurnInst(bytes memory inst) internal pure returns (BurnInstData memory) {
        BurnInstData memory data;
        data.meta = uint8(inst[0]);
        data.shard = uint8(inst[1]);
        address token;
        address payable to;
        uint amount;
        bytes32 itx;
        assembly {
        // skip first 0x20 bytes (stored length of inst)
            token := mload(add(inst, 0x22)) // [3:34]
            to := mload(add(inst, 0x42)) // [34:66]
            amount := mload(add(inst, 0x62)) // [66:98]
            itx := mload(add(inst, 0x82)) // [98:130]
        }
        data.token = token;
        data.to = to;
        data.amount = amount;
        data.itx = itx;
        return data;
    }

    /**
     * @dev Verifies that a burn instruction is valid
     * @notice All params except inst are the list of 2 elements corresponding to
     * the proof on beacon and bridge
     * @notice All params are the same as in `withdraw`
     */
    function _verifyInst(
        bytes memory inst,
        uint heights,
        bytes32[] memory instPaths,
        bool[] memory instPathIsLefts,
        bytes32 instRoots,
        bytes32 blkData,
        uint[] memory sigIdxs,
        uint8[] memory sigVs,
        bytes32[] memory sigRs,
        bytes32[] memory sigSs
    ) view internal {
        // Each instruction can only by redeemed once
        bytes32 beaconInstHash = keccak256(abi.encodePacked(inst, heights));

        // Verify instruction on beacon
        require(Incognito(_incognito()).instructionApproved(
                true, // Only check instruction on beacon
                beaconInstHash,
                heights,
                instPaths,
                instPathIsLefts,
                instRoots,
                blkData,
                sigIdxs,
                sigVs,
                sigRs,
                sigSs
            ), "ERC20: invalid input mint data");
    }

    /**
     * @dev Mint prv by providing a burn proof over at Incognito Chain
     * @notice This function takes a burn instruction on Incognito Chain, checks
     * for its validity and mint coressponding prv token
     * @notice This only works when the contract is not Paused
     * @param inst: the decoded instruction as a list of bytes
     * @param heights: the blocks containing the instruction
     * @param instPaths: merkle path of the instruction
     * @param instPathIsLefts: whether each node on the path is the left or right child
     * @param instRoots: root of the merkle tree contains all instructions
     * @param blkData: merkle has of the block body
     * @param sigIdxs: indices of the validators who signed this block
     * @param sigVs: part of the signatures of the validators
     * @param sigRs: part of the signatures of the validators
     * @param sigSs: part of the signatures of the validators
     */
    function mint(
        bytes memory inst,
        uint heights,
        bytes32[] memory instPaths,
        bool[] memory instPathIsLefts,
        bytes32 instRoots,
        bytes32 blkData,
        uint[] memory sigIdxs,
        uint8[] memory sigVs,
        bytes32[] memory sigRs,
        bytes32[] memory sigSs
    ) external virtual returns (bool) {

        require(heights > BEACON_HEIGHT, "ERC20: invalid beacon height");
        require(inst.length >= 130, "ERC20: invalid inst");
        BurnInstData memory data = _parseBurnInst(inst);
        // Check instruction type
        require(data.meta == MINT_METADATA && data.shard == 1, "ERC20: invalid inst's data");
        // Check token address
        require(data.token == address(this), "ERC20: invalid token");
        // Not withdrawed
        require(!withdrawed[data.itx], "ERC20: tx is already used");
        withdrawed[data.itx] = true;
        // verify proof
        _verifyInst(
            inst,
            heights,
            instPaths,
            instPathIsLefts,
            instRoots,
            blkData,
            sigIdxs,
            sigVs,
            sigRs,
            sigSs
        );
        _mint(data.to, data.amount);

        return true;
    }

    /**
     * @dev Destroys `amount` tokens from `account`, reducing the
     * total supply.
     *
     * Emits a {Transfer} event with `to` set to the zero address.
     *
     * Requirements:
     *
     * - `account` cannot be the zero address.
     * - `account` must have at least `amount` tokens.
     */
    function burn(string calldata incognitoAddress, uint256 amount) external virtual returns (bool) {
        _burn(_msgSender(), amount);

        emit Deposit(address(this), incognitoAddress, amount);

        return true;
    }

    function migrate(uint256 amount) external {
        OLD_PRV.transferFrom(_msgSender(), address(0x0), amount);
        _mint(_msgSender(), amount);

        emit Migrate(amount);
    }

    /**
     * @dev Get the address `account` is currently delegating to. By default will be itself.
     */
    function delegates(address account) public view override returns (address) {
        address temp = super.delegates(account);
        return temp == address(0) ? account : temp;
    }

    function decimals() public view override returns (uint8) {
        return 9;
    }
}
