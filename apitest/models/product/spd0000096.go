package models

type PD000096I struct {
	ProductId      string `json:"productId" description:"Product Id"`
	ApproveStatus  string `json:"approveStatus" description:"Approve Status(0-awaiting;1-approved;2-rejected)"`
	CheckerComment string `json:"checkerComment" description:"Checker comment"`
}

type PD000096O struct {
	Status string `json:"status" description:"status(0-inactive;1-active;2-expired;3-deleted;4-updated)"`
}

func (*PD000096I) GetServiceKey() string {
	return "PD000096"
}
