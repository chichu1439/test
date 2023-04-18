//Version: v0.0.1
package models

import "github.com/shopspring/decimal"

type PI000008I struct {
	UpdateStrategyName  bool            `json:"updateStrategyName"`
	DelCurrencyFlag     bool            `json:"delCurrencyFlag"`
	CurrentCurrencyList []string        `json:"currentCurrencyList"`
	ProductId           string          `json:"productId"`
	NewProductName      string          `json:"newProductName"`
	OldStatus           string          `json:"oldStatus"`
	RecordList          []PI000008IItem `json:"recordList"`
}
type PI000008IItem struct {
	ModifySeq               string           `json:"modifySeq" description:"Modify sequnce number"`
	StrategyId              string           `json:"strategyId" validate:"required" description:"Strategy Id"`
	StrategyName            string           `json:"strategyName" description:"Strategy Name"`
	SystemId                string           `json:"systemId" description:"Strategy Name(SV-Saving LN-Loan)"` //SV-Saving LN-Loan
	Term                    string           `json:"term" description:"Term"`
	Currency                string           `json:"currency" description:"Currency(THB)"`
	InterestId              string           `json:"interestId" description:"Interest Id"`
	FloatType               string           `json:"floatType" description:"Float type(0-percent 1-BP)"`        //0-percent 1-BP
	FloatDirection          string           `json:"floatDirection" description:"Float direction(0-up 1-down)"` //0-up 1-down
	MaxFloat                decimal.Decimal  `json:"maxFloat" description:"Max float"`
	MinFloat                decimal.Decimal  `json:"minFloat" description:"Min float"`
	MaxInterest             decimal.Decimal  `json:"maxInterest" description:"Max interest"`
	MinInterest             decimal.Decimal  `json:"minInterest" description:"Min interest"`
	CalculateType           string           `json:"calculateType" description:"Calculate type(0-Actual days 1-Year to month to day 2-Year to month)"` //0-Actual days 1-Year to month to day 2-Year to month
	RateExecType            string           `json:"rateExecType" description:"Rate execute type(0-By Value Date 1-By Calc Date)"`                     //0-By Value Date 1-By Calc Date
	FloatExecType           string           `json:"floatExecType" description:"Float execute type(0-By Value Date 1-By Calc Date 2-Fixed float)"`     //0-By Value Date 1-By Calc Date 2-Fixed float
	MonthDays               string           `json:"monthDays" description:"Month days(0-30;1-Actual days)"`
	YearDays                string           `json:"yearDays" description:"Year days(0-360;1-365;2-Actual days)"`
	DecimalFlag             string           `json:"decimalFlag" description:"Decimal flag(0-no;1-yes)"`
	AccrualPeriodType       string           `json:"accrualPeriodType" validate:"required" description:"Accrual period type(01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly)"`
	AccrualPeriodValue      int              `json:"accrualPeriodValue" validate:"required" description:"Accrual period value"`
	AccrualPeriodDay        int              `json:"accrualPeriodDay" description:"Accrual period day"`
	SettlePeriodType        string           `json:"settlePeriodType" validate:"required" description:"Settle period type(01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly)"`
	SettlePeriodValue       int              `json:"settlePeriodValue" validate:"required" description:"Settle period value"`
	SettlePeriodDay         int              `json:"settlePeriodDay" description:"Settle period day"`
	MultiMatchType          string           `json:"multiMatchType" description:"Multi match type(0-min float;1-max float;2-sum)"`
	MultiFloatType          string           `json:"multiFloatType" description:"Multi float type(0-percent;1-BP;2-none")"`
	ListParmMatch           []ListParmMatch  `json:"listParmMatch" description:"List parm match"`
	ListParmFloat           []ListParmFloat  `json:"listParmFloat" description:"List parm float"`
	EffectiveDate           string           `json:"effectiveDate" description:"Effective date"`
	ExpireDate              string           `json:"expireDate" description:"Expire date"`
	RecordType              string           `json:"recordType"`
	ProductId               string           `json:"productId"`
	InterestType            string           `json:"interestType"`
	InterestCalcFlag        string           `json:"interestCalcFlag"`
	LevelFlag               string           `json:"levelFlag"`
	StartCalcInterestAmount decimal.Decimal  `json:"startCalcInterestAmount"`
	TierType                string           `json:"tierType"`
	TierAmountMethod        string           `json:"tierAmountMethod"`
	FloatSource             string           `json:"floatSource"`
	InterestUseFlag         string           `json:"interestUseFlag"`
	FixedRate               decimal.Decimal  `json:"fixedRate"`
	ListLevel               []PI000008ILevel `json:"listLevel"`
	IsDefaultRate           string           `json:"isDefaultRate"`
	RetainedBalance         decimal.Decimal  `json:"retainedBalance"`
}
type PI000008ILevel struct {
	StrategyId      string          `json:"strategyId" description:"Strategy id"`
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
	TierDays        int             `json:"tierDays"`
}
type PI000008O struct {
}

func (*PI000008I) GetServiceKey() string {
	return "PI000008"
}
