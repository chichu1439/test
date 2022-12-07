/**
 * @Author: fangwen
 * @Description:
 * @File: sfp9000049
 * @Version: 1.0.0
 * @Date: 2022/3/18 2:01 AM
 */

package models

type FP9000049I struct {
	ClearingCode []string `json:"participantClearingCode"`
	PageSize     int      `json:"pageSize"`
	Page         int      `json:"page"`
	Details      bool     `json:"details"`
	Label        string   `json:"label"`
}

type FP9000049O struct {
	Data []ParticipantList `json:"data"`
	Rows int64             `json:"rows"`
}

func (*FP9000049I) GetServiceKey() string {
	return "FP900049"
}
