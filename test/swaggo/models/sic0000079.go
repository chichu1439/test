package models

type IC000079Request struct {
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
}

type IC000079Response struct {
	ExistModifyFlag        string       `json:"existModifyFlag"`
	EditLoanApplySerialNum string       `json:"editLoanApplySerialNum"`
	LoanReceiptNum         string       `json:"loanReceiptNum"`
	ModifyType             string       `json:"modifyType"`
	Maker                  string       `json:"maker"`
	ApplyStatus            string       `json:"applyStatus"`
	MakerComment           string       `json:"makerComment"`
	ApplyDate              string       `json:"applyDate"`
	ChangeList             []ChangeInfo `json:"changeList"`
}
type ChangeInfo struct {
	LoanElement               string `json:"loanElement"`
	ElementOldValue           string `json:"elementOldValue"`
	ElementNewValue           string `json:"elementNewValue"`
	LoanElementClassification string `json:"loanElementClassification"`
}

func (*IC000079Request) GetServiceKey() string {
	return "IC000079"
}
