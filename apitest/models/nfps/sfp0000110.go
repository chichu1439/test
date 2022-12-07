/**
 * @Author: fangwen
 * @Description:
 * @File: sfp0000110
 * @Version: 1.0.0
 * @Date: 2022/3/28 11:08 AM
 */

package models

type FP000110I struct {
	PartClearingCode string `json:"partClearingCode"`
	Currency         string `json:"currency"`
	RequestStatus    string `json:"requestStatus"`
	RequestStartDate string `json:"requestStartDate"`
	RequestEndDate   string `json:"requestEndDate"`
	PageNo           int    `json:"pageNo"`
	PageSize         int    `json:"pageSize"`
}

type FP000110O struct {
	TotalCount int             `json:"totalCount"`
	Records    []BalAdjustItem `json:"records"`
}
type BalAdjustItem struct {
	Id               int64   `json:"id"`
	PartClearingCode string  `json:"partClearingCode"`
	Currency         string  `json:"currency"`
	AdjustDirection  string  `json:"adjustDirection"`
	AdjustAmount     float64 `json:"adjustAmount"`
	RequestDate      string  `json:"requestDate"`
	RequestStatus    string  `json:"requestStatus"`
	Action           string  `json:"action"`
	Comment          string  `json:"comment"`
}

func (*FP000110I) GetServiceKey() string {
	return "FP000110"
}
