package models

type FP900113I struct {
	Participant string `json:"participant"`
}

type FP900113O struct {
	Participant string `json:"participant"`
	PrivateKey  string `json:"privateKey"` // base 64
}

func (*FP900113I) GetServiceKey() string {
	return "FP900113"
}
