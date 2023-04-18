package models

type FP110027I struct {
	MTI                  string                    `json:"mti"`
	PrimaryAcctNum       string                    `json:"primaryAcctNum" validate:"required"`
	ProcessingCode       string                    `json:"processingCode" validate:"required"`
	TransAmt             float64                   `json:"transAmt" validate:"required"`
	TransmDateTime       string                    `json:"transmDateTime" validate:"required"`
	SysTraceAuditNum     string                    `json:"sysTraceAuditNum" validate:"required"`
	LocalTransTime       string                    `json:"localTransTime" validate:"required"`
	LocalTransDate       string                    `json:"localTransDate" validate:"required"`
	SettleDate           string                    `json:"settleDate"`
	MerchantType         string                    `json:"merchantType" validate:"required"`
	POSEntryCode         string                    `json:"posEntryCode" validate:"required"`
	CardSeqNum           string                    `json:"cardSeqNum"`
	SendingID            string                    `json:"sendingId" validate:"required"`
	ForwardingID         string                    `json:"forwardingId" validate:"required"`
	Track2Data           string                    `json:"track2Data"`
	RetrievalRefNum      string                    `json:"retrievalRefNum" validate:"required"`
	ResponseCode         string                    `json:"responseCode"`
	CardAcceptorLocation CardAcceptorLocation      `json:"cardAcceptorLocation" validate:"required"`
	ITMXSegment          ITMXNotifIcationStransfer `json:"itmxSegment" validate:"required"`
	TransCurrency        string                    `json:"transCurrency" validate:"required"` // eg: 764
	PINBlock             string                    `json:"pinBlock" validate:"required"`
	EMVData              string                    `json:"emvData"`
	ITMXPOSInfo          ITMXPOSInfo               `json:"itmxPosInfo" validate:"required"`
	NetworkCode          string                    `json:"networkCode"`
	FromAcct             string                    `json:"fromAcct"`
	ToAcct               string                    `json:"toAcct"`
	MessageContent       []byte                    `json:"messageContent"`
}

func (*FP110027I) GetServiceKey() string {
	return "FP110027"
}

type FP110027O struct {
	MTI                  string                    `json:"mti"`
	PrimaryAcctNum       string                    `json:"primaryAcctNum"`
	ProcessingCode       string                    `json:"processingCode"`
	TransAmt             float64                   `json:"transAmt"`
	TransmDateTime       string                    `json:"transmDateTime"`
	SysTraceAuditNum     string                    `json:"sysTraceAuditNum"`
	LocalTransTime       string                    `json:"localTransTime"`
	LocalTransDate       string                    `json:"localTransDate"`
	SettleDate           string                    `json:"settleDate"`
	MerchantType         string                    `json:"merchantType"`
	POSEntryCode         string                    `json:"posEntryCode"`
	CardSeqNum           string                    `json:"cardSeqNum"`
	SendingID            string                    `json:"sendingId"`
	ForwardingID         string                    `json:"forwardingId"`
	Track2Data           string                    `json:"track2Data"`
	RetrievalRefNum      string                    `json:"retrievalRefNum"`
	ResponseCode         string                    `json:"responseCode"`
	CardAcceptorLocation CardAcceptorLocation      `json:"cardAcceptorLocation"`
	ITMXSegment          ITMXNotifIcationStransfer `json:"itmxSegment"`
	TransCurrency        string                    `json:"transCurrency"` // eg: 764
	PINBlock             string                    `json:"pinBlock"`
	EMVData              string                    `json:"emvData"`
	ITMXPOSInfo          ITMXPOSInfo               `json:"itmxPosInfo"`
	NetworkCode          string                    `json:"networkCode"`
	FromAcct             string                    `json:"fromAcct"`
	ToAcct               string                    `json:"toAcct"`
}

type ITMXNotifIcationStransfer struct {
	SenderFee                   float64 `json:"senderFee" validate:""`
	FromAcctFee                 float64 `json:"fromAcctFee" validate:""`
	ToAcctFee                   float64 `json:"toAcctFee"`
	SendingBank                 string  `json:"sendingBank" validate:"required"`
	FromAcctNo                  string  `json:"fromAcctNo" validate:"required"`
	FromAcctName                string  `json:"fromAcctName" validate:"required"`
	BillerID                    string  `json:"billerID" validate:"required"`
	ProxyType                   string  `json:"proxyType" validate:"required"`
	SponsorBankOfBiller         string  `json:"sponsorBankOfBiller"`
	ToAcctNo                    string  `json:"toAcctNo"`
	ToAcctDispName              string  `json:"toAcctDispName"`
	ToAcctName                  string  `json:"toAcctName"`
	PINBlockType                string  `json:"pinBlockType" validate:"required"`
	OneTimePassword             string  `json:"oneTimePassword"`
	AdditionalNote              string  `json:"additionalNote"`
	CreditNotificationReference string  `json:"transRef" validate:"required"`
	SenderID                    string  `json:"transType" validate:"required"`
	SenderProxyType             string  `json:"senderId" validate:"required"`
	CreditGuaranteeDocumentURL  string  `json:"creditGuaranteeDocumentUrl"`
}
