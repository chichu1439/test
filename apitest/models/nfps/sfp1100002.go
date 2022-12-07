// Package models will define request and response message struct
// Version: v0.0.1
package models

import (
	"net/http"
	"strings"
)

type FP110002I struct {
	ClrSysRef             string  `json:"clrSysRef" `
	ClrSysId              string  `json:"clrSysId" `
	MsgId                 string  `json:"msgId" `
	FromMemberId          string  `json:"fromMemberId" `
	ToMemberId            string  `json:"toMemberId" `
	CreateDatetime        string  `json:"createDatetime" `
	NumOfTrans            string  `json:"numOfTrans" `
	StlmtMethod           string  `json:"stlmtMethod" `
	InstructionId         string  `json:"instructionId" `
	EndToEndId            string  `json:"endToEndId" `
	TransId               string  `json:"transId" `
	AcctVerifyOption      string  `json:"acctVerifyOption" `
	CategoryPurpose       string  `json:"categoryPurpose" `
	InterbankStlmtAmt     float64 `json:"interbankStlmtAmt" `
	InterbankStlmtCcy     string  `json:"interbankStlmtCcy" `
	InterbankStlmtDate    string  `json:"interbankStlmtDate" `
	CreditDatetime        string  `json:"creditDatetime" `
	InstructedAmt         float64 `json:"instructedAmt" `
	InstructedCcy         string  `json:"instructedCcy" `
	ChargeBearerType      string  `json:"chargeBearerType" `
	ChargeAmt             float64 `json:"chargeAmt" `
	ChargeCcy             string  `json:"chargeCcy" `
	ChargePartBic         string  `json:"chargePartBic" `
	ChargeMemberId        string  `json:"chargeMemberId" `
	MandateId             string  `json:"mandateId" `
	DrName                string  `json:"drName" `
	DrOrgBic              string  `json:"drOrgBic" `
	DrId                  string  `json:"drId" `
	DrSchemeName          string  `json:"drSchemeName" `
	DrMemberName          string  `json:"drMemberName" `
	DrCustId              string  `json:"drCustId" `
	DrCustSchemeName      string  `json:"drCustSchemeName" `
	DrCustMemberName      string  `json:"drCustMemberName" `
	DrMobileNum           string  `json:"drMobileNum" `
	DrEmailAddr           string  `json:"drEmailAddr" `
	DrAcctNo              string  `json:"drAcctNo" `
	DrAcctType            string  `json:"drAcctType" `
	DrPartBic             string  `json:"drPartBic" `
	DrMemberId            string  `json:"drMemberId" `
	CrPartBic             string  `json:"crPartBic" `
	CrMemberId            string  `json:"crMemberId" `
	CrName                string  `json:"crName" `
	CrOrgBic              string  `json:"crOrgBic" `
	CrId                  string  `json:"crId" `
	CrSchemeName          string  `json:"crSchemeName" `
	CrMemberName          string  `json:"crMemberName" `
	CrCustId              string  `json:"crCustId" `
	CrCustSchemeName      string  `json:"crCustSchemeName" `
	CrCustMemberName      string  `json:"crCustMemberName" `
	CrMobileNum           string  `json:"crMobileNum" `
	CrEmailAddr           string  `json:"crEmailAddr" `
	CrAcctNo              string  `json:"crAcctNo" `
	CrAcctType            string  `json:"crAcctType" `
	PaymentPurposeCode    string  `json:"paymentPurposeCode" `
	PaymentPurposeType    string  `json:"paymentPurposeType" `
	Unstructured          string  `json:"unstructured" `
	TransNo               string  `json:"transNo" `
	SendType              string  `json:"sendType" `
	HandleStatusCode      string  `json:"handleStatusCode" `
	ProcessCode           string  `json:"processCode" `
	StlmtDate             string  `json:"stlmtDate" `
	StlmtTime             string  `json:"stlmtTime" `
	StlmtStatus           string  `json:"stlmtStatus" `
	AcctVerifyType        string  `json:"acctVerifyType" `
	ReconcileFlag         string  `json:"reconcileFlag" `
	RejectReason          string  `json:"rejectReason" `
	RejectReasonInfo      string  `json:"rejectReasonInfo" `
	TransStatus           string  `json:"transStatus" `
	TransRejectReason     string  `json:"transRejectReason" `
	TransRejectReasonInfo string  `json:"transRejectReasonInfo" `
	AcceptDatetime        string  `json:"acceptDatetime" `
	MessageContent        []byte  `json:"messageContent"`
}

func (*FP110002I) GetServiceKey() string {
	return "FP110002"
}

type FP110002O struct {
	ClrSysRef             string `json:"clrSysRef" `
	ClrSysId              string `json:"clrSysId" `
	MsgId                 string `json:"msgId" `
	FromMemberId          string `json:"fromMemberId" `
	ToMemberId            string `json:"toMemberId" `
	CreateDatetime        string `json:"createDatetime" `
	NumOfTrans            string `json:"numOfTrans" `
	HandleStatusCode      string `json:"handleStatusCode" `
	ProcessCode           string `json:"processCode" `
	StlmtDate             string `json:"stlmtDate" `
	StlmtTime             string `json:"stlmtTime" `
	StlmtStatus           string `json:"stlmtStatus" `
	AcctVerifyType        string `json:"acctVerifyType" `
	ReconcileFlag         string `json:"reconcileFlag" `
	RejectReason          string `json:"rejectReason" `
	RejectReasonInfo      string `json:"rejectReasonInfo" `
	TransStatus           string `json:"transStatus" `
	TransRejectReason     string `json:"transRejectReason" `
	TransRejectReasonInfo string `json:"transRejectReasonInfo" `
	AcceptDatetime        string `json:"acceptDatetime" `
}

type FP110002Outgress struct {
	Body []byte `json:"body"`
}

func (*FP110002Outgress) GetServiceKey() string {
	return "FP210002"
}

func (o *FP110002OutgressResponse) GetHeaderIgnoreCase(key string) (value string) {

	for k, vs := range o.Headers {
		if len(vs) == 0 {
			return ""
		}
		if strings.EqualFold(k, key) {
			return vs[0]
		}
	}

	return ""
}

type FP110002OutgressResponse struct {
	Headers http.Header `json:"headers"`
	Body    []byte      `json:"body"`
}
