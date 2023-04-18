package models

type FP010013I struct {
	MsgId    string `json:"msgId" validate:"required"`
	TransId  string `json:"transId" validate:"required"`
	PartCode string `json:"partCode" validate:"required"`
}

type FP010013O struct {
	TransactionType        string  `json:"transactionType"`
	Status                 string  `json:"status"`
	TransactionId          string  `json:"transactionId"`
	MessageId              string  `json:"messageId"`
	BalanceType            string  `json:"balanceType"` //Increase/Decrease
	TransactionDateTime    string  `json:"transactionDateTime"`
	PayerPartClearCode     string  `json:"payerPartClearCode"`
	PayerPartName          string  `json:"payerPartName"`
	Currency               string  `json:"currency"`
	PaymentCategoryPurpose string  `json:"paymentCategoryPurpose"`
	Amount                 float64 `json:"amount"`
	PayeePartClearCode     string  `json:"payeePartClearCode"`
	PayeePartName          string  `json:"payeePartName"`
}

func (*FP010013I) GetServiceKey() string {
	return "FP010013"
}
