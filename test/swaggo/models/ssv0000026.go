//Version: v0.0.1
package models

type SV000026I struct {
	DepositProofNumber string `json:"depositProofNumber" validate:"required"`
	CustomerID         string `json:"customerId" validate:"required"`
}

type SV000026O struct {
	Status string `json:"status"`
}

func (*SV000026I) GetServiceKey() string {
	return "SV000026"
}
