package models

// desc /auth/login/customerNo
type SMP0000031I struct {
	AcctNo   string `json:"acctNo" validate:"required,max=20"` // Y 20 Account no to pay
	ReqRefNo string `json:"reqRefNo" validate:"max=40"`        // Y 20 Request Reference No
}

type SMP0000031O struct {
	Code      string `json:"code"`
	RespCode  string `json:"respCode"`
	RespDesc  string `json:"respDesc"`
	ReqRefNo  string `json:"reqRefNo"`
	RespRefNo string `json:"respRefNo"`
}

func (i *SMP0000031I) GetServiceKey() string {
	return "MP000031"
}
