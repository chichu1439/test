// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v07

import (
	"reflect"

	"isaving-swagger/util"
)

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

func (r ExternalAccountIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalAccountIdentification1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalCashAccountType1Code string

func (r ExternalCashAccountType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalCashAccountType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

func (r ExternalClearingSystemIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 5 {
		return util.NewErrTextLengthInvalid("ExternalClearingSystemIdentification1Code", 1, 5)
	}
	return nil
}

// Must be at least 1 items long
type ExternalEnquiryRequestType1Code string

func (r ExternalEnquiryRequestType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalEnquiryRequestType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

func (r ExternalFinancialInstitutionIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalFinancialInstitutionIdentification1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

func (r ExternalOrganisationIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalOrganisationIdentification1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalPaymentControlRequestType1Code string

func (r ExternalPaymentControlRequestType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalPaymentControlRequestType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

func (r ExternalPersonIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalPersonIdentification1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalSystemBalanceType1Code string

func (r ExternalSystemBalanceType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalSystemBalanceType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalMarketInfrastructure1Code string

func (r ExternalMarketInfrastructure1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 3 {
		return util.NewErrTextLengthInvalid("ExternalMarketInfrastructure1Code", 1, 3)
	}
	return nil
}

// Must be at least 1 items long
type ExternalSystemErrorHandling1Code string

func (r ExternalSystemErrorHandling1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalSystemErrorHandling1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalSystemEventType1Code string

func (r ExternalSystemEventType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalSystemEventType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalAgentInstruction1Code string

func (r ExternalAgentInstruction1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalAgentInstruction1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalCashClearingSystem1Code string

func (r ExternalCashClearingSystem1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 3 {
		return util.NewErrTextLengthInvalid("ExternalCashClearingSystem1Code", 1, 3)
	}
	return nil
}

// Must be at least 1 items long
type ExternalCategoryPurpose1Code string

func (r ExternalCategoryPurpose1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalCategoryPurpose1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalCreditorAgentInstruction1Code string

func (r ExternalCreditorAgentInstruction1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalCreditorAgentInstruction1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalDiscountAmountType1Code string

func (r ExternalDiscountAmountType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalDiscountAmountType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalDocumentLineType1Code string

func (r ExternalDocumentLineType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalDocumentLineType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalGarnishmentType1Code string

func (r ExternalGarnishmentType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalGarnishmentType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalLocalInstrument1Code string

func (r ExternalLocalInstrument1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 35 {
		return util.NewErrTextLengthInvalid("ExternalLocalInstrument1Code", 1, 35)
	}
	return nil
}

// Must be at least 1 items long
type ExternalMandateSetupReason1Code string

func (r ExternalMandateSetupReason1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalMandateSetupReason1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalProxyAccountType1Code string

func (r ExternalProxyAccountType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalProxyAccountType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalPurpose1Code string

func (r ExternalPurpose1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalPurpose1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalServiceLevel1Code string

func (r ExternalServiceLevel1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalServiceLevel1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalTaxAmountType1Code string

func (r ExternalTaxAmountType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalTaxAmountType1Code", 1, 4)
	}
	return nil
}

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

func (r PreferredContactMethod1Code) Validate() error {
	for _, vv := range []string{
		"LETT", "MAIL", "PHON", "FAXX", "CELL",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("PreferredContactMethod1Code")
}

// May be one of ALLL, CHNG, MODF, DELD
type QueryType2Code string

func (r QueryType2Code) Validate() error {
	for _, vv := range []string{
		"ALLL", "CHNG", "MODF", "DELD",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("QueryType2Code")
}

// May be one of BILA, MULT
type BalanceCounterparty1Code string

func (r BalanceCounterparty1Code) Validate() error {
	for _, vv := range []string{
		"BILA", "MULT",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("BalanceCounterparty1Code")
}

// May be one of MULT, BILI, MAND, DISC, NELI, INBI, GLBL, DIDB, SPLC, SPLF, TDLC, TDLF, UCDT, ACOL, EXGT
type LimitType3Code string

func (r LimitType3Code) Validate() error {
	for _, vv := range []string{
		"MULT", "BILI", "MAND", "DISC", "NELI", "INBI", "GLBL", "DIDB", "SPLC", "SPLF", "TDLC", "TDLF", "UCDT", "ACOL", "EXGT",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("LimitType3Code")
}

// May be one of BHOL, SMTN, NOOP, RCVR, ADTW
type SystemClosureReason1Code string

func (r SystemClosureReason1Code) Validate() error {
	for _, vv := range []string{
		"BHOL", "SMTN", "NOOP", "RCVR", "ADTW",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("SystemClosureReason1Code")
}

// May be one of SUSP, ACTV, CLSD, CLSG
type SystemStatus2Code string

func (r SystemStatus2Code) Validate() error {
	for _, vv := range []string{
		"SUSP", "ACTV", "CLSD", "CLSG",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("SystemStatus2Code")
}

// May be one of CBS, BCK, BAL, CLS, CTR, CBH, CBP, DPG, DPN, EXP, TCH, LMT, LIQ, DPP, DPH, DPS, STF, TRP, TCS, LOA, LOR, TCP, OND, MGL
type PaymentType3Code string

func (r PaymentType3Code) Validate() error {
	for _, vv := range []string{
		"CBS", "BCK", "BAL", "CLS", "CTR", "CBH", "CBP", "DPG", "DPN", "EXP", "TCH", "LMT", "LIQ", "DPP", "DPH", "DPS", "STF", "TRP", "TCS", "LOA", "LOR", "TCP", "OND", "MGL",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("PaymentType3Code")
}

// May be one of DEBT, CRED, SHAR, SLEV
type ChargeBearerType1Code string

func (r ChargeBearerType1Code) Validate() error {
	for _, vv := range []string{
		"DEBT", "CRED", "SHAR", "SLEV",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("ChargeBearerType1Code")
}

// May be one of RTGS, RTNS, MPNS, BOOK
type ClearingChannel2Code string

func (r ClearingChannel2Code) Validate() error {
	for _, vv := range []string{
		"RTGS", "RTNS", "MPNS", "BOOK",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("ClearingChannel2Code")
}

// May be one of RADM, RPIN, FXDR, DISP, PUOR, SCOR
type DocumentType3Code string

func (r DocumentType3Code) Validate() error {
	for _, vv := range []string{
		"RADM", "RPIN", "FXDR", "DISP", "PUOR", "SCOR",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("DocumentType3Code")
}

// May be one of MSIN, CNFA, DNFA, CINV, CREN, DEBN, HIRI, SBIN, CMCN, SOAC, DISP, BOLD, VCHR, AROI, TSUT, PUOR
type DocumentType6Code string

func (r DocumentType6Code) Validate() error {
	for _, vv := range []string{
		"MSIN", "CNFA", "DNFA", "CINV", "CREN", "DEBN", "HIRI", "SBIN", "CMCN", "SOAC", "DISP", "BOLD", "VCHR", "AROI", "TSUT", "PUOR",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("DocumentType6Code")
}

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA, FRTN
type Frequency6Code string

func (r Frequency6Code) Validate() error {
	for _, vv := range []string{
		"YEAR", "MNTH", "QURT", "MIAN", "WEEK", "DAIL", "ADHO", "INDA", "FRTN",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("Frequency6Code")
}

// May be one of PHOA, TELA
type Instruction4Code string

func (r Instruction4Code) Validate() error {
	for _, vv := range []string{
		"PHOA", "TELA",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("Instruction4Code")
}

// May be one of CHK, TRF, DD, TRA
type PaymentMethod4Code string

func (r PaymentMethod4Code) Validate() error {
	for _, vv := range []string{
		"CHK", "TRF", "DD", "TRA",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("PaymentMethod4Code")
}

// May be one of HIGH, NORM
type Priority2Code string

func (r Priority2Code) Validate() error {
	for _, vv := range []string{
		"HIGH", "NORM",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("Priority2Code")
}

// May be one of FRST, RCUR, FNAL, OOFF, RPRE
type SequenceType3Code string

func (r SequenceType3Code) Validate() error {
	for _, vv := range []string{
		"FRST", "RCUR", "FNAL", "OOFF", "RPRE",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("SequenceType3Code")
}

// May be one of INDA, INGA, COVE, CLRG
type SettlementMethod1Code string

func (r SettlementMethod1Code) Validate() error {
	for _, vv := range []string{
		"INDA", "INGA", "COVE", "CLRG",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("SettlementMethod1Code")
}

// May be one of MM01, MM02, MM03, MM04, MM05, MM06, MM07, MM08, MM09, MM10, MM11, MM12, QTR1, QTR2, QTR3, QTR4, HLF1, HLF2
type TaxRecordPeriod1Code string

func (r TaxRecordPeriod1Code) Validate() error {
	for _, vv := range []string{
		"MM01", "MM02", "MM03", "MM04", "MM05", "MM06", "MM07", "MM08", "MM09", "MM10", "MM11", "MM12", "QTR1", "QTR2", "QTR3", "QTR4", "HLF1", "HLF2",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("TaxRecordPeriod1Code")
}

// May be one of IN01, IN02, IN03, IN04, IN05, IN06, IN07, IN08, IN09, IN10, IN11, IN12, IN13, IN14, IN15, IN16, IN17, IN18, IN19, MM20, MM21, MM22, MM25, MM26, MM27, MM28, MM29, MM30, MM31, MM32, IN33, MM34, MM35, IN36, IN37, IN38, IN39, NARR
type UnableToApplyIncorrectInformation4Code string

func (r UnableToApplyIncorrectInformation4Code) Validate() error {
	for _, vv := range []string{
		"IN01", "IN02", "IN03", "IN04", "IN05", "IN06", "IN07", "IN08", "IN09", "IN10", "IN11", "IN12", "IN13", "IN14", "IN15", "IN16", "IN17", "IN18", "IN19", "MM20", "MM21", "MM22", "MM25", "MM26", "MM27", "MM28", "MM29", "MM30", "MM31", "MM32", "IN33", "MM34", "MM35", "IN36", "IN37", "IN38", "IN39", "NARR",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("UnableToApplyIncorrectInformation4Code")
}

// May be one of MS01, MS02, MS03, MS04, MS05, MS06, MS07, MS08, MS09, MS10, MS11, MS12, MS13, MS14, MS15, MS16, MS17, NARR
type UnableToApplyMissingInformation3Code string

func (r UnableToApplyMissingInformation3Code) Validate() error {
	for _, vv := range []string{
		"MS01", "MS02", "MS03", "MS04", "MS05", "MS06", "MS07", "MS08", "MS09", "MS10", "MS11", "MS12", "MS13", "MS14", "MS15", "MS16", "MS17", "NARR",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("UnableToApplyMissingInformation3Code")
}
