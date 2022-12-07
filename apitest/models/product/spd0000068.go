package models

import "github.com/shopspring/decimal"

type PD000068I struct {
	SystemId                  string                `json:"systemId" validate:"required,max=2" description:"system Id:SV-Saving LN-Loan"`
	ProductType               string                `json:"productType" validate:"required,max=2" description:"product Type:*Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan"`
	ProductName               string                `json:"productName" description:"Product name" validate:"required"`
	Currency                  string                `json:"currency" description:"Currency" validate:"required"`
	Version                   string                `json:"version" description:"Version"`
	Remark                    string                `json:"remark" description:"Remark"`
	EffectiveDate             string                `json:"effectiveDate" description:"Effective date" validate:"required"`
	ExpireDate                string                `json:"expireDate" description:"Expire date" validate:"required"`
	ValueDateMethod           string                `json:"valueDateMethod" description:"value Date Method:01-Approval Date 02-Account Opening Day 03-Loan Date 04-Loan Date plus Interest-free period"`                                                      //01-Approval Date 02-Account Opening Day 03-Loan Date 04-Loan Date plus Interest-free period
	MaturityDateMethod        string                `json:"maturityDateMethod" description:"maturity Date Method:01- Value Date and Loan Term 02- Same as the Last Interest Collection Date 03- Loan Date and Loan Term 04- Same as the Last Repayment Date"` //
	RepaymentMethod           string                `json:"repaymentMethod" description:"repayment Method:1-Fixed Principal； 2-Fixed Installment； 3-Periodic Interest Payment & Principal Repayment at Maturity"`
	RepaymentFrequency        string                `json:"repaymentFrequency" description:"repayment Frequency:1-Daily; 2-Weekly; 3-Double Week; 4-Monthly; 5-Quarterly; 6-Harf a Year; 7-Yearly"`
	AllowEarlyRepayment       string                `json:"allowEarlyRepayment" validate:"required" description:"allow Early Repayment:0-no 1-yes"` //0-no 1-yes
	AllowExtension            string                `json:"allowExtension" validate:"required" description:"allow Extension:0-no 1-yes"`            //0-no 1-yes
	AllowGracePeriod          string                `json:"allowGracePeriod" validate:"required" description:"allow Grace Period:0-no 1-yes"`       //0-no 1-yes
	GuaranteeMethod           string                `json:"guaranteeMethod" validate:"required" description:"guarantee Method:0-No Guarantee 1-Guarantee""`
	ApplicantNumberType       string                `json:"applicantNumberType" description:"applicant Number Type:0-single customer"`
	DefaultGracePeriod        int                   `json:"defaultGracePeriod" description:"Default grace period"`
	ListInterest              []MicroLoanInterest02 `json:"listInterest" description:"List interest" validate:"dive"`
	ListFee                   []MicroLoanFee02      `json:"listFee" description:"List fee"`
	ListNotification          []PD000068ListNoti    `json:"listNotification" description:"List notification"`
	MaxGracePeriod            int                   `json:"maxGracePeriod" description:"Max grace period"`
	MaxLoanAmount             float64               `json:"maxLoanAmount" description:"Max loan amount"`
	MinLoanAmount             float64               `json:"minLoanAmount" description:"Min loan amount"`
	MaxExtensionTimes         int                   `json:"maxExtensionTimes" description:"Max extension times"`
	TotalExtensionDaysOption  string                `json:"totalExtensionDaysOption" description:"Total extension days option"`
	MaxLoanPeriods            int                   `json:"maxLoanPeriods" description:"Max loan periods"`
	MinLoanPeriods            int                   `json:"minLoanPeriods" description:"Min loan periods"`
	RevolvingFlag             string                `json:"revolvingFlag" description:"0-no;1-yes"`
	InstallmentMethod         string                `json:"installmentMethod" description:"Installment method"`
	InstallmentRoundDirection string                `json:"installmentRoundDirection" description:"Installment round direction"`
	InstallmentRoundDigits    int                   `json:"installmentRoundDigits" description:"Installment round digits"`
	CompoundInterestMethod    string                `json:"compoundInterestMethod" description:"compound Interest Method:0-simple interest;1-compound interest"`
	ImpairedDays              int                   `json:"impairedDays" description:"Impaired days"`
	MaxQuotaAmount            float64               `json:"maxQuotaAmount" description:"Max quota amount"`
	ReleaseQuotaMethod        string                `json:"releaseQuotaMethod" description:"release Quota Method:0-Each repayment;1-loan closed"`
	RepaymentOrderMethod      string                `json:"repaymentOrderMethod" description:"repayment Order Method:0-period first;1-amount first"`
	RepaymentOrderAmount      string                `json:"repaymentOrderAmount" description:"Repayment order amount"`
	MakerComment              string                `json:"makerComment" description:"Maker comment"`
	MaxEarlyRepaymentTimes    int                   `json:"maxEarlyRepaymentTimes" description:"Max early repayment times"`
	InstallmentList           string                `json:"installmentList" description:"Installment list"`
	MaxContractCount          int                   `json:"maxContractCount" description:"Max contract count"`
	CustomerType              string                `json:"customerType" description:"Customer type"`
}

