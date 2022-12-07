package models

type SMP1000009I struct {
	OldPwd   string `json:"oldPwd" validate:"required,max=1024"`
	NewPwd   string `json:"newPwd" validate:"required,max=1024"`
	ReqRefNo string `json:"reqRefNo" validate:"omitempty,max=20"`
}

type SMP1000009O struct {
	RespCode  string `json:"respCode"`
	RespDesc  string `json:"respDesc"`
	ReqRefNo  string `json:"reqRefNo"`
	RespRefNo string `json:"respRefNo"`
}

func (i *SMP1000009I) GetServiceKey() string {
	return "MP100009"
}
