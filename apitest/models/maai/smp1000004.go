package models

type SMP1000004I struct {
	Password   string `json:"password" validate:"required,max=1024"`
	NameTh     string `json:"nameTh" validate:"required,max=50"`
	LastnameTh string `json:"lastnameTh" validate:"required,max=50"`
	BirthDate  string `json:"birthDate" validate:"required,max=8"`
	CifNo      string `json:"cifNo" validate:"required,max=13"`
	CifLaserNo string `json:"cifLaserNo" validate:"required,max=12"`
	ReqRefNo   string `json:"reqRefNo" validate:"omitempty,max=20"`
}

type SMP1000004O struct {
	CustomerNo string `json:"customerNo"`
	Token      string `json:"token"`
	RespCode   string `json:"respCode"`
	RespDesc   string `json:"respDesc"`
	ReqRefNo   string `json:"reqRefNo"`
	RespRefNo  string `json:"respRefNo"`
}

func (i *SMP1000004I) GetServiceKey() string {
	return "MP100004"
}
