/**
 * @Author: fangwen
 * @Description:
 * @File: sfp0000099
 * @Version: 1.0.0
 * @Date: 2022/3/27 1:18 PM
 */

package models

import "github.com/go-playground/validator/v10"

type FP000099I struct {
	TestScenarioId int64 `json:"testScenarioId" validate:"required"`
	// TestApiId      int64 `json:"testApiId" validate:"required"`
	ScenarioStepId int64          `json:"scenarioStepId" validate:"required"`
	Type           string         `json:"type"`
	TcpStatus      string         `json:"TcpStatus"`
	Fields         []FieldUpdates `json:"fields"`
}

type FP000099O struct {
	HeaderFields map[string]interface{} `json:"headerFields"`
	RequestData  []byte                 `json:"requestData"`
	ApiInfo      RequestApiInfo         `json:"sedTemplate"`
	TraceId      string                 `json:"traceId"`
	Types        string                 `json:"type"`
	Simulation   bool                   `json:"simulation"`
}

func (*FP000099I) GetServiceKey() string {
	return "FP000099"
}

func (d *FP000099I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
