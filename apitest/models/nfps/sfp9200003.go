package models

type FP920003I struct {
	AccountKeyNum           string `json:"accountKeyNum"`
	Lang                    string `json:"lang"`
	FullName                string `json:"fullName"`
	DisplayName             string `json:"displayName"`
	AccountNum              string `json:"accountNum"`
	BankCode                string `json:"bankCode"`
	DefaultIndicator        string `json:"defaultIndicator"`
	PurposeCode             string `json:"purposeCode"`
	CustomerType            string `json:"customerType"`
	SupportOptionCode       string `json:"supportOptionCode"`
	ProxyUnbindFlag         string `json:"proxyUnbindFlag"`
	PayOrder                int    `json:"payOrder"`
	ParticipantClearingCode string `json:"participantClearingCode"`
	CustomerId              string `json:"customer_id"`
}
type FP920003O struct {
}

func (*FP920003I) GetServiceKey() string {
	return "FP920003"
}
