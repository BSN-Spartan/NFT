# Spartan-NFT | contract

This document is a guide about the official NFT contract of the spartan network.

## Overview

Spartan-NFT contract contains ERC721 contract and ERC1155 contract.

The ERC721 contract is based on the Ethereum `eip-721` standard. The interface {IERC721} implements all the core functions required for compliance implementation, and has been extended to provide more comprehensive functions such as mint. {ERC721} fully implements {IERC721} interface. (eip-721 referrence: https://eips.ethereum.org/EIPS/eip-721)

The ERC1155 contract is based on the Ethereum `eip-1155` standard, {ERC1155} implements all the functions provided by the {IERC1155} interface, and provides more convenient functions such as batch generation. (eip-721 reference: https://eips.ethereum.org/EIPS/eip-1155)

Spartan-NFT contract applies the ERC1967 Proxy contract，This contract implements an upgradable proxy. It is upgradable because the call is delegated to the implementation address which can be changed. The address is stored in storage at the location specified by `eip1967`.

**How to Use**

Get the NFT contract source code by command:

```
$ git clone git@github.com:BSN-Spartan/NFT.git
```

The tool`Hardhat` is used to develop the Spartan NFT contract, and this contract support custom development. Developers can modify the contract directly in the contacts folder.



Compile the source code:

```
$ npx hardhat compile
```

You can use the commands that come with Hardhat to compile,  please remember to load the required packages by `npm install`before compiling.



If you want to refer to the interface of the Spartan-NFT contract for development, you can directly import the {IERC721} and {IERC1155} interfaces.

```solidity
pragma solidity ^0.8.0;

import "https://github.com/BSN-Spartan/NFT/blob/main/contracts/interface/ERC721/IERC721.sol";

contract ERC721 is IERC721{
    
}
```

If you are not familiar with the smart contracts development, you can directly import the Spartan-NFT contract into Remix for contract deployment.

## License

Spartan-NFT Contracts is released under the [Spartan License](https://github.com/BSN-Spartan/NFT/blob/main/LICENSE).

