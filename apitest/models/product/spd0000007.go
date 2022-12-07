package models

import "github.com/shopspring/decimal"

type SPD0000007I struct {
	SystemId            string         `json:"systemId" description:"System id" validate:"required,max=2"`       //SV-Saving LN-Loan
	ProductType         string         `json:"productType" description:"Product type" validate:"required,max=2"` //*Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan
	ProductName         string         `json:"productName" description:"Product name" validate:"required"`
	EffectiveDate       string         `json:"effectiveDate" description:"Effective date" validate:"required"`
	ExpireDate          string         `json:"expireDate" description:"Expire date" validate:"required"`
	Currency            string         `json:"currency" description:"Currency" validate:"required"`
	Version             string         `json:"version" description:"Version" validate:"required"`
	Remark              string         `json:"remark" description:"Remark" validate:"required"`
	BalanceDirection    string         `json:"balanceDirection" description:"Balance direction" validate:"required"`        //D-Debit C-Credit
	CashCurrency        string         `json:"cashCurrency" description:"Cash currency" validate:"required"`                //C-Cash T-Currency
	AllowMaterialMedium string         `json:"allowMaterialMedium" description:"Allow material medium" validate:"required"` //0-no 1-yes
	MediumType          string         `json:"mediumType" description:"Medium type" validate:"required"`                    //0-debit card 1-credit card 2-passbook 3-checkbook
	WithdrawalMethod    string         `json:"withdrawalMethod" description:"Withdrawal method" validate:"required"`        //0- password1-seal2-signature3-fingerprint
	CrossBranch         string         `json:"crossBranch" description:"Cross branch" validate:"required"`                  //0-no 1-yes
	InterestFlag        string         `json:"interestFlag" description:"Interest flag" validate:"required"`                //0-no 1-yes
	ListInterest        []InterestInfo `json:"listInterest" description:"List interest" validate:"dive"`
	ListFee             []FeeInfo      `json:"listFee" description:"List fee" validate:"dive"`
}

type InterestInfo struct {
	InterestType string             `json:"interestType" description:"Interest type" validate:"required"` //0-normal 1-Withdrawal in Advance
	ListStrategy []InterestStrategy `json:"listStrategy" description:"List strategy" validate:"dive"`
}

type InterestStrategy struct {
	InterestStrategy string `json:"interestStrategy" description:"Interest strategy" validate:"required"`
}

type FeeInfo struct {
	FeeType               string            `json:"feeType" description:"Fee type" validate:"required"` //0-account opening
	MaxFee                decimal.Decimal   `json:"maxFee" description:"Max fee" validate:"required"`
	MinFee                decimal.Decimal   `json:"minFee" description:"Min fee" validate:"required"`
	DiscountMethod        string            `json:"discountMethod" description:"Discount method" validate:"required"`                //0-By Each Fee 1-By Total Fee
	DiscountGroupType     string            `json:"discountGroupType" description:"Discount group type" validate:"required"`         //0-Min Discount 1-Max Discount 2-Sum
	DiscountCalculateType string            `json:"discountCalculateType" description:"Discount calculate type" validate:"required"` //0-rate 1-fixed amount
	MaxDiscount           decimal.Decimal   `json:"maxDiscount" description:"Max discount" validate:"required"`
	MinDiscount           decimal.Decimal   `json:"minDiscount" description:"Min discount" validate:"required"`
	ListStrategy          []StrategyFeeInfo `json:"listStrategy" description:"List strategy" validate:"dive"`
}

type StrategyFeeInfo struct {
	FeeStrategy string `json:"feeStrategy" description:"Fee strategy"`
}

type SPD0000007O struct {
	ProductId string `json:"productId" description:"Product id"`
}

func (*SPD0000007I) GetServiceKey() string {
	return "spd0000007"
}
