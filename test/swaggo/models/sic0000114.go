package models

import "github.com/shopspring/decimal"

type SIC0000114I struct {
	LoanReceiptNum    string                  `json:"loanReceiptNum"`
	ModifyType        string                  `json:"modifyType"`
	SkipPeriodCount   int                     `json:"skipPeriodCount"`
	ReducePercentage  decimal.Decimal         `json:"reducePercentage"`
	ReducePeriodCount int                     `json:"reducePeriodCount"`
	NewTotalPeriodNum int                     `json:"newTotalPeriodNum"`
	CurrentPeriod     int                     `json:"currentPeriod"`
	Installment       decimal.Decimal         `json:"installment"`
	MakerComment      string                  `json:"makerComment"`
	AppDate           string                  `json:"appDate"`
	ListAmountType    []SIC0000114IAmountType `json:"listAmountType"`
}
type SIC0000114IAmountType struct {
	AmountType        string          `json:"amountType"`
	Amount            decimal.Decimal `json:"amount"`
	WaiveAmount       decimal.Decimal `json:"waiveAmount"`
	ConvertAmount     decimal.Decimal `json:"convertAmount"`
	ConvertAmountType string          `json:"convertAmountType"`
}
type SIC0000114O struct {
	LoanReceiptNum string `json:"loanReceiptNum"`
	ApplySerialNum string `json:"applySerialNum"`
}

func (*SIC0000114I) GetServiceKey() string {
	return "IC000114"
}
