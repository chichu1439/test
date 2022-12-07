package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"math/rand"
	"net/http"
	"time"
)

//var Db map[string]config.Db
var O *xorm.Engine
var serveraddr string
var addr string
var user string
var pwd string
var database string

var maxIdleConns int
var maxOpenConns int
var defaultQueryLimit int
var maxLimitValue int
var maxLifeValue int

func main() {

	flag.StringVar(&serveraddr, "serveraddr", "127.0.0.1:8081", "server Url")
	flag.StringVar(&addr, "address", "127.0.0.1:3306", " Url")
	flag.StringVar(&user, "user", "root", "user")
	flag.StringVar(&pwd, "password", "root", "Password")
	flag.StringVar(&database, "database", "payment", "database")
	flag.IntVar(&maxIdleConns, "maxIdleConns", 50, "maxIdleConns")
	flag.IntVar(&maxOpenConns, "maxOpenConns", 50, "maxOpenConns")
	flag.IntVar(&defaultQueryLimit, "defaultQueryLimit", 30, "defaultQueryLimit")
	flag.IntVar(&maxLimitValue, "maxLimitValue", 50000, "maxLimitValue")
	flag.IntVar(&maxLifeValue, "maxLifeValue", 540, "maxLifeValue")
	flag.Parse()
	//Db = map[string]config.Db{}
	//Db["db"] = config.Db{
	//	Name:     "",
	//	Type:     "mysql",
	//	Su:       "su",
	//	Topics:   nil,
	//	Default:  false,
	//	Addr:     addr,
	//	User:     user,
	//	Password: pwd,
	//	Database: database,
	//	Params:   "?charset=utf8&loc=Local",
	//	Debug:    false,
	//	Pool: struct {
	//		MaxIdleConns      int
	//		MaxOpenConns      int
	//		DefaultQueryLimit int
	//		MaxLimitValue     int
	//		MaxLifeValue      int
	//	}{MaxIdleConns: maxIdleConns,
	//		MaxOpenConns:      maxOpenConns,
	//		DefaultQueryLimit: defaultQueryLimit,
	//		MaxLimitValue:     maxLimitValue,
	//		MaxLifeValue:      maxLifeValue,
	//	},
	//}
	//if err := db.InitDBManagerForXorm(Db); err != nil {
	//	fmt.Println(err)
	//}
	//if nil == O {
	//	o, err := db.GetXormEngine("su", "")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	O = o
	//}
	dsn := user + ":" + pwd + "@tcp(" + addr + ")/" + database

	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}
	engine.SetMaxIdleConns(maxIdleConns)
	engine.SetMaxOpenConns(maxOpenConns)
	engine.SetConnMaxLifetime(time.Duration(maxLifeValue) * time.Second)
	engine.ShowSQL(false)
	O = engine
	http.HandleFunc("/", testSql)
	server := &http.Server{
		Addr: serveraddr,
	}
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}

}

