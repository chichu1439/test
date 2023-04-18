package models

import "github.com/go-playground/validator/v10"

type FP000048I struct {
	ParticipantClearingCode string `json:"participantClearingCode"`
	Id                      int64  `json:"id" validate:"required"`
}

type FP000048O struct {
	AssociatedStates bool `json:"associatedStates"`
}

func (*FP000048I) GetServiceKey() string {
	return "FP000048"
}

func (d *FP000048I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
