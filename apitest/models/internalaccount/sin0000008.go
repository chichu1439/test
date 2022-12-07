//Version: v1.0.0.0
package models

type IN000008I struct {
	AccountingItemNumber   string  `json:"accountingItemNumber" description:"Accounting item number"`      // 科目号
	SerialNumber           string  `json:"serialNumber" description:"Serial number"`                       // 内部户顺序号
	AccountName            string  `json:"accountName" description:"Account name"`                         // 账户名称
	BranchesRange          string  `json:"branchesRange" description:"Branches range"`                     // 核算机构范围
	CurrencyRange          string  `json:"currencyRange" description:"Currency range"`                     // 货币范围
	DebitCreditFlag        string  `json:"debitCreditFlag" description:"Debit credit flag"`                // 余额方向控制
	BalanceNegativeNumber  string  `json:"balanceNegativeNumber" description:"Balance negative number"`    // 余额红字标识
	TypeApArRegistered     string  `json:"typeApArRegistered" description:"Type ap ar registered"`         // 挂账类别
	TypeApAr               int     `json:"typeApAr" description:"Type ap ar"`                              // 挂销账种类
	DeadlineApAr           string  `json:"deadlineApAr" description:"Deadline ap ar"`                      // 挂账期限
	TermUnit               string  `json:"termUnit" description:"Term unit"`                               // 挂帐期限单位
	DebitLimit             float64 `json:"debitLimit" description:"Debit limit"`                           // 借方限制额度
	CreditLimit            float64 `json:"creditLimit" description:"Credit limit"`                         // 贷方限制额度
	CheckBalanceIndicator  string  `json:"checkBalanceIndicator" description:"Check balance indicator"`    // 余额检查标识
	EffectiveDate          string  `json:"effectiveDate" description:"Effective date"`                     // 生效日期
	ExpiryDate             string  `json:"expiryDate" description:"Expiry date"`                           // 失效日期
	AutoAccountOpeningFlag string  `json:"autoAccountOpeningFlag" description:"Auto account opening flag"` // 自动开户标识
}

type IN000008O struct {
}

func (*IN000008I) GetServiceKey() string {
	return "IN000008"
}
