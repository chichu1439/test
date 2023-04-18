package models

type IC000011I struct {
	CustomerId     string `json:"customerId" description:"Customer id" validate:"required"`
	ApproveStatus  string `json:"approveStatus" validate:"required" description:"Approve status:1-to be approved 2-approved 3-rejected"` //1-待审批 2-审批通过 3-审批拒绝
	ApproveComment string `json:"approveComment" description:"Approve Comment"`
}

type IC000011O struct {
	CustomerId string `json:"customerId" description:"Customer id"`
}

func (*IC000011I) GetServiceKey() string {
	return "ic000011"
}
