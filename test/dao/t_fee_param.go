package dao

import "time"

type TFeeParam struct {
	EffectiveDate      string    `json:"effective_date" xorm:"DATE"`                     //生效日期
	TransType          string    `json:"trans_type" xorm:"VARCHAR(20)"`                  //交易类型
	Currency           string    `json:"currency" xorm:"VARCHAR(3)"`                     //收费币种
	PaymentPurposeCode string    `json:"payment_purpose_code" xorm:"VARCHAR(10)"`        //支付用途分类代码
	BillingObject      string    `json:"billing_object" xorm:"default NULL VARCHAR(10)"` //收费对象 Payer Payee
	Mode               string    `json:"mode" xorm:"default NULL VARCHAR(1)"`            //收费模式O-固定金额 F-按费率
	Unit               string    `json:"unit" xorm:"default NULL VARCHAR(3)"`            //收费单位
	FeeAmt             float64   `json:"fee_amt" xorm:"default NULL DECIMAL(19,2)"`      //收费金额
	FeeRate            float64   `json:"fee_rate" xorm:"default NULL DECIMAL(19,2)"`     //费率
	MinFeeAmt          float64   `json:"min_fee_amt" xorm:"default NULL DECIMAL(19,2)"`  //最低收费金额
	MaxFeeAmt          float64   `json:"max_fee_amt" xorm:"default NULL DECIMAL(19,2)"`  //最高收费金额
	Status             string    `json:"status" xorm:"default NULL VARCHAR(1)"`          //状态
	CreateTime         time.Time `json:"create_time" xorm:"created DATETIME"`
	UpdateTime         time.Time `json:"update_time" xorm:"updated DATETIME"`
}

func (*TFeeParam) TableName() string {
	return "t_fee_param"
}
