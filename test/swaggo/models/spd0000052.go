package models

import "github.com/shopspring/decimal"

type SPD0000052I struct {
	ProductId string `json:"productId" description:"Product id" validate:"required"`
}

type SPD0000052O struct {
	ProductId           string                `json:"productId" description:"Product id"`
	SystemId            string                `json:"systemId" description:"system Id:SV-Saving LN-Loan"`                                                 //SV-Saving LN-Loan
	ProductType         string                `json:"productType" description:"product Type:Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan"` //*Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan
	ProductName         string                `json:"productName" description:"Product name"`
	EffectiveDate       string                `json:"effectiveDate" description:"Effective date"`
	ExpireDate          string                `json:"expireDate" description:"Expire date"`
	Currency            string                `json:"currency" description:"Currency"`
	Version             string                `json:"version" description:"Version"`
	Remark              string                `json:"remark" description:"Remark"`
	Status              string                `json:"status"`
	CustomerType        string                `json:"customerType"`
	CashCurrency        string                `json:"cashCurrency" description:"cashCurrency flag:C-Cash T-Currency"`                         //C-Cash T-Currency
	BindingMedium       string                `json:"bindingMedium" description:"0-no 1-yes"`                                                 //0-no 1-yes
	MediumType          string                `json:"mediumType" description:"medium Type:0-debit card 1-credit card 2-passbook 3-checkbook"` //0-debit card 1-passbook 2-checkbook
	AllowOverdraft      string                `json:"allowOverdraft"`
	AllowCashTrans      string                `json:"allowCashTrans"`
	AllowFreezing       string                `json:"allowFreezing"`
	IsIssueRecoDoc      string                `json:"isIssueRecoDoc"`
	RecoDocMethod       string                `json:"recoDocMethod"`
	IsControlMaxBalance string                `json:"isControlMaxBalance"`
	MaxBalance          decimal.Decimal       `json:"maxBalance"`
	TemplateType        string                `json:"templateType"`
	RetainedBalance     decimal.Decimal       `json:"retainedBalance"`
	ListCategory        []ProdCategoryNameBo  `json:"listCategory"`
	ListInterest        []SPD0000052OInterest `json:"listInterest" description:"List interest"`
	ListCurrency        []string              `json:"listCurrency" description:"List Currency"`
}

func (*SPD0000052I) GetServiceKey() string {
	return "PD000052"
}

type SPD0000052OInterest struct {
	RetainedBalance         decimal.Decimal               `json:"retainedBalance"`
	InterestType            string                        `json:"interestType" description:"interestType:0-normal loan 1-overdue loan"` //0-normal loan 1-overdue loan
	InterestCalcMethod      string                        `json:"interestCalcMethod"`                                                   // interest calculation method
	TierType                string                        `json:"tierType"`
	StartCalcInterestAmount decimal.Decimal               `json:"startCalcInterestAmount"` // amount to start calculating interest
	TierAmountMethod        string                        `json:"tierAmountMethod"`        // 默认2
	DefaultInterest         string                        `json:"defaultInterest"`
	TierDays                string                        `json:"tierDays"`
	SettlePeriodType        string                        `json:"settlePeriodType"`                 // Interest Settlement Period
	YearDays                string                        `json:"yearDays" description:"Year days"` // interest calculation basis
	SettlePeriodDay         int                           `json:"settlePeriodDay"`                  // Interest Settlement Date
	FloatSource             string                        `json:"floatSource"`                      // 默认2
	StrategyId              string                        `json:"strategyId" description:"Strategy id"`
	LevelFlag               string                        `json:"levelFlag"`
	InterestUid             int                           `json:"interestUid" description:"interestUid"`
	StrategyUid             int                           `json:"strategyUid" description:"strategyUid"`
	StrategyList            []SPD0000052OInterestStrategy `json:"strategyList" description:"List interest strategy"`
	Currency                string                        `json:"currency"`
}
type SPD0000052OInterestStrategy struct {
	InterestId       string          `json:"interestId" description:"Interest id"`
	LevelNum         string          `json:"levelNum" description:"0-All customers 1-Level 1 2-Level 2 3-Level 3"`
	BaseInterestRate decimal.Decimal `json:"baseInterestRate" description:"Base interest rate"`
	MinInterestRate  decimal.Decimal `json:"minInterestRate" description:"Min interest rate"`
	MaxInterestRate  decimal.Decimal `json:"maxInterestRate" description:"Max interest rate"`
	FloatType        string          `json:"floatType" description:"float Type:0-By Percent 1-BP Value"`             //0-By Percent 1-BP Value
	FloatDirection   string          `json:"floatDirection" description:"float Direction:0-None 1-Up 2-Down 3-Both"` //0-None 1-Up 2-Down 3-Both
	MinFloatValue    decimal.Decimal `json:"minFloatValue" description:"Min float value"`
	MaxFloatValue    decimal.Decimal `json:"maxFloatValue" description:"Max float value"`
	FloatValue       decimal.Decimal `json:"floatValue"`
	FixInterestRate  decimal.Decimal `json:"fixInterestRate"` // interest Type是fix interest 时有值
	// 前端是：tiered interest type,枚举是0-default interest,1-base interest 2-fixed interest
	// 前端是：interest Type：枚举是1-base interest 2-fixed interest
	TieredInterestType string          `json:"tieredInterestType"`
	TierLimitAmount    decimal.Decimal `json:"tierLimitAmount"` // tiered limit amount, default interest 为空
	TierDays           int             `json:"tierDays"`
	LevelAmount        decimal.Decimal `json:"levelAmount"`
}
