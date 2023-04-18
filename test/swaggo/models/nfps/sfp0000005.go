package models

//UnSuspend Participant

type FP000005I struct {
	PartClearingCode string `json:"partClearingCode" description:"Part clearing code" validate:"required"`
	PartType         string `json:"partType" description:"Part type"`
	Status           string `json:"status" description:"Status" validate:"required"`
}

type FP000005O struct {
}

func (*FP000005I) GetServiceKey() string {
	return "FP000005"
}
