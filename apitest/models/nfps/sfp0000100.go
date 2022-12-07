/**
 * @Author: fangwen
 * @Description:
 * @File: sfp0000100
 * @Version: 1.0.0
 * @Date: 2022/3/28 11:08 AM
 */

package models


type FP000100I struct {
}

type FP000100O struct {
	Data string `json:"data"`
}

func (*FP000100I) GetServiceKey() string {
	return "FP000100"
}
