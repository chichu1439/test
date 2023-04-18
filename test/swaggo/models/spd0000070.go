package models

import "github.com/shopspring/decimal"

type PD000070I struct {
	// basis
	ProductId     string `json:"productId" validate:"required" description:"Product Id"`
	SystemId      string `json:"systemId" validate:"required" description:"system Id(LN-icredit;SV-isaving)"`
	ProductType   string `json:"productType" validate:"required" description:"product Type:Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan"` //*Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan
	ProductName   string `json:"productName" description:"Product name" validate:"required"`
	CustomerType  string `json:"customerType" description:"customer Type:0-personal;1-corporate"`
	Version       string `json:"version" description:"Version"`
	Remark        string `json:"remark" description:"Remark"`
	EffectiveDate string `json:"effectiveDate" description:"Effective date" validate:"required"`
	ExpireDate    string `json:"expireDate" description:"Expire date" validate:"required"`
	// Guarantee
	GuaranteeMethod string `json:"guaranteeMethod" description:"guarantee Method:0-No Guarantee 1-Guarantee"` //0-No Guarantee 1-Guarantee
	// product detail
	Currency           string `json:"currency" validate:"required" description:"Currency(THB)"`
	MaxContractCount   int    `json:"maxContractCount" description:"Max contract count"`
	InstallmentMethod  string `json:"installmentMethod" description:"installment Method:01-Fixed Installment;02-Fixed Principal"`
	InstallmentList    string `json:"installmentList" description:"Installment list"`
	RepaymentFrequency string `json:"repaymentFrequency" description:"repayment Frequency:01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly"`
	// repayment
	FirstRepayDateMethod string `json:"firstRepayDateMethod"` // 0-Automatic；1-Manual
	RepaymentMethod      string `json:"repaymentMethod" description:"repayment Method:01-Fixed Installment;02-Fixed Principal"`
	RepaymentOrderMethod string `json:"repaymentOrderMethod" description:"repayment Order Method:0-period first;1-amount first"`
	// bill payment
	PrincipalRule []PrincipalRule `json:"principalRule" description:"Principal rule"`
	// Compound Interest
	CompoundInterestMethod string `json:"compoundInterestMethod" description:"0-simple interest;1-compound interest"`
	// loan Periods
	MaxLoanPeriods int `json:"maxLoanPeriods" description:"Max loan periods"`
	MinLoanPeriods int `json:"minLoanPeriods" description:"Min loan periods"`
	// Early Repayment
	AllowEarlyRepayment    string `json:"allowEarlyRepayment" validate:"required" description:"allow Early Repayment:0-no;1-yes"`
	MaxEarlyRepaymentTimes int    `json:"maxEarlyRepaymentTimes" description:"Max early repayment times"`
	// Late Repayment
	AllowGracePeriod string `json:"allowGracePeriod" validate:"required" description:"allow Grace Period:0-no;1-yes"`
	GraceType        string `json:"graceType"`
	MaxGracePeriod   int    `json:"maxGracePeriod" description:"Max grace period"`
	ImpairedDays     int    `json:"impairedDays" description:"Impaired days"`
	// Amount
	ReleaseQuotaMethod string          `json:"releaseQuotaMethod" description:"release Quota Method:0-Each repayment;1-loan closed"`
	MaxLoanAmount      decimal.Decimal `json:"maxLoanAmount" description:"Max loan amount"`
	MinLoanAmount      decimal.Decimal `json:"minLoanAmount" description:"Min loan amount"`
	MaxContractAmount  decimal.Decimal `json:"maxContractAmount" description:"Max contract amount"`
	MinContractAmount  decimal.Decimal `json:"minContractAmount" description:"min contract amount"`
	// Overdue Credit Line Control
	//0-不影响额度 1-按逾期天数比例调减 2-全额冻结
	QuotaControlOverdue string        `json:"quotaControlOverdue" description:"0-Overdue does not affect the quota 1-Decrease in proportion to the number of overdue days 2-Frozen in full"`
	OverDueRule         []OverDueRule `json:"overDueRule" description:"OverDue rule"`
	// interest
	ListInterest []MicroLoanInterest `json:"listInterest" description:"List interest"`
	// fee
	ListFee []MicroLoanFee `json:"listFee" description:"List fee"`
	// notification
	ListNotification []PD000070IListNoti `json:"listNotification" description:"List notification"`
	// repayment hierarchy
	RepaymentHierarchy []PD000070RepaymentHierarchy `json:"repaymentHierarchy"`
	// maker
	Maker         string `json:"maker"`
	MakerComment  string `json:"makerComment" description:"Maker comment"`
	UpdateFlag    string `json:"updateFlag" description:"Update flag"`
	Status        string `json:"status" description:"status:0-inactive;1-active;2-expired;3-deleted;4-updated"`
	RevolvingFlag string `json:"revolvingFlag" description:"revolving Flag:0-no;1-yes"`

	DefaultGracePeriod        int    `json:"defaultGracePeriod" description:"Default grace period"`
	MaturityDateMethod        string `json:"maturityDateMethod" description:"maturity Date Method:01- Value Date and Loan Term;02- Same as the Last Interest Collection Date;03- Loan Date and Loan Term;04- Same as the Last Repayment Date"`
	ApplicantNumberType       string `json:"applicantNumberType" description:"applicant Number Type:0-single customer"` //0-single customer
	InstallmentRoundDirection string `json:"installmentRoundDirection" description:"Installment round direction"`
	InstallmentRoundDigits    int    `json:"installmentRoundDigits" description:"Installment round digits"`
	ValueDateMethod           string `json:"valueDateMethod" description:"value Date Method:01-Approval Date 02-Account Opening Day 03-Loan Date 04-Loan Date plus Interest-free period"` //01-Approval Date 02-Account Opening Day 03-Loan Date 04-Loan Date plus Interest-free period
}

