package models

type PF000007I struct {
	FeeTypeId      string `json:"feeTypeId"`
	FeeTypeName    string `json:"feeTypeName"`
	AccountingCode string `json:"accountingCode"`
	IsAmortization string `json:"isAmortization"`
}

type PF000007O struct {
	FeeTypeId   string `json:"feeTypeId"`
	FeeTypeName string `json:"feeTypeName"`
}

func (*PF000007I) GetServiceKey() string {
	return "PF000007"
}
