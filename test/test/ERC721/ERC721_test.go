package ERC721

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"spartan_nft_test/chain/server"
	"testing"
)

const (
	url     = "" //Node URL
	key     = "" //Private key
	Address = "" //Contract address
	owner   = "" //owner account
	account = "" //ordinary account
)

func TestMint(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC721Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	tx, err := session.Mint(common.HexToAddress(owner), "sparton_nft", "sparton_nft", "sparton_nft")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("tx Hash: %s", tx.Hash().String()))
}

func TestSafeMint(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC721Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	data := []byte{0x1}
	tx, err := session.SafeMint(common.HexToAddress(owner), "sparton_nft", "sparton_nft", "sparton_nft", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("tx Hash: %s", tx.Hash().String()))
}

func TestApprove(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC721Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	tokenId := new(big.Int).SetUint64(4)
	tx, err := session.Approve(common.HexToAddress(account), tokenId)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("tx Hash: %s", tx.Hash().String()))
}

func TestGetApproved(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC721Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	tokenId := new(big.Int).SetUint64(4)
	tx, err := session.GetApproved(tokenId)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("Account Address: %s", tx.String()))
}

func TestSetApprovalForAll(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC721Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	tx, err := session.SetApprovalForAll(common.HexToAddress(account), true)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("tx Hash: %s", tx.Hash().String()))
}

func TestIsApprovedForAll(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC721Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	tx, err := session.IsApprovedForAll(common.HexToAddress(owner), common.HexToAddress(account))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("Is ApprovedForAll:%t", tx))
}

func TestSafeTransferFrom(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC721Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	data := []byte{0x1}
	tokenId := new(big.Int).SetUint64(1)
	tx, err := session.SafeTransferFrom(common.HexToAddress(owner), common.HexToAddress(account), tokenId, data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("tx Hash: %s", tx.Hash().String()))
}

func TestTransferFrom(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC721Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	tokenId := new(big.Int).SetUint64(3)
	tx, err := session.TransferFrom(common.HexToAddress(owner), common.HexToAddress(account), tokenId)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("tx Hash: %s", tx.Hash().String()))
}

func TestBurn(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC721Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	tokenId := new(big.Int).SetUint64(2)
	tx, err := session.Burn(tokenId)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("tx Hash: %s", tx.Hash().String()))
}

func TestBalanceOf(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC721Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	tx, err := session.BalanceOf(common.HexToAddress(owner))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("nft amount: %s", tx))
}

func TestOwnerOf(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC721Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	tokenId := new(big.Int).SetUint64(2)
	tx, err := session.OwnerOf(tokenId)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("Owner address: %s", tx.String()))
}

func TestTokenName(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC721Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	tokenId := new(big.Int).SetUint64(4)
	tx, err := session.TokenName(tokenId)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("NFT name: %s", tx))
}

func TestTokenSymbol(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC721Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	tokenId := new(big.Int).SetUint64(4)
	tx, err := session.TokenSymbol(tokenId)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("NFT symbol: %s", tx))
}

func TestTokenURI(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC721Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	tokenId := new(big.Int).SetUint64(4)
	tx, err := session.TokenURI(tokenId)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("NFT URI: %s", tx))
}

func TestGetLatestTokenID(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC721Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	tx, err := session.GetLatestTokenId()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("tokenId: %s", tx))
}
