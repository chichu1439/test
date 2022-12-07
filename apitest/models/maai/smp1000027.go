package models

type SMP1000027I struct {
	CifNo    string `json:"cifNo" validate:"required,max=20"`
	ReqRefNo string `json:"reqRefNo" validate:"omitempty,max=20"`
}

type SMP1000027O struct {
	NameTh     string `json:"nameTh"`
	LastnameTh string `json:"lastnameTh"`
	AccountNo  string `json:"accountNo"`
	WalletID   string `json:"walletId"`
	RespCode   string `json:"respCode"`
	RespDesc   string `json:"respDesc"`
	ReqRefNo   string `json:"reqRefNo"`
	RespRefNo  string `json:"respRefNo"`
}

func (i *SMP1000027I) GetServiceKey() string {
	return "MP100027"
}
