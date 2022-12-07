package models

import "github.com/go-playground/validator/v10"

type FP000047I struct {
	ParticipantClearingCode string `json:"participantClearingCode"  validate:"required"`
	ContactName             string `json:"contactName"  validate:"required"`
	ContactMethod           string `json:"contactMethod"  validate:"required"`
	ContactInformation      string `json:"contactInformation"  validate:"required"`
	Priority                string `json:"priority"`
	Necessary               string `json:"necessary"`
	Remark                  string `json:"remark"`
	Label                   string `json:"label"`
	Id                      int64  `json:"id"`
}

type FP000047O struct {
	AssociatedStates bool `json:"associatedStates"`
}

func (*FP000047I) GetServiceKey() string {
	return "FP000045"
}

func (d *FP000047I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
