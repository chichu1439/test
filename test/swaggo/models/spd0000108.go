//Version: v1
package models

type PD000108Request struct {
	Code string `json:"code" validate:"required" description:"code"`
}

type PD000108Response struct {
}

func (*PD000108Request) GetServiceKey() string {
	return "PD000108"
}
