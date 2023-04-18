package constant

// Define common error codes
const (
	SYSPANIC    = "SY00000001"
	SYSTEMPANIC = "SY00000001" // System abnormal
	// SYSCODE1 = "SY00000002"

	// SYERRCONT = "SY00000004"

	REQPACKERR   = "SY00000006"
	REQUNPACKERR = "SY00000007"
	RSPPACKERR   = "SY00000008"
	RSPUNPACKERR = "SY00000009"

	PROXYREGFAILD = "SY00000010" // Register DTS Proxy failed. %v
	PROXYFAILD    = "SY00000011" // DTS proxy executed service failed, %v
	REMOTEFAILD   = "SY00000012" // Remote call %s failed, error: %v
	INVALIDERR    = "SY00000013" //
	NOTFOUNDTOPIC = "SY00000014" // Can't found destination topic. serviceKey: [%v]

	DASNOTFOUND        = "DA00000002"
	REMOTE_CALL_FAILED = "SY00000002"
	PK_CONFLICT        = "SY00000003" // 主键冲突
	DB_ERR             = "SY00000004" // 数据库连接失败
	DB_EXEC_FAILED     = "SY00000005"
)

const (
	ProductInActive = "0"
	ProductActive   = "1"
	ProductExpired  = "2"
	ProductDeleted  = "3"
	ProductUpdated  = "4"

	LimitTargetLoanAmount  = "L0" //L0-Loan Amount
	LimitTargetGracePeriod = "L1" //L1-Grace Period
	LimitTargetLoanPeriods = "L2" //L2-Loan Periods

	CustomerStatusNormal    = "0"
	CustomerStatusSuspended = "3"
	CustomerStatusPending   = "4"
	CustomerStatusCancel    = "9"

	FeeStatusNormal = "0"

	GuarantorMethodBorrower  = "1" //1: "Borrower's own credit guarantee",
	GuarantorMethodGuarantor = "2" //2: 'Guarantor',
	GuarantorMethodPledge    = "3" //3: 'Pledge',
	GuarantorMethodMortgage  = "4" //4: 'Mortgage',

	NoGuarantor   = "0"
	NeedGuarantor = "1"

	ContactTypeGuarantor  = "1" //1 - guarantor 2 - Credit card
	ContactTypeCreditCard = "2"

	ContactTypeHomePhone          = "01"
	ContactTypeWorkPhone          = "02"
	ContactTypePhoneOnlyWalletUse = "03"
	ContactTypeEmail              = "04"
	ContactTypeFax                = "05"
	ContactTypeWetChat            = "06"
	ContactTypeQQ                 = "07"
	ContactTypePhone              = "09"
	ContactTypeOther              = "99"

	ContactTypeEmergency    = "1"
	ContactTypeGuarantee    = "2"
	ContactTypeSpouse       = "3"
	ContactRelationSpouse   = "201"
	ContactRelationChildren = "202"
	ContactRelationParent   = "203"
	ContactRelationHimself  = "205"
	ContactRelationBrothor  = "206"

	LoanRequest             = "01" // 01-Request 02-Approved 03-Denied 10-Closed；
	LoanApproved            = "02"
	LoanDenied              = "03"
	LoanContractSignSuccess = "04"
	LoanLending             = "05"
	LoanSuccessfulLoan      = "06"
	LoanLoanFailure         = "09"
	LoanClosed              = "10"
	LoanWriteoff            = "11"

	OperateStatusPending  = "1" //1 - pending 2 - approved 3 - rejected
	OperateStatusApproved = "2"
	OperateStatusRejected = "3"

	ApplyStatusPending  = "1" //1 - pending 2 - approved 3 - rejected
	ApplyStatusApproved = "2"
	ApplyStatusRejected = "3"

	ApplyTypeLoanApplication             = "1" //1 - loan application 2 - modification application 3 - repayment application 4 - secondary loan application
	ApplyTypeLoanModificationApplication = "2"
	ApplyTypeLoanRepaymentApplication    = "3"
	ApplyTypeScecondaryLoanApplication   = "4"
	ApplyTypeRestructuring               = "5"
	ApplyTypeWriteOff                    = "6"
	ApplyTypeRefinance                   = "7"
	ApplyTypeSkipAndReduce               = "8"
	ApplyTypePrepaymentDiscount          = "9"
	ApplyTypeCancelTransaction           = "10"
	ApplyTypeCloseRepayment              = "11"

	ReceiptStatusNormal         = "01" // 01-正常 02-逾期 03-减值 04-部分核销 05-核销 06-结清 07-减免结清
	ReceiptStatusOverdue        = "02"
	ReceiptStatusImpairment     = "03"
	ReceiptStatusPartOfWriteOff = "04"
	ReceiptStatusWriteOff       = "05"
	ReceiptStatusSettle         = "06"
	ReceiptStatusReliefSettle   = "07"

	LoanApprovedAgree       = "01" // 01-同意;02-退回;03-终止;04-弃权;05-拒绝
	LoanApprovedReturn      = "02"
	LoanApprovedTermination = "03"
	LoanApprovedAbstain     = "04"
	LoanApprovedReject      = "05"

	ApproveViewAgree  = "2"
	ApproveViewReject = "3"

	DayEndStatus    = "F"
	DayOnlineStatus = "N"

	Modify_LoanDeined          = "0" // 0-All// 0-All
	Modify_Fee                 = "1" // 1-Fee
	Modify_InterestRate        = "2" // 2-Interest Rate
	Modify_PeriodDistibution   = "3" // 3-Period Distributi
	Modify_LoanExtension       = "4" // 4-Loan Extension
	Modify_BillRule            = "5" // bill rule information
	Modify_TransactionAccount  = "6"
	Modify_Guarantee           = "7"
	Modify_CreditLine          = "8"
	Modify_CreditLineStatus    = "9"
	LoanApprovedStatus_Pending = "1" // 1-待审批
	LoanApprovedStatus_Agree   = "2" // 2-审批通过
	LoanApprovedStatus_Reject  = "3" // 3-审批拒绝"

	LoanApplyContactType = "30011"
	//0-正常;1-逾期;2-减值;3-核销;4-结清
	AccountStatusNormal     = "0"
	AccountStatusOverdue    = "1"
	AccountStatusImpairment = "2"
	AccountStatusWirteOff   = "3"
	AccountStatusClosed     = "4"
	AccountStatusPreClosed  = "5"
	WriteOffBeApproved      = "01"
	WriteOffApproved        = "02"
	WriteOffRejected        = "03"
	WriteOffCancel          = "04"

	LoanContractWriteoff = "05"

	ApproveWriteOffApproved = "2" //approved
	ApproveWriteOffRejected = "3" //rejected
	//计息标识
	InterestCalcTier   = "0" //分层计息
	InterestCalcNoTier = "1" //不分层计息
	InterestCalcNoInt  = "2" //不计息
	//分层方法
	TierFullAmount   = "0" //全额累进
	TierExcessAmount = "1" //超额累进
	TierLimitAmount  = "2" //限额靠档
	TierFullTime     = "3"
	TierExcessTime   = "4"

	//tier type
	TierByAmount   = "1"
	TierByTime     = "2"
	TierByAum      = "3"
	TierByCustomer = "4"

	//tier cal method
	FullProgressive   = "0"
	ExcessProgressive = "1"

	//利率使用标识
	InterestUseDefaultLvl = "0" //默认
	InterestUseRateNum    = "1" //利率编号
	InterestUseFix        = "2" //固定利率
	//利率浮动类型
	FloatTypePercent = "0" //百分比浮动
	FloatTypeFix     = "1" //
	//浮动方向
	FloatDirNo   = "0" //不浮动
	FloatDirUp   = "1" //向上浮动
	FloatDirDown = "2" //向下浮动

	BehaviorRepayment          = "01"
	BehaviorCloseRepayment     = "02"
	BehaviorRefinanceRepayment = "03"
)
const (
	CustomerType   = "CU0006"
	IDType         = "idType"
	CustomerStatus = "customerStatus"
	Gender         = "gender"
	Nationality    = "country"
	MaritalStatus  = "maritalStatus"
	Occupation     = "jobList"
	Education      = "educationLevel"
	CustomerGrade  = "customerGrade"
	LangIdentifyEn = "en"
	LangIdentifyFr = "fr"
	LangIdentifySp = "sp"
	LangIdentifyTh = "th"
	ReceiptStatus  = "receiptStatus"

	CSV = "CSV"

	ProductStatus           = "CM0048"
	ProductType             = "CM0049"
	ProductCategory         = "productCategory"
	RepaymentMethod         = "repaymentMethodOption"
	EarlyRepaymentAllowance = "CM0052"
	GuaranteeMethod         = "CM0053"
	CurrencyCode            = "CM0001"
	IL003                   = "IL003"
	IL002                   = "IL002"
	RepayCycle              = "repayCycle"
	RepayType               = "repayType"

	LoanStatus          = "IL0013"
	DisbursementChannel = "IL0012"
	ArrangementPurpose  = "IL0010"
	BillPaymentType     = "BillPaymentType"
	Cycle               = "repaymentFrequencyOption"
)
const (
	RepayMethodEqualInstallment                   = "01" //01-Equal Installment 等额本息
	RepayMethodEqualPrincipal                     = "02" //02-Equal Principal等额本金
	RepayMethodPrincipalAndInterestOnMaturityDate = "03" //03-Principal and Interest Installment on Maturity Date
	RepayMethodPrincipalOnMaturityDate            = "04" //04-Interest Installment but Principal on Maturity Date 到期还本周期还息(分期付息到期还本（先息后本))
	RepayMethodEqualPrincipalAndInterest          = "05" //05-equal principal and interest 等本等息（每期还本还息还款额都相等，每期计息的本金为贷款总本金）
	RepayMethodIndependentBillDate                = "06" //06-Independent Bill Date 账单日还款
	RepayFrequencyDoubleWeek                      = "03"
	RepayFrequencyMonthly                         = "04"
)
const (
	CancelTransTypeRepayment    = "03" //03- repayment
	CancelTransTypeDisbursement = "04" //04- disbursement
)

