//Version: v0.0.1
package models

import "github.com/shopspring/decimal"

type AL000067I struct {
	AccountingAccountNum string          `json:"accountingAccountNum" description:"Accounting account number" validate:"required"`
	LoanReceiptNum       string          `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	BusinessSerialNum    string          `json:"businessSerialNum" description:"Business serial number" validate:"required"`
	CustomerId           string          `json:"customerId" description:"Customer id" validate:"required"`
	TransDate            string          `json:"transDate" description:"Trans date" validate:"required"`
	TransAmount          decimal.Decimal `json:"transAmount" description:"Trans Amount" validate:"required"`
	TransAbstract        string          `json:"transAbstract"description:"Trans abstract" `
}

type AL000067O struct {
	Balance decimal.Decimal `json:"balance" description:"balance"`
}

func (*AL000067I) GetServiceKey() string {
	return "AL000067"
}
