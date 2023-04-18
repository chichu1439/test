package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type SPD0000053I struct {
	ProductId string `json:"productId" description:"Product id" validate:"required"`
}

type SPD0000053O struct {
	SystemId                  string                        `json:"systemId" description:"system Id:SV-Saving LN-Loan"` //SV-Saving LN-Loan
	ProductId                 string                        `json:"productId" description:"Product id"`
	Status                    string                        `json:"status" description:"product status:0-inactive;1-active;2-expired;3-deleted;4-updated"`
	ProductName               string                        `json:"productName" description:"Product name"`
	CustomerType              string                        `json:"customerType" description:"customer Type:0-personal;1-corporate"`
	ProductType               string                        `json:"productType" description:"product Type:01: 01-Micro Loan 02-Revolving Loan,03-BNPL"`
	EffectiveDate             string                        `json:"effectiveDate" description:"Effective date"`
	ExpireDate                string                        `json:"expireDate" description:"Expire date"`
	Remark                    string                        `json:"remark" description:"Remark"`
	GuaranteeMethod           string                        `json:"guaranteeMethod" description:"guarantee Method:0-No Guarantee 1-Guarantee"` //0-No Guarantee 1-Guarantee
	Currency                  string                        `json:"currency" description:"Currency"`
	MaxContractCount          int                           `json:"maxContractCount" description:"Max contract count"`
	ValueDateMethod           string                        `json:"valueDateMethod" description:"value Date Method:01-Approval Date 02-Account Opening Day 03-Loan Date 04-Loan Date plus Interest-free period"` //01-Approval Date 02-Account Opening Day 03-Loan Date 04-Loan Date plus Interest-free period
	InstallmentList           string                        `json:"installmentList" description:"Installment list"`
	RepaymentFrequency        string                        `json:"repaymentFrequency" description:"repayment Frequency:01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly"`
	RepaymentMethod           string                        `json:"repaymentMethod" description:"repayment Method:01-Fixed Installment;02-Fixed Principal"`
	RepaymentOrderMethod      string                        `json:"repaymentOrderMethod" description:"repayment Order Method:0-period first;1-amount first"`
	PrincipalRule             []PrincipalRule               `json:"principalRule" description:"Principal rule"`
	InstallmentMethod         string                        `json:"installmentMethod" description:"installment Method:01-Fixed Installment;02-Fixed Principal"`
	AllowEarlyRepayment       string                        `json:"allowEarlyRepayment" validate:"required" description:"allow Early Repayment:0-no;1-yes"`
	MaxEarlyRepaymentTimes    int                           `json:"maxEarlyRepaymentTimes" description:"Max early repayment times"`
	AllowGracePeriod          string                        `json:"allowGracePeriod" validate:"required" description:"allow Grace Period:0-no;1-yes"`
	MaxGracePeriod            int                           `json:"maxGracePeriod" description:"Max grace period"`
	ImpairedDays              int                           `json:"impairedDays" description:"Impaired days"`
	QuotaControlOverdue       string                        `json:"quotaControlOverdue" description:"0-Overdue does not affect the quota 1-Decrease in proportion to the number of overdue days 2-Frozen in full"`
	OverDueRule               []OverDueRule                 `json:"overDueRule" description:"OverDue rule"`
	ListInterest              []SPD0000053OInterest         `json:"listInterest" description:"List interest"`
	ListNotification          []SPD0000053OListNoti         `json:"listNotification" description:"List notification"`
	MakerComment              string                        `json:"makerComment" description:"Maker comment"`
	Maker                     string                        `json:"maker" description:"Maker user"`
	ApproveComment            string                        `json:"approveComment" description:"approve comment"`
	Checker                   string                        `json:"checker" description:"Checker user"`
	ApproveView               string                        `json:"approveView" description:"Approve Status(0-awaiting;1-approved;2-rejected)"`
	RevolvingFlag             string                        `json:"revolvingFlag" description:"revolving Flag:0-no;1-yes"`
	MaxContractAmount         decimal.Decimal               `json:"maxContractAmount" description:"Max contract amount"`
	MinContractAmount         decimal.Decimal               `json:"minContractAmount" description:"min contract amount"`
	MaxLoanAmount             decimal.Decimal               `json:"maxLoanAmount" description:"Max loan amount"`
	MinLoanAmount             decimal.Decimal               `json:"minLoanAmount" description:"Min loan amount"`
	Version                   string                        `json:"version" description:"Version"`
	MaturityDateMethod        string                        `json:"maturityDateMethod" description:"maturity Date Method:01- Value Date and Loan Term;02- Same as the Last Interest Collection Date;03- Loan Date and Loan Term;04- Same as the Last Repayment Date"`
	ApplicantNumberType       string                        `json:"applicantNumberType" description:"applicant Number Type:0-single customer"` //0-single customer
	DefaultGracePeriod        int                           `json:"defaultGracePeriod" description:"Default grace period"`
	ListCategory              []ProdCategoryNameBo          `json:"listCategory" description:"List category"`
	MaxLoanPeriods            int                           `json:"maxLoanPeriods" description:"Max loan periods"`
	MinLoanPeriods            int                           `json:"minLoanPeriods" description:"Min loan periods"`
	InstallmentRoundDirection string                        `json:"installmentRoundDirection" description:"Installment round direction"`
	InstallmentRoundDigits    int                           `json:"installmentRoundDigits" description:"Installment round digits"`
	CompoundInterestMethod    string                        `json:"compoundInterestMethod" description:"0-simple interest;1-compound interest"`
	ReleaseQuotaMethod        string                        `json:"releaseQuotaMethod" description:"release Quota Method:0-Each repayment;1-loan closed"`
	FirstRepayDateMethod      string                        `json:"firstRepayDateMethod"` // 0-Automatic；1-Manual
	GraceType                 string                        `json:"graceType"`
	RepaymentHierarchy        []MicroLoanRepaymentHierarchy `json:"repaymentHierarchy"`
	ListFee                   []SPD0000053OFee              `json:"listFee" description:"List fee"`
	DueRuleRule               []DueRuleRule                 `json:"dueRuleRule"`
}
type MicroLoanRepaymentHierarchy struct {
	RepaymentHierarchyType string `json:"repaymentHierarchyType" description:"repayment hierarchy type" `
	RepaymentHierarchyName string `json:"repaymentHierarchyName"`
	Behavior               string `json:"behavior"`
	RepaymentOrder         string `json:"repaymentOrder" description:"repayment order ,format； repayment priority：repayment Hierarchy" `
	Code                   string `json:"code"`
	Description            string `json:"description"`
}
type SPD0000053OFee struct {
	FeeId           string                  `json:"feeId"`
	FeeTypeId       string                  `json:"feeTypeId"`
	TransactionType string                  `json:"transactionType"`
	CvOnTop         string                  `json:"cvOnTop"`
	SystemId        string                  `json:"systemId"`
	Currency        string                  `json:"currency"`
	MaxFeeAmount    decimal.Decimal         `json:"maxFeeAmount"`
	MinFeeAmount    decimal.Decimal         `json:"minFeeAmount"`
	FeeCalcMethod   string                  `json:"feeCalcMethod"`
	Remark          string                  `json:"remark"`
	FeeTypeName     string                  `json:"feeTypeName"`
	AccountingCode  string                  `json:"accountingCode"`
	IsAmortization  string                  `json:"isAmortization"`
	FeeName         string                  `json:"feeName"`
	CalcBasisType   string                  `json:"calcBasisType"`
	PayerType       string                  `json:"payerType"`
	FinalModifyDate time.Time               `json:"finalModifyDate"`
	FinalModifyOrgz string                  `json:"finalModifyOrgz"`
	FinalModifyTime time.Time               `json:"finalModifyTime"`
	FinalModifyUser string                  `json:"finalModifyUser"`
	FeeRules        []PF000002OFeeRulesItem `json:"feeRules"`
}
type DueRuleRule struct {
	StartDay string `json:"startDay"`
	EndDay   string `json:"endDay"`
	RepayDay string `json:"repayDay"`
	DueMonth string `json:"dueMonth"`
}
type SPD0000053OListNoti struct {
	Uid            int              `json:"uid" description:"Uid"`
	NotifyType     string           `json:"notifyType" description:"01-Loan Account Reminder;02-Loan Account Late Notification;03-Account Statement"`
	TargetType     string           `json:"targetType" description:"0-Customer;1-Internal"`
	StrategyId     int              `json:"strategyId" description:"Strategy id"`
	EffectiveDate  string           `json:"effectiveDate" description:"Effective date"`
	ExpireDate     string           `json:"expireDate" description:"Expire date"`
	StrategyName   string           `json:"strategyName" description:"Strategy name"`
	TriggerType    string           `json:"triggerType" description:"0-Transaction;1-Contract Status;2-Natural Period"`
	PeriodUnit     string           `json:"periodUnit" description:"0-Day;1-Month;2-Quarter;3-Year"`
	PeriodValue    int              `json:"periodValue" description:"Period value"`
	PeriodDay      int              `json:"periodDay" description:"Period day"`
	NotifyTimeType string           `json:"notifyTimeType" description:"0-Instant;1-Before;2-After"`
	TimeUnit       string           `json:"timeUnit" description:"Time unit"`
	TimeValue      int              `json:"timeValue" description:"Time value"`
	IntervalUnit   string           `json:"intervalUnit" description:"0-Day;1-Hour;2-Minitue;3-Second"`
	IntervalValue  int              `json:"intervalValue" description:"Interval value"`
	NotifyTimes    int              `json:"notifyTimes" description:"Notify times"`
	ChannelSelect  string           `json:"channelSelect" description:"0-Customize;1-Fixed"`
	ChannelSend    string           `json:"channelSend" description:"0-Must;1-At Least"`
	AtLeastCount   int              `json:"atLeastCount" description:"At least count"`
	ListChannelObj []ListChannelObj `json:"listChannelObj" description:"List channel obj"`
}

