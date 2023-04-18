// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP113016I struct {
	ClrSysRef             string  `json:"clrSysRef"`
	ClrSysId              string  `json:"clrSysId"`
	MmbId                 string  `json:"mmbId"`
	MsgId                 string  `json:"msgId"`
	FromMemberId          string  `json:"fromMemberId"`
	ToMemberId            string  `json:"toMemberId"`
	CreateDatetime        string  `json:"createDatetime"`
	NumOfTrans            string  `json:"numOfTrans"`
	StlmtMethod           string  `json:"stlmtMethod"`
	InstructionId         string  `json:"instructionId"`
	EndToEndId            string  `json:"endToEndId"`
	TransId               string  `json:"transId"`
	TransNo               string  `json:"transNo"`
	InterbankStlmtAmt     float64 `json:"interbankStlmtAmt"`
	InterbankStlmtCcy     string  `json:"interbankStlmtCcy"`
	InterbankStlmtDate    string  `json:"interbankStlmtDate"`
	CreditDatetime        string  `json:"creditDatetime"`
	InstructedAmt         float64 `json:"instructedAmt"`
	InstructedCcy         string  `json:"instructedCcy"`
	DrName                string  `json:"drName"`
	DrOrgBic              string  `json:"drOrgBic"`
	DrId                  string  `json:"drId"`
	DrAcctNo              string  `json:"drAcctNo"`
	DrAcctType            string  `json:"drAcctType"`
	DrPartBic             string  `json:"drPartBic"`
	LookupProxyId         string  `json:"lookupProxyId"`
	LookupProxyType       string  `json:"lookupProxyType"`
	LookupProxyVal        string  `json:"lookupProxyVal"`
	LookupAccId           string  `json:"lookupAccId"`
	LookupAccDisNm        string  `json:"lookupAccDisNm"`
	LookupAccAgtMmbId     string  `json:"lookupAccAgtMmbId"`
	LookupAccAcId         string  `json:"lookupAccAcId"`
	LookupAccAcNm         string  `json:"lookupAccAcNm"`
	AdditionalData        string  `json:"additionalData"`
	AcctVerifyType        string  `json:"acctVerifyType"`
	ReconcileFlag         string  `json:"reconcileFlag"`
	RejectReason          string  `json:"rejectReason"`
	RejectReasonInfo      string  `json:"rejectReasonInfo"`
	TransStatus           string  `json:"transStatus"`
	TransRejectReason     string  `json:"transRejectReason"`
	TransRejectReasonInfo string  `json:"transRejectReasonInfo"`
	AcceptDatetime        string  `json:"acceptDatetime"`
	MessageContent        []byte  `json:"messageContent"`
}

func (*FP113016I) GetServiceKey() string {
	return "FP113016"
}

type FP113016O struct {
	MsgId                 string `json:"msgId"`
	FromMemberId          string `json:"fromMemberId" `
	ToMemberId            string `json:"toMemberId" `
	TransStatus           string `json:"transStatus"`
	TransCode             string `json:"transCode"`
	TransRejectReason     string `json:"transRejectReason"`
	TransRejectReasonInfo string `json:"transRejectReasonInfo"`
	MessageContent        []byte `json:"messageContent"`
}
