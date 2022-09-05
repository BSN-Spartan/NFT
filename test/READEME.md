# Spartan-NFT | Test

This document is a guide to test the Spartan official NFT contract.

**Caution: This test needs to be run in the go 1.13 or above environment! ! !**

## How to Use

Get the NFT contract source code by command:

```
$ git clone git@github.com:BSN-Spartan/NFT.git
```



Run the following command on the command line to load the dependencies required by `go`.

```
$ go mod tidy
```

There are three test files in the `test` directory where the NFT project is located:

- deploy_test.go: The file used to deploy ERC721 and ERC1155 contracts.
- ERC721_test.go: The file used to test the methods in ERC721 contracts.
- ERC1155_test.go: The file used to test the methods in ERC1155 contracts.



Open file`deploy_test.go`, and modify the parameters in `const`：

```go
const (
	url     = "" //Node URL
	key     = "" //Private key
)
```

Run the test method `TestDeploy` to get the addresses of the ERC721 contract and the ERC1155 contract.

#### Test ERC721Contract

Open file`ERC721_test.go`，and modify the parameters in `const`：

```go
const (
	url     = "" //Node URL
	key     = "" //Private key
	Address = "" //ERC721 contract address
	owner   = "" //account address
	account = "" //account address
)
```

Tips: The owner and account parameters can be any wallet address.

Test and run the ERC721 contract through the following methods:

- TestSafeMint: Securely mint a token. 

- TestApporve：Grant the token operation permission to a user.

- TestGetApproved：Obtain an authorizer account.

- TestSetApprovalForAll：Grant operation permission to a user.

- TestIsApprovedForAll：Check if the user is authorized.

- TestSafeTransferFrom：Securely transfer a token.

- TestTransferFrom：Transfer a token.

- TestBurn：Burn a token.

- TestBalanceOf：Query the token balance of the wallet.

- TestOwnerOf：Query the owner of the token.

- TestTokenName：Query the token name.

- TestTokenSymbol：Query the token symbol.

- TestTokenURI：Query the token URI.

- TestGetLatestTokenID：Query the latest token Id.

  

#### Test ERC1155 Contract

Open file`ERC1155_test.go`，and modify the parameters in `const`：

```go
const (
	url     = "" //Node URL
	key     = "" //Private key
	Address = "" //ERC721 contract address
	owner   = "" //account address
	account = "" //account address
)
```

Tips: The owner and account parameters can be any wallet address.

Test and run the ERC1155 contract through the following methods:

- TestSafeMint: Securely mint multiple tokens.

- TestSafeMintBatch：Securely mint multiple tokens in batches.
- TestSetApprovalForAll：Grant the token operation permission to a user.
- TestIsApprovedForAll：Check if the user is authorized.
- TestSafeTransferFrom：Securely transfer multiple tokens.

- TestSafeBatchTransferFrom: Securely transfer multiple tokens in batches.

- TestBurn：Burn multiple tokens.

- TestBurnBatch：Burn multiple tokens in batches.

- TestBalanceOf：Query the token balance of the wallet.

- TestBalanceOfBatch: Batch query the number of tokens corresponding to the owners.

- TestTokenName：Query the token name.

- TestTokenSymbol：Query the token symbol.

- TestTokenURI：Query the token URI.

- TestGetLatestTokenID：Query the latest token Id.
