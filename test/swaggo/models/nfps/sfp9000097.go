/**
 * @Author: fangwen
 * @Description:
 * @File: sfp9000097
 * @Version: 1.0.0
 * @Date: 2022/3/27 1:19 PM
 */

package models

type FP900097I struct {
	ResponseCode string `json:"responseCode"`
	TraceId      string `json:"traceId"  validate:"required"`
	Data         string `json:"data"  validate:"required"`
}

type FP900097O struct {
	Data string `json:"data"`
}

func (*FP900097I) GetServiceKey() string {
	return "FP900097"
}
