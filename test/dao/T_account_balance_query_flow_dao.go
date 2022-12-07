package dao

import (
	"time"
)

type TAccountBalanceQueryFlow struct {
	MsgId         string    `orm:"column(msg_id);size(35)" description:"Message Identification"`
	CreDtTm       time.Time `orm:"column(cre_dt_tm);type(datetime);null" description:"Creation Date Time"`
	ReqdMsgNmId   string    `orm:"column(reqd_msg_nm_id);size(35);null" description:"Requested Message Name Identification"`
	AcctId        string    `orm:"column(acct_id);size(35);null" description:"Account Identification 账户编号"`
	MmbId         string    `orm:"column(mmb_id);size(35);null" description:"Member Identification 银行行号"`
	Ccy           string    `orm:"column(ccy);size(3);null" description:"Currency"`
	PartBic       string    `orm:"column(part_bic);size(11);null" description:"participant BIC	"`
	FrMmbId       string    `orm:"column(fr_mmb_id);size(35);null" description:"Sender Participant Member Identification"`
	ToMmbId       string    `orm:"column(to_mmb_id);size(35);null" description:"Reveiver Participant Member Identification"`
	HandStatusCd  string    `orm:"column(hand_status_cd);size(2);null" description:"处理状态代码  01-待处理  02-成功  03-失败"`
	LastMaintDate time.Time `orm:"column(last_maint_date);type(date);null" description:"最后更新日期"`
	LastMaintTime time.Time `orm:"column(last_maint_time);type(time);null" description:"最后更新时间"`
	LastMaintBrno string    `orm:"column(last_maint_brno);size(20);null" description:"最后更新机构"`
	LastMaintTell string    `orm:"column(last_maint_tell);size(30);null" description:"最后更新柜员"`
	TccState      int       `orm:"column(tcc_state);null" description:"tcc状态 0-Normal，1-Insert"`
	TranNo        string    `orm:"column(tran_no);pk" description:"交易流水号"`
}

func (t *TAccountBalanceQueryFlow) TableName() string {
	return "t_account_balance_enquiry_request_flow_file"
}
