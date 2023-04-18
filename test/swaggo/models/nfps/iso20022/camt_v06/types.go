// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v06

import (
	"reflect"

	"isaving-swagger/util"
)

// Must be at least 1 items long
type ExternalEnquiryRequestType1Code string

func (r ExternalEnquiryRequestType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalEnquiryRequestType1Code", 1, 4)
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
type ExternalSystemErrorHandling1Code string

func (r ExternalSystemErrorHandling1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalSystemErrorHandling1Code", 1, 4)
	}
	return nil
}

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
type ExternalFinancialInstitutionIdentification1Code string

func (r ExternalFinancialInstitutionIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalFinancialInstitutionIdentification1Code", 1, 4)
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
type ExternalSystemEventType1Code string

func (r ExternalSystemEventType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalSystemEventType1Code", 1, 4)
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
type ExternalMarketInfrastructure1Code string

func (r ExternalMarketInfrastructure1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 3 {
		return util.NewErrTextLengthInvalid("ExternalMarketInfrastructure1Code", 1, 3)
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
type ExternalPurpose1Code string

func (r ExternalPurpose1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalPurpose1Code", 1, 4)
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

// May be one of HIGH, NORM, LOWW
type Priority1Code string

func (r Priority1Code) Validate() error {
	for _, vv := range []string{
		"HIGH", "NORM", "LOWW",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("Priority1Code")
}

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA, OVNG
type Frequency2Code string

func (r Frequency2Code) Validate() error {
	for _, vv := range []string{
		"YEAR", "MNTH", "QURT", "MIAN", "WEEK", "DAIL", "ADHO", "INDA", "OVNG",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("Frequency2Code")
}

// May be one of NFND, NAUT, UKNW, PCOR, WMSG, RNCR, MROI
type InvestigationRejection1Code string

func (r InvestigationRejection1Code) Validate() error {
	for _, vv := range []string{
		"NFND", "NAUT", "UKNW", "PCOR", "WMSG", "RNCR", "MROI",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("InvestigationRejection1Code")
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

// May be one of ENAB, DISA, DELD, REQD, BLKD
type ReservationStatus1Code string

func (r ReservationStatus1Code) Validate() error {
	for _, vv := range []string{
		"ENAB", "DISA", "DELD", "REQD", "BLKD",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("ReservationStatus1Code")
}

// May be one of CARE, UPAR, NSSR, HPAR, THRE, BLKD
type ReservationType2Code string

func (r ReservationType2Code) Validate() error {
	for _, vv := range []string{
		"CARE", "UPAR", "NSSR", "HPAR", "THRE", "BLKD",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("ReservationType2Code")
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

// May be one of RCBD, RCVD, NRCD
type NotificationStatus3Code string

func (r NotificationStatus3Code) Validate() error {
	for _, vv := range []string{
		"RCBD", "RCVD", "NRCD",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("NotificationStatus3Code")
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

// Must be at least 1 items long
type ExternalServiceLevel1Code string

func (r ExternalServiceLevel1Code) Validate() error {
	if len(string(r)) < 1 {
		return util.NewErrTextLengthInvalid("ExternalServiceLevel1Code", 1, 0)
	}
	return nil
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

// May be one of RJCR, ACCR, PDCR
type CancellationIndividualStatus1Code string

func (r CancellationIndividualStatus1Code) Validate() error {
	for _, vv := range []string{
		"RJCR", "ACCR", "PDCR",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("CancellationIndividualStatus1Code")
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

// May be one of PACR, RJCR, ACCR, PDCR
type GroupCancellationStatus1Code string

func (r GroupCancellationStatus1Code) Validate() error {
	for _, vv := range []string{
		"PACR", "RJCR", "ACCR", "PDCR",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("GroupCancellationStatus1Code")
}

// May be one of ACTC, RJCT, PDNG, ACCP, ACSP, ACSC, ACCR, ACWC
type TransactionIndividualStatus1Code string

func (r TransactionIndividualStatus1Code) Validate() error {
	for _, vv := range []string{
		"ACTC", "RJCT", "PDNG", "ACCP", "ACSP", "ACSC", "ACCR", "ACWC",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("TransactionIndividualStatus1Code")
}

// Must be at least 1 items long
type ExternalMandateSetupReason1Code string

func (r ExternalMandateSetupReason1Code) Validate() error {
	if len(string(r)) < 1 {
		return util.NewErrTextLengthInvalid("ExternalMandateSetupReason1Code", 1, 0)
	}
	return nil
}

// Must be at least 1 items long
type ExternalLocalInstrument1Code string

func (r ExternalLocalInstrument1Code) Validate() error {
	if len(string(r)) < 1 {
		return util.NewErrTextLengthInvalid("ExternalLocalInstrument1Code", 1, 0)
	}
	return nil
}

// May be one of LEGL, AGNT, CUST, ARDT, NOAS, NOOR, AC04, AM04
type PaymentCancellationRejection2Code string

func (r PaymentCancellationRejection2Code) Validate() error {
	for _, vv := range []string{
		"LEGL", "AGNT", "CUST", "ARDT", "NOAS", "NOOR", "AC04", "AM04",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("PaymentCancellationRejection2Code")
}

// Must be at least 1 items long
type ExternalCategoryPurpose1Code string

func (r ExternalCategoryPurpose1Code) Validate() error {
	if len(string(r)) < 1 {
		return util.NewErrTextLengthInvalid("ExternalCategoryPurpose1Code", 1, 0)
	}
	return nil
}

// Must be at least 1 items long
type ExternalChargeType1Code string

func (r ExternalChargeType1Code) Validate() error {
	if len(string(r)) < 1 {
		return util.NewErrTextLengthInvalid("ExternalChargeType1Code", 1, 0)
	}
	return nil
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

// Must be at least 1 items long
type ExternalCashClearingSystem1Code string

func (r ExternalCashClearingSystem1Code) Validate() error {
	if len(string(r)) < 1 {
		return util.NewErrTextLengthInvalid("ExternalCashClearingSystem1Code", 1, 0)
	}
	return nil
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

// May be one of CNCL, MODI, IPAY, ICOV, MCOV, INFO, CONF, CWFW, MWFW, UWFW, PECR, PDCR, RJCR, SMTC, SMTI, CHRG, PURP, IDUP
type InvestigationExecutionConfirmation3Code string

func (r InvestigationExecutionConfirmation3Code) Validate() error {
	for _, vv := range []string{
		"CNCL", "MODI", "IPAY", "ICOV", "MCOV", "INFO", "CONF", "CWFW", "MWFW",
		"UWFW", "PECR", "PDCR", "RJCR", "SMTC", "SMTI", "CHRG", "PURP", "IDUP",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("InvestigationExecutionConfirmation3Code")
}

// May be one of UM01, UM02, UM03, UM04, UM05, UM06, UM07, UM08, UM09, UM10, UM11, UM12, UM13, UM14, UM15, UM16, UM17, UM18, UM19, UM20, UM21, UM22, UM23, UM24, UM25, UM26, UM27
type ModificationRejection2Code string

func (r ModificationRejection2Code) Validate() error {
	for _, vv := range []string{
		"UM01", "UM02", "UM03", "UM04", "UM05", "UM06", "UM07", "UM08", "UM09", "UM10",
		"UM11", "UM12", "UM13", "UM14", "UM15", "UM16", "UM17", "UM18", "UM19", "UM20",
		"UM21", "UM22", "UM23", "UM24", "UM25", "UM26", "UM27",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("ModificationRejection2Code")
}
