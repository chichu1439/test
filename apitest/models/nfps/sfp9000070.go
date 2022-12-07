package models

type FP900070I struct {
	Id                  int64  `json:"id" validate:"required"`
	Step                int64  `json:"step"`
	Name                string `json:"name"`
	Result              bool   `json:"result"`
	Description         string `json:"description"`
	PreRequestScript    string `json:"preRequestScript"`
	PostResponseScript  string `json:"postResponseScript"`
	TestAssertionScript string `json:"testAssertionScript"`
}

type FP900070O struct {
	Status string `json:"status"`
}

func (*FP900070I) GetServiceKey() string {
	return "FP900070"
}
