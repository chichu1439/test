//Version: v1.0.0.0
package models

type SSV0000009I struct {
	FreezeBusinessNumber string `json:"freezeBusinessNumber" validate:"required,max=30" description:"Frozen business number"` //冻结业务编号
	FreezeAgreement      string `json:"freezeAgreement" validate:"required,max=30" description:"Freeze contract number"`      //冻结合约号
	FreezeAgreementType  string `json:"freezeAgreementType" validate:"required,max=5" description:"Freeze contract type"`     //冻结合约类型 "产品种类+产品细类 10001-个人活期存款（一类）；10002-个人活期存款（二类）；10003-个人活期存款（三类）"
	CustomerId           string `json:"customerId" validate:"required,max=20" description:"Frozen customer number"`           //冻结客户编号
	CustomerType         string `json:"customerType" validate:"required,max=2" description:"Frozen customer type"`            //冻结客户类型
}

type SSV0000009O struct {
	FreezeBusinessNumber string  `json:"freezeBusinessNumber" validate:"max=30" description:"Frozen business number"`                             //冻结业务编号
	FreezeType           string  `json:"freezeType" validate:"max=1" description:"Freezing method"`                                               // 冻结方式
	EffectiveDate        string  `json:"effectiveDate" description:"Effective date of freezing"`                                                  //冻结生效日
	FreezeOweDate        string  `json:"freezeOweDate" description:"Freeze expiry date"`                                                          //冻结到期日
	FreezeDate           string  `json:"freezeDate" description:"Freeze date"`                                                                    //冻结日期
	FreezeOrgnize        string  `json:"freezeOrgnize" validate:"max=30" description:"Frozen shop"`                                               //冻结行所
	FreezeEmail          string  `json:"freezeEmail" validate:"max=30" description:"Freeze teller"`                                               //冻结柜员
	FreezeAuthorityEmail string  `json:"freezeAuthorityEmail" validate:"max=30" description:"Freeze authorized tellers"`                          // 冻结授权柜员
	FreezeOrgnizeType    string  `json:"freezeOrgnizeType" validate:"max=1" description:"Freeze initiator"`                                       //冻结发起方
	FreezeOrgnizeName    string  `json:"freezeOrgnizeName" validate:"max=60" description:"Name of freezing initiator"`                            //冻结发起方名称
	FreezeMode           string  `json:"freezeMode" validate:"max=1" description:"Freezing mode"`                                                 //冻结方式
	CustomerId           string  `json:"customerId" validate:"max=20" description:"Freeze customer number"`                                       //冻结客户编号
	CustomerType         string  `json:"customerType" validate:"max=2" description:"Freeze customer type"`                                        //冻结客户类型
	MediumType           string  `json:"mediumType" validate:"max=3" description:"Media type"`                                                    //介质类型
	MediumNumber         string  `json:"mediumNumber" validate:"max=60" description:"Media number"`                                               //介质号码
	FreezeAgreement      string  `json:"freezeAgreement" validate:"max=30" description:"Freeze contract number"`                                  //冻结合约号
	FreezeAgreementType  string  `json:"freezeAgreementType" validate:"max=5" description:"Frozen contract type"`                                 //冻结合约类型
	FreezeCurrency       string  `json:"freezeCurrency" validate:"max=3" description:"currency"`                                                  //币别
	FreezeAmount         float64 `json:"freezeAmount" description:"Frozen amount"`                                                                //冻结金额
	FreezeDocnumber      string  `json:"freezeDocnumber" validate:"max=60" description:"Frozen document number"`                                  //冻结文书号
	FreezeReason         string  `json:"freezeReason" validate:"max=60" description:"Reasons for freezing"`                                       //冻结原因
	EnforcerName1        string  `json:"enforcerName1" validate:"max=60" description:"Name of law enforcement officer 1"`                         //执法人员名称1
	EnforcerId1          string  `json:"enforcerId1" validate:"max=60" description:"Law enforcement officer certificate number 1"`                //执法人员证件号码1
	EnforcerName2        string  `json:"enforcerName2" validate:"max=60" description:"ame of law enforcement officer 2"`                          //执法人员名称2
	EnforcerId2          string  `json:"enforcerId2" validate:"max=60" description:"Law enforcement officer certificate number 2"`                //执法人员证件号码2
	UnfreezeAmount       float64 `json:"unfreezeAmount" description:"Unfreezing amount"`                                                          //解冻金额
	UnfreezeDocnumber    string  `json:"unfreezeDocnumber" validate:"max=60" description:"Unfreezing document number"`                            //解冻文书号
	UnfreezeReason       string  `json:"unfreezeReason" validate:"max=60" description:"Reasons for thawing"`                                      //解冻原因
	InsfezDate           string  `json:"insfezDate" description:"Renewal date"`                                                                   //续冻日期
	FrFreezeAmount       float64 `json:"frFreezeAmount" description:"Original frozen amount"`                                                     //原冻结金额
	UnfreezeFwNumber     int     `json:"unfreezeFwNumber" description:"Thaw priority"`                                                            // 解冻优先次序
	UfEnforcerName1      string  `json:"ufEnforcerName1" validate:"max=60" description:"Name of unfreezing law enforcement officer 1"`            // 解冻执法人员名称1
	UfEnforcerId1        string  `json:"ufEnforcerId1" validate:"max=60" description:"Unfreezing law enforcement personnel certificate number 1"` //解冻执法人员证件号码1
	UfEnforcerName2      string  `json:"ufEnforcerName2" validate:"max=60" description:"Name of unfreezing law enforcement officer 2"`            //解冻执法人员名称2
	UfEnforcerId2        string  `json:"ufEnforcerId2" validate:"max=60" description:"Unfreezing law enforcement personnel certificate number 2"` //解冻执法人员证件号码2
}

func (*SSV0000009I) GetServiceKey() string {
	return "ssv0000009"
}
