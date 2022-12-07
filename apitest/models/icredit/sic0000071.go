package models

import (
	"github.com/shopspring/decimal"
)

type IC000071Request struct {
	LoanAmount  decimal.Decimal `json:"loanAmount" description:"Loan amount" validate:"required"`
	PeriodNum   int             `json:"periodNum" description:"Period number" validate:"required"`
	CurrentDate string          `json:"currentDate" description:"Current date" validate:"required,datetime=2006-01-02"`
	ProductId   string          `json:"productId" description:"Product ID" validate:"required"`
}

type IC000071Response struct {
	RepayMethod                string            `json:"repayMethod"  description:"Repayment method"`
	InterestCalculateStartDate string            `json:"InterestCalculateStartDate" description:"Interest calculate start date"`
	InterestCalculateEndDate   string            `json:"interestCalculateEndDate" description:"Interest calculate end date"`
	PeriodNum                  int               `json:"periodNum" description:"Period number"`
	LoanAmount                 decimal.Decimal   `json:"loanAmount" description:"Loan amount"`
	PlanRepayTotalAmount       decimal.Decimal   `json:"planRepayTotalAmount" description:"Plan repayment total amount"`
	PlanRepayTotalInterest     decimal.Decimal   `json:"planRepayTotalInterest" description:"Plan repayment total interest"`
	InterestRate               decimal.Decimal   `json:"interestRate" description:"Interest rate"`
	Records                    []RepayPlanRecord `json:"records" description:"records"`
}

func (*IC000071Request) GetServiceKey() string {
	return "IC000071"
}
