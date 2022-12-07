package models

type PD000002I struct {
	TreeIndex string `json:"treeIndex" description:"Tree index" validate:"required"`
}

type PD000002O struct {
	ListCategory []ListCategory `json:"listCategory" description:"List category"`
}

type ListCategory struct {
	CategoryId    string `json:"categoryId" description:"Category id"`
	CategoryName  string `json:"categoryName" description:"Category name"`
	IsLeaf        string `json:"isLeaf" description:"is Leaf flag:0-no ,1-yes"`
	CategoryLevel int    `json:"categoryLevel" description:"Category level"`
	TreeIndex     string `json:"treeIndex" description:"Tree index"`
}

func (*PD000002I) GetServiceKey() string {
	return "PD000002"
}
