//Version: v1.0.0.0
package models

type IN000010I struct {
	AccountingItemNumber string `json:"accountingItemNumber" description:"Accounting item number"` // 科目号
	SerialNumber         string `json:"serialNumber" description:"Serial number"`                  // 内部户顺序号
}

type IN000010O struct {
}

func (*IN000010I) GetServiceKey() string {
	return "IN000010"
}
