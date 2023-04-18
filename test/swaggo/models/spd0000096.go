package models

type PD000096I struct {
	ProductId      string `json:"productId" description:"Product Id"`
	ApproveView    string `json:"approveView" description:"approve view (00-awaiting;01-agree;02-rejected)"`
	ApproveComment string `json:"approveComment" description:"approve comment"`
}

type PD000096O struct {
	Status    string `json:"status" description:"status(0-inactive;1-active;2-expired;3-deleted;4-updated)"`
	ProductId string `json:"productId" description:"Product Id"`
}

func (*PD000096I) GetServiceKey() string {
	return "PD000096"
}
