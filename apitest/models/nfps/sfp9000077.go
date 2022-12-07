/**
 * @Author: fangwen
 * @Description:
 * @File: sfp9000077
 * @Version: 1.0.0
 * @Date: 2022/4/22 7:39 AM
 */

package models

type FP900077I struct {

}

type FP900077O struct {
	ScenarioInfoEnum []ScenarioInfoEnum `json:"scenarioInfoEnum"`
}

func (*FP900077I) GetServiceKey() string {
	return "FP900077"
}
