package models

import "github.com/shopspring/decimal"

type PD000087I struct {
	StrDate         string `json:"strDate" description:"Start date"`
	EndDate         string `json:"endDate" description:"End date"`
	ProductId       string `json:"productId" description:"Product id" validate:"required"`
	PageNum         int    `json:"pageNum" description:"Page number"`
	PageRecordCount int    `json:"pageRecordCount" description:"Page record count"`
}

type PD000087O struct {
	HistoryRecords  PD000087OItem  `json:"historyRecords"`
	PlanningRecords PD000087OItem2 `json:"planningRecords"`
}
type PD000087OItem struct {
	PageNum         int                   `json:"pageNum" description:"Page number"`
	PageRecordCount int                   `json:"pageRecordCount" description:"Page record count"`
	TotalCount      int                   `json:"totalCount" description:"Total count"`
	Records         []ProdInterestTypeHis `json:"records" description:"Records"`
}
type PD000087OItem2 struct {
	TotalCount int                   `json:"totalCount" description:"Total count"`
	Records    []ProdInterestTypeHis `json:"records" description:"Records"`
}
type ProdInterestTypeHis struct {
	Uid                int             `json:"uid" description:"Uid"`
	InterestType       string          `json:"interestType" description:"interest Type:0-normal loan;1-overdue loan"`
	InterestId         string          `json:"interestId" description:"Interest id"`
	CustomerGrade      string          `json:"customerGrade" description:"0-All customers 1-Level 1 2-Level 2 3-Level 3"`
	MinInterestRate    decimal.Decimal `json:"minInterestRate" description:"Min interest rate"`
	MaxInterestRate    decimal.Decimal `json:"maxInterestRate" description:"Max interest rate"`
	MinFloatValue      decimal.Decimal `json:"minFloatValue" description:"Min float value"`
	MaxFloatValue      decimal.Decimal `json:"maxFloatValue" description:"Max float value"`
	EffectiveDate      string          `json:"effectiveDate" description:"Effective date"`
	ExpireDate         string          `json:"expireDate" description:"Expire date"`
	ModifySeq          string          `json:"modifySeq" description:"Modify seq"`
	Status             string          `json:"status" description:"status:0-inactive;1-active;2-expired;3-deleted;4-updated"`
	BaseInterestRate   decimal.Decimal `json:"baseInterestRate" description:"Base interest rate"`
	StrategyId         string          `json:"strategyId" description:"Strategy id"`
	FloatDirection     string          `json:"floatDirection" description:"float Direction:0-None;1-Up;2-Down"`
	FloatMethod        string          `json:"floatMethod" description:"float Method:0-By Percent;1-BP Value"`
	CalculateType      string          `json:"calculateType" description:"caculate Type:0-Actual days;1-Year to month to day;2-Year to month"`
	RateExecType       string          `json:"rateExecType" description:"rate Exec Type:0-By Value Date;1-By Calc Date"`
	FloatExecType      string          `json:"floatExecType" description:"float ExecType:0-By Value Date;1-By Calc Date;2-Fixed float"`
	MonthDays          string          `json:"monthDays" description:"month Days:0-30;1-Actual days"`
	YearDays           string          `json:"yearDays" description:"year Days:0-360;1-365;2-Actual days"`
	DecimalFlag        string          `json:"decimalFlag" description:"decimal Flag:0-No;1-Yes"`
	AccrualPeriodType  string          `json:"accrualPeriodType" description:"Accrual period type"`
	AccrualPeriodValue int             `json:"accrualPeriodValue" description:"Accrual period value"`
	AccrualPeriodDay   int             `json:"accrualPeriodDay" description:"Accrual period day"`
	SettlePeriodType   string          `json:"settlePeriodType" description:"settle Period Type:01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly"`
	SettlePeriodValue  int             `json:"settlePeriodValue" description:"Settle period value"`
	SettlePeriodDay    int             `json:"settlePeriodDay" description:"Settle period day"`
	RecordTime         string          `json:"recordTime"`
	RecordDate         string          `json:"recordDate"`
}

func (*PD000087I) GetServiceKey() string {
	return "PD000087"
}
