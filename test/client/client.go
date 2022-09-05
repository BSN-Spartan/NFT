package client

import "github.com/ethereum/go-ethereum/common"

type Clienter interface {
	Deploy(args ...common.Address) (addr common.Address, err error)
	Connection(addr common.Address) error
	Session() interface{}
}
