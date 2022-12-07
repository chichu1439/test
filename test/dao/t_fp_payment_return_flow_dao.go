package dao

import (
	"time"
)

type T_fp_payment_return_flow struct {
	ChannelCode      string    `orm:"column(channel_code);size(4);null" description:"Channel Code"`
	MsgId            string    `orm:"column(msg_id);size(35);null"`
	CreDtTm          time.Time `orm:"column(cre_dt_tm);type(datetime);null" description:"Creation Date Time"`
	NbOfTxs          int       `orm:"column(nb_of_txs);null" description:"Number Of Transactions"`
	SttlmMtd         string    `orm:"column(sttlm_mtd);size(4);null" description:"CLRG-ClearingSystem"`
	ClrSysId         string    `orm:"column(clr_sys_id);size(35);null" description:"FPS"`
	ClrSysRef        string    `orm:"column(clr_sys_ref);pk" description:"Clearing System Reference"`
	FrMmbId          string    `orm:"column(fr_mmb_id);size(35);null" description:"Sender Participant Member Identification"`
	ToMmbId          string    `orm:"column(to_mmb_id);size(35);null" description:"Receiver Participant Member Identification"`
	RtrId            string    `orm:"column(rtr_id);size(35);null" description:"Return Identification"`
	OrgnlInstrId     string    `orm:"column(orgnl_instr_id);size(35);null" description:"Original Instruction Identification"`
	OrgnlEndToEndId  string    `orm:"column(orgnl_end_to_end_id);size(35);null" description:"Original EndToEnd Identification"`
	OrgnlTxId        string    `orm:"column(orgnl_tx_id);size(35);null" description:"Original Transaction Identification"`
	OrgnlClrSysRef   string    `orm:"column(orgnl_clr_sys_ref);size(35);null" description:"Original Clearing System Reference"`
	RtrdIntrBkSttlmA float64   `orm:"column(rtrd_intr_bk_sttlm_a);null;digits(18);decimals(5)" description:"Returned Interbank Settlement Amount"`
	RtrdIntrBkSttlmC string    `orm:"column(rtrd_intr_bk_sttlm_c);size(3);null" description:"Returned Interbank Settlement Amount Ccy"`
	RtrdIntrBkSttlmD time.Time `orm:"column(rtrd_intr_bk_sttlm_d);type(date);null" description:"Returned Interbank Settlement Date"`
	RtrdInstdAmt     float64   `orm:"column(rtrd_instd_amt);null;digits(18);decimals(5)" description:"Returned Instructed Amount"`
	RtrdInstdAmtCcy  string    `orm:"column(rtrd_instd_amt_ccy);size(3);null" description:"Returned Instructed Amount Ccy"`
	ChrgBrType       string    `orm:"column(chrg_br_type);size(4);null" description:"Charge Bearer DEBT-BorneByDebtor CRED-BorneByCreditor SHAR-Shared SLEV-FollowingServiceLevel"`
	ChrgAmt          float64   `orm:"column(chrg_amt);null;digits(18);decimals(5)" description:"Charge Amount"`
	ChrgAmtCcy       string    `orm:"column(chrg_amt_ccy);size(20);null" description:"Charge Amount Curreny"`
	ChrgPartBic      string    `orm:"column(chrg_part_bic);size(11);null" description:"Charge Participant BIC"`
	ChrgMmbId        string    `orm:"column(chrg_mmb_id);size(35);null" description:"Charge Member Identification"`
	RtrCode          string    `orm:"column(rtr_code);size(30);null" description:"Return Code"`
	RtrRsn           string    `orm:"column(rtr_rsn);size(105);null" description:"Return Reason"`
	IntrBkSttlmAmt   float64   `orm:"column(intr_bk_sttlm_amt);null;digits(18);decimals(5)" description:"Interbank Settlement Amount"`
	IntrBkSttlmAmtCc string    `orm:"column(intr_bk_sttlm_amt_cc);size(20);null" description:"Interbank Settlement Amount Curreny"`
	IntrBkSttlmDt    time.Time `orm:"column(intr_bk_sttlm_dt);type(date);null" description:"Interbank Settlement Date"`
	CtgyPurp         string    `orm:"column(ctgy_purp);size(35);null" description:"Category Purpose"`
	MndtId           string    `orm:"column(mndt_id);size(35);null" description:"Mandate Identification"`
	Ustrd            string    `orm:"column(ustrd);size(140);null" description:"Unstructured"`
	DbtrNam          string    `orm:"column(dbtr_nam);size(140);null" description:"Debtor Name"`
	DbtrOrgBic       string    `orm:"column(dbtr_org_bic);size(11);null" description:"Debtor Organization BIC"`
	DbtrId           string    `orm:"column(dbtr_id);size(35);null" description:"Debtor Identification"`
	DbtrSchmeNam     string    `orm:"column(dbtr_schme_nam);size(4);null" description:"Debtor Scheme Name"`
	DbtrMmbNam       string    `orm:"column(dbtr_mmb_nam);size(35);null" description:"Debtor Member Name"`
	DbtrCustId       string    `orm:"column(dbtr_cust_id);size(35);null" description:"Debtor Customer Identification"`
	DbtrCustSchmeNam string    `orm:"column(dbtr_cust_schme_nam);size(4);null" description:"Debtor Customer Scheme name"`
	DbtrCustMmbNam   string    `orm:"column(dbtr_cust_mmb_nam);size(35);null" description:"Debtor Customer Member Name"`
	DbtrMobileNb     string    `orm:"column(dbtr_mobile_nb);size(35);null" description:"Debtor Mobile Number"`
	DbtrEmailAdr     string    `orm:"column(dbtr_email_adr);size(2048);null" description:"Debtor Email Address"`
	DbtrAcctNo       string    `orm:"column(dbtr_acct_no);size(35);null" description:"Debtor Account No"`
	DbtrAcctType     string    `orm:"column(dbtr_acct_type);size(4);null" description:"Debtor Account Type"`
	DbtrPartBic      string    `orm:"column(dbtr_part_bic);size(11);null" description:"Debtor Participant BIC"`
	DbtrMmbId        string    `orm:"column(dbtr_mmb_id);size(35);null" description:"Debtor Member Identification"`
	CdtrPartBic      string    `orm:"column(cdtr_part_bic);size(11);null" description:"Creditor Participant BIC"`
	CdtrMmbId        string    `orm:"column(cdtr_mmb_id);size(35);null" description:"Creditor Member Identification"`
	CdtrNam          string    `orm:"column(cdtr_nam);size(140);null" description:"Creditor Name"`
	CdtrOrgBic       string    `orm:"column(cdtr_org_bic);size(11);null" description:"Creditor Organization BIC"`
	CdtrId           string    `orm:"column(cdtr_id);size(35);null" description:"Creditor Identification "`
	CdtrSchmeNam     string    `orm:"column(cdtr_schme_nam);size(4);null" description:"Creditor Scheme Name"`
	CdtrMmbNam       string    `orm:"column(cdtr_mmb_nam);size(35);null" description:"Creditor Member Name"`
	CdtrCustId       string    `orm:"column(cdtr_cust_id);size(35);null" description:"Creditor Customer Identification"`
	CdtrCustSchmeNam string    `orm:"column(cdtr_cust_schme_nam);size(4);null" description:"Creditor Customer Scheme Name"`
	CdtrCustMmbNam   string    `orm:"column(cdtr_cust_mmb_nam);size(35);null" description:"Creditor Customer Member Name"`
	CdtrMobileNb     string    `orm:"column(cdtr_mobile_nb);size(35);null" description:"Creditor Mobile Number"`
	CdtrEmailAdr     string    `orm:"column(cdtr_email_adr);size(2048);null" description:"Creditor Email Address"`
	CdtrAcctNo       string    `orm:"column(cdtr_acct_no);size(35);null" description:"Creditor Account No"`
	CdtrAcctType     string    `orm:"column(cdtr_acct_type);size(4);null" description:"Creditor Account Type"`
	TranNo           string    `orm:"column(tran_no);size(40);null" description:"交易流水号"`
	SendType         string    `orm:"column(send_type);size(2);null" description:"发送模式  1-批量 2-实时"`
	HandStatusCd     string    `orm:"column(hand_status_cd);size(2);null" description:"处理状态代码  01-待处理  02-成功  03-失败"`
	RtrHandProCd     string    `orm:"column(rtr_hand_pro_cd);size(2);null" description:"01-流水新增成功 02-结算系统结算失败 03-结算系统结算成功  04-发送服务发送报文失败 05-发送服务发送报文成功"`
	SttlmTime        time.Time `orm:"column(sttlm_time);type(time);null" description:"清算时间"`
	SttlmDate        time.Time `orm:"column(sttlm_date);type(date);null" description:"清算日期"`
	SttlmStatus      string    `orm:"column(sttlm_status);size(2);null" description:"清算状态 01-成功  02-失败"`
	RjctRsn          string    `orm:"column(rjct_rsn);size(35);null" description:"Reject Reason"`
	RjctRsnInf       string    `orm:"column(rjct_rsn_inf);size(105);null" description:"Reject Reason Information"`
	TxStatus         string    `orm:"column(tx_status);size(5);null" description:"Transaction status"`
	TxRjctRsn        string    `orm:"column(tx_rjct_rsn);size(35);null" description:"Transaction Reject Reason"`
	TxRjctRsnInf     string    `orm:"column(tx_rjct_rsn_inf);size(105);null" description:"Transaction Reject Reason Information"`
	LastMaintDate    time.Time `orm:"column(last_maint_date);type(date);null" description:"最后更新日期"`
	LastMaintTime    time.Time `orm:"column(last_maint_time);type(time);null" description:"最后更新时间"`
	LastMaintBrno    string    `orm:"column(last_maint_brno);size(20);null" description:"最后更新机构"`
	LastMaintTell    string    `orm:"column(last_maint_tell);size(30);null" description:"最后更新柜员"`
	TccState         int       `orm:"column(tcc_state);null" description:"tcc状态 0-Normal，1-Insert"`
}

func (t *T_fp_payment_return_flow) TableName() string {
	return "t_return_refund_flow_file"
}
