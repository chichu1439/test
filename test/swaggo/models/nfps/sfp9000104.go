/**
 * @Author: fangwen
 * @Description:
 * @File: sfp0000103
 * @Version: 1.0.0
 * @Date: 2022/3/28 11:08 AM
 */

package models

type FP900104I struct {
	NettingIntervals           int    //unit hour,default 24 Hours
	StartSwitchTime            string //format HH:MM:SS
	NettingEnableStatus        string
	NettingBatchDelayIntervals int
	EffectiveDate              string
	ExpireDate                 string
	ExecutionFrequency         string
}

type FP900104O struct {
}

func (*FP900104I) GetServiceKey() string {
	return "FP900104"
}