type PD000070IListNoti struct {
	Uid           int    `json:"uid" description:"Uid"`
	NotifyType    string `json:"notifyType" description:"notify Type:01-Loan Account Reminder;02-Loan Account Late Notification;03-Account Statement"`
	TargetType    string `json:"targetType" description:"target Type:0-Customer;1-Internal"`
	StrategyId    int    `json:"strategyId" description:"Strategy id"`
	RecordType    string `json:"recordType" description:"Record type"`
	EffectiveDate string `json:"effectiveDate" description:"Effective date" validate:"required"`
	ExpireDate    string `json:"expireDate" description:"Expire date" validate:"required"`
}

type MicroLoanInterest struct {
	InterestType      string                  `json:"interestType" validate:"required" description:"interest Type:0-normal 1-Withdrawal in Advance"` //0-normal 1-Withdrawal in Advance
	CustomerGradeType string                  `json:"customerGradeType"`
	UseInterestFlag   bool                    `json:"useInterestFlag"`
	StrategyList      []MicroLoanInterestItem `json:"strategyList"`
}
type MicroLoanInterestItem struct {
	InterestUid      int             `json:"interestUid" description:"Uid" validate:"required"`
	InterestId       string          `json:"interestId" description:"Interest id" validate:"required"`
	CustomerGrade    string          `json:"customerGrade" description:"0-All customers 1-Level 1 2-Level 2 3-Level 3"`
	StrategyId       string          `json:"strategyId" description:"Strategy id"`
	BaseInterestRate decimal.Decimal `json:"baseInterestRate" description:"Base interest rate"`
	MinInterestRate  decimal.Decimal `json:"minInterestRate" description:"Min interest rate"`
	MaxInterestRate  decimal.Decimal `json:"maxInterestRate" description:"Max interest rate"`
	FloatDirection   string          `json:"floatDirection" description:"Float direction"`
	FloatType        string          `json:"floatType" description:"0-By Percent;1-BP Value"`
	MinFloatValue    decimal.Decimal `json:"minFloatValue" description:"Min float value"`
	MaxFloatValue    decimal.Decimal `json:"maxFloatValue" description:"Max float value"`
	RecordType       string          `json:"recordType" validate:"required" description:"record Type:A-Add;U-Update;D-Delete;N-No Change"`
	StrategyUid      int             `json:"strategyUid" description:"Strategy uid" validate:"required"`
	YearDays         string          `json:"yearDays" description:"Year days"`
	EffectiveDate    string          `json:"effectiveDate" description:"Effective date" validate:"required"`
	ExpireDate       string          `json:"expireDate" description:"Expire date" validate:"required"`
}
type PD000070RepaymentHierarchy struct {
	RepaymentHierarchyType string `json:"repaymentHierarchyType" description:"repayment hierarchy type" `
	RepaymentHierarchyName string `json:"repaymentHierarchyName"`
	Code                   string `json:"code"`
	RepaymentOrder         string `json:"repaymentOrder" description:"repayment order ,format； repayment priority：repayment Hierarchy" `
	Description            string `json:"description"`
	Behavior               string `json:"behavior"`
}
type PD000070O struct {
	ProductId string `json:"productId" description:"Product id" `
}

func (*PD000070I) GetServiceKey() string {
	return "PD000070"
}
