package models

type SIC0000028I struct {
	ModifyType        string                  `json:"modifyType" validate:"required" description:"Modify type:2-Interest Rate;3-Period Distribution,4-Loan Extension,5- Bill rule"` //1-Loan 2-Transaction Account 3-Marital Detail 4-Contact 5-Guarantor 6-Fee 7-emergencyInfo
	LoanReceiptNum    string                  `json:"loanReceiptNum" validate:"required" description:"Loan receipt number"`
	CustomerId        string                  `json:"customerId" validate:"required" description:"Customer id"`
	MakerComment      string                  `json:"makerComment" validate:"required" description:"Maker comment"`
	ChangeContextList []LoanChangeContextList `json:"changeContextList" description:"Change context list" validate:"required"`
}

type LoanChangeContextList struct {
	ChangeItem  string `json:"changeItem" description:"Change item"`
	BeforeValue string `json:"beforeValue" description:"Before value"`
	AfterValue  string `json:"afterValue" description:"After value"`
}

type SIC0000028O struct {
	LoanReceiptNum         string `json:"loanReceiptNum" description:"Loan receipt number"`
	EditLoanApplySerialNum string `json:"editLoanApplySerialNum" `
}

func (*SIC0000028I) GetServiceKey() string {
	return "IC000028"
}
