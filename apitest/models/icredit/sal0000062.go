package models

import "github.com/shopspring/decimal"

type AL000062I struct {
	AccountingAccountNum       string          `json:"accountingAccountNum" description:"Accounting account number" validate:"required"`
	LoanReceiptNum             string          `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	CustomerId                 string          `json:"customerId" description:"Customer id" validate:"required"`
	CustomerName               string          `json:"customerName" description:"Customer name" validate:"required"`
	ProductId                  string          `json:"productId" description:"Product id" validate:"required"`
	Currency                   string          `json:"currency" description:"Currency" validate:"required" `
	LoanAmount                 decimal.Decimal `json:"loanAmount" description:"Loan amount" validate:"required"`
	MatureDate                 string          `json:"matureDate" description:"Mature date" `
	LoanPeriod                 int             `json:"loanPeriod" description:"Loan period" validate:"required"`
	RepaymentMethod            string          `json:"repaymentMethod" description:"Repayment method" validate:"required"`
	RepaymentCycle             string          `json:"repaymentCycle" description:"Repayment cycle" `
	GracePeriod                int             `json:"gracePeriod" description:"Grace period"`
	AutoRepaymentFlag          string          `json:"autoRepaymentFlag" description:"Auto repayment flag"`
	OtherBankFlag              string          `json:"otherBankFlag" description:"Other bank flag"`
	RepaymentAccountBankNumber string          `json:"repaymentAccountBankNumber" description:"Repayment account bank number"`
	RepaymentAccountBankName   string          `json:"repaymentAccountBankName" description:"Repayment account bank name"`
	RepaymentAccountNumber     string          `json:"repaymentAccountNumber" description:"Repayment account number"`
	RepaymentAccountName       string          `json:"repaymentAccountName" description:"Repayment account name"`
	RepaymentAccountOrgz       string          `json:"repaymentAccountOrgz" description:"Repayment account organization"`
	ActualMakeLoanDate         string          `json:"actualMakeLoanDate" description:"Actual make loan date"`
	RepayDay                   string          `json:"repayDay" description:"Repay day"`
	OpenAccountDate            string          `json:"openAccountDate" description:"Open account date"`
	OpenAccountOrgz            string          `json:"openAccountOrgz" description:"Open account organization"`
	ImpairedDays               int             `json:"impairedDays" description:"Impaired days"`
	CompoundInterestFlag       string          `json:"compoundInterestFlag" description:"Compound interest flag"`
}

type AL000062O struct {
	Status string `json:"status" description:"Status"`
}

func (*AL000062I) GetServiceKey() string {
	return "AL000062"
}
