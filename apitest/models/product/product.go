package models

import "github.com/shopspring/decimal"

type TProdLimitControlBo struct {
	Uid               int    `json:"uid" description:"Uid"`
	ServiceId         string `json:"serviceId" description:"Service id"` //
	FeatureId         string `json:"featureId" description:"Feature id"`
	LimitTarget       string `json:"limitTarget" description:"Limit target"`              // 0-Balance 1-Trans Times 2-Trans Amount 3-Free Fee Times 4-Free Fee Amount 5-Free Interest Times 6-Free Interest Amount 7-Free Tax Times 8-Free Tax Amount ",
	LimitMethod       string `json:"limitMethod" description:"Limit method"`              // 0-Direct Limit 1-By Strategy ",
	ControlMethod     string `json:"controlMethod" description:"Control method"`          // 0-Pass 1-Reject 2-By Strategy ",
	UpperLimitControl string `json:"upperLimitControl" description:"Upper limit control"` // 0-Reject 1-Prompt 2-Record Only ",
	UpperLimitValue   int    `json:"upperLimitValue" description:"Upper limit value"`     //
	LowerLimitControl string `json:"lowerLimitControl" description:"Lower limit control"` // 0-Reject 1-Prompt 2-Record Only ",
	LowerLimitValue   int    `json:"lowerLimitValue" description:"Lower limit value"`     //
	LimitStrategyId   string `json:"limitStrategyId" description:"Limit strategy id"`     //
	EffectiveDate     string `json:"effectiveDate" description:"Effective date"`          //
	ExpireDate        string `json:"expireDate" description:"Expire date"`                //
}

type TProdServicePoolBo struct {
	ServiceId       string `json:"serviceId" description:"Service id"`             //
	ServiceName     string `json:"serviceName" description:"Service name"`         // 01-Approval Date 02-Account Opening Day 03-Loan Date 04-Loan Date plus Interest-free period ",
	SystemId        string `json:"systemId" description:"System id"`               // SV-Saving LN-Loan ",
	ProductType     string `json:"productType" description:"Product type"`         // *Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan ",
	CustomerType    string `json:"customerType" description:"Customer type"`       // 0-Personal 1-Corporate ",
	ServiceType     string `json:"serviceType" description:"Service type"`         // 0-Opening 1-Closing 2-Management 3-Query 4-Finance Trans 5-Notification ",
	TransactionType string `json:"transactionType" description:"Transaction type"` // 0-Online 1-Batch ",
	EffectiveDate   string `json:"effectiveDate" description:"Effective date"`     //
	ExpireDate      string `json:"expireDate" description:"Expire date"`           //
}

type InterestInfoBo struct {
	InterestType     string          `json:"interestType" description:"Interest type:0-Normal Loan 1-Overdue Loan 2-Impaired Loan"` //0-Normal Loan 1-Overdue Loan 2-Impaired Loan
	InterestId       string          `json:"interestId" description:"Interest id"`
	StrategyId       string          `json:"strategyId" description:"Strategy id"`
	BaseInterestRate decimal.Decimal `json:"baseInterestRate" description:"Base interest rate"`
	MinInterestRate  decimal.Decimal `json:"minInterestRate" description:"Min interest rate"`
	MaxInterestRate  decimal.Decimal `json:"maxInterestRate" description:"Max interest rate"`
	FloatType        string          `json:"floatType" description:"Float type:0-By Percent 1-BP Value"`             //0-By Percent 1-BP Value
	FloatDirection   string          `json:"floatDirection" description:"Float direction:0-None 1-Up 2-Down 3-Both"` //0-None 1-Up 2-Down 3-Both
	MinFloatValue    decimal.Decimal `json:"minFloatValue" description:"Min float value"`
	MaxFloatValue    decimal.Decimal `json:"maxFloatValue" description:"Max float value"`
}

type ProdCategoryNameBo struct {
	CategoryId   string `json:"categoryId" description:"Category id"`
	CategoryName string `json:"categoryName" description:"Category name"`
}
