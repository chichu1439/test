package models

type Transaction struct {
	TransactionId     string  `json:"transactionId"`
	MessageId         string  `json:"messageId"`
	Currency          string  `json:"currency"`
	TransactionAmount float64 `json:"transactionAmount"`
	DateTime          string  `json:"dateTime"`
	RecordType        string  `json:"recordType"`
	NettingId         string  `json:"nettingId"`
	TransactionType   string  `json:"transactionType"`
	BankId            string  `json:"bankId"`
}

type UploadTransactionReq struct {
	Transaction map[string][]Transaction `json:"transaction"`
}

type BankTransaction struct {
	TransactionId     string  `json:"transactionId"`
	MessageId         string  `json:"messageId"`
	Currency          string  `json:"currency"`
	TransactionAmount float64 `json:"transactionAmount"`
	DateTime          string  `json:"dateTime"`
	RecordType        string  `json:"recordType"`
	TransactionType   string  `json:"transactionType"`
	BankId            string  `json:"bankId"`
}

type BankTransactionData struct {
	Transaction []BankTransaction `json:"transaction"`
}

func (*UploadTransactionReq) GetServiceKey() string {
	return "DUS-UploadTransaction"
}
