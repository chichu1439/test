package dao

type TFpPartBalanceControl struct {
	PartId          string  `orm:"column(part_id);pk"`
	Ccy             string  `orm:"column(ccy)"`
	PartType        string  `orm:"column(part_type);null"`
	StmiAccountNo   string  `orm:"column(stmi_account_no);null"`
	UpperPartId     string  `orm:"column(upper_part_id);null"`
	DrCtrl          string  `orm:"column(dr_ctrl);null"`
	CrCtrl          string  `orm:"column(cr_ctrl);null"`
	Status          string  `orm:"column(status);null"`
	OverdraftLimit  float64 `orm:"column(overdraft_limit);null;digits(18);decimals(5)"`
	CurrentBal      float64 `orm:"column(current_bal);null;digits(18);decimals(5)"`
	CirDepAmt       float64 `orm:"column(cir_dep_amt);null;digits(18);decimals(5)"`
	ContrlMinBal    float64 `orm:"column(contrl_min_bal);null;digits(18);decimals(5)"`
	BalAlertNtcFlag string  `orm:"column(bal_alert_ntc_flag);null"`
	LastMaintDate   string  `orm:"column(last_maint_date);type(date);null"`
	LastMaintTime   string  `orm:"column(last_maint_time);type(time);null"`
	LastMaintBrno   string  `orm:"column(last_maint_brno);size(20);null"`
	LastMaintTell   string  `orm:"column(last_maint_tell);size(30);null"`
	TccState        int     `orm:"column(tcc_state);null" description:"tcc status 0-Normal，1-Insert"`
}

func (t *TFpPartBalanceControl) TableName() string {
	return "t_fp_part_balance_control"
}
