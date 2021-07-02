package model

type ExtrinsicDecodeResponse struct {
	AccountId          string                 `json:"account_id"`
	CallCode           string                 `json:"call_code"`
	CallModule         string                 `json:"call_module"`
	Era                string                 `json:"era"`
	Nonce              int64                  `json:"nonce"`
	VersionInfo        string                 `json:"version_info"`
	Signature          string                 `json:"signature"`
	Params             []ExtrinsicDecodeParam `json:"params"`
	CallModuleFunction string                 `json:"call_module_function"`
	Length             int                    `json:"length"`
}
