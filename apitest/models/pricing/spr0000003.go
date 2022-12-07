//Version: v0.0.1
package models

type SPR0000003I struct {
	ProductId string             `json:"productId" validate:"required" description:"Product Id"`
	ItemTerm  int                `json:"itemTerm" validate:"required" description:"Item term"`
	Currency  string             `json:"currency" validate:"required" description:"Currency(THB)"`
	AParm     []SPR0000003IAParm `json:"AParm" description:"A parm" validate:"required"`
}

type SPR0000003IAParm struct {
	ParmCode  string `json:"parmCode" validate:"required" description:"Parameter Code"`
	ParmValue string `json:"parmValue" validate:"required" description:"Parameter Value"`
}

type SPR0000003O struct {
	AResult []SPR0000003OAResult `json:"aResult" description:"A result"`
}

type SPR0000003OAResult struct {
	InterestType       string  `json:"interestType" description:"Interest type(0-normal loan 1-overdue loan)"`
	StrategyId         string  `json:"strategyId" description:"Strategy ID"`
	ItemId             string  `json:"itemId" description:"Item Id"`
	BaseRate           float64 `json:"baseRate" description:"Base interest rate"`
	Rate               float64 `json:"rate" description:"Interest rate ID"`
	FloatValue         float64 `json:"floatValue" description:"Float value"`
	FloatType          string  `json:"floatType" description:"Float type(0-percent;1-BP)"`
	InterestMax        float64 `json:"interestMax" description:"Max interest"`
	InterestMin        float64 `json:"interestMin" description:"Min interest"`
	CalculateType      string  `json:"calculateType" description:"Calculate type(0-Actual days;1-Year to month to day;2-Year to month)"`
	ExecuteType        string  `json:"executeType" description:"Execute type(0-By Value Date;1-By Calc Date)"`
	MonthDay           string  `json:"monthDay" validate:"required" description:"Month days"`
	YearDay            string  `json:"yearDay" validate:"required" description:"Year days"`
	DecimalFlag        string  `json:"decimalFlag" validate:"required" description:"Decimal flag(0-no;1-yes)"`
	AccrualPeriodType  string  `json:"accrualPeriodType" validate:"required" description:"Accrual period type(01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly)"`
	AccrualPeriodValue int     `json:"accrualPeriodValue" validate:"required" description:"Accrual period value"`
	SettlePeriodType   string  `json:"settlePeriodType" validate:"required" description:"Settle period type(01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly)"`
	SettlePeriodValue  int     `json:"settlePeriodValue" validate:"required" description:"Settle period value"`
}

func (*SPR0000003I) GetServiceKey() string {
	return "spr0000003"
}