func testSql(w http.ResponseWriter, req *http.Request) {
	var err error
	defer func() {
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
		}
	}()
	now := time.Now()
	dateTime := now.Format("20060102150405")
	randomStr := GetRandomString(8)
	MsgTransDetail := TMsgTransDetail{
		ClrSysRef:             "FRN20211110YVE36PM0Y045340LRKTGV08Q",
		ClrSysId:              "",
		MsgId:                 "11" + dateTime,
		FromMemberId:          "006",
		ToMemberId:            "005",
		CreateDatetime:        now,
		NumOfTrans:            "",
		StlmtMethod:           "",
		InstructionId:         "",
		EndToEndId:            "12" + dateTime,
		TransId:               "12" + dateTime + randomStr,
		AcctVerifyOption:      "",
		CategoryPurpose:       "",
		InterbankStlmtAmt:     0.01,
		InterbankStlmtCcy:     "THB",
		InterbankStlmtDate:    now,
		CreditDatetime:        now,
		InstructedAmt:         0.00,
		InstructedCcy:         "",
		ChargeBearerType:      "",
		ChargeAmt:             0,
		ChargeCcy:             "",
		ChargePartBic:         "",
		ChargeMemberId:        "",
		MandateId:             "",
		DrName:                "",
		DrOrgBic:              "",
		DrId:                  "",
		DrSchemeName:          "",
		DrMemberName:          "",
		DrCustId:              "",
		DrCustSchemeName:      "",
		DrCustMemberName:      "",
		DrMobileNum:           "",
		DrEmailAddr:           "",
		DrAcctNo:              "",
		DrAcctType:            "",
		DrPartBic:             "",
		DrMemberId:            "",
		CrPartBic:             "",
		CrMemberId:            "",
		CrName:                "",
		CrOrgBic:              "",
		CrId:                  "",
		CrSchemeName:          "",
		CrMemberName:          "",
		CrCustId:              "",
		CrCustSchemeName:      "",
		CrCustMemberName:      "",
		CrMobileNum:           "",
		CrEmailAddr:           "",
		CrAcctNo:              "",
		CrAcctType:            "",
		PaymentPurposeCode:    "",
		PaymentPurposeType:    "",
		Unstructured:          "",
		TransNo:               "",
		SendType:              "",
		StatusCode:            "",
		ProcessCode:           "",
		StlmtDate:             now,
		StlmtTime:             now,
		StlmtStatus:           "",
		AcctVerifyType:        "",
		ReconcileFlag:         "",
		RejectReason:          "",
		RejectReasonInfo:      "",
		TransType:             "",
		TransDate:             now,
		TransStatus:           "",
		TransRejectReason:     "",
		TransRejectReasonInfo: "",
		AcceptDatetime:        now,
		RecordType:            "",
		CreateTime:            now,
		UpdateTime:            now,
		TccState:              0,
		PrimaryAcctNum:        "",
		SysTraceAuditNum:      "",
		MerchantType:          "",
		PosEntryCode:          "",
		CardSeqNum:            "",
		ForwardingId:          "",
		RetrievalRefNum:       "",
		CardAcceptorName:      "",
		SenderFee:             0,
		FromAcctFee:           0,
		ToAcctFee:             0,
		ProxyType:             "",
		ProxyId:               "",
		PinBlockType:          "",
		TransRef:              "",
		TransTaxType:          "",
		SenderProxyId:         "",
		SenderProxyType:       "",
		SenderRefNo:           "",
		SenderTaxId:           "",
		SendBranchCode:        "",
		ReceiverTaxId:         "",
		ReceiverBranchCode:    "",
		BankRefNum:            "",
		VatRates:              0,
		VatAmt:                0,
		IncomeType:            "",
		WithholdingTaxRates:   0,
		WithholdingTaxAmt:     0,
		WithholdingTaxCond:    "",
		SenderType:            "",
		ReceiverType:          "",
		BillRef1:              "",
		BillRef2:              "",
		BillRef3:              "",
		BillDueDate:           now,
		EmvData:               "",
		PosInfo:               "",
		NetworkCode:           ""}
	_, err = O.Insert(&MsgTransDetail)
	if err != nil {
		return
	}
	transId := "22" + dateTime + GetRandomString(8)
	MsgTransDetail.Id = 0
	MsgTransDetail.TransId = transId
	_, err = O.Insert(&MsgTransDetail)
	if err != nil {
		return
	}
	detail := TMsgTransDetail{
		TransId: transId,
	}
	_, err = O.Get(&detail)
	if err != nil {
		return
	}
	_, err = O.Update(TMsgTransDetail{StatusCode: "01"}, TMsgTransDetail{TransId: transId})
	if err != nil {
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("success"))
}

