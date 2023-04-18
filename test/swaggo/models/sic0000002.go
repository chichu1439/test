package models

import "github.com/shopspring/decimal"

type IC000002I struct {
	LoanReceiptNum string          `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	ApplySerialNum string          `json:"applySerialNum" description:"Apply serial number" validate:"required"`
	ApproveComment string          `json:"approveComment" description:"Checker comment"`
	TransAmount    decimal.Decimal `json:"transAmount" description:"Trans amount" validate:"required,gt=0"`
	ApproveView    string          `json:"approveView" description:"Apply status" validate:"required,lte=3,gt=0"`
}

type IC000002O struct {
	Status         string `json:"status" description:"Status"`
	ApplySerialNum string `json:"applySerialNum" description:"Apply serial number"`
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan receipt number" `
}

func (*IC000002I) GetServiceKey() string {
	return "IC000002"
}
