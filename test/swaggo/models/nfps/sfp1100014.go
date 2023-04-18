// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP110014I struct {
	FromMemberId       string  `json:"fromMemberId" `
	ToMemberId         string  `json:"toMemberId" `
	MsgDefId           string  `json:"msgDefId" `
	BizSvc             string  `json:"bizSvc" `
	MsgId              string  `json:"msgId" `
	CreateDatetime     string  `json:"createDatetime" `
	NumOfTrans         string  `json:"numOfTrans" `
	StlmtMethod        string  `json:"stlmtMethod" `
	ClrSysId           string  `json:"clrSysId" `
	InstructionId      string  `json:"instructionId" `
	EndToEndId         string  `json:"endToEndId" `
	TransId            string  `json:"transId" `
	ClrSysRef          string  `json:"clrSysRef" `
	CategoryPurpose    string  `json:"categoryPurpose" `
	InterbankStlmtAmt  float64 `json:"interbankStlmtAmt" `
	InterbankStlmtCcy  string  `json:"interbankStlmtCcy" `
	InterbankStlmtDate string  `json:"interbankStlmtDate" `
	InstructedAmt      float64 `json:"instructedAmt" `
	InstructedCcy      string  `json:"instructedCcy" `
	ChargeBearerType   string  `json:"chargeBearerType" `
	ChargeAmt          float64 `json:"chargeAmt" `
	ChargeCcy          string  `json:"chargeCcy" `
	ChargePartBic      string  `json:"chargePartBic" `
	ChargeMemberId     string  `json:"chargeMemberId" `
	MandateId          string  `json:"mandateId" `
	CrName             string  `json:"crName" `
	CrOrgBic           string  `json:"crOrgBic" `
	CrId               string  `json:"crId" `
	CrSchemeName       string  `json:"crSchemeName" `
	CrMemberName       string  `json:"crMemberName" `
	CrCustId           string  `json:"crCustId" `
	CrCustSchemeName   string  `json:"crCustSchemeName" `
	CrCustMemberName   string  `json:"crCustMemberName" `
	CrMobileNum        string  `json:"crMobileNum" `
	CrEmailAddr        string  `json:"crEmailAddr" `
	CrAcctNo           string  `json:"crAcctNo" `
	CrAcctType         string  `json:"crAcctType" `
	CrPartBic          string  `json:"crPartBic" `
	CrMemberId         string  `json:"crMemberId" `
	DrName             string  `json:"drName" `
	DrOrgBic           string  `json:"drOrgBic" `
	DrId               string  `json:"drId" `
	DrSchemeName       string  `json:"drSchemeName" `
	DrMemberName       string  `json:"drMemberName" `
	DrCustId           string  `json:"drCustId" `
	DrCustSchemeName   string  `json:"drCustSchemeName" `
	DrCustMemberName   string  `json:"drCustMemberName" `
	DrMobileNum        string  `json:"drMobileNum" `
	DrEmailAddr        string  `json:"drEmailAddr" `
	DrAcctNo           string  `json:"drAcctNo" `
	DrAcctType         string  `json:"drAcctType" `
	DrPartBic          string  `json:"drPartBic" `
	DrMemberId         string  `json:"drMemberId" `
	PaymentPurposeCode string  `json:"paymentPurposeCode" `
	PaymentPurposeType string  `json:"paymentPurposeType" `
	Unstructured       string  `json:"unstructured" `
}

type FP110014O struct {
	Status string `json:"status"`
}

func (*FP110014I) GetServiceKey() string {
	return "FP110014"
}
