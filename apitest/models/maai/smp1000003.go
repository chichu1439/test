package models

type SMP1000003I struct {
	IDImageFront string `json:"idImageFront" validate:"required"`
	IDImageBack  string `json:"idImageBack" validate:"required"`
	ReqRefNo     string `json:"reqRefNo" validate:"omitempty,max=20"`
}

type SMP1000003Data struct {
	CifNo           string `json:"cifNo"`
	NameTh          string `json:"nameTh"`
	LastnameTh      string `json:"lastnameTh"`
	NameEn          string `json:"nameEn"`
	LastnameEn      string `json:"lastnameEn"`
	Addr            string `json:"addr"`
	BirthDate       string `json:"birthDate"`
	BirthDateText   string `json:"birthDateText"`
	IssuedDate      string `json:"issuedDate"`
	IssuedDateText  string `json:"issuedDateText"`
	ExpiredDate     string `json:"expiredDate"`
	ExpiredDateText string `json:"expiredDateText"`
	LaserID         string `json:"laserId"`
}

type SMP1000003O struct {
	RespCode  string         `json:"respCode"`
	RespDesc  string         `json:"respDesc"`
	ReqRefNo  string         `json:"reqRefNo"`
	RespRefNo string         `json:"respRefNo"`
	Data      SMP1000003Data `json:"data"`
}

func (i *SMP1000003I) GetServiceKey() string {
	return "MP100003"
}
