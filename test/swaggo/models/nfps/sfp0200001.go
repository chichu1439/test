package models

type FP020001I struct {
	AccountKeyNum           string `json:"accountKeyNum" validate:"required"`
	AccountKeyType          string `json:"accountKeyType" validate:"required"`
	Lang                    string `json:"lang"`
	FullName                string `json:"fullName"`
	CustomerId              string `json:"customerId"`
	AccountNum              string `json:"accountNum" validate:"required"`
	BankCode                string `json:"bankCode" validate:"required"`
	DefaultIndicator        string `json:"defaultIndicator"`
	PurposeCode             string `json:"purposeCode"`
	CustomerType            string `json:"customerType"`
	SupportOptionCode       string `json:"supportOptionCode"`
	ProxyUnbindFlag         string `json:"proxyUnbindFlag"`
	PayOrder                int    `json:"payOrder"`
	ParticipantClearingCode string `json:"participantClearingCode" validate:"required"`
}
type FP020001O struct {
	// AddressingInfo
}

func (*FP020001I) GetServiceKey() string {
	return "FP020001"
}
