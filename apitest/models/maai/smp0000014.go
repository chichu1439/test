package models

type SMP0000014I struct {
	PhoneNm         string                   `json:"phoneNm" validate:"required,max=12"`
	CustomerNo      string                   `json:"customerNo" validate:"required,max=50"`
	OtpRefNo        string                   `json:"otpRefNo" validate:"max=60"`
	Agreement       []Agreement              `json:"agreement"`
	Gender          string                   `json:"gender" validate:"max=10"`
	Pin             string                   `json:"pin" validate:"max=1024"`
	Occupation      string                   `json:"occupation" validate:"required,max=100"`
	OccupationOther string                   `json:"occupationOther" validate:"max=100"`
	AddressType     []SMP0000014IAddressType `json:"addressType"`
}
type SMP0000014IAddressType struct {
	AddressType string `json:"addressType"` //1-personal address 2-work address 3-contact address
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
type Agreement struct {
	AgreementType string `json:"agreementType"`
	AgreementVer  int    `json:"agreementVer"`
	AgreementId   string `json:"agreementId"`
	SignFlag      string `json:"signFlag"`
	AgreementLang string `json:"agreementLang"`
}
type SMP0000014O struct {
	CustomerNo string `json:"customerNo"`
}

func (i *SMP0000014I) GetServiceKey() string {
	return "MP000014"
}
