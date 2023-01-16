// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./iopensea.sol";
import { OrderComponents, OfferItem, ConsiderationItem } from "./ConsiderationStructs.sol";
import { ItemType } from "./ConsiderationEnums.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155Receiver.sol";

interface IERC721 {
    /**
     * @dev Transfers `tokenId` token from `from` to `to`.
     *
     * WARNING: Usage of this method is discouraged, use {safeTransferFrom} whenever possible.
     *
     * Requirements:
     *
     * - `from` cannot be the zero address.
     * - `to` cannot be the zero address.
     * - `tokenId` token must be owned by `from`.
     * - If the caller is not `from`, it must be approved to move this token by either {approve} or {setApprovalForAll}.
     *
     * Emits a {Transfer} event.
     */
    function transferFrom(address from, address to, uint256 tokenId) external;
}

interface IERC1155 {
    /**
     * @dev Transfers `amount` tokens of token type `id` from `from` to `to`.
     *
     * Emits a {TransferSingle} event.
     *
     * Requirements:
     *
     * - `to` cannot be the zero address.
     * - If the caller is not `from`, it must be have been approved to spend ``from``'s tokens via {setApprovalForAll}.
     * - `from` must have a balance of tokens of type `id` of at least `amount`.
     * - If `to` refers to a smart contract, it must implement {IERC1155Receiver-onERC1155Received} and return the
     * acceptance magic value.
     */
    function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes calldata data) external;
}

interface Vault {
    function deposit(string calldata incognitoAddress, bytes32 txId, bytes calldata signData) payable external;
}

interface WETH {
    function approve(address spender, uint amount) external;
    function withdraw(uint wad) external;
    function deposit() payable external;
    function allowance(address owner, address spender) external view returns (uint256);
}

