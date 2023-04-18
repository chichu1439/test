package models

type PD000001I struct {
	CategoryName     string `json:"categoryName" description:"Category name" validate:"required,max=100"`
	IsLeaf           string `json:"isLeaf" validate:"required,max=1" description:"Is Leaf flag:0-no ,1-yes"` //0-no ,1-yes
	ParentCategoryId string `json:"parentCategoryId" description:"Parent category id"`
}

type PD000001O struct {
	CategoryId string `json:"categoryId" description:"Category id"`
}

func (*PD000001I) GetServiceKey() string {
	return "PD000001"
}
