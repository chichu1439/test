package models

type IC000133I struct {
	LoanReceiptNum      string `json:"loanReceiptNum"`
	ReasonForCancelling string `json:"reasonForCancelling"` // 0-Operation Error; 1-Borrower Applies for Transaction Cancellation; 2-Other Reasons;
	GlobalSerialNum     string `json:"globalSerialNum" description:"global serial num"`
	BusinessSerialNum   string `json:"businessSerialNum" description:"business serial num"`
	MakerComment        string `json:"makerComment"`
}

type IC000133O struct {
	LoanReceiptNum string `json:"loanReceiptNum"`
	TransSerialNum string `json:"transSerialNum"`
}

func (*IC000133I) GetServiceKey() string {
	return "IC000133"
}
