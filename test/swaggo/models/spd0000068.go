package models

import "github.com/shopspring/decimal"

type PD000068I struct {
	// basis
	SystemId      string `json:"systemId" validate:"required,max=2" description:"system Id:SV-Saving LN-Loan"`
	ProductName   string `json:"productName" description:"Product name" validate:"required"`
	CustomerType  string `json:"customerType" description:"Customer type"`
	ProductType   string `json:"productType" validate:"required,max=2" description:"product Type:*Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan"`
	EffectiveDate string `json:"effectiveDate" description:"Effective date" validate:"required"`
	ExpireDate    string `json:"expireDate" description:"Expire date" validate:"required"`
	Remark        string `json:"remark" description:"Remark"`
	// Guarantee
	GuaranteeMethod string `json:"guaranteeMethod" validate:"required" description:"guarantee Method:0-No Guarantee 1-Guarantee"`
	// product detail
	Currency           string `json:"currency" description:"Currency" validate:"required"`
	MaxContractCount   int    `json:"maxContractCount" description:"Max contract count"`
	InstallmentMethod  string `json:"installmentMethod" description:"Installment method"`
	InstallmentList    string `json:"installmentList" description:"Installment list"`
	RepaymentFrequency string `json:"repaymentFrequency" description:"repayment Frequency:1-Daily; 2-Weekly; 3-Double Week; 4-Monthly; 5-Quarterly; 6-Harf a Year; 7-Yearly"`
	// repayment
	FirstRepayDateMethod string `json:"firstRepayDateMethod"` // 0-Automatic；1-Manual
	RepaymentMethod      string `json:"repaymentMethod" description:"repayment Method:01-Equal Installment 02-Equal Principal; 03-Principal and Interest Installment on Maturity Day 04-Interest Installment but Principal on Maturity Day; 06-Independent Bill Date；"`
	RepaymentOrderMethod string `json:"repaymentOrderMethod" description:"repayment Order Method:0-period first;1-amount first"`
	// bill payment
	PrincipalRule []PrincipalRule `json:"principalRule" description:"Principal rule"` //0-不影响额度 1-按逾期天数比例调减 2-全额冻结
	// Compound Interest
	CompoundInterestMethod string `json:"compoundInterestMethod" description:"compound Interest Method:0-simple interest;1-compound interest"`
	// loan Periods
	MaxLoanPeriods int `json:"maxLoanPeriods" description:"Max loan periods"`
	MinLoanPeriods int `json:"minLoanPeriods" description:"Min loan periods"`
	// Early Repayment
	AllowEarlyRepayment    string `json:"allowEarlyRepayment" validate:"required" description:"allow Early Repayment:0-no 1-yes"` //0-no 1-yes
	MaxEarlyRepaymentTimes int    `json:"maxEarlyRepaymentTimes" description:"Max early repayment times"`
	// Late Repayment
	AllowGracePeriod string `json:"allowGracePeriod" validate:"required" description:"allow Grace Period:0-no 1-yes"` //0-no 1-yes
	GraceType        string `json:"graceType"`
	MaxGracePeriod   int    `json:"maxGracePeriod" description:"Max grace period"`
	ImpairedDays     int    `json:"impairedDays" description:"Impaired days"`
	// Amount
	ReleaseQuotaMethod string          `json:"releaseQuotaMethod" description:"release Quota Method:0-Each repayment;1-loan closed"` // Revolving Loan
	MaxLoanAmount      decimal.Decimal `json:"maxLoanAmount" description:"Max loan amount"`
	MinLoanAmount      decimal.Decimal `json:"minLoanAmount" description:"Min loan amount"`
	MaxContractAmount  decimal.Decimal `json:"maxContractAmount" description:"Max contract amount"`
	MinContractAmount  decimal.Decimal `json:"minContractAmount" description:"min contract amount"`
	// Overdue Credit Line Control
	QuotaControlOverdue string        `json:"quotaControlOverdue" description:"0-Overdue does not affect the quota 1-Decrease in proportion to the number of overdue days 2-Frozen in full"`
	OverDueRule         []OverDueRule `json:"overDueRule" description:"OverDue rule"`
	// interest
	ListInterest []MicroLoanInterest02 `json:"listInterest" description:"List interest" validate:"dive"`
	// fee
	ListFee []MicroLoanFee `json:"listFee" description:"List fee"`
	// notification
	ListNotification []PD000068ListNoti `json:"listNotification" description:"List notification"`
	// repayment hierarchy
	RepaymentHierarchy []PD000070RepaymentHierarchy `json:"repaymentHierarchy"`
	// maker
	MakerComment  string `json:"makerComment" description:"Maker comment"`
	Version       string `json:"version" description:"Version"`
	RevolvingFlag string `json:"revolvingFlag" description:"0-no;1-yes"`
	// dont not where
	DefaultGracePeriod        int    `json:"defaultGracePeriod" description:"Default grace period"`
	MaturityDateMethod        string `json:"maturityDateMethod" description:"maturity Date Method:01- Value Date and Loan Term 02- Same as the Last Interest Collection Date 03- Loan Date and Loan Term 04- Same as the Last Repayment Date"` //
	ApplicantNumberType       string `json:"applicantNumberType" description:"applicant Number Type:0-single customer"`
	InstallmentRoundDirection string `json:"installmentRoundDirection" description:"Installment round direction"`
	InstallmentRoundDigits    int    `json:"installmentRoundDigits" description:"Installment round digits"`
	ValueDateMethod           string `json:"valueDateMethod" description:"value Date Method:01-Approval Date 02-Account Opening Day 03-Loan Date 04-Loan Date plus Interest-free period"` //01-Approval Date 02-Account Opening Day 03-Loan Date 04-Loan Date plus Interest-free period
}

