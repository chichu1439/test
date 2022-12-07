package models

import "github.com/shopspring/decimal"

type PD000070I struct {
	ProductId                 string              `json:"productId" validate:"required" description:"Product Id"`
	SystemId                  string              `json:"systemId" validate:"required" description:"system Id(LN-icredit;SV-isaving)"`
	ProductType               string              `json:"productType" validate:"required" description:"product Type:Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan"` //*Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan
	ProductName               string              `json:"productName" description:"Product name" validate:"required"`
	Currency                  string              `json:"currency" validate:"required" description:"Currency(THB)"`
	Version                   string              `json:"version" description:"Version"`
	Remark                    string              `json:"remark" description:"Remark"`
	EffectiveDate             string              `json:"effectiveDate" description:"Effective date" validate:"required"`
	ExpireDate                string              `json:"expireDate" description:"Expire date" validate:"required"`
	ValueDateMethod           string              `json:"valueDateMethod" description:"value Date Method:01-Approval Date 02-Account Opening Day 03-Loan Date 04-Loan Date plus Interest-free period"` //01-Approval Date 02-Account Opening Day 03-Loan Date 04-Loan Date plus Interest-free period
	MaturityDateMethod        string              `json:"maturityDateMethod" description:"maturity Date Method:01- Value Date and Loan Term;02- Same as the Last Interest Collection Date;03- Loan Date and Loan Term;04- Same as the Last Repayment Date"`
	RepaymentMethod           string              `json:"repaymentMethod" description:"repayment Method:01-Fixed Installment;02-Fixed Principal"`
	RepaymentFrequency        string              `json:"repaymentFrequency" description:"repayment Frequency:01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly"`
	AllowEarlyRepayment       string              `json:"allowEarlyRepayment" validate:"required" description:"allow Early Repayment:0-no;1-yes"`
	AllowExtension            string              `json:"allowExtension" validate:"required" description:"allow Extension:0-no;1-yes"`
	AllowGracePeriod          string              `json:"allowGracePeriod" validate:"required" description:"allow Grace Period:0-no;1-yes"`
	GuaranteeMethod           string              `json:"guaranteeMethod" description:"guarantee Method:0-No Guarantee 1-Guarantee"` //0-No Guarantee 1-Guarantee
	Status                    string              `json:"status" description:"status:0-inactive;1-active;2-expired;3-deleted;4-updated"`
	ApplicantNumberType       string              `json:"applicantNumberType" description:"applicant Number Type:0-single customer"` //0-single customer
	DefaultGracePeriod        int                 `json:"defaultGracePeriod" description:"Default grace period"`
	ListInterest              []MicroLoanInterest `json:"listInterest" description:"List interest"`
	ListFee                   []MicroLoanFee      `json:"listFee" description:"List fee"`
	ListNotification          []PD000070IListNoti `json:"listNotification" description:"List notification"`
	MaxGracePeriod            int                 `json:"maxGracePeriod" description:"Max grace period"`
	MaxLoanAmount             float64             `json:"maxLoanAmount" description:"Max loan amount"`
	MinLoanAmount             float64             `json:"minLoanAmount" description:"Min loan amount"`
	MaxExtensionTimes         int                 `json:"maxExtensionTimes" description:"Max extension times"`
	TotalExtensionDaysOption  string              `json:"totalExtensionDaysOption" description:"Total extension days option"`
	MaxLoanPeriods            int                 `json:"maxLoanPeriods" description:"Max loan periods"`
	MinLoanPeriods            int                 `json:"minLoanPeriods" description:"Min loan periods"`
	RevolvingFlag             string              `json:"revolvingFlag" description:"revolving Flag:0-no;1-yes"`
	InstallmentMethod         string              `json:"installmentMethod" description:"installment Method:01-Fixed Installment;02-Fixed Principal"`
	InstallmentRoundDirection string              `json:"installmentRoundDirection" description:"Installment round direction"`
	InstallmentRoundDigits    int                 `json:"installmentRoundDigits" description:"Installment round digits"`
	CompoundInterestMethod    string              `json:"compoundInterestMethod" description:"0-simple interest;1-compound interest"`
	ImpairedDays              int                 `json:"impairedDays" description:"Impaired days"`
	MaxQuotaAmount            float64             `json:"maxQuotaAmount" description:"Max quota amount"`
	ReleaseQuotaMethod        string              `json:"releaseQuotaMethod" description:"release Quota Method:0-Each repayment;1-loan closed"`
	RepaymentOrderMethod      string              `json:"repaymentOrderMethod" description:"repayment Order Method:0-period first;1-amount first"`
	RepaymentOrderAmount      string              `json:"repaymentOrderAmount" description:"Repayment order amount"`
	MakerComment              string              `json:"makerComment" description:"Maker comment"`
	UpdateFlag                string              `json:"updateFlag" description:"Update flag"`
	InstallmentList           string              `json:"installmentList" description:"Installment list"`
	MaxEarlyRepaymentTimes    int                 `json:"maxEarlyRepaymentTimes" description:"Max early repayment times"`
	MaxContractCount          int                 `json:"maxContractCount" description:"Max contract count"`
	CustomerType              string              `json:"customerType" description:"customer Type:0-personal;1-corporate"`
}

