// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v09

import (
	"encoding/xml"

	"apitest/models/nfps/iso20022/common"
	"apitest/util"
)

type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier     `xml:"IBAN"`
	Othr GenericAccountIdentification1 `xml:"Othr"`
}

func (r AccountIdentification4Choice) Validate() error {
	return util.Validate(&r)
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                   `xml:"Prtry"`
}

func (r AccountSchemeName1Choice) Validate() error {
	return util.Validate(&r)
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                             `xml:",chardata"`
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

func (r ActiveOrHistoricCurrencyAndAmount) Validate() error {
	return util.Validate(&r)
}

type AddressType3Choice struct {
	Cd    common.AddressType2Code `xml:"Cd"`
	Prtry GenericIdentification30 `xml:"Prtry"`
}

func (r AddressType3Choice) Validate() error {
	return util.Validate(&r)
}

type AmendmentInformationDetails13 struct {
	OrgnlMndtId      *common.Max35Text                             `xml:"OrgnlMndtId,omitempty" json:",omitempty"`
	OrgnlCdtrSchmeId *PartyIdentification135                       `xml:"OrgnlCdtrSchmeId,omitempty" json:",omitempty"`
	OrgnlCdtrAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlCdtrAgt,omitempty" json:",omitempty"`
	OrgnlCdtrAgtAcct *CashAccount38                                `xml:"OrgnlCdtrAgtAcct,omitempty" json:",omitempty"`
	OrgnlDbtr        *PartyIdentification135                       `xml:"OrgnlDbtr,omitempty" json:",omitempty"`
	OrgnlDbtrAcct    *CashAccount38                                `xml:"OrgnlDbtrAcct,omitempty" json:",omitempty"`
	OrgnlDbtrAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlDbtrAgt,omitempty" json:",omitempty"`
	OrgnlDbtrAgtAcct *CashAccount38                                `xml:"OrgnlDbtrAgtAcct,omitempty" json:",omitempty"`
	OrgnlFnlColltnDt *common.ISODate                               `xml:"OrgnlFnlColltnDt,omitempty" json:",omitempty"`
	OrgnlFrqcy       *Frequency36Choice                            `xml:"OrgnlFrqcy,omitempty" json:",omitempty"`
	OrgnlRsn         *MandateSetupReason1Choice                    `xml:"OrgnlRsn,omitempty" json:",omitempty"`
	OrgnlTrckgDays   *common.Exact2NumericText                     `xml:"OrgnlTrckgDays,omitempty" json:",omitempty"`
}

func (r AmendmentInformationDetails13) Validate() error {
	return util.Validate(&r)
}

type AmountType4Choice struct {
	InstdAmt ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`
	EqvtAmt  EquivalentAmount2                 `xml:"EqvtAmt"`
}

func (r AmountType4Choice) Validate() error {
	return util.Validate(&r)
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"FinInstnId"`
	BrnchId    *BranchData3                         `xml:"BrnchId,omitempty" json:",omitempty"`
}

func (r BranchAndFinancialInstitutionIdentification6) Validate() error {
	return util.Validate(&r)
}

type BranchData3 struct {
	Id      *common.Max35Text     `xml:"Id,omitempty" json:",omitempty"`
	LEI     *common.LEIIdentifier `xml:"LEI,omitempty" json:",omitempty"`
	Nm      *common.Max140Text    `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr *PostalAddress24      `xml:"PstlAdr,omitempty" json:",omitempty"`
}

func (r BranchData3) Validate() error {
	return util.Validate(&r)
}

type CancellationReason33Choice struct {
	Cd    ExternalCancellationReason1Code `xml:"Cd"`
	Prtry common.Max35Text                `xml:"Prtry"`
}

func (r CancellationReason33Choice) Validate() error {
	return util.Validate(&r)
}

type Case5 struct {
	Id             common.Max35Text `xml:"Id"`
	Cretr          Party40Choice    `xml:"Cretr"`
	ReopCaseIndctn bool             `xml:"ReopCaseIndctn,omitempty" json:",omitempty"`
}

func (r Case5) Validate() error {
	return util.Validate(&r)
}

type CaseAssignment5 struct {
	Id      common.Max35Text   `xml:"Id"`
	Assgnr  Party40Choice      `xml:"Assgnr"`
	Assgne  Party40Choice      `xml:"Assgne"`
	CreDtTm common.ISODateTime `xml:"CreDtTm"`
}

func (r CaseAssignment5) Validate() error {
	return util.Validate(&r)
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice         `xml:"Id"`
	Tp   *CashAccountType2Choice              `xml:"Tp,omitempty" json:",omitempty"`
	Ccy  *common.ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty" json:",omitempty"`
	Nm   *common.Max70Text                    `xml:"Nm,omitempty" json:",omitempty"`
	Prxy *ProxyAccountIdentification1         `xml:"Prxy,omitempty" json:",omitempty"`
}

func (r CashAccount38) Validate() error {
	return util.Validate(&r)
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

func (r CashAccountType2Choice) Validate() error {
	return util.Validate(&r)
}

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

func (r CategoryPurpose1Choice) Validate() error {
	return util.Validate(&r)
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                          `xml:"Prtry"`
}

func (r ClearingSystemIdentification2Choice) Validate() error {
	return util.Validate(&r)
}

type ClearingSystemIdentification3Choice struct {
	Cd    ExternalCashClearingSystem1Code `xml:"Cd"`
	Prtry common.Max35Text                `xml:"Prtry"`
}

func (r ClearingSystemIdentification3Choice) Validate() error {
	return util.Validate(&r)
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty" json:",omitempty"`
	MmbId    common.Max35Text                     `xml:"MmbId"`
}

func (r ClearingSystemMemberIdentification2) Validate() error {
	return util.Validate(&r)
}

type Contact4 struct {
	NmPrfx    *common.NamePrefix2Code      `xml:"NmPrfx,omitempty" json:",omitempty"`
	Nm        *common.Max140Text           `xml:"Nm,omitempty" json:",omitempty"`
	PhneNb    *common.PhoneNumber          `xml:"PhneNb,omitempty" json:",omitempty"`
	MobNb     *common.PhoneNumber          `xml:"MobNb,omitempty" json:",omitempty"`
	FaxNb     *common.PhoneNumber          `xml:"FaxNb,omitempty" json:",omitempty"`
	EmailAdr  *common.Max2048Text          `xml:"EmailAdr,omitempty" json:",omitempty"`
	EmailPurp *common.Max35Text            `xml:"EmailPurp,omitempty" json:",omitempty"`
	JobTitl   *common.Max35Text            `xml:"JobTitl,omitempty" json:",omitempty"`
	Rspnsblty *common.Max35Text            `xml:"Rspnsblty,omitempty" json:",omitempty"`
	Dept      *common.Max70Text            `xml:"Dept,omitempty" json:",omitempty"`
	Othr      []OtherContact1              `xml:"Othr,omitempty" json:",omitempty"`
	PrefrdMtd *PreferredContactMethod1Code `xml:"PrefrdMtd,omitempty" json:",omitempty"`
}

func (r Contact4) Validate() error {
	return util.Validate(&r)
}

type ControlData1 struct {
	NbOfTxs *common.Max15NumericText `xml:"NbOfTxs,omitempty" json:",omitempty"`
	CtrlSum float64                  `xml:"CtrlSum,omitempty" json:",omitempty"`
}

func (r ControlData1) Validate() error {
	return util.Validate(&r)
}

type CreditTransferMandateData1 struct {
	MndtId       *common.Max35Text          `xml:"MndtId,omitempty" json:",omitempty"`
	Tp           *MandateTypeInformation2   `xml:"Tp,omitempty" json:",omitempty"`
	DtOfSgntr    *common.ISODate            `xml:"DtOfSgntr,omitempty" json:",omitempty"`
	DtOfVrfctn   *common.ISODateTime        `xml:"DtOfVrfctn,omitempty" json:",omitempty"`
	ElctrncSgntr *common.Max10KBinary       `xml:"ElctrncSgntr,omitempty" json:",omitempty"`
	FrstPmtDt    *common.ISODate            `xml:"FrstPmtDt,omitempty" json:",omitempty"`
	FnlPmtDt     *common.ISODate            `xml:"FnlPmtDt,omitempty" json:",omitempty"`
	Frqcy        *Frequency36Choice         `xml:"Frqcy,omitempty" json:",omitempty"`
	Rsn          *MandateSetupReason1Choice `xml:"Rsn,omitempty" json:",omitempty"`
}

func (r CreditTransferMandateData1) Validate() error {
	return util.Validate(&r)
}

type CreditorReferenceInformation2 struct {
	Tp  *CreditorReferenceType2 `xml:"Tp,omitempty" json:",omitempty"`
	Ref *common.Max35Text       `xml:"Ref,omitempty" json:",omitempty"`
}

func (r CreditorReferenceInformation2) Validate() error {
	return util.Validate(&r)
}

type CreditorReferenceType1Choice struct {
	Cd    DocumentType3Code `xml:"Cd"`
	Prtry common.Max35Text  `xml:"Prtry"`
}

func (r CreditorReferenceType1Choice) Validate() error {
	return util.Validate(&r)
}

type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"CdOrPrtry"`
	Issr      *common.Max35Text            `xml:"Issr,omitempty" json:",omitempty"`
}

func (r CreditorReferenceType2) Validate() error {
	return util.Validate(&r)
}

type CustomerPaymentCancellationRequestV09 struct {
	XMLName     xml.Name                  `xml:"CstmrPmtCxlReq"`
	Assgnmt     CaseAssignment5           `xml:"Assgnmt"`
	Case        *Case5                    `xml:"Case,omitempty" json:",omitempty"`
	CtrlData    *ControlData1             `xml:"CtrlData,omitempty" json:",omitempty"`
	Undrlyg     []UnderlyingTransaction27 `xml:"Undrlyg" json:",omitempty"`
	SplmtryData []SupplementaryData1      `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r CustomerPaymentCancellationRequestV09) Validate() error {
	return util.Validate(&r)
}

type DateAndDateTime2Choice struct {
	Dt   *common.ISODate     `xml:"Dt,omitempty" json:",omitempty"`
	DtTm *common.ISODateTime `xml:"DtTm,omitempty" json:",omitempty"`
}

func (r DateAndDateTime2Choice) Validate() error {
	return util.Validate(&r)
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     common.ISODate     `xml:"BirthDt"`
	PrvcOfBirth *common.Max35Text  `xml:"PrvcOfBirth,omitempty" json:",omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth"`
}

func (r DateAndPlaceOfBirth1) Validate() error {
	return util.Validate(&r)
}

type DatePeriod2 struct {
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt"`
}

func (r DatePeriod2) Validate() error {
	return util.Validate(&r)
}

type DiscountAmountAndType1 struct {
	Tp  *DiscountAmountType1Choice        `xml:"Tp,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

func (r DiscountAmountAndType1) Validate() error {
	return util.Validate(&r)
}

type DiscountAmountType1Choice struct {
	Cd    ExternalDiscountAmountType1Code `xml:"Cd"`
	Prtry common.Max35Text                `xml:"Prtry"`
}

func (r DiscountAmountType1Choice) Validate() error {
	return util.Validate(&r)
}

type DocumentLineIdentification1 struct {
	Tp     *DocumentLineType1 `xml:"Tp,omitempty" json:",omitempty"`
	Nb     *common.Max35Text  `xml:"Nb,omitempty" json:",omitempty"`
	RltdDt *common.ISODate    `xml:"RltdDt,omitempty" json:",omitempty"`
}

func (r DocumentLineIdentification1) Validate() error {
	return util.Validate(&r)
}

type DocumentLineInformation1 struct {
	Id   []DocumentLineIdentification1 `xml:"Id" json:",omitempty"`
	Desc *common.Max2048Text           `xml:"Desc,omitempty" json:",omitempty"`
	Amt  *RemittanceAmount3            `xml:"Amt,omitempty" json:",omitempty"`
}

func (r DocumentLineInformation1) Validate() error {
	return util.Validate(&r)
}

type DocumentLineType1 struct {
	CdOrPrtry DocumentLineType1Choice `xml:"CdOrPrtry"`
	Issr      *common.Max35Text       `xml:"Issr,omitempty" json:",omitempty"`
}

func (r DocumentLineType1) Validate() error {
	return util.Validate(&r)
}

type DocumentLineType1Choice struct {
	Cd    ExternalDocumentLineType1Code `xml:"Cd"`
	Prtry common.Max35Text              `xml:"Prtry"`
}

func (r DocumentLineType1Choice) Validate() error {
	return util.Validate(&r)
}

type EquivalentAmount2 struct {
	Amt      ActiveOrHistoricCurrencyAndAmount   `xml:"Amt"`
	CcyOfTrf common.ActiveOrHistoricCurrencyCode `xml:"CcyOfTrf"`
}

func (r EquivalentAmount2) Validate() error {
	return util.Validate(&r)
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                                `xml:"Prtry"`
}

func (r FinancialIdentificationSchemeName1Choice) Validate() error {
	return util.Validate(&r)
}

type FinancialInstitutionIdentification18 struct {
	BICFI       *common.BICFIDec2014Identifier       `xml:"BICFI,omitempty" json:",omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:",omitempty"`
	LEI         *common.LEIIdentifier                `xml:"LEI,omitempty" json:",omitempty"`
	Nm          *common.Max140Text                   `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr     *PostalAddress24                     `xml:"PstlAdr,omitempty" json:",omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr,omitempty" json:",omitempty"`
}

func (r FinancialInstitutionIdentification18) Validate() error {
	return util.Validate(&r)
}

type Frequency36Choice struct {
	Tp     Frequency6Code      `xml:"Tp"`
	Prd    FrequencyPeriod1    `xml:"Prd"`
	PtInTm FrequencyAndMoment1 `xml:"PtInTm"`
}

func (r Frequency36Choice) Validate() error {
	return util.Validate(&r)
}

type FrequencyAndMoment1 struct {
	Tp     Frequency6Code           `xml:"Tp"`
	PtInTm common.Exact2NumericText `xml:"PtInTm"`
}

func (r FrequencyAndMoment1) Validate() error {
	return util.Validate(&r)
}

type FrequencyPeriod1 struct {
	Tp        Frequency6Code `xml:"Tp"`
	CntPerPrd float64        `xml:"CntPerPrd"`
}

func (r FrequencyPeriod1) Validate() error {
	return util.Validate(&r)
}

type Garnishment3 struct {
	Tp                GarnishmentType1                   `xml:"Tp"`
	Grnshee           *PartyIdentification135            `xml:"Grnshee,omitempty" json:",omitempty"`
	GrnshmtAdmstr     *PartyIdentification135            `xml:"GrnshmtAdmstr,omitempty" json:",omitempty"`
	RefNb             *common.Max140Text                 `xml:"RefNb,omitempty" json:",omitempty"`
	Dt                *common.ISODate                    `xml:"Dt,omitempty" json:",omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
	FmlyMdclInsrncInd bool                               `xml:"FmlyMdclInsrncInd,omitempty" json:",omitempty"`
	MplyeeTermntnInd  bool                               `xml:"MplyeeTermntnInd,omitempty" json:",omitempty"`
}

func (r Garnishment3) Validate() error {
	return util.Validate(&r)
}

type GarnishmentType1 struct {
	CdOrPrtry GarnishmentType1Choice `xml:"CdOrPrtry"`
	Issr      *common.Max35Text      `xml:"Issr,omitempty" json:",omitempty"`
}

func (r GarnishmentType1) Validate() error {
	return util.Validate(&r)
}

type GarnishmentType1Choice struct {
	Cd    ExternalGarnishmentType1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

func (r GarnishmentType1Choice) Validate() error {
	return util.Validate(&r)
}

type GenericAccountIdentification1 struct {
	Id      common.Max34Text          `xml:"Id"`
	SchmeNm *AccountSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text         `xml:"Issr,omitempty" json:",omitempty"`
}

func (r GenericAccountIdentification1) Validate() error {
	return util.Validate(&r)
}

type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                          `xml:"Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text                         `xml:"Issr,omitempty" json:",omitempty"`
}

func (r GenericFinancialIdentification1) Validate() error {
	return util.Validate(&r)
}

type GenericIdentification30 struct {
	Id      common.Exact4AlphaNumericText `xml:"Id"`
	Issr    common.Max35Text              `xml:"Issr"`
	SchmeNm *common.Max35Text             `xml:"SchmeNm,omitempty" json:",omitempty"`
}

func (r GenericIdentification30) Validate() error {
	return util.Validate(&r)
}

type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                             `xml:"Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text                            `xml:"Issr,omitempty" json:",omitempty"`
}

func (r GenericOrganisationIdentification1) Validate() error {
	return util.Validate(&r)
}

type GenericPersonIdentification1 struct {
	Id      common.Max35Text                       `xml:"Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text                      `xml:"Issr,omitempty" json:",omitempty"`
}

func (r GenericPersonIdentification1) Validate() error {
	return util.Validate(&r)
}

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

func (r LocalInstrument2Choice) Validate() error {
	return util.Validate(&r)
}

type MandateClassification1Choice struct {
	Cd    common.MandateClassification1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

func (r MandateClassification1Choice) Validate() error {
	return util.Validate(&r)
}

type MandateRelatedData1Choice struct {
	DrctDbtMndt *MandateRelatedInformation14 `xml:"DrctDbtMndt,omitempty" json:",omitempty"`
	CdtTrfMndt  *CreditTransferMandateData1  `xml:"CdtTrfMndt,omitempty" json:",omitempty"`
}

func (r MandateRelatedData1Choice) Validate() error {
	return util.Validate(&r)
}

type MandateRelatedInformation14 struct {
	MndtId        *common.Max35Text              `xml:"MndtId,omitempty" json:",omitempty"`
	DtOfSgntr     *common.ISODate                `xml:"DtOfSgntr,omitempty" json:",omitempty"`
	AmdmntInd     bool                           `xml:"AmdmntInd,omitempty" json:",omitempty"`
	AmdmntInfDtls *AmendmentInformationDetails13 `xml:"AmdmntInfDtls,omitempty" json:",omitempty"`
	ElctrncSgntr  *common.Max1025Text            `xml:"ElctrncSgntr,omitempty" json:",omitempty"`
	FrstColltnDt  *common.ISODate                `xml:"FrstColltnDt,omitempty" json:",omitempty"`
	FnlColltnDt   *common.ISODate                `xml:"FnlColltnDt,omitempty" json:",omitempty"`
	Frqcy         *Frequency36Choice             `xml:"Frqcy,omitempty" json:",omitempty"`
	Rsn           *MandateSetupReason1Choice     `xml:"Rsn,omitempty" json:",omitempty"`
	TrckgDays     *common.Exact2NumericText      `xml:"TrckgDays,omitempty" json:",omitempty"`
}

func (r MandateRelatedInformation14) Validate() error {
	return util.Validate(&r)
}

type MandateSetupReason1Choice struct {
	Cd    ExternalMandateSetupReason1Code `xml:"Cd"`
	Prtry common.Max70Text                `xml:"Prtry"`
}

func (r MandateSetupReason1Choice) Validate() error {
	return util.Validate(&r)
}

type MandateTypeInformation2 struct {
	SvcLvl    *ServiceLevel8Choice          `xml:"SvcLvl,omitempty" json:",omitempty"`
	LclInstrm *LocalInstrument2Choice       `xml:"LclInstrm,omitempty" json:",omitempty"`
	CtgyPurp  *CategoryPurpose1Choice       `xml:"CtgyPurp,omitempty" json:",omitempty"`
	Clssfctn  *MandateClassification1Choice `xml:"Clssfctn,omitempty" json:",omitempty"`
}

func (r MandateTypeInformation2) Validate() error {
	return util.Validate(&r)
}

type OrganisationIdentification29 struct {
	AnyBIC *common.AnyBICDec2014Identifier      `xml:"AnyBIC,omitempty" json:",omitempty"`
	LEI    *common.LEIIdentifier                `xml:"LEI,omitempty" json:",omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

func (r OrganisationIdentification29) Validate() error {
	return util.Validate(&r)
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                        `xml:"Prtry"`
}

func (r OrganisationIdentificationSchemeName1Choice) Validate() error {
	return util.Validate(&r)
}

type OriginalGroupHeader15 struct {
	GrpCxlId     *common.Max35Text            `xml:"GrpCxlId,omitempty" json:",omitempty"`
	Case         *Case5                       `xml:"Case,omitempty" json:",omitempty"`
	OrgnlMsgId   common.Max35Text             `xml:"OrgnlMsgId"`
	OrgnlMsgNmId common.Max35Text             `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm *common.ISODateTime          `xml:"OrgnlCreDtTm,omitempty" json:",omitempty"`
	NbOfTxs      *common.Max15NumericText     `xml:"NbOfTxs,omitempty" json:",omitempty"`
	CtrlSum      float64                      `xml:"CtrlSum,omitempty" json:",omitempty"`
	GrpCxl       bool                         `xml:"GrpCxl,omitempty" json:",omitempty"`
	CxlRsnInf    []PaymentCancellationReason5 `xml:"CxlRsnInf,omitempty" json:",omitempty"`
}

func (r OriginalGroupHeader15) Validate() error {
	return util.Validate(&r)
}

type OriginalGroupInformation29 struct {
	OrgnlMsgId   common.Max35Text    `xml:"OrgnlMsgId"`
	OrgnlMsgNmId common.Max35Text    `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm *common.ISODateTime `xml:"OrgnlCreDtTm,omitempty" json:",omitempty"`
}

func (r OriginalGroupInformation29) Validate() error {
	return util.Validate(&r)
}

type OriginalPaymentInstruction36 struct {
	PmtCxlId      *common.Max35Text            `xml:"PmtCxlId,omitempty" json:",omitempty"`
	Case          *Case5                       `xml:"Case,omitempty" json:",omitempty"`
	OrgnlPmtInfId common.Max35Text             `xml:"OrgnlPmtInfId"`
	OrgnlGrpInf   *OriginalGroupInformation29  `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	NbOfTxs       *common.Max15NumericText     `xml:"NbOfTxs,omitempty" json:",omitempty"`
	CtrlSum       float64                      `xml:"CtrlSum,omitempty" json:",omitempty"`
	PmtInfCxl     bool                         `xml:"PmtInfCxl,omitempty" json:",omitempty"`
	CxlRsnInf     []PaymentCancellationReason5 `xml:"CxlRsnInf,omitempty" json:",omitempty"`
	TxInf         []PaymentTransaction124      `xml:"TxInf,omitempty" json:",omitempty"`
}

func (r OriginalPaymentInstruction36) Validate() error {
	return util.Validate(&r)
}

type OriginalTransactionReference31 struct {
	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount            `xml:"IntrBkSttlmAmt,omitempty" json:",omitempty"`
	Amt            *AmountType4Choice                            `xml:"Amt,omitempty" json:",omitempty"`
	IntrBkSttlmDt  *common.ISODate                               `xml:"IntrBkSttlmDt,omitempty" json:",omitempty"`
	ReqdColltnDt   *common.ISODate                               `xml:"ReqdColltnDt,omitempty" json:",omitempty"`
	ReqdExctnDt    *DateAndDateTime2Choice                       `xml:"ReqdExctnDt,omitempty" json:",omitempty"`
	CdtrSchmeId    *PartyIdentification135                       `xml:"CdtrSchmeId,omitempty" json:",omitempty"`
	SttlmInf       *SettlementInstruction7                       `xml:"SttlmInf,omitempty" json:",omitempty"`
	PmtTpInf       *PaymentTypeInformation27                     `xml:"PmtTpInf,omitempty" json:",omitempty"`
	PmtMtd         *PaymentMethod4Code                           `xml:"PmtMtd,omitempty" json:",omitempty"`
	MndtRltdInf    *MandateRelatedData1Choice                    `xml:"MndtRltdInf,omitempty" json:",omitempty"`
	RmtInf         *RemittanceInformation16                      `xml:"RmtInf,omitempty" json:",omitempty"`
	UltmtDbtr      *Party40Choice                                `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	Dbtr           *Party40Choice                                `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAcct       *CashAccount38                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	DbtrAgt        *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	DbtrAgtAcct    *CashAccount38                                `xml:"DbtrAgtAcct,omitempty" json:",omitempty"`
	CdtrAgt        *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty" json:",omitempty"`
	CdtrAgtAcct    *CashAccount38                                `xml:"CdtrAgtAcct,omitempty" json:",omitempty"`
	Cdtr           *Party40Choice                                `xml:"Cdtr,omitempty" json:",omitempty"`
	CdtrAcct       *CashAccount38                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	UltmtCdtr      *Party40Choice                                `xml:"UltmtCdtr,omitempty" json:",omitempty"`
	Purp           *Purpose2Choice                               `xml:"Purp,omitempty" json:",omitempty"`
}

func (r OriginalTransactionReference31) Validate() error {
	return util.Validate(&r)
}

type OtherContact1 struct {
	ChanlTp common.Max4Text    `xml:"ChanlTp"`
	Id      *common.Max128Text `xml:"Id,omitempty" json:",omitempty"`
}

func (r OtherContact1) Validate() error {
	return util.Validate(&r)
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"OrgId"`
	PrvtId PersonIdentification13       `xml:"PrvtId"`
}

func (r Party38Choice) Validate() error {
	return util.Validate(&r)
}

type Party40Choice struct {
	Pty *PartyIdentification135                       `xml:"Pty,omitempty" json:",omitempty"`
	Agt *BranchAndFinancialInstitutionIdentification6 `xml:"Agt,omitempty" json:",omitempty"`
}

func (r Party40Choice) Validate() error {
	return util.Validate(&r)
}

type PartyIdentification135 struct {
	Nm        *common.Max140Text  `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr   *PostalAddress24    `xml:"PstlAdr,omitempty" json:",omitempty"`
	Id        *Party38Choice      `xml:"Id,omitempty" json:",omitempty"`
	CtryOfRes *common.CountryCode `xml:"CtryOfRes,omitempty" json:",omitempty"`
	CtctDtls  *Contact4           `xml:"CtctDtls,omitempty" json:",omitempty"`
}

func (r PartyIdentification135) Validate() error {
	return util.Validate(&r)
}

type PaymentCancellationReason5 struct {
	Orgtr    *PartyIdentification135     `xml:"Orgtr,omitempty" json:",omitempty"`
	Rsn      *CancellationReason33Choice `xml:"Rsn,omitempty" json:",omitempty"`
	AddtlInf []common.Max105Text         `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r PaymentCancellationReason5) Validate() error {
	return util.Validate(&r)
}

type PaymentTransaction124 struct {
	CxlId             *common.Max35Text                  `xml:"CxlId,omitempty" json:",omitempty"`
	Case              *Case5                             `xml:"Case,omitempty" json:",omitempty"`
	OrgnlInstrId      *common.Max35Text                  `xml:"OrgnlInstrId,omitempty" json:",omitempty"`
	OrgnlEndToEndId   *common.Max35Text                  `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	OrgnlUETR         *common.UUIDv4Identifier           `xml:"OrgnlUETR,omitempty" json:",omitempty"`
	OrgnlInstdAmt     *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty" json:",omitempty"`
	OrgnlReqdExctnDt  *DateAndDateTime2Choice            `xml:"OrgnlReqdExctnDt,omitempty" json:",omitempty"`
	OrgnlReqdColltnDt *common.ISODate                    `xml:"OrgnlReqdColltnDt,omitempty" json:",omitempty"`
	CxlRsnInf         []PaymentCancellationReason5       `xml:"CxlRsnInf,omitempty" json:",omitempty"`
	OrgnlTxRef        *OriginalTransactionReference31    `xml:"OrgnlTxRef,omitempty" json:",omitempty"`
	SplmtryData       []SupplementaryData1               `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r PaymentTransaction124) Validate() error {
	return util.Validate(&r)
}

type PaymentTypeInformation27 struct {
	InstrPrty *Priority2Code          `xml:"InstrPrty,omitempty" json:",omitempty"`
	ClrChanl  *ClearingChannel2Code   `xml:"ClrChanl,omitempty" json:",omitempty"`
	SvcLvl    []ServiceLevel8Choice   `xml:"SvcLvl,omitempty" json:",omitempty"`
	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm,omitempty" json:",omitempty"`
	SeqTp     *SequenceType3Code      `xml:"SeqTp,omitempty" json:",omitempty"`
	CtgyPurp  *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty" json:",omitempty"`
}

func (r PaymentTypeInformation27) Validate() error {
	return util.Validate(&r)
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1          `xml:"DtAndPlcOfBirth,omitempty" json:",omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

func (r PersonIdentification13) Validate() error {
	return util.Validate(&r)
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

func (r PersonIdentificationSchemeName1Choice) Validate() error {
	return util.Validate(&r)
}

type PostalAddress24 struct {
	AdrTp       *AddressType3Choice `xml:"AdrTp,omitempty" json:",omitempty"`
	Dept        *common.Max70Text   `xml:"Dept,omitempty" json:",omitempty"`
	SubDept     *common.Max70Text   `xml:"SubDept,omitempty" json:",omitempty"`
	StrtNm      *common.Max70Text   `xml:"StrtNm,omitempty" json:",omitempty"`
	BldgNb      *common.Max16Text   `xml:"BldgNb,omitempty" json:",omitempty"`
	BldgNm      *common.Max35Text   `xml:"BldgNm,omitempty" json:",omitempty"`
	Flr         *common.Max70Text   `xml:"Flr,omitempty" json:",omitempty"`
	PstBx       *common.Max16Text   `xml:"PstBx,omitempty" json:",omitempty"`
	Room        *common.Max70Text   `xml:"Room,omitempty" json:",omitempty"`
	PstCd       *common.Max16Text   `xml:"PstCd,omitempty" json:",omitempty"`
	TwnNm       *common.Max35Text   `xml:"TwnNm,omitempty" json:",omitempty"`
	TwnLctnNm   *common.Max35Text   `xml:"TwnLctnNm,omitempty" json:",omitempty"`
	DstrctNm    *common.Max35Text   `xml:"DstrctNm,omitempty" json:",omitempty"`
	CtrySubDvsn *common.Max35Text   `xml:"CtrySubDvsn,omitempty" json:",omitempty"`
	Ctry        *common.CountryCode `xml:"Ctry,omitempty" json:",omitempty"`
	AdrLine     []common.Max70Text  `xml:"AdrLine,omitempty" json:",omitempty"`
}

func (r PostalAddress24) Validate() error {
	return util.Validate(&r)
}

type ProxyAccountIdentification1 struct {
	Tp *ProxyAccountType1Choice `xml:"Tp,omitempty" json:",omitempty"`
	Id common.Max2048Text       `xml:"Id"`
}

func (r ProxyAccountIdentification1) Validate() error {
	return util.Validate(&r)
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"Cd"`
	Prtry common.Max35Text              `xml:"Prtry"`
}

func (r ProxyAccountType1Choice) Validate() error {
	return util.Validate(&r)
}

type Purpose2Choice struct {
	Cd    ExternalPurpose1Code `xml:"Cd"`
	Prtry common.Max35Text     `xml:"Prtry"`
}

func (r Purpose2Choice) Validate() error {
	return util.Validate(&r)
}

type ReferredDocumentInformation7 struct {
	Tp       *ReferredDocumentType4     `xml:"Tp,omitempty" json:",omitempty"`
	Nb       *common.Max35Text          `xml:"Nb,omitempty" json:",omitempty"`
	RltdDt   *common.ISODate            `xml:"RltdDt,omitempty" json:",omitempty"`
	LineDtls []DocumentLineInformation1 `xml:"LineDtls,omitempty" json:",omitempty"`
}

func (r ReferredDocumentInformation7) Validate() error {
	return util.Validate(&r)
}

type ReferredDocumentType3Choice struct {
	Cd    DocumentType6Code `xml:"Cd"`
	Prtry common.Max35Text  `xml:"Prtry"`
}

func (r ReferredDocumentType3Choice) Validate() error {
	return util.Validate(&r)
}

type ReferredDocumentType4 struct {
	CdOrPrtry ReferredDocumentType3Choice `xml:"CdOrPrtry"`
	Issr      *common.Max35Text           `xml:"Issr,omitempty" json:",omitempty"`
}

func (r ReferredDocumentType4) Validate() error {
	return util.Validate(&r)
}

type RemittanceAmount2 struct {
	DuePyblAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty" json:",omitempty"`
	DscntApldAmt      []DiscountAmountAndType1           `xml:"DscntApldAmt,omitempty" json:",omitempty"`
	CdtNoteAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty" json:",omitempty"`
	TaxAmt            []TaxAmountAndType1                `xml:"TaxAmt,omitempty" json:",omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1              `xml:"AdjstmntAmtAndRsn,omitempty" json:",omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
}

func (r RemittanceAmount2) Validate() error {
	return util.Validate(&r)
}

type RemittanceAmount3 struct {
	DuePyblAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty" json:",omitempty"`
	DscntApldAmt      []DiscountAmountAndType1           `xml:"DscntApldAmt,omitempty" json:",omitempty"`
	CdtNoteAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty" json:",omitempty"`
	TaxAmt            []TaxAmountAndType1                `xml:"TaxAmt,omitempty" json:",omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1              `xml:"AdjstmntAmtAndRsn,omitempty" json:",omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
}

func (r RemittanceAmount3) Validate() error {
	return util.Validate(&r)
}

type RemittanceInformation16 struct {
	Ustrd []common.Max140Text                 `xml:"Ustrd,omitempty" json:",omitempty"`
	Strd  []StructuredRemittanceInformation16 `xml:"Strd,omitempty" json:",omitempty"`
}

func (r RemittanceInformation16) Validate() error {
	return util.Validate(&r)
}

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"Cd"`
	Prtry common.Max35Text          `xml:"Prtry"`
}

func (r ServiceLevel8Choice) Validate() error {
	return util.Validate(&r)
}

type SettlementInstruction7 struct {
	SttlmMtd             SettlementMethod1Code                         `xml:"SttlmMtd"`
	SttlmAcct            *CashAccount38                                `xml:"SttlmAcct,omitempty" json:",omitempty"`
	ClrSys               *ClearingSystemIdentification3Choice          `xml:"ClrSys,omitempty" json:",omitempty"`
	InstgRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"InstgRmbrsmntAgt,omitempty" json:",omitempty"`
	InstgRmbrsmntAgtAcct *CashAccount38                                `xml:"InstgRmbrsmntAgtAcct,omitempty" json:",omitempty"`
	InstdRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"InstdRmbrsmntAgt,omitempty" json:",omitempty"`
	InstdRmbrsmntAgtAcct *CashAccount38                                `xml:"InstdRmbrsmntAgtAcct,omitempty" json:",omitempty"`
	ThrdRmbrsmntAgt      *BranchAndFinancialInstitutionIdentification6 `xml:"ThrdRmbrsmntAgt,omitempty" json:",omitempty"`
	ThrdRmbrsmntAgtAcct  *CashAccount38                                `xml:"ThrdRmbrsmntAgtAcct,omitempty" json:",omitempty"`
}

func (r SettlementInstruction7) Validate() error {
	return util.Validate(&r)
}

type StructuredRemittanceInformation16 struct {
	RfrdDocInf  []ReferredDocumentInformation7 `xml:"RfrdDocInf,omitempty" json:",omitempty"`
	RfrdDocAmt  *RemittanceAmount2             `xml:"RfrdDocAmt,omitempty" json:",omitempty"`
	CdtrRefInf  *CreditorReferenceInformation2 `xml:"CdtrRefInf,omitempty" json:",omitempty"`
	Invcr       *PartyIdentification135        `xml:"Invcr,omitempty" json:",omitempty"`
	Invcee      *PartyIdentification135        `xml:"Invcee,omitempty" json:",omitempty"`
	TaxRmt      *TaxInformation7               `xml:"TaxRmt,omitempty" json:",omitempty"`
	GrnshmtRmt  *Garnishment3                  `xml:"GrnshmtRmt,omitempty" json:",omitempty"`
	AddtlRmtInf []common.Max140Text            `xml:"AddtlRmtInf,omitempty" json:",omitempty"`
}

func (r StructuredRemittanceInformation16) Validate() error {
	return util.Validate(&r)
}

type SupplementaryData1 struct {
	PlcAndNm *common.Max350Text         `xml:"PlcAndNm,omitempty" json:",omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"Envlp"`
}

func (r SupplementaryData1) Validate() error {
	return util.Validate(&r)
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",innerxml"`
}

func (r SupplementaryDataEnvelope1) Validate() error {
	return util.Validate(&r)
}

type TaxAmount2 struct {
	Rate         float64                            `xml:"Rate,omitempty" json:",omitempty"`
	TaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty" json:",omitempty"`
	TtlAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty" json:",omitempty"`
	Dtls         []TaxRecordDetails2                `xml:"Dtls,omitempty" json:",omitempty"`
}

func (r TaxAmount2) Validate() error {
	return util.Validate(&r)
}

type TaxAmountAndType1 struct {
	Tp  *TaxAmountType1Choice             `xml:"Tp,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

func (r TaxAmountAndType1) Validate() error {
	return util.Validate(&r)
}

type TaxAmountType1Choice struct {
	Cd    ExternalTaxAmountType1Code `xml:"Cd"`
	Prtry common.Max35Text           `xml:"Prtry"`
}

func (r TaxAmountType1Choice) Validate() error {
	return util.Validate(&r)
}

type TaxAuthorisation1 struct {
	Titl *common.Max35Text  `xml:"Titl,omitempty" json:",omitempty"`
	Nm   *common.Max140Text `xml:"Nm,omitempty" json:",omitempty"`
}

func (r TaxAuthorisation1) Validate() error {
	return util.Validate(&r)
}

type TaxInformation7 struct {
	Cdtr            *TaxParty1                         `xml:"Cdtr,omitempty" json:",omitempty"`
	Dbtr            *TaxParty2                         `xml:"Dbtr,omitempty" json:",omitempty"`
	UltmtDbtr       *TaxParty2                         `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	AdmstnZone      *common.Max35Text                  `xml:"AdmstnZone,omitempty" json:",omitempty"`
	RefNb           *common.Max140Text                 `xml:"RefNb,omitempty" json:",omitempty"`
	Mtd             *common.Max35Text                  `xml:"Mtd,omitempty" json:",omitempty"`
	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty" json:",omitempty"`
	TtlTaxAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty" json:",omitempty"`
	Dt              *common.ISODate                    `xml:"Dt,omitempty" json:",omitempty"`
	SeqNb           float64                            `xml:"SeqNb,omitempty" json:",omitempty"`
	Rcrd            []TaxRecord2                       `xml:"Rcrd,omitempty" json:",omitempty"`
}

func (r TaxInformation7) Validate() error {
	return util.Validate(&r)
}

type TaxParty1 struct {
	TaxId  *common.Max35Text `xml:"TaxId,omitempty" json:",omitempty"`
	RegnId *common.Max35Text `xml:"RegnId,omitempty" json:",omitempty"`
	TaxTp  *common.Max35Text `xml:"TaxTp,omitempty" json:",omitempty"`
}

func (r TaxParty1) Validate() error {
	return util.Validate(&r)
}

type TaxParty2 struct {
	TaxId   *common.Max35Text  `xml:"TaxId,omitempty" json:",omitempty"`
	RegnId  *common.Max35Text  `xml:"RegnId,omitempty" json:",omitempty"`
	TaxTp   *common.Max35Text  `xml:"TaxTp,omitempty" json:",omitempty"`
	Authstn *TaxAuthorisation1 `xml:"Authstn,omitempty" json:",omitempty"`
}

func (r TaxParty2) Validate() error {
	return util.Validate(&r)
}

type TaxPeriod2 struct {
	Yr     *common.ISODate       `xml:"Yr,omitempty" json:",omitempty"`
	Tp     *TaxRecordPeriod1Code `xml:"Tp,omitempty" json:",omitempty"`
	FrToDt *DatePeriod2          `xml:"FrToDt,omitempty" json:",omitempty"`
}

func (r TaxPeriod2) Validate() error {
	return util.Validate(&r)
}

type TaxRecord2 struct {
	Tp       *common.Max35Text  `xml:"Tp,omitempty" json:",omitempty"`
	Ctgy     *common.Max35Text  `xml:"Ctgy,omitempty" json:",omitempty"`
	CtgyDtls *common.Max35Text  `xml:"CtgyDtls,omitempty" json:",omitempty"`
	DbtrSts  *common.Max35Text  `xml:"DbtrSts,omitempty" json:",omitempty"`
	CertId   *common.Max35Text  `xml:"CertId,omitempty" json:",omitempty"`
	FrmsCd   *common.Max35Text  `xml:"FrmsCd,omitempty" json:",omitempty"`
	Prd      *TaxPeriod2        `xml:"Prd,omitempty" json:",omitempty"`
	TaxAmt   *TaxAmount2        `xml:"TaxAmt,omitempty" json:",omitempty"`
	AddtlInf *common.Max140Text `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r TaxRecord2) Validate() error {
	return util.Validate(&r)
}

type TaxRecordDetails2 struct {
	Prd *TaxPeriod2                       `xml:"Prd,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

func (r TaxRecordDetails2) Validate() error {
	return util.Validate(&r)
}

type UnderlyingTransaction27 struct {
	OrgnlGrpInfAndCxl *OriginalGroupHeader15         `xml:"OrgnlGrpInfAndCxl,omitempty" json:",omitempty"`
	OrgnlPmtInfAndCxl []OriginalPaymentInstruction36 `xml:"OrgnlPmtInfAndCxl,omitempty" json:",omitempty"`
}

func (r UnderlyingTransaction27) Validate() error {
	return util.Validate(&r)
}

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd *common.CreditDebitCode           `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	Rsn       *common.Max4Text                  `xml:"Rsn,omitempty" json:",omitempty"`
	AddtlInf  *common.Max140Text                `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r DocumentAdjustment1) Validate() error {
	return util.Validate(&r)
}

type FIToFIPaymentCancellationRequestV09 struct {
	XMLName     xml.Name                  `xml:"FIToFIPmtCxlReq"`
	Assgnmt     CaseAssignment5           `xml:"Assgnmt"`
	Case        *Case5                    `xml:"Case,omitempty" json:",omitempty"`
	CtrlData    *ControlData1             `xml:"CtrlData,omitempty" json:",omitempty"`
	Undrlyg     []UnderlyingTransaction26 `xml:"Undrlyg" json:",omitempty"`
	SplmtryData []SupplementaryData1      `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r FIToFIPaymentCancellationRequestV09) Validate() error {
	return util.Validate(&r)
}

type PaymentTransaction120 struct {
	CxlId               *common.Max35Text                             `xml:"CxlId,omitempty" json:",omitempty"`
	Case                *Case5                                        `xml:"Case,omitempty" json:",omitempty"`
	OrgnlGrpInf         *OriginalGroupInformation29                   `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	OrgnlInstrId        *common.Max35Text                             `xml:"OrgnlInstrId,omitempty" json:",omitempty"`
	OrgnlEndToEndId     *common.Max35Text                             `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	OrgnlTxId           *common.Max35Text                             `xml:"OrgnlTxId,omitempty" json:",omitempty"`
	OrgnlUETR           *common.UUIDv4Identifier                      `xml:"OrgnlUETR,omitempty" json:",omitempty"`
	OrgnlClrSysRef      *common.Max35Text                             `xml:"OrgnlClrSysRef,omitempty" json:",omitempty"`
	OrgnlIntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount            `xml:"OrgnlIntrBkSttlmAmt,omitempty" json:",omitempty"`
	OrgnlIntrBkSttlmDt  *common.ISODate                               `xml:"OrgnlIntrBkSttlmDt,omitempty" json:",omitempty"`
	Assgnr              *BranchAndFinancialInstitutionIdentification6 `xml:"Assgnr,omitempty" json:",omitempty"`
	Assgne              *BranchAndFinancialInstitutionIdentification6 `xml:"Assgne,omitempty" json:",omitempty"`
	CxlRsnInf           []PaymentCancellationReason5                  `xml:"CxlRsnInf,omitempty" json:",omitempty"`
	OrgnlTxRef          *OriginalTransactionReference31               `xml:"OrgnlTxRef,omitempty" json:",omitempty"`
	SplmtryData         []SupplementaryData1                          `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r PaymentTransaction120) Validate() error {
	return util.Validate(&r)
}

type UnderlyingTransaction26 struct {
	OrgnlGrpInfAndCxl *OriginalGroupHeader15  `xml:"OrgnlGrpInfAndCxl,omitempty" json:",omitempty"`
	TxInf             []PaymentTransaction120 `xml:"TxInf,omitempty" json:",omitempty"`
}

func (r UnderlyingTransaction26) Validate() error {
	return util.Validate(&r)
}

type AdditionalPaymentInformationV09 struct {
	XMLName     xml.Name                         `xml:"AddtlPmtInf"`
	Assgnmt     CaseAssignment5                  `xml:"Assgnmt"`
	Case        *Case5                           `xml:"Case,omitempty" json:",omitempty"`
	Undrlyg     UnderlyingTransaction5Choice     `xml:"Undrlyg"`
	Inf         PaymentComplementaryInformation8 `xml:"Inf"`
	SplmtryData []SupplementaryData1             `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AdditionalPaymentInformationV09) Validate() error {
	return util.Validate(&r)
}

type InstructionForCreditorAgent1 struct {
	Cd       *Instruction3Code  `xml:"Cd,omitempty" json:",omitempty"`
	InstrInf *common.Max140Text `xml:"InstrInf,omitempty" json:",omitempty"`
}

func (r InstructionForCreditorAgent1) Validate() error {
	return util.Validate(&r)
}

type InstructionForNextAgent1 struct {
	Cd       *Instruction4Code  `xml:"Cd,omitempty" json:",omitempty"`
	InstrInf *common.Max140Text `xml:"InstrInf,omitempty" json:",omitempty"`
}

func (r InstructionForNextAgent1) Validate() error {
	return util.Validate(&r)
}

type PaymentComplementaryInformation8 struct {
	InstrId           *common.Max35Text                             `xml:"InstrId,omitempty" json:",omitempty"`
	EndToEndId        *common.Max35Text                             `xml:"EndToEndId,omitempty" json:",omitempty"`
	TxId              *common.Max35Text                             `xml:"TxId,omitempty" json:",omitempty"`
	PmtTpInf          *PaymentTypeInformation27                     `xml:"PmtTpInf,omitempty" json:",omitempty"`
	ReqdExctnDt       *DateAndDateTime2Choice                       `xml:"ReqdExctnDt,omitempty" json:",omitempty"`
	ReqdColltnDt      *common.ISODate                               `xml:"ReqdColltnDt,omitempty" json:",omitempty"`
	IntrBkSttlmDt     *common.ISODate                               `xml:"IntrBkSttlmDt,omitempty" json:",omitempty"`
	Amt               *AmountType4Choice                            `xml:"Amt,omitempty" json:",omitempty"`
	IntrBkSttlmAmt    *ActiveOrHistoricCurrencyAndAmount            `xml:"IntrBkSttlmAmt,omitempty" json:",omitempty"`
	ChrgBr            *ChargeBearerType1Code                        `xml:"ChrgBr,omitempty" json:",omitempty"`
	UltmtDbtr         *PartyIdentification135                       `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	Dbtr              *PartyIdentification135                       `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAcct          *CashAccount38                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	DbtrAgt           *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	DbtrAgtAcct       *CashAccount38                                `xml:"DbtrAgtAcct,omitempty" json:",omitempty"`
	SttlmInf          *SettlementInstruction7                       `xml:"SttlmInf,omitempty" json:",omitempty"`
	IntrmyAgt1        *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty" json:",omitempty"`
	IntrmyAgt1Acct    *CashAccount38                                `xml:"IntrmyAgt1Acct,omitempty" json:",omitempty"`
	IntrmyAgt2        *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty" json:",omitempty"`
	IntrmyAgt2Acct    *CashAccount38                                `xml:"IntrmyAgt2Acct,omitempty" json:",omitempty"`
	IntrmyAgt3        *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty" json:",omitempty"`
	IntrmyAgt3Acct    *CashAccount38                                `xml:"IntrmyAgt3Acct,omitempty" json:",omitempty"`
	CdtrAgt           *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty" json:",omitempty"`
	CdtrAgtAcct       *CashAccount38                                `xml:"CdtrAgtAcct,omitempty" json:",omitempty"`
	Cdtr              *PartyIdentification135                       `xml:"Cdtr,omitempty" json:",omitempty"`
	CdtrAcct          *CashAccount38                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	UltmtCdtr         *PartyIdentification135                       `xml:"UltmtCdtr,omitempty" json:",omitempty"`
	Purp              *Purpose2Choice                               `xml:"Purp,omitempty" json:",omitempty"`
	InstrForDbtrAgt   *common.Max140Text                            `xml:"InstrForDbtrAgt,omitempty" json:",omitempty"`
	PrvsInstgAgt1     *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt1,omitempty" json:",omitempty"`
	PrvsInstgAgt1Acct *CashAccount38                                `xml:"PrvsInstgAgt1Acct,omitempty" json:",omitempty"`
	PrvsInstgAgt2     *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt2,omitempty" json:",omitempty"`
	PrvsInstgAgt2Acct *CashAccount38                                `xml:"PrvsInstgAgt2Acct,omitempty" json:",omitempty"`
	PrvsInstgAgt3     *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt3,omitempty" json:",omitempty"`
	PrvsInstgAgt3Acct *CashAccount38                                `xml:"PrvsInstgAgt3Acct,omitempty" json:",omitempty"`
	InstrForNxtAgt    []InstructionForNextAgent1                    `xml:"InstrForNxtAgt,omitempty" json:",omitempty"`
	InstrForCdtrAgt   []InstructionForCreditorAgent1                `xml:"InstrForCdtrAgt,omitempty" json:",omitempty"`
	RmtInf            *RemittanceInformation16                      `xml:"RmtInf,omitempty" json:",omitempty"`
}

func (r PaymentComplementaryInformation8) Validate() error {
	return util.Validate(&r)
}

type UnderlyingGroupInformation1 struct {
	OrgnlMsgId         common.Max35Text    `xml:"OrgnlMsgId"`
	OrgnlMsgNmId       common.Max35Text    `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm       *common.ISODateTime `xml:"OrgnlCreDtTm,omitempty" json:",omitempty"`
	OrgnlMsgDlvryChanl *common.Max35Text   `xml:"OrgnlMsgDlvryChanl,omitempty" json:",omitempty"`
}

func (r UnderlyingGroupInformation1) Validate() error {
	return util.Validate(&r)
}

type UnderlyingPaymentInstruction5 struct {
	OrgnlGrpInf     *UnderlyingGroupInformation1      `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	OrgnlPmtInfId   *common.Max35Text                 `xml:"OrgnlPmtInfId,omitempty" json:",omitempty"`
	OrgnlInstrId    *common.Max35Text                 `xml:"OrgnlInstrId,omitempty" json:",omitempty"`
	OrgnlEndToEndId *common.Max35Text                 `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	OrgnlUETR       *common.UUIDv4Identifier          `xml:"OrgnlUETR,omitempty" json:",omitempty"`
	OrgnlInstdAmt   ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt"`
	ReqdExctnDt     *DateAndDateTime2Choice           `xml:"ReqdExctnDt,omitempty" json:",omitempty"`
	ReqdColltnDt    *common.ISODate                   `xml:"ReqdColltnDt,omitempty" json:",omitempty"`
	OrgnlTxRef      *OriginalTransactionReference28   `xml:"OrgnlTxRef,omitempty" json:",omitempty"`
}

func (r UnderlyingPaymentInstruction5) Validate() error {
	return util.Validate(&r)
}

type UnderlyingPaymentTransaction4 struct {
	OrgnlGrpInf         *UnderlyingGroupInformation1       `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	OrgnlInstrId        *common.Max35Text                  `xml:"OrgnlInstrId,omitempty" json:",omitempty"`
	OrgnlEndToEndId     *common.Max35Text                  `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	OrgnlTxId           *common.Max35Text                  `xml:"OrgnlTxId,omitempty" json:",omitempty"`
	OrgnlUETR           *common.UUIDv4Identifier           `xml:"OrgnlUETR,omitempty" json:",omitempty"`
	OrgnlIntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt"`
	OrgnlIntrBkSttlmDt  *common.ISODate                    `xml:"OrgnlIntrBkSttlmDt"`
	OrgnlTxRef          *OriginalTransactionReference28    `xml:"OrgnlTxRef,omitempty" json:",omitempty"`
}

func (r UnderlyingPaymentTransaction4) Validate() error {
	return util.Validate(&r)
}

type UnderlyingStatementEntry3 struct {
	OrgnlGrpInf *OriginalGroupInformation29 `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	OrgnlStmtId *common.Max35Text           `xml:"OrgnlStmtId,omitempty" json:",omitempty"`
	OrgnlNtryId *common.Max35Text           `xml:"OrgnlNtryId,omitempty" json:",omitempty"`
	OrgnlUETR   *common.UUIDv4Identifier    `xml:"OrgnlUETR,omitempty" json:",omitempty"`
}

func (r UnderlyingStatementEntry3) Validate() error {
	return util.Validate(&r)
}

type UnderlyingTransaction5Choice struct {
	Initn    *UnderlyingPaymentInstruction5 `xml:"Initn,omitempty" json:",omitempty"`
	IntrBk   *UnderlyingPaymentTransaction4 `xml:"IntrBk,omitempty" json:",omitempty"`
	StmtNtry *UnderlyingStatementEntry3     `xml:"StmtNtry,omitempty" json:",omitempty"`
}

func (r UnderlyingTransaction5Choice) Validate() error {
	return util.Validate(&r)
}

type ActiveCurrencyAndAmount struct {
	Value float64                   `xml:",chardata"`
	Ccy   common.ActiveCurrencyCode `xml:"Ccy,attr"`
}

func (r ActiveCurrencyAndAmount) Validate() error {
	return util.Validate(&r)
}

type CancellationStatusReason3Choice struct {
	Cd    *ExternalPaymentCancellationRejection1Code `xml:"Cd,omitempty" json:",omitempty"`
	Prtry *common.Max35Text                          `xml:"Prtry,omitempty" json:",omitempty"`
}

func (r CancellationStatusReason3Choice) Validate() error {
	return util.Validate(&r)
}

type CancellationStatusReason4 struct {
	Orgtr    *PartyIdentification135          `xml:"Orgtr,omitempty" json:",omitempty"`
	Rsn      *CancellationStatusReason3Choice `xml:"Rsn,omitempty" json:",omitempty"`
	AddtlInf []common.Max105Text              `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r CancellationStatusReason4) Validate() error {
	return util.Validate(&r)
}

type ChargeType3Choice struct {
	Cd    *ExternalChargeType1Code `xml:"Cd,omitempty" json:",omitempty"`
	Prtry *GenericIdentification3  `xml:"Prtry,omitempty" json:",omitempty"`
}

func (r ChargeType3Choice) Validate() error {
	return util.Validate(&r)
}

type Charges6 struct {
	TtlChrgsAndTaxAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlChrgsAndTaxAmt,omitempty" json:",omitempty"`
	Rcrd              []ChargesRecord3                   `xml:"Rcrd,omitempty" json:",omitempty"`
}

func (r Charges6) Validate() error {
	return util.Validate(&r)
}

type Charges7 struct {
	Amt ActiveOrHistoricCurrencyAndAmount            `xml:"Amt"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
}

func (r Charges7) Validate() error {
	return util.Validate(&r)
}

type ChargesRecord3 struct {
	Amt         *ActiveOrHistoricCurrencyAndAmount            `xml:"Amt"`
	CdtDbtInd   *common.CreditDebitCode                       `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	ChrgInclInd bool                                          `xml:"ChrgInclInd,omitempty" json:",omitempty"`
	Tp          *ChargeType3Choice                            `xml:"Tp,omitempty" json:",omitempty"`
	Rate        float64                                       `xml:"Rate,omitempty" json:",omitempty"`
	Br          *ChargeBearerType1Code                        `xml:"Br,omitempty" json:",omitempty"`
	Agt         *BranchAndFinancialInstitutionIdentification6 `xml:"Agt,omitempty" json:",omitempty"`
	Tax         *TaxCharges2                                  `xml:"Tax,omitempty" json:",omitempty"`
}

func (r ChargesRecord3) Validate() error {
	return util.Validate(&r)
}

type ClaimNonReceipt2 struct {
	DtPrcd      common.ISODate                                `xml:"DtPrcd"`
	OrgnlNxtAgt *BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlNxtAgt,omitempty" json:",omitempty"`
}

func (r ClaimNonReceipt2) Validate() error {
	return util.Validate(&r)
}

type ClaimNonReceipt2Choice struct {
	Accptd *ClaimNonReceipt2                   `xml:"Accptd,omitempty" json:",omitempty"`
	Rjctd  *ClaimNonReceiptRejectReason1Choice `xml:"Rjctd,omitempty" json:",omitempty"`
}

func (r ClaimNonReceipt2Choice) Validate() error {
	return util.Validate(&r)
}

type ClaimNonReceiptRejectReason1Choice struct {
	Cd    *ExternalClaimNonReceiptRejection1Code `xml:"Cd,omitempty" json:",omitempty"`
	Prtry *common.Max35Text                      `xml:"Prtry,omitempty" json:",omitempty"`
}

func (r ClaimNonReceiptRejectReason1Choice) Validate() error {
	return util.Validate(&r)
}

type Compensation2 struct {
	Amt     ActiveCurrencyAndAmount                      `xml:"Amt"`
	DbtrAgt BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt"`
	CdtrAgt BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt"`
	Rsn     CompensationReason1Choice                    `xml:"Rsn"`
}

func (r Compensation2) Validate() error {
	return util.Validate(&r)
}

type CompensationReason1Choice struct {
	Cd    *ExternalPaymentCompensationReason1Code `xml:"Cd,omitempty" json:",omitempty"`
	Prtry *common.Max35Text                       `xml:"Prtry,omitempty" json:",omitempty"`
}

func (r CompensationReason1Choice) Validate() error {
	return util.Validate(&r)
}

type CorrectiveGroupInformation1 struct {
	MsgId   common.Max35Text    `xml:"MsgId"`
	MsgNmId common.Max35Text    `xml:"MsgNmId"`
	CreDtTm *common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
}

func (r CorrectiveGroupInformation1) Validate() error {
	return util.Validate(&r)
}

type CorrectiveInterbankTransaction2 struct {
	GrpHdr         *CorrectiveGroupInformation1      `xml:"GrpHdr,omitempty" json:",omitempty"`
	InstrId        *common.Max35Text                 `xml:"InstrId,omitempty" json:",omitempty"`
	EndToEndId     *common.Max35Text                 `xml:"EndToEndId,omitempty" json:",omitempty"`
	TxId           *common.Max35Text                 `xml:"TxId,omitempty" json:",omitempty"`
	UETR           *common.UUIDv4Identifier          `xml:"UETR,omitempty" json:",omitempty"`
	IntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt"`
	IntrBkSttlmDt  common.ISODate                    `xml:"IntrBkSttlmDt"`
}

func (r CorrectiveInterbankTransaction2) Validate() error {
	return util.Validate(&r)
}

type CorrectivePaymentInitiation4 struct {
	GrpHdr       *CorrectiveGroupInformation1      `xml:"GrpHdr,omitempty" json:",omitempty"`
	PmtInfId     *common.Max35Text                 `xml:"PmtInfId,omitempty" json:",omitempty"`
	InstrId      *common.Max35Text                 `xml:"InstrId,omitempty" json:",omitempty"`
	EndToEndId   *common.Max35Text                 `xml:"EndToEndId,omitempty" json:",omitempty"`
	UETR         *common.UUIDv4Identifier          `xml:"UETR,omitempty" json:",omitempty"`
	InstdAmt     ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`
	ReqdExctnDt  *DateAndDateTime2Choice           `xml:"ReqdExctnDt,omitempty" json:",omitempty"`
	ReqdColltnDt *common.ISODate                   `xml:"ReqdColltnDt,omitempty" json:",omitempty"`
}

func (r CorrectivePaymentInitiation4) Validate() error {
	return util.Validate(&r)
}

type CorrectiveTransaction4Choice struct {
	Initn  *CorrectivePaymentInitiation4    `xml:"Initn,omitempty" json:",omitempty"`
	IntrBk *CorrectiveInterbankTransaction2 `xml:"IntrBk,omitempty" json:",omitempty"`
}

func (r CorrectiveTransaction4Choice) Validate() error {
	return util.Validate(&r)
}

type GenericIdentification3 struct {
	Id   common.Max35Text  `xml:"Id"`
	Issr *common.Max35Text `xml:"Issr,omitempty" json:",omitempty"`
}

func (r GenericIdentification3) Validate() error {
	return util.Validate(&r)
}

type InvestigationStatus5Choice struct {
	Conf           *ExternalInvestigationExecutionConfirmation1Code `xml:"Conf,omitempty" json:",omitempty"`
	RjctdMod       []ModificationStatusReason1Choice                `xml:"RjctdMod,omitempty" json:",omitempty"`
	DplctOf        *Case5                                           `xml:"DplctOf,omitempty" json:",omitempty"`
	AssgnmtCxlConf bool                                             `xml:"AssgnmtCxlConf,omitempty" json:",omitempty"`
}

func (r InvestigationStatus5Choice) Validate() error {
	return util.Validate(&r)
}

type ModificationStatusReason1Choice struct {
	Cd    *ExternalPaymentModificationRejection1Code `xml:"Cd,omitempty" json:",omitempty"`
	Prtry *common.Max35Text                          `xml:"Prtry,omitempty" json:",omitempty"`
}

func (r ModificationStatusReason1Choice) Validate() error {
	return util.Validate(&r)
}

type ModificationStatusReason2 struct {
	Orgtr    *PartyIdentification135          `xml:"Orgtr,omitempty" json:",omitempty"`
	Rsn      *ModificationStatusReason1Choice `xml:"Rsn,omitempty" json:",omitempty"`
	AddtlInf []common.Max105Text              `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r ModificationStatusReason2) Validate() error {
	return util.Validate(&r)
}

type NumberOfCancellationsPerStatus1 struct {
	DtldNbOfTxs common.Max15NumericText           `xml:"DtldNbOfTxs"`
	DtldSts     CancellationIndividualStatus1Code `xml:"DtldSts"`
	DtldCtrlSum float64                           `xml:"DtldCtrlSum,omitempty" json:",omitempty"`
}

func (r NumberOfCancellationsPerStatus1) Validate() error {
	return util.Validate(&r)
}

type NumberOfTransactionsPerStatus1 struct {
	DtldNbOfTxs common.Max15NumericText          `xml:"DtldNbOfTxs"`
	DtldSts     TransactionIndividualStatus1Code `xml:"DtldSts"`
	DtldCtrlSum float64                          `xml:"DtldCtrlSum,omitempty" json:",omitempty"`
}

func (r NumberOfTransactionsPerStatus1) Validate() error {
	return util.Validate(&r)
}

type OriginalGroupHeader14 struct {
	OrgnlGrpCxlId    *common.Max35Text                `xml:"OrgnlGrpCxlId,omitempty" json:",omitempty"`
	RslvdCase        *Case5                           `xml:"RslvdCase,omitempty" json:",omitempty"`
	OrgnlMsgId       common.Max35Text                 `xml:"OrgnlMsgId"`
	OrgnlMsgNmId     common.Max35Text                 `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm     *common.ISODateTime              `xml:"OrgnlCreDtTm,omitempty" json:",omitempty"`
	OrgnlNbOfTxs     *common.Max15NumericText         `xml:"OrgnlNbOfTxs,omitempty" json:",omitempty"`
	OrgnlCtrlSum     float64                          `xml:"OrgnlCtrlSum,omitempty" json:",omitempty"`
	GrpCxlSts        *GroupCancellationStatus1Code    `xml:"GrpCxlSts,omitempty" json:",omitempty"`
	CxlStsRsnInf     []CancellationStatusReason4      `xml:"CxlStsRsnInf,omitempty" json:",omitempty"`
	NbOfTxsPerCxlSts []NumberOfTransactionsPerStatus1 `xml:"NbOfTxsPerCxlSts,omitempty" json:",omitempty"`
}

func (r OriginalGroupHeader14) Validate() error {
	return util.Validate(&r)
}

type OriginalPaymentInstruction30 struct {
	OrgnlPmtInfCxlId *common.Max35Text                 `xml:"OrgnlPmtInfCxlId,omitempty" json:",omitempty"`
	RslvdCase        *Case5                            `xml:"RslvdCase,omitempty" json:",omitempty"`
	OrgnlPmtInfId    common.Max35Text                  `xml:"OrgnlPmtInfId"`
	OrgnlGrpInf      *OriginalGroupInformation29       `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	OrgnlNbOfTxs     *common.Max15NumericText          `xml:"OrgnlNbOfTxs,omitempty" json:",omitempty"`
	OrgnlCtrlSum     float64                           `xml:"OrgnlCtrlSum,omitempty" json:",omitempty"`
	PmtInfCxlSts     *GroupCancellationStatus1Code     `xml:"PmtInfCxlSts,omitempty" json:",omitempty"`
	CxlStsRsnInf     []CancellationStatusReason4       `xml:"CxlStsRsnInf,omitempty" json:",omitempty"`
	NbOfTxsPerCxlSts []NumberOfCancellationsPerStatus1 `xml:"NbOfTxsPerCxlSts,omitempty" json:",omitempty"`
	TxInfAndSts      []PaymentTransaction103           `xml:"TxInfAndSts,omitempty" json:",omitempty"`
}

func (r OriginalPaymentInstruction30) Validate() error {
	return util.Validate(&r)
}

type OriginalTransactionReference28 struct {
	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount            `xml:"IntrBkSttlmAmt,omitempty" json:",omitempty"`
	Amt            *AmountType4Choice                            `xml:"Amt,omitempty" json:",omitempty"`
	IntrBkSttlmDt  *common.ISODate                               `xml:"IntrBkSttlmDt,omitempty" json:",omitempty"`
	ReqdColltnDt   *common.ISODate                               `xml:"ReqdColltnDt,omitempty" json:",omitempty"`
	ReqdExctnDt    *DateAndDateTime2Choice                       `xml:"ReqdExctnDt,omitempty" json:",omitempty"`
	CdtrSchmeId    *PartyIdentification135                       `xml:"CdtrSchmeId,omitempty" json:",omitempty"`
	SttlmInf       *SettlementInstruction7                       `xml:"SttlmInf,omitempty" json:",omitempty"`
	PmtTpInf       *PaymentTypeInformation27                     `xml:"PmtTpInf,omitempty" json:",omitempty"`
	PmtMtd         *PaymentMethod4Code                           `xml:"PmtMtd,omitempty" json:",omitempty"`
	MndtRltdInf    *MandateRelatedInformation14                  `xml:"MndtRltdInf,omitempty" json:",omitempty"`
	RmtInf         *RemittanceInformation16                      `xml:"RmtInf,omitempty" json:",omitempty"`
	UltmtDbtr      *Party40Choice                                `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	Dbtr           *Party40Choice                                `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAcct       *CashAccount38                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	DbtrAgt        *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	DbtrAgtAcct    *CashAccount38                                `xml:"DbtrAgtAcct,omitempty" json:",omitempty"`
	CdtrAgt        *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty" json:",omitempty"`
	CdtrAgtAcct    *CashAccount38                                `xml:"CdtrAgtAcct,omitempty" json:",omitempty"`
	Cdtr           *Party40Choice                                `xml:"Cdtr,omitempty" json:",omitempty"`
	CdtrAcct       *CashAccount38                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	UltmtCdtr      *Party40Choice                                `xml:"UltmtCdtr,omitempty" json:",omitempty"`
	Purp           *Purpose2Choice                               `xml:"Purp,omitempty" json:",omitempty"`
}

func (r OriginalTransactionReference28) Validate() error {
	return util.Validate(&r)
}

type PaymentTransaction102 struct {
	CxlStsId            *common.Max35Text                  `xml:"CxlStsId,omitempty" json:",omitempty"`
	RslvdCase           *Case5                             `xml:"RslvdCase,omitempty" json:",omitempty"`
	OrgnlGrpInf         *OriginalGroupInformation29        `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	OrgnlInstrId        *common.Max35Text                  `xml:"OrgnlInstrId,omitempty" json:",omitempty"`
	OrgnlEndToEndId     *common.Max35Text                  `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	OrgnlTxId           *common.Max35Text                  `xml:"OrgnlTxId,omitempty" json:",omitempty"`
	OrgnlClrSysRef      *common.Max35Text                  `xml:"OrgnlClrSysRef,omitempty" json:",omitempty"`
	OrgnlUETR           *common.UUIDv4Identifier           `xml:"OrgnlUETR,omitempty" json:",omitempty"`
	TxCxlSts            *CancellationIndividualStatus1Code `xml:"TxCxlSts,omitempty" json:",omitempty"`
	CxlStsRsnInf        []CancellationStatusReason4        `xml:"CxlStsRsnInf,omitempty" json:",omitempty"`
	RsltnRltdInf        *ResolutionData1                   `xml:"RsltnRltdInf,omitempty" json:",omitempty"`
	OrgnlIntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty" json:",omitempty"`
	OrgnlIntrBkSttlmDt  *common.ISODate                    `xml:"OrgnlIntrBkSttlmDt,omitempty" json:",omitempty"`
	Assgnr              *Party40Choice                     `xml:"Assgnr,omitempty" json:",omitempty"`
	Assgne              *Party40Choice                     `xml:"Assgne,omitempty" json:",omitempty"`
	OrgnlTxRef          *OriginalTransactionReference28    `xml:"OrgnlTxRef,omitempty" json:",omitempty"`
}

func (r PaymentTransaction102) Validate() error {
	return util.Validate(&r)
}

type PaymentTransaction103 struct {
	CxlStsId          *common.Max35Text                  `xml:"CxlStsId,omitempty" json:",omitempty"`
	RslvdCase         *Case5                             `xml:"RslvdCase,omitempty" json:",omitempty"`
	OrgnlInstrId      *common.Max35Text                  `xml:"OrgnlInstrId,omitempty" json:",omitempty"`
	OrgnlEndToEndId   *common.Max35Text                  `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	UETR              *common.UUIDv4Identifier           `xml:"UETR,omitempty" json:",omitempty"`
	TxCxlSts          *CancellationIndividualStatus1Code `xml:"TxCxlSts,omitempty" json:",omitempty"`
	CxlStsRsnInf      []CancellationStatusReason4        `xml:"CxlStsRsnInf,omitempty" json:",omitempty"`
	OrgnlInstdAmt     *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty" json:",omitempty"`
	OrgnlReqdExctnDt  *DateAndDateTime2Choice            `xml:"OrgnlReqdExctnDt,omitempty" json:",omitempty"`
	OrgnlReqdColltnDt *common.ISODate                    `xml:"OrgnlReqdColltnDt,omitempty" json:",omitempty"`
	OrgnlTxRef        *OriginalTransactionReference28    `xml:"OrgnlTxRef,omitempty" json:",omitempty"`
}

func (r PaymentTransaction103) Validate() error {
	return util.Validate(&r)
}

type PaymentTransaction107 struct {
	ModStsId            *common.Max35Text                  `xml:"ModStsId,omitempty" json:",omitempty"`
	RslvdCase           *Case5                             `xml:"RslvdCase,omitempty" json:",omitempty"`
	OrgnlGrpInf         *OriginalGroupInformation29        `xml:"OrgnlGrpInf"`
	OrgnlPmtInfId       *common.Max35Text                  `xml:"OrgnlPmtInfId,omitempty" json:",omitempty"`
	OrgnlInstrId        *common.Max35Text                  `xml:"OrgnlInstrId,omitempty" json:",omitempty"`
	OrgnlEndToEndId     *common.Max35Text                  `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	OrgnlTxId           *common.Max35Text                  `xml:"OrgnlTxId,omitempty" json:",omitempty"`
	OrgnlClrSysRef      *common.Max35Text                  `xml:"OrgnlClrSysRef,omitempty" json:",omitempty"`
	OrgnlUETR           *common.UUIDv4Identifier           `xml:"OrgnlUETR,omitempty" json:",omitempty"`
	ModStsRsnInf        []ModificationStatusReason2        `xml:"ModStsRsnInf,omitempty" json:",omitempty"`
	RsltnRltdInf        *ResolutionData1                   `xml:"RsltnRltdInf,omitempty" json:",omitempty"`
	OrgnlIntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty" json:",omitempty"`
	OrgnlIntrBkSttlmDt  *common.ISODate                    `xml:"OrgnlIntrBkSttlmDt,omitempty" json:",omitempty"`
	Assgnr              *Party40Choice                     `xml:"Assgnr,omitempty" json:",omitempty"`
	Assgne              *Party40Choice                     `xml:"Assgne,omitempty" json:",omitempty"`
	OrgnlTxRef          *OriginalTransactionReference28    `xml:"OrgnlTxRef,omitempty" json:",omitempty"`
}

func (r PaymentTransaction107) Validate() error {
	return util.Validate(&r)
}

type ResolutionData1 struct {
	EndToEndId     *common.Max35Text                  `xml:"EndToEndId,omitempty" json:",omitempty"`
	TxId           *common.Max35Text                  `xml:"TxId,omitempty" json:",omitempty"`
	UETR           *common.UUIDv4Identifier           `xml:"UETR,omitempty" json:",omitempty"`
	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt,omitempty" json:",omitempty"`
	IntrBkSttlmDt  *common.ISODate                    `xml:"IntrBkSttlmDt,omitempty" json:",omitempty"`
	ClrChanl       *ClearingChannel2Code              `xml:"ClrChanl,omitempty" json:",omitempty"`
	Compstn        *Compensation2                     `xml:"Compstn,omitempty" json:",omitempty"`
	Chrgs          []Charges7                         `xml:"Chrgs,omitempty" json:",omitempty"`
}

func (r ResolutionData1) Validate() error {
	return util.Validate(&r)
}

type ResolutionOfInvestigationV09 struct {
	XMLName       xml.Name                      `xml:"RsltnOfInvstgtn"`
	Assgnmt       CaseAssignment5               `xml:"Assgnmt"`
	RslvdCase     *Case5                        `xml:"RslvdCase,omitempty" json:",omitempty"`
	Sts           InvestigationStatus5Choice    `xml:"Sts"`
	CxlDtls       []UnderlyingTransaction22     `xml:"CxlDtls,omitempty" json:",omitempty"`
	ModDtls       *PaymentTransaction107        `xml:"ModDtls,omitempty" json:",omitempty"`
	ClmNonRctDtls *ClaimNonReceipt2Choice       `xml:"ClmNonRctDtls,omitempty" json:",omitempty"`
	StmtDtls      *StatementResolutionEntry4    `xml:"StmtDtls,omitempty" json:",omitempty"`
	CrrctnTx      *CorrectiveTransaction4Choice `xml:"CrrctnTx,omitempty" json:",omitempty"`
	RsltnRltdInf  *ResolutionData1              `xml:"RsltnRltdInf,omitempty" json:",omitempty"`
	SplmtryData   []SupplementaryData1          `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ResolutionOfInvestigationV09) Validate() error {
	return util.Validate(&r)
}

type StatementResolutionEntry4 struct {
	OrgnlGrpInf *OriginalGroupInformation29        `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	OrgnlStmtId *common.Max35Text                  `xml:"OrgnlStmtId,omitempty" json:",omitempty"`
	UETR        *common.UUIDv4Identifier           `xml:"UETR,omitempty" json:",omitempty"`
	AcctSvcrRef *common.Max35Text                  `xml:"AcctSvcrRef,omitempty" json:",omitempty"`
	CrrctdAmt   *ActiveOrHistoricCurrencyAndAmount `xml:"CrrctdAmt,omitempty" json:",omitempty"`
	Chrgs       []Charges6                         `xml:"Chrgs,omitempty" json:",omitempty"`
	Purp        *Purpose2Choice                    `xml:"Purp,omitempty" json:",omitempty"`
}

func (r StatementResolutionEntry4) Validate() error {
	return util.Validate(&r)
}

type TaxCharges2 struct {
	Id   *common.Max35Text                  `xml:"Id,omitempty" json:",omitempty"`
	Rate float64                            `xml:"Rate,omitempty" json:",omitempty"`
	Amt  *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty" json:",omitempty"`
}

func (r TaxCharges2) Validate() error {
	return util.Validate(&r)
}

type UnderlyingTransaction22 struct {
	OrgnlGrpInfAndSts *OriginalGroupHeader14         `xml:"OrgnlGrpInfAndSts,omitempty" json:",omitempty"`
	OrgnlPmtInfAndSts []OriginalPaymentInstruction30 `xml:"OrgnlPmtInfAndSts,omitempty" json:",omitempty"`
	TxInfAndSts       []PaymentTransaction102        `xml:"TxInfAndSts,omitempty" json:",omitempty"`
}

func (r UnderlyingTransaction22) Validate() error {
	return util.Validate(&r)
}
