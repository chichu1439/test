package models

import "github.com/shopspring/decimal"

type PD000097I struct {
	ProductIdList []string `json:"productIdList" description:"Product id list"`
}

type PD000097O struct {
	ListMircoLoanProduct []MircoLoanProductInfo `json:"listMircoLoanProduct" description:"List mirco loan product"`
}

type MircoLoanProductInfo struct {
	ProductId                 string          `json:"productId" description:"Product id"`
	ValueDateMethod           string          `json:"valueDateMethod" description:"Value date method"`
	MaturityDateMethod        string          `json:"maturityDateMethod" description:"Maturity date method"`
	RepaymentMethod           string          `json:"repaymentMethod" description:"Repayment method"`
	RepaymentFrequency        string          `json:"repaymentFrequency" description:"Repayment frequency"`
	GuaranteeMethod           string          `json:"guaranteeMethod" description:"Guarantee method"`
	ApplicantNumberType       string          `json:"applicantNumberType" description:"Applicant number type"`
	AllowEarlyRepayment       string          `json:"allowEarlyRepayment" description:"Allow early repayment"`
	AllowExtension            string          `json:"allowExtension" description:"Allow extension"`
	AllowGracePeriod          string          `json:"allowGracePeriod" description:"Allow grace period"`
	DefaultGracePeriod        int             `json:"defaultGracePeriod" description:"Default grace period"`
	MaxGracePeriod            int             `json:"maxGracePeriod" description:"Max grace period"`
	MaxLoanAmount             decimal.Decimal `json:"maxLoanAmount" description:"Max loan amount"`
	MinLoanAmount             decimal.Decimal `json:"minLoanAmount" description:"Min loan amount"`
	MaxExtensionTimes         int             `json:"maxExtensionTimes" description:"Max extension times"`
	TotalExtensionDaysOption  string          `json:"totalExtensionDaysOption" description:"Total extension days option"`
	MaxLoanPeriods            int             `json:"maxLoanPeriods" description:"Max loan periods"`
	MinLoanPeriods            int             `json:"minLoanPeriods" description:"Min loan periods"`
	RevolvingFlag             string          `json:"revolvingFlag" description:"Revolving flag"`
	InstallmentMethod         string          `json:"installmentMethod" description:"Installment method"`
	InstallmentRoundDirection string          `json:"installmentRoundDirection" description:"Installment round direction"`
	InstallmentRoundDigits    int             `json:"installmentRoundDigits" description:"Installment round digits"`
	CompoundInterestMethod    string          `json:"compoundInterestMethod" description:"Compound interest method"`
	ImpairedDays              int             `json:"impairedDays" description:"Impaired days"`
	MaxContractAmount         decimal.Decimal `json:"maxContractAmount" description:"Max contract amount"`
	MinContractAmount         decimal.Decimal `json:"minContractAmount" description:"Max contract amount"`
	ReleaseQuotaMethod        string          `json:"releaseQuotaMethod" description:"Release quota method"`
	RepaymentOrderMethod      string          `json:"repaymentOrderMethod" description:"Repayment order method"`
	RepaymentOrderAmount      string          `json:"repaymentOrderAmount" description:"Repayment order amount"`
	ExpireDate                string          `json:"expireDate" description:"Expire date"`
	EffectiveDate             string          `json:"effectiveDate" description:"Effective date"`
	Status                    string          `json:"status" description:"Status"`
	ModifySeq                 string          `json:"modifySeq" description:"Modify seq"`
	InstallmentList           string          `json:"installmentList" description:"Installment list"`
	MaxContractCount          int             `json:"maxContractCount" description:"Max contract count"`
	MaxEarlyRepaymentTimes    int             `json:"maxEarlyRepaymentTimes" description:"Max early repayment times"`
	PrincipalRule             []PrincipalRule `json:"principalRule" description:"Principal rule"`
	QuotaControlOverdue       string          `json:"quotaControlOverdue" description:"0-Overdue does not affect the quota 1-Decrease in proportion to the number of overdue days 2-Frozen in full"`
	OverDueRule               []OverDueRule   `json:"overDueRule" description:"OverDue rule"`
}

func (*PD000097I) GetServiceKey() string {
	return "PD000097"
}
