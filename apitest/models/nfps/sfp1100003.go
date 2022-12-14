// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP110003I struct {
	MsgId                  string  `json:"msgId" `
	CreateDatetime         string  `json:"createDatetime" `
	NumOfTrans             string  `json:"numOfTrans" `
	StlmtMethod            string  `json:"stlmtMethod" `
	ClrSysId               string  `json:"clrSysId" `
	ClrSysRef              string  `json:"clrSysRef"`
	RetId                  string  `json:"retId" `
	FromMemberId           string  `json:"fromMemberId"`
	ToMemberId             string  `json:"toMemberId"`
	CategoryPurpose        string  `json:"categoryPurpose"`
	OrigInstructionId      string  `json:"origInstructionId" `
	OrigEndToEndId         string  `json:"origEndToEndId" `
	OrigTransId            string  `json:"origTransId" `
	OrigClrSysRef          string  `json:"origClrSysRef" `
	RetInterbankStlmtAmt   float64 `json:"retInterbankStlmtAmt" `
	RetInterbankStlmtCcy   string  `json:"retInterbankStlmtCcy" `
	InterbankStlmtDate     string  `json:"interbankStlmtDate" `
	RetInstructedAmt       float64 `json:"retInstructedAmt" `
	RetInstructedCcy       string  `json:"retInstructedCcy" `
	ChargeBearerType       string  `json:"chargeBearerType" `
	ChargeAmt              float64 `json:"chargeAmt" `
	ChargeCcy              string  `json:"chargeCcy" `
	ChargePartBic          string  `json:"chargePartBic" `
	ChargeMemberId         string  `json:"chargeMemberId" `
	RetReasonCode          string  `json:"retReasonCode" `
	RetReasonInfo          string  `json:"retReasonInfo" `
	OrigInterbankStlmtAmt  float64 `json:"origInterbankStlmtAmt" `
	OrigInterbankStlmtCcy  string  `json:"origInterbankStlmtCcy" `
	OrigInterbankStlmtDate string  `json:"origInterbankStlmtDate" `
	OrigCategoryPurpose    string  `json:"origCategoryPurpose" `
	OrigMandateId          string  `json:"origMandateId" `
	OrigUnstructured       string  `json:"origUnstructured" `
	OrigDrName             string  `json:"origDrName" `
	OrigDrOrgBic           string  `json:"origDrOrgBic" `
	OrigDrId               string  `json:"origDrId" `
	OrigDrSchemeName       string  `json:"origDrSchemeName" `
	OrigDrMemberName       string  `json:"origDrMemberName" `
	OrigDrCustId           string  `json:"origDrCustId" `
	OrigDrCustSchemeName   string  `json:"origDrCustSchemeName" `
	OrigDrCustMemberName   string  `json:"origDrCustMemberName" `
	OrigDrMobileNum        string  `json:"origDrMobileNum" `
	OrigDrEmailAddr        string  `json:"origDrEmailAddr" `
	OrigDrAcctNo           string  `json:"origDrAcctNo" `
	OrigDrAcctType         string  `json:"origDrAcctType" `
	OrigDrPartBic          string  `json:"origDrPartBic" `
	OrigDrMemberId         string  `json:"origDrMemberId" `
	OrigCrPartBic          string  `json:"origCrPartBic" `
	OrigCrMemberId         string  `json:"origCrMemberId" `
	OrigCrName             string  `json:"origCrName" `
	OrigCrOrgBic           string  `json:"origCrOrgBic" `
	OrigCrId               string  `json:"origCrId" `
	OrigCrSchemeName       string  `json:"origCrSchemeName" `
	OrigCrMemberName       string  `json:"origCrMemberName" `
	OrigCrCustId           string  `json:"origCrCustId" `
	OrigCrCustSchemeName   string  `json:"origCrCustSchemeName" `
	OrigCrCustMemberName   string  `json:"origCrCustMemberName" `
	OrigCrMobileNum        string  `json:"origCrMobileNum" `
	OrigCrEmailAddr        string  `json:"origCrEmailAddr" `
	OrigCrAcctNo           string  `json:"origCrAcctNo" `
	OrigCrAcctType         string  `json:"origCrAcctType" `
}

func (*FP110003I) GetServiceKey() string {
	return "FP110003"
}

type FP110003O struct {
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