package models

type FP920005I struct {
	AccountKeyNum           string `json:"accountKeyNum"`
	AccountKeyType          string `json:"accountKeyType"`
	CustomerId              string `json:"customerId"`
	AccountNum              string `json:"accountNum"`
	BankCode                string `json:"bankCode"`
	RegisterStartDate       string `json:"registerStartDate" validate:"datetime=2006-01-02"`
	RegisterEndDate         string `json:"registerEndDate" validate:"datetime=2006-01-02"`
	ParticipantClearingCode string `json:"participantClearingCode"`
	PageNo                  int    `json:"pageNo"`
	PageSize                int    `json:"pageSize"`
}
type FP920005O struct {
	TotalCount     int              `json:"totalCount"`
	AddressingList []AddressingInfo `json:"addressingList"`
}
type FP920005List struct {
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
	ParticipantClearingName string `json:"participantClearingName"`
}

func (*FP920005I) GetServiceKey() string {
	return "FP920005"
}
