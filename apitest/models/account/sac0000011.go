//Version: v1.0.0.0
package models

type AC000011I struct {
	AcctNo                 string `json:"acctNo" validate:"required,max=32" description:"Accounting account"`                                            //核算账号
	OriginalTransDate      string `json:"originalTransDate" validate:"required,max=10" description:"Original transaction date"`                          //原交易日期
	OriginalTransFlow      string `json:"originalTransFlow" validate:"required,max=40" description:"Original transaction flow"`                          //原交易流水
	OriginalGlobalBizSeqNo string `json:"originalGlobalBizSeqNo" validate:"required,max=40" description:"The original transaction global stream number"` //原交易全局流水号
	TransDate              string `json:"transDate" validate:"required,max=10" description:"Trade date"`                                                 //交易日期
	TransFlow              string `json:"transFlow" validate:"required,max=10" description:"Trade time"`                                                 //交易时间
}

type AC000011O struct {
	AmountFreezing  float64 `json:"amountFreezing" description:"Freez amount"`                   //冻结金额
	AmountCurrent   float64 `json:"amountCurrent" description:"Current balance"`                 //当前余额
	AmountLast      float64 `json:"amountLast" description:"The balance of the previous period"` //上期余额
	AmountAvailable float64 `json:"amountAvailable" description:"Available balance"`             //可用余额
	AccountStatus   string  `json:"accountStatus" description:"Account status"`                  //账户状态 1-有效 9-无效
}

func (*AC000011I) GetServiceKey() string {
	return "AC000011"
}
