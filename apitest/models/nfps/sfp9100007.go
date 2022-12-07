package models

type FP910007I struct {
	TransId      string `json:"transId"`
	MsgId        string `json:"msgId"`
	PartCode     string `json:"partCode"`
	PageNo       int    `json:"pageNo" `
	PageRecCount int    `json:"pageRecCount" `
}
type FP910007O struct {
	//TotalCount    int             `json:"totalCount"`
	//JournalList   []FP910007OItem `json:"journalList"`
	TotalRecCount int             `json:"totalRecCount"`
	TotalPage     int             `json:"totalPage"`
	Records       []FP910007OItem `json:"records"`
}
type FP910007OItem struct {
	Id                    int64   `json:"id"`
	MsgId                 string  `json:"msgId"`
	MsgType               string  `json:"msgType"`
	TransId               string  `json:"transId"`
	InstructionId         string  `json:"instructionId"`
	EndToEndId            string  `json:"endToEndId"`
	TransNo               string  `json:"transNo"`
	CreateDatetime        string  `json:"createDatetime"`
	TransCurrency         string  `json:"transCurrency"`
	TransAmt              float64 `json:"transAmt"`
	TransType             string  `json:"transType"`
	PaymentPurposeCode    string  `json:"paymentPurposeCode"`
	ProcessCode           string  `json:"processCode"`
	DebtorMemberId        string  `json:"debtorMemberId"`
	DebtorMemberName      string  `json:"debtorMemberName"`
	DebtorAcctId          string  `json:"debtorAcctId"`
	DebtorAcctName        string  `json:"debtorAcctName"`
	DebtorAcctType        string  `json:"debtorAcctType"`
	DebtorMobileNum       string  `json:"debtorMobileNum"`
	DebtorEmailAddr       string  `json:"debtorEmailAddr"`
	DebtorProxyId         string  `json:"debtorProxyId"`
	DebtorProxyType       string  `json:"debtorProxyType"`
	CreditorMemberId      string  `json:"creditorMemberId"`
	CreditorMemberName    string  `json:"creditorMemberName"`
	CreditorAcctId        string  `json:"creditorAcctId"`
	CreditorAcctName      string  `json:"creditorAcctName"`
	CreditorAcctType      string  `json:"creditorAcctType"`
	CreditorMobileNum     string  `json:"creditorMobileNum"`
	CreditorEmailAddr     string  `json:"creditorEmailAddr"`
	CreditorProxyId       string  `json:"creditorProxyId"`
	CreditorProxyType     string  `json:"creditorProxyType"`
	StatusCode            string  `json:"statusCode"`
	TransRef              string  `json:"transRef"`
	TransDate             string  `json:"transDate"`
	TransStatus           string  `json:"transStatus"`
	TransRejectCode       string  `json:"transRejectCode"`
	TransRejectReason     string  `json:"transRejectReason"`
	TransRejectReasonInfo string  `json:"transRejectReasonInfo"`
	SettleDate            string  `json:"settleDate"`
	SettleTime            string  `json:"settleTime"`
	SettleStatus          string  `json:"settleStatus"`
	SettleMethod          string  `json:"settleMethod"`
	AcctVerifyOption      string  `json:"acctVerifyOption"`
	CategoryPurpose       string  `json:"categoryPurpose"`
	InterbankStlmtAmt     float64 `json:"interbankStlmtAmt"`
	InterbankStlmtCcy     string  `json:"interbankStlmtCcy"`
	InterbankStlmtDate    string  `json:"interbankStlmtDate"`
	InstructedAmt         float64 `json:"instructedAmt"`
	InstructedCcy         string  `json:"instructedCcy"`
	ChargeBearerType      string  `json:"chargeBearerType"`
	ChargeAmt             float64 `json:"chargeAmt"`
	ChargeCcy             string  `json:"chargeCcy"`
	ChargePartBic         string  `json:"chargePartBic"`
	ChargeMemberId        string  `json:"chargeMemberId"`
	MandateId             string  `json:"mandateId"`
	PaymentPurposeType    string  `json:"paymentPurposeType"`
	Unstructured          string  `json:"unstructured"`
	AcceptDatetime        string  `json:"acceptDatetime"`
	PrimaryAcctNum        string  `json:"primaryAcctNum"`
	SysTraceAuditNum      string  `json:"sysTraceAuditNum"`
	MerchantType          string  `json:"merchantType"`
	PosEntryCode          string  `json:"posEntryCode"`
	CardSeqNum            string  `json:"cardSeqNum"`
	ForwardingId          string  `json:"forwardingId"`
	RetrievalRefNum       string  `json:"retrievalRefNum"`
	CardAcceptorName      string  `json:"cardAcceptorName"`
	SenderFee             float64 `json:"senderFee"`
	FromAcctFee           float64 `json:"fromAcctFee"`
	ToAcctFee             float64 `json:"toAcctFee"`
	PinBlockType          string  `json:"pinBlockType"`
	TransTaxType          string  `json:"transTaxType"`
	SenderRefNo           string  `json:"senderRefNo"`
	SenderTaxId           string  `json:"senderTaxId"`
	SendBranchCode        string  `json:"sendBranchCode"`
	ReceiverTaxId         string  `json:"receiverTaxId"`
	ReceiverBranchCode    string  `json:"receiverBranchCode"`
	BankRefNum            string  `json:"bankRefNum"`
	VatRates              float64 `json:"vatRates"`
	VatAmt                float64 `json:"vatAmt"`
	IncomeType            string  `json:"incomeType"`
	WithholdingTaxRates   float64 `json:"withholdingTaxRates"`
	WithholdingTaxAmt     float64 `json:"withholdingTaxAmt"`
	WithholdingTaxCond    string  `json:"withholdingTaxCond"`
	SenderType            string  `json:"senderType"`
	ReceiverType          string  `json:"receiverType"`
	BillRef1              string  `json:"billRef1"`
	BillRef2              string  `json:"billRef2"`
	BillRef3              string  `json:"billRef3"`
	PresentURL            string  `json:"presentURL"`
	BillDueDate           string  `json:"billDueDate"`
	EmvData               string  `json:"emvData"`
	PosInfo               string  `json:"posInfo"`
	NetworkCode           string  `json:"networkCode"`
	RecordType            string  `json:"recordType"`
	CreateTime            string  `json:"createTime"`
	UpdateTime            string  `json:"updateTime"`
}

func (*FP910007I) GetServiceKey() string {
	return "FP910007"
}
