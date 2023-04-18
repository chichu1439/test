// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package auth_v02

import (
	"reflect"
	"regexp"

	"isaving-swagger/util"
)

// May be one of WIBO, TREA, TIBO, TLBO, SWAP, STBO, PRBO, PFAN, NIBO, MAAA, MOSP, LIBO, LIBI, JIBA, ISDA, GCFR, FUSW, EUCH, EUUS, EURI, EONS, EONA, CIBO, CDOR, BUBO, BBSW
type BenchmarkCurveName2Code string

func (r BenchmarkCurveName2Code) Validate() error {
	for _, vv := range []string{
		"WIBO", "TREA", "TIBO", "TLBO", "SWAP", "STBO", "PRBO", "PFAN", "NIBO", "MAAA", "MOSP", "LIBO", "LIBI",
		"JIBA", "ISDA", "GCFR", "FUSW", "EUCH", "EUUS", "EURI", "EONS", "EONA", "CIBO", "CDOR", "BUBO", "BBSW",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("BenchmarkCurveName2Code")
}

// May be one of CRDT, DBIT
type CreditDebit3Code string

func (r CreditDebit3Code) Validate() error {
	for _, vv := range []string{
		"CRDT", "DBIT",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("CreditDebit3Code")
}

// May be one of FITE, CALL
type DepositType1Code string

func (r DepositType1Code) Validate() error {
	for _, vv := range []string{
		"FITE", "CALL",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("DepositType1Code")
}

// Must match the pattern [0-9]
type Exact1NumericText string

func (r Exact1NumericText) Validate() error {
	reg := regexp.MustCompile(`[0-9]`)
	if !reg.MatchString(string(r)) {
		return util.NewErrValueInvalid("Exact1NumericText")
	}
	return nil
}

// Must match the pattern [0-9]{5}
type Exact5NumericText string

func (r Exact5NumericText) Validate() error {
	reg := regexp.MustCompile(`[0-9]{5}`)
	if !reg.MatchString(string(r)) {
		return util.NewErrValueInvalid("Exact5NumericText")
	}
	return nil
}

// May be one of SPOT, SALE, AGRD
type ExchangeRateType1Code string

func (r ExchangeRateType1Code) Validate() error {
	for _, vv := range []string{
		"SPOT", "SALE", "AGRD",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("ExchangeRateType1Code")
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
type ExternalClearingSystemIdentification1Code string

func (r ExternalClearingSystemIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 5 {
		return util.NewErrTextLengthInvalid("ExternalClearingSystemIdentification1Code", 1, 5)
	}
	return nil
}

// Must be at least 1 items long
type ExternalContractBalanceType1Code string

func (r ExternalContractBalanceType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalContractBalanceType1Code", 1, 4)
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

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

func (r ISINOct2015Identifier) Validate() error {
	reg := regexp.MustCompile(`[A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}`)
	if !reg.MatchString(string(r)) {
		return util.NewErrValueInvalid("ISINOct2015Identifier")
	}
	return nil
}

// May be one of CNTR, ESTM
type PaymentScheduleType1Code string

func (r PaymentScheduleType1Code) Validate() error {
	for _, vv := range []string{
		"CNTR", "ESTM",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("PaymentScheduleType1Code")
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

// May be one of DAYS, MNTH, WEEK, YEAR
type RateBasis1Code string

func (r RateBasis1Code) Validate() error {
	for _, vv := range []string{
		"DAYS", "MNTH", "WEEK", "YEAR",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("RateBasis1Code")
}

// May be one of EMAL, FAXI, FILE, ONLI, PHON, POST, PROP, SWMT, SWMX
type CommunicationMethod4Code string

func (r CommunicationMethod4Code) Validate() error {
	for _, vv := range []string{
		"EMAL", "FAXI", "FILE", "ONLI", "PHON", "POST", "PROP", "SWMT", "SWMX",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("CommunicationMethod4Code")
}

// May be one of NONE, MASA, MISA, SISA, IISA, CUYP, PRYP, ASTR, EMPY, EMCY, EPRY, ECYE, NFPI, NFQP, DECP, IRAC, IRAR, KEOG, PFSP, 401K, SIRA, 403B, 457X, RIRA, RIAN, RCRF, RCIP, EIFP, EIOP
type TaxExemptReason1Code string

func (r TaxExemptReason1Code) Validate() error {
	for _, vv := range []string{
		"NONE", "MASA", "MISA", "SISA", "IISA", "CUYP", "PRYP", "ASTR", "EMPY", "EMCY", "EPRY", "ECYE", "NFPI",
		"NFQP", "DECP", "IRAC", "IRAR", "KEOG", "PFSP", "401K", "SIRA", "403B", "457X", "RIRA", "RIAN", "RCRF",
		"RCIP", "EIFP", "EIOP",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("TaxExemptReason1Code")
}

// Must be at least 1 items long
type ExternalContractClosureReason1Code string

func (r ExternalContractClosureReason1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalContractClosureReason1Code", 1, 4)
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
type ExternalProxyAccountType1Code string

func (r ExternalProxyAccountType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalProxyAccountType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalShipmentCondition1Code string

func (r ExternalShipmentCondition1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalShipmentCondition1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalValidationRuleIdentification1Code string

func (r ExternalValidationRuleIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalValidationRuleIdentification1Code", 1, 4)
	}
	return nil
}

// May be one of ALLL, CHNG, MODF
type QueryType3Code string

func (r QueryType3Code) Validate() error {
	for _, vv := range []string{
		"ALLL", "CHNG", "MODF",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("QueryType3Code")
}

// May be one of LFBK, LTBK, SUPP
type SupportDocumentType1Code string

func (r SupportDocumentType1Code) Validate() error {
	for _, vv := range []string{
		"LFBK", "LTBK", "SUPP",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("SupportDocumentType1Code")
}

// Must be at least 1 items long
type ExternalStatusReason1Code string

func (r ExternalStatusReason1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalStatusReason1Code", 1, 4)
	}
	return nil
}

// May be one of ACPT, ACTC, PART, PDNG, RCVD, RJCT, RMDR, INCF, CRPT
type StatisticalReportingStatus1Code string

func (r StatisticalReportingStatus1Code) Validate() error {
	for _, vv := range []string{
		"ACPT", "ACTC", "PART", "PDNG", "RCVD", "RJCT", "RMDR", "INCF", "CRPT",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("StatisticalReportingStatus1Code")
}
