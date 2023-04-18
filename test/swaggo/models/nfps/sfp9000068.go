//Version: v1
package models

type FP900068Request struct {
	ID int64 `json:"id" validate:"required"`
}

type FP900068Response struct {
	Status string `json:"status"`
}

func (*FP900068Request) GetServiceKey() string {
	return "FP900068"
}
