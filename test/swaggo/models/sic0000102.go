package models

import "github.com/shopspring/decimal"

type SIC0000102I struct {
	CustomerId        string                    `json:"customerId"`
	CustomerName      string                    `json:"customerName"`
	IdType            string                    `json:"idType"`
	IdNum             string                    `json:"idNum"`
	ProductName       string                    `json:"productName"`
	LoanReceiptNum    string                    `json:"loanReceiptNum"`
	LoanPurpose       string                    `json:"loanPurpose"`
	Currency          string                    `json:"currency"`
	RepayMethod       string                    `json:"repayMethod"`
	RepayCycle        string                    `json:"repayCycle"`
	PeriodNum         string                    `json:"periodNum"`
	RepayDay          string                    `json:"repayDay"`
	DrawDate          string                    `json:"drawDate"`
	MatureDate        string                    `json:"matureDate"`
	ListInterest      []SIC0000102IListInterest `json:"listInterest"`
	BankAccountNumber string                    `json:"bankAccountNumber"`
	BankAccountName   string                    `json:"bankAccountName"`
	BankName          string                    `json:"bankName"`
	PrincipalAmount   decimal.Decimal           `json:"principalAmount"`
	CreditLine        decimal.Decimal           `json:"creditLine"`
	BillDay           string                    `json:"billDay"`
	ProductType       string                    `json:"productType"`
	PPI               decimal.Decimal           `json:"ppi"`
}
type SIC0000102IListInterest struct {
	InterestType string          `json:"interestType"`
	InterestRate decimal.Decimal `json:"interestRate"`
}
type SIC0000102O struct {
	Status string `json:"status"`
}

func (*SIC0000102I) GetServiceKey() string {
	return "IC000102"
}
