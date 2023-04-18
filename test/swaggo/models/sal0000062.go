package models

import "github.com/shopspring/decimal"

type AL000062I struct {
	NeedOpenAccountFlag        string                   `json:"needOpenAccountFlag" description:"need open account flag Y-need open account,N-no need open account" validate:"required"`
	LoanReceiptNum             string                   `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	GlobalSerialNum            string                   `json:"globalSerialNum" description:"global serial num"`
	CustomerId                 string                   `json:"customerId" description:"Customer id" validate:"required"`
	CustomerName               string                   `json:"customerName" description:"Customer name" validate:"required"`
	ProductId                  string                   `json:"productId" description:"Product id" validate:"required"`
	Currency                   string                   `json:"currency" description:"Currency" validate:"required" `
	LoanAmount                 decimal.Decimal          `json:"loanAmount" description:"Loan amount" validate:"required"`
	RefinanceAmount            decimal.Decimal          `json:"refinanceAmount" description:"Refinance amount"`
	MatureDate                 string                   `json:"matureDate" description:"Mature date" `
	LoanPeriod                 int                      `json:"loanPeriod" description:"Loan period"`
	RepaymentMethod            string                   `json:"repaymentMethod" description:"Repayment method" validate:"required"`
	RepaymentCycle             string                   `json:"repaymentCycle" description:"Repayment cycle" `
	GracePeriod                int                      `json:"gracePeriod" description:"Grace period"`
	AutoRepaymentFlag          string                   `json:"autoRepaymentFlag" description:"Auto repayment flag"`
	OtherBankFlag              string                   `json:"otherBankFlag" description:"Other bank flag"`
	RepaymentAccountBankNumber string                   `json:"repaymentAccountBankNumber" description:"Repayment account bank number"`
	RepaymentAccountBankName   string                   `json:"repaymentAccountBankName" description:"Repayment account bank name"`
	RepaymentAccountNumber     string                   `json:"repaymentAccountNumber" description:"Repayment account number"`
	RepaymentAccountName       string                   `json:"repaymentAccountName" description:"Repayment account name"`
	RepaymentAccountOrgz       string                   `json:"repaymentAccountOrgz" description:"Repayment account organization"`
	ActualMakeLoanDate         string                   `json:"actualMakeLoanDate" description:"Actual make loan date"`
	RepayDay                   string                   `json:"repayDay" description:"Repay day"`
	OpenAccountDate            string                   `json:"openAccountDate" description:"Open account date"`
	OpenAccountOrgz            string                   `json:"openAccountOrgz" description:"Open account organization"`
	ImpairedDays               int                      `json:"impairedDays" description:"Impaired days"`
	CompoundInterestFlag       string                   `json:"compoundInterestFlag" description:"Compound interest flag"`
	GraceType                  string                   `json:"graceType"`
	FirstRepaymentDate         string                   `json:"firstRepaymentDate"`
	LoanBillRule               LoanBillRule             `json:"loanBillRule" description:"loan bill rule info"`
	ListTLoanFeeType           []TLoanFeeType           `json:"listTLoanFeeType" description:"Loan fee type list"`
	ListLoanInterestType       []TLoanInterestType02    `json:"listLoanInterestType" description:"Loan interest type list"`
	ListFee                    []FeeTransDetail         `json:"listFee"`
	AcctRepaymentHierarchy     []AcctRepaymentHierarchy `json:"acctRepaymentHierarchy"`
}

type AcctAmortizationPeriod struct {
	AccountingAccountNum   string          `json:"accountingAccountNum"`
	AmountType             string          `json:"amountType"`
	CurrentStatus          string          `json:"currentStatus"`
	Period                 int             `json:"period"`
	CurrentAccountStatus   string          `json:"currentAccountStatus"`
	CurrentPeriodBeginDate string          `json:"currentPeriodBeginDate"`
	CurrentPeriodEndDate   string          `json:"currentPeriodEndDate"`
	CurrentRepayDate       string          `json:"currentRepayDate"`
	PlanRepayPrincipal     decimal.Decimal `json:"planRepayPrincipal"`
	ActualRepayPrincipal   decimal.Decimal `json:"actualRepayPrincipal"`
	CurrentUnpaidPrincipal decimal.Decimal `json:"currentUnpaidPrincipal"`
	PlanRepayInterest      decimal.Decimal `json:"planRepayInterest"`
	ActualRepayInterest    decimal.Decimal `json:"actualRepayInterest"`
	ActualRepayDate        string          `json:"actualRepaydate"`
	CloseDate              string          `json:"closeDate"`
	LastActiveDate         string          `json:"lastActiveDate"`
	MaintainPrinciple      string          `json:"maintainPrinciple"`
}
type FeeTransDetail struct {
	FeeType     string          `json:"feeType"`
	FeeAmount   decimal.Decimal `json:"feeAmount"`
	OnTop       string          `json:"onTop"`
	OnTopPeriod int             `json:"onTopPeriod"`
}
type AcctRepaymentHierarchy struct {
	LoanReceiptNum         string `json:"loanReceiptNum"`
	RepaymentHierarchyType string `json:"repaymentHierarchyType"`
	RepaymentHierarchyName string `json:"repaymentHierarchyName"`
	RepaymentOrder         string `json:"repaymentOrder"`
	Behavior               string `json:"behavior"`
}
type LoanBillRule struct {
	BillDay               string          `json:"billDay" description:"Bill day"`
	NextBillDate          string          `json:"nextBillDate" description:"Next bill date"`
	SameDayTreatment      string          `json:"sameDayTreatment" description:"Bill day and repayment day in the same day(Y-same date;N-repaymen date in the next month of bill date)"`
	BillRepaymentType     string          `json:"billRepaymentType" description:"Bill repayment type(0- Percentage of Principal；1- Free Minimum Amount；2- Interest Minimum Amount；3- Remaining Outstanding Amount；)"`
	PercentageOfPrincipal decimal.Decimal `json:"percentageOfPrincipal" description:"Percentage of principal"`
	FixMinimumAmount      decimal.Decimal `json:"fixMinimumAmount" description:"Fix minimum amount"`
}

type TLoanInterestType02 struct {
	InterestType       string          `json:"interestType" description:"Interest type:O-Normal Loan 1-Overdue Loan"`
	CustomerGrade      string          `json:"customerGrade" description:"Customer grade"`
	InterestId         string          `json:"interestId" description:"Interest id"`
	InterestStrategyId string          `json:"interestStrategyId" description:"Interest strategy id"`
	BaseInterestRate   decimal.Decimal `json:"baseInterestRate" description:"Base Interest Rate"`
	FloatValue         decimal.Decimal `json:"floatValue" description:"Float value"`
	InterestRate       decimal.Decimal `json:"interestRate" description:"Interest rate"`
	FloatDirection     string          `json:"floatDirection" description:"Float direction"`
	FloatMethod        string          `json:"floatMethod" description:"Float method"`
	YearDays           string          `json:"yearDays" description:"Year Days"`
	MonthDays          string          `json:"monthDays" description:"Month Days"`
}
type TLoanFeeType struct {
	FeeId               string             `json:"feeId"`
	FeeType             string             `json:"feeType"`
	CalcBasisType       string             `json:"calcBasisType"`
	FeeUseScene         string             `json:"FeeUseScene"`
	AccountingCode      string             `json:"accountingCode"`
	CalcValue           string             `json:"calcValue"`
	Currency            string             `json:"currency"`
	FeeCalcMethod       string             `json:"feeCalcMethod"`
	IsAmortization      string             `json:"isAmortization"`
	MaxFeeAmount        decimal.Decimal    `json:"maxFeeAmount"`
	MinFeeAmount        decimal.Decimal    `json:"minFeeAmount"`
	ListLoanFeeCalcRule []TLoanFeeCalcRule `json:"listLoanFeeCalcRule"`
}
type TLoanFeeCalcRule struct {
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
type AL000062O struct {
	Status               string `json:"status" description:"status"`
	AccountingAccountNum string `json:"accountingAccountNum" description:"Accounting account number" `
}

func (*AL000062I) GetServiceKey() string {
	return "AL000062"
}
