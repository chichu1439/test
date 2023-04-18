/**
 * @Author: fangwen
 * @Description:
 * @File: sfp0000103
 * @Version: 1.0.0
 * @Date: 2022/3/28 11:08 AM
 */

package models

type FP000107I struct {
	Id               int64  `json:"id"`
	PartClearingCode string `json:"partClearingCode"`
	Currency         string `json:"currency"`
}

type FP000107O struct {
}

func (*FP000107I) GetServiceKey() string {
	return "FP000107"
}
