package models

type FP010014I struct {
	MsgId    string `json:"msgId" validate:"required"`
	TransId  string `json:"transId" validate:"required"`
	PartCode string `json:"partCode" validate:"required"`
}

type FP010014O struct {
	TransactionType        string  `json:"transactionType"`
	Status                 string  `json:"status"`
	ReferenceNumber        string  `json:"referenceNumber"`
	TransactionId          string  `json:"transactionId"`
	EndToEndId             string  `json:"endToEndId"`
	InstructionId          string  `json:"instructionId"`
	MessageId              string  `json:"messageId"`
	PayerPartCode          string  `json:"payerPartCode"`
	TransactionDateTime    string  `json:"transactionDateTime"`
	PayerPartName          string  `json:"payerPartName"`
	PayeePartCode          string  `json:"payeePartCode"`
	PayeePartName          string  `json:"payeePartName"`
	Currency               string  `json:"currency"`
	PaymentCategoryPurpose string  `json:"paymentCategoryPurpose"`
	Amount                 float64 `json:"amount"`
	BalanceType            string  `json:"balanceType"`
}

func (*FP010014I) GetServiceKey() string {
	return "FP010014"
}
