package models

type SPD0000052I struct {
	ProductId string `json:"productId" description:"Product id" validate:"required"`
}

type SPD0000052O struct {
	SystemId            string         `json:"systemId" description:"system Id:SV-Saving LN-Loan"`                                                 //SV-Saving LN-Loan
	ProductType         string         `json:"productType" description:"product Type:Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan"` //*Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan
	ProductName         string         `json:"productName" description:"Product name"`
	EffectiveDate       string         `json:"effectiveDate" description:"Effective date"`
	ExpireDate          string         `json:"expireDate" description:"Expire date"`
	Currency            string         `json:"currency" description:"Currency"`
	Version             string         `json:"version" description:"Version"`
	Remark              string         `json:"remark" description:"Remark"`
	BalanceDirection    string         `json:"balanceDirection" description:"balance Direction:D-Debit C-Credit"`                                //D-Debit C-Credit
	CashCurrency        string         `json:"cashCurrency" description:"cashCurrency flag:C-Cash T-Currency"`                                   //C-Cash T-Currency
	AllowMaterialMedium string         `json:"allowMaterialMedium" description:"allowMaterialMedium:0-no 1-yes"`                                 //0-no 1-yes
	MediumType          string         `json:"mediumType" description:"medium Type:0-debit card 1-credit card 2-passbook 3-checkbook"`           //0-debit card 1-credit card 2-passbook 3-checkbook
	WithdrawalMethod    string         `json:"withdrawalMethod" description:"withdrawal Method:0- password1 - seal2 - signature3 - fingerprint"` //0- password1 - seal2 - signature3 - fingerprint
	CrossBranch         string         `json:"crossBranch" description:"cross Branch:0-no 1-yes"`                                                //0-no 1-yes
	InterestFlag        string         `json:"interestFlag" description:"interest Flag:0-no 1-yes"`                                              //0-no 1-yes
	ListInterest        []InterestInfo `json:"listInterest" description:"List interest"`
	ListFee             []FeeInfo      `json:"listFee" description:"List fee"`
}

func (*SPD0000052I) GetServiceKey() string {
	return "spd0000052"
}
