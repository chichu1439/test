package models

type FP000024I struct {
	ServiceId         string  `json:"serviceId"`
	FeatureId         string  `json:"featureId"`
	Currency          string  `json:"currency"`
	ProductId         string  `json:"productId"`
	LimitTarget       string  `json:"limitTarget"`
	LimitMethod       string  `json:"limitMethod"`
	LimitStrategyId   string  `json:"limitStrategyId"`
	UpperLimitControl string  `json:"upperLimitControl"`
	UpperLimitValue   float64 `json:"upperLimitValue"`
	LowerLimitControl string  `json:"lowerLimitControl"`
	LowerLimitValue   float64 `json:"lowerLimitValue"`
	EffectiveDate     string  `json:"effectiveDate"`
	ExpireDate        string  `json:"expireDate"`
}
type FP000024O struct {
	ServiceId         string  `json:"serviceId"`
	FeatureId         string  `json:"featureId"`
	Currency          string  `json:"currency"`
	ProductId         string  `json:"productId"`
	LimitTarget       string  `json:"limitTarget"`
	LimitMethod       string  `json:"limitMethod"`
	LimitStrategyId   string  `json:"limitStrategyId"`
	UpperLimitControl string  `json:"upperLimitControl"`
	UpperLimitValue   float64 `json:"upperLimitValue"`
	LowerLimitControl string  `json:"lowerLimitControl"`
	LowerLimitValue   float64 `json:"lowerLimitValue"`
	EffectiveDate     string  `json:"effectiveDate"`
	ExpireDate        string  `json:"expireDate"`
	LimitStatus       string  `json:"limitStatus"`
}

func (*FP000024I) GetServiceKey() string {
	return "FP000024"
}
