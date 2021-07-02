package service

import "github.com/bitarkpro/bitark-go/client"

func TransactionTransferBatch(
	endpoint string,
	network uint8,
	privateKey string,
	addressAmount map[string]uint64) (hash string, err error) {
	c, err := client.New(endpoint)
	if err != nil {
		panic(err)
	}
	sig := makeTransactionTransferBatch(c, network, privateKey, addressAmount)
	var result interface{}
	err = c.C.Client.Call(&result, "author_submitExtrinsic", sig)
	if err != nil {
		panic(err)
	}
	return result.(string), nil
}

