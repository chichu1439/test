package models

type FP000025I struct {
	Currency    string `json:"currency"`
	ProductId   string `json:"productId"`
	ServiceId   string `json:"serviceId"`
	FeatureId   string `json:"featureId"`
	LimitStatus string `json:"limitStatus"`
}
type FP000025O struct {
	TotalCount int             `json:"totalCount"`
	Records    []FP000025IItem `json:"records"`
}

type FP000025IItem struct {
	ServiceId              string  `json:"serviceId"`
	FeatureId              string  `json:"featureId"`
	Currency               string  `json:"currency"`
	ProductId              string  `json:"productId"`
	LimitTarget            string  `json:"limitTarget"`
	LimitMethod            string  `json:"limitMethod"`
	LimitStrategyId        string  `json:"limitStrategyId"`
	UpperLimitControl      string  `json:"upperLimitControl"`
	UpperLimitValue        float64 `json:"upperLimitValue"`
	LowerLimitControl      string  `json:"lowerLimitControl"`
	LowerLimitValue        float64 `json:"lowerLimitValue"`
	EffectiveDate          string  `json:"effectiveDate"`
	ExpireDate             string  `json:"expireDate"`
	LimitStatus            string  `json:"limitStatus"`
	TransactionType        string  `json:"transactionType"`
	PaymentCategoryPurpose string  `json:"paymentCategoryPurpose"`
}

func (*FP000025I) GetServiceKey() string {
	return "FP000025"
}
