//Version: v1
package models

import "github.com/shopspring/decimal"

type AL000068Request struct {
	AccountingAccountNum string `json:"accountingAccountNum" validate:"required"`
	LoanReceiptNum       string `json:"loanReceiptNum" validate:"required"`
	TransDate            string `json:"transDate" validate:"required"`
	NextBillDate         string `json:"nextBillDate"`
	RepaymentDay         string `json:"repaymentDay"`
	BillDay              string `json:"billDay"`
	//RepayDay              string             `json:"repayDay"`
	SameDayTreatment      string          `json:"sameDayTreatment"`
	BillCycle             string          `json:"billCycle"`
	BillRepaymentType     string          `json:"billRepaymentType"`
	PercentageOfPrincipal decimal.Decimal `json:"percentageOfPrincipal"`
	FixMinimumAmount      decimal.Decimal `json:"fixMinimumAmount"`
	FinalModifyDate       string          `json:"finalModifyDate"`
	FinalModifyTime       string          `json:"finalModifyTime"`
	FinalModifyOrgz       string          `json:"finalModifyOrgz"`
	FinalModifyUser       string          `json:"finalModifyUser"`
	BankAccountName       string          `json:"bankAccountName" description:"bank account name" `
	BankAccountNumber     string          `json:"bankAccountNumber" description:"bank account number"`
	BankName              string          `json:"bankName" description:"bank name"`
	Bic                   string          `json:"bic" description:"bic"`
	AccountChangeFlage    string          `json:"accountChangeFlage"`
}

type AL000068Response struct {
}

func (*AL000068Request) GetServiceKey() string {
	return "AL000068"
}
