package models

import "github.com/shopspring/decimal"

type AL000071Request struct {
	AccountingAccountNum string          `json:"accountingAccountNum" description:"Accounting account number" validate:"required"`
	LoanReceiptNum       string          `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	TransAmount          decimal.Decimal `json:"transAmount" description:"Trans amount" validate:"required"`
	TransDate            string          `json:"transDate" description:"Trans date"`
	RepaymentStatus      string          `json:"repaymentStatus" description:"01-total 02-Partial 03-by bill 04-down payment 05-overdue payment" validate:"required"`
}

type AL000071Response struct {
	Balance         decimal.Decimal `json:"balance" description:"Balance"`
	DrawTotalAmount decimal.Decimal `json:"drawTotalAmount" description:"Draw total amount"`
}

func (*AL000071Request) GetServiceKey() string {
	return "AL000071"
}
