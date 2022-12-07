package models

type SPD0000020I struct {
	ProductId     string `json:"productId" description:"Product Id"`
	ProductName   string `json:"productName" description:"Product Name"`
	SystemId      string `json:"systemId" description:"system Id:LN-Loan;SV-Saving"`
	ProductType   string `json:"productType" description:"product Type:Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan"` //*Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan
	Currency      string `json:"currency" description:"Currency"`
	Version       string `json:"version" description:"Version"`
	Remark        string `json:"remark" description:"Remark"`
	EffectiveDate string `json:"effectiveDate" description:"Effective date"`
	ExpireDate    string `json:"expireDate" description:"Expire date"`
	Status        string `json:"status" description:"status:0-Normal"`
}

type SPD0000020O struct {
}

func (*SPD0000020I) GetServiceKey() string {
	return "spd0000020"
}
