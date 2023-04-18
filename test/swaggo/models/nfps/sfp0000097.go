/**
 * @Author: fangwen
 * @Description:
 * @File: sfp0000097
 * @Version: 1.0.0
 * @Date: 2022/3/27 1:18 PM
 */

package models

import "github.com/go-playground/validator/v10"

type FP000097I struct {
	TraceId string `json:"traceId"  validate:"required"`
	Data    string `json:"data"  validate:"required"`
}

type FP000097O struct {
	Data string `json:"data"`
}

func (*FP000097I) GetServiceKey() string {
	return "FP000097"
}

func (d *FP000097I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
