/**
 * @Author: fangwen
 * @Description:
 * @File: sfp1000007
 * @Version: 1.0.0
 * @Date: 2022/3/18 1:45 AM
 */

package models

type FP100010I struct {
	ParticipantClearingCode string `json:"participantClearingCode"`
	ContactName             string `json:"contactName"`
	ContactMethod           string `json:"contactMethod"`
	ContactInformation      string `json:"contactInformation"`
	Priority                string `json:"priority"`
	Necessary               string `json:"necessary"`
	Label                   string `json:"label"`
	Remark                  string `json:"remark"`
	Id                      int64  `json:"id"`
}

type FP100010O struct {
	AssociatedStates bool `json:"associatedStates"`
}

func (*FP100010I) GetServiceKey() string {
	return "FP100010"
}
