package models

type FP900089I struct {
	PartClearingCode string   `json:"partClearingCode" validation:"required"`
	Currency         string   `json:"currency" validation:"required"`
	AlertBalance     *float64 `json:"alertBalance"`
	AlertMinuteTime  *int     `json:"alertMinuteTime"`
	AlertHightWater  *float64 `json:"alertHightWater"`
	AlertLowWater    *float64 `json:"alertLowWater"`
	Status           string   `json:"status"`
}
type FP900089O struct {
}

func (*FP900089I) GetServiceKey() string {
	return "FP900089"
}
