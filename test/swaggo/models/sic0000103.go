package models

import "github.com/shopspring/decimal"

type SIC0000103I struct {
	RestructuringApplySerialNum string `json:"restructuringApplySerialNum"`
	LoanReceiptNum              string `json:"loanReceiptNum"`
}
type SIC0000103IListInterest struct {
	InterestType string          `json:"interestType"`
	InterestRate decimal.Decimal `json:"interestRate"`
}
type SIC0000103O struct {
	ContractDetail    ContractDetail    `json:"ContractDetail"`
	CustomerInfo      CustomerInfo      `json:"customerInfo"`
	ProductInfo       ProductInfo       `json:"productInfo"`
	AccountInfo       AccountInfo       `json:"accountInfo"`
	RepaymentPlanList []SIC0000103OList `json:"repaymentPlanList"`
}
type SIC0000103OList struct {
	AccountingAccountNum   string
	Period                 int
	CurrentAccountStatus   string
	PeriodRepayStatus      string
	PrincipalRepayStatus   string
	FeeRepayStatus         string
	CurrentPeriodBeginDate string
	CurrentPeriodEndDate   string
	CurrentRepayDate       string
	Currency               string
	PlanRepayPrincipal     decimal.Decimal
	PlanRepayInterest      decimal.Decimal
	ActualRepayDate        string
	ActualRepayPrincipal   decimal.Decimal
	ActualRepayInterest    decimal.Decimal
	InterestWaived         decimal.Decimal
	PrincipalWaived        decimal.Decimal
	PrincipalImpaired      decimal.Decimal
	PrincipalWriteoff      decimal.Decimal
	FeeAmount              decimal.Decimal
	FeeRepay               decimal.Decimal
	FeeWaived              decimal.Decimal
	OverdueStartDate       string
	OverdueDays            int
	CloseDate              string
	CurrentUnpaidPrincipal decimal.Decimal
	MaintainPrinciple      decimal.Decimal
	LastActiveDate         string
	InterestWriteOff       decimal.Decimal
}

func (*SIC0000103I) GetServiceKey() string {
	return "IC000103"
}
