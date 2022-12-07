// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP910031I struct {
	MsgId                 string `json:"msgId" validate:"required"`
	StatusCode            string `json:"statusCode"`
	TransDate             string `json:"transDate"`
	TransStatus           string `json:"transStatus"`
	TransRejectCode       string `json:"transRejectCode"`
	TransRejectReason     string `json:"transRejectReason"`
	TransRejectReasonInfo string `json:"transRejectReasonInfo"`
	SettleDate            string `json:"settleDate"`
	SettleTime            string `json:"settleTime"`
	SettleStatus          string `json:"settleStatus"`
	MessageContent        []byte `json:"messageContent"`

	ClrSysRef          string  `json:"clrSysRef,omitempty" validate:""` // Deprecated
	ClrSysId           string  `json:"clrSysId,omitempty"`              // Deprecated
	FromMemberId       string  `json:"fromMemberId,omitempty"`          // Deprecated
	ToMemberId         string  `json:"toMemberId,omitempty"`            // Deprecated
	CreateDatetime     string  `json:"createDatetime,omitempty"`        // Deprecated
	NumOfTrans         string  `json:"numOfTrans,omitempty"`            // Deprecated
	StlmtMethod        string  `json:"stlmtMethod,omitempty"`           // Deprecated
	InstructionId      string  `json:"instructionId,omitempty"`         // Deprecated
	EndToEndId         string  `json:"endToEndId,omitempty"`            // Deprecated
	TransId            string  `json:"transId,omitempty" validate:""`   // Deprecated
	AcctVerifyOption   string  `json:"acctVerifyOption,omitempty"`      // Deprecated
	CategoryPurpose    string  `json:"categoryPurpose,omitempty"`       // Deprecated
	InterbankStlmtAmt  float64 `json:"interbankStlmtAmt,omitempty"`     // Deprecated
	InterbankStlmtCcy  string  `json:"interbankStlmtCcy,omitempty"`     // Deprecated
	InterbankStlmtDate string  `json:"interbankStlmtDate,omitempty"`    // Deprecated
	CreditDatetime     string  `json:"creditDatetime,omitempty"`        // Deprecated
	InstructedAmt      float64 `json:"instructedAmt,omitempty"`         // Deprecated
	InstructedCcy      string  `json:"instructedCcy,omitempty"`         // Deprecated
	ChargeBearerType   string  `json:"chargeBearerType,omitempty"`      // Deprecated
	ChargeAmt          float64 `json:"chargeAmt,omitempty"`             // Deprecated
	ChargeCcy          string  `json:"chargeCcy,omitempty"`             // Deprecated
	ChargePartBic      string  `json:"chargePartBic,omitempty"`         // Deprecated
	ChargeMemberId     string  `json:"chargeMemberId,omitempty"`        // Deprecated
	MandateId          string  `json:"mandateId,omitempty"`             // Deprecated
	DrName             string  `json:"drName,omitempty"`                // Deprecated
	DrOrgBic           string  `json:"drOrgBic,omitempty"`              // Deprecated
	DrId               string  `json:"drId,omitempty"`                  // Deprecated
	DrSchemeName       string  `json:"drSchemeName,omitempty"`          // Deprecated
	DrMemberName       string  `json:"drMemberName,omitempty"`          // Deprecated
	DrCustId           string  `json:"drCustId,omitempty"`              // Deprecated
	DrCustSchemeName   string  `json:"drCustSchemeName,omitempty"`      // Deprecated
	DrCustMemberName   string  `json:"drCustMemberName,omitempty"`      // Deprecated
	DrMobileNum        string  `json:"drMobileNum,omitempty"`           // Deprecated
	DrEmailAddr        string  `json:"drEmailAddr,omitempty"`           // Deprecated
	DrAcctNo           string  `json:"drAcctNo,omitempty"`              // Deprecated
	DrAcctType         string  `json:"drAcctType,omitempty"`            // Deprecated
	DrPartBic          string  `json:"drPartBic,omitempty"`             // Deprecated
	DrMemberId         string  `json:"drMemberId,omitempty"`            // Deprecated
	CrPartBic          string  `json:"crPartBic,omitempty"`             // Deprecated
	CrMemberId         string  `json:"crMemberId,omitempty"`            // Deprecated
	CrName             string  `json:"crName,omitempty"`                // Deprecated
	CrOrgBic           string  `json:"crOrgBic,omitempty"`              // Deprecated
	CrId               string  `json:"crId,omitempty"`                  // Deprecated
	CrSchemeName       string  `json:"crSchemeName,omitempty"`          // Deprecated
	CrMemberName       string  `json:"crMemberName,omitempty"`          // Deprecated
	CrCustId           string  `json:"crCustId,omitempty"`              // Deprecated
	CrCustSchemeName   string  `json:"crCustSchemeName,omitempty"`      // Deprecated
	CrCustMemberName   string  `json:"crCustMemberName,omitempty"`      // Deprecated
	CrMobileNum        string  `json:"crMobileNum,omitempty"`           // Deprecated
	CrEmailAddr        string  `json:"crEmailAddr,omitempty"`           // Deprecated
	CrAcctNo           string  `json:"crAcctNo,omitempty"`              // Deprecated
	CrAcctType         string  `json:"crAcctType,omitempty"`            // Deprecated
	PaymentPurposeCode string  `json:"paymentPurposeCode,omitempty"`    // Deprecated
	PaymentPurposeType string  `json:"paymentPurposeType,omitempty"`    // Deprecated
	Unstructured       string  `json:"unstructured,omitempty"`          // Deprecated
	TransNo            string  `json:"transNo,omitempty"`               // Deprecated
	SendType           string  `json:"sendType,omitempty"`              // Deprecated
	HandleStatusCode   string  `json:"handleStatusCode,omitempty"`      // Deprecated
	ProcessCode        string  `json:"processCode,omitempty"`           // Deprecated
	StlmtDate          string  `json:"stlmtDate,omitempty"`             // Deprecated
	StlmtTime          string  `json:"stlmtTime,omitempty"`             // Deprecated
	StlmtStatus        string  `json:"stlmtStatus,omitempty"`           // Deprecated
	AcctVerifyType     string  `json:"acctVerifyType,omitempty"`        // Deprecated
	ReconcileFlag      string  `json:"reconcileFlag,omitempty"`         // Deprecated
	RejectReason       string  `json:"rejectReason,omitempty"`          // Deprecated
	RejectReasonInfo   string  `json:"rejectReasonInfo,omitempty"`      // Deprecated
	TransType          string  `json:"transType,omitempty"`             // Deprecated
	AcceptDatetime     string  `json:"acceptDatetime,omitempty"`        // Deprecated
}

func (*FP910031I) GetServiceKey() string {
	return "FP910031"
}

type FP910031O struct {
	Status string `json:"status"`
}
