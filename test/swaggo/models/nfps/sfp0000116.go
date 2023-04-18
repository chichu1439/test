package models

import "github.com/go-playground/validator/v10"

type FP000116I struct {
	Participant string `json:"participant" validate:"required"`
}

type FP000116O struct {
	Participant string `json:"participant"`
	PrivateKey  string `json:"privateKey"`
}

func (*FP000116I) GetServiceKey() string {
	return "FP000116"
}

func (d *FP000116I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
