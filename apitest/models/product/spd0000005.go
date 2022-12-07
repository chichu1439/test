package models

type PD000005I struct {
	CategoryId string `json:"categoryId" description:"Category id" validate:"required,max=10"`
}

type PD000005O struct {
	ProductList []ProductInfo `json:"productList" description:"Product list"`
}

type ProductInfo struct {
	ProductId   string `json:"productId" description:"Product id"`
	ProductName string `json:"productName" description:"Product name"`
	SystemId    string `json:"systemId" description:"System id"`                                                                   //SV-Saving LN-Loan
	ProductType string `json:"productType" description:"product Type:Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan"` //*Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan
	Currency    string `json:"currency" description:"Currency"`
}

func (*PD000005I) GetServiceKey() string {
	return "PD000005"
}
