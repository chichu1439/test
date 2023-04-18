/**
 * @Author: fangwen
 * @Description:
 * @File: sfp9000100
 * @Version: 1.0.0
 * @Date: 2022/3/28 11:05 AM
 */

package models

type FP900101I struct {
}

type FP900101O struct {
	NettingId           string `json:"netting_id"`
	NettingDate         string `json:"nettingDate"`
	NettingSequence     int    `json:"nettingSequence"`
	NettingEnableStatus string `json:"nettingEnableStatus"`
	EffectiveDate       string `json:"effective_date"`
	ExpireDate          string `json:"expire_date"`
}

func (*FP900101I) GetServiceKey() string {
	return "FP900101"
}
