//Version: v0.0.1
package models

type PI000006I struct {
	SystemId       string                `json:"systemId" description:"System Id"`
	Term           string                `json:"term" description:"Term"`
	Currency       string                `json:"currency" description:"Currency(THB)"`
	Status         string                `json:"status" description:"status(0-normal)"`
	ListStrategyId []PI000006IStrategyId `json:"listStrategyId" description:"List strategy id"`
}

type PI000006IStrategyId struct {
	StrategyId string `json:"strategyId" description:"Strategy id"`
}

type PI000006O struct {
	ListStrategy []PI000006OStrategy `json:"listStrategy" description:"List strategy"`
}

type PI000006OStrategy struct {
	StrategyId         string  `json:"strategyId" description:"Strategy id"`
	StrategyName       string  `json:"strategyName" description:"strategy name"`
	SystemId           string  `json:"systemId" description:"System Id"`
	Term               string  `json:"term" description:"Term"`
	Currency           string  `json:"currency" description:"Currency(THB)"`
	InterestId         string  `json:"interestId" description:"Interest Id"`
	InterestRate       float64 `json:"interestRate" description:"Interest rate"`
	FloatType          string  `json:"floatType" description:"Float type(0-percent;1-BP)"`
	FloatDirection     string  `json:"floatDirection" description:"float Direction：0-none;1-up;2-down"`
	MaxFloat           float64 `json:"maxFloat" description:"Max float"`
	MinFloat           float64 `json:"minFloat" description:"Min float"`
	MaxInterest        float64 `json:"maxInterest" description:"Max interest"`
	MinInterest        float64 `json:"minInterest" description:"Min interest"`
	CaculateType       string  `json:"caculateType" description:"caculate Type：0-Actual days;-Year to month to day;2-Year to month"`
	RateExecType       string  `json:"rateExecType" description:"rateExec Type：0-By Value Date;1-By Calc Date"`
	FloatExecType      string  `json:"floatExecType" description:"float Exec Type：0-By Value Date;1-By Calc Date;2-Fixed float"`
	MonthDays          string  `json:"monthDays" description:"Month days(0-30;1-Actual days)"`
	YearDays           string  `json:"yearDays" description:"Year days(0-360;1-365;2-Actual days)"`
	DecimalFlag        string  `json:"decimalFlag" description:"Decimal flag(0-no;1-yes)"`
	AccrualPeriodType  string  `json:"accrualPeriodType" validate:"required" description:"Accrual period type(01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly)"`
	AccrualPeriodValue int     `json:"accrualPeriodValue" validate:"required" description:"Accrual period value"`
	AccrualPeriodDay   int     `json:"accrualPeriodDay" description:"Accrual period day"`
	SettlePeriodType   string  `json:"settlePeriodType" validate:"required" description:"Settle period type(01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly)"`
	SettlePeriodValue  int     `json:"settlePeriodValue" validate:"required" description:"Settle period value"`
	SettlePeriodDay    int     `json:"settlePeriodDay" description:"Settle period day"`
	MultiMatchType     string  `json:"multiMatchType" description:"Multi match type(0-min float;1-max float;2-sum)"`
	MultiFloatType     string  `json:"multiFloatType" description:"Multi float type(0-percent;1-BP;2-none")"`
	EffectiveDate      string  `json:"effectiveDate" description:"Effective date"`
	ExpireDate         string  `json:"expireDate" description:"Expire date"`
	Status             string  `json:"status" description:"0-normal"`
	LastModifyOrg      string  `json:"lastModifyOrg" description:"Last modify org"`
	LastModifyUser     string  `json:"lastModifyUser" description:"Last modify user"`
	LastModifyDate     string  `json:"lastModifyDate" description:"Last modify date"`
	LastModifyTime     string  `json:"lastModifyTime" description:"Last modify time"`
}

func (*PI000006I) GetServiceKey() string {
	return "PI000006"
}
