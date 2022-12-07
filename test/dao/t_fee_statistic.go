package dao

import "time"

type TFeeStatistic struct {
	PartId             string    `json:"part_id" xorm:"VARCHAR(20)"`                //参与者
	Currency           string    `json:"currency" xorm:"VARCHAR(3)"`                //币种
	TransType          string    `json:"trans_type" xorm:"VARCHAR(20)"`             //交易类型
	PaymentPurposeCode string    `json:"payment_purpose_code" xorm:"VARCHAR(10)"`   //支付用途分类代码
	WhetherBilling     string    `json:"whether_billing" xorm:"VARCHAR(1)"`         //是否收费
	BillingObject      string    `json:"billing_object" xorm:"VARCHAR(10)"`         //收费对象
	TotAmt             float64   `json:"tot_amt" xorm:"default NULL DECIMAL(19,2)"` //总交易金额
	TotNum             int       `json:"tot_num" xorm:"default NULL INT"`           //总交易笔数
	Charge             float64   `json:"charge" xorm:"default NULL DECIMAL(19,2)"`  //费用金额
	CreateTime         time.Time `json:"create_time" xorm:"created DATETIME"`
	UpdateTime         time.Time `json:"update_time" xorm:"updated DATETIME"`
}

func (*TFeeStatistic) TableName() string {
	return "t_fee_statistic"
}
