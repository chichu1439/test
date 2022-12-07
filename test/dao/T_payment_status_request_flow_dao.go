package dao

import (
	"time"
)

type T_payment_status_request_flow struct {
	MsgId         string    `orm:"column(msg_id);size(35)" description:"Message Identification"`
	CreDtTm       time.Time `orm:"column(cre_dt_tm);type(datetime);null" description:"Creation Date Time"`
	StsReqId      string    `orm:"column(sts_req_id);size(35);null" description:"Status Request Identification"`
	OrgnlMsgId    string    `orm:"column(orgnl_msg_id);size(35);null" description:"Original Message Identification"`
	OrgnlMsgNmId  string    `orm:"column(orgnl_msg_nm_id);size(35);null" description:"Original Message Name Identification"`
	OrgnlTxId     string    `orm:"column(orgnl_tx_id);size(35);null" description:"Original Transaction Identification"`
	ClrSysRef     string    `orm:"column(clr_sys_ref);size(35);null" description:"Clearing System Reference"`
	DbtrPartBic   string    `orm:"column(dbtr_part_bic);size(35);null" description:"Debtor Participant BIC"`
	DbtrMmbId     string    `orm:"column(dbtr_mmb_id);size(35);null" description:"Debtor Member Identification"`
	FrMmbId       string    `orm:"column(fr_mmb_id);size(35);null" description:"Sender Participant Member Identification"`
	ToMmbId       string    `orm:"column(to_mmb_id);size(35);null" description:"Reveiver Participant Member Identification"`
	CdtrPartBic   string    `orm:"column(cdtr_part_bic);size(35);null" description:"Creditor Participant BIC"`
	CdtrMmbId     string    `orm:"column(cdtr_mmb_id);size(35);null" description:"Creditor Member Identification"`
	HandStatusCd  string    `orm:"column(hand_status_cd);size(2);null" description:"处理状态代码  01-待处理  02-成功  03-失败"`
	LastMaintDate time.Time `orm:"column(last_maint_date);type(date);null" description:"最后更新日期"`
	LastMaintTime time.Time `orm:"column(last_maint_time);type(time);null" description:"最后更新时间"`
	LastMaintBrno string    `orm:"column(last_maint_brno);size(20);null" description:"最后更新机构"`
	LastMaintTell string    `orm:"column(last_maint_tell);size(30);null" description:"最后更新柜员"`
	TccState      int       `orm:"column(tcc_state);null" description:"tcc状态 0-Normal，1-Insert"`
	TranNo        string    `orm:"column(tran_no);pk" description:"交易流水号"`
}

func (t *T_payment_status_request_flow) TableName() string {
	return "t_payment_status_request_flow_file"
}