type ListChannelObj struct {
	ChannelType     string `json:"channelType" description:"channel Type:01-Email;02-SMS;03-Site;04-Mobile App;05-Call"`
	TemplateId      int    `json:"templateId" description:"Template id"`
	RelateMcId      string `json:"relateMcId" description:"Relate mc id"`
	TemplateName    string `json:"templateName" description:"Template name"`
	TemplateContent string `json:"templateContent" description:"Template content"`
	MustSend        string `json:"mustSend" description:"mustSend:0-no;1-yes"`
	Order           int    `json:"order" description:"Order"`
	RetryTimes      int    `json:"retryTimes" description:"Retry times"`
}

type SPD0000053OInterest struct {
	InterestType      string                        `json:"interestType" description:"interestType:0-normal loan 1-overdue loan"` //0-normal loan 1-overdue loan
	CustomerGradeType string                        `json:"customerGradeType"`
	UseInterestFlag   bool                          `json:"useInterestFlag"`
	StrategyList      []SPD0000053OInterestStrategy `json:"strategyList" description:"List interest strategy"`
}

type SPD0000053OInterestStrategy struct {
	InterestUid        int             `json:"interestUid" description:"interestUid"`
	InterestId         string          `json:"interestId" description:"Interest id"`
	CustomerGrade      string          `json:"customerGrade" description:"0-All customers 1-Level 1 2-Level 2 3-Level 3"`
	EffectiveDate      string          `json:"effectiveDate" description:"Effective date"`
	ExpireDate         string          `json:"expireDate" description:"Expire date"`
	StrategyUid        int             `json:"strategyUid" description:"strategyUid"`
	StrategyId         string          `json:"strategyId" description:"Strategy id"`
	BaseInterestRate   decimal.Decimal `json:"baseInterestRate" description:"Base interest rate"`
	MinInterestRate    decimal.Decimal `json:"minInterestRate" description:"Min interest rate"`
	MaxInterestRate    decimal.Decimal `json:"maxInterestRate" description:"Max interest rate"`
	FloatType          string          `json:"floatType" description:"float Type:0-By Percent 1-BP Value"`             //0-By Percent 1-BP Value
	FloatDirection     string          `json:"floatDirection" description:"float Direction:0-None 1-Up 2-Down 3-Both"` //0-None 1-Up 2-Down 3-Both
	MinFloatValue      decimal.Decimal `json:"minFloatValue" description:"Min float value"`
	MaxFloatValue      decimal.Decimal `json:"maxFloatValue" description:"Max float value"`
	CalculateType      string          `json:"calculateType" description:"calculate Type:0-Actual days;1-Year to month to day;2-Year to month"`
	RateExecType       string          `json:"rateExecType" description:"rate Exec Type:0-By Value Date;1-By Calc Date"`
	FloatExecType      string          `json:"floatExecType" description:"float Exec Type:0-By Value Date;1-By Calc Date;2-Fixed float"`
	MonthDays          string          `json:"monthDays" description:"month Days:0-30;1-Actual days"`
	YearDays           string          `json:"yearDays" description:"year Days:0-360;1-365;2-Actual days"`
	DecimalFlag        string          `json:"decimalFlag" description:"decimal Flag:0-No;1-Yes""`
	AccrualPeriodType  string          `json:"accrualPeriodType" description:"accrual Period Type:01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly"`
	AccrualPeriodValue int             `json:"accrualPeriodValue" description:"Accrual period value"`
	AccrualPeriodDay   int             `json:"accrualPeriodDay" description:"Accrual period day"`
	SettlePeriodType   string          `json:"settlePeriodType" description:"settle Period Type:01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly"`
	SettlePeriodValue  int             `json:"settlePeriodValue" description:"Settle period value"`
	SettlePeriodDay    int             `json:"settlePeriodDay" description:"Settle period day"`
}

func (*SPD0000053I) GetServiceKey() string {
	return "PD000053"
}
