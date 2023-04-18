package models

//Suspend Participant
type FP900005I struct {
	PartClearingCode string `json:"partClearingCode"`
	PartType         string `json:"partType"` //no need
	Status           string `json:"status"`
}
type FP900005O struct {
}

func (*FP900005I) GetServiceKey() string {
	return "FP900005"
}
