package models

type IC000049I struct {
	TranDate string `json:"TranDate" description:"Tran date"` //业务日期
	IsDaily  bool   `json:"isDaily" description:"Is daily"`   //是否日报
}

type IC000049O struct {
	FileKey string `json:"fileKey" description:"File key"`
}

type ReportLoanBO struct {
	SeqNo               string `json:"seqNo" description:"Seq no"`
	CIF                 string `json:"cif" description:"Cif"`
	CustomerName        string `json:"customerName" description:"Customer name"`
	LoanReferenceNumber string `json:"loanReferenceNumber" description:"Loan reference number"`
	LoanStatus          string `json:"loanStatus" description:"Loan status"`
	IDType              string `json:"iDType" description:"Id type"`
	CertificateID       string `json:"certificateID" description:"Certificate id"`
	CustomerPhoneNumber string `json:"customerPhoneNumber" description:"Customer phone number"`
	ProductID           string `json:"productID" description:"Product id"`
	ProductName         string `json:"productName" description:"Product name"`
	DrawdownDate        string `json:"drawdownDate" description:"Drawdown date"`
	MaturityDate        string `json:"maturityDate" description:"Maturity date"`
	PrincipalAmount     string `json:"principalAmount" description:"Principal Amount"`
	DisbursementAmount  string `json:"disbursementAmount" description:"Disbursement Amount"`
	StampDuty           string `json:"stampDuty" description:"Stamp Duty"`
	FrontEndFee         string `json:"frontEndFee" description:"Front End Fee"`
	InterestRate        string `json:"interestRate" description:"Interest rate"`
	RepaymentDate       string `json:"repaymentDate" description:"Repayment date"`
	RepaymentMethod     string `json:"repaymentMethod" description:"Repayment method"`
	Currency            string `json:"currency" description:"Currency"`
	Tenor               string `json:"tenor" description:"Tenor"`
	RepaymentFrequency  string `json:"repaymentFrequency" description:"Repayment frequency"`
	DisbursementChannel string `json:"disbursementChannel" description:"Disbursement channel"`
	BankName            string `json:"bankName" description:"Bank name"`
	BankAccountName     string `json:"bankAccountName" description:"Bank account name"`
	BankAccountNumber   string `json:"bankAccountNumber" description:"Bank account number"`
	GuaranteeMethod     string `json:"guaranteeMethod" description:"Guarantee method"`
	ArrangementPurpose  string `json:"arrangementPurpose" description:"Arrangement purpose"`
}

func (*IC000049I) GetServiceKey() string {
	return "IC000049"
}
