//Version: v0.0.1
package models

type SAC0000002I struct {
	AList []string `json:"AList" description:"A list" validate:"required"`
}

type SAC0000002O struct {
	AResult []AcctInfo `json:"aResult" description:"A result"`
}

type AcctInfo struct {
	Account          string  `json:"account" description:"Account"`
	AccountingId     string  `json:"accountingId" description:"Accounting id"`
	AmountFrozen     float64 `json:"amountFrozen" description:"Amount frozen"`
	AmountCurrent    float64 `json:"amountCurrent" description:"Amount current"`
	AmountLast       float64 `json:"amountLast" description:"Amount last"`
	OpenDate         string  `json:"openDate" description:"Open date"`
	LastActiveDate   string  `json:"lastActiveDate" description:"Last active date"`
	AccountStatus    string  `json:"accountStatus" description:"Account status"`
	AvalAmount       float64 `json:"avalAmount" description:"Aval amount"`
	LastSettleAmount float64 `json:"lastSettleAmount" description:"Last settle amount"`
	LastSettleDate   string  `json:"lastSettleDate" description:"Last settle date"`
}

func (*SAC0000002I) GetServiceKey() string {
	return "sac0000002"
}
