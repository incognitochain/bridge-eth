// this is a contract used for local testing. It is not for live deployment.
pragma solidity ^0.6.12;
pragma experimental ABIEncoderV2;

import "./IERC20.sol";
import "./vault.sol";
import {ERC20 as OZERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "hardhat/console.sol";

contract TestingVault {
    using SafeMath for uint;
    /**
     * @dev Storage slot with the incognito proxy.
     * This is the keccak-256 hash of "eip1967.proxy.incognito." subtracted by 1
     */
    bytes32 private constant _INCOGNITO_SLOT = 0x62135fc083646fdb4e1a9d700e351b886a4a5a39da980650269edd1ade91ffd2;
    address constant public ETH_TOKEN = 0x0000000000000000000000000000000000000000;

    /**
     * @dev Storage variables for Vault
     * This section is APPEND-ONLY, in order to preserve upgradeability
     * since we use Proxy Pattern
     */
    mapping(bytes32 => bool) public withdrawed;
    mapping(bytes32 => bool) public sigDataUsed;
    // address => token => amount
    mapping(address => mapping(address => uint)) public withdrawRequests;
    mapping(address => mapping(address => bool)) public migration;
    mapping(address => uint) public totalDepositedToSCAmount;
    Withdrawable public prevVault;
    bool public notEntered = true;
    bool public isInitialized = false;

    /**
    * @dev Added in Storage Layout version : 2.0
    */
    uint16 public storageLayoutVersion;
    string public recoveryAddress;

    /**
    * @dev END Storage variables
    */

    bool public allowWithdraw = true; // for testing

    struct BurnInstData {
        uint8 meta; // type of the instruction
        uint8 shard; // ID of the Incognito shard containing the instruction, must be 1
        address token; // ETH address of the token contract (0x0 for ETH)
        address payable to; // ETH address of the receiver of the token
        uint amount; // burned amount (on Incognito)
        bytes32 itx; // Incognito's burning tx
    }

    enum Prefix {
        EXECUTE_SIGNATURE,
        REQUEST_WITHDRAW_SIGNATURE
    }

    // error code
    enum Errors {
        EMPTY,
        NO_REENTRANCE,
        MAX_UINT_REACHED,
        VALUE_OVER_FLOW,
        INTERNAL_TX_ERROR,
        ALREADY_USED,
        INVALID_DATA,
        TOKEN_NOT_ENOUGH,
        WITHDRAW_REQUEST_TOKEN_NOT_ENOUGH,
        INVALID_RETURN_DATA,
        NOT_EQUAL,
        NULL_VALUE,
        ONLY_PREVAULT,
        PREVAULT_NOT_PAUSED,
        SAFEMATH_EXCEPTION,
        ALREADY_INITIALIZED,
        INVALID_SIGNATURE,
        ALREADY_UPGRADED
    }

    event Deposit(address token, string incognitoAddress, uint amount);
    event Withdraw(address token, address to, uint amount);
    event UpdateTokenTotal(address[] assets, uint[] amounts);
    event UpdateIncognitoProxy(address newIncognitoProxy);

    /**
     * modifier for contract version
     */
     modifier onlyPreVault(){
        require(address(prevVault) != address(0x0) && msg.sender == address(prevVault), errorToString(Errors.ONLY_PREVAULT));
        _;
     }

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

    function setAllowWithdraw(bool b) public {
        allowWithdraw = b;
    }

    /**
     * @dev Creates new Vault to hold assets for Incognito Chain
     * @param _prevVault: previous version of the Vault to refer back if necessary
     * After migrating all assets to a new Vault, we still need to refer
     * back to previous Vault to make sure old withdrawals aren't being reused
     */
    function initialize(address _prevVault) external {
        require(!isInitialized, errorToString(Errors.ALREADY_INITIALIZED));
        prevVault = Withdrawable(_prevVault);
        isInitialized = true;
        notEntered = true;
    }

    function upgradeVaultStorageLayout() external {
        // storageLayoutVersion is a new variable introduced in this storage layout version, then set to 2 to match the storage layout version itself
        require(storageLayoutVersion == 2, errorToString(Errors.ALREADY_UPGRADED));
        // make sure the version increase can only happen once
        storageLayoutVersion = 3;
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
     * @dev Makes a ETH deposit to the vault to mint pETH over at Incognito Chain
     * @notice This only works when the contract is not Paused
     * @notice The maximum amount to deposit is capped since Incognito balance is stored as uint64
     * @param incognitoAddress: Incognito Address to receive pETH
     */
    function deposit(string calldata incognitoAddress) external payable nonReentrant {
        require(address(this).balance <= 10 ** 27, errorToString(Errors.MAX_UINT_REACHED));
        emit Deposit(ETH_TOKEN, incognitoAddress, msg.value);
    }

    /**
     * @dev Makes a ERC20 deposit to the vault to mint pERC20 over at Incognito Chain
     * @notice This only works when the contract is not Paused
     * @notice The maximum amount to deposit is capped since Incognito balance is stored as uint64
     * @notice Before calling this function, enough ERC20 must be allowed to
     * tranfer from msg.sender to this contract
     * @param token: address of the ERC20 token
     * @param amount: to deposit to the vault and mint on Incognito Chain
     * @param incognitoAddress: Incognito Address to receive pERC20
     */
    function depositERC20(address token, uint amount, string calldata incognitoAddress) external nonReentrant {
        IERC20 erc20Interface = IERC20(token);
        uint8 decimals = getDecimals(address(token));
        uint tokenBalance = erc20Interface.balanceOf(address(this));
        uint beforeTransfer = tokenBalance;
        uint emitAmount = amount;
        if (decimals > 9) {
            emitAmount = emitAmount / (10 ** (uint(decimals) - 9));
            tokenBalance = tokenBalance / (10 ** (uint(decimals) - 9));
        }
        require(emitAmount <= 10 ** 18 && tokenBalance <= 10 ** 18 && emitAmount.safeAdd(tokenBalance) <= 10 ** 18, errorToString(Errors.VALUE_OVER_FLOW));
        erc20Interface.transferFrom(msg.sender, address(this), amount);
        require(checkSuccess(), errorToString(Errors.INTERNAL_TX_ERROR));
        require(balanceOf(token).safeSub(beforeTransfer) == amount, errorToString(Errors.NOT_EQUAL));

        emit Deposit(token, incognitoAddress, emitAmount);
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
        if (withdrawed[hash]) {
            return true;
        } else if (address(prevVault) == address(0)) {
            return false;
        }
        return prevVault.isWithdrawed(hash);
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
     * @dev Withdraws pETH/pIERC20 by providing a burn proof over at Incognito Chain
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
        require(data.meta == 241 && data.shard == 1, errorToString(Errors.INVALID_DATA)); // Check instruction type

        // Not withdrawed
        require(!isWithdrawed(data.itx), errorToString(Errors.ALREADY_USED));
        withdrawed[data.itx] = true;

        // Check if balance is enough
        if (data.token == ETH_TOKEN) {
            require(address(this).balance >= data.amount.safeAdd(totalDepositedToSCAmount[data.token]), errorToString(Errors.TOKEN_NOT_ENOUGH));
        } else {
            uint8 decimals = getDecimals(data.token);
            if (decimals > 9) {
                data.amount = data.amount.safeMul(10 ** (uint(decimals) - 9));
            }
            require(IERC20(data.token).balanceOf(address(this)) >= data.amount.safeAdd(totalDepositedToSCAmount[data.token]), errorToString(Errors.TOKEN_NOT_ENOUGH));
        }

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
        if (allowWithdraw) {
            if (data.token == ETH_TOKEN) {
              (bool success, ) =  data.to.call{value: data.amount}("");
              require(success, errorToString(Errors.INTERNAL_TX_ERROR));
            } else {
                IERC20(data.token).transfer(data.to, data.amount);
                require(checkSuccess(), errorToString(Errors.INTERNAL_TX_ERROR));
            }
            emit Withdraw(data.token, data.to, data.amount);
        } else {
            // simulate the effect of an attack; here user will lose pToken without receiving anything on Ethereum
            // we use this to test pausing & upgrading features
            console.log("cannot withdraw");
        }
    }

    /**
     * @dev Burnt Proof is submited to store burnt amount of p-token/p-ETH and receiver's address
     * Receiver then can call withdrawRequest to withdraw these token to he/she incognito address.
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
    function submitBurnProof(
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
        require(data.meta == 243 && data.shard == 1, errorToString(Errors.INVALID_DATA)); // Check instruction type

        // Not withdrawed
        require(!isWithdrawed(data.itx), errorToString(Errors.ALREADY_USED));
        withdrawed[data.itx] = true;

        // Check if balance is enough
        if (data.token == ETH_TOKEN) {
            require(address(this).balance >= data.amount.safeAdd(totalDepositedToSCAmount[data.token]), errorToString(Errors.TOKEN_NOT_ENOUGH));
        } else {
            uint8 decimals = getDecimals(data.token);
            if (decimals > 9) {
                data.amount = data.amount.safeMul(10 ** (uint(decimals) - 9));
            }
            require(IERC20(data.token).balanceOf(address(this)) >= data.amount.safeAdd(totalDepositedToSCAmount[data.token]), errorToString(Errors.TOKEN_NOT_ENOUGH));
        }

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

        withdrawRequests[data.to][data.token] = withdrawRequests[data.to][data.token].safeAdd(data.amount);
        totalDepositedToSCAmount[data.token] = totalDepositedToSCAmount[data.token].safeAdd(data.amount);
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
        } else if (address(prevVault) == address(0)) {
            return false;
        }
        return prevVault.isSigDataUsed(hash);
    }

    struct PreSignData {
        Prefix prefix;
        address token;
        bytes timestamp;
        uint amount;
    }

    function newPreSignData(Prefix prefix, address token, bytes calldata timestamp, uint amount) pure internal returns (PreSignData memory) {
        PreSignData memory psd = PreSignData(prefix, token, timestamp, amount);
        return psd;
    }

    /**
     * @dev User requests withdraw token contains in withdrawRequests.
     * Deposit event will be emitted to let incognito recognize and mint new p-tokens for the user.
     * @param incognitoAddress: incognito's address that will receive minted p-tokens.
     * @param token: ethereum's token address (eg., ETH, DAI, ...)
     * @param amount: amount of the token in ethereum's denomination
     * @param signData: signature of an unique data that is signed by an account which is generated from user's incognito privkey
     * @param timestamp: unique data generated from client (timestamp for example)
     */
    function requestWithdraw(
        string calldata incognitoAddress,
        address token,
        uint amount,
        bytes calldata signData,
        bytes calldata timestamp
    ) external nonReentrant {
        // verify owner signs data
        address verifier = verifySignData(abi.encode(newPreSignData(Prefix.REQUEST_WITHDRAW_SIGNATURE, token, timestamp, amount), incognitoAddress), signData);

        // migrate from preVault
        migrateBalance(verifier, token);

        require(withdrawRequests[verifier][token] >= amount, errorToString(Errors.WITHDRAW_REQUEST_TOKEN_NOT_ENOUGH));
        withdrawRequests[verifier][token] = withdrawRequests[verifier][token].safeSub(amount);
        totalDepositedToSCAmount[token] = totalDepositedToSCAmount[token].safeSub(amount);

        // convert denomination from ethereum's to incognito's (pcoin)
        uint emitAmount = amount;
        if (token != ETH_TOKEN) {
            uint8 decimals = getDecimals(token);
            if (decimals > 9) {
                emitAmount = amount / (10 ** (uint(decimals) - 9));
            }
        }

        emit Deposit(token, incognitoAddress, emitAmount);
    }

    /**
     * @dev execute is a general function that plays a role as proxy to interact to other smart contracts.
     * @param token: ethereum's token address (eg., ETH, DAI, ...)
     * @param amount: amount of the token in ethereum's denomination
     * @param recipientToken: received token address.
     * @param exchangeAddress: address of targeting smart contract that actually executes the desired logics like trade, invest, borrow and so on.
     * @param callData: encoded with signature and params of function from targeting smart contract.
     * @param timestamp: unique data generated from client (timestamp for example)
     * @param signData: signature of an unique data that is signed by an account which is generated from user's incognito privkey
     */
    function execute(
        address token,
        uint amount,
        address recipientToken,
        address exchangeAddress,
        bytes calldata callData,
        bytes calldata timestamp,
        bytes calldata signData
    ) external payable nonReentrant {
        //verify ower signs data from input
        address verifier = verifySignData(abi.encode(newPreSignData(Prefix.EXECUTE_SIGNATURE, token, timestamp, amount), recipientToken, exchangeAddress, callData), signData);

        // migrate from preVault
        migrateBalance(verifier, token);
        require(withdrawRequests[verifier][token] >= amount, errorToString(Errors.WITHDRAW_REQUEST_TOKEN_NOT_ENOUGH));

        // update balance of verifier
        totalDepositedToSCAmount[token] = totalDepositedToSCAmount[token].safeSub(amount);
        withdrawRequests[verifier][token] = withdrawRequests[verifier][token].safeSub(amount);

        // define number of eth spent for forwarder.
        uint ethAmount = msg.value;
        if (token == ETH_TOKEN) {
            ethAmount = ethAmount.safeAdd(amount);
        } else {
            // transfer token to exchangeAddress.
            require(IERC20(token).balanceOf(address(this)) >= amount, errorToString(Errors.TOKEN_NOT_ENOUGH));
            IERC20(token).transfer(exchangeAddress, amount);
            require(checkSuccess(), errorToString(Errors.INTERNAL_TX_ERROR));
        }
        uint returnedAmount = callExtFunc(recipientToken, ethAmount, callData, exchangeAddress);

        // update withdrawRequests
        withdrawRequests[verifier][recipientToken] = withdrawRequests[verifier][recipientToken].safeAdd(returnedAmount);
        totalDepositedToSCAmount[recipientToken] = totalDepositedToSCAmount[recipientToken].safeAdd(returnedAmount);
    }

    /**
     * @dev single trade
     */
    function callExtFunc(address recipientToken, uint ethAmount, bytes memory callData, address exchangeAddress) internal returns (uint) {
         // get balance of recipient token before trade to compare after trade.
        uint balanceBeforeTrade = balanceOf(recipientToken);
        if (recipientToken == ETH_TOKEN) {
            balanceBeforeTrade = balanceBeforeTrade.safeSub(msg.value);
        }
        require(address(this).balance >= ethAmount, errorToString(Errors.TOKEN_NOT_ENOUGH));
        (bool success, bytes memory result) = exchangeAddress.call{value: ethAmount}(callData);
        require(success, errorToString(Errors.INTERNAL_TX_ERROR));

        (address returnedTokenAddress, uint returnedAmount) = abi.decode(result, (address, uint));
        require(returnedTokenAddress == recipientToken && balanceOf(recipientToken).safeSub(balanceBeforeTrade) == returnedAmount, errorToString(Errors.INVALID_RETURN_DATA));

        return returnedAmount;
    }

    /**
     * @dev verify sign data
     */
     function verifySignData(bytes memory data, bytes memory signData) internal returns(address){
        bytes32 hash = keccak256(data);
        require(!isSigDataUsed(hash), errorToString(Errors.ALREADY_USED));
        address verifier = sigToAddress(signData, hash);
        // reject when verifier equals zero
        require(verifier != address(0x0), errorToString(Errors.INVALID_SIGNATURE));
        // mark data hash of sig as used
        sigDataUsed[hash] = true;

        return verifier;
     }

    /**
      * @dev migrate balance from previous vault
      * Note: uncomment for next version
      */
    function migrateBalance(address owner, address token) internal {
        if (address(prevVault) != address(0x0) && !migration[owner][token]) {
            withdrawRequests[owner][token] = withdrawRequests[owner][token].safeAdd(prevVault.getDepositedBalance(token, owner));
            migration[owner][token] = true;
       }
    }

    /**
     * @dev Get the amount of specific coin for specific wallet
     */
    function getDepositedBalance(
        address token,
        address owner
    ) public view returns (uint) {
        if (address(prevVault) != address(0x0) && !migration[owner][token]) {
            return withdrawRequests[owner][token].safeAdd(prevVault.getDepositedBalance(token, owner));
        }
        return withdrawRequests[owner][token];
    }

    /**
     * @dev Move total number of assets to newVault
     * @notice This only works when the preVault is Paused
     * @notice This can only be called by preVault
     * @param assets: address of the ERC20 tokens to move, 0x0 for ETH
     * @param amounts: total number of the ERC20 tokens to move, 0x0 for ETH
     */
    function updateAssets(address[] calldata assets, uint[] calldata amounts) external onlyPreVault returns(bool) {
        require(assets.length == amounts.length,  errorToString(Errors.NOT_EQUAL));
        require(Withdrawable(prevVault).paused(), errorToString(Errors.PREVAULT_NOT_PAUSED));
        for (uint i = 0; i < assets.length; i++) {
            totalDepositedToSCAmount[assets[i]] = totalDepositedToSCAmount[assets[i]].safeAdd(amounts[i]);
        }
        emit UpdateTokenTotal(assets, amounts);

        return true;
    }

    /**
     * @dev Payable receive function to receive Ether from oldVault when migrating
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

    /**
     * @dev Get the decimals of an ERC20 token, return 0 if it isn't defined
     * We check the returndatasize to covert both cases that the token has
     * and doesn't have the function decimals()
     */
    function getDecimals(address token) public view returns (uint8) {
        IERC20 erc20 = IERC20(token);
        return uint8(erc20.decimals());
    }

    /**
     * @dev Get the amount of coin deposited to this smartcontract
     */
    function balanceOf(address token) public view returns (uint) {
        if (token == ETH_TOKEN) {
            return address(this).balance;
        }
        return IERC20(token).balanceOf(address(this));
    }
}

/**
 * @dev Vault contract code from previous version : v1.1
 */
contract PreviousVersionVault {
    using SafeMath for uint;
    /**
     * @dev Storage slot with the incognito proxy.
     * This is the keccak-256 hash of "eip1967.proxy.incognito." subtracted by 1
     */
    bytes32 private constant _INCOGNITO_SLOT = 0x62135fc083646fdb4e1a9d700e351b886a4a5a39da980650269edd1ade91ffd2;
    address constant public ETH_TOKEN = 0x0000000000000000000000000000000000000000;

    /**
     * @dev Storage variables for Vault
     * This section is APPEND-ONLY, in order to preserve upgradeability
     * since we use Proxy Pattern
     */
    mapping(bytes32 => bool) public withdrawed;
    mapping(bytes32 => bool) public sigDataUsed;
    // address => token => amount
    mapping(address => mapping(address => uint)) public withdrawRequests;
    mapping(address => mapping(address => bool)) public migration;
    mapping(address => uint) public totalDepositedToSCAmount;
    Withdrawable public prevVault;
    bool public notEntered = true;
    bool public isInitialized = false;

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

    enum Prefix {
        EXECUTE_SIGNATURE,
        REQUEST_WITHDRAW_SIGNATURE
    }

    // error code
    enum Errors {
        EMPTY,
        NO_REENTRANCE,
        MAX_UINT_REACHED,
        VALUE_OVER_FLOW,
        INTERNAL_TX_ERROR,
        ALREADY_USED,
        INVALID_DATA,
        TOKEN_NOT_ENOUGH,
        WITHDRAW_REQUEST_TOKEN_NOT_ENOUGH,
        INVALID_RETURN_DATA,
        NOT_EQUAL,
        NULL_VALUE,
        ONLY_PREVAULT,
        PREVAULT_NOT_PAUSED,
        SAFEMATH_EXCEPTION,
        ALREADY_INITIALIZED,
        INVALID_SIGNATURE
    }

    event Deposit(address token, string incognitoAddress, uint amount);
    event Withdraw(address token, address to, uint amount);
    event UpdateTokenTotal(address[] assets, uint[] amounts);
    event UpdateIncognitoProxy(address newIncognitoProxy);

    /**
     * modifier for contract version
     */
     modifier onlyPreVault(){
        require(address(prevVault) != address(0x0) && msg.sender == address(prevVault), errorToString(Errors.ONLY_PREVAULT));
        _;
     }

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
     * @param _prevVault: previous version of the Vault to refer back if necessary
     * After migrating all assets to a new Vault, we still need to refer
     * back to previous Vault to make sure old withdrawals aren't being reused
     */
    function initialize(address _prevVault) external {
        require(!isInitialized, errorToString(Errors.ALREADY_INITIALIZED));
        prevVault = Withdrawable(_prevVault);
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
     * @dev Makes a ETH deposit to the vault to mint pETH over at Incognito Chain
     * @notice This only works when the contract is not Paused
     * @notice The maximum amount to deposit is capped since Incognito balance is stored as uint64
     * @param incognitoAddress: Incognito Address to receive pETH
     */
    function deposit(string calldata incognitoAddress) external payable nonReentrant {
        require(address(this).balance <= 10 ** 27, errorToString(Errors.MAX_UINT_REACHED));
        emit Deposit(ETH_TOKEN, incognitoAddress, msg.value);
    }

    /**
     * @dev Makes a ERC20 deposit to the vault to mint pERC20 over at Incognito Chain
     * @notice This only works when the contract is not Paused
     * @notice The maximum amount to deposit is capped since Incognito balance is stored as uint64
     * @notice Before calling this function, enough ERC20 must be allowed to
     * tranfer from msg.sender to this contract
     * @param token: address of the ERC20 token
     * @param amount: to deposit to the vault and mint on Incognito Chain
     * @param incognitoAddress: Incognito Address to receive pERC20
     */
    function depositERC20(address token, uint amount, string calldata incognitoAddress) external nonReentrant {
        IERC20 erc20Interface = IERC20(token);
        uint8 decimals = getDecimals(address(token));
        uint tokenBalance = erc20Interface.balanceOf(address(this));
        uint beforeTransfer = tokenBalance;
        uint emitAmount = amount;
        if (decimals > 9) {
            emitAmount = emitAmount / (10 ** (uint(decimals) - 9));
            tokenBalance = tokenBalance / (10 ** (uint(decimals) - 9));
        }
        require(emitAmount <= 10 ** 18 && tokenBalance <= 10 ** 18 && emitAmount.safeAdd(tokenBalance) <= 10 ** 18, errorToString(Errors.VALUE_OVER_FLOW));
        erc20Interface.transferFrom(msg.sender, address(this), amount);
        require(checkSuccess(), errorToString(Errors.INTERNAL_TX_ERROR));
        require(balanceOf(token).safeSub(beforeTransfer) == amount, errorToString(Errors.NOT_EQUAL));

        emit Deposit(token, incognitoAddress, emitAmount);
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
        if (withdrawed[hash]) {
            return true;
        } else if (address(prevVault) == address(0)) {
            return false;
        }
        return prevVault.isWithdrawed(hash);
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
     * @dev Withdraws pETH/pIERC20 by providing a burn proof over at Incognito Chain
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
        require(data.meta == 241 && data.shard == 1, errorToString(Errors.INVALID_DATA)); // Check instruction type

        // Not withdrawed
        require(!isWithdrawed(data.itx), errorToString(Errors.ALREADY_USED));
        withdrawed[data.itx] = true;

        // Check if balance is enough
        if (data.token == ETH_TOKEN) {
            require(address(this).balance >= data.amount.safeAdd(totalDepositedToSCAmount[data.token]), errorToString(Errors.TOKEN_NOT_ENOUGH));
        } else {
            uint8 decimals = getDecimals(data.token);
            if (decimals > 9) {
                data.amount = data.amount.safeMul(10 ** (uint(decimals) - 9));
            }
            require(IERC20(data.token).balanceOf(address(this)) >= data.amount.safeAdd(totalDepositedToSCAmount[data.token]), errorToString(Errors.TOKEN_NOT_ENOUGH));
        }

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
        if (data.token == ETH_TOKEN) {
          (bool success, ) =  data.to.call{value: data.amount}("");
          require(success, errorToString(Errors.INTERNAL_TX_ERROR));
        } else {
            IERC20(data.token).transfer(data.to, data.amount);
            require(checkSuccess(), errorToString(Errors.INTERNAL_TX_ERROR));
        }
        emit Withdraw(data.token, data.to, data.amount);
    }

    /**
     * @dev Burnt Proof is submited to store burnt amount of p-token/p-ETH and receiver's address
     * Receiver then can call withdrawRequest to withdraw these token to he/she incognito address.
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
    function submitBurnProof(
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
        require(data.meta == 243 && data.shard == 1, errorToString(Errors.INVALID_DATA)); // Check instruction type

        // Not withdrawed
        require(!isWithdrawed(data.itx), errorToString(Errors.ALREADY_USED));
        withdrawed[data.itx] = true;

        // Check if balance is enough
        if (data.token == ETH_TOKEN) {
            require(address(this).balance >= data.amount.safeAdd(totalDepositedToSCAmount[data.token]), errorToString(Errors.TOKEN_NOT_ENOUGH));
        } else {
            uint8 decimals = getDecimals(data.token);
            if (decimals > 9) {
                data.amount = data.amount.safeMul(10 ** (uint(decimals) - 9));
            }
            require(IERC20(data.token).balanceOf(address(this)) >= data.amount.safeAdd(totalDepositedToSCAmount[data.token]), errorToString(Errors.TOKEN_NOT_ENOUGH));
        }

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

        withdrawRequests[data.to][data.token] = withdrawRequests[data.to][data.token].safeAdd(data.amount);
        totalDepositedToSCAmount[data.token] = totalDepositedToSCAmount[data.token].safeAdd(data.amount);
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
        } else if (address(prevVault) == address(0)) {
            return false;
        }
        return prevVault.isSigDataUsed(hash);
    }

    struct PreSignData {
        Prefix prefix;
        address token;
        bytes timestamp;
        uint amount;
    }

    function newPreSignData(Prefix prefix, address token, bytes calldata timestamp, uint amount) pure internal returns (PreSignData memory) {
        PreSignData memory psd = PreSignData(prefix, token, timestamp, amount);
        return psd;
    }

    /**
     * @dev User requests withdraw token contains in withdrawRequests.
     * Deposit event will be emitted to let incognito recognize and mint new p-tokens for the user.
     * @param incognitoAddress: incognito's address that will receive minted p-tokens.
     * @param token: ethereum's token address (eg., ETH, DAI, ...)
     * @param amount: amount of the token in ethereum's denomination
     * @param signData: signature of an unique data that is signed by an account which is generated from user's incognito privkey
     * @param timestamp: unique data generated from client (timestamp for example)
     */
    function requestWithdraw(
        string calldata incognitoAddress,
        address token,
        uint amount,
        bytes calldata signData,
        bytes calldata timestamp
    ) external nonReentrant {
        // verify owner signs data
        address verifier = verifySignData(abi.encode(newPreSignData(Prefix.REQUEST_WITHDRAW_SIGNATURE, token, timestamp, amount), incognitoAddress), signData);

        // migrate from preVault
        migrateBalance(verifier, token);

        require(withdrawRequests[verifier][token] >= amount, errorToString(Errors.WITHDRAW_REQUEST_TOKEN_NOT_ENOUGH));
        withdrawRequests[verifier][token] = withdrawRequests[verifier][token].safeSub(amount);
        totalDepositedToSCAmount[token] = totalDepositedToSCAmount[token].safeSub(amount);

        // convert denomination from ethereum's to incognito's (pcoin)
        uint emitAmount = amount;
        if (token != ETH_TOKEN) {
            uint8 decimals = getDecimals(token);
            if (decimals > 9) {
                emitAmount = amount / (10 ** (uint(decimals) - 9));
            }
        }

        emit Deposit(token, incognitoAddress, emitAmount);
    }

    /**
     * @dev execute is a general function that plays a role as proxy to interact to other smart contracts.
     * @param token: ethereum's token address (eg., ETH, DAI, ...)
     * @param amount: amount of the token in ethereum's denomination
     * @param recipientToken: received token address.
     * @param exchangeAddress: address of targeting smart contract that actually executes the desired logics like trade, invest, borrow and so on.
     * @param callData: encoded with signature and params of function from targeting smart contract.
     * @param timestamp: unique data generated from client (timestamp for example)
     * @param signData: signature of an unique data that is signed by an account which is generated from user's incognito privkey
     */
    function execute(
        address token,
        uint amount,
        address recipientToken,
        address exchangeAddress,
        bytes calldata callData,
        bytes calldata timestamp,
        bytes calldata signData
    ) external payable nonReentrant {
        //verify ower signs data from input
        address verifier = verifySignData(abi.encode(newPreSignData(Prefix.EXECUTE_SIGNATURE, token, timestamp, amount), recipientToken, exchangeAddress, callData), signData);

        // migrate from preVault
        migrateBalance(verifier, token);
        require(withdrawRequests[verifier][token] >= amount, errorToString(Errors.WITHDRAW_REQUEST_TOKEN_NOT_ENOUGH));

        // update balance of verifier
        totalDepositedToSCAmount[token] = totalDepositedToSCAmount[token].safeSub(amount);
        withdrawRequests[verifier][token] = withdrawRequests[verifier][token].safeSub(amount);

        // define number of eth spent for forwarder.
        uint ethAmount = msg.value;
        if (token == ETH_TOKEN) {
            ethAmount = ethAmount.safeAdd(amount);
        } else {
            // transfer token to exchangeAddress.
            require(IERC20(token).balanceOf(address(this)) >= amount, errorToString(Errors.TOKEN_NOT_ENOUGH));
            IERC20(token).transfer(exchangeAddress, amount);
            require(checkSuccess(), errorToString(Errors.INTERNAL_TX_ERROR));
        }
        uint returnedAmount = callExtFunc(recipientToken, ethAmount, callData, exchangeAddress);

        // update withdrawRequests
        withdrawRequests[verifier][recipientToken] = withdrawRequests[verifier][recipientToken].safeAdd(returnedAmount);
        totalDepositedToSCAmount[recipientToken] = totalDepositedToSCAmount[recipientToken].safeAdd(returnedAmount);
    }

    /**
     * @dev single trade
     */
    function callExtFunc(address recipientToken, uint ethAmount, bytes memory callData, address exchangeAddress) internal returns (uint) {
         // get balance of recipient token before trade to compare after trade.
        uint balanceBeforeTrade = balanceOf(recipientToken);
        if (recipientToken == ETH_TOKEN) {
            balanceBeforeTrade = balanceBeforeTrade.safeSub(msg.value);
        }
        require(address(this).balance >= ethAmount, errorToString(Errors.TOKEN_NOT_ENOUGH));
        (bool success, bytes memory result) = exchangeAddress.call{value: ethAmount}(callData);
        require(success, errorToString(Errors.INTERNAL_TX_ERROR));

        (address returnedTokenAddress, uint returnedAmount) = abi.decode(result, (address, uint));
        require(returnedTokenAddress == recipientToken && balanceOf(recipientToken).safeSub(balanceBeforeTrade) == returnedAmount, errorToString(Errors.INVALID_RETURN_DATA));

        return returnedAmount;
    }

    /**
     * @dev verify sign data
     */
     function verifySignData(bytes memory data, bytes memory signData) internal returns(address){
        bytes32 hash = keccak256(data);
        require(!isSigDataUsed(hash), errorToString(Errors.ALREADY_USED));
        address verifier = sigToAddress(signData, hash);
        // reject when verifier equals zero
        require(verifier != address(0x0), errorToString(Errors.INVALID_SIGNATURE));
        // mark data hash of sig as used
        sigDataUsed[hash] = true;

        return verifier;
     }

    /**
      * @dev migrate balance from previous vault
      * Note: uncomment for next version
      */
    function migrateBalance(address owner, address token) internal {
        if (address(prevVault) != address(0x0) && !migration[owner][token]) {
            withdrawRequests[owner][token] = withdrawRequests[owner][token].safeAdd(prevVault.getDepositedBalance(token, owner));
            migration[owner][token] = true;
       }
    }

    /**
     * @dev Get the amount of specific coin for specific wallet
     */
    function getDepositedBalance(
        address token,
        address owner
    ) public view returns (uint) {
        if (address(prevVault) != address(0x0) && !migration[owner][token]) {
            return withdrawRequests[owner][token].safeAdd(prevVault.getDepositedBalance(token, owner));
        }
        return withdrawRequests[owner][token];
    }

    /**
     * @dev Move total number of assets to newVault
     * @notice This only works when the preVault is Paused
     * @notice This can only be called by preVault
     * @param assets: address of the ERC20 tokens to move, 0x0 for ETH
     * @param amounts: total number of the ERC20 tokens to move, 0x0 for ETH
     */
    function updateAssets(address[] calldata assets, uint[] calldata amounts) external onlyPreVault returns(bool) {
        require(assets.length == amounts.length,  errorToString(Errors.NOT_EQUAL));
        require(Withdrawable(prevVault).paused(), errorToString(Errors.PREVAULT_NOT_PAUSED));
        for (uint i = 0; i < assets.length; i++) {
            totalDepositedToSCAmount[assets[i]] = totalDepositedToSCAmount[assets[i]].safeAdd(amounts[i]);
        }
        emit UpdateTokenTotal(assets, amounts);

        return true;
    }

    /**
     * @dev Payable receive function to receive Ether from oldVault when migrating
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

    /**
     * @dev Get the decimals of an ERC20 token, return 0 if it isn't defined
     * We check the returndatasize to covert both cases that the token has
     * and doesn't have the function decimals()
     */
    function getDecimals(address token) public view returns (uint8) {
        IERC20 erc20 = IERC20(token);
        return uint8(erc20.decimals());
    }

    /**
     * @dev Get the amount of coin deposited to this smartcontract
     */
    function balanceOf(address token) public view returns (uint) {
        if (token == ETH_TOKEN) {
            return address(this).balance;
        }
        return IERC20(token).balanceOf(address(this));
    }
}

contract Token1 is OZERC20 {
    constructor () OZERC20("Test Token 1", "TT1") public {
        _setupDecimals(12);
    }
    function mint(address account, uint256 amount) public payable {
        _mint(account, amount);
    }
}
contract Token2 is OZERC20 {
    constructor () OZERC20("Test Token 2", "TT2") public {
        _setupDecimals(8);
    }
    function mint(address account, uint256 amount) public payable {
        _mint(account, amount);
    }
}
contract Token3 is OZERC20 {
    constructor () OZERC20("Test Token 3", "TT3") public {
        _setupDecimals(16);
    }
    function mint(address account, uint256 amount) public payable {
        _mint(account, amount);
    }
}

contract TestingExchange {
    // swap 1-1 for everything
    constructor() public {}
    receive() external payable {}
    // kyber interface
    // trade() will not be called
    function trade(IERC20 src, uint srcAmount, IERC20 dest, address destAddress, uint maxDestAmount, uint minConversionRate, address walletId) external payable returns(uint) {
        dest.transfer(destAddress, srcAmount);
        return srcAmount;
    }
    function swapTokenToToken(IERC20 src, uint srcAmount, IERC20 dest, uint minConversionRate) external returns(uint) {
        src.transferFrom(msg.sender, address(this), srcAmount);
        dest.transfer(msg.sender, srcAmount);
        return srcAmount;
    }
    function swapEtherToToken(IERC20 token, uint minConversionRate) external payable returns(uint) {
        uint val = msg.value;
        // msg.sender.transfer(val);
        token.transfer(msg.sender, val);
        return val;
    }
    function swapTokenToEther(IERC20 token, uint srcAmount, uint minConversionRate) external returns(uint) {
        token.transferFrom(msg.sender, address(this), srcAmount);
        msg.sender.call{value: srcAmount}("");
        return srcAmount;
    }
    function getExpectedRate(IERC20 src, IERC20 dest, uint srcQty) external view returns(uint expectedRate, uint slippageRate) {
        return (1000000000000000000, 1000000000000000000);
    }

    // uniswap interface
    function factory() external pure returns (address) { return address(0x0); }
    function WETH() external pure returns (address) { return address(0x0); }
    function WETH9() external pure returns (address) { return address(0x0); }

    // path length of 2 only
    function swapExactTokensForTokens(uint amountIn, uint amountOutMin, address[] calldata path, address to, uint deadline) external returns (uint[] memory amounts) {
        uint len = path.length;
        uint[] memory result = new uint[](len);
        result[0] = amountIn;
        result[1] = amountOutMin;
        IERC20(path[0]).transferFrom(msg.sender, address(this), amountIn);
        IERC20(path[len-1]).transfer(to, amountOutMin);
        return result;
    }
    function swapExactETHForTokens(uint amountOutMin, address[] calldata path, address to, uint deadline) external payable returns (uint[] memory amounts) {
        uint len = path.length;
        uint[] memory result = new uint[](len);
        result[0] = msg.value;
        result[1] = amountOutMin;
        IERC20(path[len-1]).transfer(to, amountOutMin);
        return result;
    }
    function swapExactTokensForETH(uint amountIn, uint amountOutMin, address[] calldata path, address to, uint deadline) external returns (uint[] memory amounts) {
        uint len = path.length;
        uint[] memory result = new uint[](len);
        result[0] = amountIn;
        result[1] = amountOutMin;

        IERC20(path[0]).transferFrom(msg.sender, address(this), amountIn);
        address payable recv = payable(to);
        uint tempNum = amountOutMin;
        recv.call{value: amountOutMin}("");
        return result;
    }
    function getAmountsOut(uint amountIn, address[] calldata path) external view returns (uint[] memory amounts) {
        uint[] memory result = new uint[](2);
        result[0] = amountIn;
        result[1] = amountIn;
        return result;
    }
}