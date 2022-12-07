// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP910013I struct {
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
	StlmtDate              string  `json:"stlmtDate"`
	StlmtTime              string  `json:"stlmtTime"`
	StlmtStatus            string  `json:"stlmtStatus"`
	AcceptTime             string  `json:"acceptTime"`
	TransDate              string  `json:"transDate"`
	TransStatus            string  `json:"transStatus"`
	TransReason            string  `json:"transReason"`
	TransReasonInfo        string  `json:"transReasonInfo"`
}

type FP910013O struct {
	Status string `json:"status"`
}

func (*FP910013I) GetServiceKey() string {
	return "FP910013"
}