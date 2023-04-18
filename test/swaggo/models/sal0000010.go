package models

import "github.com/shopspring/decimal"

type AL000010I struct {
	LoanReceiptNum       string             `json:"loanReceiptNum" description:"Loan receipt number"`
	SuId                 string             `json:"suId" description:"Su id"`
	ListLoanInterestType []LoanInterestType `json:"listLoanInterestType" description:"loan interest type:O-Normal Loan 1-Overdue Loan"`
	RepayMethod          string             `json:"repayMethod" validate:"required" description:"Repay method:01-equal principal 02-equal principal and interest 03-due (one-time) repayment of principal and interest 04-due repayment of principal and interest 05-equal principal and interest 06-repayment with borrowing"`
	BillRuleInfo         *BillRuleInfo      `json:"billRuleInof"`
	LoanFeeRuleList      []LoanFeeCalcRule  `json:"loanFeeRuleList"`
	LoanFeeTypeList      []LoanFeeType      `json:"loanFeeTypeList"`
}

type LoanInterestType struct {
	AccountingAccountNum string          `json:"accountingAccountNum" description:"Accounting account number"`
	CustomerGrade        string          `json:"customerGrade" description:"Customer grade"`
	AccrualPeriodDay     int             `json:"accrualPeriodDay" description:"Accrual period day"`
	AccrualPeriodType    string          `json:"accrualPeriodType" description:"Accrual period type"`
	AccrualPeriodValue   int             `json:"accrualPeriodValue" description:"Accrual period value"`
	CalculateType        string          `json:"calculateType" description:"Calculate type:0-Actual days 1-Year to month to day 2-Year to month"`
	DecimalFlag          string          `json:"decimalFlag" description:"Decimal flag:0-no 1-yes"`
	InterestId           string          `json:"interestId" description:"Interest id"`
	BaseInterestRate     decimal.Decimal `json:"baseInterestRate" description:"Base interest rate"`
	InterestRate         decimal.Decimal `json:"interestRate" description:"Interest rate"`
	InterestType         string          `json:"interestType" description:"Interest type"`
	FloatDirection       string          `json:"floatDirection" description:"Float direction"`
	FloatExecType        string          `json:"floatExecType" description:"Float exec type:0-By Value Date;1-By Calc Date;2-Fixed float"`
	FloatMethod          string          `json:"floatMethod" description:"Float method:0-By Percent 1-BP Value"`
	FloatValue           decimal.Decimal `json:"floatValue" description:"Float value"`
	MonthDays            string          `json:"monthDays" description:"Month days"`
	YearDays             string          `json:"yearDays" description:"Year days"`
	RateExecType         string          `json:"rateExecType" description:"Rate exec type:0-By Value Date;1-By Calc Date"`
	SettlePeriodDay      int             `json:"settlePeriodDay" description:"Settle period day"`
	SettlePeriodType     string          `json:"settlePeriodType" description:"Settle period type"`
	SettlePeriodValue    int             `json:"settlePeriodValue" description:"Settle period value"`
	Status               string          `json:"status" description:"Status"`
	StrategyId           string          `json:"strategyId" description:"Strategy id"`
	TaskOrgzNum          string          `json:"taskOrgzNum" description:"Task orgz number"`
	TaskUserId           string          `json:"taskUserId" description:"Task user id"`
	TaskUserName         string          `json:"taskUserName" description:"Task user name"`
	FinalModifyOrgz      string          `json:"finalModifyOrgz" description:"Final modify organization"`
	FinalModifyUser      string          `json:"finalModifyUser" description:"Final modify user"`
}

type LoanFeeCalcRule struct {
	FeeId          string          `json:"feeId"`
	IntervalNo     int             `json:"intervalNo"`
	LoanReceiptNum string          `json:"loanReceiptNum"`
	LowerCondition string          `json:"lowerCondition"`
	LowerInterval  decimal.Decimal `json:"lowerInterval"`
	UpperCondition string          `json:"upperCondition"`
	UpperInterval  decimal.Decimal `json:"upperInterval"`
	CalcValue      decimal.Decimal `json:"calcValue"`
	ThresholdFlag  string          `json:"thresholdFlag"`
	ThresholdType  string          `json:"thresholdType"`
}
type LoanFeeType struct {
	FeeId          string          `json:"feeId"`
	FeeType        string          `json:"feeType"`
	CalcBasisType  string          `json:"calcBasisType"`
	FeeUseScene    string          `json:"FeeUseScene"`
	AccountingCode string          `json:"accountingCode"`
	CalcValue      string          `json:"calcValue"`
	Currency       string          `json:"currency"`
	FeeCalcMethod  string          `json:"feeCalcMethod"`
	IsAmortization string          `json:"isAmortization"`
	MaxFeeAmount   decimal.Decimal `json:"maxFeeAmount"`
	MinFeeAmount   decimal.Decimal `json:"minFeeAmount"`
}
type AL000010O struct {
}

func (*AL000010I) GetServiceKey() string {
	return "AL000010"
}
