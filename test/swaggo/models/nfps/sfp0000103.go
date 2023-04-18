/**
 * @Author: fangwen
 * @Description:
 * @File: sfp0000103
 * @Version: 1.0.0
 * @Date: 2022/3/28 11:08 AM
 */

package models

type FP000103I struct {
}

type FP000103O struct {
	NettingIntervals           int    `json:"nettingIntervals"` //unit hour,default 24 Hours
	StartSwitchTime            string `json:"startSwitchTime"`  //format HH:MM
	NettingEnableStatus        string `json:"nettingEnableStatus"`
	NettingBatchDelayIntervals int    `json:"nettingBatchDelayIntervals"`
	EffectiveDate              string `json:"effectiveDate"`
	ExpireDate                 string `json:"expireDate"`
}

func (*FP000103I) GetServiceKey() string {
	return "FP000103"
}
