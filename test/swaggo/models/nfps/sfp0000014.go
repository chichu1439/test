package models

//verify random code

type FP000014I struct {
	RandomCode string `json:"randomCode" description:"Random code" validate:"required"`
}

type FP000014O struct {
}

func (*FP000014I) GetServiceKey() string {
	return "FP000014"
}
