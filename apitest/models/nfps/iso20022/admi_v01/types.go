// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package admi_v01

import (
	"reflect"

	"apitest/util"
)

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

func (r ExternalAccountIdentification1Code) Validate() error {
	if len(string(r)) < 1 {
		return util.NewErrTextLengthInvalid("ExternalAccountIdentification1Code", 1, 0)
	}
	return nil
}

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

func (r ExternalClearingSystemIdentification1Code) Validate() error {
	if len(string(r)) < 1 {
		return util.NewErrTextLengthInvalid("ExternalClearingSystemIdentification1Code", 1, 0)
	}
	return nil
}

// Must be at least 1 items long
type ExternalEnquiryRequestType1Code string

func (r ExternalEnquiryRequestType1Code) Validate() error {
	if len(string(r)) < 1 {
		return util.NewErrTextLengthInvalid("ExternalEnquiryRequestType1Code", 1, 0)
	}
	return nil
}

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

func (r ExternalFinancialInstitutionIdentification1Code) Validate() error {
	if len(string(r)) < 1 {
		return util.NewErrTextLengthInvalid("ExternalFinancialInstitutionIdentification1Code", 1, 0)
	}
	return nil
}

// Must be at least 1 items long
type ExternalPaymentControlRequestType1Code string

func (r ExternalPaymentControlRequestType1Code) Validate() error {
	if len(string(r)) < 1 {
		return util.NewErrTextLengthInvalid("ExternalPaymentControlRequestType1Code", 1, 0)
	}
	return nil
}

// Must be at least 1 items long
type ExternalSystemBalanceType1Code string

func (r ExternalSystemBalanceType1Code) Validate() error {
	if len(string(r)) < 1 {
		return util.NewErrTextLengthInvalid("ExternalSystemBalanceType1Code", 1, 0)
	}
	return nil
}

// Must be at least 1 items long
type ExternalSystemEventType1Code string

func (r ExternalSystemEventType1Code) Validate() error {
	if len(string(r)) < 1 {
		return util.NewErrTextLengthInvalid("ExternalSystemEventType1Code", 1, 0)
	}
	return nil
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
