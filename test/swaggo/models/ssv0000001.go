//Version: v0.0.1
package models

import "github.com/shopspring/decimal"

type SSV0000001I struct {
	ProductId             string          `validate:"required" json:"productId" description:"Product id"`             //产品号
	ProductNumber         int             ` json:"productNumber" description:"Product number" validate:"required" `   //产品顺序号
	Currency              string          `validate:"required" json:"currency" description:"Currency"`                //币种
	CashTransFlag         string          `validate:"required" json:"cashTransFlag" description:"Cash trans flag"`    //钞汇标识 cash currency
	AccountingFlag        string          ` json:"accountingFlag" description:"Accounting flag"`                      // 是否计息标志
	MaturityDate          string          `validate:"required" json:"maturityDate" description:"Maturity date"`       //到期日期
	CustomerId            string          `validate:"required" json:"customerId" description:"Customer id"`           //客户编号
	CustomerGrade         string          `  json:"customerGrade" description:"Customer Grade"`                       //客户等级
	CustomerType          string          ` json:"customerType" description:"Customer type"`                          //客户类型
	ShareCode             string          `json:"shareCode" description :"Invite Code"`                               //分享码
	AccountName           string          `validate:"required" json:"accountName" description:"Account name"`         //账户名称/客户名称
	UsageCode             string          ` json:"usageCode" description:"Usage code"`                                //用途代码
	WithdrawMethod        string          `validate:"required" json:"withdrawMethod" description:"Withdraw method"`   //支取方式
	KeyVersion            string          `validate:"required" json:"keyVersion" description:"Key version"`           //公钥版本号
	AccountPassword       string          `validate:"required" json:"accountPassword" description:"Account password"` //账户密码
	OpenAmount            float64         ` json:"openAmount" description:"Open amount"`                              //开户金额
	AccountCancelFlag     string          ` json:"accountCancelFlag" description:"Account cancel flag"`               //销户启用标识
	AutoCancel            string          ` json:"autoCancel" description:"Auto cancel"`                              //是否允许自动销户
	GraceCancel           int             ` json:"graceCancel" description:"Grace cancel"`                            //自动销户宽限天数
	StatementFlag         string          `validate:"required" json:"statementFlag" description:"Statement flag"`     //对账单标识 0-是 1-否 0-Yes 1-No
	StatementMethod       string          `json:"statementMethod"`                                                    //
	StatementFrequency    string          `json:"statementFrequency"`                                                 //
	StatementDay          string          `json:"statementDay"`                                                       //
	StatementAddress      string          `json:"statementAddress"`                                                   //
	StatementEmail        string          `json:"statementEmail"`                                                     //
	AllowCashTransactions string          `json:"allowCashTransactions" description:"Allow Cash Transactions"`        //是否允许现金交易 0-允许现金存取 1-允许现金存，不允许现金取 2-允许现金取，不允许现金存 3-不允许现金存取 0 - Allow cash deposits and withdrawals 1- Allowed cash deposit ，not allowed cash withdrawal 2- Allowed cash withdrawal  , not allowed cash deposit is 3- Not allowed cash deposits and withdrawals
	RetainedBalance       decimal.Decimal `json:"retainedBalance" description:"Retained Balance"`                     //留存余额
	ControlMaximumBalance string          `json:"controlMaximumBalance" description:"Control Maximum Balance"`        //是否最大余额控制 0-控制 1-不控制 0-Yes 1-No
	MaximumBalance        decimal.Decimal `json:"maximumBalance" description:"Maximum Balance"`                       //最大余额
	OverDrawFlag          string          `json:"overDrawFlag" description:"Over draw flag"`                          //透支标识
	NationCode            string          `json:"nationCode" description:"Nation code"`                               //国家
	Address               string          `json:"address" description:"Address" validate:"required"`                  //客户联系地址
	Address1              string          `json:"address1" description:"Address1"`                                    //
	Address2              string          `json:"address2" description:"Address2"`                                    //
	Address3              string          `json:"address3" description:"Address3"`                                    //
	//District                 string  `json:"district" description:"District"`                                                    //
	PostalCode               string `json:"postalCode" description:"Postal code"`                                                  //
	CustomerContactTelephone string `validate:"required" json:"customerContactTelephone" description:"Customer contact telephone"` //客户联系电话
	CustomerContactEmail     string `validate:"required" json:"customerContactEmail" description:"Customer contact email"`         //邮箱地址
	//BranchCreateAccount      string `validate:"required" json:"branchCreateAccount" description:"Branch create account"`           //开户行所
	//BranchAccount            string `validate:"required" json:"branchAccount" description:"Branch account"`                        //账务行所
	AccountOpenDate string `validate:"required" json:"accountOpenDate" description:"Account open date"` //开户日期
	//AccountOpenEmail string `validate:"required" json:"accountOpenEmail" description:"Account open email"` //开户柜员
	//CheckDTransFlag          string  `validate:"required" json:"checkDTransFlag" description:"Check d trans flag"`                  //是否需要堵重检查标志
	//CheckDTransCode          string  `json:"checkDTransCode" description:"Check d trans code"`                                      //堵重随机码
	//RequestId                string  `json:"requestId" description:"Request id"`                                                    //外围流水号
	GlobalBizSeqNo string `json:"globalBizSeqNo" description:"Global biz seq no"` //全局流水号
	SrcBizSeqNo    string `json:"srcBizSeqNo" description:"Src biz seq no"`       //服务流水号
	DepositNature  string `json:"depositNature" description:"Deposit nature"`     //账户性质 00-基本账户;01-一般账户
	MediumType     string `json:"mediumType" description:"Medium type"`           //开户介质类型
	MediumNumber   string `json:"mediumNumber" description:"Medium number"`       //开户介质号码
	//InterestType             string  `json:"interestType" description:"Interest type"`                                              //利率类型
	//StrategyId               string  `json:"strategyId" description:"Strategy id"`                                                  //利率策略ID
	//ItemId                   string  `json:"itemId" description:"Item id"`                                                          //利率ID
	//Rate                     float64 `json:"rate" description:"Rate"`                                                               //执行利率
	//FloatValue               float64 `json:"floatValue" description:"Float value"`                                                  //浮动值
	//InterestMax              float64 `json:"interestMax" description:"Interest max" validate:"required"`                            //最高利息
	//InterestMin              float64 `json:"interestMin" description:"Interest min" validate:"required"`                            //最低利息
	//CalculateType            string  `json:"calculateType" description:"Calculate type" validate:"required"`                        //计息算法
	//ExecuteType              string  `json:"executeType" description:"Execute type" validate:"required"`                            //利率执行方式
	//MonthDay                 string  `json:"monthDay" description:"Month day" validate:"required"`                                  //月计息天数
	//YearDay                  string  `json:"yearDay" description:"Year day" validate:"required"`                                    //年计息天数
	//DecimalFlag              string  `json:"decimalFlag" description:"Decimal flag" validate:"required"`                            //小数是否计息
	//AccrualPeriodType        string  `json:"accrualPeriodType" description:"Accrual period type" validate:"required"`               //计提周期类型
	//AccrualPeriodValue       int     `json:"accrualPeriodValue" description:"Accrual period value" validate:"required"`             //计息周期值
	//SettlePeriodType         string  `json:"settlePeriodType" description:"Settle period type" validate:"required"`                 //计息周期类型
	//SettlePeriodValue        int     `json:"settlePeriodValue" description:"Settle period value" validate:"required"`               //计息周期值
}

type SSV0000001O struct {
	AgreementID   string `json:"agreementID" description:"Agreement i d"`    //合约号
	AgreementType string `json:"agreementType" description:"Agreement type"` //合约类型
	AccountNum    string `json:"accountNum" description:"Account num"`       //账户账号
	Prompt        string `json:"prompt" description:"prompt"`
}

func (*SSV0000001I) GetServiceKey() string {
	return "ssv0000001"
}
