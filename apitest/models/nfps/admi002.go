package models

import "encoding/xml"

type FPSAdmi002 struct {
	XMLName  xml.Name `xml:"fps:FpsMsg"`
	NumOfMsg string   `xml:"fps:NbOfMsgs"`
	XMLHead  Head001  `xml:"fps:FpsPylds>fps:BizData>ah:AppHdr"`
	Document Admi002  `xml:"fps:FpsPylds>fps:BizData>doc:Document>doc:admi.002.001.01"`
}

type Admi002 struct {
	// Refers to the identification of the message previously received and for which the rejection is notified.
	RelatedRef string `xml:"doc:RltdRef>doc:Ref"`
	// Reason of the rejection provided by the rejecting party
	RejectReason string `xml:"doc:Rsn>doc:RjctgPtyRsn"`
	// Date and time at which the rejection was generated.
	RejectDatetime string `xml:"doc:Rsn>doc:RjctnDtTm"`
	// Description of the precise location of the potential error in a message.
	ErrorLocation string `xml:"doc:Rsn>doc:ErrLctn,omitempty"`
	// Detailed description of the rejection reason.
	ReasonDesc string `xml:"doc:Rsn>doc:RsnDesc"`
	// Additional information related to the rejection and meant to allow for the precise identification of the rejection reason. This could include a copy of the rejected message in part or in full.
	AdditionalData *AdditionalData `xml:"doc:Rsn>doc:AddtlData"`
}

type AdditionalData struct {
	XMLName xml.Name `xml:"doc:AddtlData"`
	Text    string   `xml:",cdata"`
}