const (
	LevelFlagNo  = "0" //"0-no
	LevelFlagYes = "1" //1-yes"
	Level0       = "0"
	Level1       = "1"
	Level2       = "2"
	Level3       = "3"
	Level4       = "4"
	Level5       = "5"

	SerialCode_ACCT = "ACCT"
	AC900001        = "AC900001"

	TieredInterestTypeDefaultInterest = "0"
	TieredInterestTypeBaseInterest    = "1"
	TieredInterestTypeFixedInterest   = "2"
)

const (
	InterestNormal  = "0" //0-normal loan
	InterestOverdue = "1" //1-overdue loan

)
const (
	ProductTypeMicroLoan     = "01" //01-Micro Loan
	ProductTypeRevolvingLoan = "02" //02-Revolving Loan
	ProductTypeAutoLoan      = "04" //04-AutoLoan
	ProductTypeBNPL          = "03" //03-bnpl
	RevolvingFlagNo          = "0"  //0-Not revolving
	RevolvingFlagYes         = "1"  //1-revolving

	//FeeTypeStampDuty     = "0" //0-Stamp Duty,
	//FeeTypeFrontEndFee   = "1" //1-Front End Fee,
	//FeeTypeCollectionFee = "2" //2-Collection Fee,
	//FeeTypePrepaymentFee = "3" //3-Prepayment Fee
	FeeTypeStampDuty            = "F001" //F001-Stamp Duty,
	FeeTypeFrontEndFee          = "F002" //F002-Front End Fee,
	FeeTypeCollectionFee        = "F003" //F003-Collection Fee,
	FeeTypePrepaymentFee        = "F004" //F004-Prepayment Fee
	FeeTypePPI                  = "F005" //F005-PPI
	FeeTypeDeferredRepaymentFee = "F006" //F005-Deferred Repayment Fee
	FeeTypeTransferFee          = "5"    //5-Transfer Fee
)
const (
	GlsCreate      = "GlsCreate"
	GlsCreateDxc   = "GlsCreateDxc"
	GlsUpdate      = "GlsUpdate"
	GlsUpdateDxc   = "GlsUpdateDxc"
	GlsRemove      = "GlsRemove"
	GlsRemoveDxc   = "GlsRemoveDxc"
	GlsExist       = "GlsExist"
	GlsSuList      = "GlsSuList"
	GlsSpecChanged = "GlsSpecChanged"
)
const (
	GLS_TYPE_CMM   = "CMM"
	GLS_TYPE_SRV   = "SRV"
	GLS_ID_COMMON  = "common"
	GLS_TYPE_LNM   = "LNM"    // 贷款切片类型
	GLS_LNM_COMMON = "ilcomm" // 贷款公共ID
	GLS_TYPE_SCM   = "SCM"    // 存款的公共
	GTS_TOPIC_AGM  = "AGM"    // 合约号
	GTS_TOPIC_MED  = "MED"    // 介质号
	GLS_TYPE_CON   = "CON"
	GLS_TYPE_ACC   = "ACC"
	GLS_TYPE_ACT   = "ACT"
	GLS_TYPE_CUS   = "CUS"
	GLS_TYPE_ACG   = "ACG"
	GLS_TYPE_PRD   = "PRD"
	GLS_TYPE_PHN   = "PHN"
	GLS_TYPE_IDN   = "IDN"
	GLS_TYPE_ODR   = "ODR" // FOR ORDER
	GLS_TYPE_MER   = "MER" // FOR merchant
	GLS_TYPE_PPS   = "PPS" // For PP service
	GLS_TYPE_HUI   = "HUI" // For HU id
	GLS_TYPE_ORG   = "ORG" // For organization
	GLS_TYPE_HCD   = "HCD" // For Hcode
	GLS_TYPE_WID   = "WID" // For wallet WalletId
	GLS_TYPE_CNT   = "CNT" // 合同号切片类型
	GLS_TYPE_DBT   = "DBT" // 借据号切片类型
	GLS_TYPE_AGM   = "AGM" // 存款的-合约号切片类型
	GLS_TYPE_MED   = "MED"
	GLS_TYPE_QTA   = "QTA" //quota credit Apply serial number
	GLS_TYPE_MAS   = "MAS" //makeLoanApplySerialNum放款申请流水号
)

