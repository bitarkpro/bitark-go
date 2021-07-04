package service

import (
	"bufio"
	"io/ioutil"
	"mime/multipart"

	shell "github.com/ipfs/go-ipfs-api"
)

type IPFSService struct {
	Shell *shell.Shell
}

func NewIPFSService(endport string) *IPFSService {
	return &IPFSService{
		Shell: shell.NewShell(endport),
	}
}

func (i *IPFSService) Upload(file multipart.File) (hash string, err error) {
	hash, err = i.Shell.Add(bufio.NewReader(file))
	return
}

func (i *IPFSService) DownLoad(cid string) ([]byte, error) {
	read, err := i.Shell.Cat(cid)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(read)
	return body, nil
}
