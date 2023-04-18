package models

type FP920012I struct {
	AccountKeyNum  string `json:"accountKeyNum"`
	AccountKeyType string `json:"accountKeyType"`
}
type FP920012O struct {
	AccountNum string `json:"accountNum"`
	BankCode   string `json:"bankCode"`
}

func (*FP920012I) GetServiceKey() string {
	return "FP920012"
}
