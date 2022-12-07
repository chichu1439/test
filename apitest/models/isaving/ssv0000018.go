//Version: v1.0.0.0
package models

type SV000018I struct {
	DdctnTyp             string  `validate:"required,max=1,oneof=D F" json:"ddctnTyp" description:"Ddctn typ"`        //扣划方式 D-直接扣划，F-冻结扣划
	FreezeBusinessNumber string  `validate:"max=30" json:"freezeBusinessNumber" description:"Freeze business number"` //原冻结业务编号
	DdcDate              string  `validate:"required" json:"ddcDate" description:"Ddc date"`                          //扣划日期
	DdcAccTyp            string  `validate:"required,max=1,oneof=A M" json:"ddcAccTyp" description:"Ddc acc typ"`     //扣划账户类型 A-合约，M-介质
	DdcAgrmt             string  `validate:"max=30" json:"ddcAgrmt" description:"Ddc agrmt"`                          //扣划合约号
	FreezeAgreementType  string  `validate:"max=5" json:"freezeAgreementType" description:"Freeze agreement type"`    //冻结合约类型
	MediumNumber         string  `validate:"required,max=30" json:"mediumNumber" description:"Medium number"`         //扣划介质号码
	MediumType           string  `validate:"required,max=3" json:"mediumType" description:"Medium type"`              //扣划介质类型
	DdcDocNm             string  `validate:"max=30" json:"ddcDocNm" description:"Ddc doc nm"`                         //扣划文书号
	DdcAmt               float64 `validate:"max=99999999999999999" json:"ddcAmt" description:"Ddc amt"`               //扣划金额
	DdcReason            string  `validate:"max=60" json:"ddcReason" description:"Ddc reason"`                        //扣划原因
	DdcMech              string  `validate:"max=10" json:"ddcMech" description:"Ddc mech"`                            //扣划机构
	DdcEm                string  `validate:"max=10" json:"ddcEm" description:"Ddc em"`                                //扣划柜员
	DdcAuthEm            string  `validate:"max=10" json:"ddcAuthEm" description:"Ddc auth em"`                       //扣划授权柜员
	EnforcerName1        string  `validate:"required,max=60" json:"enforcerName1" description:"Enforcer name1"`       //执法人员名称1
	EnforcerId1          string  `validate:"required,max=60" json:"enforcerId1" description:"Enforcer id1"`           //执法人员证件号码1
	EnforcerName2        string  `validate:"max=60" json:"enforcerName2" description:"Enforcer name2"`                //执法人员名称2
	EnforcerId2          string  `validate:"max=60" json:"enforcerId2" description:"Enforcer id2"`                    //执法人员证件号码2
	RvAcc                string  `validate:"required,max=60" json:"rvAcc" description:"Rv acc"`                       //收款账号
	RvBn                 string  `validate:"required,max=60" json:"rvBn" description:"Rv bn"`                         //收款银行
	Currency             string  `validate:"max=3" json:"currency" description:"Currency"`                            //币别 当查询输入类型是2时，本字段必输
	CashTransFlag        string  `validate:"max=4" json:"cashTransFlag" description:"Cash trans flag"`                //炒汇标识 当查询输入类型是2时，本字段必输
	ChannelNumber        string  `validate:"max=4" json:"channelNumber" description:"Channel number"`                 //渠道号 当查询输入类型是2时，本字段必输
	UsageCode            string  `validate:"max=3" json:"usageCode" description:"Usage code"`                         //用途代码 当查询输入类型是2时，本字段必输

}

type SV000018O struct {
	State string `json:"state" description:"State"`
}

func (*SV000018I) GetServiceKey() string {
	return "ssv0000018"
}
