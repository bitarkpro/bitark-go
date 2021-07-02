package model

type Header struct {
	ParentHash     string      `json:"parentHash"`
	Number         string      `json:"number"`
	StateRoot      string      `json:"stateRoot"`
	ExtrinsicsRoot string      `json:"extrinsicsRoot"`
	Digest         interface{} `json:"digest"`
}
