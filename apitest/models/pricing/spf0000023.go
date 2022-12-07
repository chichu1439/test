package models

import (
	"github.com/shopspring/decimal"
)

type PF000023I struct {
	StrDate   string   `json:"strDate" description:"Start date"`
	EndDate   string   `json:"endDate" description:"End date"`
	FeeIdList []string `json:"feeIdList" description:"Fee id list"`
}

type PF000023O struct {
	ListProdFeeItemHis []ProdFeeItemHis `json:"prodFeeItemHis" description:"List prod fee item his"`
}

type ProdFeeItemHis struct {
	SystemId           string          `json:"systemId" description:"System Id(LN-icredit;SV-isaving)"`
	FeeId              string          `json:"feeId" description:"Fee id"`
	FeeName            string          `json:"feeName" description:"Fee name"`
	TriggerInfo        string          `json:"triggerInfo" description:"Trigger information"`
	AccountingCode     string          `json:"accountingCode" description:"Accounting code"`
	AmortizationPeriod string          `json:"amortizationPeriod" description:"Amortization period"`
	AmortizationTimes  int             `json:"amortizationTimes" description:"Amortization times"`
	CalcValue          decimal.Decimal `json:"calcValue" description:"Calc value"`
	ChargePeriod       string          `json:"chargePeriod" description:"Charge period"`
	ChargePeriodDay    int             `json:"chargePeriodDay" description:"Charge period day"`
	ChargePeriodValue  int             `json:"chargePeriodValue" description:"Charge period value"`
	ChargeType         string          `json:"chargeType" description:"Charge type"`
	Currency           string          `json:"currency" description:"Currency"`
	FeeCalcMethod      string          `json:"feeCalcMethod" description:"Fee calc method"`
	IsAmortization     string          `json:"isAmortization" description:"Is amortization"`
	MaxFeeAmount       decimal.Decimal `json:"maxFeeAmount" description:"Max fee amount"`
	MinFeeAmount       decimal.Decimal `json:"minFeeAmount" description:"Min fee amount"`
	PayerType          string          `json:"payerType" description:"Payer type"`
	PriorityFlag       string          `json:"priorityFlag" description:"Priority flag"`
	EffectiveDate      string          `json:"effectiveDate" description:"Effective date"`
	ExpireDate         string          `json:"expireDate" description:"Expire date"`
	Status             string          `json:"status" description:"Status"`
	ModifySeq          string          `json:"modifySeq" description:"Modify seq"`
	FinalModifyDate    string          `json:"finalModifyDate" description:"Final modify date"`
	FinalModifyOrgz    string          `json:"finalModifyOrgz" description:"Final modify organization"`
	FinalModifyTime    string          `json:"finalModifyTime" description:"Final modify time"`
	FinalModifyUser    string          `json:"finalModifyUser" description:"Final modify user"`
}

func (*PF000023I) GetServiceKey() string {
	return "PF000023"
}
