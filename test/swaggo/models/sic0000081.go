//Version: v1
package models

type IC000081Request struct {
	LoanReceiptNum string   `json:"loanReceiptNum" description:"Loan receipt number" validate:"required"`
	TransType      string   `json:"transType" description:" trans type : 03-RPYM， 04 - DISB"  validate:"required"`
	Behavior       string   `json:"behavior" description:"behavior,01-repayment ,02-close repayment"`
	ApproveView    []string `json:"approveView" description:" apply status: 1 - pending 2 - approved 3 - rejected "`
}

type IC000081Response struct {
	ListRegisterRecord []RegisterRecord `json:"listRegisterRecord" description:" Register records"`
}
type RegisterRecord struct {
	ApplyDate                 string `json:"applyDate"  description:" apply date"`
	ApplySerialNum            string `json:"applySerialNum" description:" apply serial num"`
	LoanReceiptNum            string `json:"loanReceiptNum" description:"loan receipt num "`
	CustomerId                string `json:"customerId" description:"customer Id "`
	ProductId                 string `json:"productId" description:"product id "`
	TransDate                 string `json:"transDate" description:"trans date "`
	TransTime                 string `json:"transTime" description:" trans time"`
	TransAmount               string `json:"transAmount"  description:"trans amount "`
	TransPrinciple            string `json:"transPrinciple" description:"trans principle "`
	TransInterest             string `json:"transInterest" description:"trans interest "`
	NormalInterest            string `json:"normalInterest" description:"normal interest"`
	PenaltyInterest           string `json:"penaltyInterest" description:"penalty interest"`
	CompoundInterest          string `json:"compoundInterest" description:"compound interest"`
	PenaltyCompoundInterest   string `json:"penaltyCompoundInterest" description:"penalty compound interest"`
	LoanChannel               string `json:"loanChannel" description:" loan channel"`
	BankName                  string `json:"bankName" description:" bank name"`
	ChannelName               string `json:"channelName" description:"channel name "`
	CounterpartyAccountName   string `json:"counterpartyAccountName" description:"Counterparty Account Name "`
	CounterpartyAccountNum    string `json:"counterpartyAccountNum" description:"Counterparty Account Num "`
	CounterpartyBankName      string `json:"counterpartyBankName" description:" Counterparty Bank Name"`
	CounterpartyBankNum       string `json:"counterpartyBankNum" description:" Counterparty Bank Num"`
	CounterpartyLocalBankFlag string `json:"counterpartyLocalBankFlag" description:"Counterparty Local Bank Flag "`
	TransType                 string `json:"transType" description:" Trans Type"`
	TransDescribe             string `json:"transDescribe" description:" Trans Describe"`
	Maker                     string `json:"makerUserId" description:" Maker User Id "`
	MakerComment              string `json:"makerComment" description:"Maker Comment "`
	Checker                   string `json:"checker" description:"Checker User Id "`
	ApproveComment            string `json:"approveComment" description:"Checker Comment "`
	ApproveView               string `json:"approveView" description:" Apply Status"`
	LoanStatus                string `json:"loanStatus" description:"loan status "`
	DiscountPercentage        string `json:"discountPercentage" description:"Discount percentage"`
}
