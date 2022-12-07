package models

import (
	"github.com/shopspring/decimal"
)

type SAL0000004I struct {
	AccountingAccountNum string `json:"accountingAccountNum" description:"Accounting account number" validate:"required"`
}

type SAL0000004O struct {
	AccountingAccountStatus string          `json:"accountingAccountStatus" description:"Accounting account status:0:normal 1:overdue 2:impaired 3:write_off 4:closed"`
	Balance                 decimal.Decimal `json:"balance" description:"Balance"`
	CurrentPeriod           int             `json:"currentPeriod" description:"Current period"`
	DrawTotalAmount         decimal.Decimal `json:"drawTotalAmount" description:"Draw total amount"`
	DrawTotalPeriod         int             `json:"drawTotalPeriod" description:"Draw total period"`
	LoanAmount              decimal.Decimal `json:"loanAmount" description:"Loan amount"`
	LoanReceiptNum          string          `json:"loanReceiptNum" description:"Loan receipt number"`
	NextRepaymentDate       string          `json:"nextRepaymentDate" description:"Next repayment date"`
	CurrentUnpaidPrincipal  decimal.Decimal `json:"currentUnpaidPrincipal" description:"Current unpaid principal"`
	AccumulateAccrualAmount decimal.Decimal `json:"accumulateAccrualAmount" description:"Accumulate accrual amount"`
	FixedInterestRate       decimal.Decimal `json:"fixedInterestRate" description:"Fixed interest rate"`
	PlanRepayInterest       decimal.Decimal `json:"planRepayInterest" description:"Plan repay interest"`   //add mvp2-sprint1-LOAN1732
	PlanRepayPrincipal      decimal.Decimal `json:"planRepayPrincipal" description:"Plan repay principal"` //add mvp2-sprint1-LOAN1732
}

func (*SAL0000004I) GetServiceKey() string {
	return "sal0000004"
}
