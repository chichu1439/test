package dao

type TFPFeeCalcResult struct {
	PartId          string  `orm:"column(part_id);pk;" json:"PartId,omitempty"`
	TransactionType string  `orm:"column(transaction_type);" json:"TransactionType,omitempty"`
	Currency        string  `orm:"column(currency);" json:"Currency,omitempty"`
	TransDate       string  `orm:"column(trans_date);" json:"TransDate,omitempty"`
	ChargeCount     int     `orm:"column(charge_count);" json:"ChargeCount,omitempty"`
	ChargeAmount    float64 `orm:"column(charge_amount);" json:"ChargeAmount,omitempty"`
	CalcTime        string  `orm:"column(calc_time);" json:"CalcTime,omitempty"`
}

//
//  获取表名
//  Get table name.
//
func (this *TFPFeeCalcResult) TableName() string {
	return "t_fp_fee_calc_result"
} //end
