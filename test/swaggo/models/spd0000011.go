package models

import "github.com/shopspring/decimal"

type SPD0000011I struct {
	SystemId            string             `json:"systemId" validate:"required,max=2" description:"system Id:LN-Loan;SV-Saving"`                                                 //SV-Saving LN-Loan
	ProductType         string             `json:"productType" validate:"required,max=2" description:"product Type:Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan"` //*Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan
	ProductName         string             `json:"productName" description:"Product name" validate:"required"`
	Currency            string             `json:"currency" description:"Currency" validate:"required"`
	Version             string             `json:"version" description:"Version" validate:"required"`
	Remark              string             `json:"remark" description:"Remark"`
	EffectiveDate       string             `json:"effectiveDate" description:"Effective date" validate:"required"`
	ExpireDate          string             `json:"expireDate" description:"Expire date" validate:"required"`
	ValueDateMethod     string             `json:"valueDateMethod" description:"value Date Method:01-Approval Date 02-Account Opening Day 03-Loan Date 04-Loan Date plus Interest-free period"` //01-Approval Date 02-Account Opening Day 03-Loan Date 04-Loan Date plus Interest-free period
	MaturityDateMethod  string             `json:"maturityDateMethod" description:"maturity Date Method:01- Value Date and Loan Term;02- Same as the Last Interest Collection Date;03- Loan Date and Loan Term;04- Same as the Last Repayment Date"`
	RepaymentMethod     string             `json:"repaymentMethod" description:"repayment Method:01-Fixed Installment;02-Fixed Principal"`
	RepaymentFrequency  string             `json:"repaymentFrequency" description:"Repayment frequency"`
	AllowEarlyRepayment string             `json:"allowEarlyRepayment" validate:"required" description:"allow Early Repayment:0-no;1-yes"`
	AllowExtension      string             `json:"allowExtension" validate:"required" description:"allow Extension:0-no;1-yes"`
	AllowGracePeriod    string             `json:"allowGracePeriod" validate:"required" description:"allow Grace Period:0-no;1-yes"`
	GuaranteeMethod     string             `json:"guaranteeMethod" description:"guarantee Method:0-No Guarantee 1-Guarantee"` //0-No Guarantee 1-Guarantee
	ListInterest        []InterestLoanInfo `json:"listInterest" description:"interest List " validate:"dive"`
	ListFee             []FeeLoanInfo      `json:"listFee" description:"List fee" validate:"dive"`
}

type InterestLoanInfo struct {
	InterestType string             `json:"interestType" validate:"required" description:"interest Type:0-normal 1-Withdrawal in Advance"` //0-normal 1-Withdrawal in Advance
	ListStrategy []InterestStrategy `json:"listStrategy" description:"List strategy" validate:"required"`
}

type FeeLoanInfo struct {
	FeeType               string             `json:"feeType" validate:"required" description:"fee Type:0-account opening"` //0-account opening
	MaxFee                decimal.Decimal    `json:"maxFee" description:"Max fee" validate:"required"`
	MinFee                decimal.Decimal    `json:"minFee" description:"Min fee" validate:"required"`
	DiscountMethod        string             `json:"discountMethod" validate:"required" description:"discount Method:0-By Each Fee 1-By Total Fee"`               //0-By Each Fee 1-By Total Fee
	DiscountGroupType     string             `json:"discountGroupType" validate:"required" description:"discount Group Type:0-Min Discount 1-Max Discount 2-Sum"` //0-Min Discount 1-Max Discount 2-Sum
	DiscountCalculateType string             `json:"discountCalculateType" validate:"required" description:"discount Calculate Type:0-rate 1-fixed amount"`       //0-rate 1-fixed amount
	MaxDiscount           decimal.Decimal    `json:"maxDiscount" description:"Max discount" validate:"max Discount:required"`
	MinDiscount           decimal.Decimal    `json:"minDiscount" description:"Min discount" validate:"min Discount:required"`
	ListStrategy          []StrategyLoanInfo `json:"listStrategy" description:"List strategy" validate:"list Strategy:required,dive"`
}

type StrategyLoanInfo struct {
	FeeStrategy string `json:"feeStrategy" description:"Fee strategy" validate:"required"`
}

type SPD0000011O struct {
	ProductId string `json:"productId" description:"Product id"`
}

func (*SPD0000011I) GetServiceKey() string {
	return "PD000011"
}
