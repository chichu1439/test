//Version: v1
package models

type BP000006Request struct {
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan Receipt Number" validate:"required"`
}

type BP000006Response struct {
	RepaymentPlanList []BP000006RepaymentPlan `json:"repaymentPlanList" description:"Repayment Plan List"`
}

type BP000006RepaymentPlan struct {
	LoanReceiptNum       string `json:"loanReceiptNum" description:"Loan Receipt Number"`
	RecalculateDate      string `json:"recalculateDate" description:"Recalculate Date"`
	LoanPeriod           string `json:"loanPeriod" description:"Loan Period"`
	CustomerId           string `json:"customerId" description:"Customer Id"`
	RepayDate            string `json:"repayDate" description:"Repay Date"`
	Installment          string `json:"installment" description:"Installment"`
	RepayPrincipalAmount string `json:"repayPrincipalAmount" description:"Repay Principal Amount"`
	RepayInterestAmount  string `json:"repayInterestAmount" description:"Repay Interest Amount"`
	CurrentBeginDate     string `json:"currentBeginDate" description:"Current Begin Date"`
	CurrentEndDate       string `json:"currentEndDate" description:"Current End Date"`
	RemainAmount         string `json:"remainAmount" description:"Remain Amount"`
	CurrentAccountStatus string `json:"currentAccountStatus" description:"Current Account Status"`
	RepayStatus          string `json:"repayStatus" description:"Repay Status"`
}

func (*BP000006Request) GetServiceKey() string {
	return "BP000006"
}
