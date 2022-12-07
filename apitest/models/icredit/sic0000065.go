package models

import "github.com/shopspring/decimal"

type IC000065Request struct {
	CustomerId             string          `json:"customerId"`
	CreditApplySerialNum   string          `json:"creditApplySerialNum"`
	NewPermanentQuotaTotal decimal.Decimal `json:"newPermanentQuotaTotal"`
	OperationType          string          `json:"operationType"` //- Create; 2- Increase; 3- Decrease; 4- Freeze; 5- Unfreeze; 6- Expire
}

type IC000065Response struct {
}
