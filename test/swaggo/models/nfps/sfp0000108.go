/**
 * @Author: fangwen
 * @Description:
 * @File: sfp0000103
 * @Version: 1.0.0
 * @Date: 2022/3/28 11:08 AM
 */

package models

type FP000108I struct {
	PartClearingCode string `json:"partClearingCode"`
	Currency         string `json:"currency"`
	PageNo           int    `json:"pageNo"`
	PageSize         int    `json:"pageSize"`
}

type FP000108O struct {
	TotalCount int             `json:"totalCount"`
	Records    []BalAdjustItem `json:"records"`
}

func (*FP000108I) GetServiceKey() string {
	return "FP000108"
}
