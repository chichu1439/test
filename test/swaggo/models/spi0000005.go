//Version: v0.0.1
package models

import "github.com/shopspring/decimal"

type PI000005I struct {
	RecordList map[string][]PI000005IItem2 `json:"recordList"`
}

type PI000005IItem2 struct {
	Index                   int             `json:"index"`
	ModifySeq               string          `json:"modifySeq" description:"Modify seq"`
	StrategyName            string          `json:"strategyName" description:"Strategy name"`
	SystemId                string          `json:"systemId" description:"System id"`
	Term                    string          `json:"term" description:"Term"`
	Currency                string          `json:"currency" description:"Currency"`
	InterestId              string          `json:"interestId" description:"Interest id"`
	FloatType               string          `json:"floatType" description:"Float type(0-percent;1-BP)"`
	FloatDirection          string          `json:"floatDirection" description:"Float Direction：0-none;1-up;2-down"`
	MaxFloat                decimal.Decimal `json:"maxFloat" description:"Max float"`
	MinFloat                decimal.Decimal `json:"minFloat" description:"Min float"`
	MaxInterest             decimal.Decimal `json:"maxInterest" description:"Max interest"`
	MinInterest             decimal.Decimal `json:"minInterest" description:"Min interest"`
	CalculateType           string          `json:"calculateType" description:"calculate Type：0-Actual days;-Year to month to day;2-Year to month"`
	RateExecType            string          `json:"rateExecType" description:"rate Exec Type：0-By Value Date;1-By Calc Date"`
	FloatExecType           string          `json:"floatExecType" description:"float Exec Type：0-By Value Date;1-By Calc Date;2-Fixed float"`
	MonthDays               string          `json:"monthDays" description:"Month days(0-30;1-Actual days)"`
	YearDays                string          `json:"yearDays" description:"Year days(0-360;1-365;2-Actual days)"`
	DecimalFlag             string          `json:"decimalFlag" description:"Decimal flag(0-no;1-yes)"`
	AccrualPeriodType       string          `json:"accrualPeriodType" validate:"required" description:"Accrual period type(01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly)"`
	AccrualPeriodValue      int             `json:"accrualPeriodValue" validate:"required" description:"Accrual period value"`
	AccrualPeriodDay        int             `json:"accrualPeriodDay" description:"Accrual period day"`
	SettlePeriodType        string          `json:"settlePeriodType" validate:"required" description:"Settle period type(01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly)"`
	SettlePeriodValue       int             `json:"settlePeriodValue" validate:"required" description:"Settle period value"`
	SettlePeriodDay         int             `json:"settlePeriodDay" description:"Settle period day"`
	MultiMatchType          string          `json:"multiMatchType" description:"Multi match type(0-min float;1-max float;2-sum)"`
	MultiFloatType          string          `json:"multiFloatType" description:"Multi float type(0-percent;1-BP;2-none"`
	ListParmMatch           []ListParmMatch `json:"listParmMatch" description:"List parm match"`
	ListParmFloat           []ListParmFloat `json:"listParmFloat" description:"List parm float"`
	EffectiveDate           string          `json:"effectiveDate" description:"Effective date"`
	ExpireDate              string          `json:"expireDate" description:"Expire date"`
	Status                  string          `json:"status"`
	ProductId               string          `json:"productId"`
	InterestType            string          `json:"interestType"`
	InterestCalcFlag        string          `json:"interestCalcFlag"`
	LevelFlag               string          `json:"levelFlag"`
	StartCalcInterestAmount decimal.Decimal `json:"startCalcInterestAmount"`
	TierType                string          `json:"tierType"`
	IsDefaultRate           string          `json:"isDefaultRate"`
	TierAmountMethod        string          `json:"tierAmountMethod"`
	FloatSource             string          `json:"floatSource"`
	InterestUseFlag         string          `json:"interestUseFlag"`
	FixedRate               decimal.Decimal `json:"fixedRate"`
	ListLevel               []ListLevel     `json:"listLevel"`
	RetainedAmount          decimal.Decimal `json:"retainedAmount"`
}
type ListLevel struct {
	LevelNum        string          `json:"levelNum"`
	LevelAmount     decimal.Decimal `json:"levelAmount"`
	TierLimitAmount decimal.Decimal `json:"tierLimitAmount"`
	InterestUseFlag string          `json:"interestUseFlag"`
	FixedRate       decimal.Decimal `json:"fixedRate"`
	TierDays        int             `json:"tierDays"`
	InterestId      string          `json:"interestId"`
	FloatType       string          `json:"floatType"`
	FloatDirection  string          `json:"floatDirection"`
	FloatValue      decimal.Decimal `json:"floatValue"`
	MaxFloat        decimal.Decimal `json:"maxFloat"`
	MinFloat        decimal.Decimal `json:"minFloat"`
	MaxInterest     decimal.Decimal `json:"maxInterest"`
	MinInterest     decimal.Decimal `json:"minInterest"`
}
type ListParmMatch struct {
	ParmCode           string      `json:"parmCode" description:"Parm code"`
	ParmValue          string      `json:"parmValue" description:"Parm value"`
	ParmMatchType      string      `json:"parmMatchType" description:"Parm match type"`
	RangeCalculateType string      `json:"rangeCalculateType" description:"Range calculate type"`
	ListRange          []ListRange `json:"listRange" description:"List range"`
}

type ListParmFloat struct {
	ParmCode           string          `json:"parmCode" description:"Parm code"`
	ParmType           string          `json:"parmType" description:"Parm type"`
	ParmValue          string          `json:"parmValue" description:"Parm value"`
	GroupMatchType     string          `json:"groupMatchType" description:"Group match type"`
	ParmMatchType      string          `json:"parmMatchType" description:"Parm match type"`
	RangeCalculateType string          `json:"rangeCalculateType" description:"Range calculate type"`
	FloatValue         decimal.Decimal `json:"floatValue" description:"Float value"`
	ListRange          []ListRange     `json:"listRange" description:"List range"`
	ListGroup          []ListGroup     `json:"listGroup" description:"List group"`
}

type ListGroup struct {
	GroupId            string      `json:"groupId" description:"Group id"`
	GroupName          string      `json:"groupName" description:"Group name"`
	ParmCode           string      `json:"parmCode" description:"Parm code"`
	ParmValue          string      `json:"parmValue" description:"Parm value"`
	ParmMatchType      string      `json:"parmMatchType" description:"Parm match type"`
	RangeCalculateType string      `json:"rangeCalculateType" description:"Range calculate type"`
	ListRange          []ListRange `json:"listRange" description:"List range"`
}

type ListRange struct {
	ParmCode   string `json:"parmCode" description:"Parm code"`
	RangeStart string `json:"rangeStart" description:"Range start"`
	RangeEnd   string `json:"rangeEnd" description:"Range end"`
	Result     string `json:"result" description:"Result"`
	ResultType string `json:"resultType" description:"Result type"`
}

type PI000005O struct {
	StrategyIdList map[string]map[int]string `json:"strategyIdList"`
}

func (*PI000005I) GetServiceKey() string {
	return "PI000005"
}
