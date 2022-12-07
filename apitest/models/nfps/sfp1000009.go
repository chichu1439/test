/**
 * @Author: fangwen
 * @Description:
 * @File: sfp1000007
 * @Version: 1.0.0
 * @Date: 2022/3/18 1:45 AM
 */

package models

type FP100009I struct {
	ParticipantClearingCode string `json:"participantClearingCode"`
	ContactName             string `json:"contactName"`
	ContactMethod           string `json:"contactMethod"`
	ContactInformation      string `json:"contactInformation"`
	Priority                string `json:"priority"`
	Necessary               string `json:"necessary"`
	Remark                  string `json:"remark"`
	Label                   string `json:"label"`
}

type FP100009O struct {
	AssociatedStates bool `json:"associatedStates"`
}

func (*FP100009I) GetServiceKey() string {
	return "FP100009"
}
