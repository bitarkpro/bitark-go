package service

import (
	"github.com/bitarkpro/bitark-go/client"
	"github.com/centrifuge/go-substrate-rpc-client/v3/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v3/types"
)

//GetAccountInfo wss://rococo.api.onfinality.io/public-ws
func GetAccountInfo(endpoint string, address string) (account *types.AccountInfo, err error) {
	c, err := client.New(endpoint)
	if err != nil {
		return nil, err
	}
	return c.GetAccountInfo(address)
}

func GetAccountInfoByPrivateKey(network uint8, privateKey string) (address string, publicKey string, err error) {
	k, err := signature.KeyringPairFromSecret(privateKey, network)
	return k.Address, string(k.PublicKey), err
}
