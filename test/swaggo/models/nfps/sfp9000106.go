/**
 * @Author: fangwen
 * @Description:
 * @File: sfp0000103
 * @Version: 1.0.0
 * @Date: 2022/3/28 11:08 AM
 */

package models

type FP900106I struct {
	Id               int64   `json:"id"`
	PartClearingCode string  `json:"partClearingCode"`
	Currency         string  `json:"currency"`
	AdjustDirection  string  `json:"adjustDirection"`
	AdjustAmount     float64 `json:"adjustAmount"`
	RequestDate      string  `json:"requestDate"`
	RequestStatus    string  `json:"requestStatus"`
	Action           string  `json:"action"`
	Comment          string  `json:"comment"`
	NewRequestStatus string  `json:"newRequestStatus"`
}

type FP900106O struct {
}

func (*FP900106I) GetServiceKey() string {
	return "FP900106"
}
