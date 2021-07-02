package model

type ExtrinsicResponse struct {
	Type            string `json:"type"`   //Transfer or another
	Status          string `json:"status"` //success or fail
	Txid            string `json:"txid"`
	FromAddress     string `json:"from_address"`
	ToAddress       string `json:"to_address"`
	Amount          string `json:"amount"`
	Fee             string `json:"fee"`
	Signature       string `json:"signature"`
	Nonce           int64  `json:"nonce"`
	Era             string `json:"era"`
	ExtrinsicIndex  int    `json:"extrinsic_index"`
	EventIndex      int    `json:"event_index"`
	ExtrinsicLength int    `json:"extrinsic_length"`
}
