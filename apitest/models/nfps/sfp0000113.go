package models

type FP000113I struct {
}

type FP000113O struct {
	PartList []PartList `json:"partList"`
}

func (*FP000113I) GetServiceKey() string {
	return "FP000113"
}
