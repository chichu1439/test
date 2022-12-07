package models

type SMP9CM0004I struct {
	AgreementType string `json:"agreementType"`
	Language      string `json:"language"`
}
type SMP9CM0004O struct {
	AgreementId          string `json:"agreementId"`
	AgreementType        string `json:"agreementType"`
	AgreementVer         int    `json:"agreementVer"`
	Language             string `json:"language"`
	AgreementContent     string `json:"agreementContent"`
	AgreementLongContent string `json:"agreementLongContent"`
	AgreementStatus      string `json:"agreementStatus"`
	AgreementEffDateTime string `json:"agreementEffDateTime"`
	AgreementExpDateTime string `json:"agreementExpDateTime"`
}
type Agreements struct {
}

func (i SMP9CM0004I) GetServiceKey() string {
	return "MP9CM004"
}
