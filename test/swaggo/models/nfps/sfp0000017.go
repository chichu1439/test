package models

//delete Agency Relationship for Indirect Participant

type FP000017I struct {
	DirectPartClearingCode   string `json:"directPartClearingCode" description:"Direct part clearing code" validate:"required"`
	IndirectPartClearingCode string `json:"indirectPartClearingCode" description:"Indirect part clearing code" validate:"required"`
	Currency                 string `json:"currency" description:"Currency" validate:"required"`
}

type FP000017O struct {
}

func (*FP000017I) GetServiceKey() string {
	return "FP000017"
}
