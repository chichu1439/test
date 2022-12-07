package models

type SMP0000005I struct {
	PhoneNo     string `json:"phoneNo"`
	IdNo        string `json:"idNo"`
	Email       string `json:"email"`
	RequestType string `json:"requestType" validate:"required"`
}

type SMP0000005O struct {
	ResponseDate  string `json:"responsedate"` // yyyy-MM-ddTHH:mm:ss.SSSZ
	RequestNumber string `json:"requestnumber"`
	SentTo        string `json:"sentTo"`
	SuccessCode   int    `json:"successcode"`
	SystemRefNo   string `json:"systemrefno"`
	OtpPrefix     string `json:"otpPrefix"`
	ExpiryTime    string `json:"expiryTime"`
	ApSerNo       string `json:"apSerNo"`
}

func (i *SMP0000005I) GetServiceKey() string {
	return "MP000005"
}
