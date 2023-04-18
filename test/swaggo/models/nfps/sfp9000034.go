package models

type FP900034I struct {
	ParticipantCode []string `json:"participantCode"`
}
type FP900034O struct {
	TotalCount      int                     `json:"totalCount"`
	ParticipantInfo map[string]FP900034Item `json:"participantInfo"`
}
type FP900034Item struct {
	PartName string `json:"partName"`
	PartType string `json:"partType"`
}

func (*FP900034I) GetServiceKey() string {
	return "FP900034"
}
