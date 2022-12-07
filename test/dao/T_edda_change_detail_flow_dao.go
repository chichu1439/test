package dao

import (
	"time"
)

//func init() {
//	orm.RegisterModel(
//		new(T_edda_change_detail_flow),
//	)
//}

type T_edda_change_detail_flow struct {
	ClrSysRef        string    `orm:"column(clr_sys_ref);pk;size(35);null" description:"Clearing System Reference"`
	FrMmbId          string    `orm:"column(fr_mmb_id);size(35);null" description:"Sender Participant Mnmber Identification"`
	ToMmbId          string    `orm:"column(to_mmb_id);size(35);null" description:"Reveiver Participant Mnmber Identification"`
	TranNo           string    `orm:"column(tran_no);size(40);null" description:"Transaction serial number"`
	TranFlowNo       string    `orm:"column(tran_flow_no);size(40);null" description:"Transaction serial flow number"`
	MndtId           string    `orm:"column(mndt_id);size(35);null" description:"Mandate Identification"`
	MndtReqId        string    `orm:"column(mndt_req_id);size(35);null" description:"Mandate Request Identification"`
	TypeCode         string    `orm:"column(type_code);size(20);null" description:"Type Code  "`
	SeqType          string    `orm:"column(seq_type);size(20);null" description:"Sequence Type"`
	FrqcyType        string    `orm:"column(frqcy_type);size(20);null" description:"Frequency Type"`
	FrDt             time.Time `orm:"column(fr_dt);type(time);null" description:"From Date "`
	ToDt             time.Time `orm:"column(to_dt);type(time);null" description:"To Date"`
	ColltnAmt        float64   `orm:"column(colltn_amt);null;digits(18);decimals(2)" description:"Collection Amount  "`
	ColltnAmtCcy     string    `orm:"column(colltn_amt_ccy);size(20);null" description:"ColltnAmtCcy"`
	MaxAmt           float64   `orm:"column(max_amt);null;digits(18);decimals(2)" description:"Maximum Amount"`
	MaxAmtCcy        string    `orm:"column(max_amt_ccy);size(20);null" description:"MaxAmtCcy"`
	CdtrNam          string    `orm:"column(cdtr_nam);size(140);null" description:"Creditor Name"`
	CdtrId           string    `orm:"column(cdtr_id);size(35);null" description:"Creditor ID"`
	CdtrSchmeNam     string    `orm:"column(cdtr_schme_nam);size(20);null" description:"Creditor Scheme Name"`
	CdtrPrvtId       string    `orm:"column(cdtr_prvt_id);size(35);null" description:"Creditor Private Identification"`
	CdtrPrvtTyp      string    `orm:"column(cdtr_prvt_typ);size(20);null" description:"Creditor Private Type"`
	CdtrAcctNo       string    `orm:"column(cdtr_acct_no);size(35);null" description:"Creditor Account"`
	CdtrAcctCtry     string    `orm:"column(cdtr_acct_ctry);size(20);null" description:"Creditor  Account Country"`
	CdtrMmbId        string    `orm:"column(cdtr_mmb_id);size(20);null" description:"Creditor Member Identification"`
	DbtrNam          string    `orm:"column(dbtr_nam);size(140);null" description:"Debtor Name"`
	DbtrId           string    `orm:"column(dbtr_id);size(35);null" description:"Debtor ID"`
	DbtrSchmeNam     string    `orm:"column(dbtr_schme_nam);size(20);null" description:"Debtor Scheme Name"`
	DbtrPrvtId       string    `orm:"column(dbtr_prvt_id);size(35);null" description:"Debtor Private Identification"`
	DbtrPrvtTyp      string    `orm:"column(dbtr_prvt_typ);size(20);null" description:"Debtor Private Type"`
	DbtrAcctNo       string    `orm:"column(dbtr_acct_no);size(35);null" description:"Debtor Account"`
	DbtrAcctCtry     string    `orm:"column(dbtr_acct_ctry);size(20);null" description:"Debtor Account Country"`
	DbtrMmbId        string    `orm:"column(dbtr_mmb_id);size(20);null" description:"Debtor  Member Identification"`
	CredtTm          time.Time `orm:"column(credt_tm);size(2);null" description:"Create TIME"`
	Status           string    `orm:"column(status)";type(timestamp);null" description:"Status"`
	LstUpdTm         time.Time `orm:"column(lst_upd_tm);type(timestamp);null" description:"Last update time"`
	CdtrRef          string    `orm:"column(cdtr_ref);size(35);null" description:"Creditor Reference"`
	DbtrAcctType     string    `orm:"column(dbtr_acct_type);size(4);null" description:"Debtor Account Type"`
	CdtrAcctType     string    `orm:"column(cdtr_acct_type);size(4);null" description:"Creditor  Account Type"`
	CdtrProxyId      string    `orm:"column(cdtr_proxy_id);size(34);null" description:"Creditor Proxy  Identification"`
	CdtrProxyType    string    `orm:"column(cdtr_proxy_type);size(4);null" description:"Creditor Proxy  Type"`
	DbtrProxyId      string    `orm:"column(dbtr_proxy_id);size(34);null" description:"Debtor  Proxy Identification"`
	DbtrProxyType    string    `orm:"column(dbtr_proxy_type);size(4);null" description:"Debtor  Proxy Type"`
	CdtrCustId       string    `orm:"column(cdtr_cust_id);size(35);null" description:"Creditor Customer Identification"`
	DbtrCustId       string    `orm:"column(dbtr_cust_id);size(35);null" description:"Debtor Customer Identification"`
	RjctRsn          string    `orm:"column(rjct_rsn);size(10);null" description:"Reject Reason"`
	RjctRsnInf       string    `orm:"column(rjct_rsn_inf);size(150);null" description:"Reject Reason Information"`
	XactSys          string    `orm:"column(xact_sys);size(20);null" description:"External Aash Clearing System"`
	DbtrAcctNickname string    `orm:"column(dbtr_acct_nickname);size(35);null" description:"Debtor Account Nickname"`
	ProcStatus       string    `orm:"column(proc_status);size(20);null" description:"Product Status"`
	UlmtDbtrName     string    `orm:"column(ulmt_dbtr_name);size(200);null" description:"Ultimate Debtor"`
	CntPerPro        int       `orm:"column(cnt_per_pro);null" description:"Count Per Period"`
	SuspEndDt        time.Time `orm:"column(susp_end_dt);type(date);null" description:"suspension end date"`
	TrcIdc           string    `orm:"column(trc_idc);size(20);null" description:"Tracking Indicator"`
	CxlRsn           string    `orm:"column(cxl_rsn);size(100);null" description:"Cancel Reason"`
	CxlDt            time.Time `orm:"column(cxl_dt);type(date);null" description:"Cancel Date"`
	PermNetId        string    `orm:"column(perm_net_id);size(30);null" description:"Perman Net Identification"`
	MaintCd          string    `orm:"column(maint_cd);size(10);null" description:"Maintenance Code"`
	IpAddr           string    `orm:"column(ip_addr);size(20);null" description:"IP address"`
	BankMndtId       string    `orm:"column(bank_mndt_id);size(64);null" description:"Bank Mandate Identification"`
	MobileNb         string    `orm:"column(mobile_nb);size(35);null" description:"Mobile Number"`
	CreDtTm          time.Time `orm:"column(cre_dt_tm);type(datetime);null" description:"Creation Date Time"`
	SendType         string    `orm:"column(send_type);size(2);null" description:"发送模式  1-批量 2-实时"`
	ChgeType         string    `orm:"column(chge_type);size(3);null" description:"01-启动 02-修改 03-暂停 04-取消"`
	MsgType          string    `orm:"column(msg_type);size(35);null" description:"报文类型"`
	MsgId            string    `orm:"column(msg_id);size(35);null" description:"报文标识号"`
	HandStatusCd     string    `orm:"column(hand_status_cd);size(2);null" description:"处理状态代码 01-待处理  02-成功  03-失败"`
	HandProCd        string    `orm:"column(hand_pro_cd);size(2);null" description:"报文处理状态 01-申请已接收 02-发送获取报告 03-发出验证 04-收到验证报告 05-已向发出方发出验证报告 06-已发验证报告到接收方"`
	SspnsnReqId      string    `orm:"column(sspnsn_req_id);size(35);null" description:"暂停请求ID"`
	SsonsRsnCd       string    `orm:"column(ssons_rsn_cd);size(35);null" description:"暂停原因代码"`
	SsonsRsnInf      string    `orm:"column(ssons_rsn_inf);size(105);null" description:"暂停详细原因"`
	AmdmntRsnCd      string    `orm:"column(amdmnt_rsn_cd);size(4);null" description:"修改原因代码"`
	AmdmntRsnInf     string    `orm:"column(amdmnt_rsn_inf);size(105);null" description:"修改详细原因"`
	CxlRsnInf        string    `orm:"column(cxl_rsn_inf);size(105);null" description:"取消详细原因"`
	Remark1          string    `orm:"column(remark1);size(320);null" description:"备注1"`
	Remark2          string    `orm:"column(remark2);size(320);null" description:"备注2"`
	Remark3          string    `orm:"column(remark3);size(320);null" description:"备注3"`
	LastMaintDate    time.Time `orm:"column(last_maint_date);type(date);null" description:"最后更新日期"`
	LastMaintTime    time.Time `orm:"column(last_maint_time);type(time);null" description:"最后更新时间"`
	LastMaintBrno    string    `orm:"column(last_maint_brno);size(20);null" description:"最后更新机构"`
	LastMaintTell    string    `orm:"column(last_maint_tell);size(30);null" description:"最后更新柜员"`
	TccState         int       `orm:"column(tcc_state)"`
}

func (t *T_edda_change_detail_flow) TableName() string {
	return "t_edda_change_detail_flow"
}
