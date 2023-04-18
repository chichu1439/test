/**
 * @Author: fangwen
 * @Description:
 * @File: sfp9000078
 * @Version: 1.0.0
 * @Date: 2022/4/22 4:55 PM
 */

package models

type FP900078I struct {
	ApiId string `json:"apiId"`
}

type FP900078O struct {
	SendTemplate    string `json:"sendTemplate"`
	ReceiveTemplate string `json:"receiveTemplate"`
}

func (*FP900078I) GetServiceKey() string {
	return "FP900078"
}
