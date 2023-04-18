// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v01

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
type ExternalMarketInfrastructure1Code string

func (r ExternalMarketInfrastructure1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 3 {
		return util.NewErrTextLengthInvalid("ExternalFinancialInstitutionIdentification1Code", 1, 3)
	}
	return nil
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
type ExternalSystemEventType1Code string

func (r ExternalSystemEventType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return util.NewErrTextLengthInvalid("ExternalSystemEventType1Code", 1, 4)
	}
	return nil
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

// May be one of LQMG, LMMG, PYMG, REDR, BKMG, STMG
type PaymentRole1Code string

func (r PaymentRole1Code) Validate() error {
	for _, vv := range []string{
		"LQMG", "LMMG", "PYMG", "REDR", "BKMG", "STMG",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return util.NewErrValueInvalid("PaymentRole1Code")
}
