package models

import "github.com/shopspring/decimal"

type BP000003Request struct {
	CustomerId         string          `json:"customerId" description:"Customer Id" validate:"required"`
	OrderId            string          `json:"orderId" description:"Order Id" validate:"required"`
	PeriodNum          int             `json:"periodNum" description:"Period Number" validate:"required"`
	TotalAmount        decimal.Decimal `json:"totalAmount" description:"Total Amount" validate:"required"`
	MerchantId         string          `json:"merchantId" description:"Merchant Id" validate:"required"`
	MerchantName       string          `json:"merchantName" description:"Merchant name" validate:"required"`
	BindAcctNo         string          `json:"bindAcctNo" description:"Bind Account Number" validate:"required"`
	BindAcctТуре       string          `json:"bindAcctТуре" description:"Bind Account Туре" validate:"required"`
	BindAcctOweBank    string          `json:"bindAcctOweBank" description:"Bind Account Owe Bank" validate:"required"`
	BnplProviderId     string          `json:"bnplProviderId" description:"Bnpl Provider Id" validate:"required"`
	Currency           string          `json:"currency" description:"Currency" validate:"required"`
	LoanProductId      string          `json:"loanProductId" description:"Loan Product Id" validate:"required"`
	LoanPurpose        string          `json:"loanPurpose" description:"Loan Purpose" validate:"required"`
	RepaymentFrequency string          `json:"repaymentFrequency" description:"Repayment Frequency" validate:"required"`
	RepaymentMethod    string          `json:"repaymentMethod" description:"Repayment Method" validate:"required"`
	RepaymentType      string          `json:"repaymentType" description:"Repayment Type"`
}

type BP000003Response struct {
	Status         string `json:"status" description:"Status"`
	OrderId        string `json:"orderId" description:"Order Id"`
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan Receipt Number"`
	CustomerId     string `json:"customerId" description:"Customer Id"`
	CustomerName   string `json:"customerName" description:"Customer Name"`
	LoanCustomerId string `json:"loanCustomerId" description:"Loan Customer Id"`
}

func (*BP000003Request) GetServiceKey() string {
	return "BP000003"
}
