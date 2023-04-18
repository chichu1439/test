package models

import "github.com/shopspring/decimal"

type SPD0000007I struct {
	ProductName         string               `json:"productName" description:"Product name" validate:"required"`
	SystemId            string               `json:"systemId" description:"System id" validate:"required,max=2"`       // SV-Saving LN-Loan
	ProductType         string               `json:"productType" description:"Product type" validate:"required,max=2"` // 04-个人活期储蓄账户(1类)  05-个人活期储蓄账户(2类) 06-个人活期储蓄账户(3类)
	CustomerType        string               `json:"customerType"`                                                     // 0-Personal 1-Corporate
	Currency            string               `json:"currency" description:"Currency" validate:"required"`
	Remark              string               `json:"remark" description:"Remark"`
	EffectiveDate       string               `json:"effectiveDate" description:"Effective date" validate:"required"`
	ExpireDate          string               `json:"expireDate" description:"Expire date" validate:"required"`
	CashCurrency        string               `json:"cashCurrency" description:"Cash currency" validate:"required"`   //0-Cash 1-Currency
	BindingMedium       string               `json:"bindingMedium" description:"Binding Medium" validate:"required"` //0-no 1-yes
	MediumType          string               `json:"mediumType" description:"Medium type"`                           //0-debit card 1-passbook 2-checkbook
	AllowOverdraft      string               `json:"allowOverdraft"`
	AllowCashTrans      string               `json:"allowCashTrans"`
	AllowFreezing       string               `json:"allowFreezing"`
	IsIssueRecoDoc      string               `json:"isIssueRecoDoc"`
	RecoDocMethod       string               `json:"recoDocMethod"`
	RetainedBalance     decimal.Decimal      `json:"retainedBalance"`
	IsControlMaxBalance string               `json:"isControlMaxBalance"`
	MaxBalance          decimal.Decimal      `json:"maxBalance"`
	TemplateType        string               `json:"templateType"`
	ListInterest        []SPD0000007Interest `json:"listInterest"`
	ListCurrency        []string             `json:"listCurrency" description:"List Currency"`
}
type SPD0000007Interest struct {
	InterestType            string               `json:"interestType" validate:"required"` // 默认 0-normal
	InterestCalcMethod      string               `json:"interestCalcMethod"`               // interest calculation method
	TierType                string               `json:"tierType"`
	RetainedBalance         decimal.Decimal      `json:"retainedBalance"`
	DefaultInterest         string               `json:"defaultInterest"`
	StartCalcInterestAmount decimal.Decimal      `json:"startCalcInterestAmount"`          // amount to start calculating interest
	TierAmountMethod        string               `json:"tierAmountMethod"`                 // 默认2
	SettlePeriodType        string               `json:"settlePeriodType"`                 // Interest Settlement Period
	YearDays                string               `json:"yearDays" description:"Year days"` // interest calculation basis
	SettlePeriodDay         int                  `json:"settlePeriodDay"`                  // Interest Settlement Date
	FloatSource             string               `json:"floatSource"`                      // 默认2
	StrategyList            []SPD0000007Strategy `json:"strategyList" description:"List strategy" validate:"dive"`
	Currency                string               `json:"currency" description:"Currency"`
}
type SPD0000007Strategy struct {
	MinInterestRate    decimal.Decimal `json:"minInterestRate" description:"Min interest rate"`
	MaxInterestRate    decimal.Decimal `json:"maxInterestRate" description:"Max interest rate"`
	MinFloatValue      decimal.Decimal `json:"minFloatValue" description:"Min float value"`
	MaxFloatValue      decimal.Decimal `json:"maxFloatValue" description:"Max float value"`
	FloatDirection     string          `json:"floatDirection" description:"Float direction"`
	FloatType          string          `json:"floatType" description:"Float type"` // 前端是：float method
	FixInterestRate    decimal.Decimal `json:"fixInterestRate"`                    // interest Type是fix interest 时有值
	FloatValue         decimal.Decimal `json:"floatValue"`
	TieredInterestType string          `json:"tieredInterestType"`
	// 前端是：tiered interest type,枚举是0-default interest,1-base interest 2-fixed interest
	// 前端是：interest Type：枚举是1-base interest 2-fixed interest
	TierLimitAmount  decimal.Decimal `json:"tierLimitAmount"` // tiered limit amount, default interest 为空
	InterestId       string          `json:"interestId"`      // interest Type是base interest时有值
	LevelNum         string          `json:"levelNum"`        // default interest默认0-all ,tiered是1-level1 2-level2 ...
	BaseInterestRate decimal.Decimal `json:"baseInterestRate"`
	TierDays         int             `json:"tierDays"`
}

type SPD0000007O struct {
	ProductId string `json:"productId" description:"Product id"`
}

func (*SPD0000007I) GetServiceKey() string {
	return "PD000007"
}
