//Version: v1.0.0.0
package models

type SSV0000053I struct {
	AgreementId string `validate:"required,max=30" json:"agreementId" description:"Agreement id"`
}

type SSV0000053O struct {
	Status string `json:"status" description:"Status"`
}

func (*SSV0000053I) GetServiceKey() string {
	return "SV000053"
}
