package models

import "github.com/shopspring/decimal"

type SIC0000022I struct {
	MakeLoanApplySerialNum string `json:"makeLoanApplySerialNum" description:"Make loan apply serial number" description："Make loan apply serial number"`
	LoanReceiptNum         string `json:"loanReceiptNum" description:"Loan receipt number"`
	CustomerId             string `json:"customerId" description:"Customer id"`
	PageNum                int    `json:"pageNum" description:"Page number"`
	PageRecordCount        int    `json:"pageRecordCount" description:"Page record count"`
}

type SIC0000022O struct {
	PageNum         int        `json:"pageNum" description:"Page number"`
	PageRecordCount int        `json:"pageRecordCount" description:"Page record count"`
	TotalCount      int        `json:"totalCount" description:"Total count"`
	Records         []LoanList `json:"records" description:"Records"`
}

type LoanList struct {
	MakeLoanApplySerialNum string          `json:"makeLoanApplySerialNum" description:"Make loan apply serial number"`
	LoanReceiptNum         string          `json:"loanReceiptNum" description:"Loan receipt number"`
	LoanProductId          string          `json:"loanProductId" description:"Loan product id"`
	CustomerId             string          `json:"customerId" description:"Customer id"`
	CustomerName           string          `json:"customerName" description:"Customer name"`
	IdType                 string          `json:"idType" description:"Id type"`
	IdNum                  string          `json:"idNum" description:"Id number"`
	LoanApplyStatus        string          `json:"loanApplyStatus" description:"Loan apply status"`
	ActualMakeLoanDate     string          `json:"actualMakeLoanDate" description:"Actual make loan date"`
	MatureDate             string          `json:"matureDate" description:"Mature date"`
	ContractAmount         decimal.Decimal `json:"contractAmount"`
	ProductType            string          `json:"productType"`
	ExecuteInterest        decimal.Decimal `json:"executeInterest" description:"Execute interest"`
	LoanCurrencyCode       string          `json:"loanCurrencyCode" description:"Loan currency code"`
	LoanProductName        string          `json:"loanProductName" description:"Loan product name"`
	DrawDate               string          `json:"drawDate"`
}

func (*SIC0000022I) GetServiceKey() string {
	return "IC000022"
}
