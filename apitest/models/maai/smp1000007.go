package models

type SMP1000007I struct {
	ReqRefNo string `json:"reqRefNo"`
}
type Wallets struct {
	AcctNo             string `json:"acctNo"`
	CurrencyCode       string `json:"currencyCode"`
	AcctStatusCode     string `json:"acctStatusCode"`
	OutstandingBalance string `json:"outstandingBalance"`
	AvailableBalance   string `json:"availableBalance"`
	AcctType           string `json:"acctType"`
	Prior              string `json:"prior"`
	ApprovalLimit      string `json:"approvalLimit"`
	MinBalance         string `json:"minBalance"`
	MaxBalance         string `json:"maxBalance"`
	CorpCode           string `json:"corpCode"`
	Sof                []Sof
}
type Sof struct {
	SofId           string `json:"sofId"`
	SofSourceCode   string `json:"sofSourceCode"`
	SofAcctNo       string `json:"sofAcctNo"`
	BotBankCode     string `json:"botBankCode"`
	ActivatedFlag   string `json:"activatedFlag"`
	DefaultFlag     string `json:"defaultFlag"`
	EmailMatched    string `json:"emailMatched"`
	MobileNoMatched string `json:"mobileNoMatched"`
	UserIdMatched   string `json:"userIdMatched"`
}
type SMP1000007O struct {
	Wallets   []Wallets
	RespCode  string `json:"respCode"`
	RespDesc  string `json:"respDesc"`
	ReqRefNo  string `json:"reqRefNo"`
	RespRefNo string `json:"respRefNo"`
}
type SMP1000007IOutgress struct {
	SMP1000007I
}

func (i *SMP1000007I) GetServiceKey() string {
	return "MP100007"
}

func (i *SMP1000007IOutgress) GetServiceKey() string {
	return "MP300007"
}
