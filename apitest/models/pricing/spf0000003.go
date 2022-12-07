package models

import "github.com/shopspring/decimal"

type PF000003I struct {
	ModifySeq          string          `json:"modifySeq" description:"Modify seq"`
	SystemId           string          `json:"systemId" validate:"required" description:"System Id(LN-icredit;SV-isaving)"`
	FeeId              string          `json:"feeId" description:"Fee Id"`
	FeeName            string          `json:"feeName" description:"Fee Name"`
	TriggerInfo        string          `json:"triggerInfo" description:"Trigger information"`
	AccountingCode     string          `json:"accountingCode" description:"Accounting code"`
	AmortizationPeriod string          `json:"amortizationPeriod" validate:"required,max=2" description:"Amortization period(01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly)"` //0-No 1-Yes
	AmortizationTimes  int             `json:"amortizationTimes" description:"Amortization times"`                                                                                                                                                     //0-Day 1-Week 2-Month 3-Quarter 4-Half Year 5-Year
	CalcValue          decimal.Decimal `json:"calcValue" description:"Calculate value"`
	ChargePeriod       string          `json:"chargePeriod" validate:"required" description:"Charge period(01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly)" ` //0-Instant 1-By Period
	ChargePeriodDay    int             `json:"chargePeriodDay" description:"Charge period day"`
	ChargePeriodValue  int             `json:"chargePeriodValue" description:"Charge period value"`
	ChargeType         string          `json:"chargeType" description:"Charge type(0-Instant;1-By Period)"`
	Currency           string          `json:"currency" validate:"required,max=3" description:"Currency(THB)"`
	FeeCalcMethod      string          `json:"feeCalcMethod" description:"Fee calculate method(0-Percent 1-Fixed Amount 2-By Strategy)"` //0-Percent 1-Fixed Amount 2-By Strategy
	IsAmortization     string          `json:"isAmortization" validate:"required,max=2" description:"Is amortization(0-No;1-Yes)"`
	MaxFeeAmount       decimal.Decimal `json:"maxFeeAmount" description:"Max fee amount"`
	MinFeeAmount       decimal.Decimal `json:"minFeeAmount" description:"Min fee amount"`
	PayerType          string          `json:"payerType" description:"Payer type(01-Payee Participant;02-Payer Participant)"`
	PriorityFlag       string          `json:"priorityFlag" description:"Priority flag(0-priority 1-normal)"` //0-priority 1-normal
	EffectiveDate      string          `json:"effectiveDate" description:"Effective date"`
	ExpireDate         string          `json:"expireDate" description:"Expire date"`
	Status             string          `json:"status" description:"0-Normal"`
	FinalModifyOrgz    string          `json:"finalModifyOrgz" description:"Final modify organization"`
	FinalModifyUser    string          `json:"finalModifyUser" description:"Final modify user"`
	ListStrategy       []ListStrategy  `json:"listStrategy" description:"List strategy"`
}

type PF000003O struct {
}

func (*PF000003I) GetServiceKey() string {
	return "PF000003"
}
