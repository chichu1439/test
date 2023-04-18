package models

type SV000037I struct {
	AgreementId    string `json:"agreementId" validate:"required"`
	AgreementType  string `json:"agreementType" validate:"required"`
	BusinessNumber string `json:"businessNumber" validate:"required"`
	DocNumber      string `json:"docNumber" validate:"required"`
}

type SV000037O struct {
	FreezeBusinessNumber string  `json:"freezeBusinessNumber"` // 冻结业务编号自生成
	CnfreezeDate         string  `json:"cnfreezeDate"`         // 续冻日期
	CnfreezeTime         string  `json:"cnfreezeTime"`         // 续冻时间
	CnfreezeMode         string  `json:"cnfreezeMode"`         // 原冻结方式
	CnfezAmount          float64 `json:"cnfezAmount"`          // 续冻金额
	CnfezDocNumber       string  `json:"cnfezDocNumber"`       // 续冻文书号
	CnfreezeReason       string  `json:"cnfreezeReason"`       // 续冻原因
	FrFreezeAmount       float64 `json:"frFreezeAmount"`       // 原冻结金额
	CnfezOrganize        string  `json:"cnfezOrganize"`        // 续冻行所
	CnfreezeTeller       string  `json:"cnfreezeTeller"`       // 续冻柜员
	CnfezAuthorityTeller string  `json:"cnfezAuthorityTeller"` // 续冻授权柜员
	CnfEnforcerName1     string  `json:"cnfEnforcerName1"`     // 续冻执法人员名称1
	CnfEnforcerId1       string  `json:"cnfEnforcerId1"`       // 续冻执法人员证件号码1
	CnfEnforcerName2     string  `json:"cnfEnforcerName2"`     // 续冻执法人员名称2
	CnfEnforcerId2       string  `json:"cnfEnforcerId2"`       // 续冻执法人员证件号码2
}

func (*SV000037I) GetServiceKey() string {
	return "SV000037"
}
