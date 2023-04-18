package models

import "github.com/shopspring/decimal"

type PD000100I struct {
	ProductId string `json:"productId" description:"Product Id"`
}

type PD000100O struct {
	ChangeOldRecord PD0001000Temp `json:"changeOldRecord" description:"Old record before change"`
	Maker           string        `json:"maker" description:"Maker user"`
	MakerComment    string        `json:"makerComment" description:"Maker comment"`
	MakerDateTime   string        `json:"makerDateTime" description:"Maker date time"`
	Checker         string        `json:"checker" description:"Checker user"`
	ApproveComment  string        `json:"approveComment" description:"Checker comment"`
	CheckerDateTime string        `json:"checkerDateTime" description:"Checker date time"`
}

type PD0001000Temp struct {
	ProductId              string                       `json:"productId" validate:"required" description:"Product Id"`
	SystemId               string                       `json:"systemId" validate:"required" description:"System Id(LN-icredit;SV-isaving)"`
	ProductType            string                       `json:"productType" validate:"required" description:"product Type:Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan"` //*Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan
	ProductName            string                       `json:"productName" validate:"required" description:"Product Name"`
	CustomerType           string                       `json:"customerType" description:"0-personal;1-corporate"`
	Version                string                       `json:"version" description:"Version"`
	Remark                 string                       `json:"remark" description:"Remark"`
	EffectiveDate          string                       `json:"effectiveDate" description:"Effective date" validate:"required"`
	ExpireDate             string                       `json:"expireDate" description:"Expire date" validate:"required"`
	GuaranteeMethod        string                       `json:"guaranteeMethod" description:"guarantee Method:0-No Guarantee 1-Guarantee"` //0-No Guarantee 1-Guarantee
	Currency               string                       `json:"currency" validate:"required" description:"Curreny(THB)"`
	MaxContractCount       int                          `json:"maxContractCount" description:"Max contract count"`
	InstallmentMethod      string                       `json:"installmentMethod" description:"installment Method:01-Fixed Installment;02-Fixed Principal"`
	InstallmentList        string                       `json:"installmentList" description:"Installment list"`
	RepaymentFrequency     string                       `json:"repaymentFrequency" description:"repayment Frequency:01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly"`
	FirstRepayDateMethod   string                       `json:"firstRepayDateMethod"` // 0-Automatic；1-Manual
	RepaymentMethod        string                       `json:"repaymentMethod" description:"repayment Method:01-Fixed Installment;02-Fixed Principal"`
	RepaymentOrderMethod   string                       `json:"repaymentOrderMethod" description:"repayment Order Method:0-period first;1-amount first"`
	PrincipalRule          []PrincipalRule              `json:"principalRule" description:"Principal rule"`
	CompoundInterestMethod string                       `json:"compoundInterestMethod" description:"0-simple interest;1-compound interest"`
	MaxLoanPeriods         int                          `json:"maxLoanPeriods" description:"Max loan periods"`
	MinLoanPeriods         int                          `json:"minLoanPeriods" description:"Min loan periods"`
	AllowEarlyRepayment    string                       `json:"allowEarlyRepayment" validate:"required" description:"allow Early Repayment:0-No;1-Yes"`
	MaxEarlyRepaymentTimes int                          `json:"maxEarlyRepaymentTimes" description:"Max early repayment times"`
	AllowGracePeriod       string                       `json:"allowGracePeriod" validate:"required" description:"allow Grace Period:0-No;1-Yes"`
	GraceType              string                       `json:"graceType"`
	MaxGracePeriod         int                          `json:"maxGracePeriod" description:"Max grace period"`
	ImpairedDays           int                          `json:"impairedDays" description:"Impaired days"`
	ReleaseQuotaMethod     string                       `json:"releaseQuotaMethod" description:"release Quota Method:0-Each repayment;1-loan closed"`
	MaxLoanAmount          decimal.Decimal              `json:"maxLoanAmount" description:"Max loan amount"`
	MinLoanAmount          decimal.Decimal              `json:"minLoanAmount" description:"Min loan amount"`
	MaxContractAmount      decimal.Decimal              `json:"maxContractAmount" description:"Max contract amount"`
	MinContractAmount      decimal.Decimal              `json:"minContractAmount" description:"min contract amount"`
	QuotaControlOverdue    string                       `json:"quotaControlOverdue" description:"0-Overdue does not affect the quota 1-Decrease in proportion to the number of overdue days 2-Frozen in full"`
	OverDueRule            []OverDueRule                `json:"overDueRule" description:"OverDue rule"`
	ListInterest           []MicroLoanInterest          `json:"listInterest" description:"List interest"`
	ListFee                []SPD0000053OFee             `json:"listFee" description:"List fee"`
	ListNotification       []PD000100ListNoti           `json:"listNotification" description:"List notification"`
	RepaymentHierarchy     []PD000070RepaymentHierarchy `json:"repaymentHierarchy"`
	Maker                  string                       `json:"maker"`
	MakerComment           string                       `json:"makerComment" description:"Maker comment"`
	UpdateFlag             string                       `json:"updateFlag" description:"Update flag"`
	Status                 string                       `json:"status" description:"status:0-inactive;1-active;2-expired;3-deleted;4-updated"`
	RevolvingFlag          string                       `json:"revolvingFlag" description:"revolving Flag:0-No;1-Yes"`

	DefaultGracePeriod        int    `json:"defaultGracePeriod" description:"Default grace period"`
	MaturityDateMethod        string `json:"maturityDateMethod" description:"maturity Date Method:01- Value Date and Loan Term;02- Same as the Last Interest Collection Date;03- Loan Date and Loan Term;04- Same as the Last Repayment Date"`
	ApplicantNumberType       string `json:"applicantNumberType" description:"applicant Number Type:0-single customer"` //0-single customer
	InstallmentRoundDirection string `json:"installmentRoundDirection" description:"installment Round Direction:0-round;1-ceil;2-floor"`
	InstallmentRoundDigits    int    `json:"installmentRoundDigits" description:"Installment round digits"`
	ValueDateMethod           string `json:"valueDateMethod" description:"value Date Method:01-Approval Date 02-Account Opening Day 03-Loan Date 04-Loan Date plus Interest-free period"` //01-Approval Date 02-Account Opening Day 03-Loan Date 04-Loan Date plus Interest-free period
}

