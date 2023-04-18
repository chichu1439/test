package models

type FP920001I struct {
	RegisterDateTime        string `json:"registerDateTime"`
	LastUpdateDateTime      string `json:"lastUpdateDateTime"`
	AccountKeyNum           string `json:"accountKeyNum"`
	AccountKeyType          string `json:"accountKeyType"`
	Lang                    string `json:"lang"`
	FullName                string `json:"fullName"`
	DisplayName             string `json:"displayName"`
	CustomerId              string `json:"customerId"`
	AccountNum              string `json:"accountNum"`
	BankCode                string `json:"bankCode"`
	DefaultIndicator        string `json:"defaultIndicator"`
	PurposeCode             string `json:"purposeCode"`
	CustomerType            string `json:"customerType"`
	SupportOptionCode       string `json:"supportOptionCode"`
	ProxyUnbindFlag         string `json:"proxyUnbindFlag"`
	PayOrder                int    `json:"payOrder"`
	ParticipantClearingCode string `json:"participantClearingCode"`
	UpdateIndicatorFlag     string `json:"updateIndicatorFlag"`
}
type FP920001O struct {
	// AddressingInfo
}

func (*FP920001I) GetServiceKey() string {
	return "FP920001"
}
