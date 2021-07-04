package service

import (
	"github.com/bitarkpro/bitark-go/chain"
	"github.com/bitarkpro/bitark-go/chain/crypto"
	"github.com/bitarkpro/bitark-go/client"
	"github.com/centrifuge/go-substrate-rpc-client/v3/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v3/types"
)

type ChainService struct {
	ChainEndport string
	ChainNetwork uint8
	Client       *client.Client
}

func NewChainService(endpoint string, network uint8) *ChainService {
	c, err := client.New(endpoint)
	if err != nil {
		panic(err)
	}
	return &ChainService{
		ChainEndport: endpoint,
		ChainNetwork: network,
		Client:       c,
	}
}

func (n *ChainService) TransactionTransferBatch(privateKey string, addressAmount map[string]uint64) (hash string, err error) {
	sig := makeTransactionTransferBatch(n.Client, n.ChainNetwork, privateKey, addressAmount)
	var result interface{}
	err = n.Client.C.Client.Call(&result, "author_submitExtrinsic", sig)
	if err != nil {
		panic(err)
	}
	return result.(string), nil
}

func (n *ChainService) TransactionRemark(privateKey string, message interface{}) (hash string, err error) {
	sig := makeTransactionRemark(n.Client, n.ChainNetwork, privateKey, message)
	var result interface{}
	err = n.Client.C.Client.Call(&result, "author_submitExtrinsic", sig)
	if err != nil {
		panic(err)
	}
	return result.(string), nil
}

func (n *ChainService) TransactionTransfer(privateKey string, toAddress string, amount uint64) (hash string, err error) {
	sig := makeTransactionTransfer(n.Client, n.ChainNetwork, privateKey, toAddress, amount)
	var result interface{}
	err = n.Client.C.Client.Call(&result, "author_submitExtrinsic", sig)
	if err != nil {
		panic(err)
	}
	return result.(string), nil
}

func (n *ChainService) GetAccountInfo(address string) (account *types.AccountInfo, err error) {
	return n.Client.GetAccountInfo(address)
}

func (n *ChainService) GetAccountInfoByPrivateKey(privateKey string) (address string, publicKey string, err error) {
	k, err := signature.KeyringPairFromSecret(privateKey, n.ChainNetwork)
	return k.Address, string(k.PublicKey), err
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

func transactionBase(c *client.Client, network uint8, privateKey string) (*chain.Transaction, *chain.MetadataExpand) {
	k, err := signature.KeyringPairFromSecret(privateKey, network)
	accountInfo, err := c.GetAccountInfo(k.Address)
	if err != nil {
		panic(err)
	}
	nonce := uint64(accountInfo.Nonce)
	transaction := chain.NewTransaction(k.Address, nonce)
	ed, err := chain.NewMetadataExpand(c.Meta)
	if err != nil {
		panic(err)
	}
	return transaction, ed
}
