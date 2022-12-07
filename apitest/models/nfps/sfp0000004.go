package models

//Suspend Participant

type FP000004I struct {
	PartClearingCode string `json:"partClearingCode" description:"Part clearing code" validate:"required"`
	PartType         string `json:"partType" description:"Part type"`
	Status           string `json:"status" description:"Status" validate:"required"` //P-待生效，N-正常，S-停用
}

type FP000004O struct {
}

func (*FP000004I) GetServiceKey() string {
	return "FP000004"
}
