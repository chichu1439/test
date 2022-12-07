package models

type IC000050I struct {
	TranDate string `json:"TranDate" description:"Tran date"` //业务日期
}

type IC000050O struct {
	FileKey string `json:"fileKey" description:"File key"`
}

type MonthlyClosedLoanReport struct {
	SeqNo                string `json:"seqNo" description:"Seq no"`
	CIF                  string `json:"cif" description:"Cif"`
	CustomerName         string `json:"customerName" description:"Customer name"`
	LoanReferenceNumber  string `json:"loanReferenceNumber" description:"Loan reference number"`
	LoanStatus           string `json:"loanStatus" description:"Loan status"`
	IDType               string `json:"iDType" description:"Id type"`
	CertificateID        string `json:"certificateID" description:"Certificate id"`
	ProductID            string `json:"productID" description:"Product id"`
	ProductName          string `json:"productName" description:"Product name"`
	DrawdownDate         string `json:"drawdownDate" description:"Drawdown date"`
	MaturityDate         string `json:"maturityDate" description:"Maturity date"`
	ClosedDate           string `json:"closedDate" description:"Closed date"`
	Currency             string `json:"currency" description:"Currency"`
	InterestRate         string `json:"interestRate" description:"Interest rate"`
	TotalRepaymentAmount string `json:"totalRepaymentAmount" description:"Total repayment amount"`
	PrincipalUSD         string `json:"principalUSD" description:"Principal USD"`
	InterestUSD          string `json:"interestUSD" description:"Interest USD"`
	FeeUSD               string `json:"feeUSD" description:"Fee USD"`
	Tenor                string `json:"tenor" description:"Tenor"`
	GuaranteeMethod      string `json:"guaranteeMethod" description:"Guarantee method:0-No Guarantee 1-Guarantee"`
}

func (*IC000050I) GetServiceKey() string {
	return "IC000050"
}
