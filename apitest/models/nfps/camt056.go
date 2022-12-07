package models

import "encoding/xml"

type FPSCamt056 struct {
	XMLName  xml.Name `xml:"FpsMsg"`
	NumOfMsg string   `xml:"NbOfMsgs"`
	XMLHead  Head001  `xml:"FpsPylds>BizData>AppHdr"`
	Document Camt056  `xml:"FpsPylds>BizData>Document"`
}

// Camt056
// The FIToFI Payment Cancellation Request message is sent by a case creator/case assigner to a case assignee.
// This message is used to request the cancellation of an original payment instruction. The FIToFI Payment Cancellation Request message is exchanged between the instructing agent and the instructed agent to request the cancellation of a interbank payment message previously sent (such as FIToFICustomerCreditTransfer, FIToFICustomerDirectDebit or FinancialInstitutionCreditTransfer).
// Usage
// The FIToFI Payment Cancellation Request message must be answered with a:
// - Resolution Of Investigation message with a positive final outcome when the case assignee can perform the requested cancellation
// - Resolution Of Investigation message with a negative final outcome when the case assignee may perform the requested cancellation but fails to do so (too late, irrevocable instruction, ...)
// - Reject Investigation message when the case assignee is unable or not authorised to perform the requested cancellation
// - Notification Of Case Assignment message to indicate whether the case assignee will take on the case himself or reassign the case to a subsequent party in the payment processing chain.
// A FIToFI Payment Cancellation Request message concerns one and only one original payment instruction at a time.
// When a case assignee successfully performs a cancellation, it must return the corresponding funds to the case assigner. It may provide some details about the return in the Resolution Of Investigation message.
// The processing of a FIToFI Payment Cancellation Request message case may lead to a Debit Authorisation Request message sent to the creditor by its account servicing institution.
// The FIToFI Payment Cancellation Request message may be used to escalate a case after an unsuccessful request to modify the payment. In this scenario, the case identification rem
type Camt056 struct {

	// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	// Uniquely identifies the case assignment.
	MsgId string `xml:"FIToFIPmtCxlReq>Assgnmt>Id,omitempty"`
	// Party who assigns the case.
	// Usage: This is also the sender of the message.
	AssignerClrSysId string `xml:"FIToFIPmtCxlReq>Assgnmt>Assgnr>Agt>FinInstnId>ClrSysMmbId>MmbId,omitempty"`
	// Party to which the case is assigned.
	// Usage: This is also the receiver of the message.
	AssigneeClrSysId string `xml:"FIToFIPmtCxlReq>Assgnmt>Assgne>Agt>FinInstnId>ClrSysMmbId>MmbId,omitempty"`
	// Date and time at which the assignment was created.
	CreationDateTime      string `xml:"FIToFIPmtCxlReq>Assgnmt>CreDtTm,omitempty"`
	InstructionId         string `xml:"FIToFIPmtCxlReq>Case>Id,omitempty"`                                     //itmx add
	CreatorMemberId       string `xml:"FIToFIPmtCxlReq>Case>Cretr>Agt>FinInstnId>ClrSysMmbId>MmbId,omitempty"` //itmx add
	OriginalMessageId     string `xml:"FIToFIPmtCxlReq>Undrlyg>OrgnlGrpInfAndCxl>OrgnlMsgId,omitempty"`        //itmx add
	OriginalMessageNameId string `xml:"FIToFIPmtCxlReq>Undrlyg>OrgnlGrpInfAndCxl><OrgnlMsgNmId,omitempty"`     //itmx add

	// Identifies the payment instruction to be cancelled.
	// Set of elements used to provide information on the original transactions to which the cancellation request message refers.
	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionId string `xml:"FIToFIPmtCxlReq>Undrlyg>TxInf>OrgnlInstrId,omitempty"`
	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndId string `xml:"FIToFIPmtCxlReq>Undrlyg>TxInf>OrgnlEndToEndId,omitempty"`
	// Unique identification, as assigned by the original first instructing agent, to umambiguously identify the transaction.
	OriginalTransId string `xml:"FIToFIPmtCxlReq>Undrlyg>TxInf>OrgnlTxId,omitempty"`
	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClrSysRef          string `xml:"FIToFIPmtCxlReq>Undrlyg>TxInf>OrgnlClrSysRef,omitempty"`
	OriginalInterbankStlmtAmt  string `xml:"FIToFIPmtCxlReq>Undrlyg>TxInf>OrgnlIntrBkSttlmAmt,omitempty"`
	OriginalInterbankStlmtDate string `xml:"FIToFIPmtCxlReq>Undrlyg>TxInf>OrgnlIntrBkSttlmDt,omitempty"`
	CancelReasonCode           string `xml:"FIToFIPmtCxlReq>Undrlyg>TxInf>CxlRsnInf>Rsn>Cd,omitempty"`
	//// Set of elements used to provide detailed information on the cancellation reason.
	CancelReason string `xml:"FIToFIPmtCxlReq>Undrlyg>TxInf>CxlRsnInf>Rsn>Prtry,omitempty"` // removed
	//// Further details on the cancellation request reason.
	CancelReasonInfo string `xml:"FIToFIPmtCxlReq>Undrlyg>TxInf>CxlRsnInf>Rsn>AddtlInf,omitempty"` // removed
}
