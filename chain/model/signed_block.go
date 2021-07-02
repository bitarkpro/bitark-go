package model

type SignedBlock struct {
	Block         Block `json:"block"`
	Justification Bytes `json:"justification"`
}
