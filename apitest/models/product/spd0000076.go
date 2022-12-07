package models

import "github.com/shopspring/decimal"

type SPD0000076I struct {
	TransDate string `json:"transDate" description:"Trans date"`
}

type SPD0000076O struct {
	ListMicroLoanProduct []MicroLoanProduct `json:"listMicroLoanProduct" description:"List micro loan product"`
}

type MicroLoanProduct struct {
	ProductId           string          `json:"productId" description:"Product id"`
	ValueDateMethod     string          `json:"valueDateMethod" description:"Value date method"`
	MaturityDateMethod  string          `json:"maturityDateMethod" description:"Maturity date method"`
	RepaymentMethod     string          `json:"repaymentMethod" description:"Repayment method"`
	RepaymentFrequency  string          `json:"repaymentFrequency" description:"Repayment frequency"`
	AllowEarlyRepayment string          `json:"allowEarlyRepayment" description:"Allow early repayment"`
	AllowExtension      string          `json:"allowExtension" description:"Allow extension"`
	AllowGracePeriod    string          `json:"allowGracePeriod" description:"Allow grace period"`
	GuaranteeMethod     string          `json:"guaranteeMethod" description:"Guarantee method"`
	EffectiveDate       string          `json:"effectiveDate" description:"Effective date"`
	ExpireDate          string          `json:"expireDate" description:"Expire date"`
	Status              string          `json:"status" description:"Status"`
	ProductName         string          `json:"productName" description:"Product name"`
	SystemId            string          `json:"systemId" description:"System id"`
	ProductType         string          `json:"productType" description:"Product type"`
	Currency            string          `json:"currency" description:"Currency"`
	Version             string          `json:"version" description:"Version"`
	Remark              string          `json:"remark" description:"Remark"`
	ServiceId           string          `json:"serviceId" description:"Service id"`
	LimitTarget         string          `json:"limitTarget" description:"Limit target"`
	LimitMethod         string          `json:"limitMethod" description:"Limit method"`
	ControlMethod       string          `json:"controlMethod" description:"Control method"`
	UpperLimitControl   string          `json:"upperLimitControl" description:"Upper limit control"`
	UpperLimitValue     decimal.Decimal `json:"upperLimitValue" description:"Upper limit value"`
	LowerLimitControl   string          `json:"lowerLimitControl" description:"Lower limit control"`
	LowerLimitValue     decimal.Decimal `json:"lowerLimitValue" description:"Lower limit value"`
	LimitStrategyId     string          `json:"limitStrategyId" description:"Limit strategy id"`
	InterestType        string          `json:"interestType" description:"Interest type"`
	InterestId          string          `json:"interestId" description:"Interest id"`
	MinInterestRate     decimal.Decimal `json:"minInterestRate" description:"Min interest rate"`
	MaxInterestRate     decimal.Decimal `json:"maxInterestRate" description:"Max interest rate"`
	MinFloatValue       decimal.Decimal `json:"minFloatValue" description:"Min float value"`
	MaxFloatValue       decimal.Decimal `json:"maxFloatValue" description:"Max float value"`
}

func (*SPD0000076I) GetServiceKey() string {
	return "PD000076"
}
