//Version: v1
package models

type IC000087Request struct {
	LoanReceiptNum  string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	StartDate       string `json:"startDate" description:"Start date"`
	EndDate         string `json:"endDate" description:"Start date"`
	PageNum         int    `json:"pageNum"`
	PageRecordCount int    `json:"pageRecordCount"`
}

type IC000087Response struct {
	PageNum         int              `json:"pageNum" description:"Page number"`
	PageRecordCount int              `json:"pageRecordCount" description:"Page record count"`
	TotalCount      int              `json:"totalCount" description:"Total count"`
	LoanReceiptNum  string           `json:"loanReceiptNum" description:"Loan receipt number"`
	CustomerId      string           `json:"customerId" description:"Customer id"`
	Records         []RegisterRecord `json:"records" description:"Records"`
}

func (*IC000087Request) GetServiceKey() string {
	return "IC000087"
}
