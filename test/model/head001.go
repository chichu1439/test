package models

type Head001 struct {
	FromMemberId   string `xml:"Fr>FIId>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"required,max=35"`
	FromBranchId   string `xml:"Fr>BrnchId>Id,omitempty" validate:"max=35"` //itmx add
	ToMemberId     string `xml:"To>FIId>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"required,max=35"`
	ToBranchId     string `xml:"To>BrnchId>Id,omitempty" validate:"max=35"` //itmx add
	BizMsgId       string `xml:"BizMsgIdr,omitempty" validate:"required,max=35"`
	MsgDefId       string `xml:"MsgDefIdr,omitempty" validate:"required,max=35"`
	BizSvc         string `xml:"BizSvc,omitempty" validate:"required,max=35"`
	CreateDatetime string `xml:"CreDt,omitempty" validate:"required,max=35"`
	CopyDup        string `xml:"CpyDplct,omitempty" validate:"max=35"` //itmx add
}

// AppHdr ...
type AppHdr *BusinessApplicationHeaderV03

// BusinessApplicationHeader7 ...
type BusinessApplicationHeader7 struct {
	CharSet    string                        `xml:"CharSet"`
	Fr         *Party44Choice                `xml:"Fr"`
	To         *Party44Choice                `xml:"To"`
	BizMsgIdr  string                        `xml:"BizMsgIdr"`
	MsgDefIdr  string                        `xml:"MsgDefIdr"`
	BizSvc     string                        `xml:"BizSvc"`
	MktPrctc   *ImplementationSpecification1 `xml:"MktPrctc"`
	CreDt      string                        `xml:"CreDt"`
	BizPrcgDt  string                        `xml:"BizPrcgDt"`
	CpyDplct   string                        `xml:"CpyDplct"`
	PssblDplct bool                          `xml:"PssblDplct"`
	Prty       string                        `xml:"Prty"`
	Sgntr      *SignatureEnvelope            `xml:"Sgntr"`
}

// BusinessApplicationHeaderV03 ...
type BusinessApplicationHeaderV03 struct {
	CharSet    string                        `xml:"CharSet"`
	Fr         *Party44Choice                `xml:"Fr"`
	To         *Party44Choice                `xml:"To"`
	BizMsgIdr  string                        `xml:"BizMsgIdr"`
	MsgDefIdr  string                        `xml:"MsgDefIdr"`
	BizSvc     string                        `xml:"BizSvc"`
	MktPrctc   *ImplementationSpecification1 `xml:"MktPrctc"`
	CreDt      string                        `xml:"CreDt"`
	BizPrcgDt  string                        `xml:"BizPrcgDt"`
	CpyDplct   string                        `xml:"CpyDplct"`
	PssblDplct bool                          `xml:"PssblDplct"`
	Prty       string                        `xml:"Prty"`
	Sgntr      *SignatureEnvelope            `xml:"Sgntr"`
	Rltd       []*BusinessApplicationHeader7 `xml:"Rltd"`
}

// BusinessMessagePriorityCode ...
type BusinessMessagePriorityCode string

// CopyDuplicate1Code ...
type CopyDuplicate1Code string

// ImplementationSpecification1 ...
type ImplementationSpecification1 struct {
	Regy string `xml:"Regy"`
	Id   string `xml:"Id"`
}

// Party44Choice ...
type Party44Choice struct {
	OrgId *PartyIdentification135                       `xml:"OrgId"`
	FIId  *BranchAndFinancialInstitutionIdentification6 `xml:"FIId"`
}

// SignatureEnvelope ...
type SignatureEnvelope struct {
}

// UnicodeChartsCode ...
type UnicodeChartsCode string

// YesNoIndicator ...
type YesNoIndicator bool
