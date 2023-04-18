package models

import "github.com/shopspring/decimal"

type SPD0000012I struct {
	ProductId string `json:"productId" description:"Product id" validate:"required"`
}

type SPD0000012O struct {
	SystemId                  string          `json:"systemId" validate:"required" description:"system Id:LN-Loan;SV-Saving"`
	ProductId                 string          `json:"productId" description:"Product id"`                                                                 //
	ProductType               string          `json:"productType" description:"product Type:Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan"` //*Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan
	ProductName               string          `json:"productName" description:"Product name"`                                                             //
	Currency                  string          `json:"currency" description:"Currency"`                                                                    //
	Version                   string          `json:"version" description:"Version"`                                                                      //
	Remark                    string          `json:"remark" description:"Remark"`                                                                        //
	EffectiveDate             string          `json:"effectiveDate" description:"Effective date"`                                                         //
	ExpireDate                string          `json:"expireDate" description:"Expire date"`                                                               //
	MaturityDateMethod        string          `json:"maturityDateMethod" description:"maturity Date Method:01- Value Date and Loan Term;02- Same as the Last Interest Collection Date;03- Loan Date and Loan Term;04- Same as the Last Repayment Date"`
	GuaranteeMethod           string          `json:"guaranteeMethod" description:"guarantee Method:0-No Guarantee 1-Guarantee"` //0-No Guarantee 1-Guarantee
	MaxContractCount          int             `json:"maxContractCount" description:"Max contract count"`
	ValueDateMethod           string          `json:"valueDateMethod" description:"value Date Method:01-Approval Date 02-Account Opening Day 03-Loan Date 04-Loan Date plus Interest-free period"` //01-Approval Date 02-Account Opening Day 03-Loan Date 04-Loan Date plus Interest-free period
	InstallmentList           string          `json:"installmentList" description:"Installment list"`
	RepaymentFrequency        string          `json:"repaymentFrequency" description:"repayment Frequency:01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly"`
	RepaymentMethod           string          `json:"repaymentMethod" description:"repayment Method:01-Fixed Installment;02-Fixed Principal"`
	RepaymentOrderMethod      string          `json:"repaymentOrderMethod" description:"repayment Order Method:0-period first;1-amount first"`
	PrincipalRule             []PrincipalRule `json:"principalRule" description:"Principal rule"`
	InstallmentMethod         string          `json:"installmentMethod" description:"installment Method:01-Fixed Installment;02-Fixed Principal"`
	AllowEarlyRepayment       string          `json:"allowEarlyRepayment" validate:"required" description:"allow Early Repayment:0-no;1-yes"`
	MaxEarlyRepaymentTimes    int             `json:"maxEarlyRepaymentTimes" description:"Max early repayment times"`
	AllowExtension            string          `json:"allowExtension" validate:"required" description:"allow Extension:0-no;1-yes"`
	MaxExtensionTimes         int             `json:"maxExtensionTimes" description:"Max extension times"`
	TotalExtensionDaysOption  string          `json:"totalExtensionDaysOption" description:"Total extension days option"`
	AllowGracePeriod          string          `json:"allowGracePeriod" validate:"required" description:"allow Grace Period:0-no;1-yes"`
	MaxGracePeriod            int             `json:"maxGracePeriod" description:"Max grace period"`
	ImpairedDays              int             `json:"impairedDays" description:"Impaired days"`
	QuotaControlOverdue       string          `json:"quotaControlOverdue" description:"0-Overdue does not affect the quota 1-Decrease in proportion to the number of overdue days 2-Frozen in full"`
	OverDueRule               []OverDueRule   `json:"overDueRule" description:"OverDue rule"`
	MakerComment              string          `json:"makerComment" description:"Maker comment"`
	Maker                     string          `json:"maker" description:"Maker user"`
	ApproveComment            string          `json:"approveComment" description:"approve comment"`
	Checker                   string          `json:"checker" description:"Checker user"`
	ApproveView               string          `json:"approveView" description:"Approve Status(0-awaiting;1-approved;2-rejected)"`
	RevolvingFlag             string          `json:"revolvingFlag" description:"revolving Flag:0-no;1-yes"`
	MaxContractAmount         decimal.Decimal `json:"maxContractAmount" description:"Max contract amount"`
	MinContractAmount         decimal.Decimal `json:"minContractAmount" description:"min contract amount"`
	MaxLoanAmount             decimal.Decimal `json:"maxLoanAmount" description:"Max loan amount"`
	MinLoanAmount             decimal.Decimal `json:"minLoanAmount" description:"Min loan amount"`
	ApplicantNumberType       string          `json:"applicantNumberType" description:"applicant Number Type:0-single customer"` //0-single customer
	DefaultGracePeriod        int             `json:"defaultGracePeriod" description:"Default grace period"`
	MaxLoanPeriods            int             `json:"maxLoanPeriods" description:"Max loan periods"`
	MinLoanPeriods            int             `json:"minLoanPeriods" description:"Min loan periods"`
	InstallmentRoundDirection string          `json:"installmentRoundDirection" description:"Installment round direction"`
	InstallmentRoundDigits    int             `json:"installmentRoundDigits" description:"Installment round digits"`
	CompoundInterestMethod    string          `json:"compoundInterestMethod" description:"0-simple interest;1-compound interest"`
	ReleaseQuotaMethod        string          `json:"releaseQuotaMethod" description:"release Quota Method:0-Each repayment;1-loan closed"`
	RepaymentOrderAmount      string          `json:"repaymentOrderAmount" description:"Repayment order amount"`
	Status                    string          `json:"status" description:"status:0-inactive;1-active;2-expired;3-deleted;4-updated"`
}

func (*SPD0000012I) GetServiceKey() string {
	return "PD000012"
}
