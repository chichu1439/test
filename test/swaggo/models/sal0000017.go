package models

import "github.com/shopspring/decimal"

type AL000017I struct {
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan receipt number"`
}

type AL000017O struct {
	ListLoanInterestType   []GetLoanInterestType    `json:"listLoanInterestType" description:"Loan interest type list"`
	BillRuleInfo           *BillRuleInfo            `json:"billRuleInfo" description:"Bill rule info"`
	AcctRepaymentHierarchy []AcctRepaymentHierarchy `json:"acctRepaymentHierarchy"`
	ListTLoanFeeType       []TLoanFeeType           `json:"listTLoanFeeType" description:"Loan fee type list"`
}

type GetLoanInterestType struct {
	LoanReceiptNum   string          `json:"loanReceiptNum" description:"Loan receipt number"`
	InterestId       string          `json:"interestId" description:"Interest id"`
	BaseInterestRate decimal.Decimal `json:"baseInterestRate" description:"Base interest rate"`
	InterestRate     decimal.Decimal `json:"interestRate" description:"Interest rate"`
	InterestType     string          `json:"interestType" description:"Interest type"`
	FloatDirection   string          `json:"floatDirection" description:"Float direction"`
	FloatExecType    string          `json:"floatExecType" description:"Float exec type"`
	FloatMethod      string          `json:"floatMethod" description:"Float method"`
	FloatValue       decimal.Decimal `json:"floatValue" description:"Float value"`
	StrategyId       string          `json:"strategyId" description:"Strategy id"`
	YearDays         int             `json:"yearDays" description:" days of year(0-360 1-365 2-Actual days)"`
	MonthDays        int             `json:"monthDays"`
	CustomerGrade    string          `json:"customerGrade" description:"Customer grade"`
}

func (*AL000017I) GetServiceKey() string {
	return "AL000017"
}
