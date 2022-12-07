/**
 * @Author: fangwen
 * @Description:
 * @File: sfp0000103
 * @Version: 1.0.0
 * @Date: 2022/3/28 11:08 AM
 */

package models

type FP000106I struct {
	Id               int64   `json:"id"`
	PartClearingCode string  `json:"partClearingCode"`
	Currency         string  `json:"currency"`
	AdjustDirection  string  `json:"adjustDirection"`
	AdjustAmount     float64 `json:"adjustAmount"`
	RequestDate      string  `json:"requestDate"`
}

type FP000106O struct {
}

func (*FP000106I) GetServiceKey() string {
	return "FP000106"
}
