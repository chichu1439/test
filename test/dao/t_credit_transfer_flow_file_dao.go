package dao

type T_credit_transfer_flow_file struct {
	ClrSysRef          string  `orm:"column(clr_sys_ref);pk" description:"Clearing System Reference"`
	MsgId              string  `orm:"column(msg_id);size(35)" description:"Message Identification"`
	FrMmbId            string  `orm:"column(fr_mmb_id);size(35);null" description:"Sender Participant Member Identification"`
	ToMmbId            string  `orm:"column(to_mmb_id);size(35);null" description:"Reveiver Participant Member Identification"`
	CreDtTm            string  `orm:"column(cre_dt_tm);type(datetime);null" description:"Creation Date Time"`
	NbOfTxs            int     `orm:"column(nb_of_txs);null" description:"Number Of Transactions"`
	SttlmMtd           string  `orm:"column(sttlm_mtd);size(2);null" description:"Settlement Method 01-InstructedAgent 02-InstructingAgent 03-ClearingSystem"`
	ClrSysId           string  `orm:"column(clr_sys_id);size(35);null" description:"Clearing System Identification"`
	InstrId            string  `orm:"column(instr_id);size(35);null" description:"Instruction Identification"`
	EndToEndId         string  `orm:"column(end_to_end_id);size(35);null" description:"EndToEnd Identification"`
	TxId               string  `orm:"column(tx_id);size(35);null" description:"Transaction Identification"`
	AcctVerifiOption   string  `orm:"column(acct_verifi_option);size(2);null" description:"Account Verification Option 01-SkipPyeVrf 02-PerformPyeVrf 03-MdSttl"`
	CtgyPurp           string  `orm:"column(ctgy_purp);size(35);null" description:"Category Purpose"`
	IntrBkSttlmAmt     float64 `orm:"column(intr_bk_sttlm_amt);null;digits(18);decimals(5)" description:"Interbank Settlement Amount"`
	IntrBkSttlmAmtCc   string  `orm:"column(intr_bk_sttlm_amt_cc);size(20);null" description:"Interbank Settlement Amount Curreny"`
	IntrBkSttlmDt      string  `orm:"column(intr_bk_sttlm_dt);type(date);null" description:"Interbank Settlement Date"`
	CdtDtTm            string  `orm:"column(cdt_dt_tm);type(datetime);null" description:"Credit Date Time"`
	InstdAmt           float64 `orm:"column(instd_amt);null;digits(18);decimals(5)" description:"Instructed Amount"`
	InstdAmtCcy        string  `orm:"column(instd_amt_ccy);size(20);null" description:"Instructed Amount Curreny"`
	ChrgBrType         string  `orm:"column(chrg_br_type);size(2);null" description:"Charge Bearer 01-BorneByDebtor 02-BorneByCreditor 03-Shared 04-FollowingServiceLevel"`
	ChrgAmt            float64 `orm:"column(chrg_amt);null;digits(18);decimals(5)" description:"Charge Amount"`
	ChrgAmtCcy         string  `orm:"column(chrg_amt_ccy);size(20);null" description:"Charge Amount Curreny"`
	ChrgPartBic        string  `orm:"column(chrg_part_bic);size(11);null" description:"Charge Participant BIC"`
	ChrgMmbId          string  `orm:"column(chrg_mmb_id);size(35);null" description:"Charge Member Identification"`
	DbtrNam            string  `orm:"column(dbtr_nam);size(140);null" description:"Debtor Name"`
	DbtrOrgBic         string  `orm:"column(dbtr_org_bic);size(11);null" description:"Debtor Organization BIC"`
	DbtrId             string  `orm:"column(dbtr_id);size(35);null" description:"Debtor Identification"`
	DbtrSchmeNam       string  `orm:"column(dbtr_schme_nam);size(4);null" description:"Debtor Scheme Name"`
	DbtrMmbNam         string  `orm:"column(dbtr_mmb_nam);size(35);null" description:"Debtor Member Name"`
	DbtrCustId         string  `orm:"column(dbtr_cust_id);size(35);null" description:"Debtor Customer Identification"`
	DbtrCustSchmeNam   string  `orm:"column(dbtr_cust_schme_nam);size(4);null" description:"Debtor Customer Scheme name"`
	DbtrCustMmbNam     string  `orm:"column(dbtr_cust_mmb_nam);size(35);null" description:"Debtor Customer Member Name"`
	DbtrMobileNb       string  `orm:"column(dbtr_mobile_nb);size(35);null" description:"Debtor Mobile Number"`
	DbtrEmailAdr       string  `orm:"column(dbtr_email_adr);size(2048);null" description:"Debtor Email Address"`
	DbtrAcctNo         string  `orm:"column(dbtr_acct_no);size(35);null" description:"Debtor Account No"`
	DbtrAcctType       string  `orm:"column(dbtr_acct_type);size(4);null" description:"Debtor Account Type"`
	DbtrPartBic        string  `orm:"column(dbtr_part_bic);size(11);null" description:"Debtor Participant BIC"`
	DbtrMmbId          string  `orm:"column(dbtr_mmb_id);size(35);null" description:"Debtor Member Identification"`
	CdtrPartBic        string  `orm:"column(cdtr_part_bic);size(11);null" description:"Creditor Participant BIC"`
	CdtrMmbId          string  `orm:"column(cdtr_mmb_id);size(35);null" description:"Creditor Member Identification"`
	CdtrNam            string  `orm:"column(cdtr_nam);size(140);null" description:"Creditor Name"`
	CdtrOrgBic         string  `orm:"column(cdtr_org_bic);size(11);null" description:"Creditor Organization BIC"`
	CdtrId             string  `orm:"column(cdtr_id);size(35);null" description:"Creditor Identification "`
	CdtrSchmeNam       string  `orm:"column(cdtr_schme_nam);size(4);null" description:"Creditor Scheme Name"`
	CdtrMmbNam         string  `orm:"column(cdtr_mmb_nam);size(35);null" description:"Creditor Member Name"`
	CdtrCustId         string  `orm:"column(cdtr_cust_id);size(35);null" description:"Creditor Customer Identification"`
	CdtrCustSchmeNam   string  `orm:"column(cdtr_cust_schme_nam);size(4);null" description:"Creditor Customer Scheme Name"`
	CdtrCustMmbNam     string  `orm:"column(cdtr_cust_mmb_nam);size(35);null" description:"Creditor Customer Member Name"`
	CdtrMobileNb       string  `orm:"column(cdtr_mobile_nb);size(35);null" description:"Creditor Mobile Number"`
	CdtrEmailAdr       string  `orm:"column(cdtr_email_adr);size(2048);null" description:"Creditor Email Address"`
	CdtrAcctNo         string  `orm:"column(cdtr_acct_no);size(35);null" description:"Creditor Account No"`
	CdtrAcctType       string  `orm:"column(cdtr_acct_type);size(4);null" description:"Creditor Account Type"`
	PaymentPurposeCode string  `orm:"column(payment_purpose_code);size(4);null" description:"Payment Purpose Code"`
	PaymentPurposeType string  `orm:"column(payment_purpose_type);size(35);null" description:"Payment Purpose Type"`
	Ustrd              string  `orm:"column(ustrd);size(140);null" description:"Unstructured"`
	TranNo             string  `orm:"column(tran_no);size(40);null" description:"交易流水号"`
	SendType           string  `orm:"column(send_type);size(2);null" description:"发送模式  1-批量 2-实时"`
	HandStatusCd       string  `orm:"column(hand_status_cd);size(2);null" description:"处理状态代码  01-待处理  02-成功  03-失败"`
	HandProCd          string  `orm:"column(hand_pro_cd);size(2);null" description:"处理流程代码  01-申请已发送 02-已发送message reject 03-发出验证 04-收到Account Verification Result 05-已发送payee participant payment Cancellation Request 06-已发送payer participant payment status report(pacs.002) 07-已发送payee participant credit tansfer 08-已发送payee participant payment status report(pacs.002)"`
	SttlmDate          string  `orm:"column(sttlm_date);type(date);null" description:"清算日期"`
	SttlmTime          string  `orm:"column(sttlm_time);type(time);null" description:"清算时间"`
	SttlmStatus        string  `orm:"column(sttlm_status);size(2);null" description:"清算状态 01-成功  02-失败"`
	AcctVerifiType     string  `orm:"column(acct_verifi_type);size(2);null" description:"验证状态 01-成功  02-失败"`
	ReconcFlag         string  `orm:"column(reconc_flag);size(2);null" description:"对账标识"`
	RjctRsn            string  `orm:"column(rjct_rsn);size(35);null" description:"Reject Reason"`
	RjctRsnInf         string  `orm:"column(rjct_rsn_inf);size(105);null" description:"Reject Reason Information"`
	TxStatus           string  `orm:"column(tx_status);size(4);null" description:"Transaction status"`
	TxRjctRsn          string  `orm:"column(tx_rjct_rsn);size(35);null" description:"Transaction Reject Reason"`
	TxRjctRsnInf       string  `orm:"column(tx_rjct_rsn_inf);size(105);null" description:"Transaction Reject Reason Information"`
	AccptncDtTm        string  `orm:"column(accptnc_dt_tm);type(datetime);null" description:"Acceptance Date Time"`
	LastMaintDate      string  `orm:"column(last_maint_date);type(date);null" description:"最后更新日期"`
	LastMaintTime      string  `orm:"column(last_maint_time);type(time);null" description:"最后更新时间"`
	LastMaintBrno      string  `orm:"column(last_maint_brno);size(20);null" description:"最后更新机构"`
	LastMaintTell      string  `orm:"column(last_maint_tell);size(30);null" description:"最后更新柜员"`
	TccState           int     `orm:"column(tcc_state);null" description:"tcc状态 0-Normal，1-Insert"`
}

func (t *T_credit_transfer_flow_file) TableName() string {
	return "t_credit_transfer_flow_file"
}
