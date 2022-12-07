package models

type SMP1000002I struct {
	Channel     string     `json:"channel" validate:"required,max=20"`
	Date        string     `json:"date" validate:"required,max=8"`
	Code        string     `json:"code" validate:"required,max=20"`
	Description string     `json:"description" validate:"omitempty,max=50"`
	Version     string     `json:"version" validate:"required,max=20"`
	MobileNo    []MobileNo `json:"mobileNo" validate:"gt=0"`
	ReqRefNo    string     `json:"reqRefNo" validate:"omitempty,max=20"`
}
type MobileNo struct {
	Subject string `json:"subject" validate:"required,max=100"`
	Action  bool   `json:"action" validate:"required"`
}

type SMP1000002O struct {
	RespCode  string `json:"respCode"`
	RespDesc  string `json:"respDesc"`
	ReqRefNo  string `json:"reqRefNo"`
	RespRefNo string `json:"respRefNo"`
}

func (i *SMP1000002I) GetServiceKey() string {
	return "MP100002"
}
