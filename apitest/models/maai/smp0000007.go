package models

type SMP0000007I struct {
	SystemRefNo string `json:"systemRefNo" validate:"required"`
}

type SMP0000007O struct {
	RequestDate   string `json:"requestDate"`
	RequestNumber string `json:"requestnumber"`
	SuccessCode   int    `json:"successcode"`
	SystemRefNo   string `json:"systemrefno"`
	SentTo        string `json:"sentTo"`
	OtpPrefix     string `json:"otpPrefix"`
	Retry         int    `json:"retry"`
}

func (i *SMP0000007I) GetServiceKey() string {
	return "MP000007"
}
