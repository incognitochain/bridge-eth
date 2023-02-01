// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155Receiver.sol";

enum Side { Buy, Sell }
enum SignatureVersion { Single, Bulk }
enum AssetType { ERC721, ERC1155 }

struct Fee {
    uint16 rate;
    address payable recipient;
}

struct Order {
    address trader;
    Side side;
    address matchingPolicy;
    address collection; // token address
    uint256 tokenId; // identifier
    uint256 amount;
    address paymentToken;
    uint256 price;
    uint256 listingTime;
    /* Order expiration timestamp - 0 for oracle cancellations. */
    uint256 expirationTime;
    Fee[] fees;
    uint256 salt;
    bytes extraParams;
}

struct Input {
    Order order;
    uint8 v;
    bytes32 r;
    bytes32 s;
    bytes extraSignature;
    SignatureVersion signatureVersion;
    uint256 blockNumber;
}

struct Execution {
    Input sell;
    Input buy;
}

struct OpenseaTrades {
    uint256 value;
    bytes tradeData;
}

struct ERC20Details {
    address[] tokenAddrs;
    uint256[] amounts;
}

struct ERC1155Details {
    address tokenAddr;
    uint256[] ids;
    uint256[] amounts;
}

struct ConverstionDetails {
    bytes conversionData;
}

struct AffiliateDetails {
    address affiliate;
    bool isActive;
}

struct SponsoredMarket {
    uint256 marketId;
    bool isActive;
}

struct TradeDetails {
    uint256 marketId;
    uint256 value;
    bytes tradeData;
}

struct Market {
    address proxy;
    bool isLib;
    bool isActive;
}

struct TokenInfo {
    AssetType tokenType;
    address collection;
    uint256 identifier;
    uint256 amount;
}

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

interface IMatchingPolicy {
    function canMatchMakerAsk(Order calldata makerAsk, Order calldata takerBid)
    external
    view
    returns (
        bool,
        uint256,
        uint256,
        uint256,
        AssetType
    );

    function canMatchMakerBid(Order calldata makerBid, Order calldata takerAsk)
    external
    view
    returns (
        bool,
        uint256,
        uint256,
        uint256,
        AssetType
    );
}

interface IBlur {
    function batchBuyWithETH(
        TradeDetails[] memory tradeDetails
    ) payable external;

    function batchBuyWithERC20s(
        ERC20Details memory erc20Details,
        TradeDetails[] memory tradeDetails,
        ConverstionDetails[] memory converstionDetails,
        address[] memory dustTokens
    ) payable external;

    function execute(Input calldata sell, Input calldata buy) external payable;

    function bulkExecute(Execution[] calldata executions) external payable;
}

contract BuyBlurProxy {
    IBlur constant public blur1 = IBlur(0x9A886f7cB772Fc990dd8265cb1a9547DB6CfE2e3);
    IBlur constant public blur2 = IBlur(0x39da41747a83aeE658334415666f3EF92DD0D541);

    // process buy
    function batchBuyWithETH(TradeDetails[] calldata tradeDetails, address recipient, TokenInfo[] memory listTokens) external payable {
        blur2.batchBuyWithETH{value: msg.value}(tradeDetails);
        // transfer NFT token
        for (uint256 i = 0; i < listTokens.length; i++) {
            _transferNFT(listTokens[i].tokenType, recipient, listTokens[i].collection, listTokens[i].identifier, listTokens[i].amount);
        }
    }

    function execute(Input calldata sell, Input calldata buy, address recipient) external payable {
        blur1.execute{value: msg.value}(sell, buy);
        _processExecute(sell.order, buy.order, recipient);
    }

    function bulkExecute(Execution[] calldata executions, address recipient) external payable {
        blur1.bulkExecute{value: msg.value}(executions);
        for (uint i = 0; i < executions.length; i++) {
            Execution memory temp = executions[i];
            _processExecute(temp.sell.order, temp.buy.order, recipient);
        }
    }

    function _processExecute(Order memory sell, Order memory buy, address recipient) internal {
        (uint256 tokenId, uint256 amount, AssetType assetType) = _getAssetType(sell, buy);
        _transferNFT(assetType, recipient, sell.collection, tokenId, amount);
    }

    function _transferNFT(AssetType assetType, address to, address token, uint256 identifier, uint256 amount) internal {
        if (assetType == AssetType.ERC721) {
            IERC721(token).transferFrom(address(this), to, identifier);
        } else if (assetType == AssetType.ERC1155) {
            IERC1155(token).safeTransferFrom(address(this), to, identifier, amount, bytes(""));
        } else {
            revert("PB: the item type not supported");
        }
    }

    function _getAssetType(Order memory sell, Order memory buy)
    internal
    view
    returns (uint256 tokenId, uint256 amount, AssetType assetType)
    {
        if (sell.listingTime <= buy.listingTime) {
            /* Seller is maker. */
            (,, tokenId, amount, assetType) = IMatchingPolicy(sell.matchingPolicy).canMatchMakerAsk(sell, buy);
        } else {
            /* Buyer is maker. */
            (,, tokenId, amount, assetType) = IMatchingPolicy(buy.matchingPolicy).canMatchMakerBid(buy, sell);
        }
    }

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