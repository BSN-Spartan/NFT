package server

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"spartan_nft_test/client"
	"spartan_nft_test/contract/eth/ERC1155"
	"spartan_nft_test/contract/eth/ERC721"
)

func NewERC721Session(cli *client.ETHClient, addr common.Address) (*ERC721.ERC721Session, error) {
	session, err := ERC721.NewERC721(addr, cli.Client())
	if err != nil {
		return nil, err
	}
	transactOpts := cli.GetTrandOpts()
	callOpts := bind.CallOpts{From: transactOpts.From}
	return &ERC721.ERC721Session{session, callOpts, *transactOpts}, err
}

func NewERC1155Session(cli *client.ETHClient, addr common.Address) (*ERC1155.ERC1155Session, error) {
	session, err := ERC1155.NewERC1155(addr, cli.Client())
	if err != nil {
		return nil, err
	}
	transactOpts := cli.GetTrandOpts()
	callOpts := bind.CallOpts{From: transactOpts.From}
	return &ERC1155.ERC1155Session{session, callOpts, *transactOpts}, err
}
