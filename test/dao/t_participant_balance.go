package dao

type TPartBalance struct {
	PartClearingCode string  `json:"part_clearing_code"`
	Currency         string  `json:"currency"`
	PartType         string  `json:"part_type"`
	SettAcctNo       string  `json:"sett_acct_no"`
	AcctBalance      float64 `json:"acct_balance"`
	DpBalance        float64 `json:"dp_balance"`
	IpBalance        float64 `json:"ip_balance"`
	CreateTime       string  `json:"create_time"`
	UpdateTime       string  `json:"update_time"`
}

func (t *TPartBalance) TableName() string {
	return "t_part_balance"
}
