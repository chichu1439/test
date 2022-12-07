package models

type SMP9CU0005I struct {
	PhoneNo  string `json:"phoneNo" validate:"required"`
	OtpRefNo string `json:"otpRefNo" validate:"required"`
	DeviceId string `json:"deviceId" validate:"required"`
}

type SMP9CU0005O struct{}

func (i *SMP9CU0005I) GetServiceKey() string {
	return "MP9CU005"
}
