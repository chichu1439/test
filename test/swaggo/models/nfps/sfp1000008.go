/**
 * @Author: fangwen
 * @Description:
 * @File: sfp1000007
 * @Version: 1.0.0
 * @Date: 2022/3/18 1:45 AM
 */

package models

type FP100008I struct {
	Keyword  map[string]string
	Sort     string `json:"sort"`
	Order    string `json:"order"`
	PageSize int64  `json:"pageSize"`
	PageNum  int64  `json:"pageNum"`
}

type FP100008O struct {
	EnvId string `json:"envId"`
	AigId string `json:"aigId"`
}

func (*FP100008I) GetServiceKey() string {
	return "FP100008"
}

type AigDetil struct {
	ServiceID   string `json:"ServiceID"`
	ServiceName string `json:"ServiceName"`
	Status      string `json:"Status"`
}
