//Version: v1
package models

type FP900067Request struct {
	ID           int64              `json:"id" validate:"required"`
	Name         string             `json:"name"`
	Description  string             `json:"description"`
	ScenarioStep []ScenarioStepList `json:"scenarioStepList"`
	ClickStep    []int              `json:"clickStep"`
}

type FP900067Response struct {
	Status string `json:"status"`
}

func (*FP900067Request) GetServiceKey() string {
	return "FP900067"
}
