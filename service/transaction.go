package service

import (
	"github.com/bitarkpro/bitark-go/chain"
	"github.com/bitarkpro/bitark-go/chain/crypto"
	"github.com/bitarkpro/bitark-go/client"
)

func makeTransactionRemark(c *client.Client, network uint8, privateKey string, message interface{}) string {
	transaction, ed := transactionBase(c, network, privateKey)
	smCallIdx, err := ed.MV.GetCallIndex("System", "remark")
	call, err := chain.NewCall(smCallIdx, message)
	if err != nil {
		panic(err)
	}
	transaction.SetGenesisHashAndBlockHash(c.GetGenesisHash(), c.GetGenesisHash()).
		SetSpecAndTxVersion(uint32(c.SpecVersion), uint32(c.TransactionVersion)).
		SetCall(call)
	sig, err := transaction.SignTransaction(privateKey, crypto.Sr25519Type)
	if err != nil {
		panic(err)
	}
	return sig
}

func makeTransactionTransfer(c *client.Client, network uint8, privateKey string, toAddress string, amount uint64) string {
	transaction, ed := transactionBase(c, network, privateKey)
	call, err := ed.BalanceTransferCall(toAddress, amount)
	if err != nil {
		panic(err)
	}
	transaction.
		SetGenesisHashAndBlockHash(c.GetGenesisHash(), c.GetGenesisHash()).
		SetSpecAndTxVersion(uint32(c.SpecVersion), uint32(c.TransactionVersion)).
		SetCall(call)
	sig, err := transaction.SignTransaction(privateKey, crypto.Sr25519Type)
	if err != nil {
		panic(err)
	}
	return sig
}

func makeTransactionTransferBatch(c *client.Client, network uint8, privateKey string, addressAmount map[string]uint64) string {
	transaction, ed := transactionBase(c, network, privateKey)
	call, err := ed.UtilityBatchTxCall(addressAmount, false)
	if err != nil {
		panic(err)
	}
	transaction.
		SetGenesisHashAndBlockHash(c.GetGenesisHash(), c.GetGenesisHash()).
		SetSpecAndTxVersion(uint32(c.SpecVersion), uint32(c.TransactionVersion)).
		SetCall(call)
	sig, err := transaction.SignTransaction(privateKey, crypto.Sr25519Type)
	if err != nil {
		panic(err)
	}
	return sig
}

func transactionBase(c *client.Client, network uint8, privateKey string) (*chain.Transaction, *chain.MetadataExpand) {
	address, _, err := GetAccountInfoByPrivateKey(network, privateKey)
	accountInfo, err := c.GetAccountInfo(address)
	if err != nil {
		panic(err)
	}
	nonce := uint64(accountInfo.Nonce)
	transaction := chain.NewTransaction(address, nonce)
	ed, err := chain.NewMetadataExpand(c.Meta)
	if err != nil {
		panic(err)
	}
	return transaction, ed
}
