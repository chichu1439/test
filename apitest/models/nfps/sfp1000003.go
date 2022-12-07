package models

type FP100003I struct {
	MessHeader                string `json:"messHeader"`
	Bitmap                    []byte `json:"bitmap"`
	PrimaryAccountNumber      string `json:"primaryAccountNumber"`
	ProcessingCode            int    `json:"processingCode"`
	TransactionAmount         int    `json:"transactionAmount"`
	SettlementAmount          int    `json:"settlementAmount"`
	TransmissionDateTime      int    `json:"transmissionDateTime"`
	ConversionRateSettlement  string `json:"conversionRateSettlement"`
	SysTraceAuditNo           int    `json:"sysTraceAuditNo"`
	LocalTransactionTime      int    `json:"localTransactionTime"`
	LocalTransactionDate      int    `json:"localTransactionDate"`
	SettlementDate            int    `json:"settlementDate"`
	ConversionDate            int    `json:"conversionDate"`
	MerchantType              int    `json:"merchantType"`
	POSEntryCode              string `json:"pOSEntryCode"`
	CardSequenceNo            int    `json:"cardSequenceNo"`
	AmountAcquirerFee         string `json:"amountAcquirerFee"`
	AmountIssuerFee           string `json:"amountIssuerFee"`
	AcquirerId                string `json:"acquirerId"`
	ForwardingId              string `json:"forwardingId"`
	Track2Data                string `json:"track2Data"`
	RetrievalReferenceNo      string `json:"retrievalReferenceNo"`
	AuthNo                    string `json:"authNo"`
	ResponseCode              string `json:"responseCode"`
	CardAcceptorTerminalId    string `json:"cardAcceptorTerminalId"`
	CardAcquirerId            string `json:"cardAcquirerId"`
	CardAcceptorNameLocation  string `json:"cardAcceptorNameLocation"`
	AdditionalDataPrivate     string `json:"additionalDataPrivate"`
	CurrencyCodeOfTransaction int    `json:"currencyCodeOfTransaction"`
	CurrencyCodeOfSettlement  string `json:"currencyCodeOfSettlement"`
	PINBlock                  string `json:"pINBlock"`
	AdditionalAmounts         string `json:"additionalAmounts"`
	EMVData                   string `json:"eMVData"`
	ReversalCode              string `json:"reversalCode"`
	ITMXAdditionalPOSInfo     string `json:"iTMXAdditionalPOSInfo"`
	NetworkCode               string `json:"networkCode"`
	OrigData                  string `json:"origData"`
	ReplacementAmount         string `json:"replacementAmount"`
	FromAccount               string `json:"fromAccount"`
	ToAccount                 string `json:"toAccount"`
	NetworkManagementInfo     string `json:"networkManagementInfo"`
	MessAuthCodeField         string `json:"messAuthCodeField"`
}
type FP100003O struct {
	Buf []byte `json:"buf"`
	Str string `json:"str"`
}

func (*FP100003I) GetServiceKey() string {
	return "FP100003"
}
