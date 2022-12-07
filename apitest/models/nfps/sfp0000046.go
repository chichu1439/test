package models

import "github.com/go-playground/validator/v10"

type FP000046I struct {
	ParticipantClearingCode string `json:"participantClearingCode"  validate:"required"`
	ContactName             string `json:"contactName"  validate:"required"`
	ContactMethod           string `json:"contactMethod"  validate:"required"`
	ContactInformation      string `json:"contactInformation"  validate:"required"`
	Priority                string `json:"priority"  validate:"required"`
	Necessary               string `json:"necessary"` //0/1
	Remark                  string `json:"remark"`
	Label                   string `json:"label"`
}

type FP000046O struct {
	AssociatedStates bool `json:"associatedStates"`
}

func (*FP000046I) GetServiceKey() string {
	return "FP000046"
}

func (d *FP000046I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
