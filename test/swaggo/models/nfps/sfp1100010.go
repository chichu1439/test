package models

//Settlement Account balance inquiry(redis)
type FP110010I struct {
	PartClearingCode string `json:"partClearingCode" validate:"required"`
	Currency         string `json:"currency" validate:"required"`
}
type FP110010O struct {
	AvailBal float64 `json:"availBal"`
}

func (*FP110010I) GetServiceKey() string {
	return "FP110010"
}
