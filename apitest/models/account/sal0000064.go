package models

type AL000064I struct {
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan receipt number"` //
}

type AL000064O struct {
	Records []AL000064OData `json:"records" description:"Records"`
}

type AL000064OData struct {
	LoanReceiptNum       string `json:"loanReceiptNum" description:"Loan receipt number"` // ",
	RecalculateDate      string `json:"recalculateDate" description:"Recalculate date"`   // 重算日期 ",
	LoanPeriod           string `json:"loanPeriod" description:"Loan period"`             // ",
	CustomerId           string `json:"customerId" description:"Customer id"`             // ",
	RepayDate            string `json:"repayDate" description:"Repay date"`               // ",
	Installment          string `json:"installment" description:"Installment"`
	RepayPrincipalAmount string `json:"repayPrincipalAmount" description:"Repay principal amount"` // ",
	RepayInterestAmount  string `json:"repayInterestAmount" description:"Repay interest amount"`   // ",
	CurrentBeginDate     string `json:"currentBeginDate" description:"Current begin date"`         // ",
	CurrentEndDate       string `json:"currentEndDate" description:"Current end date"`             // ",
	RemainAmount         string `json:"remainAmount" description:"Remain amount"`                  // ",
	RepayStatus          string `json:"repayStatus" description:"Repay status"`                    // 0-未还 1-已还 ",
}

func (*AL000064I) GetServiceKey() string {
	return "AL000064"
}
