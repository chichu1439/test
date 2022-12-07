package models

type FP910027I struct {
	MsgId   string `json:"msgId" validate:"required"`
	TransId string `json:"transId" validate:"required"`
}

type FP910027O struct {
	TransactionType        string  `json:"transactionType"`
	Status                 string  `json:"status"`
	TransactionId          string  `json:"transactionId"`
	MessageId              string  `json:"messageId"`
	TransactionDateTime    string  `json:"transactionDateTime"`
	Currency               string  `json:"currency"`
	PaymentCategoryPurpose string  `json:"paymentCategoryPurpose"`
	Amount                 float64 `json:"amount"`
	PayeePartCode          string  `json:"payeePartCode"`
	PayeePartName          string  `json:"payeePartName"`
	ReferenceNumber        string  `json:"referenceNumber"`
	EndToEndId             string  `json:"endToEndId"`
	InstructionId          string  `json:"instructionId"`
	PayerPartCode          string  `json:"payerPartCode"`
	PayerPartName          string  `json:"payerPartName"`
	BalanceType            string  `json:"balanceType"`
	BalanceMemberId        string  `json:"balanceMemberId"`
}

func (*FP910027I) GetServiceKey() string {
	return "FP910027"
}
