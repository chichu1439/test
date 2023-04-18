//Version: v1.0.0
package models

type CM000053I struct {
	ShareCode string `json:"shareCode"`
}

type CM000053O struct {
	AgreementId   string `json:"agreementId"`
	AgreementType string `json:"agreementType"`
}

func (*CM000053I) GetServiceKey() string {
	return "CM000053"
}
