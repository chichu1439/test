package models

type IC000030I struct {
	//"1001"-Account Statement "1002"-Transaction History "1003"-repayment plan "1004"-Late Payment Notification
	OprateType     string `json:"oprateType" validate:"required" description:"Oprate type:1001-Account Statement 1002-Transaction History 1003-repayment plan 1004-Late Payment Notification"`
	CustomerId     string `json:"customerId" description:"Customer id" validate:"required"`
	LoanReceiptNum string `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	MailFlag       string `json:"mailFlag" validate:"required" description:"Mail flag:Y-send mail N-no mail"` //Y- send mail N- no mail
}

type IC000030O struct {
	FileKey string `json:"fileKey" description:"File key"`
	Data    map[string]interface{}
}

func (*IC000030I) GetServiceKey() string {
	return "IC000030"
}
