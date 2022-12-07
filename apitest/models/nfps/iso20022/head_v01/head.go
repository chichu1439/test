// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package head_v01

import (
	"encoding/xml"
	"time"

	"apitest/models/nfps/iso20022/common"
	"apitest/util"
)

type BusinessApplicationHeaderV01 struct {
	XMLName    xml.Name                     `xml:"AppHdr"`
	Attrs      []xml.Attr                   `xml:",any,attr,omitempty" json:",omitempty"`
	CharSet    string                       `xml:"CharSet,omitempty" json:",omitempty"`
	Fr         Party9Choice                 `xml:"Fr"`
	To         Party9Choice                 `xml:"To"`
	BizMsgIdr  common.Max35Text             `xml:"BizMsgIdr"`
	MsgDefIdr  common.Max35Text             `xml:"MsgDefIdr"`
	BizSvc     *common.Max35Text            `xml:"BizSvc,omitempty" json:",omitempty"`
	CreDt      common.ISONormalisedDateTime `xml:"CreDt"`
	CpyDplct   *common.CopyDuplicate1Code   `xml:"CpyDplct,omitempty" json:",omitempty"`
	PssblDplct bool                         `xml:"PssblDplct,omitempty" json:",omitempty"`
	Prty       string                       `xml:"Prty,omitempty" json:",omitempty"`
	Sgntr      *SignatureEnvelope           `xml:"Sgntr,omitempty" json:",omitempty"`
	Rltd       *BusinessApplicationHeader1  `xml:"Rltd,omitempty" json:",omitempty"`
}

func (r BusinessApplicationHeaderV01) Validate() error {
	return util.Validate(&r)
}

type BranchAndFinancialInstitutionIdentification5 struct {
	FinInstnId FinancialInstitutionIdentification8 `xml:"FinInstnId"`
	BrnchId    *BranchData2                        `xml:"BrnchId,omitempty" json:",omitempty"`
}

func (r BranchAndFinancialInstitutionIdentification5) Validate() error {
	return util.Validate(&r)
}

type BranchData2 struct {
	Id      *common.Max35Text  `xml:"Id,omitempty" json:",omitempty"`
	Nm      *common.Max140Text `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr *PostalAddress6    `xml:"PstlAdr,omitempty" json:",omitempty"`
}

func (r BranchData2) Validate() error {
	return util.Validate(&r)
}

type BusinessApplicationHeader1 struct {
	CharSet    string                       `xml:"CharSet,omitempty" json:",omitempty"`
	Fr         Party9Choice                 `xml:"Fr"`
	To         Party9Choice                 `xml:"To"`
	BizMsgIdr  common.Max35Text             `xml:"BizMsgIdr"`
	MsgDefIdr  common.Max35Text             `xml:"MsgDefIdr"`
	BizSvc     *common.Max35Text            `xml:"BizSvc,omitempty" json:",omitempty"`
	CreDt      common.ISONormalisedDateTime `xml:"CreDt"`
	CpyDplct   *common.CopyDuplicate1Code   `xml:"CpyDplct,omitempty" json:",omitempty"`
	PssblDplct bool                         `xml:"PssblDplct,omitempty" json:",omitempty"`
	Prty       string                       `xml:"Prty,omitempty" json:",omitempty"`
	Sgntr      *SignatureEnvelope           `xml:"Sgntr,omitempty" json:",omitempty"`
}

func (r BusinessApplicationHeader1) Validate() error {
	return util.Validate(&r)
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                          `xml:"Prtry"`
}

func (r ClearingSystemIdentification2Choice) Validate() error {
	return util.Validate(&r)
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty" json:",omitempty"`
	MmbId    common.Max35Text                     `xml:"MmbId"`
}

func (r ClearingSystemMemberIdentification2) Validate() error {
	return util.Validate(&r)
}

type ContactDetails2 struct {
	NmPrfx   *common.NamePrefix1Code `xml:"NmPrfx,omitempty" json:",omitempty"`
	Nm       *common.Max140Text      `xml:"Nm,omitempty" json:",omitempty"`
	PhneNb   *common.PhoneNumber     `xml:"PhneNb,omitempty" json:",omitempty"`
	MobNb    *common.PhoneNumber     `xml:"MobNb,omitempty" json:",omitempty"`
	FaxNb    *common.PhoneNumber     `xml:"FaxNb,omitempty" json:",omitempty"`
	EmailAdr *common.Max2048Text     `xml:"EmailAdr,omitempty" json:",omitempty"`
	Othr     *common.Max35Text       `xml:"Othr,omitempty" json:",omitempty"`
}

func (r ContactDetails2) Validate() error {
	return util.Validate(&r)
}

type DateAndPlaceOfBirth struct {
	BirthDt     common.ISODate     `xml:"BirthDt"`
	PrvcOfBirth *common.Max35Text  `xml:"PrvcOfBirth,omitempty" json:",omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth"`
}

func (r DateAndPlaceOfBirth) Validate() error {
	return util.Validate(&r)
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                                `xml:"Prtry"`
}

func (r FinancialIdentificationSchemeName1Choice) Validate() error {
	return util.Validate(&r)
}

