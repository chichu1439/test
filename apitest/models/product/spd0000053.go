package models

import "github.com/shopspring/decimal"

type SPD0000053I struct {
	ProductId string `json:"productId" description:"Product id" validate:"required"`
}

type SPD0000053O struct {
	SystemId                  string                `json:"systemId" description:"system Id:SV-Saving LN-Loan"` //SV-Saving LN-Loan
	ProductId                 string                `json:"productId" description:"Product id"`
	ProductType               string                `json:"productType" description:"product Type:Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan"` //*Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan
	CustomerType              string                `json:"customerType" description:"customer Type:0-personal;1-corporate"`
	ProductName               string                `json:"productName" description:"Product name"`
	Currency                  string                `json:"currency" description:"Currency"`
	Version                   string                `json:"version" description:"Version"`
	Remark                    string                `json:"remark" description:"Remark"`
	EffectiveDate             string                `json:"effectiveDate" description:"Effective date"`
	ExpireDate                string                `json:"expireDate" description:"Expire date"`
	ValueDateMethod           string                `json:"valueDateMethod" description:"value Date Method:01-Approval Date 02-Account Opening Day 03-Loan Date 04-Loan Date plus Interest-free period"` //01-Approval Date 02-Account Opening Day 03-Loan Date 04-Loan Date plus Interest-free period
	MaturityDateMethod        string                `json:"maturityDateMethod" description:"maturity Date Method:01- Value Date and Loan Term;02- Same as the Last Interest Collection Date;03- Loan Date and Loan Term;04- Same as the Last Repayment Date"`
	RepaymentMethod           string                `json:"repaymentMethod" description:"repayment Method:01-Fixed Installment;02-Fixed Principal"`
	RepaymentFrequency        string                `json:"repaymentFrequency" description:"repayment Frequency:01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly"`
	AllowEarlyRepayment       string                `json:"allowEarlyRepayment" validate:"required" description:"allow Early Repayment:0-no;1-yes"`
	AllowExtension            string                `json:"allowExtension" validate:"required" description:"allow Extension:0-no;1-yes"`
	AllowGracePeriod          string                `json:"allowGracePeriod" validate:"required" description:"allow Grace Period:0-no;1-yes"`
	GuaranteeMethod           string                `json:"guaranteeMethod" description:"guarantee Method:0-No Guarantee 1-Guarantee"` //0-No Guarantee 1-Guarantee
	Status                    string                `json:"status" description:"status:0-inactive;1-active;2-expired;3-deleted;4-updated"`
	ApplicantNumberType       string                `json:"applicantNumberType" description:"applicant Number Type:0-single customer"` //0-single customer
	DefaultGracePeriod        int                   `json:"defaultGracePeriod" description:"Default grace period"`
	ListCategory              []ProdCategoryNameBo  `json:"listCategory" description:"List category"`
	ListInterest              []SPD0000053OInterest `json:"listInterest" description:"List interest"`
	ListFee                   []ProdFeeService      `json:"listFee" description:"List fee"`
	ListNotification          []SPD0000053OListNoti `json:"listNotification" description:"List notification"`
	MaxGracePeriod            int                   `json:"maxGracePeriod" description:"Max grace period"`
	MaxLoanAmount             float64               `json:"maxLoanAmount" description:"Max loan amount"`
	MinLoanAmount             float64               `json:"minLoanAmount" description:"Min loan amount"`
	MaxExtensionTimes         int                   `json:"maxExtensionTimes" description:"Max extension times"`
	TotalExtensionDaysOption  string                `json:"totalExtensionDaysOption" description:"Total extension days option"`
	MaxLoanPeriods            int                   `json:"maxLoanPeriods" description:"Max loan periods"`
	MinLoanPeriods            int                   `json:"minLoanPeriods" description:"Min loan periods"`
	RevolvingFlag             string                `json:"revolvingFlag" description:"revolving Flag:0-no;1-yes"`
	InstallmentMethod         string                `json:"installmentMethod" description:"installment Method:01-Fixed Installment;02-Fixed Principal"`
	InstallmentRoundDirection string                `json:"installmentRoundDirection" description:"Installment round direction"`
	InstallmentRoundDigits    int                   `json:"installmentRoundDigits" description:"Installment round digits"`
	CompoundInterestMethod    string                `json:"compoundInterestMethod" description:"0-simple interest;1-compound interest"`
	ImpairedDays              int                   `json:"impairedDays" description:"Impaired days"`
	MaxQuotaAmount            float64               `json:"maxQuotaAmount" description:"Max quota amount"`
	ReleaseQuotaMethod        string                `json:"releaseQuotaMethod" description:"release Quota Method:0-Each repayment;1-loan closed"`
	RepaymentOrderMethod      string                `json:"repaymentOrderMethod" description:"repayment Order Method:0-period first;1-amount first"`
	RepaymentOrderAmount      string                `json:"repaymentOrderAmount" description:"Repayment order amount"`
	MaxContractCount          int                   `json:"maxContractCount" description:"Max contract count"`
	MaxEarlyRepaymentTimes    int                   `json:"maxEarlyRepaymentTimes" description:"Max early repayment times"`
	InstallmentList           string                `json:"installmentList" description:"Installment list"`
	ApproveStatus             string                `json:"approveStatus" description:"Approve Status(0-awaiting;1-approved;2-rejected)"`
	MakerComment              string                `json:"makerComment" description:"Maker comment"`
	MakerUser                 string                `json:"makerUser" description:"Maker user"`
	CheckerComment            string                `json:"checkerComment" description:"Checker comment"`
	CheckerUser               string                `json:"checkerUser" description:"Checker user"`
}