contract ProxyOpenSeaOffer {
    ConsiderationInterface constant public seaport = ConsiderationInterface(0x00000000006c3852cbEf3e08E8dF289169EdE581);
    bytes32 constant public domainSeparator = 0x712fdde1f147adcbb0fabb1aeb9c2d317530a46d266db095bc40c606fe94f0c2;
    Vault constant public vault = Vault(0xc157CC3077ddfa425bae12d2F3002668971A4e3d);
    WETH constant public weth = WETH(0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6);
    mapping(bytes32 => Offer) public offers;

    // offer data structure
    struct Offer {
        string otaKey;
        address signer;
        bytes signature;
        uint256 startTime;
        uint256 endTime;
        uint256 offerAmount;
        address recipient;
        uint256 claimed;
    }

    // new offer
    function offer(OrderComponents calldata order, string calldata otaKey, bytes calldata signature, address conduit, address recipient) payable external {
        // verify offerer
        require(order.offerer == address(this) && order.offer.length == 1 && order.startTime != 0, "OpenseaOffer: invalid offer");
        require(recipient != address(0), "OpenseaOffer: recipient must not 0");
        bytes32 signOfferHash = toTypedDataHash(domainSeparator, seaport.getOrderHash(order));
        require(offers[signOfferHash].startTime == 0, "OpenseaOffer: offer existed");
        address signer = recoverSigner(signOfferHash, signature);

        // verify eth amount
        uint256 offerAmount;
        for(uint8 i = 0; i < order.offer.length; i++) {
            OfferItem memory item = order.offer[i];
            require(item.startAmount == item.endAmount, "OpenseaOffer: invalid start end amount");
            offerAmount = offerAmount + item.startAmount;
        }
        require(msg.value == offerAmount, "OpenseaOffer: invalid offer amount");
        offers[signOfferHash] = Offer(otaKey, signer, signature, order.startTime, order.endTime, offerAmount, recipient, 0);
        weth.deposit{value: msg.value}();

        // approve to corresponding conduit key
        weth.approve(conduit, weth.allowance(address(this), conduit) + offerAmount);
    }

    function cancelOffer(OrderComponents calldata order,  bytes calldata offerSignature, bytes32 txId, bytes calldata regulatorSignData) external {
        bytes32 orderHash = seaport.getOrderHash(order);
        (, bool isCanceled, uint256 totalFilled, uint256 totalSize) = seaport.getOrderStatus(orderHash);
        Offer memory offerTemp = offers[toTypedDataHash(domainSeparator, orderHash)];
        require(offerTemp.startTime != 0 && offerTemp.offerAmount != 0 && !isCanceled && (totalSize > totalFilled || totalSize == 0), "OpenseaOffer: invalid offer");
        // incase the offer not expired must verify offerrer signature
        if (block.timestamp < offerTemp.endTime) {
            require(offerTemp.signer == recoverSigner(orderHash, offerSignature), "OpenseaOffer: invalid signature");
        }
        // call seaport to cancel offer
        OrderComponents[] memory orders = new OrderComponents[](1);
        orders[0] = order;
        require(seaport.cancel(orders), "OpenseaOffer: execute cancel on seaport failed");
        uint256 withdrawAmount;
        unchecked {
           withdrawAmount = totalSize == 0 ?
                offerTemp.offerAmount
                : offerTemp.offerAmount * (totalSize - totalFilled) / totalSize;
        }
        // withdraw native token
        weth.withdraw(withdrawAmount);
        // call deposit to vault contract
        vault.deposit{value: withdrawAmount}(
            offerTemp.otaKey,
            txId,
            regulatorSignData
        );
    }

    function claim(OrderComponents calldata order) external {
        bytes32 signOfferHash = toTypedDataHash(domainSeparator, seaport.getOrderHash(order));
        require(offers[signOfferHash].startTime != 0 && order.consideration.length != 0, "OpenseaOffer: invalid offer");
        ConsiderationItem memory temp = order.consideration[0];
        if (temp.itemType == ItemType.ERC721) {
            IERC721(temp.token).transferFrom(address(this), offers[signOfferHash].recipient, temp.identifierOrCriteria);
            offers[signOfferHash].claimed = 1;
        } else if (temp.itemType == ItemType.ERC1155) {
            (,, uint256 totalFilled, uint256 totalSize) = seaport.getOrderStatus(seaport.getOrderHash(order));
            if (totalFilled > 0 && totalSize > 0) {
                uint256 claimAmount = temp.endAmount * totalFilled / totalSize;
                if (claimAmount > offers[signOfferHash].claimed) {
                    uint256 transferAmount;
                    unchecked {
                        transferAmount = claimAmount - offers[signOfferHash].claimed;
                    }
                    offers[signOfferHash].claimed = claimAmount;
                    IERC1155(temp.token).safeTransferFrom(address(this), offers[signOfferHash].recipient, temp.identifierOrCriteria, transferAmount, bytes(""));
                }
            }
        } else {
            revert("OpenseaOffer: the item type not supported");
        }
    }

    // verify signature source https://eips.ethereum.org/EIPS/eip-1271
    /**
     * @notice Verifies that the signer is the owner of the signing contract.
     */
    function isValidSignature(
        bytes32 _hash,
        bytes calldata _signature
    ) external view returns (bytes4) {
        if (validateSignature(_hash, _signature)) {
            return 0x1626ba7e;
        } else {
            return 0xffffffff;
        }
    }

    function validateSignature(
        bytes32 _hash,
        bytes calldata _signature
    ) view internal returns(bool) {
        Offer memory temp = offers[_hash];
        if (temp.startTime == 0 || keccak256(temp.signature) != keccak256(_signature) || temp.endTime < block.timestamp) return false;
        return true;
    }

    /**
      * @notice Recover the signer of hash, assuming it's an EOA account
       * @dev Only for EthSign signatures
       * @param _hash       Hash of message that was signed
       * @param _signature  Signature encoded as (bytes32 r, bytes32 s, uint8 v)
       */
    function recoverSigner(
        bytes32 _hash,
        bytes memory _signature
    ) internal pure returns (address signer) {
        require(_signature.length == 65, "SignatureValidator#recoverSigner: invalid signature length");

        // Variables are not scoped in Solidity.
        uint8 v = uint8(_signature[64]);
        bytes32 s;
        bytes32 r;
        assembly {
            r := mload(add(_signature, 0x20))
            s := mload(add(_signature, 0x40))
        }

        // EIP-2 still allows signature malleability for ecrecover(). Remove this possibility and make the signature
        // unique. Appendix F in the Ethereum Yellow paper (https://ethereum.github.io/yellowpaper/paper.pdf), defines
        // the valid range for s in (281): 0 < s < secp256k1n ÷ 2 + 1, and for v in (282): v ∈ {27, 28}. Most
        // signatures from current libraries generate a unique signature with an s-value in the lower half order.
        //
        // If your library generates malleable signatures, such as s-values in the upper range, calculate a new s-value
        // with 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141 - s1 and flip v from 27 to 28 or
        // vice versa. If your library also generates signatures with 0/1 for v instead 27/28, add 27 to v to accept
        // these malleable signatures as well.
        //
        // Source OpenZeppelin
        // https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/cryptography/ECDSA.sol

        if (uint256(s) > 0x7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF5D576E7357A4501DDFE92F46681B20A0) {
            revert("SignatureValidator#recoverSigner: invalid signature 's' value");
        }

        if (v != 27 && v != 28) {
            revert("SignatureValidator#recoverSigner: invalid signature 'v' value");
        }

        // Recover ECDSA signer
        signer = ecrecover(_hash, v, r, s);

        // Prevent signer from being 0x0
        require( signer != address(0x0), "SignatureValidator#recoverSigner: INVALID_SIGNER");

        return signer;
    }

    /**
     * @dev Returns an Ethereum Signed Typed Data, created from a
     * `domainSeparator` and a `structHash`. This produces hash corresponding
     * to the one signed with the
     * https://eips.ethereum.org/EIPS/eip-712[`eth_signTypedData`]
     * JSON-RPC method as part of EIP-712.
     *
     * See {recover}.
     */
    function toTypedDataHash(bytes32 domainSeparator_, bytes32 structHash_) public pure returns (bytes32) {
        return keccak256(abi.encodePacked("\x19\x01", domainSeparator_, structHash_));
    }

    // Receive function which allows transfer eth.
    receive() external payable {}

    function onERC721Received(address, address, uint256, bytes calldata) external pure returns (bytes4) {
        return IERC721Receiver.onERC721Received.selector;
    }

    function onERC1155Received(address, address, uint256, bytes calldata) external pure returns (bytes4) {
        return IERC1155Receiver.onERC1155Received.selector;
    }

    function onERC1155BatchReceived(address, address, uint256, bytes calldata) external pure returns (bytes4) {
        return IERC1155Receiver.onERC1155BatchReceived.selector;
    }
}