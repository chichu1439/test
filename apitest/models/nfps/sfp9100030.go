// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP910030I struct {
	MsgId     string `json:"msgId"`
	TransId   string `json:"transId"`
	ClrSysRef string `json:"clrSysRef"`
}

func (t *FP910030I) GetServiceKey() string {
	return "FP910030"
}

type FP910030O struct {
	Records []FP910030Record
}

type FP910030Record struct {
	MsgId                 string  `json:"msgId,omitempty"`
	MsgType               string  `json:"msgType,omitempty"`
	TransId               string  `json:"transId,omitempty"`
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
	CreditorMemberId      string  `json:"creditorMemberId,omitempty"`
	CreditorMemberName    string  `json:"creditorMemberName,omitempty"`
	CreditorAcctId        string  `json:"creditorAcctId,omitempty"`
	CreditorAcctName      string  `json:"creditorAcctName,omitempty"`
	ProxyType             string  `json:"proxyType,omitempty"`
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

	DrName     string `json:"drName,omitempty"`     // Deprecated
	DrAcctNo   string `json:"drAcctNo,omitempty"`   // Deprecated
	DrAcctType string `json:"drAcctType,omitempty"` // Deprecated

	CrName     string `json:"crName,omitempty"`     // Deprecated
	CrAcctNo   string `json:"crAcctNo,omitempty"`   // Deprecated
	CrAcctType string `json:"crAcctType,omitempty"` // Deprecated

	ClrSysRef          string  `json:"clrSysRef,omitempty"`          // Deprecated
	ClrSysId           string  `json:"clrSysId,omitempty"`           // Deprecated
	FromMemberId       string  `json:"fromMemberId,omitempty"`       // Deprecated
	ToMemberId         string  `json:"toMemberId,omitempty"`         // Deprecated
	NumOfTrans         string  `json:"numOfTrans,omitempty"`         // Deprecated
	StlmtMethod        string  `json:"stlmtMethod,omitempty"`        // Deprecated
	AcctVerifyOption   string  `json:"acctVerifyOption,omitempty"`   // Deprecated
	CategoryPurpose    string  `json:"categoryPurpose,omitempty"`    // Deprecated
	InterbankStlmtAmt  float64 `json:"interbankStlmtAmt,omitempty"`  // Deprecated
	InterbankStlmtCcy  string  `json:"interbankStlmtCcy,omitempty"`  // Deprecated
	InterbankStlmtDate string  `json:"interbankStlmtDate,omitempty"` // Deprecated
	CreditDatetime     string  `json:"creditDatetime,omitempty"`     // Deprecated
	InstructedAmt      float64 `json:"instructedAmt,omitempty"`      // Deprecated
	InstructedCcy      string  `json:"instructedCcy,omitempty"`      // Deprecated
	ChargeBearerType   string  `json:"chargeBearerType,omitempty"`   // Deprecated
	ChargeAmt          float64 `json:"chargeAmt,omitempty"`          // Deprecated
	ChargeCcy          string  `json:"chargeCcy,omitempty"`          // Deprecated
	ChargePartBic      string  `json:"chargePartBic,omitempty"`      // Deprecated
	ChargeMemberId     string  `json:"chargeMemberId,omitempty"`     // Deprecated
	MandateId          string  `json:"mandateId,omitempty"`          // Deprecated
	DrOrgBic           string  `json:"drOrgBic,omitempty"`           // Deprecated
	DrId               string  `json:"drId,omitempty"`               // Deprecated
	DrSchemeName       string  `json:"drSchemeName,omitempty"`       // Deprecated
	DrMemberName       string  `json:"drMemberName,omitempty"`       // Deprecated
	DrCustId           string  `json:"drCustId,omitempty"`           // Deprecated
	DrCustSchemeName   string  `json:"drCustSchemeName,omitempty"`   // Deprecated
	DrCustMemberName   string  `json:"drCustMemberName,omitempty"`   // Deprecated
	DrMobileNum        string  `json:"drMobileNum,omitempty"`        // Deprecated
	DrEmailAddr        string  `json:"drEmailAddr,omitempty"`        // Deprecated
	DrPartBic          string  `json:"drPartBic,omitempty"`          // Deprecated
	DrMemberId         string  `json:"drMemberId,omitempty"`         // Deprecated
	CrPartBic          string  `json:"crPartBic,omitempty"`          // Deprecated
	CrMemberId         string  `json:"crMemberId,omitempty"`         // Deprecated
	CrOrgBic           string  `json:"crOrgBic,omitempty"`           // Deprecated
	CrId               string  `json:"crId,omitempty"`               // Deprecated
	CrSchemeName       string  `json:"crSchemeName,omitempty"`       // Deprecated
	CrMemberName       string  `json:"crMemberName,omitempty"`       // Deprecated
	CrCustId           string  `json:"crCustId,omitempty"`           // Deprecated
	CrCustSchemeName   string  `json:"crCustSchemeName,omitempty"`   // Deprecated
	CrCustMemberName   string  `json:"crCustMemberName,omitempty"`   // Deprecated
	CrMobileNum        string  `json:"crMobileNum,omitempty"`        // Deprecated
	CrEmailAddr        string  `json:"crEmailAddr,omitempty"`        // Deprecated
	PaymentPurposeType string  `json:"paymentPurposeType,omitempty"` // Deprecated
	Unstructured       string  `json:"unstructured,omitempty"`       // Deprecated
	SendType           string  `json:"sendType,omitempty"`           // Deprecated
	HandleStatusCode   string  `json:"handleStatusCode,omitempty"`   // Deprecated
	StlmtDate          string  `json:"stlmtDate,omitempty"`          // Deprecated
	StlmtTime          string  `json:"stlmtTime,omitempty"`          // Deprecated
	StlmtStatus        string  `json:"stlmtStatus,omitempty"`        // Deprecated
	AcctVerifyType     string  `json:"acctVerifyType,omitempty"`     // Deprecated
	ReconcileFlag      string  `json:"reconcileFlag,omitempty"`      // Deprecated
	RejectReason       string  `json:"rejectReason,omitempty"`       // Deprecated
	RejectReasonInfo   string  `json:"rejectReasonInfo,omitempty"`   // Deprecated
	AcceptDatetime     string  `json:"acceptDatetime,omitempty"`     // Deprecated
}