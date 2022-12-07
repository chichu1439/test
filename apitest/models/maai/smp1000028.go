package models

type SMP1000028I struct {
	CifNo       string  `json:"cifNo" validate:"required,max=20"`
	Amount      float64 `json:"amount" validate:"required,max=20"`
	Remark      string  `json:"remark" validate:"required,max=20"`
	ClientTrnNo string  `json:"clientTrnNo" validate:"required,max=20"`
	ReqRefNo    string  `json:"reqRefNo" validate:"omitempty,max=20"`
}

type TransferWallet struct {
	AcctNo             string  `json:"acctNo"`
	WalletID           string  `json:"walletId"`
	CurrencyCode       string  `json:"currencyCode"`
	AcctStatusCode     string  `json:"acctStatusCode"`
	OutstandingBalance float64 `json:"outstandingBalance"`
	AvailableBalance   float64 `json:"availableBalance"`
	AcctType           string  `json:"acctType"`
	Prior              int     `json:"prior"`
	ApprovalLimit      float64 `json:"approvalLimit"`
	MinBalance         float64 `json:"minBalance"`
	MaxBalance         float64 `json:"maxBalance"`
	CorpCode           string  `json:"corpCode"`
	Sof                string  `json:"sof"`
}

type SMP1000028O struct {
	RespCode      string         `json:"respCode"`
	RespDesc      string         `json:"respDesc"`
	ReqRefNo      string         `json:"reqRefNo"`
	RespRefNo     string         `json:"respRefNo"`
	TxStatusCode  string         `json:"txStatusCode"`
	TxStatusName  string         `json:"txStatusName"`
	SenderTrnNo   string         `json:"senderTrnNo"`
	ReceiverTrnNo string         `json:"receiverTrnNo"`
	Wallet        TransferWallet `json:"wallet"`
}

func (i *SMP1000028I) GetServiceKey() string {
	return "MP100028"
}
