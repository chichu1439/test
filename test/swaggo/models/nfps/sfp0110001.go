package models

import "github.com/go-playground/validator/v10"

type SF011001I struct {
	Fileaction     string `json:"fileaction"`
	FileSize       int64  `json:"fileSize"`
	FileName       string `json:"fileName" validate:"required"`
	FileDate       int64  `json:"fileDate"`
	SendBackNumber int    `json:"sendBackNumber"`
	InFileName     string `json:"inFileName"`
}

type SF011001O struct {
	Data string `json:"data"`
}

func (*SF011001I) GetServiceKey() string {
	return "FP011001"
}

func (d *SF011001I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