const (
	GlsCommonKey           = "common"
	GlsLoanCommonKey       = "ilcomm" // 贷款公共ID
	GlsTypeCommon          = "CMM"
	GlsTypeLoanCommon      = "LNM" // 贷款切片类型
	GlsTypeSavingCommon    = "SCM" // 存款的公共
	GlsTypeContract        = "CON"
	GlsTypeAccount         = "ACC"
	GlsTypeAccounting      = "ACT"
	GlsTypeCustomer        = "CUS"
	GlsTypeProduct         = "PRD"
	GlsTypePhone           = "PHN"
	GlsTypeId              = "IDN"
	GlsTypeOrganize        = "ORG" // For organization
	GlsTypeWalletId        = "WID" // For wallet WalletId
	GlsTypeContractNum     = "CNT" // 合同号切片类型
	GlsTypeDebtAcknowledge = "DBT" // 借据号切片类型
	GlsTypeAgreementId     = "AGM" // 存款的-合约号切片类型
	GlsTypeMediumNum       = "MED" // medium number
	GlsTypeMsg             = "MSG" // fast payment message use
)

const (
	GlsCreateResult_Exits = "664518"
)

// Define common header
const (
	DTS_TOPIC_TYPE = "DTS"
	TRN_TOPIC_TYPE = "TRN"
	KEY_TYPE_RSA   = "RSA"
	KEY_TYPE_AES   = "AES"
	DLS_TYPE_CMM   = "CMM"
	DLS_ID_COMMON  = "common"
	DLS_TYPE_LNM   = "LNM"    // 贷款切片类型
	DLS_LNM_COMMON = "ilcomm" // 贷款公共ID
	DLS_TYPE_SCM   = "SCM"    // 存款的公共
	DTS_TOPIC_AGM  = "AGM"    // 合约号
	DTS_TOPIC_MED  = "MED"    // 介质号
	DLS_TYPE_CON   = "CON"
	DLS_TYPE_ACC   = "ACC"
	DLS_TYPE_ACT   = "ACT"
	DLS_TYPE_CUS   = "CUS"
	DLS_TYPE_ACG   = "ACG"
	DLS_TYPE_PRD   = "PRD"
	DLS_TYPE_PHN   = "PHN"
	DLS_TYPE_IDN   = "IDN"
	DLS_TYPE_ODR   = "ODR" // FOR ORDER
	DLS_TYPE_MER   = "MER" // FOR merchant
	DLS_TYPE_PPS   = "PPS" // For PP service
	DLS_TYPE_HUI   = "HUI" // For HU id
	DLS_TYPE_ORG   = "ORG" // For organization
	DLS_TYPE_HCD   = "HCD" // For Hcode
	DLS_TYPE_WID   = "WID" // For wallet WalletId
	DLS_TYPE_CNT   = "CNT" // 合同号切片类型
	DLS_TYPE_DBT   = "DBT" // 借据号切片类型
	DLS_TYPE_AGM   = "AGM" // 存款的-合约号切片类型
	DLS_TYPE_MED   = "MED"
	DLS_TYPE_QTA   = "QTA" // 额度申请编号
)

