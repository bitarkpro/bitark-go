package service

import (
	"github.com/bitarkpro/bitark-go/client"
	"github.com/centrifuge/go-substrate-rpc-client/v3/types"
)

func TransactionTransferAwaitCallBack(endpoint string, network uint8, privateKey string, toAddress string, amount uint64) (hash string, err error) {
	c, err := client.New(endpoint)
	if err != nil {
		panic(err)
	}
	sig := makeTransactionTransfer(c, network, privateKey, toAddress, amount)
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

func TransactionTransfer(endpoint string, network uint8, privateKey string, toAddress string, amount uint64) (hash string, err error) {
	c, err := client.New(endpoint)
	if err != nil {
		panic(err)
	}
	sig := makeTransactionTransfer(c, network, privateKey, toAddress, amount)
	var result interface{}
	err = c.C.Client.Call(&result, "author_submitExtrinsic", sig)
	if err != nil {
		panic(err)
	}
	return result.(string), nil
}
