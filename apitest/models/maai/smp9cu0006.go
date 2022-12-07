package models

import "time"

type SMP9CU0006I struct {
	OtpId    int    `json:"otpID"`
	PhoneNo  string `json:"phoneNo"`
	OtpRefNo string `json:"otpRefNo"`
	DeviceId string `json:"deviceId"`
}

type SMP9CU0006OData struct {
	OtpId    int       `json:"otpId"`
	PhoneNo  string    `json:"phoneNo"`
	OtpRefNo string    `json:"otpRefNo"`
	DeviceId string    `json:"deviceId"`
	CreateDt time.Time `json:"createDt"`
	UpdateDt time.Time `json:"updateDt"`
}

type SMP9CU0006O struct {
	Data []SMP9CU0006OData `json:"data"`
}

func (i *SMP9CU0006I) GetServiceKey() string {
	return "MP9CU006"
}
