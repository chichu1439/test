package models

import "github.com/shopspring/decimal"

type IC000052I struct {
	TranDate string `json:"tranDate" description:"Tran date"` //业务日期
}

type IC000052O struct {
	FileKey string `json:"fileKey" description:"File key"`
}

type CurrentCreditQuotaBo struct {
	AccountingAccountNum string          `json:"accountingAccountNum" description:"Accounting account number"` //核算账号
	LoanReceiptNum       string          `json:"loanReceiptNum" description:"Loan receipt number"`             //借据号
	CurrentCreditQuota   decimal.Decimal `json:"currentCreditQuota" description:"Current credit quota"`        //当前授信额度

}

type CurrentCreditQuotaStatisticsBo struct {
	Num1     int             `json:"Num1" description:"Num1"`         //额度<= 5,000数量
	Num2     int             `json:"Num2" description:"Num2"`         //额度5,000.01-10,000.00数量
	Num3     int             `json:"Num3" description:"Num3"`         //额度10,000.01-15,000.00数量
	Num4     int             `json:"Num4" description:"Num4"`         //额度15,000.01-20,000.00数量
	Num5     int             `json:"Num5" description:"Num5"`         //额度20,000.01-25,000.00数量
	Num6     int             `json:"Num6" description:"Num6"`         //额度25,000.01-30,000.00数量
	Num7     int             `json:"Num7" description:"Num7"`         //额度30,000.01-50,000.00数量
	Num8     int             `json:"Num8" description:"Num8"`         //额度50,000.01-75,000.00数量
	Num9     int             `json:"Num9" description:"Num9"`         //额度75,000.01-100,000.00数量
	Balance1 decimal.Decimal `json:"Balance1" description:"Balance1"` //额度<= 5,000的记录的额度之和
	Balance2 decimal.Decimal `json:"Balance2" description:"Balance2"` //额度5,000.01-10,000.00的记录的额度之和
	Balance3 decimal.Decimal `json:"Balance3" description:"Balance3"` //额度10,000.01-15,000.00的记录的额度之和
	Balance4 decimal.Decimal `json:"Balance4" description:"Balance4"` //额度15,000.01-20,000.00的记录的额度之和
	Balance5 decimal.Decimal `json:"Balance5" description:"Balance5"` //额度20,000.01-25,000.00的记录的额度之和
	Balance6 decimal.Decimal `json:"Balance6" description:"Balance6"` //额度25,000.01-30,000.00的记录的额度之和
	Balance7 decimal.Decimal `json:"Balance7" description:"Balance7"` //额度30,000.01-50,000.00的记录的额度之和
	Balance8 decimal.Decimal `json:"Balance8" description:"Balance8"` //额度50,000.01-75,000.00的记录的额度之和
	Balance9 decimal.Decimal `json:"Balance9" description:"Balance9"` //额度75,000.01-100,000.00的记录的额度之和
}

func (*IC000052I) GetServiceKey() string {
	return "IC000052"
}
