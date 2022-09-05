package client

import (
	"context"
	"crypto/ecdsa"
	"strconv"
	"strings"

	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func NewETHClient(cli *ethclient.Client, pk *ecdsa.PrivateKey) (*ETHClient, error) {
	contractCli := &ETHClient{
		client: cli,
		key:    pk,
	}

	chainId, err := cli.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	contractCli.chainId = chainId

	return contractCli, nil
}

type ETHClient struct {
	client       *ethclient.Client
	session      interface{}
	key          *ecdsa.PrivateKey
	chainId      *big.Int
	transactOpts *bind.TransactOpts

	GasPrice string
	GasLimit string
}

type DeployContract func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, error)

func (d *ETHClient) GetTrandOpts() *bind.TransactOpts {
	if d.transactOpts == nil {
		d.transactOpts = NewKeyedTransactor(d.key, d.chainId, d.GasPrice, d.GasLimit)
	}
	return d.transactOpts

}

func (d *ETHClient) Deploy(deploy DeployContract) (addr common.Address, txId string, err error) {

	addr, tx, err := deploy(d.GetTrandOpts(), d.client)
	if err != nil {

		fmt.Println("deploy contract has err :", err.Error())
		return
	}

	fmt.Println(fmt.Sprintf("contract deployed success, contract address: %s, tx hash: %s",
		strings.ToLower(addr.String()), tx.Hash().String()))

	return addr, tx.Hash().String(), err
}

type NewSession func(auth *bind.TransactOpts, opts bind.CallOpts, client *ethclient.Client) (interface{}, error)

func (d *ETHClient) Connection(session NewSession) error {
	transactOpts := d.GetTrandOpts()
	callOpts := bind.CallOpts{From: transactOpts.From}
	var err error
	d.session, err = session(transactOpts, callOpts, d.client)

	return err

}

func (d *ETHClient) Client() *ethclient.Client {
	return d.client
}

func (d *ETHClient) Session() interface{} {
	return d.session
}

//1900000000
var gasPrice = new(big.Int).SetInt64(1000000000)

//var gasPrice = new(big.Int).SetInt64(1900000000)

// NewKeyedTransactor is a utility method to easily create a transaction signer
// from a single private key.
func NewKeyedTransactor(key *ecdsa.PrivateKey, chainId *big.Int, gasPrice string, gaslimit string) *bind.TransactOpts {
	keyAddr := crypto.PubkeyToAddress(key.PublicKey)
	to := &bind.TransactOpts{
		//GasPrice: gasPrice,
		//GasLimit: 7863392,

		From: keyAddr,
		Signer: func(signer1 types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {

			signer := types.NewEIP155Signer(chainId)
			if address != keyAddr {
				return nil, errors.New("not authorized to sign this account")
			}
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), key)
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}

	if gasPrice != "" {
		gp, ok := new(big.Int).SetString(gasPrice, 10)
		if ok {
			to.GasPrice = gp
		}
	}

	if gaslimit != "" {
		gl, err := strconv.ParseUint(gaslimit, 10, 64)
		if err != nil {
			to.GasLimit = gl
		}

	}

	return to
}
