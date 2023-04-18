package models

type SPD0000054I struct {
	CategoryId  string          `json:"categoryId" validate:"required" description:"Category Id"`
	ProductList []ProductIdInfo `json:"productList" description:"Product list" validate:"required,dive"`
}

type ProductIdInfo struct {
	ProductId string `json:"productId" description:"Product id" validate:"required"`
}

type SPD0000054O struct {
}

func (*SPD0000054I) GetServiceKey() string {
	return "PD000054"
}
