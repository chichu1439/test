/**
 * @Author: fangwen
 * @Description:
 * @File: sfp9000098
 * @Version: 1.0.0
 * @Date: 2022/3/27 4:15 PM
 */

package models

type FP900098I struct {
	TraceId string `json:"traceId"  validate:"required"`
}

type FP900098O struct {
	Data string `json:"data"`
}

func (*FP900098I) GetServiceKey() string {
	return "FP900098"
}
