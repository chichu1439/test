/**
 * @Author: fangwen
 * @Description:
 * @File: sfp9000100
 * @Version: 1.0.0
 * @Date: 2022/3/28 11:05 AM
 */

package models

 type FP900100I struct {
	RequestData []byte `json:"requestData"`
}

type FP900100O struct {
	Data string `json:"data"`
}

func (*FP900100I) GetServiceKey() string {
	return "FP900100"
}