type SPD0000053OListNoti struct {
	Uid            string           `json:"uid" description:"Uid"`
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
	Uid                  int                           `json:"uid" description:"Uid"`
	InterestType         string                        `json:"interestType" description:"interestType:0-normal loan 1-overdue loan"` //0-normal loan 1-overdue loan
	InterestId           string                        `json:"interestId" description:"Interest id"`
	EffectiveDate        string                        `json:"effectiveDate" description:"Effective date"`
	ExpireDate           string                        `json:"expireDate" description:"Expire date"`
	ListInterestStrategy []SPD0000053OInterestStrategy `json:"listInterestStrategy" description:"List interest strategy"`
}

type SPD0000053OInterestStrategy struct {
	Uid                int     `json:"uid" description:"Uid"`
	StrategyId         string  `json:"strategyId" description:"Strategy id"`
	BaseInterestRate   float64 `json:"baseInterestRate" description:"Base interest rate"`
	MinInterestRate    float64 `json:"minInterestRate" description:"Min interest rate"`
	MaxInterestRate    float64 `json:"maxInterestRate" description:"Max interest rate"`
	FloatType          string  `json:"floatType" description:"float Type:0-By Percent 1-BP Value"`             //0-By Percent 1-BP Value
	FloatDirection     string  `json:"floatDirection" description:"float Direction:0-None 1-Up 2-Down 3-Both"` //0-None 1-Up 2-Down 3-Both
	MinFloatValue      float64 `json:"minFloatValue" description:"Min float value"`
	MaxFloatValue      float64 `json:"maxFloatValue" description:"Max float value"`
	CaculateType       string  `json:"caculateType" description:"caculate Type:0-Actual days;1-Year to month to day;2-Year to month"`
	RateExecType       string  `json:"rateExecType" description:"rate Exec Type:0-By Value Date;1-By Calc Date"`
	FloatExecType      string  `json:"floatExecType" description:"float Exec Type:0-By Value Date;1-By Calc Date;2-Fixed float"`
	MonthDays          string  `json:"monthDays" description:"month Days:0-30;1-Actual days"`
	YearDays           string  `json:"yearDays" description:"year Days:0-360;1-365;2-Actual days"`
	DecimalFlag        string  `json:"decimalFlag" description:"decimal Flag:0-No;1-Yes""`
	AccrualPeriodType  string  `json:"accrualPeriodType" description:"accrual Period Type:01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly"`
	AccrualPeriodValue int     `json:"accrualPeriodValue" description:"Accrual period value"`
	AccrualPeriodDay   int     `json:"accrualPeriodDay" description:"Accrual period day"`
	SettlePeriodType   string  `json:"settlePeriodType" description:"settle Period Type:01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly"`
	SettlePeriodValue  int     `json:"settlePeriodValue" description:"Settle period value"`
	SettlePeriodDay    int     `json:"settlePeriodDay" description:"Settle period day"`
}

type ProdFeeService struct {
	Uid         int           `json:"uid" description:"Uid"`
	ServiceId   string        `json:"serviceId" description:"Service id"`
	FeatureId   string        `json:"featureId" description:"Feature id"`
	ServiceName string        `json:"serviceName" description:"Service name"`
	ListFeeType []ProdFeeItem `json:"listFeeType" description:"List fee type"`
}

type ProdFeeItem struct {
	Uid           int             `json:"uid" description:"Uid"`
	FeeId         string          `json:"feeId" description:"Fee id"`
	FeeName       string          `json:"feeName" description:"Fee name"`
	MaxFee        decimal.Decimal `json:"maxFee" description:"Max fee"`
	MinFee        decimal.Decimal `json:"minFee" description:"Min fee"`
	FeeCalcMethod string          `json:"feeCalcMethod" description:"Fee calc method"` //0-Percent 1-Fixed Amount
	CalcValue     decimal.Decimal `json:"calcValue" description:"Calc value"`
	PayerType     string          `json:"payerType" description:"payer Type:01-Payee Participant;02-Payer Participant" :"payer_type"`
	EffectiveDate string          `json:"effectiveDate" description:"Effective date"`
	ExpireDate    string          `json:"expireDate" description:"Expire date"`
}

func (*SPD0000053I) GetServiceKey() string {
	return "spd0000053"
}
