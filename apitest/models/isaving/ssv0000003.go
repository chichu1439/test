//Version: v0.01
package models

type SSV0000003I struct {
	QueryType     string `json:"queryType" description:"Query type" validate:"required,max=30"` //查询类型
	MediaType     string `json:"mediaType" description:"Medium type" validate:"max=4"`          //介质类型
	MediaNumber   string `json:"mediaNumber" description:"Medium number" validate:"max=40"`     //介质号码
	AgreementId   string `json:"agreementId" description:"Contract number" validate:"max=30"`   //合约号
	AgreementType string `json:"agreementType" description:"Contract type" validate:"max=5"`    //合约类型
	CustomerId    string `json:"customerId" description:"Customer number" validate:"max=20"`    //客户编号
	Currency      string `json:"currency" description:"Currency"`                               //币种
	CashtranFlag  string `json:"cashtranFlag" description:"Flag of cash and remittance"`        //钞汇标识
	ChannelNumber string `json:"channelNumber" validate:"max=3" description:"Channel number"`   //渠道号
	UsageCode     string `json:"usageCode" validate:"max=3" description:"Usage code"`           //用途代码
}

type SSV0000003O struct {
	Records []ContInfo `json:"records" description:"Contract information"` //合约信息

}

type ContInfo struct {
	AgreementType            string  `json:"agreementType" description:"Contract type"`                     //合约类型
	AgreementId              string  `json:"agreementId" description:"Contract number"`                     //合约号
	Currency                 string  `json:"currency" description:"Currency"`                               //币种
	CashtranFlag             string  `json:"cashtranFlag" description:"Flag of cash and remittance"`        //钞汇标识
	Balance                  float64 `json:"balance" description:"Balance"`                                 //余额
	FrzAmount                float64 `json:"frzAmount" description:"Frozen amount"`                         //冻结金额
	AvlBalance               float64 `json:"avlBalance" description:"Available balance"`                    //可用余额
	AccountStatus            string  `json:"accountStatus" description:"Account status"`                    //账户状态
	AccountOpenDate          string  `json:"accountOpenDate" description:"Account opening date"`            //开户日期
	CustomerContactTelephone string  `json:"customerContactTelephone" description:"Mobile phone number"`    //手机号
	CustomerContactAddress   string  `json:"customerContactAddress" description:"Customer contact address"` //联系地址
	CustomerContactEmail     string  `json:"customerContactEmail" description:"Customer email"`             //邮箱
	AccountName              string  `json:"accountName" description:"Account name"`                        //账户姓名
	FreezeType               string  `json:"freezeType" description:"Frozen state"`                         //冻结状态
	AgreementStatus          string  `json:"agreementStatus" description:"Conttract state"`                 //合约状态
	AccountingAccount        string  `json:"accountingAccount" description:"Accounting account number"`     //核算账号
	DefaultAgreement         string  `json:"defaultAgreement" description:"Default contract ID"`            //默认合约标识
	AccountPassword          string  `json:"accountPassword" description:"Account password"`                //账户密码
	WithdrawMethod           string  `json:"withdrawMethod" description:"Withdrawal method"`                //支取方式
	DepcreFlag               string  `json:"depcreFlag" description:"Debit credit control mark"`            //借贷记控制标记
	CustomerId               string  `json:"customerId" description:"Customer number"`                      //客户编号
	ProductId                string  `json:"productId" description:"Product number"`                        //产品号
	ProductNumber            int     `json:"productNumber" description:"Product serial number"`             //产品顺序号
	LastSettleAmount         float64 `json:"lastSettleAmount" description:"Last settle amount"`
	LastSettleDate           string  `json:"lastSettleDate" description:"Last settle date"`
}

func (*SSV0000003I) GetServiceKey() string {
	return "ssv0000003"
}
