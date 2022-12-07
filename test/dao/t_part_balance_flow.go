package dao

import "time"

type TPartBalanceHis struct {
	PartClearingCode string    `json:"part_clearing_code" xorm:"not null pk VARCHAR(35)"`
	BalanceDate      string    `json:"balance_date" xorm:"not null DATE"`
	Currency         string    `json:"currency" xorm:"not null VARCHAR(3)"`
	AcctBalance      float64   `json:"acct_balance" xorm:"default NULL DECIMAL(19,2)"`
	PartType         string    `json:"part_type" xorm:"default NULL VARCHAR(2)"`
	CreateTime       time.Time `json:"create_time" xorm:"default NULL created DATETIME"`
	UpdateTime       time.Time `json:"update_time" xorm:"default NULL updated DATETIME"`
}

func (t *TPartBalanceHis) TableName() string {
	return "t_part_balance_his"
}
