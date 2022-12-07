package models

type IC000027I struct {
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan receipt number"` //
}

type IC000027O struct {
	Records []IC000027OData `json:"records" description:"Records"`
}

type IC000027OData struct {
	LoanReceiptNum       string `json:"loanReceiptNum" description:"Loan receipt number"` // ",
	RecalculateDate      string `json:"recalculateDate" description:"Recalculation date"` // 重算日期 ",
	LoanPeriod           string `json:"loanPeriod" description:"Loan period"`             // ",
	CustomerId           string `json:"customerId" description:"Customer id"`             // ",
	RepayDate            string `json:"repayDate" description:"Repay date"`               // ",
	Installment          string `json:"installment" description:"Installment"`
	RepayPrincipalAmount string `json:"repayPrincipalAmount" description:"Repay principal amount"` // ",
	RepayInterestAmount  string `json:"repayInterestAmount" description:"Repay interest amount"`   // ",
	CurrentBeginDate     string `json:"currentBeginDate" description:"Current begin date"`         // ",
	CurrentEndDate       string `json:"currentEndDate" description:"Current end date"`             // ",
	RemainAmount         string `json:"remainAmount" description:"Remain amount"`                  // ",
	CurrentAccountStatus string `json:"currentAccountStatus" description:"Current account status"`
	RepayStatus          string `json:"repayStatus" description:"Repay status:0-outstanding 1-repaid"` // 0-未还 1-已还 ",
}

func (*IC000027I) GetServiceKey() string {
	return "IC000027"
}