// Request is the struct which describing the format from client
// the whole struct should be stored in SMF body
type Request struct {
	// Param is the param in HTTP request
	// especially useful in GET request
	Param map[string]string `json:"param"`
	// Header is the header will be put on HTTP request
	Header map[string]string `json:"header"`
	// Body is the body will be put on HTTP request
	Body []byte `json:"body"`
}

// Response is the struct which will be returned to client
// the whole struct will be stored in SMF body
type Response struct {
	// Status represent the HTTP status code, eg 200,401
	Status int `json:"status"`
	// Header represent HTTP header return from server
	Header map[string]interface{} `json:"header"`
	// Body represent HTTP body return from server
	Body []byte `json:"body"`
}

const HTTPSTATUSCODE = "httpStatusCode"
const (
	TIME_STAMP        = "2006-01-02 15:04:05"
	DATE_DASH_FORMAT  = "2006-01-02"
	DATEFORMAT        = "20060102"
	TIME_COLON_FORMAT = "15:04:05"
)

const (
	ContentType = "Content-Type"
	JsonContent = "application/json; charset=UTF-8"
)
const (
	SuccessCode = "0"
	SuccessMsg  = "success"
)

const (
	GLOBALBIZSEQNO = "GlobalBizSeqNo" // Global unique business serial number
	SrcBizSeqNo    = "SrcBizSeqNo"
)
const (
	SystemId_SV = "SV"
	SystemId_LN = "LN"
)
const (
	NotifyType_01        = "01" //Loan Account Reminder
	NotifyType_02        = "02" //Loan Account Late Notification
	NotifyType_03        = "03" //Account Statement
	NotifyType_Refinance = "06" //Refinance
	NotifyType_Skip      = "07" //Skip Payment & Reduce Installment
	NotifyType_Discount  = "08" //Prepayment Discount

)

