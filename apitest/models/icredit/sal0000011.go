package models

import "github.com/shopspring/decimal"

type AL000011I struct {
	LoanReceiptNum  string `json:"loanReceiptNum" description:"Loan receipt number"`
	CustomerId      string `json:"customerId" description:"Customer id"`
	PageNum         int    `json:"pageNum" description:"Page number"`
	PageRecordCount int    `json:"pageRecordCount" description:"Page record count"`
}

type AL000011O struct {
	PageNum         int                       `json:"pageNum" description:"Page number"`
	PageRecordCount int                       `json:"pageRecordCount" description:"Page record count"`
	TotalCount      int                       `json:"totalCount" description:"Total count"`
	Records         []AL000011OTLoanAcctgAcct `json:"records" description:"Records"`
}

type AL000011OTLoanAcctgAcct struct {
	AccountingAccountNum    string          `json:"accountingAccountNum" description:"Accounting account number"` //
	LoanReceiptNum          string          `json:"loanReceiptNum" description:"Loan receipt number"`             //
	CustomerId              string          `json:"customerId" description:"Customer id"`
	LoanAmount              decimal.Decimal `json:"loanAmount" description:"Loan amount"`
	Balance                 decimal.Decimal `json:"balance" description:"Balance"`
	DrawTotalAmount         decimal.Decimal `json:"drawTotalAmount" description:"Draw total amount"`
	DrawBeginDate           string          `json:"drawBeginDate" description:"Draw begin date"`
	DrawTotalPeriod         int             `json:"drawTotalPeriod" description:"Draw total period"`
	MaturityDate            string          `json:"maturityDate" description:"Maturity date"`
	AccountingAccountStatus string          `json:"accountingAccountStatus" description:"Accounting account status:0:normal 1:overdue 2:impaired 3:write_off 4:closed"` // 0-正常;1-逾期;2-减值;3-核销;4-结清
}

func (*AL000011I) GetServiceKey() string {
	return "AL000011"
}
