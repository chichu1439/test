/**
 * @Author: fangwen
 * @Description:
 * @File: sfp9000049
 * @Version: 1.0.0
 * @Date: 2022/3/18 2:01 AM
 */

package models

type FP9000054I struct {
	ParticipantClearingCode string `json:"participantClearingCode" validate:"required"`
	ContactName             string `json:"contactName"`
	ContactMethod           string `json:"contactMethod"`
	ContactInformation      string `json:"contactInformation"`
	Priority                string `json:"priority"  validate:"required"`
	Necessary               string `json:"necessary"`
	Remark                  string `json:"remark"`
	Label                   string `json:"label"`
	Id                      int64  `json:"id" validate:"required"`
}

type FP9000054O struct {
	AssociatedStates bool `json:"associatedStates"`
}

func (*FP9000054I) GetServiceKey() string {
	return "FP900054"
}
