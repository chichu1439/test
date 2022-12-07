package models

type FP900111I struct{}

type FP900111O struct {
	PartList []PartList `json:"partList"`
}

type PartList struct {
	ParamName string `json:"paramName"`
	ParamCode string `json:"paramValue"`
}

func (*FP900111I) GetServiceKey() string {
	return "FP900111"
}
