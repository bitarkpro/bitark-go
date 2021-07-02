package model

type Block struct {
	Extrinsics []string `json:"extrinsics"`
	Header     Header   `json:"header"`
}
