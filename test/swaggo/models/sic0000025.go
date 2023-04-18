package models

import "github.com/shopspring/decimal"

type IC000025I struct {
	TradeDate      string          `json:"tradeDate" validate:"required" description:"Trade date"`
	LoanReceiptNum string          `json:"loanReceiptNum" validate:"required" description:"Loan receipt number"`
	TransAmount    decimal.Decimal `json:"transAmount" validate:"required" description:"Trans amount"`
	TransDescribe  string          `json:"transDescribe" description:"Trans describe:1.Repayment 2.Disbursement"`
	MakerComment   string          `json:"makerComment" description:"Maker comment"`
}

type IC000025O struct {
	Status         string `json:"status" description:"Status"`
	ApplySerialNum string `json:"applySerialNum" description:"Apply serial number"`
}

func (*IC000025I) GetServiceKey() string {
	return "IC000025"
}
