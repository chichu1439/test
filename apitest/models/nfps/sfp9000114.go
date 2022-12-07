package models

type FP900114I struct {
	Participant string `json:"participant"`
}

type FP900114O struct {
	Status string `json:"status"`
}

func (*FP900114I) GetServiceKey() string {
	return "FP900114"
}
