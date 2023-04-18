package models

type SPD0000019I struct {
	ProductId     string `json:"productId" description:"Product id"`
	ProductName   string `json:"productName" description:"Product name"`
	SystemId      string `json:"systemId" description:"system Id:LN-Loan;SV-Saving"`
	ProductType   string `json:"productType" description:"product Type:Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan"` //*Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan
	Currency      string `json:"currency" description:"Currency"`
	Version       string `json:"version" description:"Version"`
	Remark        string `json:"remark" description:"Remark"`
	EffectiveDate string `json:"effectiveDate" description:"Effective date"`
	ExpireDate    string `json:"expireDate" description:"Expire date"`
	Status        string `json:"status" description:"0-Normal"`
}

type SPD0000019O struct {
}

func (*SPD0000019I) GetServiceKey() string {
	return "PD000019"
}
