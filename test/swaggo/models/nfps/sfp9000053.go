/**
 * @Author: fangwen
 * @Description:
 * @File: sfp9000049
 * @Version: 1.0.0
 * @Date: 2022/3/18 2:01 AM
 */

package models

import "time"

type FP9000053I struct {
	ParticipantClearingCode string    `json:"participantClearingCode" validate:"required"`
	ContactName             string    `json:"contactName" validate:"required"`
	ContactMethod           string    `json:"contactMethod" validate:"required"`
	ContactInformation      string    `json:"contactInformation" validate:"required"`
	Priority                string    `json:"priority" validate:"number"`
	Necessary               string    `json:"necessary"`
	Remark                  string    `json:"remark"`
	Label                   string    `json:"label"`
	CreateTime              time.Time `json:"create_time"`
	UpdateTime              time.Time `json:"update_time"`
}

type FP9000053O struct {
	AssociatedStates bool `json:"associatedStates"`
}

func (*FP9000053I) GetServiceKey() string {
	return "FP900053"
}
