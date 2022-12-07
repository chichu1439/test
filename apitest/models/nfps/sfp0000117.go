package models

import "github.com/go-playground/validator/v10"

type FP000117I struct {
	Participant string `json:"participant" validate:"required"`
}

type FP000117O struct {
	Status string `json:"status"`
}

func (*FP000117I) GetServiceKey() string {
	return "FP000117"
}

func (d *FP000117I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
