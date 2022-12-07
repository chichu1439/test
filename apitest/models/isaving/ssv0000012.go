//Version: v1.0.0.0
package models

type SV000012I struct {
	CheckType               string `validate:"required,max=1" json:"checkType" description:"Check type"`                      //查询方式 A-合约，F-冻结业务编号，D-直接扣划业务编号,M-介质用途
	MediaNumber             string `validate:"required,max=30" json:"chkMediumNumber" description:"Media number"`             //介质号码
	MediaType               string `validate:"required,max=3" json:"chkMediumType" description:"Media type"`                  //介质类型
	AgreementType           string `validate:"max=5" json:"agreementType" description:"Agreement type"`                       //合约类型
	AgreementId             string `validate:"max=30" json:"agreementId" description:"Agreement id"`                          //合约号码
	FreezeBusinessNumber    string `validate:"max=30" json:"freezeBusinessNumber" description:"Freeze business number"`       //冻结业务编号
	DeductionBusinessNumber string `validate:"max=30" json:"deductionBusinessNumber" description:"Deduction business number"` //扣划业务编号
	Currency                string `validate:"max=3" json:"currency" description:"Currency"`                                  //币别 当查询输入类型是M时，本字段必输
	CashTransFlag           string `validate:"max=4" json:"cashTransFlag" description:"Cash trans flag"`                      //炒汇标识 当查询输入类型是M时，本字段必输
	ChannelNumber           string `validate:"max=4" json:"channelNumber" description:"Channel number"`                       //渠道号 当查询输入类型是M时，本字段必输
	UsageCode               string `validate:"max=3" json:"usageCode" description:"Usage code"`                               //用途代码 当查询输入类型是M时，本字段必输
	PageNum                 int    `validate:"max=9999" json:"pageNum" description:"Page number"`                             //查询页数
	PageRecordCount         int    `validate:"max=9999999999" json:"pageRecordCount" description:"Page record count"`         //每页记录数
}

type SV000012O struct {
	PageNum         int               `json:"pageNum" description:"Page number"`               //查询页数
	PageRecordCount int               `json:"pageRecordCount" description:"Page record count"` //每页记录数
	TotalRecord     int               `json:"totalRecord" description:"Total record"`          //总记录数
	TotalPages      int               `json:"totalPages" description:"Total pages"`            //总页数
	FezArr          []SV000012_FezArr `json:"fezArr" description:"Fez arr"`
}

type SV000012_FezArr struct {
	UzFezArr             []SV000012_UzFezArr `json:"uzFezArr" description:"Uz fez arr"`
	RecordType           string              `json:"recordType" description:"Record type"`                      //记录类型
	FreezeBusinessNumber string              `json:"freezeBusinessNumber" description:"Freeze business number"` //冻结业务编号
	FreezeAgreement      string              `json:"freezeAgreement" description:"Freeze agreement"`            //冻结合约号
	FreezeAgreementType  string              `json:"freezeAgreementType" description:"Freeze agreement type"`   //冻结合约类型
	FreezeAccountName    string              `json:"freezeAccountName" description:"Freeze account name"`       //冻结账户名称
	FreezeMode           string              `json:"freezeMode" description:"Freeze mode"`                      //冻结方式
	FreezeType           string              `json:"freezeType" description:"Freeze type"`                      //冻结状态
	FreezeCurrency       string              `json:"freezeCurrency" description:"Freeze currency"`              //币别
	FreezeOrgnizeType    string              `json:"freezeOrgnizeType" description:"Freeze orgnize type"`       //冻结机构类型
	FreezeOrgnizeName    string              `json:"freezeOrgnizeName" description:"Freeze orgnize name"`       //冻结机构名称
	FreezeDocumentNumber string              `json:"freezeDocumentNumber" description:"Freeze document number"` //冻结文书号
	FreezeAmount         float64             `json:"freezeAmount" description:"Freeze amount"`                  //冻结金额
	FreezeReason         string              `json:"freezeReason" description:"Freeze reason"`                  //冻结原因
	EffectiveDate        string              `json:"effectiveDate" description:"Effective date"`                //生效日期
	FreezeOweDate        string              `json:"freezeOweDate" description:"Freeze owe date"`               //冻结到期日
	EnforcerName1        string              `json:"enforcerName1" description:"Enforcer name1"`                //执法人员名称1
	EnforcerId1          string              `json:"enforcerId1" description:"Enforcer id1"`                    //执法人员证件号码1
	EnforcerName2        string              `json:"enforcerName2" description:"Enforcer name2"`                //执法人员名称2
	EnforcerId2          string              `json:"enforcerId2" description:"Enforcer id2"`                    //执法人员证件号码2
	FreezeOrgnize        string              `json:"freezeOrgnize" description:"Freeze orgnize"`                //冻结机构
	FreezeEmail          string              `json:"freezeEmail" description:"Freeze email"`                    //冻结柜员
	FreezeAuthorityEmail string              `json:"freezeAuthorityEmail" description:"Freeze authority email"` //冻结授权柜员

}

