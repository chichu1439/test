package models

type MsgHeard struct {
	HeaderRecordType []byte `json:"headerRecordType"`
	HeaderBankAbbreviation []byte `json:"headerBankAbbreviation"`
	SettlementDate []byte `json:"settlementDate"`
	HeaderTitle []byte `json:"headerTitle"`
	ReservedSpace []byte `json:"reservedSpace"`
}

type MsgBody struct {
	DetailRecordType []byte `json:"detailRecordType"`
	SendingBankCode []byte `json:"sendingBankCode"`
	PAN []byte `json:"pan"`
	TransfererBankCode []byte `json:"transfererBankCode"`
	TerminalID []byte `json:"terminalId"`
	ReceiptSequenceNumber []byte `json:"receiptSequenceNumber"`
	TransactionDate []byte `json:"transactionDate"`
	TransactionTime []byte `json:"transactionTime"`
	SwitchingDate []byte `json:"switchingDate"`
	TransactionType []byte `json:"transactionType"`
	Reserve []byte `json:"reserve"`
	FormAccountNumber []byte `json:"formAccountNumber"`
	ToAccountNumber []byte `json:"toAccountNumber"`
	Amount1 []byte `json:"amount1"`
	TermState []byte `json:"termState"`
	ResponseCode []byte `json:"responseCode"`
	TransfereeFee []byte `json:"transfereeFee"`
	TransfererFee []byte `json:"transfererFee"`
	SenderFee []byte `json:"senderFee"`
	ReceivingBankCode []byte `json:"receivingBankCode"`
	DRCRFlag []byte `json:"drcrFlag"`
	Reserve2 []byte `json:"reserve"`
	Channel []byte `json:"channel"`
	ProxyType []byte `json:"proxyType"`
	ProxyID []byte `json:"proxyId"`
	TypeOfSender []byte `json:"typeOfSender"`
	TypeOfReceiver []byte `json:"typeOfReceiver"`
	SenderID []byte `json:"senderId"`
	SenderProxyType []byte `json:"senderProxyType"`
	ReservedSpace []byte `json:"reservedSpace"`
	BillReference []byte `json:"billReference"`
	BillReference2 []byte `json:"billReference2"`
	BillReference3 []byte `json:"billReference3"`
	SenderAccountName []byte `json:"senderAccountName"`
	ReservedSpace2 []byte `json:"reservedSpace2"`
}

type MgsTrailer struct {
	TrailerRecordType []byte `json:"trailerRecordType"`
	TrailerFlag []byte `json:"trailerFlag"`
	TotalNumberOfRecords []byte `json:"totalNumberOfRecords"`
	ReservedSpace []byte `json:"reservedSpace"`
}