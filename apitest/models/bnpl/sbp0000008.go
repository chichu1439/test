//Version: v1
package models

import "github.com/shopspring/decimal"

type BP000008Request struct {
	CustomerId      string          `json:"customerId" description:"Customer ID" validate:"required,max=20"`
	CustomerName    string          `json:"customerName" description:"Customer Name" validate:"required,max=320"`
	LoanCustomerId  string          `json:"loanCustomerId" description:"Loan Customer ID" validate:"required,max=20"`
	BnplProviderId  string          `json:"bnplProviderId" description:"BNPL Provider ID (Source Channel)" validate:"required,max=20"`
	TotalAmount     decimal.Decimal `json:"totalAmount" description:"Total Amount" validate:"required"`
	Currency        string          `json:"currency" description:"Currency: USD, GBP, EUR, CNY, THB, etc.." validate:"required"`
	LoanReceiptNum  string          `json:"loanReceiptNum" description:"Loan Receipt Number"`
	RepaymentType   string          `json:"repaymentType" description:"Repayment Type" validate:"required"`
	BindAcctNo      string          `json:"bindAcctNo" description:"Bind Account Number" validate:"required"`
	BindAcctType    string          `json:"bindAcctType" description:"Bind Account Type" validate:"required"`
	BindAcctOweBank string          `json:"bindAcctOweBank" description:"Bind Account Owe Bank" validate:"required"`
}

type BP000008Response struct {
	Status string `json:"status" description:"Status"`
}

func (*BP000008Request) GetServiceKey() string {
	return "BP000008"
}
