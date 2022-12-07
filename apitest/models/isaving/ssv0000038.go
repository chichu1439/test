//Version: v1.0.0.0
package models

type SV000038I struct {
	ExternalSystemFlow      string                  `json:"externalSystemFlow" description:"External system flow"`                        //外围业务系统流水号
	ExternalSystem          string                  `json:"externalSystem" description:"External system"`                                 //外围业务系统编号
	TradeFlag               string                  `json:"tradeFlag" description:"Trade flag"`                                           //反交易标志
	ReverseTradeFlag        string                  `json:"reverseTradeFlag" description:"Reverse trade flag"`                            //冲正交易标志
	OriginalTransactionDate string                  `json:"originalTransactionDate" description:"Original transaction date"`              //原交易日期
	OriginalGlobalBizseqNo  string                  `json:"originalGlobalBizseqNo" description:"Original global bizseq no"`               //原全局流水号
	OriginalSrcBizseqNo     string                  `json:"originalSrcBizseqNo" description:"Original src bizseq no"`                     //原业务流水号
	Postscript              string                  `json:"postscript" description:"Postscript" validate:"max=200"`                       //转账附言
	DlyArriveAFlag          string                  `json:"dlyArriveAFlag" description:"Dly arrive a flag" validate:"max=1"`              //延时到账标志 1-实时到账；2-延时到账
	TransType               string                  `json:"transType" description:"Trans type" validate:"required,max=1"`                 //交易方式 1-联机；2-批量
	BatchNumber             string                  `json:"batchNumber" description:"Batch number" validate:"max=10"`                     //批号 当批次调用时，需要送入次栏位
	AmountFreezeAuthority   string                  `json:"amountFreezeAuthority" description:"Amount freeze authority" validate:"max=1"` //金额冻结授权标志 Y-已经授权
	KeyVersion              string                  `json:"keyVersion" description:"Key version"`
	TransferArray           []SV000038ITransferInfo `json:"transferArray" description:"Transfer array"`
}

