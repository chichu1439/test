package models

import "github.com/shopspring/decimal"

type AL000016I struct {
	LoanReceiptNum       string                   `json:"loanReceiptNum" description:"Loan receipt number"`
	TransDate            string                   `json:"transDate" description:"Transaction date" `
	ListLoanInterestType []UpdateLoanInterestType `json:"listLoanInterestType" description:"List loan interest type"`
}

type UpdateLoanInterestType struct {
	InterestId       string          `json:"interestId" description:"Interest id"`
	BaseInterestRate decimal.Decimal `json:"baseInterestRate" description:"Base interest rate"`
	InterestRate     decimal.Decimal `json:"interestRate" description:"Interest rate"`
	InterestType     string          `json:"interestType" description:"Interest type"`
	FloatDirection   string          `json:"floatDirection" description:"Float direction"`
	FloatMethod      string          `json:"floatMethod" description:"Float method"`
	FloatValue       decimal.Decimal `json:"floatValue" description:"Float value"`
	FinalModifyOrgz  string          `json:"finalModifyOrgz" description:"Final modify organization"`
	FinalModifyUser  string          `json:"finalModifyUser" description:"Final modify user"`
}

type AL000016O struct {
}

func (*AL000016I) GetServiceKey() string {
	return "AL000016"
}
