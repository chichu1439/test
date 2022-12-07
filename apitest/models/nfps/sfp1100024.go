// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP110024I struct {
	MTI                  string               `json:"mti"`
	PrimaryAcctNum       string               `json:"primaryAcctNum" validate:"required"`
	ProcessingCode       string               `json:"processingCode" validate:"required"`
	TransAmt             float64              `json:"transAmt" validate:"required"`
	TransmDateTime       string               `json:"transmDateTime" validate:"required"`
	SysTraceAuditNum     string               `json:"sysTraceAuditNum" validate:"required"`
	LocalTransTime       string               `json:"localTransTime" validate:"required"`
	LocalTransDate       string               `json:"localTransDate" validate:"required"`
	SettleDate           string               `json:"settleDate"`
	MerchantType         string               `json:"merchantType" validate:"required"`
	POSEntryCode         string               `json:"posEntryCode" validate:"required"`
	CardSeqNum           string               `json:"cardSeqNum"`
	SendingID            string               `json:"sendingId" validate:"required"`
	ForwardingID         string               `json:"forwardingId" validate:"required"`
	Track2Data           string               `json:"track2Data"`
	RetrievalRefNum      string               `json:"retrievalRefNum" validate:"required"`
	ResponseCode         string               `json:"responseCode"`
	CardAcceptorLocation CardAcceptorLocation `json:"cardAcceptorLocation" validate:"required"`
	ITMXSegment          ITMXRequestToPay     `json:"itmxSegment" validate:"required"`
	TransCurrency        string               `json:"transCurrency" validate:"required"` // eg: 764
	PINBlock             string               `json:"pinBlock" validate:"required"`
	EMVData              string               `json:"emvData"`
	ITMXPOSInfo          ITMXPOSInfo          `json:"itmxPosInfo" validate:"required"`
	NetworkCode          string               `json:"networkCode"`
	FromAcct             string               `json:"fromAcct"`
	ToAcct               string               `json:"toAcct"`
	MessageContent       []byte               `json:"messageContent"`
}

type FP110024O struct {
	MTI                  string               `json:"mti"`
	PrimaryAcctNum       string               `json:"primaryAcctNum"`
	ProcessingCode       string               `json:"processingCode"`
	TransAmt             float64              `json:"transAmt"`
	TransmDateTime       string               `json:"transmDateTime"`
	SysTraceAuditNum     string               `json:"sysTraceAuditNum"`
	LocalTransTime       string               `json:"localTransTime"`
	LocalTransDate       string               `json:"localTransDate"`
	SettleDate           string               `json:"settleDate"`
	MerchantType         string               `json:"merchantType"`
	POSEntryCode         string               `json:"posEntryCode"`
	CardSeqNum           string               `json:"cardSeqNum"`
	SendingID            string               `json:"sendingId"`
	ForwardingID         string               `json:"forwardingId"`
	Track2Data           string               `json:"track2Data"`
	RetrievalRefNum      string               `json:"retrievalRefNum"`
	ResponseCode         string               `json:"responseCode"`
	CardAcceptorLocation CardAcceptorLocation `json:"cardAcceptorLocation"`
	ITMXSegment          ITMXRequestToPay     `json:"itmxSegment"`
	TransCurrency        string               `json:"transCurrency"` // eg: 764
	PINBlock             string               `json:"pinBlock"`
	EMVData              string               `json:"emvData"`
	ITMXPOSInfo          ITMXPOSInfo          `json:"itmxPosInfo"`
	NetworkCode          string               `json:"networkCode"`
	FromAcct             string               `json:"fromAcct"`
	ToAcct               string               `json:"toAcct"`
}

func (*FP110024I) GetServiceKey() string {
	return "FP110024"
}

//type CardAcceptorLocation struct {
//	TerminalOwnerID          string `json:"terminalOwnerId"`
//	TerminalCardAcceptorName string `json:"terminalCardAcceptorName"`
//	TerminalType             string `json:"terminalType"`
//	BranchID                 string `json:"branchId"`
//	TerminalState            string `json:"terminalState"`
//	ProvinceCode             string `json:"provinceCode"`
//	PromptPayFlag            string `json:"promptPayFlag"`
//	Filler                   string `json:"filler"`
//	CountryCode              string `json:"countryCode"`
//}

