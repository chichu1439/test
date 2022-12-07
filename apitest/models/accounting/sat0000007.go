//Version: v1.0.0.0
package models

type AT000007I struct {
	AccountingItemNumber string `json:"accountingItemNumber" description:"Accounting item number"` //科目号
	SerialNumber         string `json:"serialNumber" description:"Serial number"`                  //内部户顺序号
}

type AT000007O struct {
	AccountingItemAttribute   string `json:"accountingItemAttribute" description:"Accounting item attribute"`     //科目方向
	AccountingItemType        string `json:"accountingItemType" description:"Accounting item type"`               //科目类型
	AccountingItemName        string `json:"accountingItemName" description:"Accounting item name"`               //科目名称
	AccountingItemLevel       string `json:"accountingItemLevel" description:"Accounting item level"`             //科目级别
	AccountingItemFlag        string `json:"accountingItemFlag" description:"Accounting item flag"`               //是否最末级
	AccountingItemDistinction string `json:"accountingItemDistinction" description:"Accounting item distinction"` //分户标志
	UpperLevelAccountingItem  string `json:"upperLevelAccountingItem" description:"Upper level accounting item"`  //上级科目号
	NegativeNumberFlag        string `json:"negativeNumberFlag" description:"Negative number flag"`               //科目红字标志
	Status                    string `json:"status" description:"Status"`                                         //科目状态 0-normal 1-suspend 2-cancelled
	EffectiveDate             string `json:"effectiveDate" description:"Effective date"`                          //生效日期
	ExpiryDate                string `json:"expiryDate" description:"Expiry date"`                                //失效日期
}

func (*AT000007I) GetServiceKey() string {
	return "AT000007"
}
