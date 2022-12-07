package models

//Inquire Participant Agency Relationship info
type FP900010I struct {
	DirectPartClearingCode   string `json:"directPartClearingCode" validate:"required"`
	IndirectPartClearingCode string `json:"indirectPartClearingCode" validate:"required"`
	Currency                 string `json:"currency" validate:"required"`
}
type FP900010O struct {
	DirectPartClearingCode   string `json:"directPartClearingCode"`
	DirectPartName           string `json:"directPartName"`
	IndirectPartClearingCode string `json:"indirectPartClearingCode"`
	IndirectPartName         string `json:"indirectPartName"`
	Currency                 string `json:"currency"`
	EffectiveDate            string `json:"effectiveDate"`
	Remark                   string `json:"remark"`
	Status                   string `json:"status"`
}

func (*FP900010I) GetServiceKey() string {
	return "FP900010"
}
