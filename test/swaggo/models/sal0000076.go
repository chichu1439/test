//Version: v1
package models

type AL000076Request struct {
	ApplySerialNum       string `json:"applySerialNum" validate:"required"`
	AccountingAccountNum string `json:"accountingAccountNum" validate:"required"`
}

type AL000076Response struct {
}

func (*AL000076Request) GetServiceKey() string {
	return "AL000076"
}
