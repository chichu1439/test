//Version: v0.0.1
package models

import "github.com/shopspring/decimal"

//calculate_interest_by_product

type PI000010I struct {
	ListType    []ListType    `json:"listType" description:"List type"`
	ListBizParm []ListBizParm `json:"listBizParm" description:"List biz parm"`
}

type ListType struct {
	InterestType string           `json:"interestType" description:"Interest type(0-normal loan 1-overdue loan)"`
	ListStrategy []ListStrategyId `json:"listStrategy" description:"List strategy"`
}

type ListStrategyId struct {
	StrategyId string `json:"strategyId" description:"Strategy Id"`
}

type ListBizParm struct {
	ParmCode  string `json:"parmCode" description:"Parameter Code"`
	ParmValue string `json:"parmValue" description:"Parameter Value"`
}

type PI000010O struct {
	ListType []PI000010OListType `json:"listType" description:"List type"`
}

type PI000010OListType struct {
	InterestType string       `json:"interestType" description:"Interest type"`
	StrategyItem StrategyItem `json:"strategyItem" description:"Interest item"`
}

type StrategyItem struct {
	StrategyId         string          `json:"strategyId" description:"Strategy Id"`
	FloatType          string          `json:"floatType" description:"float Type：0-percent;1-BP;2-none"`
	FloatDirection     string          `json:"floatDirection" description:"floatDirection (0-up;1-down)"`
	MaxFloat           decimal.Decimal `json:"maxFloat" description:"Max float"`
	MinFloat           decimal.Decimal `json:"minFloat" description:"Min float"`
	MaxInterest        decimal.Decimal `json:"maxInterest" description:"Max interest"`
	MinInterest        decimal.Decimal `json:"minInterest" description:"Min interest"`
	InterestRate       decimal.Decimal `json:"interestRate" description:"Interest rate"`
	ExecInterestRate   decimal.Decimal `json:"execInterestRate" description:"Execute Interest rate"`
	FloatValue         decimal.Decimal `json:"floatValue" description:"Float value"`
	CaculateType       string          `json:"caculateType" description:"Caculate type(0-Actual days;-Year to month to day;2-Year to month)"`
	RateExecType       string          `json:"rateExecType" description:"Caculate type(0-By Value Date;1-By Calc Date)"`
	FloatExecType      string          `json:"floatExecType" description:"Float execute type(0-By Value Date;1-By Calc Date;2-Fixed float)"`
	MonthDays          string          `json:"monthDays" description:"Month days(0-30;1-Actual days)"`
	YearDays           string          `json:"yearDays" description:"Year days(0-360;1-365;2-Actual days)"`
	DecimalFlag        string          `json:"decimalFlag" description:"Decimal flag(0-no;1-yes)"`
	AccrualPeriodType  string          `json:"accrualPeriodType" validate:"required" description:"Accrual period type(01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly)"`
	AccrualPeriodValue int             `json:"accrualPeriodValue" validate:"required" description:"Accrual period value"`
	AccrualPeriodDay   int             `json:"accrualPeriodDay" description:"Accrual period day"`
	SettlePeriodType   string          `json:"settlePeriodType" validate:"required" description:"Settle period type(01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly)"`
	SettlePeriodValue  int             `json:"settlePeriodValue" validate:"required" description:"Settle period value"`
	SettlePeriodDay    int             `json:"settlePeriodDay" description:"Settle period day"`
}

func (*PI000010I) GetServiceKey() string {
	return "PI000010"
}
