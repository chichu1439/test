package models

type FP910019I struct {
	PartCode string `json:"partCode"`
	Mti      string `json:"mti"`
	FlowId   string `json:"flowId"`
	FlowText []byte `json:"flowText"`
	FlowType string `json:"flowType"`
}
type FP910019O struct {
	Status string
}

func (*FP910019I) GetServiceKey() string {
	return "FP910019"
}
