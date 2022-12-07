package models

import (
	"github.com/go-playground/validator/v10"
)

type FP000049I struct {
	AlertID      string `json:"alertID"`
	AlertType    string `json:"alertType"`
	AlertEntity  string `json:"alertEntity"`
	StrategyName string `json:"strategyName"`
	AlertDetails string `json:"alertDetails" validate:"required"`
	TimeStart    string `json:"timeStart"`
	Actions      string `json:"actions"`
	WorkSpace    string `json:"workSpace"`
	Env          string `json:"env"`
	Su           string `json:"su"`
	AlertName    string `json:"alertName"`
	AlertLevel   string `json:"alertLevel"`
	AlertLabels  string `json:"alertLabels"`
	Expression   string `json:"expression"`
	Label        string `json:"label"`
	Type         string `json:"type"`
	TraceId      string `json:"traceId"`
}

type FP000049O struct {
	AssociatedStates bool `json:"associatedStates"`
}

func (*FP000049I) GetServiceKey() string {
	return "FP000045"
}

func (d *FP000049I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
