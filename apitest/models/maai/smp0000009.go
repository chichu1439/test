package models

type SMP0000009I struct {
	CustomerNo  string `json:"customerNo" description:"customer no" validate:"required,max=50"`
	AgreementId string `json:"agreementId" description:"agreement id" validate:"required,max=20"`
	SignFlag    string `json:"signFlag" description:"sign flag, Y-sign N-reject" validate:"required,max=1,oneof=Y N"`
	Code        string `json:"code" description:"agreement code"  validate:"required,max=20"`
	Version     string `json:"version" description:"agreement version" validate:"required,max=20"`
}

type SMP0000009O struct {
}

func (i *SMP0000009I) GetServiceKey() string {
	return "MP000009"
}
