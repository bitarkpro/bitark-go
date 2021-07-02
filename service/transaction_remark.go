package service

import (
	"github.com/bitarkpro/bitark-go/client"
	"github.com/centrifuge/go-substrate-rpc-client/v3/types"
)

func TransactionRemarkAwaitCallBack(endpoint string, network uint8, privateKey string, message interface{}) (hash string, err error) {
	c, err := client.New(endpoint)
	sig := makeTransactionRemark(c, network, privateKey, message)
	var ex types.Extrinsic
	err = types.DecodeFromHexString(sig, &ex)
	sub, err := c.C.RPC.Author.SubmitAndWatchExtrinsic(ex)
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()
	for {
		status := <-sub.Chan()
		if status.IsInBlock {
			return status.AsInBlock.Hex(), nil
		}
	}
}

func TransactionRemark(endpoint string, network uint8, privateKey string, message interface{}) (hash string, err error) {
	c, err := client.New(endpoint)
	sig := makeTransactionRemark(c, network, privateKey, message)
	var result interface{}
	err = c.C.Client.Call(&result, "author_submitExtrinsic", sig)
	if err != nil {
		panic(err)
	}
	return result.(string), nil
}
