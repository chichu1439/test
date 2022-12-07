//Version: v0.0.1
package models

type SPR0000004I struct {
	Account string               `json:"account" validate:"required,max=32" description:"Account number"`
	DcnId   string               `json:"dcnId" description:"Su ID"`
	AResult []SPR0000004IAResult `json:"aResult" description:"A result"`
}

type SPR0000004IAResult struct {
	InterestType       string  `json:"interestType" description:"Interest type(0-normal loan 1-overdue loan)"`
	StrategyId         string  `json:"strategyId" validate:"required" description:"Strategy ID"`
	ItemId             string  `json:"itemId" validate:"required" description:"Item Id"`
	Rate               float64 `json:"rate" description:"Interest rate ID"`
	FloatValue         float64 `json:"floatValue" description:"Float value"`
	InterestMax        float64 `json:"interestMax" validate:"required" description:"Max interest"`
	InterestMin        float64 `json:"interestMin" validate:"required" description:"Min interest"`
	CalculateType      string  `json:"calculateType" validate:"required" description:"Calculate type(0-Actual days;1-Year to month to day;2-Year to month)"`
	ExecuteType        string  `json:"executeType" validate:"required" description:"Execute type(0-By Value Date;1-By Calc Date)"`
	MonthDay           string  `json:"monthDay" validate:"required" description:"Month days"`
	YearDay            string  `json:"yearDay" validate:"required" description:"Year days"`
	DecimalFlag        string  `json:"decimalFlag" validate:"required" description:"Decimal flag(0-no;1-yes)"`
	AccrualPeriodType  string  `json:"accrualPeriodType" validate:"required" description:"Accrual period type(01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly)"`
	AccrualPeriodValue int     `json:"accrualPeriodValue" validate:"required" description:"Accrual period value"`
	SettlePeriodType   string  `json:"settlePeriodType" validate:"required" description:"Settle period type(01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly)"`
	SettlePeriodValue  int     `json:"settlePeriodValue" validate:"required" description:"Settle period value"`
}

type SPR0000004O struct {
}

func (*SPR0000004I) GetServiceKey() string {
	return "spr0000004"
}
