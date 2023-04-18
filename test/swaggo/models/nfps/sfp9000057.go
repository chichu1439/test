/**
 * @Author: fangwen
 * @Description:
 * @File: sfp9000049
 * @Version: 1.0.0
 * @Date: 2022/3/18 2:01 AM
 */

package models

type FP9000057I struct {
	ParticipantClearingCode string `json:"participantClearingCode"`
	Id                      int64  `json:"id" validate:"required"`
}

type FP9000057O struct {
	AssociatedStates bool `json:"associatedStates"`
}

func (*FP9000057I) GetServiceKey() string {
	return "FP900057"
}
