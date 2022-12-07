/**
 * @Author: fangwen
 * @Description:
 * @File: sfp9000075
 * @Version: 1.0.0
 * @Date: 2022/3/31 11:31 AM
 */

package models

import "time"

type FP900075I struct {
	AlertID          string    `json:"alertId"`
	AlertType        string    `json:"alertType"`
	AlertEntity      string    `json:"alertEntity"`
	AlertDetails     string    `json:"alertDetails"`
	TimeStart        time.Time `json:"timeStart"`
	Actions          string    `json:"actions"`
	AlertName        string    `json:"alertName"`
	AlertLevel       string    `json:"alertLevel"`
	AlertLabels      string    `json:"alertLabels"`
	Expression       string    `json:"expression"`
	PartClearingCode string    `json:"partClearingCode"`
	ContactNum       int       `json:"contactNum"`
	Reason           string `json:"reason"`
	Lable            string `json:"lable"`
}

type FP900075O struct {
	AssociatedStates string `json:"associatedStates"`
}

func (*FP900075I) GetServiceKey() string {
	return "FP900075"
}
