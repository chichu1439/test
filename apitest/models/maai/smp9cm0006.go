package models

type SMP9CM0006I struct {
	AgreementType string `json:"agreementType"`
	AgreementId   string `json:"agreementId"`
	CustomerNo    string `json:"customerNo"`
	Language      string `json:"language"`
}
type SMP9CM0006O struct {
	AgreementId          string `json:"agreementId"`
	AgreementType        string `json:"agreementType"`
	AgreementVer         int    `json:"agreementVer"`
	AgreementContent     string `json:"agreementContent"`
	AgreementStatus      string `json:"agreementStatus"`
	AgreementEffDateTime string `json:"agreementEffDateTime"`
	AgreementExpDateTime string `json:"agreementExpDateTime"`
}

func (i SMP9CM0006I) GetServiceKey() string {
	return "MP9CM006"
}
