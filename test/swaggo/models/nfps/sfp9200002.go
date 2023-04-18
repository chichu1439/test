package models

type FP920002I struct {
	AccountKeyNum string `json:"accountKeyNum"`
	AccountNum    string `json:"accountNum"`
	BankCode      string `json:"bankCode"`
}
type FP920002O struct {
}

func (*FP920002I) GetServiceKey() string {
	return "FP920002"
}
