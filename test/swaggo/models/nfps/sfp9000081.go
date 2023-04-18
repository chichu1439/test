/**
 * @Author: fangwen
 * @Description:
 * @File: sfp9000080
 * @Version: 1.0.0
 * @Date: 2022/4/26 7:20 PM
 */

package models

type FP900081I struct {
	SceneName string `json:"sceneName"`
}

type FP900081O struct {
	ApiId []int64 `json:"apiId"`
}

func (*FP900081I) GetServiceKey() string {
	return "FP900081"
}
