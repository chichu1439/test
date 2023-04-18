package models

import "encoding/xml"

type FPSCamt054 struct {
	XMLName  xml.Name `xml:"fps:FpsMsg"`
	NbOfMsgs string   `xml:"fps:NbOfMsgs"`
	XmlHead  XmlHead  `xml:"fps:FpsPylds>fps:BizData>ah:AppHdr"`
	Document Camt054  `xml:"fps:FpsPylds>fps:BizData>doc:Document>doc:BkToCstmrDbtCdtNtfctn"`
}

type Camt054 struct {
	MsgId    string `xml:"GrpHdr>MsgId" json:"MsgId,omitempty"`
	CreDtTm  string `xml:"GrpHdr>CreDtTm" json:"CreDtTm,omitempty"`
	NoticeId string `xml:"Ntfctn>Id" json:"NoticeId,omitempty"`
	NoticeTm string `xml:"Ntfctn>CreDtTm" json:"NoticeTm,omitempty"`
	AcctId   string `xml:"Ntfctn>Acct>Id>Othr>Id" json:"AcctId,omitempty"`
	AcctTpCd string `xml:"Ntfctn>Acct>Tp>Cd" json:"AcctTpCd,omitempty"`
	// Amount of money in the cash entry
	//- This is the transaction amount.
	NtryAmt AmtField `xml:"Ntfctn>Ntry>Amt" json:"NtryAmt,omitempty"`
	// Indicates whether the entry is a credit or a debit entry
	NtryCdtDbtInd string `xml:"Ntfctn>Ntry>CdtDbtInd" json:"NtryCdtDbtInd,omitempty"`
	// Status of an entry on the books of the account servicer.
	// The only possible value is “BOOK” in NFPS.
	NtrySts string `xml:"Ntfctn>Ntry>Sts" json:"NtrySts,omitempty"`
	// Transaction code for identifying the type of transaction.
	PrtryCd    string `xml:"Ntfctn>Ntry>BkTxCd>Prtry>Cd" json:"PrtryCd,omitempty"`
	EndToEndId string `xml:"Ntfctn>Ntry>NtryDtls>TxDtls>Refs>EndToEndId" json:"EndToEndId,omitempty"`
	// Transaction ID
	TxId string `xml:"Ntfctn>Ntry>NtryDtls>TxDtls>Refs>TxId" json:"TxId,omitempty"`
	// For Sc.C, this is the NFPS Reference number of the returning payment.
	ClrSysRef string `xml:"Ntfctn>Ntry>NtryDtls>TxDtls>Refs>ClrSysRef" json:"ClrSysRef,omitempty"`
	// Amount of money in the cash entry
	//- This is the transaction amount.
	NtryDtlsAmt AmtField `xml:"Ntfctn>Ntry>NtryDtls>TxDtls>Amt" json:"NtryDtlsAmt,omitempty"`
	// Indicates whether the entry is a credit or a debit entry For Sc.C, this is the opposite of the original payment being returned.
	CdtDbtInd     string `xml:"Ntfctn>Ntry>NtryDtls>TxDtls>CdtDbtInd" json:"NtryDtlsCdtDbtInd,omitempty"`
	DbtrAgtNmbId  string `xml:"Ntfctn>Ntry>NtryDtls>TxDtls>RltdPties>DbtrAgt>FinInstnId>ClrSysMmbId>MmbId" json:"DbtrAgtNmbId,omitempty"`
	CdtrAgtNmbId  string `xml:"Ntfctn>Ntry>NtryDtls>TxDtls>RltdPties>CdtrAgt>FinInstnId>ClrSysMmbId>MmbId" json:"CdtrAgtNmbId,omitempty"`
	IntrBkSttlmDt string `xml:"Ntfctn>Ntry>NtryDtls>TxDtls>RltdDts>IntrBkSttlmDt" json:"IntrBkSttlmDt,omitempty"`
	TxDtTm        string `xml:"Ntfctn>Ntry>NtryDtls>TxDtls>RltdDts>TxDtTm" json:"TxDtTm,omitempty"`
}
