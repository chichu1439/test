//Version: v1.0.0.0
package models

type SSV0000011I struct {
	CheckType          string `validate:"max=1,oneof=1 2" json:"checkType" description:"Check type"`                     //查询输入类型 1-合约,2-介质
	KeyVersion         string `validate:"required,max=5" json:"keyVersion" description:"Key version"`                    //钥版本号
	BeVerifiedPassword string `validate:"required,max=400" json:"beVerifiedPassword" description:"Be verified password"` //旧支付密码
	NewPassword        string `validate:"required,max=400" json:"newPassword" description:"New password"`                //新支付密码
	AgreementType      string `validate:"max=5" json:"agreementType" description:"Agreement type"`                       //合约类型
	AgreementId        string `validate:"max=30" json:"agreementId" description:"Agreement id"`                          //合约号
	MediaType          string `validate:"max=3" json:"mediaType" description:"Media type"`                               //介质类型
	MediaNumber        string `validate:"max=30" json:"mediaNumber" description:"Media number"`                          //介质号码
	Currency           string `validate:"max=3" json:"currency" description:"Currency"`                                  //币别 当查询输入类型是2时，本字段必输
	CashTransFlag      string `validate:"max=4" json:"cashTransFlag" description:"Cash trans flag"`                      //炒汇标识 当查询输入类型是2时，本字段必输
	ChannelNumber      string `validate:"max=4" json:"channelNumber" description:"Channel number"`                       //渠道号 当查询输入类型是2时，本字段必输
	UsageCode          string `validate:"max=3" json:"usageCode" description:"Usage code"`                               //用途代码 当查询输入类型是2时，本字段必输
}

type SSV0000011O struct {
	Status string `json:"status" description:"Status"`
}

func (*SSV0000011I) GetServiceKey() string {
	return "SV000011"
}
