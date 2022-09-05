package Deploy

import (
	"spartan_nft_test/chain/deploy"
	"spartan_nft_test/chain/server"
	"testing"
)

const (
	url = "" //Node URL
	key = "" //Private key
)

func TestDeploy(t *testing.T) {

	cli := server.NewETHClientByURL(t, url, key)

	err := deploy.Deploy(cli)

	if err != nil {
		t.Fatal(err)
	}
}
