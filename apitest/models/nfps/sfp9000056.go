/**
 * @Author: fangwen
 * @Description:
 * @File: sfp9000049
 * @Version: 1.0.0
 * @Date: 2022/3/18 2:01 AM
 */

package models


type FP9000056I struct {
	ParticipantClearingCode string `json:"participantClearingCode"`
}


type FP9000056O struct {
	Inexistence bool `json:"inexistence"`
}

func (* FP9000056I) GetServiceKey() string {
	return "FP900056"
}

