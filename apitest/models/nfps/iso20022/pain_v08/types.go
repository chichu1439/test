// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v08

import (
	"reflect"

	"apitest/util"
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
type ExternalCategoryPurpose1Code string

func (r ExternalCategoryPurpose1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalCategoryPurpose1Code", 1, 4)
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
type ExternalDocumentFormat1Code string

func (r ExternalDocumentFormat1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalDocumentFormat1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalDocumentLineType1Code string

func (r ExternalDocumentLineType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalDocumentFormat1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalDocumentType1Code string

func (r ExternalDocumentType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalDocumentType1Code", 1, 4)
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
type ExternalOrganisationIdentification1Code string

func (r ExternalOrganisationIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalOrganisationIdentification1Code", 1, 4)
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

// Must be at least 1 items long
type ExternalPaymentGroupStatus1Code string

func (r ExternalPaymentGroupStatus1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalPaymentGroupStatus1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalPaymentTransactionStatus1Code string

func (r ExternalPaymentTransactionStatus1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalPaymentTransactionStatus1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalStatusReason1Code string

func (r ExternalStatusReason1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalStatusReason1Code", 1, 4)
	}
	return nil
}

// May be one of ADWD, ADWD
type AdviceType1Code string

func (r AdviceType1Code) Validate() error {
	for _, vv := range []string{
		"ADWD", "ADWD",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("AdviceType1Code")
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

// May be one of MLDB, MLCD, MLFA, CRDB, CRCD, CRFA, PUDB, PUCD, PUFA, RGDB, RGCD, RGFA
type ChequeDelivery1Code string

func (r ChequeDelivery1Code) Validate() error {
	for _, vv := range []string{
		"MLDB", "MLCD", "MLFA", "CRDB", "CRCD", "CRFA", "PUDB", "PUCD", "PUFA", "RGDB", "RGCD", "RGFA",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("ChequeDelivery1Code")
}

// May be one of CCHQ, CCCH, BCHQ, DRFT, ELDR
type ChequeType2Code string

func (r ChequeType2Code) Validate() error {
	for _, vv := range []string{
		"CCHQ", "CCCH", "BCHQ", "DRFT", "ELDR",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("ChequeType2Code")
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

// May be one of CHK, TRF
type PaymentMethod7Code string

func (r PaymentMethod7Code) Validate() error {
	for _, vv := range []string{
		"CHK", "TRF",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("PaymentMethod7Code")
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

// May be one of CRED, DEBT, BOTH
type RegulatoryReportingType1Code string

func (r RegulatoryReportingType1Code) Validate() error {
	for _, vv := range []string{
		"CRED", "DEBT", "BOTH",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("RegulatoryReportingType1Code")
}

// May be one of FAXI, EDIC, URID, EMAL, POST, SMSM
type RemittanceLocationMethod2Code string

func (r RemittanceLocationMethod2Code) Validate() error {
	for _, vv := range []string{
		"FAXI", "EDIC", "URID", "EMAL", "POST", "SMSM",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("RemittanceLocationMethod2Code")
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
