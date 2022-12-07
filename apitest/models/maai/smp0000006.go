package models

type SMP0000006I struct {
	SystemRefNo string `json:"systemRefNo" validate:"required"`
	Otp         string `json:"otp" validate:"required"`
	LogTypeId   string `json:"logTypeId"`
}

type SMP0000006O struct {
	ResponseDate  string `json:"responsedate"`
	RequestNumber string `json:"requestnumber"`
	SuccessCode   string `json:"successcode"`
	SystemRefNo   string `json:"systemrefno"`
}

func (i *SMP0000006I) GetServiceKey() string {
	return "MP000006"
}
