package models

//Settlement Account balance inquiry(redis)
type FP900016I struct {
	PartClearingCode string `json:"partClearingCode" validate:"required"`
	Currency         string `json:"currency" validate:"required"`
}
type FP900016O struct {
	Balance  float64 `json:"balance"`
	AvailBal float64 `json:"availBal"`
}

func (*FP900016I) GetServiceKey() string {
	return "FP900016"
}
