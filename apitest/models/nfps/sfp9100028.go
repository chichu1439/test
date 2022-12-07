package models

//Settlement Account balance inquiry(redis)
type FP910028I struct {
	PartClearingCode string `json:"partClearingCode" validate:"required"`
	Currency         string `json:"currency" validate:"required"`
}
type FP910028O struct {
	Balance  float64 `json:"balance"`
	AvailBal float64 `json:"availBal"`
}

func (*FP910028I) GetServiceKey() string {
	return "FP910028"
}
