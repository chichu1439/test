package models

type FP000027I struct {
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
	TransactionType    string `json:"transactionType"`
	PaymentCategoryPurpose string `json:"paymentCategoryPurpose"`
}

type FP000027O struct {
	FeeId     string `json:"feeId"`
	ModifySeq string `json:"modifySeq"`
}

func (*FP000027I) GetServiceKey() string {
	return "FP000027"
}
