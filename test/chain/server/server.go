package server

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"spartan_nft_test/client"
	"testing"
)

func NewETHClientByURL(t *testing.T, url, key string) *client.ETHClient {
	fmt.Println(url)
	hex, err := hexutil.Decode(key)
	if err != nil {
		t.Fatal(err)
	}
	pk, err := crypto.ToECDSA(hex)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(crypto.PubkeyToAddress(pk.PublicKey).String())
	ethCli, err := ethclient.Dial(url)
	if err != nil {
		t.Fatal(err)
	}

	cli, err := client.NewETHClient(ethCli, pk)
	if err != nil {
		t.Fatal(err)
	}

	return cli

}

func NewETHCliByURL(url, key string) (*client.ETHClient, error) {
	fmt.Println(url)
	hex, err := hexutil.Decode(key)
	if err != nil {
		return nil, err
	}
	pk, err := crypto.ToECDSA(hex)
	if err != nil {
		return nil, err
	}
	fmt.Println(crypto.PubkeyToAddress(pk.PublicKey).String())
	ethCli, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	cli, err := client.NewETHClient(ethCli, pk)
	if err != nil {
		return nil, err
	}

	return cli, nil

}
