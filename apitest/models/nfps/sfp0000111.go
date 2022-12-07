/**
 * @Author: fangwen
 * @Description:
 * @File: sfp0000110
 * @Version: 1.0.0
 * @Date: 2022/3/28 11:08 AM
 */

package models

type FP000111I struct {
	PartClearingCode string `json:"partClearingCode"`
	Currency         string `json:"currency"`
	NettingDate      string `json:"nettingDate"`
	NettingId        string `json:"nettingId"`
	PageNo           int    `json:"pageNo"`
	PageSize         int    `json:"pageSize"`
}

type FP000111O struct {
	TotalCount int                 `json:"totalCount"`
	Records    []NettingResultItem `json:"records"`
}

func (*FP000111I) GetServiceKey() string {
	return "FP000111"
}
