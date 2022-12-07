package models

type SMP1000006I struct {
	MobileNo        string      `json:"mobileNo" validate:"required"`
	Gender          string      `json:"gender"`
	Occupation      string      `json:"occupation" validate:"required"`
	OccupationOther string      `json:"occupationOther"`
	ReqRefNo        string      `json:"reqRefNo"`
	PersonalAddr    AddressType `json:"personalAddr"`
	WorkAddr        AddressType `json:"workAddr"`
	ContactAddr     AddressType `json:"contactAddr"`
}
type AddressType struct {
	CompanyName string `json:"companyName"`
	AddrNo      string `json:"addrNo" validate:"required"`
	Bld         string `json:"bld"`
	Soi         string `json:"soi"`
	Road        string `json:"road" validate:"required"`
	District    string `json:"district" validate:"required"`
	City        string `json:"city" validate:"required"`
	Province    string `json:"province" validate:"required"`
	ZipCode     string `json:"zipCode" validate:"required"`
	CountryCode string `json:"countryCode" validate:"required"`
}
type SMP1000006O struct {
	CustomerNo string `json:"customerNo"`
	RespCode   string `json:"respCode"`
	RespDesc   string `json:"respDesc"`
	ReqRefNo   string `json:"reqRefNo"`
	RespRefNo  string `json:"respRefNo"`
}
type SMP1000006IOutgress struct {
	SMP1000006I
}

func (i *SMP1000006I) GetServiceKey() string {
	return "MP100006"
}

func (i *SMP1000006IOutgress) GetServiceKey() string {
	return "MP300006"
}
