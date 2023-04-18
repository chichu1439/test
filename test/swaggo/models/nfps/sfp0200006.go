package models

type FP020006I struct {
	AccountKeyNum string `json:"accountKeyNum" validate:"required"`
}
type FP020006O struct {
	//  AddressingInfo
}

func (*FP020006I) GetServiceKey() string {
	return "FP020006"
}