type ITMXRequestToPay struct {
	SenderFee            float64 `json:"senderFee" description:"F48.1" validate:""`
	FromAcctFee          float64 `json:"fromAcctFee" description:"F48.2" validate:""`
	ToAcctFee            float64 `json:"toAcctFee" description:"F48.3"`
	RequesterBank        string  `json:"requesterBank" description:"F48.4" validate:"required"`
	RequesterAcctNo      string  `json:"requesterAcctNo" description:"F48.5" validate:"required"`
	RequesterAcctName    string  `json:"requesterAcctName" description:"F48.6" validate:"required"`
	ReceiverID           string  `json:"receiverId" description:"F48.7" validate:"required"`
	ReceiverProxyType    string  `json:"receiverProxyType" description:"F48.8" validate:"required"`
	ReceiverBank         string  `json:"receiverBank" description:"F48.9"`
	ReceiverAcctNo       string  `json:"receiverAcctNo" description:"F48.10"`
	ReceiverAcctDispName string  `json:"receiverAcctDispName" description:"F48.11"`
	ReceiverAcctName     string  `json:"receiverAcctName" description:"F48.12"`
	PINBlockType         string  `json:"pinBlockType" description:"F48.13" validate:"required"`
	OneTimePassword      string  `json:"oneTimePassword" description:"F48.14"`
	AdditionalNote       string  `json:"additionalNote" description:"F48.15"`
	TransRef             string  `json:"transRef" description:"F48.16"`
	RequesterID          string  `json:"requesterId" description:"F48.17"`
	RequesterProxyType   string  `json:"requesterProxyType" description:"F48.18"`
	CreditNotifyRefNo    string  `json:"creditNotifyRefNo" description:"F48.19"`
	BillRef1             string  `json:"billRef1" description:"F48.20"`
	BillRef2             string  `json:"billRef2" description:"F48.21"`
	BillRef3             string  `json:"billRef3" description:"F48.22"`
	BillPresentURL       string  `json:"billPresentUrl" description:"F48.23"`
	ExpiryDate           string  `json:"expiryDate" description:"F48.24"`
}

type ITMXCreditNotify struct {
	SenderFee            float64 `json:"senderFee" description:"F48.1" validate:""`
	FromAcctFee          float64 `json:"fromAcctFee" description:"F48.2" validate:""`
	ToAcctFee            float64 `json:"toAcctFee" description:"F48.3"`
	SendingBank          string  `json:"sendingBank" description:"F48.4" validate:"required"`
	FromAcctNo           string  `json:"fromAcctNo" description:"F48.5" validate:"required"`
	FromAcctName         string  `json:"fromAcctName" description:"F48.6" validate:"required"`
	ReceiverID           string  `json:"receiverId" description:"F48.7" validate:"required"`
	ReceiverProxyType    string  `json:"receiverProxyType" description:"F48.8" validate:"required"`
	SponsorBank          string  `json:"sponsorBank" description:"F48.9"`
	ToAcctNo             string  `json:"toAcctNo" description:"F48.10"`
	ToAcctDispName       string  `json:"toAcctDispName" description:"F48.11"`
	ToAcctName           string  `json:"toAcctName" description:"F48.12"`
	PINBlockType         string  `json:"pinBlockType" description:"F48.13" validate:"required"`
	OneTimePassword      string  `json:"oneTimePassword" description:"F48.14"`
	AdditionalNote       string  `json:"additionalNote" description:"F48.15"`
	TransRef             string  `json:"transRef" description:"F48.16"`
	SenderID             string  `json:"senderId" description:"F48.17"`
	SenderProxyType      string  `json:"senderProxyType" description:"F48.18"`
	CreditGurDocumentURL string  `json:"creditGurDocumentUrl" description:"F48.19"`
}

