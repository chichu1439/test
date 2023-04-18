package models

type FP900031I struct {
	FeeId                  string  `json:"feeId"`
	ProductId              string  `json:"productId"`
	ServiceId              string  `json:"serviceId"`
	FeatureId              string  `json:"featureId"`
	MaxDiscount            float64 `json:"maxDiscount"`
	MinDiscount            float64 `json:"minDiscount"`
	EffectiveDate          string  `json:"effectiveDate"`
	ExpireDate             string  `json:"expireDate"`
	ModifySeq              string  `json:"modifySeq"`
	FinalModifyOrgz        string  `json:"finalModifyOrgz"`
	FinalModifyUser        string  `json:"finalModifyUser"`
	FeeName                string  `json:"feeName"`
	SystemId               string  `json:"systemId"`
	Currency               string  `json:"currency"`
	IsAmortization         string  `json:"isAmortization"`
	AmortizationPeriod     string  `json:"amortizationPeriod"`
	AmortizationTimes      int     `json:"amortizationTimes"`
	ChargeType             string  `json:"chargeType"`
	ChargePeriod           string  `json:"chargePeriod"`
	ChargePeriodValue      int     `json:"chargePeriodValue"`
	ChargePeriodDay        int     `json:"chargePeriodDay"`
	FeeCalcMethod          string  `json:"feeCalcMethod"`
	CalcValue              float64 `json:"calcValue"`
	MaxFeeAmount           float64 `json:"maxFeeAmount"`
	MinFeeAmount           float64 `json:"minFeeAmount"`
	AccountingCode         string  `json:"accountingCode"`
	TriggerInfo            string  `json:"triggerInfo"`
	PriorityFlag           string  `json:"priorityFlag"`
	PayerType              string  `json:"payerType"`
	TransactionType        string  `json:"transactionType"`
	PaymentCategoryPurpose string  `json:"paymentCategoryPurpose"`
}
type FP900031O struct {
	Status string `json:"status"`
}

func (*FP900031I) GetServiceKey() string {
	return "FP900031"
}
