//Version: v1
package models

import "github.com/shopspring/decimal"

type PR000004Request struct {
	DcnId                         string                      `json:"dcnId"`
	AccountNum                    string                      `json:"accountNum"`
	LastModifyOrg                 string                      `json:"lastModifyOrg"`
	LastModifyUser                string                      `json:"lastModifyUser"`
	AcctInterestStrategyInfo      AcctInterestStrategy        `json:"acctInterestStrategyList"`
	AcctInterestStrategyLevelList []AcctInterestStrategyLevel `json:"acctInterestStrategyLevelList"`
}
type AcctInterestStrategy struct {
	StrategyId              string          `json:"strategyId"`
	TransDate               string          `json:"transDate"`
	StrategyName            string          `json:"strategyName"`
	SystemId                string          `json:"systemId"`
	Term                    string          `json:"term"`
	Currency                string          `json:"currency"`
	InterestCalcFlag        string          `json:"interestCalcFlag"`
	LevelFlag               string          `json:"levelFlag"`
	StartCalcInterestAmount decimal.Decimal `json:"startCalcInterestAmount"`
	TierType                string          `json:"tierType"`
	TiterAmountMethod       string          `json:"titerAmountMethod"`
	InterestUseFlag         string          `json:"interestUseFlag"`
	FixedRate               decimal.Decimal `json:"fixedRate"`
	FloatSource             string          `json:"floatSource"`
	InterestId              string          `json:"interestId"`
	FloatType               string          `json:"floatType"`
	FloatDirection          string          `json:"floatDirection"`
	MaxFloat                decimal.Decimal `json:"maxFloat"`
	MinFloat                decimal.Decimal `json:"minFloat"`
	MaxInterest             decimal.Decimal `json:"maxInterest"`
	MinInterest             decimal.Decimal `json:"minInterest"`
	CaculateType            string          `json:"caculateType"`
	RateExecType            string          `json:"rateExecType"`
	FloatExecType           string          `json:"floatExecType"`
	MonthDays               string          `json:"monthDays"`
	YearDays                string          `json:"yearDays"`
	DecimalFlag             string          `json:"decimalFlag"`
	AccrualPeriodType       string          `json:"accrualPeriodType"`
	AccrualPeriodValue      int             `json:"accrualPeriodValue"`
	AccrualPeriodDay        int             `json:"accrualPeriodDay"`
	SettlePeriodType        string          `json:"settlePeriodType"`
	SettlePeriodValue       int             `json:"settlePeriodValue"`
	SettlePeriodDay         int             `json:"settlePeriodDay"`
	MultiMatchType          string          `json:"multiMatchType"`
	MultiFloatType          string          `json:"multiFloatType"`
}
type AcctInterestStrategyLevel struct {
	TransDate       string          `json:"transDate"`
	StrategyId      string          `json:"strategyId"`
	LevelNum        string          `json:"levelNum"`
	LevelAmount     decimal.Decimal `json:"levelAmount"`
	TierLimitAmount decimal.Decimal `json:"tierLimitAmount"`
	InterestUseFlag string          `json:"interestUseFlag"`
	FixedRate       decimal.Decimal `json:"fixedRate"`
	InterestId      string          `json:"interestId"`
	FloatType       string          `json:"floatType"`
	FloatDirection  string          `json:"floatDirection"`
	FloatValue      decimal.Decimal `json:"floatValue"`
	MaxFloat        decimal.Decimal `json:"maxFloat"`
	MinFloat        decimal.Decimal `json:"minFloat"`
	MaxInterest     decimal.Decimal `json:"maxInterest"`
	MinInterest     decimal.Decimal `json:"minInterest"`
}

type PR000004Response struct {
}

func (*PR000004Request) GetServiceKey() string {
	return "PR000004"
}
