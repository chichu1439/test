package models

import "encoding/xml"

type Acmt023 struct {
	XMLName  xml.Name                             `xml:"FpsMsg"`
	NumOfMsg string                               `xml:"NbOfMsgs"`
	XMLHead  Head001                              `xml:"FpsPylds>BizData>AppHdr"`
	Document IdentificationVerificationRequestV03 `xml:"FpsPylds>BizData>Document>IdVrctnReq"`
}

// IdentificationAssignment3 ...
type IdentificationAssignment3 struct {
	MsgId   string                                        `xml:"MsgId"`
	CreDtTm string                                        `xml:"CreDtTm"`
	Cretr   *Party40Choice                                `xml:"Cretr"`
	FrstAgt *BranchAndFinancialInstitutionIdentification6 `xml:"FrstAgt"`
	Assgnr  *Party40Choice                                `xml:"Assgnr"`
	Assgne  *Party40Choice                                `xml:"Assgne"`
}

// IdentificationInformation4 ...
type IdentificationInformation4 struct {
	Pty  *PartyIdentification135                       `xml:"Pty"`
	Acct *CashAccount40                                `xml:"Acct"`
	Agt  *BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
}

// IdentificationVerification4 ...
type IdentificationVerification4 struct {
	Id           string                      `xml:"Id"`
	PtyAndAcctId *IdentificationInformation4 `xml:"PtyAndAcctId"`
}

// IdentificationVerificationRequestV03 ...
type IdentificationVerificationRequestV03 struct {
	Assgnmt     *IdentificationAssignment3     `xml:"Assgnmt"`
	Vrfctn      []*IdentificationVerification4 `xml:"Vrfctn"`
	SplmtryData []*SupplementaryData1          `xml:"SplmtryData"`
}

// Party40Choice ...
type Party40Choice struct {
	Pty *PartyIdentification135                       `xml:"Pty"`
	Agt *BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
}
