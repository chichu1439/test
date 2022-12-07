package dao

type PartFlow struct {
	PartId        string  `orm:"column(part_id);pk"`
	Ccy           string  `orm:"column(ccy);null"`
	StmiAccountNo string  `orm:"column(stmi_account_no);null"`
	TransId       string  `orm:"column(trans_id);null"`
	TransDate     string  `orm:"column(trans_date);null"`
	TransTime     string  `orm:"column(trans_time);null"`
	CrDrSign      string  `orm:"column(cr_dr_sign);null"`
	TransType     string  `orm:"column(trans_type);null"`
	TransAmt      float64 `orm:"column(trans_amt);null"`
	AcctBalance   float64 `orm:"column(acct_balance);null"`
	FrMmbId       string  `orm:"column(fr_mmb_id);null"`
	ToMmbId       string  `orm:"column(to_mmb_id);null"`
	LastMaintDate string  `orm:"column(last_maint_date);null"`
	LastMaintTime string  `orm:"column(last_maint_time);null"`
	LastMaintBrno string  `orm:"column(last_maint_brno);null"`
	LastMaintTell string  `orm:"column(last_maint_tell);null"`
	TccState      int     `orm:"column(tcc_state);null" description:"tcc status 0-Normal，1-Insert"`
}

func (t *PartFlow) TableName() string {
	return "t_fp_part_flow"
}
