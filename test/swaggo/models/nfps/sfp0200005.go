package models

type FP020005I struct {
	AccountKeyNum           string `json:"accountKeyNum"`
	AccountKeyType          string `json:"accountKeyType"`
	CustomerId              string `json:"customerId"`
	AccountNum              string `json:"accountNum"`
	BankCode                string `json:"bankCode"`
	RegisterStartDate       string `json:"registerStartDate"`
	RegisterEndDate         string `json:"registerEndDate"`
	ParticipantClearingCode string `json:"participantClearingCode"`
	ParticipantClearingName string `json:"participantClearingName"`
}
type FP020005O struct {
	TotalCount     int            `json:"totalCount"`
	AddressingList []FP920005List `json:"addressingList"`
}

func (*FP020005I) GetServiceKey() string {
	return "FP020005"
}

type AddressingInfo struct {
	RegisterDateTime        string `json:"registerDateTime"`
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
	LastUpdateDateTime      string `json:"lastUpdateDateTime"`
	ParticipantClearingCode string `json:"participantClearingCode"`
}
