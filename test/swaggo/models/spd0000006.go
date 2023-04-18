package models

type SPD0000006I struct {
	ProductType   string   `json:"productType" description:"product Type:Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan"` //*Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan
	ProductId     string   `json:"productId" description:"Product id"`
	ProductName   string   `json:"productName" description:"Product name"`
	SystemId      string   `json:"systemId" description:"System id" validate:"required"`
	ProductIdList []string `json:"productIdList" description:"Product id list"`
	Status        string   `json:"status" description:"Status:0-Inactive;1-Active"`
}

type SPD0000006O struct {
	ProductList []Product `json:"productList" description:"Product list"`
}

type Product struct {
	ProductId     string   `json:"productId" description:"Product id"`
	ProductName   string   `json:"productName" description:"Product name"`
	SystemId      string   `json:"systemId" description:"System id"`
	ProductType   string   `json:"productType" description:"product Type:Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan"` //*Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan
	Currency      string   `json:"currency" description:"Currency"`
	EffectiveDate string   `json:"effectiveDate" description:"Effective date"`
	ExpireDate    string   `json:"expireDate" description:"Expire date"`
	Status        string   `json:"status" description:"status:0-Inactive;1-Active"`
	ApproveView   string   `json:"approveView" description:"Approve Status(00-awaiting;01-approved;02-rejected)"`
	CustomerType  string   `json:"customerType" description:"customer Type:0-single;1-group"`
	ListCurrency  []string `json:"listCurrency" description:"List Currency"`
}

func (*SPD0000006I) GetServiceKey() string {
	return "pd000006"
}
