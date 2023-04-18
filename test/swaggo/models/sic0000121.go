//Version: v1
package models

type IC000121Request struct {
	LoanReceiptNum string `json:"loanReceiptNum" description:"loan receipt num" validate:"required"`
}

type IC000121Response struct {
	ApplySerialNum string `json:"applySerialNum" description:"apply serial num"`
	Status         string `json:"status" description:"status,1 - pending 2 - approved 3 - rejecte"`
	LoanReceiptNum string `json:"loanReceiptNum" description:"loan receipt num" `
}

func (*IC000121Request) GetServiceKey() string {
	return "IC000121"
}
