//Version: v0.0.1
package models

type SV000027I struct {
	CustomerID         string `json:"customerId" validate:"required"`
	DepositProofNumber string `json:"depositProofNumber" validate:"required"`
}
type SV000027O struct {
	DepositProofNumber string          `json:"depositProofNumber"`
	DepositProofType   string          `json:"depositProofType"`
	DepositProofStatus string          `json:"depositProofStatus"`
	TotalProofAmount   float64         `json:"totalProofAmount"`
	TimeProofDate      string          `json:"timeProofDate"`
	ExpiryProofDate    string          `json:"expiryProofDate"`
	RecipientName      string          `json:"recipientName"`
	DeliveryMethod     string          `json:"deliveryMethod"`
	Email              string          `json:"email"`
	PostalCode         string          `json:"postalCode"`
	NationCode         string          `json:"nationCode"`
	Address1           string          `json:"address1"`
	Address2           string          `json:"address2"`
	Address3           string          `json:"address3"`
	Address            string          `json:"address"`
	ReceiverName       string          `json:"receiverName"`
	CustomerID         string          `json:"customerId"`
	AccountName        string          `json:"accountName"`
	CreateDate         string          `json:"createDate"`
	AgreementList      []AgreementInfo `json:"agreementList"`
}

func (*SV000027I) GetServiceKey() string {
	return "SV000027"
}
