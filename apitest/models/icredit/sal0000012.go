package models

import "github.com/shopspring/decimal"

type AL000012I struct {
	LoanReceiptNum       string          `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	AccountingAccountNum string          `json:"accountingAccountNum" description:"Accounting account number" validate:"required"`
	WriteoffType         string          `json:"writeoffType" validate:"required" description:"Writeoff type:0-partial 1-total" `
	WriteoffAmount       decimal.Decimal `json:"writeoffAmount" description:"Writeoff amount" validate:"required"`
	ReasonCode           string          `json:"reasonCode" validate:"required" description:"Reason code:01-not repaid after repeated collections 02-borrower is bankrupt 03-borrower is missing 04-borrower is dead 06-other reason"`
	Description          string          `json:"description" description:"Description"`
	ProofFile            string          `json:"proofFile" description:"file url list[file1.jpg,file2.jpg…]"`
}

type AL000012O struct {
	RegId string `json:"regId" description:"Reg id"`
}

func (*AL000012I) GetServiceKey() string {
	return "AL000012"
}
