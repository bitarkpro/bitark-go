package model

type BlockResponse struct {
	Height     int64                `json:"height"`
	ParentHash string               `json:"parent_hash"`
	BlockHash  string               `json:"block_hash"`
	Timestamp  int64                `json:"timestamp"`
	Extrinsic  []*ExtrinsicResponse `json:"extrinsic"`
}
