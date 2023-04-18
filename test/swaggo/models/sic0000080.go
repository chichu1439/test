//Version: v1
package models

import (
	"github.com/shopspring/decimal"
)

type IC000080Request struct {
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	CustomerId     string `json:"customerId" description:"Customer id" validate:"required"`
	TransDate      string `json:"transDate" description:"Trans Date" validate:"required"`
	TransType      string `json:"transType" `
}

type IC000080Response struct {
	PendingApproval           string          `json:"pendingApproval" description:"Exist pending disbursement flage"`
	Maker                     string          `json:"maker"`
	ApplyStatus               string          `json:"applyStatus" description:"Apply Status"`
	MakerComment              string          `json:"makerComment" description:"Maker Comment"`
	ApplyDate                 string          `json:"applyDate" description:"Apply Date"`
	ApplySerialNum            string          `json:"applySerialNum"`
	TransType                 string          `json:"transType" `
	TransAmount               decimal.Decimal `json:"transAmount"`
	Balance                   decimal.Decimal `json:"balance"`
	PrincipleAmount           decimal.Decimal `json:"principleAmount" description:"Principle Amount"`
	CounterpartyLocalBankFlag string          `json:"counterpartyLocalBankFlag"`
	CounterpartyBankNum       string          `json:"counterpartyBankNum"`
	CounterpartyBankName      string          `json:"counterpartyBankName"`
	CounterpartyAccountNum    string          `json:"counterpartyAccountNum"`
	CounterpartyAccountName   string          `json:"counterpartyAccountName"`
	CounterpartyTransDate     string          `json:"counterpartyTransDate"`
	CounterpartySerialNum     string          `json:"counterpartySerialNum"`
	FrontEndFee               decimal.Decimal `json:"frontEndFee"`
	StampDutyFee              decimal.Decimal `json:"stampDutyFee"`
}

func (*IC000080Request) GetServiceKey() string {
	return "IC000080"
}
