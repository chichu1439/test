package models

type FP000040I struct {
	FeeId              string  `json:"feeId" validate:"required"`
	ProductId          string  `json:"productId" validate:"required"`
	ServiceId          string  `json:"serviceId" validate:"required"`
	FeatureId          string  `json:"featureId" validate:"required"`
	MaxDiscount        float64 `json:"maxDiscount"`
	MinDiscount        float64 `json:"minDiscount"`
	EffectiveDate      string  `json:"effectiveDate" validate:"required"`
	ExpireDate         string  `json:"expireDate"`
	FeeName            string  `json:"feeName"`
	SystemId           string  `json:"systemId"`
	Currency           string  `json:"currency"`
	IsAmortization     string  `json:"isAmortization"`
	AmortizationPeriod string  `json:"amortizationPeriod"`
	AmortizationTimes  int     `json:"amortizationTimes"`
	ChargeType         string  `json:"chargeType"`
	ChargePeriod       string  `json:"chargePeriod"`
	ChargePeriodValue  int     `json:"chargePeriodValue"`
	ChargePeriodDay    int     `json:"chargePeriodDay"`
	FeeCalcMethod      string  `json:"feeCalcMethod"`
	CalcValue          float64 `json:"calcValue"`
	MaxFeeAmount       float64 `json:"maxFeeAmount"`
	MinFeeAmount       float64 `json:"minFeeAmount"`
	AccountingCode     string  `json:"accountingCode"`
	TriggerInfo        string  `json:"triggerInfo"`
	PriorityFlag       string  `json:"priorityFlag"`
	PayerType          string  `json:"payerType"`
}
type FP000040O struct {
	Uid                int     `json:"uid"`
	FeeId              string  `json:"feeId"`
	ProductId          string  `json:"productId"`
	ServiceId          string  `json:"serviceId"`
	FeatureId          string  `json:"featureId"`
	Status             string  `json:"status"`
	MaxDiscount        float64 `json:"maxDiscount"`
	MinDiscount        float64 `json:"minDiscount"`
	EffectiveDate      string  `json:"effectiveDate"`
	ExpireDate         string  `json:"expireDate"`
	ModifySeq          string  `json:"modifySeq"`
	FinalModifyDate    string  `json:"finalModifyDate"`
	FinalModifyTime    string  `json:"finalModifyTime"`
	FinalModifyOrgz    string  `json:"finalModifyOrgz"`
	FinalModifyUser    string  `json:"finalModifyUser"`
	FeeName            string  `json:"feeName"`
	SystemId           string  `json:"systemId"`
	Currency           string  `json:"currency"`
	IsAmortization     string  `json:"isAmortization"`
	AmortizationPeriod string  `json:"amortizationPeriod"`
	AmortizationTimes  int     `json:"amortizationTimes"`
	ChargeType         string  `json:"chargeType"`
	ChargePeriod       string  `json:"chargePeriod"`
	ChargePeriodValue  int     `json:"chargePeriodValue"`
	ChargePeriodDay    int     `json:"chargePeriodDay"`
	FeeCalcMethod      string  `json:"feeCalcMethod"`
	CalcValue          float64 `json:"calcValue"`
	MaxFeeAmount       float64 `json:"maxFeeAmount"`
	MinFeeAmount       float64 `json:"minFeeAmount"`
	AccountingCode     string  `json:"accountingCode"`
	TriggerInfo        string  `json:"triggerInfo"`
	PriorityFlag       string  `json:"priorityFlag"`
	PayerType          string  `json:"payerType"`
}

func (*FP000040I) GetServiceKey() string {
	return "FP000040"
}
