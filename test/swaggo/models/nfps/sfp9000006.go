package models

//Add Agency Relationship for Indirect Participant
type FP900006I struct {
	DirectPartClearingCode string `json:"directPartClearingCode"`
	//DirectPartName           string `json:"directPartName"`
	IndirectPartClearingCode string `json:"indirectPartClearingCode"`
	//IndirectPartName         string `json:"indirectPartName"`
	Currency      string `json:"currency"`
	EffectiveDate string `json:"effectiveDate"`
	Remark        string `json:"remark"`
}
type FP900006O struct {
}

func (*FP900006I) GetServiceKey() string {
	return "FP900006"
}