const (
	QuotaFozen   = "2"
	QuotaExpired = "3"
	QuotaActive  = "1"
)

const (
	LoanApplyStatus_01 = "01" //01-放款申请已受理
	LoanApplyStatus_03 = "03" //03-放款审批拒绝
	LoanApplyStatus_10 = "10" //10-closed
	LoanApplyStatus_09 = "09" //09-放款失败

)

const (
	MaritalStatus_Single   = "10"
	MaritalStatus_Married  = "20"
	MaritalStatus_Widowed  = "30"
	MaritalStatus_Divorced = "40"
)

const (
	MicroLoan     = "01"
	RevolvingLoan = "02"
	AutoLoan      = "04"
)

const (
	Percentage = 100
)

const (
	RepayStatusNotPay = "0"
	RepayStatusPaid   = "1"
)

const (
	FeeDiscountYes = 1
)

const (
	//wallet transaction_type
	WalletTopUp       = 101
	WalletWithdraw    = 102
	WalletSender      = 201
	WalletReceiver    = 202
	WalletOnlinePay   = 203
	WalletPayClearing = 204
)

//for reversal transaction
const (
	TransAccountTypeDebit  = "D"
	TransAccountTypeCredit = "C"
)

//const (
//	BehaviorRepayment          = "01"
//	BehaviorCloseRepayment     = "02"
//	BehaviorRefinanceRepayment = "03"
//)

const (
	AmountOfPrincipal               = "P001"
	AmountOfPrincipalOverdue        = "P002"
	AmountOfPrincipalCurrent        = "P003"
	AmountOfPrincipalAdvance        = "P004"
	AmountOfNormalInterest          = "I001"
	AmountOfNormalInterestOverdue   = "I101"
	AmountOfNormalInterestCurrent   = "I201"
	AmountOfNormalInterestAdvance   = "I301"
	AmountOfNormalCompoundInterest  = "I002"
	AmountOfPenaltyInterest         = "I003"
	AmountOfPenaltyCompoundInterest = "I004"
	AmountOfCollectionFee           = "F003" //F003-Collection Fee,
	AmountOfPrepaymentFee           = "F004" //F004-Prepayment Fee
	AmountOfPPI                     = "F005" //F005-PPI
	AmountOfPPIOverdue              = "F105"
	AmountOfPPICurrent              = "F205"
	AmountOfPPIAdvance              = "F305"
	AmountOfDeferredRepaymentFee    = "F006" //F006-Deferred Repayment Fee =
)

const (
	ChangeTypeSharedtimes   = "1"
	ChangeTypeCustomerGrade = "2"
)
const (
	TargetTypeCustomer    = "0"
	ChannelTypeEmail      = "01"
	TargetKeyTypeCustomer = "0"
)

const (
	TccStatusSuc = 0
	TccStatusErr = 1
)
