package models

type SPD0000009I struct {
	SystemId            string `json:"systemId" description:"System id:SV-Saving LN-Loan" validate:"required,max=2"` //SV-Saving LN-Loan
	ProductId           string `json:"productId" description:"Product id" validate:"required"`
	ProductType         string `json:"productType" description:"Product type" validate:"max=2"` //*Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan
	ProductName         string `json:"productName" description:"Product name" `
	Currency            string `json:"currency" description:"Currency" `
	Version             string `json:"version" description:"Version" `
	Remark              string `json:"remark" description:"Remark"`
	EffectiveDate       string `json:"effectiveDate" description:"Effective date"`
	ExpireDate          string `json:"expireDate" description:"Expire date"`
	Status              string `json:"status" description:"Status"`
	BalanceDirection    string `json:"balanceDirection" description:"Balance direction:D-Debit C-Credit" ` //D-Debit C-Credit
	CashCurrency        string `json:"cashCurrency" description:"Cash currency:C-Cash T-Currency"`         //C-Cash T-Currency
	AllowMaterialMedium string `json:"allowMaterialMedium" description:"Allow material medium:0-no 1-yes"` //0-no 1-yes
	MediumType          string `json:"mediumType" description:"Medium type" `                              //0-debit card 1-credit card 2-passbook 3-checkbook
	WithdrawalMethod    string `json:"withdrawalMethod" description:"Withdrawal method" `                  //0- password1 - seal2 - signature3 - fingerprint
	CrossBranch         string `json:"crossBranch" description:"Cross branch"`                             //0-no 1-yes
	InterestFlag        string `json:"interestFlag" description:"Interest flag"`                           //0-no 1-yes
}

type SPD0000009O struct {
}

func (*SPD0000009I) GetServiceKey() string {
	return "spd0000009"
}
