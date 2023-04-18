package models

import (
	"github.com/shopspring/decimal"
)

type SV000030I struct {
	AgreementId   string `json:"agreementId" description:"Agreement Id" validate:"required"` //合约号
	AgreementType string `json:"agreementType" description:"Agreement type"`                 //合约类型

}

type SV000030O struct {
	AgreementType            string          `json:"agreementType" description:"Contract type"`                         //合约类型
	AgreementId              string          `json:"agreementId" description:"Contract number"`                         //合约号
	Currency                 string          `json:"currency" description:"Currency"`                                   //币种
	CashtranFlag             string          `json:"cashTransFlag" description:"Flag of cash and remittance"`           //钞汇标识
	Balance                  float64         `json:"balance" description:"Balance"`                                     //余额
	FreezeAmount             float64         `json:"freezeAmount" description:"Frozen amount"`                          //冻结金额
	AvlBalance               float64         `json:"avlBalance" description:"Available balance"`                        //可用余额
	AccountStatus            string          `json:"accountStatus" description:"Account status"`                        //账户状态
	AccountOpenDate          string          `json:"accountOpenDate" description:"Account opening date"`                //开户日期
	CustomerContactTelephone string          `json:"customerContactTelephone" description:"Customer contact telephone"` //客户联系电话
	CustomerContactEmail     string          `json:"customerContactEmail" description:"Customer contact email"`         //邮箱地址
	PostalCode               string          `json:"postalCode" description:"Postal code"`                              //
	NationCode               string          `json:"nationCode" description:"Nation code"`                              //国家
	Address1                 string          `json:"address1"`
	Address2                 string          `json:"address2"`
	Address3                 string          `json:"address3"`
	Address                  string          `json:"address" description:"Customer contact address"`            //联系地址
	AccountName              string          `json:"accountName" description:"Account name"`                    //账户姓名
	FreezeType               string          `json:"freezeType" description:"Frozen state"`                     //冻结状态
	AgreementStatus          string          `json:"agreementStatus" description:"Conttract state"`             //合约状态
	AccountingAccount        string          `json:"accountingAccount" description:"Accounting account number"` //核算账号
	AccountPassword          string          `json:"accountPassword" description:"Account password"`            //账户密码
	WithdrawMethod           string          `json:"withdrawMethod" description:"Withdrawal method"`            //支取方式
	DepcreFlag               string          `json:"depcreFlag" description:"Debit credit control mark"`        //借贷记控制标记
	CustomerId               string          `json:"customerId" description:"Customer number"`                  //客户编号
	ProductId                string          `json:"productId" description:"Product number"`                    //产品号
	ProductNumber            int             `json:"productNumber" description:"Product serial number"`         //产品顺序号
	BranchNumber             string          `json:"branchNumber"`
	LimitAmount              decimal.Decimal `json:"limitAmount" description:"Limit Amount"`
	UserId                   string          `json:"userId"`
	ExpirationDay            string          `json:"expirationDay"`
	ClosingDate              string          `json:"closingDate"`
	LastSettleAmount         float64         `json:"lastSettleAmount"`
	LastSettleDate           string          `json:"lastSettleDate"`
	CustomerGrade            string          `json:"customerGrade"`
	OverDrawFlag             string          `json:"overDrawFlag" description:"Over draw flag"`
	AllowFreezing            string          `json:"allowFreezing" description:"Allow freezing""`
	StatementFlag            string          `json:"statementFlag" description:"Statement flag"`
	DefaultInterestRate      decimal.Decimal `json:"defaultInterestRate"`      //账户默认利率
	PrefeInteresrRate        decimal.Decimal `json:"prefeInteresrRate"`        // 账户等级利率
	PrefeInterestRateLimit   decimal.Decimal `json:"prefeInterestRateLimit"`   //优惠利率限额
	AmountStartCalcuInterest decimal.Decimal `json:"amountStartCalcuInterest"` //起息金额
	AccountFloatRate         decimal.Decimal `json:"AccountFloatRate"`         //优惠利率
	AccountFloatDirection    string          `json:"accountFloatDirection"`    //等级利率浮动方向
	StatementMethod          string          `json:"statementMethod"`          //
	StatementFrequency       string          `json:"statementFrequency"`       //
	StatementDay             string          `json:"statementDay"`             //
	StatementAddress         string          `json:"statementAddress"`         //
	StatementEmail           string          `json:"statementEmail"`
	AllowCashTransactions    string          `json:"allowCashTransactions" description:"Allow Cash Transactions"` //是否允许现金交易 0-允许现金存取 1-允许现金存，不允许现金取 2-允许现金取，不允许现金存 3-不允许现金存取 0 - Allow cash deposits and withdrawals 1- Allowed cash deposit ，not allowed cash withdrawal 2- Allowed cash withdrawal  , not allowed cash deposit is 3- Not allowed cash deposits and withdrawals
	ProductName              string          `json:"productName"`
	MaxBalance               decimal.Decimal `json:"maxBalance"`          //最大金额
	RetainedBalance          decimal.Decimal `json:"retainedBalance"`     //留存金额
	IsControlMaxBalance      string          `json:"isControlMaxBalance"` //是否控制最大金额
	MediumList               []MediumInfo    `json:"mediumList"`
}

type MediumInfo struct {
	MediumType   string `json:"mediumType" description:"Medium type"`     //开户介质类型
	MediumNumber string `json:"mediumNumber" description:"Medium number"` //开户介质号码

}

func (*SV000030I) GetServiceKey() string {
	return "SV000030"
}
