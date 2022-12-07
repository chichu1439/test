/**
 * @Author: fangwen
 * @Description:
 * @File: sfp0000098
 * @Version: 1.0.0
 * @Date: 2022/3/27 1:18 PM
 */

package models

import "github.com/go-playground/validator/v10"

type FP000098I struct {
	TraceId string `json:"traceId"  validate:"required"`

}

type FP000098O struct {
	Data string `json:"data"`
}

func (*FP000098I) GetServiceKey() string {
	return "FP000098"
}

func (d *FP000098I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}