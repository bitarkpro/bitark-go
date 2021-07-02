package model

type UtilityParamsValue struct {
	CallModule   string                  `json:"call_module"`
	CallFunction string                  `json:"call_function"`
	CallIndex    string                  `json:"call_index"`
	CallArgs     []UtilityParamsValueArg `json:"call_args"`
}

type UtilityParamsValueArg struct {
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	Value    interface{} `json:"value"`
	ValueRaw string      `json:"value_raw"`
}
