package models

//gen random code

type FP000013I struct {
}

type FP000013O struct {
	RandomCode string `json:"randomCode" description:"Random code"`
}

func (*FP000013I) GetServiceKey() string {
	return "FP000013"
}
