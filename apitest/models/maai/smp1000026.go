package models

// desc /auth/login/customerNo
type SMP1000026I struct {
	AcctNo   string `json:"acctNo" validate:"required"` // Y 20 Account no to pay
	ReqRefNo string `json:"reqRefNo"`                   // Y 20 Request Reference No
}

type SMP1000026O struct {
	Code      string `json:"code"`
	RespCode  string `json:"respCode"`
	RespDesc  string `json:"respDesc"`
	ReqRefNo  string `json:"reqRefNo"`
	RespRefNo string `json:"respRefNo"`
}

type SMP1000026IOutgress struct {
	SMP1000026I
}

func (i *SMP1000026I) GetServiceKey() string {
	return "MP100026"
}

func (i *SMP1000026IOutgress) GetServiceKey() string {
	return "MP300026"
}