type SV000012_UzFezArr struct {
	UnfreezeDate           string  `json:"unfreezeDate" description:"Unfreeze date"`                      //解冻日期
	UnfreezeMode           string  `json:"unfreezeMode" description:"Unfreeze mode"`                      //解冻方式
	UnfreezeAmount         float64 `json:"unfreezeAmount" description:"Unfreeze amount"`                  //解冻金额
	UnfreezeDocnumber      string  `json:"unfreezeDocnumber" description:"Unfreeze docnumber"`            //解冻文书号
	UnfreezeReason         string  `json:"unfreezeReason" description:"Unfreeze reason"`                  //解冻原因
	FrFreezeAmount         float64 `json:"frFreezeAmount" description:"Fr freeze amount"`                 //原冻结金额
	UnfreezeFwNumber       int     `json:"unfreezeFwNumber" description:"Unfreeze fw number"`             //解冻优先次序
	UnfreezeOrgnize        string  `json:"unfreezeOrgnize" description:"Unfreeze orgnize"`                //解冻行所
	UnfreezeEmail          string  `json:"unfreezeEmail" description:"Unfreeze email"`                    //解冻柜员
	UnfreezeAuthorityEmail string  `json:"unfreezeAuthorityEmail" description:"Unfreeze authority email"` //解冻授权柜员
	UfEnforcerName1        string  `json:"ufEnforcerName1" description:"Uf enforcer name1"`               //解冻执法人员名称1
	UfEnforcerId1          string  `json:"ufEnforcerId1" description:"Uf enforcer id1"`                   //解冻执法人员证件号码1
	UfEnforcerName2        string  `json:"ufEnforcerName2" description:"Uf enforcer name2"`               //解冻执法人员名称2
	UfEnforcerId2          string  `json:"ufEnforcerId2" description:"Uf enforcer id2"`                   //解冻执法人员证件号码2
	DeductionDate          string  `json:"deductionDate" description:"Deduction date"`                    //扣划日期
	DeductionTime          string  `json:"deductionTime" description:"Deduction time"`                    //扣划时间
	DeductionMode          string  `json:"deductionMode" description:"Deduction mode"`                    //扣划方式
	DeductionAmount        string  `json:"deductionAmount" description:"Deduction amount"`                //扣划金额
	DeductionDocnumber     string  `json:"deductionDocnumber" description:"Deduction docnumber"`          //扣划文书号
	DeductionReason        string  `json:"deductionReason" description:"Deduction reason"`                //扣划原因
	//FrFreezeAmount         string //原冻结金额
	DeductionOrgnize        string `json:"deductionOrgnize" description:"Deduction orgnize"`                //扣划行所
	DeductionEmail          string `json:"deductionEmail" description:"Deduction email"`                    //扣划柜员
	DeductionAuthorityEmail string `json:"deductionAuthorityEmail" description:"Deduction authority email"` //扣划授权柜员
	DateEnforcerName1       string `json:"dateEnforcerName1" description:"Date enforcer name1"`             //扣划执法人员名称1
	DateEnforcerId1         string `json:"dateEnforcerId1" description:"Date enforcer id1"`                 //扣划执法人员证件号码1
	DateEnforcerName2       string `json:"dateEnforcerName2" description:"Date enforcer name2"`             //扣划执法人员名称2
	DateEnforcerId2         string `json:"dateEnforcerId2" description:"Date enforcer id2"`                 //扣划执法人员证件号码2

}

func (*SV000012I) GetServiceKey() string {
	return "ssv0000012"
}
