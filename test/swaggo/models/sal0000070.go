package models

import "github.com/shopspring/decimal"

type AL000070Request struct {
	SuType                     string             `json:"suType" description:"Su type"`
	LoanInterestTypeList       []LoanInterestType `json:"loanInterestTypeList" description:"Loan interest type list"`
	LoanReceiptNum             string             `json:"loanReceiptNum" description:"Loan receipt number"`
	AccountingAccountNum       string             `json:"accountingAccountNum" description:"Accounting account number"`
	CustomerId                 string             `json:"customerId" description:"Customer Id"`
	ProductId                  string             `json:"productId" description:"Product Id"`
	LoanMethod                 string             `json:"loanMethod" description:"Loan method"`
	RepayMethod                string             `json:"repayMethod" description:"Repay method"`
	LoanCycle                  string             `json:"loanCycle" description:"Loan cycle"`
	LoanPeriod                 int                `json:"loanPeriod" description:"Loan period"`
	OpenAccountDate            string             `json:"openAccountDate" description:"Open account date"`
	Currency                   string             `json:"currency" description:"Currency"`
	LoanAmount                 decimal.Decimal    `json:"loanAmount" description:"Loan amount"`
	MatureDate                 string             `json:"matureDate" description:"Mature date"`
	RevolvingFlag              string             `json:"revolvingFlag" description:"Revolving flag"`
	ImpairedDays               int                `json:"impairedDays" description:"Impaired days"`
	GracePeriod                int                `json:"gracePeriod" description:"Grace Period"`
	AutoRepaymentFlag          string             `json:"autoRepaymentFlag" description:"Auto Repayment Flag"`
	OtherBankFlag              string             `json:"otherBankFlag" description:"Other Bank Flag"`
	RepaymentAccountBankNumber string             `json:"repaymentAccountBankNumber" description:"Repayment Account BankNumber"`
	RepaymentAccountBankName   string             `json:"repaymentAccountBankName" description:"Repayment Account Bank Name"`
	RepaymentAccountNumber     string             `json:"repaymentAccountNumber" description:"Repayment Account Number"`
	RepaymentAccountName       string             `json:"repaymentAccountName" description:"Repayment Account Name"`
	CompoundInterestFlag       string             `json:"CompoundInterestFlag" description:"Compound Interest Flag"`
}

type AL000070Response struct {
	MatureDate        string          `json:"matureDate" description:"Mature date format yyyy-mm-dd"`
	RepaymentPlanList []RepaymentPlan `json:"repaymentPlanList" description:"Repayment plan"`
}

func (*AL000070Request) GetServiceKey() string {
	return "AL000070"
}
