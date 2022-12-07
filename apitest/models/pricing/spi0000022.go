package models

type PI000022I struct {
	StrDate   string   `json:"strDate" description:"Str date"`
	EndDate   string   `json:"endDate" description:"End date"`
	ModifySeq []string `json:"modifySeq" description:"Modify seq"`
}

type PI000022O struct {
	ListInterest []PriceInterestStrategyHis `json:"listInterest" description:"List interest"`
}

type PriceInterestStrategyHis struct {
	InterestRate       float64 `json:"interestRate" description:"Interest rate"`
	StrategyId         string  `json:"strategyId" description:"Strategy id"`
	StrategyName       string  `json:"strategyName" description:"Strategy name"`
	SystemId           string  `json:"systemId" description:"System id"`
	Term               string  `json:"term" description:"Term"`
	Currency           string  `json:"currency" description:"Currency"`
	InterestId         string  `json:"interestId" description:"Interest id"`
	FloatType          string  `json:"floatType" description:"Float type"`
	FloatDirection     string  `json:"floatDirection" description:"Float direction"`
	MaxFloat           float64 `json:"maxFloat" description:"Max float"`
	MinFloat           float64 `json:"minFloat" description:"Min float"`
	MaxInterest        float64 `json:"maxInterest" description:"Max interest"`
	MinInterest        float64 `json:"minInterest" description:"Min interest"`
	CaculateType       string  `json:"caculateType" description:"Caculate type"`
	RateExecType       string  `json:"rateExecType" description:"Rate exec type"`
	FloatExecType      string  `json:"floatExecType" description:"Float exec type"`
	MonthDays          string  `json:"monthDays" description:"Month days"`
	YearDays           string  `json:"yearDays" description:"Year days"`
	DecimalFlag        string  `json:"decimalFlag" description:"Decimal flag"`
	AccrualPeriodType  string  `json:"accrualPeriodType" description:"Accrual period type"`
	AccrualPeriodValue int     `json:"accrualPeriodValue" description:"Accrual period value"`
	AccrualPeriodDay   int     `json:"accrualPeriodDay" description:"Accrual period day"`
	SettlePeriodType   string  `json:"settlePeriodType" description:"Settle period type"`
	SettlePeriodValue  int     `json:"settlePeriodValue" description:"Settle period value"`
	SettlePeriodDay    int     `json:"settlePeriodDay" description:"Settle period day"`
	MultiMatchType     string  `json:"multiMatchType" description:"Multi match type"`
	MultiFloatType     string  `json:"multiFloatType" description:"Multi float type"`
	EffectiveDate      string  `json:"effectiveDate" description:"Effective date"`
	ExpireDate         string  `json:"expireDate" description:"Expire date"`
	Status             string  `json:"status" description:"Status"`
	ModifySeq          string  `json:"modifySeq" description:"Modify seq"`
	LastModifyDate     string  `json:"lastModifyDate" description:"Last modify date"`
	LastModifyTime     string  `json:"lastModifyTime" description:"Last modify time"`
	LastModifyOrg      string  `json:"lastModifyOrg" description:"Last modify org"`
	LastModifyUser     string  `json:"lastModifyUser" description:"Last modify user"`
}

func (*PI000022I) GetServiceKey() string {
	return "PI000022"
}
