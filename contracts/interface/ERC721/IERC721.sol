// SPDX-License-Identifier: Spartan
pragma solidity ^0.8.0;

/**
 * @dev Required interface of an ERC721 compliant contract.
 */
interface IERC721 {
    /**
     * @dev Emitted when `tokenId` token is transferred from `from` to `to`.
     */
    event Transfer(
        address indexed from,
        address indexed to,
        uint256 indexed tokenId
    );

    /**
     * @dev Emitted when `owner` enables `approved` to manage the `tokenId` token.
     */
    event Approval(
        address indexed owner,
        address indexed approved,
        uint256 indexed tokenId
    );

    /**
     * @dev Emitted when `owner` enables or disables (`approved`) `operator` to manage all of its assets.
     */
    event ApprovalForAll(
        address indexed owner,
        address indexed operator,
        bool approved
    );

    /**
     * @dev Mints `tokenURI` and transfers it to `to`.
     *
     * Requirements:
     *
     * - `to` cannot be the zero address.
     *
     * Emits a {Transfer} event.
     */
    function mint(
        address to,
        string memory name,
        string memory symbol,
        string memory tokenURI
    ) external;

    /**
     * @dev Safely mints `tokenURI` and transfers it to `to`.
     *
     * Requirements:
     *
     * - `to` cannot be the zero address.
     * - If `to` refers to a smart contract, it must implement {IERC721Receiver-onERC721Received}, which is called upon a safe transfer.
     *
     * Emits a {Transfer} event.
     */
    function safeMint(
        address to,
        string memory name,
        string memory symbol,
        string memory tokenURI,
        bytes memory data
    ) external;

    /**
     * @dev Gives permission to `to` to transfer `tokenId` token to another account.
     *
     * Requirements:
     *
     * - sender must be the owner only.
     * - `to` cannot be the zero address.
     * - `tokenId` must exist.
     *
     * Emits an {Approval} event.
     */
    function approve(address to, uint256 tokenId) external;

    /**
     * @dev Returns the account approved for `tokenId` token.
     *
     * Requirements:
     *
     */
    function getApproved(uint256 tokenId) external view returns (address);

    /**
     * @dev Approve or remove `operator` as an operator for the caller.
     * Operators can call {transferFrom} or {safeTransferFrom} for any token owned by the caller.
     *
     * Requirements:
     *
     * - The `operator` cannot be the caller.
     *
     * Emits an {ApprovalForAll} event.
     */
    function setApprovalForAll(address operator, bool approved) external;

    /**
     * @dev Returns if the `operator` is allowed to manage all of the assets of `owner`.
     *
     * See {setApprovalForAll}
     */
    function isApprovedForAll(address owner, address operator)
        external
        view
        returns (bool);

    /**
     * @dev Safely transfers `tokenId` token from `from` to `to`, checking first that contract recipients
     * are aware of the ERC721 protocol to prevent tokens from being forever locked.
     *
     * Requirements:
     * - sender must have call method permission.
     * - `from `&`to` are must be a available `token` account.
     * - `token` must be available.
     * - sender & from & to are must be belong to the same platform;
     * - sender must be the owner or approved.
     *
     * transfer:
     * - `from` cannot be the zero address.
     * - `to` cannot be the zero address.
     * - `tokenId` token must exist and be owned by `from`.
     * - If the caller is not `from`, it must be have been allowed to move this token by either {approve} or {setApprovalForAll}.
     * - If `to` refers to a smart contract, it must implement {IERC721Receiver-onERC721Received}, which is called upon a safe transfer.
     *
     * Emits a {Transfer} event.
     */
    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId,
        bytes memory data
    ) external;

    /**
     * @dev Transfers `tokenId` token from `from` to `to`.
     *
     * Requirements:
     * - sender must be the owner or approved.
     *
     * Emits a {Transfer} event.
     */
    function transferFrom(
        address from,
        address to,
        uint256 tokenId
    ) external;

    /**
     * @dev Burns a token.
     *
     * Requirements:
     * - sender must own `tokenId` or be an approved operator.
     */
    function burn(uint256 tokenId) external;

    /**
     * @dev Returns the number of tokens in ``owner``'s account.
     *
     * Requirements:
     *
     */
    function balanceOf(address owner) external view returns (uint256);

    /**
     * @dev Returns the owner of the `tokenId` token.
     *
     * Requirements:
     */
    function ownerOf(uint256 tokenId) external view returns (address);

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
