/**
 * @Author: fangwen
 * @Description:
 * @File: sfp9000058
 * @Version: 1.0.0
 * @Date: 2022/4/16 11:21 PM
 */

package models


type FP9000058I struct {

}



type FP9000058O struct {
	APIInfo []APIInfo
}

type APIInfo struct {
	Name string `json:"name"`
	ServiceId string `json:"serviceId"`
}

func (* FP9000058I) GetServiceKey() string {
	return "FP900058"
}
