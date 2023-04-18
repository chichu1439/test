package models

import (
	"github.com/shopspring/decimal"
)

type AL000054I struct {
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
}

type AL000054O struct {
	AccountingAccountNum    string                   `json:"accountingAccountNum" description:"Accounting account number"`
	AccountingAccountStatus string                   `json:"accountingAccountStatus" description:"Accounting account status:0-Normal 1-Overdue 2-Impaired 3-Written Off 4-Closed"` // 0-正常;1-逾期;2-减值;3-核销;4-结清 ",
	CustomerId              string                   `json:"customerId" description:"Customer id"`
	ProductId               string                   `json:"productId" description:"Product id"`
	ListPeriod              []TAcctLoanAcctgPeriodBo `json:"listPeriod" description:"List period"`
}

type TAcctLoanAcctgPeriodBo struct {
	Period                 int             `json:"period" description:"Period"`
	CurrentAccountStatus   string          `json:"currentAccountStatus" description:"Current account status"`
	CurrentPeriodBeginDate string          `json:"currentPeriodBeginDate" description:"Current period begin date"`
	CurrentPeriodEndDate   string          `json:"currentPeriodEndDate" description:"Current period end date"`
	CurrentRepayDate       string          `json:"currentRepayDate" description:"Current repay date"`
	Currency               string          `json:"currency" description:"Currency"`
	PlanRepayPrincipal     decimal.Decimal `json:"planRepayPrincipal" description:"Plan repay principal"`
	ActualRepayDate        string          `json:"actualRepayDate" description:"Actual repay date"`
	ActualRepayPrincipal   decimal.Decimal `json:"actualRepayPrincipal" description:"Actual repay principal"`
	CurrentUnpaidPrincipal decimal.Decimal `json:"currentUnpaidPrincipal" description:"Current unpaid principal"`
	LastActiveDate         string          `json:"lastActiveDate" description:"Last active date"`
	PeriodRepayStatus      string          `json:"periodRepayStatus" description:"Period repay status"`
	PrincipalRepayStatus   string          `json:"principalRepayStatus" description:"Principal repay status"`
	FeeRepayStatus         string          `json:"feeRepayStatus" description:"Fee repay status"`
	PrincipalWaived        decimal.Decimal `json:"principalWaived" description:"Principal waived"`
	PrincipalImpaired      decimal.Decimal `json:"principalImpaired" description:"Principal impaired"`
	PrincipalWriteoff      decimal.Decimal `json:"principalWriteoff" description:"Principal writeoff"`
	ActualRepayInterest    decimal.Decimal `json:"actualRepayInterest" description:"Actual repay interest"`
	InterestWaived         decimal.Decimal `json:"interestWaived" description:"Interest waived"`
	FeeAmount              decimal.Decimal `json:"feeAmount" description:"Fee amount"`
	FeeRepay               decimal.Decimal `json:"feeRepay" description:"Fee repay"`
	FeeWaived              decimal.Decimal `json:"feeWaived" description:"Fee waived"`
	OverdueStartDate       string          `json:"overdueStartDate" description:"Overdue start date"`
	OverdueDays            int             `json:"overdueDays" description:"Overdue days"`
	CloseDate              string          `json:"closeDate" description:"Close date"`
	PlanRepayInterest      decimal.Decimal `json:"planRepayInterest" description:"Plan repay interest"`
	MaintainPrinciple      decimal.Decimal `json:"maintainPrinciple" description:"Maintain principle"`
}

func (*AL000054I) GetServiceKey() string {
	return "AL000054"
}
