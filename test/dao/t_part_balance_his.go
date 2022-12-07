package dao

import "time"

type TPartBalanceFlow struct {
	Id               int       `json:"id" xorm:"not null pk int"`
	FlowId           string    `json:"flow_id" xorm:"default null VARCHAR(60)"`
	PartClearingCode string    `json:"part_clearing_code" xorm:"default null VARCHAR(35)"`
	Currency         string    `json:"currency" xorm:"default null VARCHAR(3)"`
	TransDate        string    `json:"trans_date" xorm:"default null DATE"`
	DebCrtFlag       string    `json:"deb_crt_flag" xorm:"default null VARCHAR(1)"`
	TransAmount      float64   `json:"trans_balance" xorm:"default NULL DECIMAL(19,2)"`
	AcctBalance      float64   `json:"acct_balance" xorm:"default NULL DECIMAL(19,2)"`
	CreateTime       time.Time `json:"create_time" xorm:"default NULL created DATETIME"`
	UpdateTime       time.Time `json:"update_time" xorm:"default NULL updated DATETIME"`
}

func (t *TPartBalanceFlow) TableName() string {
	return "t_part_balance_flow"
}
