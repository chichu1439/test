/**
 * @Author: fangwen
 * @Description:
 * @File: sfp9000059
 * @Version: 1.0.0
 * @Date: 2022/4/16 11:21 PM
 */

package models


type FP9000059I struct {

}



type FP9000059O struct {
	HostTestLoad []HostTestLoad `json:"hostTestLoad"`
	ScenarioType []string `json:"scenarioType"`
}

type HostTestLoad struct {
	Ip string `json:"ip"`
	Name string `json:"name"`
	Path string `json:"path"`
	Port string `json:"port"`
}


func (* FP9000059I) GetServiceKey() string {
	return "FP900059"
}