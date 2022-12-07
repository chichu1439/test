package dao

type T_edda_base_grant_auth struct {
	MndtId           string  `orm:"column(mndt_id);pk"`
	MndtReqId        string  `orm:"column(mndt_req_id)"`
	TypeCode         string  `orm:"column(type_code)"`
	SeqType          string  `orm:"column(seq_type)"`
	FrqcyType        string  `orm:"column(frqcy_type)"`
	FrDt             string  `orm:"column(fr_dt);type(timestamp);null"`
	ToDt             string  `orm:"column(to_dt);type(timestamp);null"`
	ColltnAmt        float64 `orm:"column(colltn_amt)"`
	ColltnAmtCcy     string  `orm:"column(colltn_amt_ccy)"`
	MaxAmt           float64 `orm:"column(max_amt)"`
	MaxAmtCcy        string  `orm:"column(max_amt_ccy)"`
	CdtrNam          string  `orm:"column(cdtr_nam)"`
	CdtrId           string  `orm:"column(cdtr_id)"`
	CdtrSchmeNam     string  `orm:"column(cdtr_schme_nam)"`
	CdtrPrvtId       string  `orm:"column(cdtr_prvt_id)"`
	CdtrPrvtTyp      string  `orm:"column(cdtr_prvt_typ)"`
	CdtrAcctNo       string  `orm:"column(cdtr_acct_no)"`
	CdtrAcctCtry     string  `orm:"column(cdtr_acct_ctry)"`
	CdtrMmbId        string  `orm:"column(cdtr_mmb_id)"`
	DbtrNam          string  `orm:"column(dbtr_nam)"`
	DbtrId           string  `orm:"column(dbtr_id)"`
	DbtrSchmeNam     string  `orm:"column(dbtr_schme_nam)"`
	DbtrPrvtId       string  `orm:"column(dbtr_prvt_id)"`
	DbtrPrvtTyp      string  `orm:"column(dbtr_prvt_typ)"`
	DbtrAcctNo       string  `orm:"column(dbtr_acct_no)"`
	DbtrAcctCtry     string  `orm:"column(dbtr_acct_ctry)"`
	DbtrMmbId        string  `orm:"column(dbtr_mmb_id)"`
	CredtTm          string  `orm:"column(credt_tm);type(timestamp);null"`
	Status           string  `orm:"column(status)"`
	LstUpdTm         string  `orm:"column(lst_upd_tm);type(timestamp);null"`
	CdtrRef          string  `orm:"column(cdtr_ref)"`
	CbtrAcctType     string  `orm:"column(dbtr_acct_type)"`
	CdtrAcctType     string  `orm:"column(cdtr_acct_type)"`
	CdtrProxyId      string  `orm:"column(cdtr_proxy_id)"`
	CdtrProxyType    string  `orm:"column(cdtr_proxy_type)"`
	DbtrProxyId      string  `orm:"column(dbtr_proxy_id)"`
	DbtrProxyType    string  `orm:"column(dbtr_proxy_type)"`
	CdtrCustId       string  `orm:"column(cdtr_cust_id)"`
	DbtrCustId       string  `orm:"column(dbtr_cust_id)"`
	RjctRsn          string  `orm:"column(rjct_rsn)"`
	RjctRsnInf       string  `orm:"column(rjct_rsn_inf)"`
	XactSys          string  `orm:"column(xact_sys)"`
	DbtrAcctNickname string  `orm:"column(dbtr_acct_nickname)"`
	ProcStatus       string  `orm:"column(proc_status)"`
	UlmtDbtrName     string  `orm:"column(ulmt_dbtr_name)"`
	CntPerPro        int     `orm:"column(cnt_per_pro)"`
	SuspEndDt        string  `orm:"column(susp_end_dt);type(date);null"`
	TrcIdc           string  `orm:"column(trc_idc)"`
	CxlRsn           string  `orm:"column(cxl_rsn)"`
	CxlDt            string  `orm:"column(cxl_dt);type(date);null"`
	PermNetId        string  `orm:"column(perm_net_id)"`
	MaintCd          string  `orm:"column(maint_cd)"`
	IpAddr           string  `orm:"column(ip_addr)"`
	BankMndtId       string  `orm:"column(bank_mndt_id)"`
	MobileNb         string  `orm:"column(mobile_nb)"`
	Remark1          string  `orm:"column(remark1)"`
	Remark2          string  `orm:"column(remark2)"`
	Remark3          string  `orm:"column(remark3)"`
	LastMaintDate    string  `orm:"column(last_maint_date);type(date);null"`
	LastMaintTime    string  `orm:"column(last_maint_time);type(datetime);null"`
	LastMaintBrno    string  `orm:"column(last_maint_brno);size(20);null"`
	LastMaintTell    string  `orm:"column(last_maint_tell);size(30);null"`
	TccState         int     `orm:"column(tcc_state)"`
}

func (t *T_edda_base_grant_auth) TableName() string {
	return "t_edda_base_grant_auth"
}
