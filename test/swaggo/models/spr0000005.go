//Version: v0.0.1
package models

type SPR0000005I struct {
	ProductId string `json:"productId" validate:"required" description:"Product ID"`
	ItemTerm  int    `json:"itemTerm" description:"Interest rate schedule"`
	Currency  string `json:"currency" description:"Currency"`
}

type SPR0000005O struct {
	AStrategy []SPR0000005OAStrategy `json:"aStrategy" description:"Interest rate pricing strategy array"`
}

type SPR0000005OAStrategy struct {
	StrategyID         string                  `json:"strategyId" description:"PI000006OStrategy ID"`
	StrategyName       string                  `json:"strategyName" description:"PI000006OStrategy name"`
	ProductId          string                  `json:"productId" description:"Product ID"`
	ItemTerm           int                     `json:"itemTerm" description:"Interest rate schedule"`
	Currency           string                  `json:"currency" description:"Currency"`
	InterestType       string                  `json:"interestType" description:"Interest rate type"`
	ItemId             string                  `json:"itemId" description:"Interest rate ID"`
	FloatType          string                  `json:"floatType" description:"Floating method"`
	FloatMax           float64                 `json:"floatMax" description:"Floating ceiling"`
	FloatMin           float64                 `json:"floatMin" description:"Floating lower limit"`
	InterestMax        float64                 `json:"interestMax" description:"Maximum interest"`
	InterestMin        float64                 `json:"interestMin" description:"Minimum interest"`
	BaseRate           float64                 `json:"baseRate" description:"Benchmark interest rate"`
	CalculateType      string                  `json:"calculateType" description:"Interest calculation algorithm"`
	ExecuteType        string                  `json:"executeType" description:"Interest rate execution method"`
	MonthDay           string                  `json:"monthDay" description:"Monthly interest accrual days"`
	YearDay            string                  `json:"yearDay" description:"Annual interest accrual days"`
	DecimalFlag        string                  `json:"decimalFlag" description:"Whether the decimal is interest-bearing"`
	AccrualPeriodType  string                  `json:"accrualPeriodType" description:"Type of accrual period"`
	AccrualPeriodValue int                     `json:"accrualPeriodValue" description:"Period value"`
	SettlePeriodType   string                  `json:"settlePeriodType" description:"Interest calculation period type"`
	SettlePeriodValue  int                     `json:"settlePeriodValue" description:"Interest accrual period value"`
	MultiMatchType     string                  `json:"multiMatchType" description:"Matching combination"`
	MultiFloatType     string                  `json:"multiFloatType" description:"Floating combination"`
	AParmMatch         []SPR0000005OAParmMatch `json:"aParmMatch" description:"Match parameter array"`
	AParmFloat         []SPR0000005OAParmFloat `json:"aParmFloat" description:"Floating calculation array"`
}

type SPR0000005OAParmMatch struct {
	StrategyID         string                  `json:"strategyId" description:"PI000006OStrategy ID"`
	ParmCode           string                  `json:"parmCode" description:"Parameter code"`
	ParmValue          string                  `json:"parmValue" description:"Parameter value"`
	ParmMatchType      string                  `json:"parmMatchType" description:"Parameter matching method"`
	RangeCalculateType string                  `json:"rangeCalculateType" description:"Interval calculation method"`
	AListRange         []SPR0000005OAListRange `json:"aListRange" description:"Compare interval arrays"`
}

type SPR0000005OAParmFloat struct {
	StrategyID          string                  `json:"strategyId" description:"PI000006OStrategy ID"`
	ParmType            string                  `json:"parmType" description:"Parameter Type"`
	ParmCode            string                  `json:"parmCode" description:"Parm code"`
	ParmValue           string                  `json:"parmValue" description:"Parameter value"`
	GroupMultiMatchType string                  `json:"groupMultiMatchType" description:"Combination parameter matching method"`
	ParmMatchType       string                  `json:"parmMatchType" description:"Single parameter matching method"`
	RangeCalculateType  string                  `json:"rangeCalculateType" description:"Interval calculation method"`
	FloatValue          float64                 `json:"floatValue" description:"Floating value"`
	AListRange          []SPR0000005OAListRange `json:"aListRange" description:"Compare interval arrays"`
	AListGroup          []SPR0000005OAListGroup `json:"aListGroup" description:"Combined parameter array"`
}

type SPR0000005OAListGroup struct {
	StrategyID         string                  `json:"strategyId" description:"PI000006OStrategy ID"`
	GroupId            string                  `json:"groupId" description:"Combination parameter ID"`
	ParmCode           string                  `json:"parmCode" description:"Parameter code"`
	ParmValue          string                  `json:"parmValue" description:"Parameter value"`
	ParmMatchType      string                  `json:"parmMatchType" description:"Parameter matching method"`
	RangeCalculateType string                  `json:"rangeCalculateType" description:"Interval calculation method"`
	AListRange         []SPR0000005OAListRange `json:"aListRange" description:"Compare interval arrays"`
}

type SPR0000005OAListRange struct {
	StrategyID string `json:"strategyId" description:"PI000006OStrategy ID"`
	GroupId    string `json:"groupId" description:"Combination parameter ID"`
	ParmCode   string `json:"parmCode" description:"Parameter code"`
	RangeStart string `json:"rangeStart" description:"Interval start value"`
	RangeEnd   string `json:"rangeEnd" description:"Interval end value"`
	Result     string `json:"result" description:"Hit value"`
}

func (*SPR0000005I) GetServiceKey() string {
	return "PR000005"
}
