//Version: v0.0.1
package models

type SPR0000002I struct {
	StrategyID         string           `json:"strategyId" description:"Strategy ID"`
	StrategyName       string           `json:"strategyName" validate:"required" description:"Strategy name"`
	ProductId          string           `json:"productId" validate:"required" description:"Product ID"`
	ItemTerm           int              `json:"itemTerm" validate:"required" description:"Item term"`
	Currency           string           `json:"currency" validate:"required" description:"Currency(THB)"`
	InterestType       string           `json:"interestType" validate:"required" description:"Interest type(0-normal loan 1-overdue loan)"`
	ItemId             string           `json:"itemId" validate:"required" description:"Item Id"`
	FloatType          string           `json:"floatType" validate:"required" description:"Float type(0-percent;1-BP)"`
	FloatMax           float64          `json:"floatMax" validate:"required" description:"Max float"`
	FloatMin           float64          `json:"floatMin" validate:"required" description:"Min float"`
	InterestMax        float64          `json:"interestMax" description:"Max interest"`
	InterestMin        float64          `json:"interestMin" description:"Min interest"`
	CalculateType      string           `json:"calculateType" description:"Calculate type(0-Actual days;1-Year to month to day;2-Year to month)"`
	ExecuteType        string           `json:"executeType" description:"Execute type(0-By Value Date;1-By Calc Date)"`
	MonthDay           string           `json:"monthDay" validate:"required" description:"Month days"`
	YearDay            string           `json:"yearDay" validate:"required" description:"Year days"`
	DecimalFlag        string           `json:"decimalFlag" validate:"required" description:"Decimal flag(0-no;1-yes)"`
	AccrualPeriodType  string           `json:"accrualPeriodType" validate:"required" description:"Accrual period type(01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly)"`
	AccrualPeriodValue int              `json:"accrualPeriodValue" validate:"required" description:"Accrual period value"`
	SettlePeriodType   string           `json:"settlePeriodType" validate:"required" description:"Settle period type(01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly)"`
	SettlePeriodValue  int              `json:"settlePeriodValue" validate:"required" description:"Settle period value"`
	EffectDate         string           `json:"effectDate" description:"Effect date(yyyy-mm-dd)"`
	MultiMatchType     string           `json:"multiMatchType" validate:"required" description:"Multi match type(0-any;1-all)"`
	MultiFloatType     string           `json:"multiFloatType" validate:"required" description:"Multi float type(0-min float;1-max float;2-sum)"`
	AParmMatch         []AParmMatchList `json:"aParmMatch" description:"A parm match"`
	AParmFloat         []AParmFloatList `json:"aParmFloat" description:"A parm float"`
}

type AParmMatchList struct {
	ParmCode           string       `json:"parmCode" description:"Parameter code"`
	ParmValue          string       `json:"parmValue" description:"Parameter value"`
	ParmMatchType      string       `json:"parmMatchType" description:"Parameter type"`
	RangeCalculateType string       `json:"rangeCalculateType" description:"Range calculate type(0-exclude end;1-exclude start;2-include both)"`
	ARange             []ARangeList `json:"aRange" description:"A range"`
}

type AParmFloatList struct {
	ParmCode            string       `json:"parmCode" description:"Parameter code"`
	ParmType            string       `json:"parmType" description:"Parameter type(0-single;1-group)"`
	ParmValue           string       `json:"parmValue" description:"Parameter value"`
	GroupMultiMatchType string       `json:"groupMultiMatchType" description:"Group multi match type(0-Greater than;1-Greater than or equal to;2-Less than;3-Less than or equal to;4-equal;5-not equal;6-range)"`
	ParmMatchType       string       `json:"parmMatchType" description:"Parameter match type(0-Greater than;1-Greater than or equal to;2-Less than;3-Less than or equal to;4-equal;5-not equal;6-range)"`
	RangeCalculateType  string       `json:"rangeCalculateType" description:"Range calculate type(0-exclude end;1-exclude start;2-include both)"`
	FloatValue          float64      `json:"floatValue" description:"Float value"`
	ARange              []ARangeList `json:"aRange" description:"A range"`
	AGroup              []AGroupList `json:"aGroup" description:"A group"`
}

type AGroupList struct {
	GroupId            string       `json:"groupId" description:"Group id"`
	GroupName          string       `json:"groupName" description:"Group name"`
	ParmCode           string       `json:"parmCode" description:"Parm code"`
	ParmValue          string       `json:"parmValue" description:"Parm value"`
	ParmMatchType      string       `json:"parmMatchType" description:"ParmMatch type(0-Greater than;1-Greater than or equal to;2-Less than;3-Less than or equal to;4-equal;5-not equal;6-range)"`
	RangeCalculateType string       `json:"rangeCalculateType" description:"Range calculate type(0-exclude end;1-exclude start;2-include both)"`
	ARange             []ARangeList `json:"aRange" description:"A range"`
}

type ARangeList struct {
	ParmCode   string `json:"parmCode" description:"Parm code"`
	RangeStart string `json:"rangeStart" description:"Range start"`
	RangeEnd   string `json:"rangeEnd" description:"Range end"`
	Result     string `json:"result" description:"Result"`
}

type SPR0000002O struct {
	StrategyId string `json:"strategyId" description:"Strategy id"`
}

func (*SPR0000002I) GetServiceKey() string {
	return "spr0000002"
}