type PD000100ListNoti struct {
	NotifyType     string           `json:"notifyType" description:"notify Type:01-Loan Account Reminder;02-Loan Account Late Notification;03-Account Statement"`
	TargetType     string           `json:"targetType" description:"target Type:0-Customer;1-Internal"`
	StrategyId     int              `json:"strategyId" description:"Strategy id"`
	EffectiveDate  string           `json:"effectiveDate" description:"Effective date"`
	ExpireDate     string           `json:"expireDate" description:"Expire date"`
	RecordType     string           `json:"recordType" description:"0-From Request"`
	StrategyName   string           `json:"strategyName" description:"Strategy name"`
	TriggerType    string           `json:"triggerType" description:"trigger Type:0-Transaction;1-Contract Status;2-Natural Period"`
	PeriodUnit     string           `json:"periodUnit" description:"period Unit:0-Day;1-Month;2-Quarter;3-Year"`
	PeriodValue    int              `json:"periodValue" description:"Period value"`
	PeriodDay      int              `json:"periodDay" description:"Period day"`
	NotifyTimeType string           `json:"notifyTimeType" description:"notify Time Type:0-Instant;1-Before;2-After"`
	TimeUnit       string           `json:"timeUnit" description:"time Unit:0-Day;1-Hour;2-Minitue;3-Second"`
	TimeValue      int              `json:"timeValue" description:"Time value"`
	IntervalUnit   string           `json:"intervalUnit" description:"interval Unit:0-Day;1-Hour;2-Minitue;3-Second"`
	IntervalValue  int              `json:"intervalValue" description:"Interval value"`
	NotifyTimes    int              `json:"notifyTimes" description:"Notify times"`
	ChannelSelect  string           `json:"channelSelect" description:"channel Select:0-Customize;1-Fixed"`
	ChannelSend    string           `json:"channelSend" description:"channel Send:0-Must;1-At Least"`
	AtLeastCount   int              `json:"atLeastCount" description:"At least count"`
	ListChannelObj []ListChannelObj `json:"listChannelObj" description:"List channel obj"`
}

func (*PD000100I) GetServiceKey() string {
	return "PD000100"
}
