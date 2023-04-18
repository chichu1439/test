package models

type SV000035I struct {
	AgreementId    string `json:"agreementId" validate:"required"`
	AgreementType  string `json:"agreementType" validate:"required"`
	BusinessNumber string `json:"businessNumber" validate:"required"`
	DocNumber      string `json:"docNumber" validate:"required"`
}

type SV000035O struct {
	FreezeBusinessNumber    string  `json:"freezeBusinessNumber" xorm:"not null pk comment('冻结业务编号:自生成：规则：25位 冻结方式（1位）+冻结发起方（1位）+DCN后四位（4位）+日期（19位，精确到纳秒）') VARCHAR(30)"`
	UnfreezeDate            string  `json:"unfreezeDate" xorm:"not null comment('解冻日期') DATE"`
	UnfreezeTime            string  `json:"unfreezeTime" xorm:"comment('解冻时间') TIME"`
	UnfreezeMode            string  `json:"unfreezeMode" xorm:"not null comment('解冻方式   1-账户冻结解冻，2-全部金额解冻，3-部分金额解冻') VARCHAR(1)"`
	UnfreezeAmount          float64 `json:"unfreezeAmount" xorm:"comment('解冻金额') DECIMAL(18,2)"`
	UnfreezeDocNumber       string  `json:"unfreezeDocNumber" xorm:"not null pk comment('解冻文书号') VARCHAR(60)"`
	UnfreezeReason          string  `json:"unfreezeReason" xorm:"comment('解冻原因') VARCHAR(60)"`
	FrFreezeAmount          float64 `json:"frFreezeAmount" xorm:"comment('原冻结金额') DECIMAL(18,2)"`
	UnfreezeFwNumber        int     `json:"unfreezeFwNumber" xorm:"comment('解冻优先次序') INT(6)"`
	UnfreezeOrganize        string  `json:"unfreezeOrganize" xorm:"comment('解冻行所') VARCHAR(30)"`
	UnfreezeTeller          string  `json:"unfreezeTeller" xorm:"comment('解冻柜员') VARCHAR(30)"`
	UnfreezeAuthorityTeller string  `json:"unfreezeAuthorityTeller" xorm:"comment('解冻授权柜员') VARCHAR(30)"`
	UfEnforcerName1         string  `json:"ufEnforcerName1" xorm:"comment('解冻执法人员名称1') VARCHAR(60)"`
	UfEnforcerId1           string  `json:"ufEnforcerId1" xorm:"comment('解冻执法人员证件号码1') VARCHAR(60)"`
	UfEnforcerName2         string  `json:"ufEnforcerName2" xorm:"comment('解冻执法人员名称2') VARCHAR(60)"`
	UfEnforcerId2           string  `json:"ufEnforcerId2" xorm:"comment('解冻执法人员证件号码2') VARCHAR(60)"`
}

func (*SV000035I) GetServiceKey() string {
	return "SV000035"
}
