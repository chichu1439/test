package models

import "github.com/shopspring/decimal"

type SIC0000006I struct {
	ProductId       string               `json:"productId" description:"Loan product id" validate:"required"`
	CustomerId      string               `json:"customerId" description:"Customer id" validate:"required"`
	PrincipalAmount decimal.Decimal      `json:"principalAmount" description:"Principal amount" validate:"required"`
	ContractAmount  decimal.Decimal      `json:"contractAmount" description:"Contract amount" validate:"required"`
	PeriodNum       int                  `json:"periodNum" description:"Period number" `
	LoanPurpose     string               `json:"loanPurpose" validate:"required" description:"Loan Purpose"` //
	AppDate         string               `json:"appDate" description:"App date" validate:"required"`
	DrawDate        string               `json:"drawDate" description:"Draw date" validate:"required"`
	MatureDate      string               `json:"matureDate" validate:"required" description:"Mature Date"`
	RepayDay        string               `json:"repayDay" description:"Repay day" `
	RepayCycle      string               `json:"repayCycle" description:"Repay cycle:01-daily 02-weekly 04-Monthly 08-Yearly"`      //01- daily 02- weekly 04- Monthly 08- Yearly
	AutoReplyFlag   string               `json:"autoReplyFlag" validate:"required" description:"Auto reply flag:01-Auto repayment"` //01-Auto repayment
	MakerComment    string               `json:"makerComment" validate:"required" description:"Maker Comment"`
	GracePeriod     int                  `json:"gracePeriod" description:"Grace period"`
	RepayMethod     string               `json:"repayMethod" validate:"required" description:"Repay method:01-equal principal 02-equal principal and interest 03-due (one-time) repayment of principal and interest 04-due repayment of principal and interest 05-equal principal and interest 06-repayment with borrowing"`
	ListInterest    []MicroLoanInterest2 `json:"listInterest" description:"List interest"`
	//BillRuleInfo
	BillDay               string          `json:"billDay" description:"Bill day"`
	NextBillDate          string          `json:"nextBillDate" description:"Next bill date"`
	SameDayTreatment      string          `json:"sameDayTreatment" description:"Bill day and repayment day in the same day(Y-same date;N-repaymen date in the next month of bill date)"`
	BillRepaymentType     string          `json:"billRepaymentType" description:"Bill repayment type(0- Percentage of Principal；1- Free Minimum Amount；2- Interest Minimum Amount；3- Remaining Outstanding Amount；)"`
	PercentageOfPrincipal decimal.Decimal `json:"percentageOfPrincipal" description:"Percentage of principal"`
	FixMinimumAmount      decimal.Decimal `json:"fixMinimumAmount" description:"Fix minimum amount"`
	CreditCardNumber      string          `json:"creditCardNumber" description:"Credit card number"`
	CardHoldName          string          `json:"cardHoldName" description:"card hold name"`
	ExpireMonth           string          `json:"expireMonth" description:"Expire month"`
	ExpireYear            string          `json:"expireYear" description:"expire year"`
	SecurityCode          string          `json:"securityCode" description:"security code"`
	BankAccountName       string          `json:"bankAccountName" description:"bank account name" validate:"required"`
	BankAccountNumber     string          `json:"bankAccountNumber" description:"bank account number" validate:"required"`
	BankName              string          `json:"bankName" description:"bank name" validate:"required"`
	Bic                   string          `json:"bic" description:"bic" validate:"required"`
	GuaranteeMethod       string          `json:"guaranteeMethod" description:"guarantee method" validate:"required" `
	GuaranteeInfo         GuaranteeInfo   `json:"guaranteeInfo" description:"guarantee info" validate:"required,dive" `
	SourceChannel         string          `json:"sourceChannel"`
	FirstRepaymentDate    string          `json:"firstFepaymentDate"`
	ListFee               []FeeInfo2      `json:"listFee"`
}
type GuaranteeInfo struct {
	Name        string `json:"name" description:"name" validate:"required"`
	IdType      string `json:"idType" description:"id type" validate:"required"`
	IdNation    string `json:"idNation" description:"id nation" validate:"required"`
	IdNum       string `json:"idNum" description:"id num" validate:"required"`
	PhoneNumber string `json:"phoneNumber" description:"phone number" validate:"required"`
}
type MicroLoanInterest2 struct {
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
}

type MicroLoanFee2 struct {
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
type FeeInfo2 struct {
	FeeType     string          `json:"feeType"`
	FeeAmount   decimal.Decimal `json:"feeAmount"`
	OnTop       string          `json:"onTop" description:"OnTop:0-No,1-Yes"`
	OnTopPeriod int             `json:"onTopPeriod"`
}

type SIC0000006O struct {
	MakeLoanApplySerialNum string `json:"makeLoanApplySerialNum" description:"Make loan apply serial number"`
}

func (*SIC0000006I) GetServiceKey() string {
	return "IC000006"
}
