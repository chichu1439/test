package models

type SMP9CU0002I struct {
	AgreementType string `json:"agreementType"`
	AgreementId   string `json:"agreementId"`
	CustomerNo    string `json:"customerNo"`
	Language      string `json:"language"`
}

type SMP9CU0002O struct {
	CustomerNo        string `json:"customerNo"`
	AgreementId       string `json:"agreementId"`
	AgreementType     string `json:"agreementType"`
	AgreementVer      int    `json:"agreementVer"`
	AgreementStatus   string `json:"agreementStatus"`
	AgreementSignDate string `json:"agreementSignDate"`
	AgreementEffDate  string `json:"agreementEffDate"`
	AgreementExpDate  string `json:"agreementExpDate"`
}

func (i *SMP9CU0002I) GetServiceKey() string {
	return "MP9CU002"
}
