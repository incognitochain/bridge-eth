// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./iopensea.sol";
import { OrderComponents, OfferItem } from "./ConsiderationStructs.sol";

interface Vault {
    function deposit(string calldata incognitoAddress, bytes32 txId, bytes calldata signData) payable external;
}

interface WETH {
    function approve(address spender, uint amount) external;
    function withdraw(uint wad) external;
    function deposit() payable external;
}

contract ProxyOpenSeaOffer {
    ConsiderationInterface constant public seaport = ConsiderationInterface(0x00000000006c3852cbEf3e08E8dF289169EdE581);
    address constant public executor= address(0);
    Vault constant public vault = Vault(0xc157CC3077ddfa425bae12d2F3002668971A4e3d);
    WETH constant public weth = WETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2);
    mapping(bytes32 => Offer) public offers;

    enum Status {
        empty,
        offering,
        expired
    }

    // offer data structure
    struct Offer {
        bytes otaKey;
        bool isCanceled;
        address signer;
        uint256 startTime;
        uint256 endTime;
        uint256 offerAmount;
    }

    constructor() {
        weth.approve(address(seaport), 2**256 - 1);
    }

    modifier onlyProxyOffer() {
        require(msg.sender == executor, "ProxyOfferStorage: not authorised");
        _;
    }

    // new offer
    function newOffer(OrderComponents calldata order, bytes calldata otaKey, address signer) payable external {
        // verify offerer
        require(order.offerer == address(this), "OpenseaOffer: invalid offerer");
        // verify eth amount
        uint256 offerAmount;
        for(uint8 i = 0; i < order.offer.length; i++) {
            OfferItem memory item = order.offer[i];
            require(item.startAmount == item.endAmount, "OpenseaOffer: invalid start end amount");
            offerAmount = offerAmount + item.startAmount;
        }
        require(msg.value == offerAmount, "OpenseaOffer: invalid offer amount");
        bytes32 orderHash = seaport.getOrderHash(order);
        offers[orderHash] = Offer(otaKey, false, signer, order.startTime, order.endTime, offerAmount);
        weth.deposit{value: msg.value}();
    }

    function cancelOffer(OrderComponents calldata order,  bytes calldata offerSignature, bytes32 txId, bytes calldata regulatorSignData) external {
        bytes32 orderHash = seaport.getOrderHash(order);
        require(!offers[orderHash].isCanceled, "OpenseaOffer: offer cancelled");
        // verify signature
        require(offers[orderHash].signer == recoverSigner(orderHash, offerSignature), "OpenseaOffer: invalid signature");
        // call seaport cancel offer
        OrderComponents[] memory orders = new OrderComponents[](1);
        orders[0] = order;
        require(seaport.cancel(orders), "OpenseaOffer:");
        weth.withdraw(offers[orderHash].offerAmount);
        // call deposit to vault contract
        vault.deposit{value: offers[orderHash].offerAmount}(
            string(abi.encodePacked(offers[orderHash].otaKey)),
            txId,
            regulatorSignData
        );
        offers[orderHash].isCanceled = true;
    }

    // verify signature source https://eips.ethereum.org/EIPS/eip-1271
    /**
     * @notice Verifies that the signer is the owner of the signing contract.
     */
    function isValidSignature(
        bytes32 _hash,
        bytes calldata _signature
    ) external view returns (bytes4) {
        // todo: update here
        // Validate signatures
        if (recoverSigner(_hash, _signature) == offers[_hash].signer) {
            return 0x1626ba7e;
        } else {
            return 0xffffffff;
        }
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
}