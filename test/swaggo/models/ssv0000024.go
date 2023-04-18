//Version: v0.0.1
package models

type SV000024I struct {
	DepositProofType          string          `json:"depositProofType" validate:"required"`
	CreateDate                string          `json:"createDate" validate:"required"`
	TotalProofAmount          float64         `json:"totalProofAmount" validate:"required"`
	TimeProofDate             string          `json:"timeProofDate" validate:"required"`
	ExpiryProofDate           string          `json:"expiryProofDate" `
	RecipientName             string          `json:"recipientName" validate:"required"`
	DeliveryMethod            string          `json:"deliveryMethod" validate:"required"`
	Email                     string          `json:"email" `
	PostalCode                string          `json:"postalCode" `
	NationCode                string          `json:"nationCode" `
	Address1                  string          `json:"address1" `
	Address2                  string          `json:"address2" `
	Address3                  string          `json:"address3" `
	Address                   string          `json:"address" `
	ReceiverName              string          `json:"receiverName" `
	AgreementList             []AgreementInfo `json:"agreementList" `
	DeductionFlag             string          `json:"deductionFlag" validate:"required"`
	FeeDeductionMediumNumber  string          `json:"feeDeductionMediumNumber" `
	FeeDeductionMediumType    string          `json:"feeDeductionMediumType" `
	FeeDeductionCurrency      string          `json:"feeDeductionCurrency" validate:"required"`
	FeeDeductionCashTransFlag string          `json:"feeDeductionCashTransFlag" `
	FeeDeductionChannel       string          `json:"feeDeductionChannel" `
	FeeUsageCode              string          `json:"feeUsageCode" `
	FeeDedWithdrawMethod      string          `json:"feeDedWithdrawMethod" `
	FeeNeedPasswordFlag       string          `json:"feeNeedPasswordFlag" `
	FeeAccountPassword        string          `json:"feeAccountPassword" `
	ChargeList                []ChargeInfo    `json:"chargeList" `
	CustomerID                string          `json:"customerId" `
	AccountName               string          `json:"accountName" `
}

type SV000024O struct {
	DepositProofNumber string `json:"depositProofNumber"`
}

type AgreementInfo struct {
	AgreementId          string  `json:"agreementId" validate:"required"`
	AgreementType        string  `json:"agreementType" validate:"required"`
	Currency             string  `json:"currency" validate:"required"`
	CashTransFlag        string  `json:"cashTransFlag" validate:"required"`
	ValueDate            string  `json:"valueDate" `
	ExpireDate           string  `json:"expireDate" `
	MediumType           string  `json:"mediumType" validate:"required"`
	MediumNumber         string  `json:"mediumNumber" validate:"required"`
	AccountName          string  `json:"accountName" `
	AccountingNumber     string  `json:"accountingNumber"`
	AccountBalance       float64 `json:"accountBalance" `
	AvaBalance           float64 `json:"avaBalance" `
	AccountStatus        string  `json:"accountStatus" validate:"required"`
	FreezeType           string  `json:"freezeType" validate:"required"`
	FreezeBusinessNumber string  `json:"freezeBusiness"`
	FreezeDocNumber      string  `json:"freezeDocNumber"`
	DepcreFlag           string  `json:"depcreFlag" validate:"required"`
	ProofAmount          float64 `json:"proofAmount" validate:"required"`
}

type ChargeInfo struct {
	FeeAbstractCode     string  `json:"feeAbstractCode" `
	BeFeeDeductionFlag  string  `json:"beFeeDeductionFlag" `
	FeeTobeDeducted     float64 `json:"feeTobeDeducted" `
	FeePreferential     float64 `json:"feePreferential" `
	FeeActuallyDeducted float64 `json:"feeActuallyDeducted" `
	StrategyId          string  `json:"strategyId" `
	FeeRateId           string  `json:"feeRateId" `
	FeeAmortizationFlag string  `json:"feeAmortizationFlag" `
	AmortizePeriodType  int     `json:"amortizePeriodType" `
}

func (*SV000024I) GetServiceKey() string {
	return "SV000024"
}
