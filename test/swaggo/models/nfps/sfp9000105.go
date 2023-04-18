/**
 * @Author: fangwen
 * @Description:
 * @File: sfp0000103
 * @Version: 1.0.0
 * @Date: 2022/3/28 11:08 AM
 */

package models

type FP900105I struct {
	PartClearingCode string  `json:"partClearingCode"`
	Currency         string  `json:"currency"`
	AdjustDirection  string  `json:"adjustDirection"`
	AdjustAmount     float64 `json:"adjustAmount"`
	RequestDate      string  `json:"requestDate"`
}

type FP900105O struct {
}

func (*FP900105I) GetServiceKey() string {
	return "FP900105"
}
