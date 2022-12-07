/**
 * @Author: fangwen
 * @Description:
 * @File: sfp1000013
 * @Version: 1.0.0
 * @Date: 2022/3/30 11:03 AM
 */

package models

type FP1000013I struct {
	PageNum         int    `json:"pageNum"`
	PageRecordCount int    `json:"pageRecordCount"`
	Name            string `json:"name"`
	Path            string `json:"path"`
	EventId         string `json:"eventId"`
}

type FP1000013O struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	EventId string `json:"eventId"`
	Path    string `json:"path"`
	EndPort string `json:"endPort"`
	Method  string `json:"method"`
}

func (*FP1000013I) GetServiceKey() string {
	return "FP100013"
}
