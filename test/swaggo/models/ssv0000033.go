package models

type SV000033I struct {
	AgreementId          string `json:"agreementId" description:"Agreement Id" validate:"required,max=30"`
	AgreementType        string `json:"agreementType" description:"Agreement type" validate:"required,max=5"`
	FreezeBusinessNumber string `json:"freezeBusinessNumber" validate:"required"`
}

type SV000033O struct {
	FreezeBusinessNumber  string  `json:"freezeBusinessNumber" xorm:"not null pk comment('冻结业务编号 19位:冻结方式（1位）+冻结发起方（1位）+DCN后四位（4位）+19位纳秒时间戳') VARCHAR(30)"`
	FreezeType            string  `json:"freezeType" xorm:"comment('冻结状态 0-解冻;1-合约冻结;2-金额冻结-3暂禁') VARCHAR(1)"`
	EffectiveDate         string  `json:"effectiveDate" xorm:"not null comment('冻结生效日') DATE"`
	FreezeOweDate         string  `json:"freezeOweDate" xorm:"comment('冻结到期日') DATE"`
	FreezeDate            string  `json:"freezeDate" xorm:"not null comment('冻结日期') DATE"`
	FreezeTime            string  `json:"freezeTime" xorm:"comment('冻结时间') TIME"`
	FreezeOrganize        string  `json:"freezeOrganize" xorm:"comment('冻结行所') VARCHAR(30)"`
	FreezeTeller          string  `json:"freezeTeller" xorm:"comment('冻结柜员') VARCHAR(30)"`
	FreezeAuthorityTeller string  `json:"freezeAuthorityTeller" xorm:"comment('冻结授权柜员') VARCHAR(30)"`
	FreezeOrganizeType    string  `json:"freezeOrganizeType" xorm:"not null comment('冻结发起方 1-权利机构冻结;2-银行冻结') VARCHAR(1)"`
	FreezeOrganizeName    string  `json:"freezeOrganizeName" xorm:"comment('冻结发起方名称') VARCHAR(60)"`
	FreezeMode            string  `json:"freezeMode" xorm:"comment('冻结方式 CM0002 1-账户冻结;2-金额冻结;3-暂禁;N-无冻结;P-保付支票冻结;S-止付票据;C-圈存;Z-全部冻结+金额冻结（只用于查询接口）;-') VARCHAR(1)"`
	CustomerId            string  `json:"customerId" xorm:"comment('冻结客户编号') VARCHAR(20)"`
	CustomerType          string  `json:"customerType" xorm:"comment('冻结客户类型 01-个人客户;02-法人客户;03-联名客户') VARCHAR(2)"`
	MediumType            string  `json:"mediumType" xorm:"comment('介质类型') VARCHAR(3)"`
	MediumNumber          string  `json:"mediumNumber" xorm:"comment('介质号码') VARCHAR(60)"`
	FreezeAgreement       string  `json:"freezeAgreement" xorm:"not null comment('冻结合约号') VARCHAR(30)"`
	FreezeAgreementType   string  `json:"freezeAgreementType" xorm:"not null comment('冻结合约类型 产品种类+产品细类 10001-个人活期存款（一类）；10002-个人活期存款（二类）；10003-个人活期存款（三类）') VARCHAR(5)"`
	FreezeCurrency        string  `json:"freezeCurrency" xorm:"not null comment('币别') VARCHAR(3)"`
	FreezeAmount          float64 `json:"freezeAmount" xorm:"comment('冻结金额') DECIMAL(18,2)"`
	FreezeDocNumber       string  `json:"freezeDocNumber" xorm:"not null comment('冻结文书号') VARCHAR(60)"`
	FreezeReason          string  `json:"freezeReason" xorm:"comment('冻结原因') VARCHAR(200)"`
	EnforcerName1         string  `json:"enforcerName1" xorm:"comment('执法人员名称1') VARCHAR(60)"`
	EnforcerId1           string  `json:"enforcerId1" xorm:"comment('执法人员证件号码1') VARCHAR(60)"`
	EnforcerName2         string  `json:"enforcerName2" xorm:"comment('执法人员名称2') VARCHAR(60)"`
	EnforcerId2           string  `json:"enforcerId2" xorm:"comment('执法人员证件号码2') VARCHAR(60)"`
	UnfreezeFwNumber      int     `json:"unfreezeFwNumber" xorm:"not null comment('解冻优先次序') INT(6)"`
	BankFreezeType        string  `json:"bankFreezeType" xorm:"comment('银行冻结区分M-Internal management;T-tansation') VARCHAR(1)"`
	FreezeStatus          string  `json:"freezeStatus" xorm:"comment('冻结生效标志Y-effect；N-no effect') VARCHAR(1)"`
}

func (*SV000033I) GetServiceKey() string {
	return "SV000033"
}
