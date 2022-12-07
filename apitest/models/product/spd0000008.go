package models

type SPD0000008I struct {
	ProductId string `json:"productId" description:"Product id"`
}

type SPD0000008O struct {
	ProductId           string `json:"productId" description:"Product id"`
	BalanceDirection    string `json:"balanceDirection" description:"Balance direction"`
	CashCurrency        string `json:"cashCurrency" description:"Cash currency"`
	AllowMaterialMedium string `json:"allowMaterialMedium" description:"Allow material medium"`
	MediumType          string `json:"mediumType" description:"Medium type"`
	WithdrawalMethod    string `json:"withdrawalMethod" description:"Withdrawal method"`
	CrossBranch         string `json:"crossBranch" description:"Cross branch"`
	InterestFlag        string `json:"interestFlag" description:"Interest flag"`
	EffectiveDate       string `json:"effectiveDate" description:"Effective date"`
	ExpireDate          string `json:"expireDate" description:"Expire date"`
	Status              string `json:"status" description:"Status"`
	ProductName         string `json:"productName" description:"Product name"`
	SystemId            string `json:"systemId" description:"System id"`
	ProductType         string `json:"productType" description:"Product type"`
	Currency            string `json:"currency" description:"Currency"`
	Version             string `json:"version" description:"Version"`
	Remark              string `json:"remark" description:"Remark"`
}

func (*SPD0000008I) GetServiceKey() string {
	return "spd0000008"
}
