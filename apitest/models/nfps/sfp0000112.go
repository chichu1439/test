package models

import "github.com/go-playground/validator/v10"

type FP000112I struct {
	ParamGroups []string `json:"paramGroup"  validate:"required,min=1"` // if not sed query all parameters
}

type FP000112O struct {
	EnumList []EnumList `json:"enumList"`
}

func (*FP000112I) GetServiceKey() string {
	return "FP000112"
}

func (d *FP000112I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
