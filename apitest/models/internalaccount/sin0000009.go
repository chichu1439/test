//Version: v1.0.0.0
package models

type IN000009I struct {
	AccountingItemNumber string `json:"accountingItemNumber" description:"Accounting item number"` // 科目号
	SerialNumber         string `json:"serialNumber" description:"Serial number"`                  // 内部户顺序号
}

type IN000009O struct {
}

func (*IN000009I) GetServiceKey() string {
	return "IN000009"
}
