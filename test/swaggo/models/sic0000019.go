package models

import "github.com/shopspring/decimal"

type SIC0000019File struct {
	SerialNum                  string          `json:"serialNum"`
	ContractSuID               string          `json:"contractSuId"`
	LoanReceiptNum             string          `json:"loanReceiptNum"`
	CustomerId                 string          `json:"customerId"`
	CustomerName               string          `json:"customerName"`
	ProductId                  string          `json:"productId"`
	Currency                   string          `json:"currency"`
	LoanAmount                 decimal.Decimal `json:"loanAmount"`
	MatureDate                 string          `json:"matureDate"`
	LoanPeriod                 int             `json:"loanPeriod"`
	RepaymentMethod            string          `json:"repaymentMethod"`
	RepaymentCycle             string          `json:"repaymentCycle"`
	RepaymentFrequency         string          `json:"repaymentFrequency"`
	GracePeriod                int             `json:"gracePeriod"`
	InterestAccrualFlag        string          `json:"interestAccrualFlag"`
	AutoRepaymentFlag          string          `json:"autoRepaymentFlag"`
	OtherBankFlag              string          `json:"otherBankFlag"`
	RepaymentAccountBankNumber string          `json:"repaymentAccountBankNumber"`
	RepaymentAccountBankName   string          `json:"repaymentAccountBankName"`
	RepaymentAccountNumber     string          `json:"repaymentAccountNumber"`
	RepaymentAccountName       string          `json:"repaymentAccountName"`
	RepaymentAccountOrgz       string          `json:"repaymentAccountOrgz"`
	AccountingAccountNum       string          `json:"accountingAccountNum"`
	ActualMakeLoanDate         string          `json:"actualMakeLoanDate"`
	RepayDay                   string          `json:"repayDay"`
	CompoundInterestFlag       string          `json:"compoundInterestFlag"`
	ImpairedDays               int             `json:"impairedDays"`
}
