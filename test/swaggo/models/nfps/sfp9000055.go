/**
 * @Author: fangwen
 * @Description:
 * @File: sfp9000049
 * @Version: 1.0.0
 * @Date: 2022/3/18 2:01 AM
 */

package models

import "time"

type FP900055I struct {
	ParticipantClearingCode []string `json:"participantClearingCode"`
	FilterLabel             string   `json:"filterLabel"`
}

type FP900055O struct {
	Data []ParticipantLists `json:"data"`
}

type ParticipantLists struct {
	ParticipantClearingCode string    `json:"participantClearingCode"`
	ContactName             string    `json:"contactName"`
	ContactMethod           string    `json:"contactMethod"`
	ContactInformation      string    `json:"contactInformation"`
	Priority                string    `json:"priority"`
	Necessary               string    `json:"necessary"`
	Remark                  string    `json:"remark"`
	Label                   string    `json:"label"`
	CreateTime              time.Time `json:"createTime"`
	UpdateTime              time.Time `json:"updateTime"`
}

func (*FP900055I) GetServiceKey() string {
	return "FP900055"
}
