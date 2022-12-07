/**
 * @Author: fangwen
 * @Description:
 * @File: sfp1000007
 * @Version: 1.0.0
 * @Date: 2022/3/18 1:45 AM
 */

package models

type FP100007I struct {
	ClearingCodes []string `json:"clearingCodes"`
	PageSize      int      `json:"pageSize"`
	Page          int      `json:"page"`
	Details       bool     `json:"details"`
	Roule         string   `json:"roule"`
	Label         string   `json:"label"`
}

type FP100007O struct {
	Data []ParticipantList `json:"data"`
	Rows int64             `json:"rows"`
}

func (*FP100007I) GetServiceKey() string {
	return "FP100007"
}
