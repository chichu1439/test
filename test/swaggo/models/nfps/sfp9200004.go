package models

type FP920004I struct {
	AccountKeyNum string `json:"accountKeyNum"`
}
type FP920004O struct {
	// AddressingInfo
}

func (*FP920004I) GetServiceKey() string {
	return "FP920004"
}
