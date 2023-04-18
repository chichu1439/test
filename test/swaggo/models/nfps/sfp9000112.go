package models

type FP900112I struct {
	Participant string `json:"participant"`
	PrivateKey  string `json:"privateKey"`
}

type FP900112O struct {
	Status string `json:"status"`
}

func (*FP900112I) GetServiceKey() string {
	return "FP900112"
}
