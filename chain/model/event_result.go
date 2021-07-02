package model

type EventResult struct {
	From         string `json:"from"`
	To           string `json:"to"`
	Amount       string `json:"amount"`
	ExtrinsicIdx int    `json:"extrinsic_idx"`
	EventIdx     int    `json:"event_idx"`
	Status       string `json:"status"`
	Weight       int64  `json:"weight"` //权重
}
