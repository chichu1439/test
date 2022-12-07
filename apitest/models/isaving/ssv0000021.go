//Version: v1.0.0.0
package models

type SV000021I struct {
	CanTpy          string `validate:"required,max=1,oneof=0 1" json:"canTpy" description:"Can tpy"`  //销户类型 0-介质，1-合约
	MediaType       string `validate:"required,max=3" json:"mediaType" description:"Media type"`      //介质类型 VCD-虚拟卡
	MediaNumber     string `validate:"required,max=30" json:"mediaNumber" description:"Media number"` //介质号码
	AgreementId     string `validate:"max=30" json:"agreementId" description:"Agreement id"`          //合约号 当销户类型为1-合约时，本字段必输
	AgreementType   string `validate:"max=5" json:"agreementType" description:"Agreement type"`       //合约类型 当销户类型为1-合约时，本字段必输
	TranAccTy       string `validate:"max=1" json:"tranAccTy" description:"Tran acc ty"`              //资金去向 A-行外，B-行内
	CltMediaNumber  string `validate:"max=30" json:"cltMediaNumber" description:"Clt media number"`   //收款介质号码 当资金去向为B-行内时，本字段必输
	CltMediaType    string `validate:"max=3" json:"cltMediaType" description:"Clt media type"`        //收款介质类型 当资金去向为B-行内时，本字段必输
	KeyVersion      string `validate:"max=8" json:"keyVersion" description:"Key version"`             //公钥版本号 合约支付方式为0-密码时，本字段为必送
	AccountPassword string `validate:"max=400" json:"accountPassword" description:"Account password"` //支付密码 当合约支付方式为0-密码时，本字段为必送
	AccountName     string `validate:"max=200" json:"accountName" description:"Account name"`         //收款户名
	BnName          string `validate:"max=200" json:"bnName" description:"Bn name"`                   //收款银行
	CanBankNumber   string `validate:"max=200" json:"canBankNumber" description:"Can bank number"`    //销户行号
	SmryFlag        string `validate:"max=3" json:"smryFlag" description:"Smry flag"`                 //交易摘要
	Currency        string `validate:"max=3" json:"currency" description:"Currency"`                  //币别 当查询输入类型是0时，本字段必输
	CashTransFlag   string `validate:"max=4" json:"cashTransFlag" description:"Cash trans flag"`      //炒汇标识 当查询输入类型是0时，本字段必输
	ChannelNumber   string `validate:"max=4" json:"channelNumber" description:"Channel number"`       //渠道号 当查询输入类型是0时，本字段必输
	UsageCode       string `validate:"max=3" json:"usageCode" description:"Usage code"`               //用途代码 当查询输入类型是0时，本字段必输
}

type SV000021O struct {
	State string `json:"state" description:"State"`
}

func (*SV000021I) GetServiceKey() string {
	return "ssv0000021"
}
