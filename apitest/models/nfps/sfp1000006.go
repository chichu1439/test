package models

type FP100006I struct {
	ReceivingBankId string `json:"receivingBankId"`
	Data            string `json:"data"`
}

type FP100006O struct {
	EncryptData string `json:"encryptData"`
}

func (*FP100006I) GetServiceKey() string {
	return "FP100006"
}