//Version: v1.0.0.0
package models

type SSV0000010I struct {
	CheckType         string `validate:"max=1,oneof=1 2" json:"checkType" description:"Check type"`              //查询输入类型 1-合约,2-介质
	DccMediumNumber   string `validate:"required,max=30" json:"dccMediumNumber" description:"Dcc medium number"` //介质号码
	DccMediumType     string `validate:"required,max=3" json:"dccMediumType" description:"Dcc medium type"`      //介质类型
	DccAgreement      string `validate:"max=30" json:"dccAgreement" description:"Dcc agreement"`                 //合约号
	DccAgreementType  string `validate:"max=5" json:"dccAgreementType" description:"Dcc agreement type"`         //合约类型
	Currency          string `validate:"max=3" json:"currency" description:"Currency"`                           //币别 当查询输入类型是2时，本字段必输
	CashTransFlag     string `validate:"max=4" json:"cashTransFlag" description:"Cash trans flag"`               //炒汇标识 当查询输入类型是2时，本字段必输
	ChannelNumber     string `validate:"max=4" json:"channelNumber" description:"Channel number"`                //渠道号 当查询输入类型是2时，本字段必输
	UsageCode         string `validate:"max=3" json:"usageCode" description:"Usage code"`                        //用途代码 当查询输入类型是2时，本字段必输
	DccMode           string `validate:"required,max=2,oneof=D C F N" json:"dccMode" description:"Dcc mode"`     //控制方向设置
	Remark            string `json:"remark" description:"Remark"`                                                //备注
	DccDate           string `json:"dccDate" description:"Dcc date"`                                             //控制日期
	DccOrganize       string `json:"dccOrganize" description:"Dcc organize"`                                     //控制机构
	DccOrganizeName   string `validate:"max=60" json:"dccOrganizeName" description:"Dcc organize name"`          //控制机构名称
	DccDocumentNumber string `validate:"max=30" json:"dccDocumentNumber" description:"Dcc document number"`      //控制文书号
}

type SSV0000010O struct {
	Status string `json:"status" description:"Status"`
}

func (*SSV0000010I) GetServiceKey() string {
	return "ssv0000011"
}