type PD000068ListNoti struct {
	NotifyType string `json:"notifyType" description:"notify Type:01-Loan Account Reminder;02-Loan Account Late Notification;03-Account Statement"`
	TargetType string `json:"targetType" description:"target Type:0-Customer;1-Internal"`
	StrategyId int    `json:"strategyId" description:"Strategy id"`
}

type OverDueRule struct {
	OverdueDays int `json:"overdueDays" description:"Overdue days"`
	//当逾期额度控制为“1-按逾按逾期天数调减”时必输，在每个日期后输入相应的值
	ReducePercent decimal.Decimal `json:"reducePercent" description:"Reduce percent"`
}

type PrincipalRule struct {
	//当还款方式为“5-账单日还款”时必输。
	//When the repayment method is "5-bill day repayment"，must be entered .
	BillPaymentType string `json:"billPaymentType" description:"0-Percentage of principal 1- Add or subtract a fixed amount based on the bill interest 2- by bill interest"`
	//当最低还款额类型为“0-按本金百分比时，1-按账单利息加减固定金额”，时必输。如果选“0-All customers”只需输入一个值（百分比值或固定金额值）适用所有客户。
	//如果不选“0-All customers”时每一个等级需设定一个值 （百分比值或固定金额值）。0与1，2，3不能同时选择。
	//When the minimum repayment type is "0-based on principal percentage, 1-based on bill interest plus or minus fixed amount",
	//it must be entered. If you select "0-All customers", you only need to enter a value (percentage value or fixed amount value) to apply to all customers.
	//If "0-All customers" is not selected, a value (percentage value or fixed amount value) must be set for each level.0 and 1, 2, 3 cannot be selected at the same time.
	CustomerGrade string `json:"customerGrade" description:"0-All customers 1-Level 1 2-Level 2 3-Level 3"`
	//当最低还款额类型为“0-按本金百分比”，时必输。
	PercentageValue decimal.Decimal `json:"percentageValue" description:"When the minimum repayment type is  0-Percentage of principal , it must be entered."`
	//当最低还款额类型为“1-账单利息加减固定金额”时必输
	FixedAmountValue decimal.Decimal `json:"fixedAmountValue" description:"When the minimum repayment amount type is  1-Bill interest plus or minus fixed amount , it must be entered."`
}

type MicroLoanInterest02 struct {
	InterestType string                    `json:"interestType" validate:"required" description:"interest Type:0-normal 1-Withdrawal in Advance"` //0-normal 1-Withdrawal in Advance
	StrategyList []InterestAndStrategyList `json:"strategyList"`
}
type InterestAndStrategyList struct {
	CustomerGrade    string          `json:"customerGrade" description:"0-All customers 1-Level 1 2-Level 2 3-Level 3"`
	BaseInterestRate decimal.Decimal `json:"baseInterestRate" description:"Base interest rate"`
	MinInterestRate  decimal.Decimal `json:"minInterestRate" description:"Min interest rate"`
	MaxInterestRate  decimal.Decimal `json:"maxInterestRate" description:"Max interest rate"`
	FloatDirection   string          `json:"floatDirection" description:"Float direction"`
	FloatType        string          `json:"floatType" description:"Float type"`
	MinFloatValue    decimal.Decimal `json:"minFloatValue" description:"Min float value"`
	MaxFloatValue    decimal.Decimal `json:"maxFloatValue" description:"Max float value"`
	YearDays         string          `json:"yearDays" description:"Year days"`
	InterestId       string          `json:"interestId"`
}
type MicroLoanFee struct {
	FeeId           string `json:"feeId"`
	FeeTypeId       string `json:"feeTypeId"`
	TransactionType string `json:"transactionType"`
	CvOnTop         string `json:"cvOnTop"`
}

type PD000068O struct {
	ProductId string `json:"productId" description:"Product id"`
}

func (*PD000068I) GetServiceKey() string {
	return "PD000068"
}
