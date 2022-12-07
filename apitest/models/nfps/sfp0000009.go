package models

//Add Agency Relationship for Indirect Participant

type FP000009I struct {
	DirectPartClearingCode   string `json:"directPartClearingCode" description:"Direct part clearing code" validate:"required"`
	IndirectPartClearingCode string `json:"indirectPartClearingCode" description:"Indirect part clearing code" validate:"required"`
	Currency                 string `json:"currency" description:"Currency" validate:"required"`
	Remark                   string `json:"remark" description:"Remark"`
}

type FP000009O struct {
}

func (*FP000009I) GetServiceKey() string {
	return "FP000009"
}
