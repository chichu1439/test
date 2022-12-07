//Version: v1.0.0.0
package models

type SV000014I struct {
	FreezeBusinessNumber      string  `json:"freezeBusinessNumber" validate:"required,max=30" description:"Original frozen business number"` //原冻结业务编号
	FreezeDate                string  `json:"freezeDate" validate:"required" description:"Thaw date"`                                        //解冻日期
	UnfreezeMode              string  `json:"unfreezeMode" validate:"required,max=1,oneof=4" description:"UnFreeze Mode"`                    //解冻方式 1-账户冻结解冻，2-全部金额解冻，3-部分金额解冻，4-解暂禁
	ChkType                   string  `json:"chkType" validate:"required,max=1,oneof=1 2" description:"Query input type"`                    //查询输入类型 1-合约，2-介质用途
	MediaType                 string  `json:"mediaType" validate:"required,max=3" description:"Media type"`                                  //介质类型
	MediaNumber               string  `json:"mediaNumber" validate:"required,max=30" description:"Media number"`                             //介质号码
	Currency                  string  `json:"currency" validate:"max=3" description:"Currency"`                                              //币种 当查询输入类型是2时
	CashTransFlag             string  `json:"cashTransFlag" validate:"omitempty,max=4,oneof=C T" description:"Currency exchange logo"`       //钞汇标识 当查询输入类型是2时
	ChannelNumber             string  `json:"channelNumber" validate:"max=3" description:"Channel number"`                                   //渠道号 当查询输入类型是2时
	FreezeAgreement           string  `json:"freezeAgreement" validate:"required,max=30" description:"Freeze contract number"`               //冻结合约号
	FreezeAgreementType       string  `json:"freezeAgreementType" validate:"required,max=5" description:"Freeze contract type"`              //冻结合约类型 "产品种类+产品细类 10001-个人活期存款（一类）；10002-个人活期存款（二类）；10003-个人活期存款（三类）"
	CustomerId                string  `json:"customerId" validate:"required,max=20" description:"Frozen customer number"`                    //冻结客户编号
	CustomerType              string  `json:"customerType" description:"Frozen customer type"`                                               //冻结客户类型
	UnfreezeOrgnizeType       string  `json:"unfreezeOrgnizeType" validate:"required,max=1,oneof=1 2" description:"UnFreeze mechanism type"` //解冻机构类型 1-有权机关，2-银行内部
	UnfreezeOrgnizeName       string  `json:"unfreezeOrgnizeName" validate:"required,max=60" description:"Thawing organization name"`        //解冻机构名称 如内部预授权解冻上送申请模块名称（如支付系统发起解冻
	UnfreezeDocumentNumber    string  `json:"unfreezeDocumentNumber" validate:"max=30" description:"Unfreeze instrument number"`             //解冻文书号 解冻结构类型为1-有权机关时必输
	UnfreezeApplicationNumber string  `json:"unfreezeApplicationNumber" validate:"max=30" description:"Application defrost number"`          //申请解冻编号 解冻结构类型为2-银行内部时必输
	UnfreezeAmount            float64 `json:"unfreezeAmount" description:"Unfreeze amount"`                                                  //解冻金额 解冻方式为3-部分金额解冻时必输
	UnfreezeReason            string  `json:"unfreezeReason" validate:"max=60" description:"Reason for thawing"`                             //解冻原因
	UnfreezeOrgnize           string  `json:"unfreezeOrgnize" validate:"max=10" description:"Unfreeze organization"`                         //解冻机构
	UnfreezeEmail             string  `json:"unfreezeEmail" validate:"max=10" description:"Unfreeze teller"`                                 //解冻柜员
	FreezeAuthorityEmail      string  `json:"freezeAuthorityEmail" validate:"max=10" description:"Unfreeze authorized teller"`               //解冻授权柜员
	EnforcerName1             string  `json:"enforcerName1" validate:"max=60" description:"Name of law enforcement officer 1"`               //解冻执法人员名称1 解冻结构类型为1-有权机关时必输
	EnforcerId1               string  `json:"enforcerId1" validate:"max=60" description:"Law enforcement officer ID number 1"`               //解冻执法人员证件号码1 解冻结构类型为1-有权机关时必输
	EnforcerName2             string  `json:"enforcerName2" validate:"max=60" description:"Name of law enforcement officer 2"`               //解冻执法人员名称2
	EnforcerId2               string  `json:"enforcerId2" validate:"max=60" description:"Law enforcement officer ID number 2"`               //解冻执法人员证件号码2
}

type SV000014O struct {
}

func (*SV000014I) GetServiceKey() string {
	return "ssv0000014"
}
