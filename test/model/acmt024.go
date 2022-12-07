package models

import "encoding/xml"

type Acmt024 struct {
	XMLName  xml.Name                            `xml:"fps:FpsMsg"`
	NumOfMsg string                              `xml:"fps:NbOfMsgs"`
	XMLHead  Head001                             `xml:"fps:FpsPylds>fps:BizData>ah:AppHdr"`
	Document IdentificationVerificationReportV03 `xml:"fps:FpsPylds>fps:BizData>doc:Document>doc:acmt.024.001.03"`
}

// ExternalVerificationReason1Code ...
type ExternalVerificationReason1Code string

// IdentificationVerificationIndicator ...
type IdentificationVerificationIndicator bool

// IdentificationVerificationReportV03 ...
type IdentificationVerificationReportV03 struct {
	Assgnmt      *IdentificationAssignment3 `xml:"Assgnmt"`
	OrgnlAssgnmt *MessageIdentification7    `xml:"OrgnlAssgnmt"`
	Rpt          []*VerificationReport4     `xml:"Rpt"`
	SplmtryData  []*SupplementaryData1      `xml:"SplmtryData"`
}

// MessageIdentification7 ...
type MessageIdentification7 struct {
	MsgId   string                                        `xml:"MsgId"`
	CreDtTm string                                        `xml:"CreDtTm"`
	FrstAgt *BranchAndFinancialInstitutionIdentification6 `xml:"FrstAgt"`
}

// VerificationReason1Choice ...
type VerificationReason1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// VerificationReport4 ...
type VerificationReport4 struct {
	OrgnlId           string                      `xml:"OrgnlId"`
	Vrfctn            bool                        `xml:"Vrfctn"`
	Rsn               *VerificationReason1Choice  `xml:"Rsn"`
	OrgnlPtyAndAcctId *IdentificationInformation4 `xml:"OrgnlPtyAndAcctId"`
	UpdtdPtyAndAcctId *IdentificationInformation4 `xml:"UpdtdPtyAndAcctId"`
}
