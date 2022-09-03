// SPDX-License-Identifier: Spartan

pragma solidity ^0.8.0;

/**
 * @title ERC721 receiver interface
 * @dev Interface for any contract that wants to support safeTransfers
 * from ERC721 asset contracts.
 */
interface IERC721Receiver {
    /**
     * @dev Whenever an {IERC721} `ERCId` ERC is transferred to this contract via {IERC721-safeTransferFrom}
     * by `operator` from `from`, this function is called.
     *
     * It must return its Solidity selector to confirm the ERC transfer.
     * If any other value is returned or the interface is not implemented by the recipient, the transfer will be reverted.
     *
     * The selector can be obtained in Solidity with `IERC721Receiver.onERC721Received.selector`.
     */
    function onERC721Received(
        address operator,
        address from,
        uint256 ERCId,
        bytes calldata data
    ) external returns (bytes4);

     /**
     * @dev Whenever an {IERC721} `ERCId` ERC is transferred to this contract via {IERC721-safeBatchTransferFrom}
     * by `operator` from `from`, this function is called.
     *
     * It must return its Solidity selector to confirm the ERC transfer.
     * If any other value is returned or the interface is not implemented by the recipient, the transfer will be reverted.
     *
     * The selector can be obtained in Solidity with `IERC721Receiver.onERC721Received.selector`.
     */
    function onERC721BatchReceived(
        address operator,
        address from,
        uint256[] calldata ERCIds,
        bytes calldata data
    ) external returns (bytes4);
}
