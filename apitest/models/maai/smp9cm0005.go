package models

type SMP9CM0005I struct {
	AgreementId string `json:"agreementId"`
}

type SMP9CM0005O struct {
	AgreementId      string `json:"agreementId"`
	AgreementType    string `json:"agreementType"`
	AgreementVer     int    `json:"agreementVer"`
	AgreementExpDate string `json:"agreementExpDate"`
	AgreementEffDate string `json:"agreementEffDate"`
}

func (i *SMP9CM0005I) GetServiceKey() string {
	return "MP9CM005"
}
