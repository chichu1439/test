/**
 * @Author: fangwen
 * @Description:
 * @File: sfp0000103
 * @Version: 1.0.0
 * @Date: 2022/3/28 11:08 AM
 */

package models

type FP900108I struct {
	PartClearingCode string   `json:"partClearingCode"`
	Currency         string   `json:"currency"`
	RequestStatus    []string `json:"requestStatus"`
	RequestStartDate string   `json:"requestStartDate"`
	RequestEndDate   string   `json:"requestEndDate"`
	PageNo           int      `json:"pageNo"`
	PageSize         int      `json:"pageSize"`
}

type FP900108O struct {
	TotalCount int             `json:"totalCount"`
	Records    []BalAdjustItem `json:"records"`
}

func (*FP900108I) GetServiceKey() string {
	return "FP900108"
}
