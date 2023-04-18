/**
 * @Author: fangwen
 * @Description:
 * @File: sfp0000103
 * @Version: 1.0.0
 * @Date: 2022/3/28 11:08 AM
 */

package models

type FP000104I struct {
	NettingIntervals           int    `json:"nettingIntervals"` //unit hour,default 24 Hours 5
	StartSwitchTime            string `json:"startSwitchTime"`  //format HH:MM 02:01
	NettingEnableStatus        string `json:"nettingEnableStatus"`
	NettingBatchDelayIntervals int    `json:"nettingBatchDelayIntervals"`
	EffectiveDate              string `json:"effectiveDate"`
	ExpireDate                 string `json:"expireDate"`
}

type FP000104O struct {
	AzkabanMsg string
}

func (*FP000104I) GetServiceKey() string {
	return "FP000104"
}
