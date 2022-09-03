// SPDX-License-Identifier: Spartan
pragma solidity ^0.8.0;

/**
 * @dev Required interface of an ERC1155 compliant contract, as defined in the
 */
interface IERC1155 {
    /**
     * @dev Emitted when `value` tokens of token type `tokenId` are transferred from `from` to `to` by `operator`.
     */
    event TransferSingle(
        address indexed operator,
        address indexed from,
        address indexed to,
        uint256 tokenId,
        uint256 amount
    );

    /**
     * @dev Equivalent to multiple {TransferSingle} events, where `operator`, `from` and `to` are the same for all
     * transfers.
     */
    event TransferBatch(
        address indexed operator,
        address indexed from,
        address indexed to,
        uint256[] tokenIds,
        uint256[] amounts
    );

    /**
     * @dev Emitted when `account` grants or revokes permission to `operator` to transfer their tokens, according to
     * `approved`.
     */
    event ApprovalForAll(
        address indexed owner,
        address indexed operator,
        bool approved
    );

    /**
     * @dev Emitted when the URI for token type `tokenId` changes to `value`, if it is a non-programmatic URI.
     *
     * If an {URI} event was emitted for `tokenId`, the standard
     * https://eips.ethereum.org/EIPS/eip-1155#metadata-extensions[guarantees] that `value` will equal the value
     * returned by {IERC1155MetadataURI-uri}.
     */
    event URI(string value, uint256 indexed tokenId);

    /**
     * @dev Creates `amount` tokens of token type `tokenId`, and assigns them to `to`.
     *
     * Emits a {TransferSingle} event.
     *
     * Requirements:
     *
     * - `to` cannot be the zero address.
     * - If `to` refers to a smart contract, it must implement {IERC1155Receiver-onERC1155Received} and return the
     */
    function safeMint(
        address to,
        string memory name,
        string memory symbol,
        uint256 amount,
        string memory tokenURI,
        bytes memory data
    ) external;

    /**
     * @dev Bulk create `amount` tokens with token type `tokenId` and assign them to `to`.
     *
     * Emits a {TransferBatch} event.
     *
     * Requirements:
     *
     * - `to` cannot be the zero address.
     * - If `to` refers to a smart contract, it must implement {IERC1155Receiver-onERC1155Received} and return the
     */
    function safeMintBatch(
        address to,
        string[] memory names,
        string[] memory symbols,
        uint256[] memory amounts,
        string[] memory tokenURIs,
        bytes memory data
    ) external;

    /**
     * @dev Authorize all tokenIds under the owner to the authorizer account
     *
     * Requirements:
     *
     */
    function setApprovalForAll(address operator, bool approved) external;

    /**
     * @dev Check whether all tokenIds under the owner are authorized to the authorizer account
     *
     * Requirements:
     *
     */
    function isApprovedForAll(address owner, address operator)
        external
        view
        returns (bool);

    /**
     * @dev Transfers `amount` tokens of token type `tokenId` from `from` to `to`.
     *
     * Requirements:
     *
     * - `to` cannot be the zero address.
     * - If the caller is not `from`, it must have been approved to spend ``from``'s tokens via {setApprovalForAll}.
     * - `from` must have a balance of tokens of type `tokenId` of at least `amount`.
     * - If `to` refers to a smart contract, it must implement {IERC1155Receiver-onERC1155Received} and return the
     * acceptance magic value.
     */
    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId,
        uint256 amount,
        bytes memory data
    ) external;

    /**
     * @dev Batch transfer "amount" tokens with token type "tokenId" from "from" to "to".
     *
     *
     * Requirements:
     *
     * - `tokenIds` and `amounts` must have the same length.
     * - If `to` refers to a smart contract, it must implement {IERC1155Receiver-onERC1155BatchReceived} and return the
     * acceptance magic value.
     */
    function safeBatchTransferFrom(
        address from,
        address to,
        uint256[] memory tokenIds,
        uint256[] memory amounts,
        bytes memory data
    ) external;

    /**
     * @dev Destroy a single token owned by the owner
     *
     * Requirements:
     *
     */
    function burn(address owner, uint256 tokenId) external;

    /**
     * @dev Destroy the token list owned by the owner in batches
     *
     * Requirements:
     *
     */
    function burnBatch(address owner, uint256[] memory tokenIds) external;

    /**
     * @dev Returns the amount of tokens of token type `tokenId` owned by `owner`.
     *
     * Requirements:
     *
     * - `owner` cannot be the zero address.
     */
    function balanceOf(address owner, uint256 tokenId)
        external
        view
        returns (uint256);

    /**
     * @dev Returns the amount of tokens of token type `tokenIds` owned by `owners`.
     *
     * Requirements:
     *
     * - `owners` and `tokenIds` must have the same length.
     */
    function balanceOfBatch(address[] memory owners, uint256[] memory tokenIds)
        external
        view
        returns (uint256[] memory);

    /**
     * @dev Returns the token name of the `tokenId` token
     *
     * Requirements:
     */
    function tokenName(uint256 tokenId) external view returns (string memory);

    /**
     * @dev Returns the token symbol of the `tokenId` token
     *
     * Requirements:
     */
    function tokenSymbol(uint256 tokenId) external view returns (string memory);

    /**
     * @dev Returns the token URI of the `tokenId` token
     *
     * Requirements:
     */
    function tokenURI(uint256 tokenId) external view returns (string memory);

    /**
     * @dev Returns the last tokenId.
     *
     * Requirements:
     */
    function getLatestTokenId() external view returns (uint256);
}
