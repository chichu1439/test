package models

import "github.com/shopspring/decimal"

type SIC0000006I struct {
	CreditApplySerialNum string              `json:"creditApplySerialNum" validate:"required" description:"Credit Apply Serial number"`
	LoanProductId        string              `json:"loanProductId" description:"Loan product id" validate:"required"`
	CustomerId           string              `json:"customerId" description:"Customer id" validate:"required"`
	LoanAmount           decimal.Decimal     `json:"loanAmount" description:"Loan amount" validate:"required"`
	PeriodNum            int                 `json:"periodNum" description:"Period number" `
//	LoanPurpose          string              `json:"loanPurpose" validate:"required" description:"Loan Purpose"` //
	AppDate              string              `json:"appDate" description:"App date" validate:"required"`
	DrawDate             string              `json:"drawDate" description:"Draw date" validate:"required"`
	MatureDate           string              `json:"matureDate" validate:"required" description:"Mature Date"`
	RepayDay             string              `json:"repayDay" description:"Repay day" `
	RepayCycle           string              `json:"repayCycle" description:"Repay cycle:01-daily 02-weekly 04-Monthly 08-Yearly"`      //01- daily 02- weekly 04- Monthly 08- Yearly
	AutoReplyFlag        string              `json:"autoReplyFlag" validate:"required" description:"Auto reply flag:01-Auto repayment"` //01-Auto repayment
	Maker                string              `json:"maker" description:"Maker" validate:"required"`
	MakerComment         string              `json:"makerComment" validate:"required" description:"Maker Comment"`
	GracePeriod          int                 `json:"gracePeriod" description:"Grace period" validate:""`
	RepayMethod          string              `json:"repayMethod" validate:"required" description:"Repay method:01-equal principal 02-equal principal and interest 03-due (one-time) repayment of principal and interest 04-due repayment of principal and interest 05-equal principal and interest 06-repayment with borrowing"`
	ListInterest         []MicroLoanInterest `json:"listInterest" description:"List interest"`
	ListFee              []MicroLoanFee      `json:"listFee" description:"List fee"`
	BillRuleInfo         *BillRuleInfo       `json:"billRuleInfo" description:"Bill rule information"`
}

type Fee struct {
	FeeType         string          `json:"feeType" description:"0-account opening"`
	Currency        string          `json:"currency" description:"Currency"`
	FeeChargeMethod string          `json:"feeChargeMethod" description:"Fee charge method"`
	Amount          decimal.Decimal `json:"amount" description:"Amount"`
	Percent         decimal.Decimal `json:"percent" description:"Percent"`
	MaxFeeAmount    decimal.Decimal `json:"maxFeeAmount" description:"Max fee amount"`
	MinFeeAmount    decimal.Decimal `json:"minFeeAmount" description:"Min fee amount"`
}

type MicroLoanInterest struct {
	InterestType       string          `json:"interestType" description:"Interest type:O-Normal Loan 1-Overdue Loan"`
	CustomerGrade      string          `json:"customerGrade" description:"Customer grade"`
	InterestId         string          `json:"interestId" description:"Interest id"`
	InterestStrategyId string          `json:"interestStrategyId" description:"Interest strategy id"`
	BaseInterestRate   decimal.Decimal `json:"baseInterestRate" description:"Base Interest Rate"`
	FloatValue         decimal.Decimal `json:"floatValue" description:"Float value"`
	InterestRate       decimal.Decimal `json:"interestRate" description:"Interest rate"`
	FloatDirection     string          `json:"floatDirection" description:"Float direction"`
	FloatMethod        string          `json:"floatMethod" description:"Float method"`
	YearDays           int             `json:"yearDays" description:"Year Days"`
}

type MicroLoanFee struct {
	LoanFeeType        string          `json:"loanFeeType"`
	FeeId              string          `json:"feeId" description:"Fee id"`
	FeeName            string          `json:"feeName" description:"Fee name"`
	StrategyId         string          `json:"strategyId" description:"Strategy id"`
	FeeAmount          string          `json:"feeAmount" description:"Fee amount"`
	DiscountAmount     decimal.Decimal `json:"discountAmount" description:"Discount amount"`
	FeeAmountActual    decimal.Decimal `json:"feeAmountActual" description:"Fee amount actual"`
	IsAmortization     string          `json:"isAmortization" description:"Is amortization"`
	AmortizationPeriod string          `json:"amortizationPeriod" description:"Amortization period"`
	AmortizationTimes  string          `json:"amortizationTimes" description:"Amortization times"`
	ChargeType         string          `json:"chargeType" description:"Charge type"`
	ChargePeriod       string          `json:"chargePeriod" description:"Charge period"`
	ChargePeriodValue  string          `json:"chargePeriodValue" description:"Charge period value"`
	ChargePeriodDay    string          `json:"chargePeriodDay" description:"Charge period day"`
	MaxFeeAmount       decimal.Decimal `json:"maxFeeAmount" description:"Max fee amount"`
	MinFeeAmount       decimal.Decimal `json:"minFeeAmount" description:"Min fee amount"`
	AccountingCode     string          `json:"accountingCode" description:"Accounting code"`
}
type BillRuleInfo struct {
	BillDay               string          `json:"billDay" description:"Bill day"`
	CurrentBillDate       string          `json:"currentBillDate" description:"Current bill date"`
	TotalBillPeriod       int             `json:"totalBillPeriod" description:"Total bill period"`
	NextBillDate          string          `json:"nextBillDate" description:"Next bill date"`
	SameDayTreatment      string          `json:"sameDayTreatment" description:"Bill day and repayment day in the same day(Y-same date;N-repaymen date in the next month of bill date)"`
	BillCycle             string          `json:"billCycle" description:"Bill cycle(04-by month)"`
	BillRepaymentType     string          `json:"billRepaymentType" description:"Bill repayment type(0- Percentage of Principal；1- Free Minimum Amount；2- Interest Minimum Amount；3- Remaining Outstanding Amount；)"`
	PercentageOfPrincipal decimal.Decimal `json:"percentageOfPrincipal" description:"Percentage of principal"`
	FixMinimumAmount      decimal.Decimal `json:"fixMinimumAmount" description:"Fix minimum amount"`
}
type SIC0000006O struct {
	MakeLoanApplySerialNum string `json:"makeLoanApplySerialNum" description:"Make loan apply serial number"`
	LoanReceiptNum         string `json:"loanReceiptNum" description:"Loan receipt number"`
}

func (*SIC0000006I) GetServiceKey() string {
	return "sic0000006"
}
