//Version: v1.0.0.0
package models

import "github.com/shopspring/decimal"

type SV000016I struct {
	AgreementID         string          `json:"agreementID" validate:"required" description:"Agreement id"`                           //合约号
	AgreementType       string          `json:"agreementType" description:"Agreement type"`                                           //合约类型
	Currency            string          `json:"currency" validate:"omitempty,max=3" description:"Currency"`                           //币种 当查询方式为A-按照介质时，必输
	CashTransFlag       string          `json:"cashTransFlag" validate:"omitempty,max=2,oneof=C T" description:"Money exchange logo"` //钞汇标志 C-钞户;T-汇户;-当查询方式为A-按照介质时，必输
	ChannelType         string          `json:"channelType" validate:"omitempty,max=4" description:"Trading channel"`                 //交易渠道 当查询方式为A-按照介质时，必输
	UsageCode           string          `json:"usageCode" validate:"omitempty,max=3" description:"Purpose code"`                      //用途代码 当查询方式为A-按照介质时，必输，EAC-电子账号
	BeginDate           string          `json:"beginDate" validate:"max=30" description:"Transaction start date"`                     //交易开始日期
	EndDate             string          `json:"endDate" validate:"max=30" description:"Transaction end date"`                         //交易结束日期
	BeginAmount         decimal.Decimal `json:"beginAmount" validate:"omitempty" description:"Starting amount"`                       //起始金额 17 2
	EndAmount           decimal.Decimal `json:"endAmount" validate:"omitempty" description:"Closing amount"`                          //结束金额 17 2
	GlobalBusinessSeqNo string          `json:"globalBusinessSeqNo" validate:"omitempty,max=64" description:"Global serial number"`   //全局流水号
	SrcBizSeqNo         string          `json:"srcBizSeqNo" description:"Business serial number"`
	PageRecCount        int             `json:"pageRecCount" validate:"max=9999,min=1" description:"Number of records per page"` //每页记录数
	PageNo              int             `json:"pageNo" validate:"max=9999999999,min=1" description:"Query pages"`                //查询页数
}

type SV000016O struct {
	PageRecCount int                  `json:"pageRecCount" validate:"len=4" description:"Number of records per page"` //每页记录数
	PageNo       int                  `json:"pageNo" validate:"len=10" description:"Query pages"`                     //查询页数
	TotalRecord  int                  `json:"totalRecord" validate:"len=10" description:"total"`                      //总记录数
	TotalPage    int                  `json:"totalPage" validate:"len=10" description:"total pages"`                  //总页数
	Records      []TransactionRecords `json:"records" description:"Records"`
}

type TransactionRecords struct {
	TranDate            string          `json:"tranDate" description:"transaction date"`                //交易日期 Y
	TranTime            string          `json:"tranTime" description:"transaction hour"`                //交易时间 Y
	GlobalBusinessSeqNo string          `json:"globalBusinessSeqNo" description:"Global serial number"` //全局流水号 40 Y
	SrcBizSeqNo         string          `json:"srcBizSeqNo" description:"Business serial number"`
	ChannelType         string          `json:"channelType" description:"Trading channel"`                              //交易渠道 4 Y
	TransDirect         string          `json:"transDirect" description:"Transaction direction"`                        //交易方向 1 Y 1、转入 2、转出
	Agreement           string          `json:"agreementId" description:"Contract ID"`                                  //合约号 30 Y
	AgreementType       string          `json:"agreementType" description:"Contract type"`                              //合约类型 5 Y
	MediumNumber        string          `json:"mediumNumber" description:"Media number"`                                //介质号码 30 N
	MediumType          string          `json:"mediumType" description:"Media type"`                                    //介质类型 3 N
	CustId              string          `json:"custId" description:"Customer Number"`                                   //客户编号 10 N
	CustTy              string          `json:"custTy" description:"Customer type"`                                     //客户类型 1 N 1-个人客户;2-内部客户
	Currency            string          `json:"currency" description:"Currency"`                                        //币别 3 Y 156-人民币
	CashTransFlag       string          `json:"cashTransFlag" description:"Money exchange logo"`                        //钞汇标志 2 N C-钞户;T-汇户;-
	PaymentAccountName  string          `json:"paymentAccountName" description:"account name"`                          //账户名称 200 N
	UsageCode           string          `json:"usageCode" description:"Purpose code"`                                   //用途代码 3 N EAC-电子账号
	TransAmount         decimal.Decimal `json:"transAmount" description:"The transaction amount"`                       //交易金额 17 2 Y
	AmountCurrent       decimal.Decimal `json:"amountCurrent" description:"Account Balance"`                            //账户余额 17 2 Y
	AmountAvailable     decimal.Decimal `json:"amountAvailable" description:"Available Balance"`                        //可用余额 17 2 Y
	TranCtrAcc          string          `json:"tranCtrAcc" description:"Counterparty account"`                          //交易对手账号 30 N
	CltAccountName      string          `json:"cltAccountName" description:"Counterparty account name"`                 //交易对手户名 200 N
	Postscript          string          `json:"postscript" description:"Transfer postscript"`                           //转账附言 200 N
	TransType           string          `json:"transType" description:"means of transaction(1-Transfer,2-Cash 3- Fee)"` //交易方式 1 Y 1-联机；2-批量
	ReverseTradeFlag    string          `json:"reverseTradeFlag" description:"Reversal transaction sign"`               //冲正交易标志 1 Y N-正常交易;C-冲销交易;B-已冲正交易
	BusinessSerial      string          `json:"businessSerial" description:"Business serial number"`                    //业务流水号
	TransactionMethod   string          `json:"transactionMethod" decription:"L-Online; B-Batch"`                       //1-客户账户；2-内部账户；3-合约
	BankName            string          `json:"bankName"`
}

func (*SV000016I) GetServiceKey() string {
	return "SV000016"
}
