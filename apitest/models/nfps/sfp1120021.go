// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP112021I struct {
	ClrSysRef             string  `json:"clrSysRef"`
	ClrSysId              string  `json:"clrSysId"`
	MmbId                 string  `json:"MmbId"`
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

func (*FP112021I) GetServiceKey() string {
	return "FP112021"
}

type FP112021O struct {
	ClrSysRef             string `json:"clrSysRef"`
	ClrSysId              string `json:"clrSysId"`
	MsgId                 string `json:"msgId"`
	FromMemberId          string `json:"fromMemberId"`
	ToMemberId            string `json:"toMemberId"`
	CreateDatetime        string `json:"createDatetime"`
	NumOfTrans            string `json:"numOfTrans"`
	HandleStatusCode      string `json:"handleStatusCode"`
	ProcessCode           string `json:"processCode"`
	AcctVerifyType        string `json:"acctVerifyType"`
	TransStatus           string `json:"transStatus"`
	TransRejectReason     string `json:"transRejectReason"`
	TransRejectReasonInfo string `json:"transRejectReasonInfo"`
	AcceptDatetime        string `json:"acceptDatetime"`
}

type FP112021Outgress struct {
	Body []byte `json:"body"`
}

func (*FP112021Outgress) GetServiceKey() string {
	return "FP212021"
}

type FP112021OutgressResponse struct {
	Body []byte `json:"body"`
}
