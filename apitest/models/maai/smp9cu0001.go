package models

type SMP9CU0001I struct {
	AgreementId string `json:"agreementId" validate:"required"`
	CustomerNo  string `json:"customerNo" validate:"required"`
	//AcceptResult     string `json:"acceptResult"`
	AgreementType    string `json:"agreementType"`
	AgreementVer     int    `json:"agreementVer"`
	SignFlag         string `json:"signFlag"`
	AgreementLang    string `json:"agreementLang"`
	AgreementExpDate string `json:"agreementExpDate"`
	AgreementEffDate string `json:"agreementEffDate"`
}

type SMP9CU0001O struct {
}

func (i *SMP9CU0001I) GetServiceKey() string {
	return "MP9CU001"
}
