//Version: v1
package models

import "github.com/shopspring/decimal"

type BP000004Request struct {
	CustomerId      string `json:"customerId" description:"Customer ID" validate:"required,max=20"`
	BnplProviderId  string `json:"bnplProvider" description:"BNPL Provider ID" validate:"required"`
	LoanProductId   string `json:"loanProductId" description:"Loan Product ID" validate:"required"`
	PageNum         int    `json:"pageNum" description:"Page Number"`
	PageRecordCount int    `json:"pageRecordCount" description:"Page Record Count"`
}

type BP000004Response struct {
	PageNum          int                `json:"pageNum" description:"Page Number"`
	PageRecordCount  int                `json:"pageRecordCount" description:"Page Record Count"`
	TotalCount       int                `json:"totalCount" description:"Total Count"`
	CustomerID       string             `json:"customerId" description:"Customer ID"`
	CustomerName     string             `json:"customerName" description:"Customer Name"`
	LoanCustomerId   string             `json:"loanCustomerId" description:"Loan Customer ID"`
	LoanContractList []LoanContractInfo `json:"loanContractList" description:"Loan Contract List"`
}

type LoanContractInfo struct {
	LoanReceiptNum  string          `json:"loanReceiptNum" description:"Loan Receipt Number"`
	LoanStatus      string          `json:"loanStatus" description:"Loan Status"`
	Currency        string          `json:"currency" description:"Currency"`
	LoanAmount      decimal.Decimal `json:"loanAmount" description:"Loan Amount"`
	OpenAccountDate string          `json:"openAccountDate" description:"Open Account Date"`
	TotalPeriod     int             `json:"totalPeriod" description:"Total Period"`
	PaidPeriod      int             `json:"paidPeriod" description:"Paid Period"`
	OrderId         string          `json:"orderId" description:"Order ID"`
	MerchantId      string          `json:"merchantId" description:"Merchant ID"`
	MerchantName    string          `json:"merchantName" description:"Merchant Name"`
}

func (*BP000004Request) GetServiceKey() string {
	return "BP000004"
}
