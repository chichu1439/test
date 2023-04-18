/**
 * @Author: fangwen
 * @Description:
 * @File: sfp0000103
 * @Version: 1.0.0
 * @Date: 2022/3/28 11:08 AM
 */

package models

type FP900107I struct {
	Id               int64  `json:"id"`
	PartClearingCode string `json:"partClearingCode"`
	Currency         string `json:"currency"`
	RequestDate      string `json:"requestDate"`
	RequestStatus    string `json:"requestStatus"`
}

type FP900107O struct {
}

func (*FP900107I) GetServiceKey() string {
	return "FP900107"
}
