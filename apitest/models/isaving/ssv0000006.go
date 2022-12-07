//Version: v1.0.0
package models

type SSV0000006I struct {
	FreezeMediumNumber      string  `json:"freezeMediumNumber" validate:"max=30" description:"Media number"`                                   //介质号码
	FreezeMediumType        string  `json:"freezeMediumType" validate:"max=3" description:"Media type"`                                        //介质类型
	FreezeAgreement         string  `json:"freezeAgreement" validate:"required,max=30" description:"Freeze contract number"`                   //冻结合约号
	FreezeAgreementType     string  `json:"freezeAgreementType" validate:"required,max=5" description:"product code"`                          //冻结合约类型 10001-个人活期账户;10002-个人定期账户;10003-个人活期保证金账户;10004-个人定期保证金账户;10005-对公活期账户;10006-对公定期账户;10007-对公活期保证金账户;10008-对公定期保证金账户;10009-内部账户;20001-个人贷款账户;20002-对公贷款账户;30001-协议存款;30002-保证金;30003-同业存款;30004-智能通知存款;30005-法人账户透支;30006-集团账户;30007-协定存款;30008-收息账号;30009-活期存款转定期约定;30010-网银开户签约;30011-贷款;30012-贷款收费;30013-普通定期存款;30014-活期保证金存款;30015-定期保证金存款;30016-普通活期存款;40001-客户对账单;-
	FreezeAccountName       string  `json:"freezeAccountName" validate:"max=200" description:"Freeze contract type"`                           //冻结账户名称
	FreezeMode              string  `json:"freezeMode" validate:"required,max=1,oneof=1 2" description:"Freezing method"`                      //冻结方式 1-账户冻结，2-金额冻结
	FreezeCurrency          string  `json:"freezeCurrency" validate:"max=3" description:"Currency"`                                            //币别 冻结方式为2-金额冻结时必输
	CashTransFlag           string  `json:"cashTransFlag" validate:"omitempty,max=2,oneof=C T" description:"Money exchange logo"`              //钞汇标志 冻结方式为2-金额冻结时必输
	FreezeOrgnizeType       string  `json:"freezeOrgnizeType" validate:"required,max=1,oneof=1 2" description:"Type of freezing organization"` //冻结机构类型 1-有权机关，2-银行内部
	FreezeOrgnizeName       string  `json:"freezeOrgnizeName" validate:"required" description:"Name of freezing organization"`                 //冻结机构名称 如内部预授权冻结上送申请模块名称（如支付系统发起冻结
	FreezeDocumentNumber    string  `json:"freezeDocumentNumber" description:"Frozen instrument number"`                                       //冻结文书号 冻结结构类型为1-有权机关时必输
	FreezeApplicationNumber string  `json:"freezeApplicationNumber" description:"Application freeze number"`                                   //申请冻结编号 冻结结构类型为2-银行内部时必输
	FreezeAmount            float64 `json:"freezeAmount" validate:"omitempty,gte=0.0" description:"Frozen amount"`                             //冻结金额 冻结方式为2-金额冻结时必输
	FreezeReason            string  `json:"freezeReason" description:"Frozen amount"`                                                          //冻结原因
	EffectiveDate           string  `json:"effectiveDate" validate:"required,nefield=FreezeOweDate" description:"effective date"`              //生效日期
	FreezeOweDate           string  `json:"freezeOweDate" validate:"required" description:"Freeze expiry date"`                                //冻结到期日
	EnforcerName1           string  `json:"enforcerName1" validate:"max=60" description:"Name of law enforcement officer one"`                 //执法人员名称1
	EnforcerId1             string  `json:"enforcerId1" validate:"max=60" description:"Law enforcement officer ID number one"`                 //执法人员证件号码1
	EnforcerName2           string  `json:"enforcerName2" validate:"max=60" description:"Name of law enforcement officer two"`                 //执法人员名称2
	EnforcerId2             string  `json:"enforcerId2" validate:"max=60" description:"Law enforcement officer ID number two"`                 //执法人员证件号码2
	FreezeOrgnize           string  `json:"freezeOrgnize" description:"Freeze agency"`                                                         //冻结机构
	FreezeEmail             string  `json:"freezeEmail" description:"Freeze teller"`                                                           //冻结柜员
	FreezeAuthorityEmail    string  `json:"freezeAuthorityEmail" description:"Freeze authorized tellers"`                                      //冻结授权柜员
	Remarks                 string  `json:"remarks" validate:"max=80" description:"Remarks"`                                                   //备注
	CustomerId              string  `json:"customerId" description:"Frozen customer number"`                                                   //冻结客户编号
	CustomerType            string  `json:"customerType" description:"Frozen customer type"`                                                   //冻结客户类型
	BankFreezeType          string  `json:"bankFreezeType" validate:"max=1" description:"Bank freeze distinction"`                             //银行冻结区分 M-Internal management;T-tansation
}

type SSV0000006O struct {
	FreezeBusinessNumber string `json:"freezeBusinessNumber" validate:"required,max=30" description:"Frozen business number"` //冻结业务编号
	UnfreezeFwNumber     int    `json:"unfreezeFwNumber" validate:"len=6" description:"Unfreeze priority"`                    //解冻优先次序
}

func (*SSV0000006I) GetServiceKey() string {
	return "ssv0000006"
}
