package models

import "github.com/shopspring/decimal"

type IC000055I struct {
	LoanReceiptNum  string `json:"loanReceiptNum" description:"Loan receipt number"`
	CustomerId      string `json:"customerId" description:"Customer id"`
	PageNum         int    `json:"pageNum" description:"Page number"`
	PageRecordCount int    `json:"pageRecordCount" description:"Page record count"`
}

type IC000055O struct {
	PageNum         int                `json:"pageNum" description:"Page number"`
	PageRecordCount int                `json:"pageRecordCount" description:"Page record count"`
	TotalCount      int                `json:"totalCount" description:"Total count"`
	Records         []LoanWriteOffList `json:"records" description:"Records"`
}

type LoanWriteOffList struct {
	AccountingAccountNum   string          `json:"accountingAccountNum" description:"Accounting account number"`
	MakeLoanApplySerialNum string          `json:"makeLoanApplySerialNum" description:"Make loan apply serial number"`
	LoanReceiptNum         string          `json:"loanReceiptNum" description:"Loan receipt number"`
	LoanProductId          string          `json:"loanProductId" description:"Loan product id"`
	CustomerId             string          `json:"customerId" description:"Customer id"`
	CustomerName           string          `json:"customerName" description:"Customer name"`
	IdType                 string          `json:"idType" description:"Id type"`
	IdNum                  string          `json:"idNum" description:"Id number"`
	ApplicationDate        string          `json:"applicationDate" description:"Application date"`
	LoanApplyStatus        string          `json:"loanApplyStatus" description:"Loan apply status"`
	ActualMakeLoanDate     string          `json:"actualMakeLoanDate" description:"Actual make loan date"`
	MatureDate             string          `json:"matureDate" description:"Mature date"`
	LoanAmount             decimal.Decimal `json:"loanAmount" description:"Loan amount"`
	ExecuteInterest        decimal.Decimal `json:"executeInterest" description:"Execute interest"`
	LoanChangeStatus       string          `json:"loanChangeStatus" description:"Loan change status"`
	LoanCurrencyCode       string          `json:"loanCurrencyCode" description:"Loan currency code"`
	LoanProductName        string          `json:"loanProductName" description:"Loan product name"`
}

func (*IC000055I) GetServiceKey() string {
	return "IC000055"
}
