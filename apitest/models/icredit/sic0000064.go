package models

import "github.com/shopspring/decimal"

type IC000064Request struct {
	CustomerId           string `json:"customerId" description:"customer id" validate:"required"`
	CreditApplySerialNum string `json:"creditApplySerialNum" description:"credit Application Serial Number" validate:"required"`
	PageNum              int    `json:"pageNum" description:"page Number of current request" validate:"required"`
	PageSize             int    `json:"PageSize" description:"number of records of each page" validate:"required"`
}

type IC000064Response struct {
	PageNum         int               `json:"pageNum" description:"page Number of current request"`
	PageRecordCount int               `json:"pageRecordCount" description:"the records of current page"`
	TotalCount      int               `json:"totalCount" description:"total records"`
	Records         []QuotaChangeinfo `json:"records" description:"Quota Change info list"`
}

type QuotaChangeinfo struct {
	OperationDateTime   string          `json:"operationDateTime" description:"Operation Date Time"`
	OperationType       string          `json:"operationType" description:"Operation Type"`
	UpdatedStatus       string          `json:"updatedStatus" description:"Updated Status"`
	CreditLine          decimal.Decimal `json:"creditLine" description:"Credit Line"`
	UsedCreditLine      decimal.Decimal `json:"usedCreditLine" description:"Used Credit Line"`
	AvailableCreditLine decimal.Decimal `json:"availableCreditLine" description:"Available Credit Line"`
	FrozenCreditLine    decimal.Decimal `json:"frozenCreditLine" description:"Frozen credit line" `
	UserID              string          `json:"userID" description:"User ID"`
}

func (*IC000064Request) GetServiceKey() string {
	return "IC000064"
}
