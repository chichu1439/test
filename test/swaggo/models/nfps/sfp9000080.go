/**
 * @Author: fangwen
 * @Description:
 * @File: sfp9000080
 * @Version: 1.0.0
 * @Date: 2022/4/26 7:20 PM
 */

package models

import "time"

type FP900080I struct {
	TransactionType string `json:"transactionType"`
	StartTime       int64  `json:"startTime"`
	EndTime         int64  `json:"endTime"`
	TestId          string `json:"testId"`
	Page            int    `json:"page"`
	PageNum         int    `json:"pageNum"`
}

type FP900080O struct {
	TestRecord []TestResult `json:"testRecord"`
	TotalRows  int64        `json:"totalRows"`
}

type TestResult struct {
	TestId          string     `json:"testId"`
	ThreadNumber    int        `json:"threadNumber"`
	DurationSecond  int        `json:"durationSecond"`
	TransactionType string     `json:"transactionType"`
	LoadHost        string     `json:"loadHost"`
	DosKey          string     `json:"dosKey"`
	Message         string     `json:"message"`
	TestResults     string     `json:"testResults"`
	CaseNumber      int        `json:"caseNumber"`
	RecordStatus    int        `json:"recordStatus"`
	CreateTime      *time.Time `json:"createTime"`
}
type TestResult2 struct {
	TestId          string                 `json:"testId"`
	ThreadNumber    int                    `json:"threadNumber"`
	DurationSecond  int                    `json:"durationSecond"`
	TransactionType string                 `json:"transactionType"`
	LoadHost        string                 `json:"loadHost"`
	Message         string                 `json:"message"`
	DosKey          string                 `json:"dosKey"`
	TestResults     map[string]interface{} `json:"testResults"`
	CaseNumber      int                    `json:"caseNumber"`
	RecordStatus    int                    `json:"recordStatus"`
	CreateTime      *time.Time             `json:"createTime"`
}

func (*FP900080I) GetServiceKey() string {
	return "FP900080"
}
