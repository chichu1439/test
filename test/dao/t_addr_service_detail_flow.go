package dao

//func init() {
//	orm.RegisterModel(new(T_addr_service_detail_flow))
//}

type T_addr_service_detail_flow struct {
	MsgId         string `orm:"column(msg_id);size(35);null" description:"Message Identification"`
	CreDtTm       string `orm:"column(cre_dt_tm);type(date);null" description:"Creation Date Time"`
	AdrReqId      string `orm:"column(adr_req_id);type(date);null" description:"Addressing Request Identification"`
	ProxId        string `orm:"column(prox_id);size(35);null" description:"Proxy Identification"`
	ProxTp        string `orm:"column(prox_tp);size(35);null" description:"Proxy Identification Type"`
	Lang          string `orm:"column(lang);size(35);null" description:"Language"`
	FullNm        string `orm:"column(full_nm);size(140);null" description:"Full Customer Name"`
	DispNm        string `orm:"column(disp_nm);size(140);null" description:"Displayed customer name"`
	CusId         string `orm:"column(cus_id);size(35);null" description:"Customer ID"`
	FrMmbId       string `orm:"column(fr_mmb_id);size(35);null" description:"Sender Participant Mnmber Identification"`
	ToMmbId       string `orm:"column(to_mmb_id);size(35);null" description:"Reveiver Participant Mnmber Identification"`
	MmbId         string `orm:"column(mmb_id);size(35);null" description:"Member Identification"`
	Dflt          string `orm:"column(dflt);size(1);null" description:"Dedault Indicator"`
	PurpCd        string `orm:"column(purp_cd);size(4);null" description:"Purpsoe CODE"`
	CusTp         string `orm:"column(cus_tp);size(4);null" description:"Customer Type"`
	SupopCd       string `orm:"column(supop_cd);size(4);null" description:"Supported Option  Code"`
	TranNo        string `orm:"column(tran_no);size(35);null" description:"Transaction Serial Number"`
	Sts           string `orm:"column(sts);size(1);null" description:"Status"`
	RsnCd         string `orm:"column(rsn_cd);size(4);null" description:"Reason Code"`
	DisplaySts    string `orm:"column(display_sts);size(1);null" description:"Display Status"`
	ReportSts     string `orm:"column(report_sts);size(1);null" description:"Report Status"`
	LastMaintDate string `orm:"column(last_maint_date);type(date);null" description:"最后更新日期"`
	LastMaintTime string `orm:"column(last_maint_time);type(time);null" description:"最后更新时间"`
	LastMaintBrno string `orm:"column(last_maint_brno);size(20);null" description:"最后更新机构"`
	LastMaintTell string `orm:"column(last_maint_tell);size(30);null" description:"最后更新柜员"`
	TccState      int    `orm:"column(tcc_state);size(2);null" description:"Tcc State"`
	ClrSysRef     string `orm:"column(clr_sys_ref);pk;size(35);null" description:"Clearing System Reference"`
}

func (t *T_addr_service_detail_flow) TableName() string {
	return "t_addr_service_detail_flow"
}
