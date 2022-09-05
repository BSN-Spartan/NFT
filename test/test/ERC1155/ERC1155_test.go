package ERC1155

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

func TestSafeMint(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC1155Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	amount := new(big.Int).SetUint64(3)
	data := []byte{0x1}
	tx, err := session.SafeMint(common.HexToAddress(owner), "sparton_nft", "sparton_nft", amount, "sparton_nft", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("tx Hash: %s", tx.Hash().String()))
}

func TestSafeMintBatch(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC1155Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	tokenName := []string{"sparton_nft_1", "sparton_nft_2"}
	tokenSymbol := []string{"sparton_nft_1", "sparton_nft_2"}
	tokenURIs := []string{"http://sparton.json", "http://sparton.json"}
	var amount []*big.Int
	amount = append(amount, new(big.Int).SetUint64(3), new(big.Int).SetUint64(2))
	data := []byte{0x1, 0x2}
	tx, err := session.SafeMintBatch(common.HexToAddress(owner), tokenName, tokenSymbol, amount, tokenURIs, data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("tx Hash: %s", tx.Hash().String()))
}

func TestSetApprovalForAll(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC1155Session(cli, common.HexToAddress(Address))
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
	session, err := server.NewERC1155Session(cli, common.HexToAddress(Address))
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
	session, err := server.NewERC1155Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	data := []byte{0x1}
	tokenId := new(big.Int).SetUint64(5)
	amount := new(big.Int).SetUint64(1)
	tx, err := session.SafeTransferFrom(common.HexToAddress(owner), common.HexToAddress(account), tokenId, amount, data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("tx Hash: %s", tx.Hash().String()))
}

func TestSafeBatchTransferFrom(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC1155Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}

	var tokenIDs []*big.Int
	tokenIDs = append(tokenIDs, new(big.Int).SetUint64(3), new(big.Int).SetUint64(4))

	var amount []*big.Int
	amount = append(amount, new(big.Int).SetUint64(1), new(big.Int).SetUint64(1))

	data := []byte{0x1, 0x2}
	tx, err := session.SafeBatchTransferFrom(common.HexToAddress(owner), common.HexToAddress(account), tokenIDs, amount, data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("tx Hash: %s", tx.Hash().String()))
}

func TestBurn(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC1155Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	tokenId := new(big.Int).SetUint64(1)
	tx, err := session.Burn(common.HexToAddress(owner), tokenId)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("tx Hash: %s", tx.Hash().String()))
}

func TestBurnBatch(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC1155Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	var tokenIDs []*big.Int
	tokenIDs = append(tokenIDs, new(big.Int).SetUint64(2), new(big.Int).SetUint64(3))

	tx, err := session.BurnBatch(common.HexToAddress(owner), tokenIDs)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("tx Hash: %s", tx.Hash().String()))
}

func TestBalanceOf(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC1155Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	tokenId := new(big.Int).SetUint64(1)
	tx, err := session.BalanceOf(common.HexToAddress(owner), tokenId)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("Account balance: %s", tx))
}

func TestBalanceOfBatch(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC1155Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	var tokenIDs []*big.Int
	tokenIDs = append(tokenIDs, new(big.Int).SetUint64(3), new(big.Int).SetUint64(3))

	var owners []common.Address
	owners = append(owners, common.HexToAddress(owner), common.HexToAddress(account))

	tx, err := session.BalanceOfBatch(owners, tokenIDs)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("owners balance: %s", tx))
}

func TestTokenName(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC1155Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	tokenId := new(big.Int).SetUint64(3)
	tx, err := session.TokenName(tokenId)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("NFT name: %s", tx))
}

func TestTokenSymbol(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC1155Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	tokenId := new(big.Int).SetUint64(2)
	tx, err := session.TokenSymbol(tokenId)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("NFT symbol: %s", tx))
}

func TestTokenURI(t *testing.T) {
	cli := server.NewETHClientByURL(t, url, key)
	session, err := server.NewERC1155Session(cli, common.HexToAddress(Address))
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
	session, err := server.NewERC1155Session(cli, common.HexToAddress(Address))
	if err != nil {
		t.Fatal(err)
	}
	tx, err := session.GetLatestTokenId()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("tokenId: %s", tx))
}
