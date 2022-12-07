/**
 * @Author: fangwen
 * @Description:
 * @File: sfp0000110
 * @Version: 1.0.0
 * @Date: 2022/3/28 11:08 AM
 */

package models

type FP000114I struct {
	PartClearingCode string   `json:"partClearingCode"`
	Currency         string   `json:"currency"`
	RequestStatus    []string `json:"requestStatus"`
	RequestStartDate string   `json:"requestStartDate"`
	RequestEndDate   string   `json:"requestEndDate"`
	PageNo           int      `json:"pageNo"`
	PageSize         int      `json:"pageSize"`
}

type FP000114O struct {
	TotalCount int             `json:"totalCount"`
	Records    []BalAdjustItem `json:"records"`
}

func (*FP000114I) GetServiceKey() string {
	return "FP000114"
}
