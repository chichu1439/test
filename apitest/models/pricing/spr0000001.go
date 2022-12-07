//Version: v0.0.1
package models

type SPR0000001I struct {
	ItemName   string  `json:"itemName" validate:"required" description:"Item name"`
	ItemTerm   int     `json:"itemTerm" validate:"required" description:"Item term"`
	Currency   string  `json:"currency" validate:"required" description:"Currency(THB)"`
	Rate       float64 `json:"rate" validate:"required" description:"rate"`
	SourceType string  `json:"sourceType" validate:"required" description:"Source type"`
	EffectDate string  `json:"effectDate" description:"Effect date(yyyy-mm-dd)"`
}

type SPR0000001O struct {
	ItemId string `json:"itemId" description:"Item Id"`
}

func (*SPR0000001I) GetServiceKey() string {
	return "spr0000001"
}
