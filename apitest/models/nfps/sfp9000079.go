/**
 * @Author: fangwen
 * @Description:
 * @File: sfp9000079
 * @Version: 1.0.0
 * @Date: 2022/4/26 7:20 PM
 */

package models

import "time"

type FP900079I struct {
	TestId string `json:"testId"`
	ThreadNumber int `json:"threadNumber"`
	DurationSecond int `json:"durationSecond"`
	TransactionType string `json:"transactionType"`
	LoadHost string `json:"loadHost"`
	DosKey string `json:"dosKey"`
	Message string `json:"message"`
	TestResults string `json:"testResults"`
	CaseNumber int `json:"caseNumber"`
	RecordStatus int `json:"recordStatus"`
	CreateTime *time.Time `json:"createTime"`
	UpdateTime *time.Time `json:"updateTime"`
}

type FP900079O struct {
}



func (*FP900079I) GetServiceKey() string {
	return "FP900079"
}