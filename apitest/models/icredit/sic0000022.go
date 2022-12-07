package models

import "github.com/shopspring/decimal"

type SIC0000022I struct {
	// LoanApplicationStartDate string   `json:"loanApplicationStartDate" description:"Loan application start date"`
	// LoanApplicationEndDate   string   `json:"loanApplicationEndDate" description:"Loan application end date"`
	MakeLoanApplySerialNum string `json:"makeLoanApplySerialNum" description:"Make loan apply serial number" description："Make loan apply serial number"`
	// LoanReceiptNum           string   `json:"loanReceiptNum" description:"Loan receipt number"`
	// CustomerName             string   `json:"customerName" description:"Customer name"`
	// IdType                   string   `json:"idType" description:"Id type"`
	// IdNum                    string   `json:"idNum" description:"Id number"`
	// CustomerId               string   `json:"customerId" description:"Customer id"`
	// LoanStatus               []string `json:"loanStatus" description:"Loan status"`
	// CurrentDateType          string   `json:"currentDateType" description:"Current date type"` //1 today;2 weekly;3 monthly
	// PageNum                  int      `json:"pageNum" description:"Page number"`
	// PageRecordCount          int      `json:"pageRecordCount" description:"Page record count"`
	CreditApplySerialNum string `json:"creditApplySerialNum" description:"Credit apply serial number"`
	LoanReceiptNum       string `json:"loanReceiptNum" description:"Loan receipt number"`
	CustomerId           string `json:"customerId" description:"Customer id"`
	PageNum              int    `json:"pageNum" description:"Page number"`
	PageRecordCount      int    `json:"pageRecordCount" description:"Page record count"`
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

func (*SIC0000022I) GetServiceKey() string {
	return "sic0000022"
}
