package models

type FP020003I struct {
	AccountKeyNum           string `json:"accountKeyNum" validate:"required"`
	Lang                    string `json:"lang"`
	FullName                string `json:"fullName"`
	DisplayName             string `json:"displayName"`
	AccountNum              string `json:"accountNum"`
	BankCode                string `json:"bankCode" validate:"required"`
	DefaultIndicator        string `json:"defaultIndicator"`
	PurposeCode             string `json:"purposeCode"`
	CustomerType            string `json:"customerType"`
	SupportOptionCode       string `json:"supportOptionCode"`
	ProxyUnbindFlag         string `json:"proxyUnbindFlag"`
	ParticipantClearingCode string `json:"participantClearingCode"`
	CustomerId              string `json:"customerId"`
}
type FP020003O struct {
	// AddressingInfo
}

func (*FP020003I) GetServiceKey() string {
	return "FP020003"
}
