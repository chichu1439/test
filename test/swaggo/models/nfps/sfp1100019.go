// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP110019I struct {
	MTI                  string               `json:"mti"`
	PrimaryAcctNum       string               `json:"primaryAcctNum" validate:"required"`
	ProcessingCode       string               `json:"processingCode" validate:"required"`
	TransAmt             float64              `json:"transAmt" validate:"required"`
	TransmDateTime       string               `json:"transmDateTime" validate:"required"`
	SysTraceAuditNum     string               `json:"sysTraceAuditNum" validate:"required"`
	LocalTransTime       string               `json:"localTransTime" validate:"required"`
	LocalTransDate       string               `json:"localTransDate" validate:"required"`
	SettleDate           string               `json:"settleDate" validate:"required"`
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

type FP110019O struct {
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

func (*FP110019I) GetServiceKey() string {
	return "FP110019"
}
