package dao

type RtgsSettlement struct {
	TxnId    string  `orm:"column(txn_id);pk"`
	TxnType  string  `orm:"column(txn_type);null"` //CRTXFR DDEBIT RTNRFD BLNSWP IPBADJ CBDRCR
	DrAcctNo string  `orm:"column(dr_acct_no);null"`
	CrAcctNo string  `orm:"column(cr_acct_no);null"`
	Ccy      string  `orm:"column(ccy);null"`
	TxnAmt   float64 `orm:"column(txn_amt);null"`
}

func (t *RtgsSettlement) TableName() string {
	return "t_fp_rtgs_settlement"
}
