/**
 * @Author: fangwen
 * @Description:
 * @File: sfp0000103
 * @Version: 1.0.0
 * @Date: 2022/3/28 11:08 AM
 */

package models

type FP000109I struct {
	PartClearingCode string `json:"partClearingCode"`
	Currency         string `json:"currency"`
	Action           string `json:"action"`
	ActionComment    string `json:"actionComment"`
	RequestDate      string `json:"requestDate"`
	//AdjustDirection  string  `json:"adjustDirection"`
	//AdjustAmount     float64 `json:"adjustAmount"`
}

type FP000109O struct {
	Balance float64
}

func (*FP000109I) GetServiceKey() string {
	return "FP000109"
}