type PD000068ListNoti struct {
	NotifyType string `json:"notifyType" description:"notify Type:01-Loan Account Reminder;02-Loan Account Late Notification;03-Account Statement"`
	TargetType string `json:"targetType" description:"target Type:0-Customer;1-Internal"`
	StrategyId int    `json:"strategyId" description:"Strategy id"`
}

type MicroLoanInterest02 struct {
	Uid              int     `json:"uid" description:"Uid"`
	InterestId       string  `json:"interestId" description:"Interest id"`
	InterestType     string  `json:"interestType" validate:"required" description:"interest Type:0-normal 1-Withdrawal in Advance"` //0-normal 1-Withdrawal in Advance
	StrategyId       string  `json:"strategyId" description:"Strategy id"`
	BaseInterestRate float64 `json:"baseInterestRate" description:"Base interest rate"`
	MinInterestRate  float64 `json:"minInterestRate" description:"Min interest rate"`
	MaxInterestRate  float64 `json:"maxInterestRate" description:"Max interest rate"`
	FloatDirection   string  `json:"floatDirection" description:"Float direction"`
	FloatType        string  `json:"floatType" description:"Float type"`
	MinFloatValue    float64 `json:"minFloatValue" description:"Min float value"`
	MaxFloatValue    float64 `json:"maxFloatValue" description:"Max float value"`
	// EffectiveDate      string  `json:"effectiveDate" description:"Effective date"`
	// ExpireDate         string  `json:"expireDate" description:"Expire date" `
	// RecordType         string  `json:"recordType" description:"Record type"`
	YearDays string `json:"yearDays" description:"Year days"`
	// InterestStrategyId string  `json:"interestStrategyId" description:"Interest strategy id"`
}

type MicroLoanFee02 struct {
	Uid         int                  `json:"uid" description:"Uid"`
	FeatureId   string               `json:"featureId" description:"Feature id"`
	ServiceId   string               `json:"serviceId" description:"Service id"`
	ListFeeType []MicroLoanFeeType02 `json:"listFeeType" description:"List fee type"`
}

type MicroLoanFeeType02 struct {
	Uid           int             `json:"uid" description:"Uid"`
	FeatureId     string          `json:"featureId" description:"Feature id"`
	FeeId         string          `json:"feeId" description:"Fee id"`
	FeeName       string          `json:"feeName" description:"Fee name"`
	MaxFee        decimal.Decimal `json:"maxFee" description:"Max fee"`
	MinFee        decimal.Decimal `json:"minFee" description:"Min fee"`
	FeeCalcMethod string          `json:"feeCalcMethod" description:"Fee calc method"`
	CalcValue     decimal.Decimal `json:"calcValue" description:"Calc value"`
	PayerType     string          `json:"payerType" description:"Payer type"`
	EffectiveDate string          `json:"effectiveDate" description:"Effective date"`
	ExpireDate    string          `json:"expireDate" description:"Expire date" `
	RecordType    string          `json:"recordType" description:"Record type"`
}

type PD000068O struct {
	ProductId string `json:"productId" description:"Product id"`
}

func (*PD000068I) GetServiceKey() string {
	return "PD000068"
}
