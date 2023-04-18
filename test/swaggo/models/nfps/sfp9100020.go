package models

type FP910020I struct {
	Mti      string `json:"mti"`
	FlowId   string `json:"flowId"`
	FlowText string `json:"flowText"`
	SndDate  string `json:"sndDate"`
	SndTime  string `json:"sndTime"`
}
type FP910020O struct {
	Status string
}

func (*FP910020I) GetServiceKey() string {
	return "FP910020"
}
