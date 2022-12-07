package models

//delete Agency Relationship for Indirect Participant
type FP900007I struct {
	DirectPartClearingCode   string `json:"directPartClearingCode"`
	IndirectPartClearingCode string `json:"indirectPartClearingCode"`
	Currency                 string `json:"currency"`
}
type FP900007O struct {
}

func (*FP900007I) GetServiceKey() string {
	return "FP900007"
}
