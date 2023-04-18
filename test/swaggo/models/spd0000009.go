package models

import "github.com/shopspring/decimal"

type SPD0000009I struct {
	BindingMedium       string               `json:"bindingMedium"`
	AllowOverdraft      string               `json:"allowOverdraft"`
	AllowCashTrans      string               `json:"allowCashTrans"`
	AllowFreezing       string               `json:"allowFreezing"`
	IsIssueRecoDoc      string               `json:"isIssueRecoDoc"`
	RecoDocMethod       string               `json:"recoDocMethod"`
	RetainedBalance     decimal.Decimal      `json:"retainedBalance"`
	IsControlMaxBalance string               `json:"isControlMaxBalance"`
	MaxBalance          decimal.Decimal      `json:"maxBalance"`
	TemplateType        string               `json:"templateType"`
	CustomerType        string               `json:"customerType"`
	SystemId            string               `json:"systemId" description:"System id:SV-Saving LN-Loan" validate:"required,max=2"` //SV-Saving LN-Loan
	ProductId           string               `json:"productId" description:"Product id" validate:"required"`
	ProductType         string               `json:"productType" description:"Product type" validate:"max=2"` //*Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan
	ProductName         string               `json:"productName" description:"Product name" `
	Currency            string               `json:"currency" description:"Currency" `
	Version             string               `json:"version" description:"Version" `
	Remark              string               `json:"remark" description:"Remark"`
	EffectiveDate       string               `json:"effectiveDate" description:"Effective date"`
	ExpireDate          string               `json:"expireDate" description:"Expire date"`
	Status              string               `json:"status" description:"Status"`
	CashCurrency        string               `json:"cashCurrency" description:"Cash currency:C-Cash T-Currency"` //C-Cash T-Currency
	MediumType          string               `json:"mediumType" description:"Medium type" `                      //0-debit card 1-credit card 2-passbook 3-checkbook
	ListInterest        []SPD0000009Interest `json:"listInterest"`
	ListCurrency        []string             `json:"listCurrency" description:"List Currency"`
}
type SPD0000009Interest struct {
	InterestType            string               `json:"interestType" validate:"required"` // 默认 0-normal
	InterestCalcMethod      string               `json:"interestCalcMethod"`               // interest calculation method
	TierType                string               `json:"tierType"`
	StartCalcInterestAmount decimal.Decimal      `json:"startCalcInterestAmount"`          // amount to start calculating interest
	TierAmountMethod        string               `json:"tierAmountMethod"`                 // 默认2
	SettlePeriodType        string               `json:"settlePeriodType"`                 // Interest Settlement Period
	YearDays                string               `json:"yearDays" description:"Year days"` // interest calculation basis
	SettlePeriodDay         int                  `json:"settlePeriodDay"`                  // Interest Settlement Date
	FloatSource             string               `json:"floatSource"`                      // 默认2
	StrategyId              string               `json:"strategyId" description:"Strategy id"`
	InterestUid             int                  `json:"interestUid" description:"interestUid"`
	StrategyUid             int                  `json:"strategyUid" description:"strategyUid"`
	RetainedBalance         decimal.Decimal      `json:"retainedBalance"`
	StrategyList            []SPD0000007Strategy `json:"strategyList" description:"List strategy" validate:"dive"`
	DefaultInterest         string               `json:"defaultInterest"`
	Currency                string               `json:"currency" description:"Currency"`
}
type SPD0000009O struct {
}

func (*SPD0000009I) GetServiceKey() string {
	return "PD000009"
}
