package dao

type TFpTransDet struct {
	TransId          string  `orm:"column(trans_id);pk"`
	CreditPart       string  `orm:"column(credit_part);null"`
	CreditPartType   string  `orm:"column(credit_part_type);null"`
	CreditCcy        string  `orm:"column(credit_ccy);null"`
	DebitPart        string  `orm:"column(debit_part);null"`
	DebitPartType    string  `orm:"column(debit_part_type);null"`
	DebitCcy         string  `orm:"column(debit_ccy);null"`
	TransAmt         float64 `orm:"column(trans_amt);null;digits(18);decimals(5)"`
	TransDate        string  `orm:"column(trans_date);type(date);null"`
	TransTime        string  `orm:"column(trans_time);type(time);null"`
	TransStatus      string  `orm:"column(trans_status);null"`
	TransType        string  `orm:"column(trans_type);null" description:"transaction type CRTXFR DDEBIT RTNRFD BLNSWP IPBADJ CBDRCR"`
	RtgsReturnResult string  `orm:"column(rtgs_return_result);null" description:"rtgs return result 01-success  02-failed 03-overtime"`
	RtgsFailReason   string  `orm:"column(rtgs_fail_reason);null"`
	DrNoticeMsg      string  `orm:"column(dr_notice_msg);null"`
	CrNoticeMsg      string  `orm:"column(cr_notice_msg);null"`
	FrMmbId          string  `orm:"column(fr_mmb_id);null"`
	ToMmbId          string  `orm:"column(to_mmb_id);null"`
	LastMaintDate    string  `orm:"column(last_maint_date);type(date);null"`
	LastMaintTime    string  `orm:"column(last_maint_time);type(time);null"`
	LastMaintBrno    string  `orm:"column(last_maint_brno);size(20);null"`
	LastMaintTell    string  `orm:"column(last_maint_tell);size(30);null"`
	TccState         int     `orm:"column(tcc_state);null" description:"tcc status 0-Normal，1-Insert"`
}

func (t *TFpTransDet) TableName() string {
	return "t_fp_trans_det"
}
