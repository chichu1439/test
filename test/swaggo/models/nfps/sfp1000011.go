/**
 * @Author: fangwen
 * @Description:
 * @File: sfp1000007
 * @Version: 1.0.0
 * @Date: 2022/3/18 1:45 AM
 */

package models

import "time"

type FP100011I struct {
	AlertID      string    `json:"alertID" validate:"required"`
	StrategyName string    `json:"strategyName"`
	AlertType    string    `json:"alertType" validate:"required"`
	AlertEntity  string    `json:"alertEntity" validate:"required"`
	AlertDetails string    `json:"alertDetails"`
	TimeStart    time.Time `json:"timeStart" validate:"required"`
	Actions      string    `json:"actions"`
	WorkSpace    string    `json:"workSpace" validate:"required"`
	Env          string    `json:"env" validate:"required"`
	Su           string    `json:"su"`
	AlertName    string    `json:"alertName" validate:"required"`
	AlertLevel   string    `json:"alertLevel" validate:"required"`
	AlertLabels  string    `json:"alertLabels"`
	Expression   string    `json:"expression"`
	Lable        string    `json:"lable"`
	Type         string    `json:"type"`
	TraceId      string    `json:"traceId"`
}

type FP100011O struct {
	AssociatedStates bool `json:"associatedStates"`
}

func (*FP100011I) GetServiceKey() string {
	return "FP100011"
}
