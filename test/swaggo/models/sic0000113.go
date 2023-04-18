package models

import "github.com/shopspring/decimal"

type SIC0000113I struct {
	LoanReceiptNum string `json:"loanReceiptNum" description:"loan receipt num"`
	ApplySerialNum string `json:"applySerialNum" description:"apply serial num"`
}

type SIC0000113O struct {
	ApplySerialNum       string                  `json:"applySerialNum" description:"apply serial num"`
	LoanReceiptNum       string                  `json:"loanReceiptNum" description:"loan receipt num"`
	ApplyDate            string                  `json:"applyDate" description:"apply date"`
	ApplyTime            string                  `json:"applyTime" description:"apply time"`
	ModifyType           string                  `json:"modifyType"`
	CurrentPeriod        int                     `json:"currentPeriod"`
	InstallmentAmount    decimal.Decimal         `json:"installmentAmount"`
	ReduceInstallment    decimal.Decimal         `json:"reduceInstallment"`
	SkipPeriodCount      int                     `json:"skipPeriodCount"`
	ReducePercentage     decimal.Decimal         `json:"reducePercentage"`
	ReducePeriodCount    int                     `json:"reducePeriodCount"`
	Checker              string                  `json:"checker"`
	ApproveComment       string                  `json:"approveComment"`
	ApproveView          string                  `json:"approveView"`
	MakerComment         string                  `json:"makerComment"`
	Maker                string                  `json:"maker"`
	CurrentAccountStatus string                  `json:"currentAccountStatus"`
	FinalModifyDate      string                  `json:"finalModifyDate"`
	FinalModifyTime      string                  `json:"finalModifyTime"`
	ListAmountType       []SIC0000113OAmountType `json:"listAmountType"`
}

type SIC0000113OAmountType struct {
	SeqNum            int             `json:"seqNum"`
	AmountType        string          `json:"amountType"`
	Amount            decimal.Decimal `json:"amount"`
	WaiveAmount       decimal.Decimal `json:"waiveAmount"`
	ConvertAmount     decimal.Decimal `json:"convertAmount"`
	ConvertAmountType string          `json:"convertAmountType"`
}

func (*SIC0000113I) GetServiceKey() string {
	return "IC000113"
}
