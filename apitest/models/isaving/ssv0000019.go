//Version: v1.0.0.0
package models

type SV000019I struct {
	ChkTpy               string `validate:"required,max=1" json:"chkTpy" description:"Chk tpy"`                      //查询类型 A-冻结信息查询，B-直接扣划冻结信息查询，C-冻结扣划信息查询
	CheckType            string `validate:"required,max=1" json:"checkType" description:"Check type"`                //查询输入类型 1-合约 2-介质用途
	FreezeBusinessNumber string `validate:"max=30" json:"freezeBusinessNumber" description:"Freeze business number"` //冻结业务编号
	FreezeAgreement      string `validate:"max=30" json:"freezeAgreement" description:"Freeze agreement"`            //冻结合约号
	FreezeAgreementType  string `validate:"max=5" json:"freezeAgreementType" description:"Freeze agreement type"`    //冻结合约类型
	MediumType           string `validate:"required,max=3" json:"mediumType" description:"Medium type"`              //介质类型
	MediumNumber         string `validate:"required,max=60" json:"mediumNumber" description:"Medium number"`         //介质号码
	Currency             string `validate:"max=3" json:"currency" description:"Currency"`                            //币别 当查询输入类型是2时，本字段必输
	CashTransFlag        string `validate:"max=4" json:"cashTransFlag" description:"Cash trans flag"`                //炒汇标识 当查询输入类型是2时，本字段必输
	ChannelNumber        string `validate:"max=4" json:"channelNumber" description:"Channel number"`                 //渠道号 当查询输入类型是2时，本字段必输
	UsageCode            string `validate:"max=3" json:"usageCode" description:"Usage code"`                         //用途代码 当查询输入类型是2时，本字段必输

}

type SV000019O struct {
	FreezeBusinessNumber string  `json:"freezeBusinessNumber" description:"Freeze business number"` //冻结业务编号
	FreezeType           string  `json:"freezeType" description:"Freeze type"`                      //冻结方式
	EffectiveDate        string  `json:"effectiveDate" description:"Effective date"`                //冻结生效日
	FreezeOweDate        string  `json:"freezeOweDate" description:"Freeze owe date"`               //冻结到期日
	FreezeDate           string  `json:"freezeDate" description:"Freeze date"`                      //冻结日期
	FreezeOrgnize        string  `json:"freezeOrgnize" description:"Freeze orgnize"`                //冻结行所
	FreezeEmail          string  `json:"freezeEmail" description:"Freeze email"`                    //冻结柜员
	FreezeAuthorityEmail string  `json:"freezeAuthorityEmail" description:"Freeze authority email"` //冻结授权柜员
	FreezeOrgnizeType    string  `json:"freezeOrgnizeType" description:"Freeze orgnize type"`       //冻结发起方
	FreezeOrgnizeName    string  `json:"freezeOrgnizeName" description:"Freeze orgnize name"`       //冻结发起方名称
	FreezeMode           string  `json:"freezeMode" description:"Freeze mode"`                      //冻结方式
	CustomerId           string  `json:"customerId" description:"Customer id"`                      //冻结客户编号
	CustomerType         string  `json:"customerType" description:"Customer type"`                  //冻结客户类型
	MediumType           string  `json:"mediumType" description:"Medium type"`                      //介质类型
	MediumNumber         string  `json:"mediumNumber" description:"Medium number"`                  //介质号码
	FreezeAgreement      string  `json:"freezeAgreement" description:"Freeze agreement"`            //冻结合约号
	FreezeAgreementType  string  `json:"freezeAgreementType" description:"Freeze agreement type"`   //冻结合约类型
	FreezeCurrency       string  `json:"freezeCurrency" description:"Freeze currency"`              //币别
	FreezeAmount         float64 `json:"freezeAmount" description:"Freeze amount"`                  //冻结金额
	FreezeDocnumber      string  `json:"freezeDocnumber" description:"Freeze docnumber"`            //冻结文书号
	FreezeReason         string  `json:"freezeReason" description:"Freeze reason"`                  //冻结原因
	EnforcerName1        string  `json:"enforcerName1" description:"Enforcer name1"`                //执法人员名称1
	EnforcerId1          string  `json:"enforcerId1" description:"Enforcer id1"`                    //执法人员证件号码1
	EnforcerName2        string  `json:"enforcerName2" description:"Enforcer name2"`                //执法人员名称2
	EnforcerId2          string  `json:"enforcerId2" description:"Enforcer id2"`                    //执法人员证件号码2
	AntDdctbleBln        float64 `json:"antDdctbleBln" description:"Ant ddctble bln"`               //账户可扣划余额
	AgreementStatus      string  `json:"agreementStatus" description:"Agreement status"`            //合约状态 0-正常;1-不动户;2-销户;
	AgreementFreezeType  string  `json:"agreementFreezeType" description:"Agreement freeze type"`   //合约冻结状态 0-正常;1-合约冻结/账户冻结;2-金额冻结;3-暂禁

}

func (*SV000019I) GetServiceKey() string {
	return "ssv0000019"
}
