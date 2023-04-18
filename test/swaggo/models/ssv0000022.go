//Version: v1.0.0.0

package models

type SV000022I struct {
	MediumType                    string `json:"mediumType"`
	MediumNumber                  string `json:"mediumNumber"`
	CustomerId                    string `json:"customerId" validate:"required"`
	CustomerIdType                string `json:"customerIdType" validate:"required"`
	Currency                      string `json:"currency" validate:"required"`
	CashTranFlag                  string `json:"cashTranFlag" validate:"required"`
	ChannelNumber                 string `json:"channelNumber" `
	UsageCode                     string `json:"usageCode"`
	AgreementID                   string `json:"agreement" validate:"required"`
	AgreementType                 string `json:"agreementType" validate:"required"`
	ActivateEmployee              string `json:"activateEmployee"`
	ActivateAuthorizationEmployee string `json:"activateAuthorizationEmployee"`
}

type SV000022O struct {
}

func (*SV000022I) GetServiceKey() string {
	return "SV000022"
}
