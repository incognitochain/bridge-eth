// SPDX-License-Identifier: MIT

pragma solidity ^0.6.12;
pragma experimental ABIEncoderV2;

import "./IERC721.sol";

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

/**
 * @dev Responsible for holding the assets and issue minting instruction to
 * Incognito Chain. Also, when presented with a burn proof created over at
 * Incognito Chain, releases the tokens back to user
 */
contract VaultNFT {

    /**
     * @dev Storage slot with the incognito proxy.
     * This is the keccak-256 hash of "eip1967.proxy.incognito." subtracted by 1
     */
    bytes32 private constant _INCOGNITO_SLOT = 0x62135fc083646fdb4e1a9d700e351b886a4a5a39da980650269edd1ade91ffd2;

    // metadata
    uint8 constant public CURRENT_NETWORK_ID = 1; // Ethereum
    uint8 constant public BURN_REQUEST_METADATA_TYPE = 241;
    uint8 constant public BURN_TO_CONTRACT_REQUEST_METADATA_TYPE = 243;
    uint8 constant public BURN_CALL_REQUEST_METADATA_TYPE = 158;

    /**
     * @dev Storage variables for Vault
     * This section is APPEND-ONLY, in order to preserve upgradeability
     * since we use Proxy Pattern
     */
    mapping(bytes32 => bool) public withdrawed;
    bool public notEntered;
    bool public isInitialized;
    /**
    * @dev Added in Storage Layout version : 1.0
    */

    struct BurnInstData {
        uint8 meta; // type of the instruction
        uint8 shard; // ID of the Incognito shard containing the instruction, must be 1
        address token; // ETH address of the token contract
        address to; // ETH address of the receiver of the token
        uint tokenId; // NFT id burned
        bytes32 itx; // Incognito's burning tx
    }

    // error code
    enum Errors {
        EMPTY,
        NO_REENTRANCE,
        INVALID_DATA,
        ALREADY_INITIALIZED,
        ALREADY_USED
    }

    event Deposit(address token, string incognitoAddress, uint id);
    event Withdraw(address token, address to, uint amount);

    /**
     * @dev Prevents a contract from calling itself, directly or indirectly.
     * Calling a `nonReentrant` function from another `nonReentrant`
     * function is not supported. It is possible to prevent this from happening
     * by making the `nonReentrant` function external, and make it call a
     * `private` function that does the actual work.
     */
    modifier nonReentrant() {
        // On the first call to nonReentrant, notEntered will be true
        require(notEntered, errorToString(Errors.NO_REENTRANCE));

        // Any calls to nonReentrant after this point will fail
        notEntered = false;

        _;

        // By storing the original value once again, a refund is triggered (see
        // https://eips.ethereum.org/EIPS/eip-2200)
        notEntered = true;
    }

    /**
     * @dev Creates new Vault to hold assets for Incognito Chain
     * After migrating all assets to a new Vault, we still need to refer
     * back to previous Vault to make sure old withdrawals aren't being reused
     */
    function initialize() external {
        require(!isInitialized, errorToString(Errors.ALREADY_INITIALIZED));
        isInitialized = true;
        notEntered = true;
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
     * @dev Makes a NFT deposit to the vault to mint pNFT over at Incognito Chain
     * @notice This only works when the contract is not Paused
     * @notice Before calling this function, NFT must be allowed to
     * tranfer from msg.sender to this contract
     * @param token: address of the ERC721 token
     * @param id: the id of NFT
     * @param incognitoAddress: Incognito Address to receive pNFT
     */
    function depositNFT(IERC721 token, uint id, string calldata incognitoAddress) external nonReentrant {
        token.transferFrom(msg.sender, address(this), id);

        emit Deposit(address(token), incognitoAddress, id);
    }

    /**
     * @dev Checks if a burn proof has been used before
     * @notice First, we check inside the storage of this contract itself. If the
     * hash has been used before, we return the result. Otherwise, we query
     * previous vault recursively until the first Vault (prevVault address is 0x0)
     * @param hash: of the burn proof
     * @return bool: whether the proof has been used or not
     */
    function isWithdrawed(bytes32 hash) public view returns(bool) {
        return withdrawed[hash];
    }

    /**
     * @dev Parses a burn instruction and returns the components
     * @param inst: the full instruction, containing both metadata and body
     */
    function parseBurnInst(bytes memory inst) public pure returns (BurnInstData memory) {
        BurnInstData memory data;
        data.meta = uint8(inst[0]);
        data.shard = uint8(inst[1]);
        address token;
        address payable to;
        uint tokenId;
        bytes32 itx;
        assembly {
        // skip first 0x20 bytes (stored length of inst)
            token := mload(add(inst, 0x22)) // [3:34]
            to := mload(add(inst, 0x42)) // [34:66]
            tokenId := mload(add(inst, 0x62)) // [66:98]
            itx := mload(add(inst, 0x82)) // [98:130]
        }
        data.token = token;
        data.to = to;
        data.tokenId = tokenId;
        data.itx = itx;
        return data;
    }

    /**
     * @dev Verifies that a burn instruction is valid
     * @notice All params except inst are the list of 2 elements corresponding to
     * the proof on beacon and bridge
     * @notice All params are the same as in `withdraw`
     */
    function verifyInst(
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
            ), errorToString(Errors.INVALID_DATA));
    }

    /**
     * @dev Withdraws pNFT by providing a burn proof over at Incognito Chain
     * @notice This function takes a burn instruction on Incognito Chain, checks
     * for its validity and returns the token back to ETH chain
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
    function withdraw(
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
    ) public nonReentrant {
        require(inst.length >= 130, errorToString(Errors.INVALID_DATA));
        BurnInstData memory data = parseBurnInst(inst);
        require(data.meta == BURN_REQUEST_METADATA_TYPE && data.shard == 1, errorToString(Errors.INVALID_DATA)); // Check instruction type

        // Not withdrawed
        require(!isWithdrawed(data.itx), errorToString(Errors.ALREADY_USED));
        withdrawed[data.itx] = true;

        // verify beacon instruction
        verifyInst(
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

        // Send and notify
        IERC721(data.token).transferFrom(address(this), data.to, data.tokenId);

        emit Withdraw(data.token, data.to, data.tokenId);
    }

    /**
     * @dev convert enum to string value
     */
    function errorToString(Errors error) internal pure returns(string memory) {
        uint8 erroNum = uint8(error);
        uint maxlength = 10;
        bytes memory reversed = new bytes(maxlength);
        uint i = 0;
        while (erroNum != 0) {
            uint8 remainder = erroNum % 10;
            erroNum = erroNum / 10;
            reversed[i++] = byte(48 + remainder);
        }
        bytes memory s = new bytes(i + 1);
        for (uint j = 0; j <= i; j++) {
            s[j] = reversed[i - j];
        }
        return string(s);
    }
}
