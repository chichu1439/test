//Version: v1.0.0.0
package models

type PR000009I struct {
	ScenarioId string           `json:"scenarioId" validate:"required" description:"Scenario Id"`
	Amount     float64          `json:"amount" validate:"required" description:"Amount"`
	Currency   string           `json:"currency" validate:"required" description:"Currency(THB)"`
	TrnDate    string           `json:"trnDate" validate:"required" description:"Trade date"`
	AParm      []PR000009IAParm `json:"AParm" description:"A parm" validate:"required"`
}

type PR000009IAParm struct {
	ParmCode  string `json:"parmCode" description:"Parameter Code"`
	ParmValue string `json:"parmValue" description:"Parameter Value"`
}

type PR000009IFormatAParm struct {
	ParmCode  string      `json:"parmCode" description:"Parameter Code"`
	ParmValue interface{} `json:"parmValue"`
}

type PR000009O struct {
	FeeTotal              float64            `json:"feeTotal" description:"Fee total"`
	DiscountTotal         float64            `json:"discountTotal" description:"Discount totall"`
	DiscountStg           string             `json:"discountStg" description:"Discount strategy"`
	DiscountSumType       string             `json:"discountSumType" description:"Discount sum type"`
	DiscountCalculateType string             `json:"discountCalculateType" description:"Discount calculate type"`
	AResult               []PR000009OAResult `json:"aResult" description:"A result"`
}

type PR000009OAResult struct {
	StrategyId         string  `json:"strategyId" description:"Strategy Id"`
	StrategyName       string  `json:"strategyName" description:"Strategy Name"`
	FeeInit            float64 `json:"feeInit" description:"Fee init"`
	Discount           float64 `json:"discount" description:"Discount amount"`
	FeeReal            float64 `json:"feeReal" description:"Fee real"`
	FeeMax             float64 `json:"feeMax" description:"Fee max"`
	FeeMin             float64 `json:"feeMin" description:"Fee min"`
	DiscountMax        float64 `json:"discountMax" description:"Fee discount max"`
	DiscountMin        float64 `json:"discountMin" description:"Fee discount mins"`
	FeeCalcAmt         float64 `json:"feeCalcAmt" description:"Fee calculate amount"`
	AmortizeFlag       string  `json:"amortizeFlag" description:"Amortize flag"`
	AmortizePeriodType string  `json:"amortizePeriodType" description:"Amortize period type"`
	AccountingCode     string  `json:"accountingCode" description:"Accounting code"`
}

func (*PR000009I) GetServiceKey() string {
	return "PR000009"
}
