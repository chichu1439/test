package models

type FP020002I struct {
	AccountKeyNum string `json:"accountKeyNum" validate:"required"`
	AccountNum    string `json:"accountNum" validate:"required"`
	BankCode      string `json:"bankCode" validate:"required"`
}
type FP020002O struct {
}

func (*FP020002I) GetServiceKey() string {
	return "FP020002"
}