type PD000070IListNoti struct {
	Uid           int    `json:"uid" description:"Uid"`
	NotifyType    string `json:"notifyType" description:"notify Type:01-Loan Account Reminder;02-Loan Account Late Notification;03-Account Statement"`
	TargetType    string `json:"targetType" description:"target Type:0-Customer;1-Internal"`
	StrategyId    int    `json:"strategyId" description:"Strategy id"`
	EffectiveDate string `json:"effectiveDate" description:"Effective date"`
	ExpireDate    string `json:"expireDate" description:"Expire date"`
	RecordType    string `json:"recordType" description:"Record type"`
}

type MicroLoanInterest struct {
	Uid              int     `json:"uid" description:"Uid" validate:"required"`
	InterestId       string  `json:"interestId" description:"Interest id" validate:"required"`
	InterestType     string  `json:"interestType" validate:"required" description:"interest Type:0-normal 1-Withdrawal in Advance"` //0-normal 1-Withdrawal in Advance
	StrategyId       string  `json:"strategyId" description:"Strategy id"`
	BaseInterestRate float64 `json:"baseInterestRate" description:"Base interest rate"`
	MinInterestRate  float64 `json:"minInterestRate" description:"Min interest rate"`
	MaxInterestRate  float64 `json:"maxInterestRate" description:"Max interest rate"`
	FloatDirection   string  `json:"floatDirection" description:"Float direction"`
	FloatType        string  `json:"floatType" description:"0-By Percent;1-BP Value"`
	MinFloatValue    float64 `json:"minFloatValue" description:"Min float value"`
	MaxFloatValue    float64 `json:"maxFloatValue" description:"Max float value"`
	EffectiveDate    string  `json:"effectiveDate" description:"Effective date"`
	ExpireDate       string  `json:"expireDate" description:"Expire date" `
	RecordType       string  `json:"recordType" validate:"required" description:"record Type:A-Add;U-Update;D-Delete;N-No Change"`
	StrategyUid      int     `json:"strategyUid" description:"Strategy uid" validate:"required"`
	YearDays         string  `json:"yearDays" description:"Year days"`
}

type MicroLoanFee struct {
	Uid         int                `json:"uid" description:"Uid" validate:"required"`
	FeatureId   string             `json:"featureId" description:"Feature id"`
	ServiceId   string             `json:"serviceId" description:"Service id"`
	ListFeeType []MicroLoanFeeType `json:"listFeeType" description:"List fee type" validate:"dive"`
}

type MicroLoanFeeType struct {
	Uid           int             `json:"uid" description:"Uid" validate:"required" :"uid"`
	FeatureId     string          `json:"featureId" description:"Feature id" :"feature_id"`
	FeeId         string          `json:"feeId" description:"Fee id" validate:"required" :"fee_id"`
	FeeName       string          `json:"feeName" description:"Fee name" :"fee_name"`
	MaxFee        decimal.Decimal `json:"maxFee" description:"Max fee" :"max_fee"`
	MinFee        decimal.Decimal `json:"minFee" description:"Min fee" :"min_fee"`
	FeeCalcMethod string          `json:"feeCalcMethod" description:"Fee calc method" :"fee_calc_method"`
	CalcValue     decimal.Decimal `json:"calcValue" description:"Calc value" :"calc_value"`
	PayerType     string          `json:"payerType" description:"01-Payee Participant;02-Payer Participant" :"payer_type"`
	EffectiveDate string          `json:"effectiveDate" description:"Effective date" :"effective_date"`
	ExpireDate    string          `json:"expireDate" description:"Expire date" :"expire_date"`
	RecordType    string          `json:"recordType" description:"A-Add;U-Update;D-Delete;N-No Change" :"record_type"`
}

type PD000070O struct {
	ProductId string `json:"productId" description:"Product id" `
}

func (*PD000070I) GetServiceKey() string {
	return "PD000070"
}
