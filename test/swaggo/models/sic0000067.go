//Version: v1
package models

type IC000067Request struct {
	LoanReceiptNum  string `json:"loanReceiptNum"`
	GlobalSerialNum string `json:"globalSerialNum" description:"global serial num"`
	ApplySerialNum  string `json:"applySerialNum"`
	CheckFlag       string `json:"checkFlag"`
}

type IC000067Response struct {
	MakerComment        string `json:"makerComment"`
	Maker               string `json:"maker"`
	IsPending           string `json:"isPending"` //Y N
	ReasonForCancelling string `json:"reasonForCancelling"`
	BusinessSerialNum   string `json:"businessSerialNum"`
	ApplySerialNum      string `json:"applySerialNum"`
}
