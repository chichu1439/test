package models

type FP900073I struct {
	Fields []Fields
}

type FP900073O struct {
	Status string `json:"status"`
}

func (*FP900073I) GetServiceKey() string {
	return "FP900073"
}