type TMsgTransDetail struct {
	Id                    int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	ClrSysRef             string    `json:"clr_sys_ref" xorm:"default 'NULL' comment('Clearing System Reference') VARCHAR(35)"`
	ClrSysId              string    `json:"clr_sys_id" xorm:"default 'NULL' comment('Clearing System Identification') VARCHAR(35)"`
	MsgId                 string    `json:"msg_id" xorm:"not null default '''' comment('Message Identification') index VARCHAR(60)"`
	FromMemberId          string    `json:"from_member_id" xorm:"default 'NULL' comment('Sender Participant Member Identification') VARCHAR(35)"`
	ToMemberId            string    `json:"to_member_id" xorm:"default 'NULL' comment('Receiver Participant Member Identification') VARCHAR(35)"`
	CreateDatetime        time.Time `json:"create_datetime" xorm:"default 'NULL' comment('Creation Date Time') DATETIME"`
	NumOfTrans            string    `json:"num_of_trans" xorm:"default 'NULL' comment('Number Of Transactions') VARCHAR(255)"`
	StlmtMethod           string    `json:"stlmt_method" xorm:"default 'NULL' comment('Settlement Method 01-InstructedAgent 02-InstructingAgent 03-ClearingSystem') VARCHAR(4)"`
	InstructionId         string    `json:"instruction_id" xorm:"default 'NULL' comment('Instruction Identification') VARCHAR(35)"`
	EndToEndId            string    `json:"end_to_end_id" xorm:"default 'NULL' comment('EndToEnd Identification') VARCHAR(35)"`
	TransId               string    `json:"trans_id" xorm:"not null comment('Transaction Identification') index VARCHAR(35)"`
	AcctVerifyOption      string    `json:"acct_verify_option" xorm:"default 'NULL' comment('Account Verification Option 01-SkipPyeVrf 02-PerformPyeVrf 03-MdSttl') VARCHAR(35)"`
	CategoryPurpose       string    `json:"category_purpose" xorm:"default 'NULL' comment('Category Purpose') VARCHAR(35)"`
	InterbankStlmtAmt     float64   `json:"interbank_stlmt_amt" xorm:"default NULL comment('Interbank Settlement Amount') DECIMAL(18,5)"`
	InterbankStlmtCcy     string    `json:"interbank_stlmt_ccy" xorm:"default 'NULL' comment('Interbank Settlement Amount Currency') VARCHAR(20)"`
	InterbankStlmtDate    time.Time `json:"interbank_stlmt_date" xorm:"default 'NULL' comment('Interbank Settlement Date') DATE"`
	CreditDatetime        time.Time `json:"credit_datetime" xorm:"default 'NULL' comment('Credit Date Time') DATETIME"`
	InstructedAmt         float64   `json:"instructed_amt" xorm:"default NULL comment('Instructed Amount') DECIMAL(18,5)"`
	InstructedCcy         string    `json:"instructed_ccy" xorm:"default 'NULL' comment('Instructed Amount Currency') VARCHAR(20)"`
	ChargeBearerType      string    `json:"charge_bearer_type" xorm:"default 'NULL' comment('Charge Bearer 01-BorneByDebtor 02-BorneByCreditor 03-Shared 04-FollowingServiceLevel') VARCHAR(4)"`
	ChargeAmt             float64   `json:"charge_amt" xorm:"default NULL comment('Charge Amount') DECIMAL(18,5)"`
	ChargeCcy             string    `json:"charge_ccy" xorm:"default 'NULL' comment('Charge Amount Currency') VARCHAR(20)"`
	ChargePartBic         string    `json:"charge_part_bic" xorm:"default 'NULL' comment('Charge Participant BIC') VARCHAR(11)"`
	ChargeMemberId        string    `json:"charge_member_id" xorm:"default 'NULL' comment('Charge Member Identification') VARCHAR(35)"`
	MandateId             string    `json:"mandate_id" xorm:"default 'NULL' comment('Mandate Identification') VARCHAR(35)"`
	DrName                string    `json:"dr_name" xorm:"default 'NULL' comment('Debtor Name') VARCHAR(140)"`
	DrOrgBic              string    `json:"dr_org_bic" xorm:"default 'NULL' comment('Debtor Organization BIC') VARCHAR(11)"`
	DrId                  string    `json:"dr_id" xorm:"default 'NULL' comment('Debtor Identification') VARCHAR(35)"`
	DrSchemeName          string    `json:"dr_scheme_name" xorm:"default 'NULL' comment('Debtor Scheme Name') VARCHAR(4)"`
	DrMemberName          string    `json:"dr_member_name" xorm:"default 'NULL' comment('Debtor Member Name') VARCHAR(35)"`
	DrCustId              string    `json:"dr_cust_id" xorm:"default 'NULL' comment('Debtor Customer Identification') VARCHAR(35)"`
	DrCustSchemeName      string    `json:"dr_cust_scheme_name" xorm:"default 'NULL' comment('Debtor Customer Scheme name') VARCHAR(4)"`
	DrCustMemberName      string    `json:"dr_cust_member_name" xorm:"default 'NULL' comment('Debtor Customer Member Name') VARCHAR(35)"`
	DrMobileNum           string    `json:"dr_mobile_num" xorm:"default 'NULL' comment('Debtor Mobile Number') VARCHAR(35)"`
	DrEmailAddr           string    `json:"dr_email_addr" xorm:"default 'NULL' comment('Debtor Email Address') VARCHAR(2048)"`
	DrAcctNo              string    `json:"dr_acct_no" xorm:"default 'NULL' comment('Debtor Account No') VARCHAR(35)"`
	DrAcctType            string    `json:"dr_acct_type" xorm:"default 'NULL' comment('Debtor Account Type') VARCHAR(4)"`
	DrPartBic             string    `json:"dr_part_bic" xorm:"default 'NULL' comment('Debtor Participant BIC') VARCHAR(11)"`
	DrMemberId            string    `json:"dr_member_id" xorm:"default 'NULL' comment('Debtor Member Identification') VARCHAR(35)"`
	CrPartBic             string    `json:"cr_part_bic" xorm:"default 'NULL' comment('Creditor Participant BIC') VARCHAR(11)"`
	CrMemberId            string    `json:"cr_member_id" xorm:"default 'NULL' comment('Creditor Member Identification') VARCHAR(35)"`
	CrName                string    `json:"cr_name" xorm:"default 'NULL' comment('Creditor Name') VARCHAR(140)"`
	CrOrgBic              string    `json:"cr_org_bic" xorm:"default 'NULL' comment('Creditor Organization BIC') VARCHAR(11)"`
	CrId                  string    `json:"cr_id" xorm:"default 'NULL' comment('Creditor Identification') VARCHAR(35)"`
	CrSchemeName          string    `json:"cr_scheme_name" xorm:"default 'NULL' comment('Creditor Scheme Name') VARCHAR(4)"`
	CrMemberName          string    `json:"cr_member_name" xorm:"default 'NULL' comment('Creditor Member Name') VARCHAR(35)"`
	CrCustId              string    `json:"cr_cust_id" xorm:"default 'NULL' comment('Creditor Customer Identification') VARCHAR(35)"`
	CrCustSchemeName      string    `json:"cr_cust_scheme_name" xorm:"default 'NULL' comment('Creditor Customer Scheme Name') VARCHAR(4)"`
	CrCustMemberName      string    `json:"cr_cust_member_name" xorm:"default 'NULL' comment('Creditor Customer Member Name') VARCHAR(35)"`
	CrMobileNum           string    `json:"cr_mobile_num" xorm:"default 'NULL' comment('Creditor Mobile Number') VARCHAR(35)"`
	CrEmailAddr           string    `json:"cr_email_addr" xorm:"default 'NULL' comment('Creditor Email Address') VARCHAR(2048)"`
	CrAcctNo              string    `json:"cr_acct_no" xorm:"default 'NULL' comment('Creditor Account No') VARCHAR(35)"`
	CrAcctType            string    `json:"cr_acct_type" xorm:"default 'NULL' comment('Creditor Account Type') VARCHAR(4)"`
	PaymentPurposeCode    string    `json:"payment_purpose_code" xorm:"default 'NULL' comment('Payment Purpose Code') VARCHAR(10)"`
	PaymentPurposeType    string    `json:"payment_purpose_type" xorm:"default 'NULL' comment('Payment Purpose Type') VARCHAR(35)"`
	Unstructured          string    `json:"unstructured" xorm:"default 'NULL' comment('Unstructured') VARCHAR(140)"`
	TransNo               string    `json:"trans_no" xorm:"default 'NULL' comment('交易流水号') VARCHAR(64)"`
	SendType              string    `json:"send_type" xorm:"default 'NULL' comment('发送模式  1-批量 2-实时') VARCHAR(2)"`
	StatusCode            string    `json:"status_code" xorm:"default 'NULL' comment('处理状态代码  01-待处理  02-成功  03-失败') VARCHAR(2)"`
	ProcessCode           string    `json:"process_code" xorm:"default 'NULL' comment('处理流程代码  01-申请已发送 02-已发送message reject 03-发出验证 04-收到Account Verification Result 05-已发送payee participant payment Cancellation Request 06-已发送payer participant payment status report(pacs.002) 07-已发送payee participant credit tansfer 08-已发送payee participant payment status report(pacs.002)') VARCHAR(2)"`
	StlmtDate             time.Time `json:"stlmt_date" xorm:"default 'NULL' comment('清算日期') DATE"`
	StlmtTime             time.Time `json:"stlmt_time" xorm:"default 'NULL' comment('清算时间') TIME"`
	StlmtStatus           string    `json:"stlmt_status" xorm:"default 'NULL' comment('清算状态 01-成功  02-失败') VARCHAR(2)"`
	AcctVerifyType        string    `json:"acct_verify_type" xorm:"default 'NULL' comment('验证状态 01-成功  02-失败') VARCHAR(2)"`
	ReconcileFlag         string    `json:"reconcile_flag" xorm:"default 'NULL' comment('对账标识') VARCHAR(2)"`
	RejectReason          string    `json:"reject_reason" xorm:"default 'NULL' comment('Reject Reason') VARCHAR(35)"`
	RejectReasonInfo      string    `json:"reject_reason_info" xorm:"default 'NULL' comment('Reject Reason Information') VARCHAR(105)"`
	TransType             string    `json:"trans_type" xorm:"not null comment('01-Credit Transfer; 02-Direct Debit; 03-Return / Refund; 04-Balance Sweeping Notification; 05-IP Balance Adjustment Notification; 06-Payment Status Enquiry; 07-Account Balance Enquiry;08-request to pay;') VARCHAR(2)"`
	TransDate             time.Time `json:"trans_date" xorm:"not null comment('Transaction date') DATE"`
	TransStatus           string    `json:"trans_status" xorm:"default 'NULL' comment('Transaction status') VARCHAR(4)"`
	TransRejectReason     string    `json:"trans_reject_reason" xorm:"default 'NULL' comment('Transaction Reject Reason') VARCHAR(35)"`
	TransRejectReasonInfo string    `json:"trans_reject_reason_info" xorm:"default 'NULL' comment('Transaction Reject Reason Information') VARCHAR(105)"`
	AcceptDatetime        time.Time `json:"accept_datetime" xorm:"default 'NULL' comment('Acceptance Date Time') DATETIME"`
	RecordType            string    `json:"record_type" xorm:"not null default '''' comment('1-双方在不同su上的第二笔记录') VARCHAR(1)"`
	CreateTime            time.Time `json:"create_time" xorm:"default 'NULL' comment('最后更新日期') DATETIME"`
	UpdateTime            time.Time `json:"update_time" xorm:"default 'NULL' comment('最后更新时间') DATETIME"`
	TccState              int       `json:"tcc_state" xorm:"default NULL comment('tcc状态 0-Normal，1-Insert') INT(11)"`
	PrimaryAcctNum        string    `json:"primary_acct_num" xorm:"default 'NULL' VARCHAR(20)"`
	SysTraceAuditNum      string    `json:"sys_trace_audit_num" xorm:"default 'NULL' VARCHAR(10)"`
	MerchantType          string    `json:"merchant_type" xorm:"default 'NULL' VARCHAR(10)"`
	PosEntryCode          string    `json:"pos_entry_code" xorm:"default 'NULL' VARCHAR(10)"`
	CardSeqNum            string    `json:"card_seq_num" xorm:"default 'NULL' VARCHAR(10)"`
	ForwardingId          string    `json:"forwarding_id" xorm:"default 'NULL' VARCHAR(20)"`
	RetrievalRefNum       string    `json:"retrieval_ref_num" xorm:"default 'NULL' VARCHAR(20)"`
	CardAcceptorName      string    `json:"card_acceptor_name" xorm:"default 'NULL' VARCHAR(40)"`
	SenderFee             float64   `json:"sender_fee" xorm:"default NULL DECIMAL(18,5)"`
	FromAcctFee           float64   `json:"from_acct_fee" xorm:"default NULL DECIMAL(18,5)"`
	ToAcctFee             float64   `json:"to_acct_fee" xorm:"default NULL DECIMAL(18,5)"`
	ProxyType             string    `json:"proxy_type" xorm:"default 'NULL' VARCHAR(20)"`
	ProxyId               string    `json:"proxy_id" xorm:"default 'NULL' VARCHAR(128)"`
	PinBlockType          string    `json:"pin_block_type" xorm:"default 'NULL' VARCHAR(1)"`
	TransRef              string    `json:"trans_ref" xorm:"default 'NULL' VARCHAR(20)"`
	TransTaxType          string    `json:"trans_tax_type" xorm:"default 'NULL' VARCHAR(1)"`
	SenderProxyId         string    `json:"sender_proxy_id" xorm:"default 'NULL' VARCHAR(128)"`
	SenderProxyType       string    `json:"sender_proxy_type" xorm:"default 'NULL' VARCHAR(20)"`
	SenderRefNo           string    `json:"sender_ref_no" xorm:"default 'NULL' VARCHAR(20)"`
	SenderTaxId           string    `json:"sender_tax_id" xorm:"default 'NULL' VARCHAR(13)"`
	SendBranchCode        string    `json:"send_branch_code" xorm:"default 'NULL' VARCHAR(6)"`
	ReceiverTaxId         string    `json:"receiver_tax_id" xorm:"default 'NULL' VARCHAR(13)"`
	ReceiverBranchCode    string    `json:"receiver_branch_code" xorm:"default 'NULL' VARCHAR(6)"`
	BankRefNum            string    `json:"bank_ref_num" xorm:"default 'NULL' VARCHAR(50)"`
	VatRates              float64   `json:"vat_rates" xorm:"default NULL DECIMAL(18,5)"`
	VatAmt                float64   `json:"vat_amt" xorm:"default NULL DECIMAL(18,5)"`
	IncomeType            string    `json:"income_type" xorm:"default 'NULL' VARCHAR(3)"`
	WithholdingTaxRates   float64   `json:"withholding_tax_rates" xorm:"default NULL DECIMAL(18,5)"`
	WithholdingTaxAmt     float64   `json:"withholding_tax_amt" xorm:"default NULL DECIMAL(18,5)"`
	WithholdingTaxCond    string    `json:"withholding_tax_cond" xorm:"default 'NULL' VARCHAR(1)"`
	SenderType            string    `json:"sender_type" xorm:"default 'NULL' VARCHAR(1)"`
	ReceiverType          string    `json:"receiver_type" xorm:"default 'NULL' VARCHAR(1)"`
	BillRef1              string    `json:"bill_ref1" xorm:"default 'NULL' VARCHAR(20)"`
	BillRef2              string    `json:"bill_ref2" xorm:"default 'NULL' VARCHAR(20)"`
	BillRef3              string    `json:"bill_ref3" xorm:"default 'NULL' VARCHAR(20)"`
	BillDueDate           time.Time `json:"bill_due_date" xorm:"default 'NULL' DATE"`
	EmvData               string    `json:"emv_data" xorm:"default 'NULL' VARCHAR(256)"`
	PosInfo               string    `json:"pos_info" xorm:"default 'NULL' VARCHAR(12)"`
	NetworkCode           string    `json:"network_code" xorm:"default 'NULL' VARCHAR(3)"`
}

func (t *TMsgTransDetail) TableName() string {
	return "t_msg_trans_detail"
}
func GetRandomString(l int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
