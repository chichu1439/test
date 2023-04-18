//Version: v1
package models

type AL000057Request struct {
	AccountingAccountNum string `json:"accountingAccountNum" description:"Accounting account number" validate:"required"`
}

type AL000057Response struct {
}

func (*AL000057Request) GetServiceKey() string {
	return "AL000057"
}
