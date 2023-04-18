package models

import "github.com/go-playground/validator/v10"

type FP000115I struct {
	Participant string `json:"participant" validate:"required"`
	PrivateKey  string `json:"privateKey"`
}

type FP000115O struct {
	Status string `json:"status"`
}

func (*FP000115I) GetServiceKey() string {
	return "FP000115"
}

func (d *FP000115I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
