//Version: v1
package models

import "github.com/shopspring/decimal"

type IC000067Request struct {
	CustomerId                string `json:"customerId"`
	ProductId                 string `json:"productId"`
	PageNum                   int    `json:"pageNum" description:"Page number"`
	PageRecordCount           int    `json:"pageRecordCount" description:"Page record count"`
	CreditLineReferenceNumber string `json:"creditLineReferenceNumber"`
}

type IC000067Response struct {
	PageNum         int                    `json:"pageNum" description:"Page number"`
	PageRecordCount int                    `json:"pageRecordCount" description:"Page record count"`
	TotalCount      int                    `json:"totalCount" description:"Total count"`
	Records         []IC000067ResponseItem `json:"records"`
}
type IC000067ResponseItem struct {
	TransactionDate     string          `json:"transactionDate"` // t_quota_history.final_modify_date
	TransactionTime     string          `json:"transactionTime"` // t_quota_history.final_modify_time
	TransactionType     string          `json:"transactionType"` // 01 Loan/Disbursement Application;02 Loan/Disbursement Rejection;03 Online Disbursement;04 Batch Disbursement;05 Online Repayment;06 Batch Repayment;07 Batch Repay-off Total;08 Write off'
	CreditLine          decimal.Decimal `json:"creditLine"`
	UsedCreditLine      decimal.Decimal `json:"usedCreditLine"`
	FrozenCreditLine    decimal.Decimal `json:"frozenCreditLine"`
	AvailableCreditLine decimal.Decimal `json:"availableCreditLine"`
	UserId              string          `json:"userId"`
}
