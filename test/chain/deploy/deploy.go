package deploy

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"spartan_nft_test/client"
	"spartan_nft_test/contract/eth/ERC1155"
	"spartan_nft_test/contract/eth/ERC721"
	"spartan_nft_test/contract/eth/proxy"
	"time"
)

const wait = 30

func DeployERC721(cli *client.ETHClient) (common.Address, string, error) {

	return cli.Deploy(func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, error) {
		addr, tx, _, err := ERC721.DeployERC721(auth, backend)

		return addr, tx, err
	})

}

func DeployERC1155(cli *client.ETHClient) (common.Address, string, error) {

	return cli.Deploy(func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, error) {
		addr, tx, _, err := ERC1155.DeployERC1155(auth, backend)

		return addr, tx, err
	})

}

func DeployProxy(cli *client.ETHClient, logicAddr common.Address) (common.Address, string, error) {

	initData, _ := hex.DecodeString("8129fc1c")

	return cli.Deploy(func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, error) {
		addr, tx, _, err := proxy.DeployProxy(auth, backend, logicAddr, initData)

		return addr, tx, err
	})

}

func Deploy(cli *client.ETHClient) (err error) {

	var erc721Logic common.Address
	var erc1155Logic common.Address

	time.Sleep(wait * time.Second)
	erc721Logic, _, err = DeployERC721(cli)
	if err != nil {
		fmt.Println("deploy erc 721 logic has error", err.Error())
		return
	}

	//fmt.Println("erc721 Logic address :", erc721Logic.String())

	time.Sleep(wait * time.Second)
	erc1155Logic, _, err = DeployERC1155(cli)
	if err != nil {
		fmt.Println("deploy erc 1155 logic has error", err.Error())
		return
	}

	//fmt.Println("erc1155 Logic address :", erc1155Logic.String())

	var erc721Proxy common.Address
	var erc1155Proxy common.Address

	time.Sleep(wait * time.Second)
	erc721Proxy, _, err = DeployProxy(cli, erc721Logic)
	if err != nil {
		fmt.Println("deploy erc 721 proxy has error", err.Error())
		return
	}

	fmt.Println("erc721 proxy address :", erc721Proxy.String())

	time.Sleep(wait * time.Second)
	erc1155Proxy, _, err = DeployProxy(cli, erc1155Logic)
	if err != nil {
		fmt.Println("deploy erc 721 proxy has error", err.Error())
		return
	}

	fmt.Println("erc1155 proxy address :", erc1155Proxy.String())

	return

}
