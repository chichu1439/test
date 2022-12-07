//Version: v1.0.0.0
package models

type SV000020I struct {
	QueryType     string `json:"queryType" description:"Query type" validate:"required,max=1"`      //查询类型 0-介质，1-合约
	MediaType     string `json:"mediaType" description:"Media type" validate:"required,max=3"`      //介质类型 VCD-虚拟卡
	MediaNumber   string `json:"mediaNumber" description:"Media number" validate:"required,max=30"` //介质号码
	AgreementId   string `json:"agreementId" description:"Agreement id" validate:"max=30"`          //合约号 当查询类型为1-合约时，本字段必输
	AgreementType string `json:"agreementType" description:"Agreement type" validate:"max=5"`       //合约类型 当查询类型为1-合约时，本字段必输
}

type SV000020O struct {
	TransDate string             `json:"transDate" description:"Trans date"` //交易日期
	Records   []SV000020ORecords `json:"records" description:"Records"`
}

type SV000020ORecords struct {
	AgreementType            string  `json:"agreementType" description:"Agreement type"`                        //合约类型
	AgreementID              string  `json:"agreementID" description:"Agreement i d"`                           //合约号
	Currency                 string  `json:"currency" description:"Currency"`                                   //币种
	CashtranFlag             string  `json:"cashtranFlag" description:"Cashtran flag"`                          //钞汇标志 C-钞户;T-汇户
	Balance                  float64 `json:"balance" description:"Balance"`                                     //余额
	FrzAmount                float64 `json:"frzAmount" description:"Frz amount"`                                //冻结金额
	AvlBalance               float64 `json:"avlBalance" description:"Avl balance"`                              //可用余额
	AccountStatus            string  `json:"accountStatus" description:"Account status"`                        //账户状态
	FreezeType               string  `json:"freezeType" description:"Freeze type"`                              //合约冻结状态 0-正常;1-合约冻结/账户冻结;2-金额冻结;3-暂禁
	AccountOpenDate          string  `json:"accountOpenDate" description:"Account open date"`                   //开户日期
	CustomerContactTelephone string  `json:"customerContactTelephone" description:"Customer contact telephone"` //手机号
	WithdrawMethod           string  `json:"withdrawMethod" description:"Withdraw method"`                      //支取方式 0-密码;
	PayInterest              float64 `json:"payInterest" description:"Pay interest"`                            //应付利息
	InterestTax              float64 `json:"interestTax" description:"Interest tax"`                            //利息税
}

func (*SV000020I) GetServiceKey() string {
	return "ssv0000020"
}
