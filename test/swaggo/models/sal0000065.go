package models

import "github.com/shopspring/decimal"

type AL000065I struct {
	OrderId        string          `json:"orderId" description:"Order id" validate:"required"`
	LinkId         string          `json:"linkId" description:"Link id" validate:"required"`
	LoanReceiptNum string          `json:"loanReceiptNum" description:"Loan receipt number" validate:"required" `
	Currency       string          `json:"currency" description:"Currency" validate:"required"`
	Amount         decimal.Decimal `json:"amount" description:"Amount" validate:"required"`
	InvoiceNo      string          `json:"invoiceNo" description:"Invoice number"`
	CardNum        string          `json:"cardNum" description:"Card number"`
	TransChannel   string          `json:"transChannel" description:"Trans channel:2c2p"`
	PaymentErrCode string          `json:"paymentErrCode" description:"Payment error code:F-Failure S-success" `
	PaymentErrMsg  string          `json:"paymentErrMsg" description:"Payment err message" `
	PaymentStatus  string          `json:"paymentStatus" description:"Payment status"`
}

type AL000065O struct {
	Status string `json:"status" description:"Status"`
}

func (*AL000065I) GetServiceKey() string {
	return "AL000065"
}
