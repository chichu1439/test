package models

type FP910032I struct {
	MsgId                 string  `json:"msgId,omitempty" validate:"required"`
	MsgType               string  `json:"msgType,omitempty" validate:"required"`
	TransId               string  `json:"transId,omitempty" validate:"required"`
	InstructionId         string  `json:"instructionId,omitempty"`
	EndToEndId            string  `json:"endToEndId,omitempty"`
	TransNo               string  `json:"transNo,omitempty"`
	CreateDatetime        string  `json:"createDatetime,omitempty"`
	TransCurrency         string  `json:"transCurrency,omitempty"`
	TransAmt              float64 `json:"transAmt,omitempty"`
	TransType             string  `json:"transType,omitempty"`
	PaymentPurposeCode    string  `json:"paymentPurposeCode,omitempty"`
	ProcessCode           string  `json:"processCode,omitempty"`
	DebtorMemberId        string  `json:"debtorMemberId,omitempty"`
	DebtorMemberName      string  `json:"debtorMemberName,omitempty"`
	DebtorAcctId          string  `json:"debtorAcctId,omitempty"`
	DebtorAcctName        string  `json:"debtorAcctName,omitempty"`
	DebtorAcctType        string  `json:"debtorAcctType,omitempty"`
	DebtorMobileNum       string  `json:"debtorMobileNum,omitempty"`
	DebtorEmailAddr       string  `json:"debtorEmailAddr,omitempty"`
	DebtorProxyId         string  `json:"debtorProxyId,omitempty"`
	DebtorProxyType       string  `json:"debtorProxyType,omitempty"`
	CreditorMemberId      string  `json:"creditorMemberId,omitempty"`
	CreditorMemberName    string  `json:"creditorMemberName,omitempty"`
	CreditorAcctId        string  `json:"creditorAcctId,omitempty"`
	CreditorAcctName      string  `json:"creditorAcctName,omitempty"`
	CreditorAcctType      string  `json:"creditorAcctType,omitempty"`
	CreditorMobileNum     string  `json:"creditorMobileNum,omitempty"`
	CreditorEmailAddr     string  `json:"creditorEmailAddr,omitempty"`
	CreditorProxyId       string  `json:"creditorProxyId,omitempty"`
	CreditorProxyType     string  `json:"creditorProxyType,omitempty"`
	StatusCode            string  `json:"statusCode,omitempty"`
	TransRef              string  `json:"transRef,omitempty"`
	TransDate             string  `json:"transDate,omitempty"`
	TransStatus           string  `json:"transStatus,omitempty"`
	TransRejectCode       string  `json:"transRejectCode,omitempty"`
	TransRejectReason     string  `json:"transRejectReason,omitempty"`
	TransRejectReasonInfo string  `json:"transRejectReasonInfo,omitempty"`
	SettleDate            string  `json:"settleDate,omitempty"`
	SettleTime            string  `json:"settleTime,omitempty"`
	SettleStatus          string  `json:"settleStatus,omitempty"`
	SettleMethod          string  `json:"settleMethod,omitempty"`
	ClrSysRef             string  `json:"clrSysRef,omitempty"`
	ClrSysId              string  `json:"clrSysId,omitempty"`
	FromMemberId          string  `json:"fromMemberId,omitempty"` // Deprecated
	ToMemberId            string  `json:"toMemberId,omitempty"`   // Deprecated
	NumOfTrans            string  `json:"numOfTrans,omitempty"`   // Deprecated
	StlmtMethod           string  `json:"stlmtMethod,omitempty"`  // Deprecated
	AcctVerifyOption      string  `json:"acctVerifyOption,omitempty"`
	CategoryPurpose       string  `json:"categoryPurpose,omitempty"`
	InterbankStlmtAmt     float64 `json:"interbankStlmtAmt,omitempty"`
	InterbankStlmtCcy     string  `json:"interbankStlmtCcy,omitempty"`
	InterbankStlmtDate    string  `json:"interbankStlmtDate,omitempty"`
	CreditDatetime        string  `json:"creditDatetime,omitempty"`
	InstructedAmt         float64 `json:"instructedAmt,omitempty"`
	InstructedCcy         string  `json:"instructedCcy,omitempty"`
	ChargeBearerType      string  `json:"chargeBearerType,omitempty"`
	ChargeAmt             float64 `json:"chargeAmt,omitempty"`
	ChargeCcy             string  `json:"chargeCcy,omitempty"`
	ChargePartBic         string  `json:"chargePartBic,omitempty"`
	ChargeMemberId        string  `json:"chargeMemberId,omitempty"`
	MandateId             string  `json:"mandateId,omitempty"`
	DrName                string  `json:"drName,omitempty"`           // Deprecated
	DrOrgBic              string  `json:"drOrgBic,omitempty"`         // Deprecated
	DrId                  string  `json:"drId,omitempty"`             // Deprecated
	DrSchemeName          string  `json:"drSchemeName,omitempty"`     // Deprecated
	DrMemberName          string  `json:"drMemberName,omitempty"`     // Deprecated
	DrCustId              string  `json:"drCustId,omitempty"`         // Deprecated
	DrCustSchemeName      string  `json:"drCustSchemeName,omitempty"` // Deprecated
	DrCustMemberName      string  `json:"drCustMemberName,omitempty"` // Deprecated
	DrMobileNum           string  `json:"drMobileNum,omitempty"`      // Deprecated
	DrEmailAddr           string  `json:"drEmailAddr,omitempty"`      // Deprecated
	DrAcctNo              string  `json:"drAcctNo,omitempty"`         // Deprecated
	DrAcctType            string  `json:"drAcctType,omitempty"`       // Deprecated
	DrPartBic             string  `json:"drPartBic,omitempty"`        // Deprecated
	DrMemberId            string  `json:"drMemberId,omitempty"`       // Deprecated
	CrPartBic             string  `json:"crPartBic,omitempty"`        // Deprecated
	CrMemberId            string  `json:"crMemberId,omitempty"`       // Deprecated
	CrName                string  `json:"crName,omitempty"`           // Deprecated
	CrOrgBic              string  `json:"crOrgBic,omitempty"`         // Deprecated
	CrId                  string  `json:"crId,omitempty"`             // Deprecated
	CrSchemeName          string  `json:"crSchemeName,omitempty"`     // Deprecated
	CrMemberName          string  `json:"crMemberName,omitempty"`     // Deprecated
	CrCustId              string  `json:"crCustId,omitempty"`         // Deprecated
	CrCustSchemeName      string  `json:"crCustSchemeName,omitempty"` // Deprecated
	CrCustMemberName      string  `json:"crCustMemberName,omitempty"` // Deprecated
	CrMobileNum           string  `json:"crMobileNum,omitempty"`      // Deprecated
	CrEmailAddr           string  `json:"crEmailAddr,omitempty"`      // Deprecated
	CrAcctNo              string  `json:"crAcctNo,omitempty"`         // Deprecated
	CrAcctType            string  `json:"crAcctType,omitempty"`       // Deprecated
	PaymentPurposeType    string  `json:"paymentPurposeType,omitempty"`
	Unstructured          string  `json:"unstructured,omitempty"`
	SendType              string  `json:"sendType,omitempty"`
	HandleStatusCode      string  `json:"handleStatusCode,omitempty"` // Deprecated
	StlmtDate             string  `json:"stlmtDate,omitempty"`        // Deprecated
	StlmtTime             string  `json:"stlmtTime,omitempty"`        // Deprecated
	StlmtStatus           string  `json:"stlmtStatus,omitempty"`      // Deprecated
	AcctVerifyType        string  `json:"acctVerifyType,omitempty"`
	ReconcileFlag         string  `json:"reconcileFlag,omitempty"`
	RejectReason          string  `json:"rejectReason,omitempty"`
	RejectReasonInfo      string  `json:"rejectReasonInfo,omitempty"`
	AcceptDatetime        string  `json:"acceptDatetime,omitempty"`
	RecordType            string  `json:"recordType,omitempty"`
	PrimaryAcctNum        string  `json:"primaryAcctNum,omitempty"`
	SysTraceAuditNum      string  `json:"sysTraceAuditNum,omitempty"`
	MerchantType          string  `json:"merchantType,omitempty"`
	PosEntryCode          string  `json:"posEntryCode,omitempty"`
	CardSeqNum            string  `json:"cardSeqNum,omitempty"`
	ForwardingId          string  `json:"forwardingId,omitempty"`
	RetrievalRefNum       string  `json:"retrievalRefNum,omitempty"`
	CardAcceptorName      string  `json:"cardAcceptorName,omitempty"`
	SenderFee             float64 `json:"senderFee,omitempty"`
	FromAcctFee           float64 `json:"fromAcctFee,omitempty"`
	ToAcctFee             float64 `json:"toAcctFee,omitempty"`
	ProxyType             string  `json:"proxyType,omitempty"`
	ProxyId               string  `json:"proxyId,omitempty"` // Deprecated
	PinBlockType          string  `json:"pinBlockType,omitempty"`
	TransTaxType          string  `json:"transTaxType,omitempty"`
	SenderProxyId         string  `json:"senderProxyId,omitempty"`   // Deprecated
	SenderProxyType       string  `json:"senderProxyType,omitempty"` // Deprecated
	SenderRefNo           string  `json:"senderRefNo,omitempty"`
	SenderTaxId           string  `json:"senderTaxId,omitempty"`
	SendBranchCode        string  `json:"sendBranchCode,omitempty"`
	ReceiverTaxId         string  `json:"receiverTaxId,omitempty"`
	ReceiverBranchCode    string  `json:"receiverBranchCode,omitempty"`
	BankRefNum            string  `json:"bankRefNum,omitempty"`
	VatRates              float64 `json:"vatRates,omitempty"`
	VatAmt                float64 `json:"vatAmt,omitempty"`
	IncomeType            string  `json:"incomeType,omitempty"`
	WithholdingTaxRates   float64 `json:"withholdingTaxRates,omitempty"`
	WithholdingTaxAmt     float64 `json:"withholdingTaxAmt,omitempty"`
	WithholdingTaxCond    string  `json:"withholdingTaxCond,omitempty"`
	SenderType            string  `json:"senderType,omitempty"`
	ReceiverType          string  `json:"receiverType,omitempty"`
	BillRef1              string  `json:"billRef1,omitempty"`
	BillRef2              string  `json:"billRef2,omitempty"`
	BillRef3              string  `json:"billRef3,omitempty"`
	BillDueDate           string  `json:"billDueDate,omitempty"`
	PresentURL            string  `json:"presentURL,omitempty"` // Contains short URL for present credit guarantee or bill slip or direct Link for payment from biller
	EmvData               string  `json:"emvData,omitempty"`
	PosInfo               string  `json:"posInfo,omitempty"`
	NetworkCode           string  `json:"networkCode,omitempty"`
}

func (*FP910032I) GetServiceKey() string {
	return "FP910032"
}

type FP910032O struct {
	Status string
}
