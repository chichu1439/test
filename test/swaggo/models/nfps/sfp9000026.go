// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP900026I struct {
	ServiceId   string `json:"serviceId"`
	Currency    string `json:"currency"`
	FeatureId   string `json:"featureId"`
	ProductId   string `json:"productId"`
	LimitTarget string `json:"limitTarget"`
}

type FP900026O struct {
	Records []FP900026ResponseRecord `json:"records"`
}

type FP900026ResponseRecord struct {
	LimitId           int     `json:"limitId"`
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
	CreateDateTime    string  `json:"createDateTime"`
}

func (*FP900026I) GetServiceKey() string {
	return "FP900026"
}