type ITMXBillPayment struct {
	SenderFee             float64 `json:"senderFee" description:"F48.1" validate:""`
	FromAcctFee           float64 `json:"fromAcctFee" description:"F48.2" validate:""`
	ToAcctFee             float64 `json:"toAcctFee" description:"F48.3"`
	SendingBank           string  `json:"sendingBank" description:"F48.4" validate:"required"`
	SendAcctNo            string  `json:"sendAcctNo" description:"F48.5" validate:"required"`
	SendAcctName          string  `json:"sendAcctName" description:"F48.6" validate:"required"`
	BillID                string  `json:"billId" description:"F48.7" validate:"required"`
	ProxyType             string  `json:"proxyType" description:"F48.8" validate:"required"`
	SponsorBank           string  `json:"sponsorBank" description:"F48.9"`
	BillAcctNo            string  `json:"billAcctNo" description:"F48.10"`
	BillAcctDispName      string  `json:"billAcctDispName" description:"F48.11"`
	BillAcctName          string  `json:"billAcctName" description:"F48.12"`
	PINBlockType          string  `json:"pinBlockType" description:"F48.13" validate:"required"`
	OneTimePassword       string  `json:"oneTimePassword" description:"F48.14"`
	AdditionalNote        string  `json:"additionalNote" description:"F48.15"`
	TransRef              string  `json:"transRef" description:"F48.16"`
	TransType             string  `json:"transType" description:"F48.17" validate:"required"`
	SenderID              string  `json:"senderId" description:"F48.18"`
	SenderProxyType       string  `json:"senderProxyType" description:"F48.19"`
	SenderRefNo           string  `json:"senderRefNo" description:"F48.20"`
	SenderTaxID           string  `json:"senderTaxId" description:"F48.21"`
	SendingBankBranchCode string  `json:"sendingBankBranchCode" description:"F48.22"`
	BillTaxID             string  `json:"billTaxId" description:"F48.23"`
	BillBankBranchCode    string  `json:"billBankBranchCode" description:"F48.24"`
	RefNum                string  `json:"refNum" description:"F48.25"`
	VATRate               float64 `json:"vatRate" description:"F48.26"`
	VATAmt                float64 `json:"vatAmt" description:"F48.27"`
	IncomeType            string  `json:"incomeType" description:"F48.28"`
	WithholdingTaxRate    float64 `json:"withholdingTaxRate" description:"F48.29"`
	WithholdingTaxAmt     float64 `json:"withholdingTaxAmt" description:"F48.30"`
	WithholdingTaxCond    string  `json:"withholdingTaxCond" description:"F48.31"`
	SenderType            string  `json:"senderType" description:"F48.32" validate:"required"`
	ReceiverType          string  `json:"receiverType" description:"F48.33"`
	BillRef1              string  `json:"billRef1" description:"F48.34" validate:"required"`
	BillRef2              string  `json:"billRef2" description:"F48.35"`
	BillRef3              string  `json:"billRef3" description:"F48.36"`
	DueDate               string  `json:"dueDate" description:"F48.37"`
}

type ITMXGenericInquiry struct {
	SenderFee        float64 `json:"senderFee" description:"F48.1" validate:""`
	FromAcctFee      float64 `json:"fromAcctFee" description:"F48.2" validate:""`
	ToAcctFee        float64 `json:"toAcctFee" description:"F48.3"`
	SendingBank      string  `json:"sendingBank" description:"F48.4" validate:"required"`
	SendAcctNo       string  `json:"sendAcctNo" description:"F48.5" validate:"required"`
	SendAcctName     string  `json:"sendAcctName" description:"F48.6" validate:"required"`
	ProxyID          string  `json:"proxyId" description:"F48.7" validate:"required"`
	ProxyType        string  `json:"proxyType" description:"F48.8" validate:"required"`
	SponsorBank      string  `json:"sponsorBank" description:"F48.9"`
	BillAcctNo       string  `json:"billAcctNo" description:"F48.10"`
	BillAcctDispName string  `json:"billAcctDispName" description:"F48.11"`
	BillAcctName     string  `json:"billAcctName" description:"F48.12"`
	PINBlockType     string  `json:"pinBlockType" description:"F48.13" validate:"required"`
	OneTimePassword  string  `json:"oneTimePassword" description:"F48.14"`
	AdditionalNote   string  `json:"additionalNote" description:"F48.15"`
	TransRef         string  `json:"transRef" description:"F48.16"`
	BillRef1         string  `json:"billRef1" description:"F48.17" validate:"required"`
	BillRef2         string  `json:"billRef2" description:"F48.18"`
	BillRef3         string  `json:"billRef3" description:"F48.19"`
	BillAmt          float64 `json:"billAmt" description:"F48.20"`
	DirectLinkURL    string  `json:"directLinkUrl" description:"F48.21"`
}

//type ITMXPOSInfo struct {
//	TerminalType     string `json:"terminalType"`
//	TerminalEntryCap string `json:"terminalEntryCap"`
//	ChipCondCode     string `json:"chipCondCode"`
//	ChipTransInd     string `json:"chipTransInd"`
//	ChipCardAuth     string `json:"chipCardAuth"`
//}
