//Version: v1
package models

import "github.com/shopspring/decimal"

type IC000124Request struct {
	CustomerId         string          `json:"customerId" validate:"required"`
	LoanReceiptNum     string          `json:"loanReceiptNum" validate:"required"`
	DiscountPercentage decimal.Decimal `json:"discountPercentage" validate:"required"`
	BankName           string          `json:"bankName"`
	Bic                string          `json:"bic"`
	BankAccountNumber  string          `json:"bankAccountNumber"`
	BankAccountName    string          `json:"bankAccountName"`
	FeeList            []FeeInfo       `json:"feeList"`
	MakerComment       string          `json:"makerComment" validate:"required"`
}

type IC000124Response struct {
	ApplySerialNum string `json:"applySerialNum"`
	LoanReceiptNum string `json:"loanReceiptNum"`
}

func (*IC000124Request) GetServiceKey() string {
	return "IC000124"
}
