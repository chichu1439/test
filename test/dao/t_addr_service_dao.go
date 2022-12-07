package dao

type T_addr_service struct {
	MsgId         string `orm:"column(msg_id);pk;size(50);null" description:"Message Identification"`
	CreDtTm       string `orm:"column(cre_dt_tm);type(date);null" description:"Creation Date Time"`
	AdrReqId      string `orm:"column(adr_req_id);size(50);null" description:"Addressing Request Identification"`
	ProxId        string `orm:"column(prox_id);size(35);null" description:"Proxy Identification"`
	ProxTp        string `orm:"column(prox_tp);size(35);null" description:"Proxy Identification Type"`
	Lang          string `orm:"column(lang);size(35);null" description:"Language"`
	FullNm        string `orm:"column(full_nm);size(140);null" description:"Full Customer Name"`
	DispNm        string `orm:"column(disp_nm);size(140);null" description:"Displayed customer name"`
	CusId         string `orm:"column(cus_id);size(35)" description:"Customer ID"`
	AcctId        string `orm:"column(acct_id);size(35)" description:"Account ID"`
	Bank          string `orm:"column(bank);size(35)" description:"Bind Bank"`
	MmbId         string `orm:"column(mmb_id);size(35)" description:"Member Identification"`
	Dflt          string `orm:"column(dflt);size(1);null" description:"Dedault Indicator"`
	PurpCd        string `orm:"column(purp_cd);size(4);null" description:"Purpsoe CODE"`
	CusTp         string `orm:"column(cus_tp);size(4);null" description:"Customer Type"`
	SupopCd       string `orm:"column(supop_cd);size(4);null" description:"Supported Option  Code"`
	ValidSts      string `orm:"column(valid_sts);size(1);null" description:"valid status"`
	UbCustFlag    string `orm:"column(ub_cust_flag);size(1);null" description:"Cust Force Unbind Flag"`
	UbProxFlag    string `orm:"column(ub_prox_flag);size(1);null" description:"Prox Force Unbind Flag"`
	DfltPayOrder  int    `orm:"column(dflt_pay_order);null" description:"default pay order"`
	LastMaintDate string `orm:"column(last_maint_date);type(date);null" description:"最后更新日期"`
	LastMaintTime string `orm:"column(last_maint_time);type(time);null" description:"最后更新时间"`
	LastMaintBrno string `orm:"column(last_maint_brno);size(20);null" description:"最后更新机构"`
	LastMaintTell string `orm:"column(last_maint_tell);size(30);null" description:"最后更新柜员"`
	TccState      int    `orm:"column(tcc_state);size(2);null" description:"Tcc State"`
}

func (t *T_addr_service) TableName() string {
	return "t_addr_service"
}
