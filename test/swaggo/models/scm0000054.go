//Version: v1
package models

type CM000054Request struct {
	ShareCode      string `json:"shareCode"`
	AgreementType  string `json:"agreementType"`
	AgreementId    string `json:"agreementId"`
	AccountNumber  string `json:"accountNumber"`
	LastModifyOrg  string `json:"lastModifyOrg"`
	LastModifyUser string `json:"lastModifyUser"`
}

type CM000054Response struct {
}

func (*CM000054Request) GetServiceKey() string {
	return "CM000054"
}