type SV000038ITransferInfo struct {
	ProcessType          string  `json:"processType" description:"Process type"`
	SourceFunds          string  `json:"sourceFunds" description:"Source funds" validate:"required,max=1"`           //资金来源 1-客户账户；2-内部账户；3-合约
	PaymentAgreement     string  `json:"paymentAgreement" description:"Payment agreement" validate:"max=30"`         //付款合约号
	PaymentAgreementType string  `json:"paymentAgreementType" description:"Payment agreement type" validate:"max=5"` //付款合约类型 10001-个人活期账户;10002-个人定期账户;10003-个人活期保证金账户;10004-个人定期保证金账户;10005-对公活期账户;10006-对公定期账户;10007-对公活期保证金账户;10008-对公定期保证金账户;10009-内部账户;20001-个人贷款账户;20002-对公贷款账户;30001-协议存款;30002-保证金;30003-同业存款;30004-智能通知存款;30005-法人账户透支;30006-集团账户;30007-协定存款;30008-收息账号;30009-活期存款转定期约定;30010-网银开户签约;30011-贷款;30012-贷款收费;30013-普通定期存款;30014-活期保证金存款;30015-定期保证金存款;30016-普通活期存款;40001-客户对账单;-
	PaymentInAccount     string  `json:"paymentInAccount" description:"Payment in account" validate:"max=30"`        //转出内部户账号
	PaymentMediumNM      string  `json:"paymentMediumNm" description:"Payment medium n m" validate:"max=30"`         //付款方介质号码
	PaymentMediumType    string  `json:"paymentMediumType" description:"Payment medium type" validate:"max=3"`       //付款方介质类型
	PaymentCustomerId    string  `json:"paymentCustomerId" description:"Payment customer id" validate:"max=20"`      //付款方客户编号
	PaymentCustomerType  string  `json:"paymentCustomerType" description:"Payment customer type" validate:"max=2"`   //付款方客户类型 1-个人客户;2-法人客户;3-联名客户;4-集团客户;5-客户关系个人信息;6-客户关系对公信息;7-内部户客户;8-集团成员客户-
	PaymentCUR           string  `json:"paymentCur" description:"Payment c u r" validate:"required,max=3"`           //付款方币别 156-人民币
	CashTransFlag        string  `json:"cashTransFlag" description:"Cash trans flag" validate:"max=2"`               //付款方钞汇标志 C-钞户;T-汇户;-
	PaymentAccountName   string  `json:"paymentAccountName" description:"Payment account name" validate:"max=200"`   //付款账户名称
	UsageCode            string  `json:"usageCode" description:"Usage code" validate:"max=3"`                        //用途代码 EAC-电子账号
	PayAbstractCode      string  `json:"payAbstractCode" description:"Pay abstract code"`
	WithdrawMethod       string  `json:"withdrawMethod" description:"Withdraw method" validate:"max=1"`               //支取方式 1-密码;
	NeedPasswordFlag     string  `json:"needPasswordFlag" description:"Need password flag" validate:"required,max=1"` //是否需要验密 Y-需要验密；N-不需要验密
	AccountPassword      string  `json:"accountPassword" description:"Account password" validate:""`                  //账户密码 当是否需要验密为Y，本字段必填
	TransAmount          float64 `json:"transAmount" description:"Trans amount" validate:"required"`                  //金额
	GoingFunds           string  `json:"goingFunds" description:"Going funds" validate:"required,max=1"`              //资金去向 1-客户账户；2-内部账户
	CltAgreement         string  `json:"cltAgreement" description:"Clt agreement" validate:"max=30"`                  //收款合约号
	CltAgreementType     string  `json:"cltAgreementType" description:"Clt agreement type" validate:"max=5"`          //收款合约类型
	CltInAccount         string  `json:"cltInAccount" description:"Clt in account" validate:"max=30"`                 //转入内部户账号
	CltMediumNM          string  `json:"cltMediumNm" description:"Clt medium n m" validate:"max=30"`                  //收款方介质号码
	CltMediumType        string  `json:"cltMediumType" description:"Clt medium type" validate:"max=3"`                //收款方介质类型
	CltCustomerId        string  `json:"cltCustomerId" description:"Clt customer id" validate:"max=20"`               //收款方客户编号
	CltCustomerType      string  `json:"cltCustomerType" description:"Clt customer type" validate:"max=2"`            //收款方客户类型
	CltCUR               string  `json:"cltCur" description:"Clt c u r" validate:"required,max=3"`                    //收款方币别
	CltCashTransFlag     string  `json:"cltCashTransFlag" description:"Clt cash trans flag" validate:"max=2"`         //收款方钞汇标志
	CltUsageCode         string  `json:"cltUsageCode" description:"Clt usage code" validate:"max=3"`                  //收款方用途代码
	CltAccountName       string  `json:"cltAccountName" description:"Clt account name" validate:"max=200"`            //收款账户名称
	CtlAbstractCode      string  `json:"ctlAbstractCode" description:"Ctl abstract code"`
}

type SV000038O struct {
	ResultArray []SV000038OResultInfo `json:"resultArray" description:"Result array"`
}

type SV000038OResultInfo struct {
	PaymentAgreement     string  `json:"paymentAgreement" description:"Payment agreement" validate:"max=30"`         //付款合约号
	PaymentAgreementType string  `json:"paymentAgreementType" description:"Payment agreement type" validate:"max=5"` //付款合约类型 10001-个人活期账户;10002-个人定期账户;10003-个人活期保证金账户;10004-个人定期保证金账户;10005-对公活期账户;10006-对公定期账户;10007-对公活期保证金账户;10008-对公定期保证金账户;10009-内部账户;20001-个人贷款账户;20002-对公贷款账户;30001-协议存款;30002-保证金;30003-同业存款;30004-智能通知存款;30005-法人账户透支;30006-集团账户;30007-协定存款;30008-收息账号;30009-活期存款转定期约定;30010-网银开户签约;30011-贷款;30012-贷款收费;30013-普通定期存款;30014-活期保证金存款;30015-定期保证金存款;30016-普通活期存款;40001-客户对账单;-
	PaymentInAccount     string  `json:"paymentInAccount" description:"Payment in account" validate:"max=30"`        //转出内部户账号
	PaymentMediumNM      string  `json:"paymentMediumNm" description:"Payment medium n m" validate:"max=30"`         //付款方介质号码
	PaymentMediumType    string  `json:"paymentMediumType" description:"Payment medium type" validate:"max=3"`       //付款方介质类型
	TueTransAmount       float64 `json:"tueTransAmount" description:"Tue trans amount"`                              //实际扣款金额
}

func (*SV000038I) GetServiceKey() string {
	return "ssv0000038"
}
