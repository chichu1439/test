//Version: v0.0.1
package models

type SV000025I struct {
	DepositProofNumber string `json:"depositProofNumber"`
	CustomerID         string `json:"customerId" `
	IDNation           string `json:"idNation"`
	IDType             string `json:"idType" validate:"required"`
	IDNumber           string `json:"idNumber" validate:"required"`
	IssuedStartDate    string `json:"issuedStartDate" `
	IssuedEndDate      string `json:"issuedEndDate" `
	PageNo             int    `json:"pageNo" `
	PageRecCount       int    `json:"pageRecCount" `
}

type SV000025O struct {
	TotalRecord int               `json:"totalRecord"`
	TotalPage   int               `json:"totalPage"`
	Records     []SV000025ORecord `json:"records"`
}

type SV000025ORecord struct {
	GlobalBusinessSeqNo string  `json:"globalBusinessSeqNo" `
	SourceBusinessSeqNo string  `json:"sourceBusinessSeqNo" `
	DepositProofNumber  string  `json:"depositProofNumber" `
	DepositProofType    string  `json:"depositProofType" `
	TotalProofAmount    float64 `json:"totalProofAmount" `
	TimeProofDate       string  `json:"timeProofDate" `
	ExpiryProofDate     string  `json:"expiryProofDate" `
	RecipientName       string  `json:"recipientName" `
	DeliveryMethod      string  `json:"deliveryMethod" `
	Email               string  `json:"email" `
	PostalCode          string  `json:"postalCode" `
	NationCode          string  `json:"nationCode" `
	Address1            string  `json:"address1" `
	Address2            string  `json:"address2" `
	Address3            string  `json:"address3" `
	Address             string  `json:"address" `
	ReceiverName        string  `json:"receiverName" `
	CustomerID          string  `json:"customerId" `
	AccountName         string  `json:"accountName" `
	CreateDate          string  `json:"createDate" `
}

func (*SV000025I) GetServiceKey() string {
	return "SV000025"
}