type FinancialInstitutionIdentification8 struct {
	BICFI       *common.BICFIIdentifier              `xml:"BICFI,omitempty" json:",omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:",omitempty"`
	Nm          *common.Max140Text                   `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr     *PostalAddress6                      `xml:"PstlAdr,omitempty" json:",omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr,omitempty" json:",omitempty"`
}

func (r FinancialInstitutionIdentification8) Validate() error {
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

type OrganisationIdentification7 struct {
	AnyBIC *common.AnyBICIdentifier             `xml:"AnyBIC,omitempty" json:",omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

func (r OrganisationIdentification7) Validate() error {
	return util.Validate(&r)
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                        `xml:"Prtry"`
}

func (r OrganisationIdentificationSchemeName1Choice) Validate() error {
	return util.Validate(&r)
}

type Party10Choice struct {
	OrgId  OrganisationIdentification7 `xml:"OrgId"`
	PrvtId PersonIdentification5       `xml:"PrvtId"`
}

func (r Party10Choice) Validate() error {
	return util.Validate(&r)
}

type Party9Choice struct {
	OrgId *PartyIdentification42                       `xml:"OrgId,omitempty" json:",omitempty"`
	FIId  BranchAndFinancialInstitutionIdentification5 `xml:"FIId"`
}

func (r Party9Choice) Validate() error {
	return util.Validate(&r)
}

type PartyIdentification42 struct {
	Nm        *common.Max140Text  `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr   *PostalAddress6     `xml:"PstlAdr,omitempty" json:",omitempty"`
	Id        *Party10Choice      `xml:"Id,omitempty" json:",omitempty"`
	CtryOfRes *common.CountryCode `xml:"CtryOfRes,omitempty" json:",omitempty"`
	CtctDtls  *ContactDetails2    `xml:"CtctDtls,omitempty" json:",omitempty"`
}

func (r PartyIdentification42) Validate() error {
	return util.Validate(&r)
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth           `xml:"DtAndPlcOfBirth,omitempty" json:",omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

func (r PersonIdentification5) Validate() error {
	return util.Validate(&r)
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

func (r PersonIdentificationSchemeName1Choice) Validate() error {
	return util.Validate(&r)
}

type PostalAddress6 struct {
	AdrTp       *common.AddressType2Code `xml:"AdrTp,omitempty" json:",omitempty"`
	Dept        *common.Max70Text        `xml:"Dept,omitempty" json:",omitempty"`
	SubDept     *common.Max70Text        `xml:"SubDept,omitempty" json:",omitempty"`
	StrtNm      *common.Max70Text        `xml:"StrtNm,omitempty" json:",omitempty"`
	BldgNb      *common.Max16Text        `xml:"BldgNb,omitempty" json:",omitempty"`
	PstCd       *common.Max16Text        `xml:"PstCd,omitempty" json:",omitempty"`
	TwnNm       *common.Max35Text        `xml:"TwnNm,omitempty" json:",omitempty"`
	CtrySubDvsn *common.Max35Text        `xml:"CtrySubDvsn,omitempty" json:",omitempty"`
	Ctry        *common.CountryCode      `xml:"Ctry,omitempty" json:",omitempty"`
	AdrLine     []common.Max70Text       `xml:"AdrLine,omitempty" json:",omitempty"`
}

func (r PostalAddress6) Validate() error {
	return util.Validate(&r)
}

type SignatureEnvelope struct {
	Item string `xml:",innerxml"`
}

func (r SignatureEnvelope) Validate() error {
	return util.Validate(&r)
}

type Head001 struct {
	FromMemberId   string    `xml:"Fr>FIId>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"required,max=35"`
	FromBranchId   string    `xml:"Fr>BrnchId>Id,omitempty" validate:"max=35"`
	ToMemberId     string    `xml:"To>FIId>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"required,max=35"`
	ToBranchId     string    `xml:"To>BrnchId>Id,omitempty" validate:"max=35"`
	BizMsgId       string    `xml:"BizMsgIdr,omitempty" validate:"required,max=35"`
	MsgDefId       string    `xml:"MsgDefIdr,omitempty" validate:"required,max=35"`
	BizSvc         string    `xml:"BizSvc,omitempty" validate:"required,max=35"`
	CreateDatetime time.Time `xml:"CreDt,omitempty" validate:"required,max=35"`
	CopyDup        string    `xml:"CpyDplct,omitempty" validate:"max=35"`
}

func NewBusinessApplicationHeaderV01(head *Head001) *BusinessApplicationHeaderV01 {
	return &BusinessApplicationHeaderV01{
		Fr: Party9Choice{
			FIId: BranchAndFinancialInstitutionIdentification5{
				FinInstnId: FinancialInstitutionIdentification8{
					BICFI: nil,
					ClrSysMmbId: &ClearingSystemMemberIdentification2{
						MmbId: common.Max35Text(head.FromMemberId),
					},
				},
				BrnchId: &BranchData2{
					Id: (*common.Max35Text)(&head.FromBranchId),
				},
			},
		},
		To: Party9Choice{
			FIId: BranchAndFinancialInstitutionIdentification5{
				FinInstnId: FinancialInstitutionIdentification8{
					BICFI: nil,
					ClrSysMmbId: &ClearingSystemMemberIdentification2{
						MmbId: common.Max35Text(head.ToMemberId),
					},
				},
				BrnchId: &BranchData2{
					Id: (*common.Max35Text)(&head.ToBranchId),
				},
			},
		},
		BizMsgIdr: common.Max35Text(head.BizMsgId),
		MsgDefIdr: common.Max35Text(head.MsgDefId),
		BizSvc:    (*common.Max35Text)(&head.BizSvc),
		CreDt:     common.ISONormalisedDateTime(head.CreateDatetime),
		CpyDplct:  (*common.CopyDuplicate1Code)(&head.CopyDup),
	}
}
