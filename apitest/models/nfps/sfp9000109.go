/**
 * @Author: fangwen
 * @Description:
 * @File: sfp0000110
 * @Version: 1.0.0
 * @Date: 2022/3/28 11:08 AM
 */

package models

type FP900109I struct {
	PartClearingCode string `json:"partClearingCode"`
	Currency         string `json:"currency"`
	NettingStartDate string `json:"nettingStartDate"`
	NettingEndDate   string `json:"nettingEndDate"`
	NettingId        string `json:"nettingId"`
	PageNo           int    `json:"pageNo"`
	PageSize         int    `json:"pageSize"`
}

type FP900109O struct {
	TotalCount int                 `json:"totalCount"`
	Records    []NettingResultItem `json:"records"`
}
type NettingResultItem struct {
	PartClearingCode string  `json:"partClearingCode"`
	Currency         string  `json:"currency"`
	NettingDirection string  `json:"nettingDirection"`
	NettingAmount    float64 `json:"nettingAmount"`
	NettingDate      string  `json:"nettingDate"`
	NettingId        string  `json:"nettingId"`
	TotDebitAmount   float64 `json:"totDebitAmount"`
	TotCreditAmount  float64 `json:"totCreditAmount"`
}

func (*FP900109I) GetServiceKey() string {
	return "FP900109"
}
