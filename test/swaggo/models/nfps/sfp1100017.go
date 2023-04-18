// Package models will define request and response message struct
// Version: v0.0.1
package models

import "fmt"

type FP110017I struct {
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
	ITMXSegment          ITMXCreditTransfer   `json:"itmxSegment" validate:"required"`
	TransCurrency        string               `json:"transCurrency" validate:"required"` // eg: 764
	PINBlock             string               `json:"pinBlock" validate:"required"`
	EMVData              string               `json:"emvData"`
	ITMXPOSInfo          ITMXPOSInfo          `json:"itmxPosInfo" validate:"required"`
	NetworkCode          string               `json:"networkCode"`
	FromAcct             string               `json:"fromAcct"`
	ToAcct               string               `json:"toAcct"`
}

type FP110017O struct {
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
	ITMXSegment          ITMXCreditTransfer   `json:"itmxSegment"`
	TransCurrency        string               `json:"transCurrency"` // eg: 764
	PINBlock             string               `json:"pinBlock"`
	EMVData              string               `json:"emvData"`
	ITMXPOSInfo          ITMXPOSInfo          `json:"itmxPosInfo"`
	NetworkCode          string               `json:"networkCode"`
	FromAcct             string               `json:"fromAcct"`
	ToAcct               string               `json:"toAcct"`
}

func (*FP110017I) GetServiceKey() string {
	return "FP110017"
}

type CardAcceptorLocation struct {
	TerminalOwnerID          string `json:"terminalOwnerId"`
	TerminalCardAcceptorName string `json:"terminalCardAcceptorName"`
	TerminalType             string `json:"terminalType"`
	BranchID                 string `json:"branchId"`
	TerminalState            string `json:"terminalState"`
	ProvinceCode             string `json:"provinceCode"`
	PromptPayFlag            string `json:"promptPayFlag"`
	Filler                   string `json:"filler"`
	CountryCode              string `json:"countryCode"`
}

type ITMXCreditTransfer struct {
	SenderFee              float64 `json:"senderFee" validate:""`
	FromAcctFee            float64 `json:"fromAcctFee" validate:""`
	ToAcctFee              float64 `json:"toAcctFee"`
	SendingBank            string  `json:"sendingBank" validate:"required"`
	FromAcctNo             string  `json:"fromAcctNo" validate:"required"`
	FromAcctName           string  `json:"fromAcctName" validate:"required"`
	ReceivingID            string  `json:"receivingId" validate:"required"`
	ProxyType              string  `json:"proxyType" validate:"required"`
	DestBank               string  `json:"destBank"`
	ToAcctNo               string  `json:"toAcctNo"`
	ToAcctDispName         string  `json:"toAcctDispName"`
	ToAcctName             string  `json:"toAcctName"`
	PINBlockType           string  `json:"pinBlockType" validate:"required"`
	OneTimePassword        string  `json:"oneTimePassword"`
	AdditionalNote         string  `json:"additionalNote"`
	TransRef               string  `json:"transRef"`
	TransType              string  `json:"transType" validate:"required"`
	SenderID               string  `json:"senderId"`
	SenderProxyType        string  `json:"senderProxyType"`
	SenderRefNo            string  `json:"senderRefNo"`
	SenderTaxID            string  `json:"senderTaxId"`
	SendingBankBranchCode  string  `json:"sendingBankBranchCode"`
	ReceiverTaxID          string  `json:"receiverTaxId"`
	ReceiverBankBranchCode string  `json:"receiverBankBranchCode"`
	RefNum                 string  `json:"refNum"`
	VATRate                float64 `json:"vatRate"`
	VATAmt                 float64 `json:"vatAmt"`
	IncomeType             string  `json:"incomeType"`
	WithholdingTaxRate     float64 `json:"withholdingTaxRate"`
	WithholdingTaxAmt      float64 `json:"withholdingTaxAmt"`
	WithholdingTaxCond     string  `json:"withholdingTaxCond"`
	SenderType             string  `json:"senderType" validate:"required"`
	ReceiverType           string  `json:"receiverType"`
}

type ITMXPOSInfo struct {
	TerminalType     string `json:"terminalType"`
	TerminalEntryCap string `json:"terminalEntryCap"`
	ChipCondCode     string `json:"chipCondCode"`
	ChipTransInd     string `json:"chipTransInd"`
	ChipCardAuth     string `json:"chipCardAuth"`
}

func (c CardAcceptorLocation) String() string {
	return c.TerminalOwnerID + c.TerminalCardAcceptorName +
		c.TerminalType + c.BranchID + c.TerminalState + c.ProvinceCode +
		c.PromptPayFlag + c.Filler + c.CountryCode
}

const ITMXPOSInfoDefaultFiller = "0"

func (i ITMXPOSInfo) String() string {
	return fmt.Sprintf("%s%s%s000%s%s0000", i.TerminalType, i.TerminalEntryCap, i.ChipCondCode,
		i.ChipTransInd, i.ChipCardAuth)
}
