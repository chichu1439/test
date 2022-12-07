package models

type SPD0000013I struct {
	SystemId            string  `json:"systemId" validate:"required" description:"system Id:LN-Loan;SV-Saving"`
	ProductId           string  `json:"productId" description:"Product id" validate:"required"`
	ProductType         string  `json:"productType" description:"product Type:Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan"` //*Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan
	ProductName         string  `json:"productName" description:"Product name"`
	Currency            string  `json:"currency" description:"Currency"`
	Version             string  `json:"version" description:"Version"`
	Remark              *string `json:"remark" description:"Remark"`
	EffectiveDate       string  `json:"effectiveDate" description:"Effective date"`
	ExpireDate          string  `json:"expireDate" description:"Expire date"`
	ValueDateMethod     string  `json:"valueDateMethod" description:"value Date Method:01-Approval Date 02-Account Opening Day 03-Loan Date 04-Loan Date plus Interest-free period"` //01-Approval Date 02-Account Opening Day 03-Loan Date 04-Loan Date plus Interest-free period
	MaturityDateMethod  string  `json:"maturityDateMethod" description:"maturity Date Method:01- Value Date and Loan Term;02- Same as the Last Interest Collection Date;03- Loan Date and Loan Term;04- Same as the Last Repayment Date"`
	RepaymentMethod     string  `json:"repaymentMethod" description:"repayment Method:01-Fixed Installment;02-Fixed Principal"`
	RepaymentFrequency  string  `json:"repaymentFrequency" description:"Repayment frequency"`
	AllowEarlyRepayment string  `json:"allowEarlyRepayment" validate:"required" description:"allow Early Repayment:0-no;1-yes"`
	AllowExtension      string  `json:"allowExtension" validate:"required" description:"allow Extension:0-no;1-yes"`
	AllowGracePeriod    string  `json:"allowGracePeriod" validate:"required" description:"allow Grace Period:0-no;1-yes"`
	GuaranteeMethod     string  `json:"guaranteeMethod" description:"guarantee Method:0-No Guarantee 1-Guarantee"` //0-No Guarantee 1-Guarantee
	Status              string  `json:"status" description:"status:0-inactive;1-active;2-expired;3-deleted;4-updated"`
}

type SPD0000013O struct {
}

func (*SPD0000013I) GetServiceKey() string {
	return "spd0000013"
}
