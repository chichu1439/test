//Version: v0.0.1
package models

type SPR0000013I struct {
	TypeDeduction             string     `json:"typeDeduction" description:"Deduction type(0-media, 1-agreement)"`                                                                                                                                 //扣款方式
	FeeDeductionAgreement     string     `json:"feeDeductionAgreement" description:"Fee deduction agreement ID"`                                                                                                                                   //扣费合约号
	FeeDeductionAgreementType string     `json:"feeDeductionAgreementType" description:"Type of Payment Contract(10001-personal current deposit (class I); 10002-personal current deposit (class II); 10003-personal current deposit(Class III))"` //扣费合约类型
	FeeDeductionMediumNM      string     `json:"feeDeductionMediumNm" description:"Payer's media number"`                                                                                                                                          //扣费介质号码
	FeeDeductionMediumType    string     `json:"feeDeductionMediumType" description:"Payer medium type(VCD-virtual card)"`                                                                                                                         //扣费介质类型
	Currency                  string     `json:"currency" description:"Pay in currency(THB)"`                                                                                                                                                      //币别
	CashTransFlag             string     `json:"cashTransFlag" description:"Payor's currency sign( C-cash account;T-remittance account)"`                                                                                                          //钞汇标志
	FeeDeductionAccountName   string     `json:"feeDeductionAccountName" description:"Payment Account Name"`                                                                                                                                       //账户名称
	UsageCode                 string     `json:"usageCode" description:"Use of the code(EAC-electronic account)"`                                                                                                                                  //用途代码
	FeeDeductionChannel       string     `json:"feeDeductionChannel" description:"Channel number"`                                                                                                                                                 //渠道号
	WithdrawMethod            string     `json:"withdrawMethod" description:"Withdrawal way(0-password;1-seal;2-signature;3-fingerprint)"`                                                                                                         //支取方式
	NeedPasswordFlag          string     `json:"needPasswordFlag" description:"Whether it needs to be inspected(Y-need to verify the secret; N-do not need to verify the secret)" `                                                                //是否需要验密
	KeyVersion                string     `json:"keyVersion" description:"Version of key"`
	AccountPassword           string     `json:"accountPassword" description:"The account password"`      //账户密码
	TransType                 string     `json:"transType" description:"Way to trade(1-online; 2-batch)"` //交易方式
	ExternalSystem            string     `json:"externalSystem" description:"external System"`            //外围系统号
	ExternalSystemFlow        string     `json:"externalSystemFlow" description:"external System Flow"`   //外围流水号
	FeeArray                  []FeeArray `json:"feeArray" description:"Fee array"`                        //
}

type FeeArray struct {
	FeeAbstractCode     string  `json:"feeAbstractCode" description:"Fee abstract code"`                                                                                                                               //费用摘要
	BeFeeDeductionFlag  string  `json:"beFeeDeductionFlag" description:"The flag of fee to be deducted(Y- to be deducted;N-deduct the fee online)"`                                                                    //转待扣标志
	FeeTobeDeducted     float64 `json:"feeTobeDeducted" description:"Fee to be deducted"`                                                                                                                              //应扣金额
	FeePreferential     float64 `json:"feePreferential" description:"Preferential fee"`                                                                                                                                //减免金额
	FeeActuallyDeducted float64 `json:"feeActuallyDeducted" description:"Fee actually deducted"`                                                                                                                       //实扣金额
	StrategyId          string  `json:"strategyId" description:"Fee strategy ID"`                                                                                                                                      //策略ID
	FeeRateId           string  `json:"feeRateId" description:"Fee rate ID"`                                                                                                                                           //费率ID
	FeeAmortizationFlag string  `json:"feeAmortizationFlag" description:"Fee amortization flag(Y- need to amortization;N-no need to amortization)"`                                                                    //转摊销标志
	AmortizePeriodType  int     `json:"amortizePeriodType" description:"Amortize period type(01- By day;02- Weekly;03- By half month/biweekly;04- monthly;05- bimonthly;06- quarterly;07- By half a year;08- yearly)"` //摊销期数
}

type SPR0000013O struct {
}

func (*SPR0000013I) GetServiceKey() string {
	return "PR000013"
}
