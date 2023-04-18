// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP110030I struct {
	ClrSysRef             string  `json:"clrSysRef" `
	ClrSysId              string  `json:"clrSysId" `
	MsgId                 string  `json:"msgId" `
	MsgDefId              string  `json:"msgDefId"`
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
	MessageType           string  `json:"messageType"`
	MemberId              string  `json:"memberId"`
	MerchantType          string  `json:"merchantType"`
	TermType              string  `json:"termType"`
	OneTimeAuth           string  `json:"oneTimeAuth"`
	OrgInstructionId      string  `json:"orgInstructionId"`
	RecvType              string  `json:"recvType"`
	RecvTaxId             string  `json:"recvTaxId"`
	SendTaxId             string  `json:"sendTaxId"`
	MerchantBillId        string  `json:"merchantBillId"`
	BillDisNameTha        string  `json:"billDisNameTha"`
	BillDisNameEng        string  `json:"billDisNameEng"`
	CustDisNameEng        string  `json:"custDisNameEng"`
	BillRef1              string  `json:"billRef1"`
	BillRef2              string  `json:"billRef2"`
	BillRef3              string  `json:"billRef3"`
	AddiNote              string  `json:"addiNote"`
	ReqExecDate           string  `json:"reqExecDate"`
	ProxyType             string  `json:"proxyType"`
	ProxyVal              string  `json:"proxyVal"`
	LocalAmount           float64 `json:"localAmount"`
	CountryCode           string  `json:"countryCode"`
	VatRate               string  `json:"vatRate"`
	Vat                   string  `json:"vat"`
	IncomeType            string  `json:"incomeType"`
	WithHoldTaxRate       string  `json:"withHoldTaxRate"`
	WithHoldTax           string  `json:"withHoldTax"`
	WithHoldTaxCond       string  `json:"withHoldTaxCond"`
	DrAgtSubId            string  `json:"drAgtSubId"`
	CrAgtSubId            string  `json:"crAgtSubId"`
	DrAddressLine         string  `json:"drAddressLine"`
	DrBirthDate           string  `json:"drBirthDate"`
	DrBirthCity           string  `json:"drBirthCity"`
	DrBirthCountry        string  `json:"drBirthCountry"`
	SendBicCode           string  `json:"sendBicCode"`
	RecvBicCode           string  `json:"recvBicCode"`
	PurposeCode           string  `json:"purposeCode"`
}

func (*FP110030I) GetServiceKey() string {
	return "FP110030"
}

type FP110030O struct {
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
