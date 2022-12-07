package models

type PD000003I struct {
	CategoryId       string `json:"categoryId" description:"Category id" validate:"required,max=10"`
	CategoryName     string `json:"categoryName" description:"Category name" validate:"required,max=100"`
	IsLeaf           string `json:"isLeaf" validate:"required,max=1" description:"is Leaf flag:0-no ,1-yes"` //0-no ,1-yes
	ParentCategoryId string `json:"parentCategoryId" description:"Parent category id"`
}

type PD000003O struct {
}

func (*PD000003I) GetServiceKey() string {
	return "PD000003"
}
