package models

type PD000004I struct {
	TreeIndex string `json:"treeIndex" description:"Tree index" validate:"required"`
}

type PD000004O struct {
}

func (*PD000004I) GetServiceKey() string {
	return "PD000004"
}
