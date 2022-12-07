// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v01

import (
	"encoding/xml"

	"apitest/models/nfps/iso20022/common"
	"apitest/util"
)

type AccountSwitchDetails1 struct {
	UnqRefNb      common.Max35Text            `xml:"UnqRefNb"`
	RtgUnqRefNb   common.Max35Text            `xml:"RtgUnqRefNb"`
	SwtchRcvdDtTm *common.ISODateTime         `xml:"SwtchRcvdDtTm,omitempty" json:",omitempty"`
	SwtchDt       common.ISODate              `xml:"SwtchDt,omitempty" json:",omitempty"`
	SwtchTp       SwitchType1Code             `xml:"SwtchTp"`
	SwtchSts      *SwitchStatus1Code          `xml:"SwtchSts,omitempty" json:",omitempty"`
	BalTrfWndw    *BalanceTransferWindow1Code `xml:"BalTrfWndw,omitempty" json:",omitempty"`
	Rspn          []ResponseDetails1          `xml:"Rspn,omitempty" json:",omitempty"`
}

func (r AccountSwitchDetails1) Validate() error {
	return util.Validate(&r)
}

type AccountSwitchTerminationSwitchV01 struct {
	XMLName       xml.Name               `xml:"AcctSwtchTermntnSwtch"`
	MsgId         MessageIdentification1 `xml:"MsgId"`
	AcctSwtchDtls AccountSwitchDetails1  `xml:"AcctSwtchDtls"`
	SplmtryData   []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountSwitchTerminationSwitchV01) Validate() error {
	return util.Validate(&r)
}

type MessageIdentification1 struct {
	Id      common.Max35Text   `xml:"Id"`
	CreDtTm common.ISODateTime `xml:"CreDtTm"`
}

func (r MessageIdentification1) Validate() error {
	return util.Validate(&r)
}

type ResponseDetails1 struct {
	RspnCd    common.Max35Text   `xml:"RspnCd"`
	AddtlDtls *common.Max350Text `xml:"AddtlDtls,omitempty" json:",omitempty"`
}

func (r ResponseDetails1) Validate() error {
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
