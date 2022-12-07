package models

//Settlement Account Opening

type FP000012I struct {
	PartClearingCode string `json:"partClearingCode" description:"Part clearing code" validate:"required"`
	//PartType         string   `json:"partType"`
	Currency []string `json:"currency" description:"Currency" validate:"required"`
}

type FP000012O struct {
	Records []FP000012OAcctItem `json:"records" description:"Records"`
}

type FP000012OAcctItem struct {
	Currency        string  `json:"currency" description:"Currency"`
	SettAcctNo      string  `json:"settAcctNo" description:"Sett acct no"`
	SettAcctBalance float64 `json:"settAcctBalance" description:"Sett acct balance"`
	AcctBalance     float64 `json:"acctBalance" description:"Acct balance"`
}

func (*FP000012I) GetServiceKey() string {
	return "FP000012"
}
