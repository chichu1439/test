package models

type SMP0000020I struct {
	CustomerNo    string `json:"customerNo" validate:"required"`
	AgreementType string `json:"agreementType" validate:"required"`
	Language      string `json:"language"`
}
type SMP0000020O struct {
	AgreementId          string `json:"agreementId"`
	AgreementType        string `json:"agreementType"`
	AgreementVer         int    `json:"agreementVer"`
	AgreementContent     string `json:"agreementContent"`
	AgreementStatus      string `json:"agreementStatus"`
	AgreementEffDateTime string `json:"agreementEffDateTime"`
	AgreementExpDateTime string `json:"agreementExpDateTime"`
}

func (i SMP0000020I) GetServiceKey() string {
	return "MP000020"
}
