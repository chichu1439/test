// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package head_v02

import (
	"reflect"

	"apitest/util"
)

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

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

func (r ExternalClearingSystemIdentification1Code) Validate() error {
	if len(string(r)) < 1 {
		return util.NewErrTextLengthInvalid("ExternalClearingSystemIdentification1Code", 1, 0)
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
type ExternalOrganisationIdentification1Code string

func (r ExternalOrganisationIdentification1Code) Validate() error {
	if len(string(r)) < 1 {
		return util.NewErrTextLengthInvalid("ExternalOrganisationIdentification1Code", 1, 0)
	}
	return nil
}

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

func (r ExternalPersonIdentification1Code) Validate() error {
	if len(string(r)) < 1 {
		return util.NewErrTextLengthInvalid("ExternalPersonIdentification1Code", 1, 0)
	}
	return nil
}
